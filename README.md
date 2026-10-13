# go-diff

File diff utility implemented in Go.

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
| `make test RUN=…` | Run only tests matching a regex, e.g. `RUN=TestLineOperationString` |
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

```bash
go-diff [--lcs] FILE1 FILE2
```

- The LCS algorithm is the default until Myers is implemented. After that,
  Myers becomes the default and `--lcs` selects LCS.

## License

See [LICENSE](LICENSE).
