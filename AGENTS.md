# AGENTS.md

go-diff is a line-based file diff CLI in Go. Only `internal/diff` holds code so far; `cmd/go-diff`, `internal/cli` and `internal/render` are placeholders.

## Run Go commands inside the devcontainer

The host has no Go toolchain. Go, `make` and `golangci-lint` live in the devcontainer defined by `.devcontainer/devcontainer.json`, which mounts the repo at `/workspaces/go-diff`, so host edits are visible in the container immediately. Run every `go` and `make` command through `docker exec`, from the repo root:

```bash
docker exec -u vscode -w /workspaces/go-diff \
  "$(docker ps -q --filter "label=devcontainer.config_file=$PWD/.devcontainer/devcontainer.json")" \
  make check
```

Swap `make check` for any other command. The `vscode` user has the host's UID 1000, so files it writes stay owned by the host user.

If `docker ps` finds no container, the devcontainer is not running: ask the user to start it (VS Code: **Dev Containers: Reopen in Container**). Where the devcontainer CLI is installed, `devcontainer up --workspace-folder .` starts it and `devcontainer exec --workspace-folder . make check` replaces the `docker exec` form.

## Done means `make check` is clean

CI runs `make check` (fmt, vet, race-enabled tests) and then `git diff --exit-code`. A change is done when `make check` passes in the container and `go fmt` leaves the tree unchanged.

- One test: `make test RUN=TestLCS`
- Fuzzing: `go test -fuzz=FuzzLCS -fuzztime=30s ./internal/diff` (without `-fuzztime` it runs until stopped)
- Lint: `make lint`

## Code conventions

- `internal/diff/edit` holds the edit-script types (`Node`, `Op`) shared by the diff steps and the renderers. Display symbols such as `+` and `-` belong in `internal/render`.
- `internal/diff` keeps tokenization private. Its public entry point is `LCS(oldLines, newLines []string) Script`, which tokenizes both inputs through one shared `tokenMap` so equal tokens always mean equal lines. New algorithms (Myers is planned) follow the same shape: take `[]string`, return a `Script`.
- Tests are table-driven: a package-level `var <name>Tests = []struct{...}` looped with `t.Run`. Algorithm tests work on tokens in `package diff`; public-API tests live in `package diff_test` and compare rendered text, never token ids.
