# go-diff

File diff utility implemented in Go.

## Project layout

```
go-diff/
├── .devcontainer/          # VS Code / Dev Containers configuration (Go 1.27)
├── cmd/
│   └── go-diff/            # package main — entry point; keep it thin
├── internal/
│   ├── cli/                # flag parsing, argument validation, wiring
│   ├── diff/               # core diff algorithm (e.g. Myers / LCS), no I/O
│   │   └── testdata/       # fixture files used by tests (ignored by go build)
│   └── render/             # output formats (unified, side-by-side, ...)
├── bin/                    # build output (git-ignored)
├── Makefile
├── go.mod                  # create with `go mod init` (see below)
└── README.md
```

Conventions followed:

- **`cmd/<name>/`** holds `main` packages. `main.go` should do little more than
  call into `internal/cli` and translate the returned error into an exit code.
- **`internal/`** packages cannot be imported by other modules, so the API can
  change freely. Promote a package to `pkg/` (or the module root) only if you
  decide to publish it as a library.
- **`testdata/`** directories are ignored by the Go toolchain and are the
  idiomatic home for test fixtures and golden files.
- Tests live next to the code they test (`foo.go` → `foo_test.go`).
- Empty directories contain a `.gitkeep` so Git tracks them; delete it once the
  directory has real files.

## Development with the devcontainer

The repo ships a [Dev Container](https://containers.dev/) based on
`mcr.microsoft.com/devcontainers/go:2-1.27-trixie`. It provides Go 1.27, `make`,
`git`, `gopls` and related Go tools, plus zsh / Oh My Zsh, running as the
non-root `vscode` user. Nothing needs to be installed on the host other than
Docker.

### VS Code

1. Install [Docker](https://docs.docker.com/get-docker/) (Docker Desktop on
   Windows/macOS; with WSL2, enable the WSL integration) and the
   [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
   extension.
2. Open this folder in VS Code.
3. Run **Dev Containers: Reopen in Container** from the command palette
   (`Ctrl+Shift+P` / `Cmd+Shift+P`), or click **Reopen in Container** on the
   prompt that appears.
4. The first build takes a few minutes. Afterwards, open a terminal in VS Code —
   it runs inside the container with the repo mounted at
   `/workspaces/go-diff`.

After editing `.devcontainer/devcontainer.json`, run
**Dev Containers: Rebuild Container** to apply the changes.

### Dev Container CLI

Without VS Code, use the [devcontainer CLI](https://github.com/devcontainers/cli):

```bash
npm install -g @devcontainers/cli
```

```bash
devcontainer up --workspace-folder .
```

```bash
devcontainer exec --workspace-folder . make check
```

## Building and testing

All common tasks are in the `Makefile`. Run `make` or `make help` to list them.

| Target            | Description                                             |
| ----------------- | ------------------------------------------------------- |
| `make build`      | Build `bin/go-diff`                                     |
| `make install`    | Install `go-diff` into `$GOBIN`                         |
| `make run ARGS=…` | Build and run, e.g. `make run ARGS="a.txt b.txt"`       |
| `make test`       | Run all tests with the race detector                    |
| `make cover`      | Run tests with coverage and print a summary             |
| `make cover-html` | Open the HTML coverage report                           |
| `make bench`      | Run benchmarks                                          |
| `make fmt`        | `go fmt ./...`                                          |
| `make vet`        | `go vet ./...`                                          |
| `make lint`       | Run `golangci-lint` (must be installed separately)      |
| `make tidy`       | `go mod tidy` and `go mod verify`                       |
| `make check`      | `fmt` + `vet` + `test` — run before committing          |
| `make clean`      | Remove `bin/`, coverage output and the test cache       |

`make build` stamps the binary with a version from `git describe`. To use it,
declare `var version = "dev"` in package `main`.

## Usage

_TODO_

## License

See [LICENSE](LICENSE).
