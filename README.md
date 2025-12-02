[![Actions Status](https://github.com/spoddub/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/spoddub/go-project-244/actions)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=spoddub_go-project-244&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=spoddub_go-project-244)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=spoddub_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=spoddub_go-project-244)

## GenDiff

`GenDiff` is a small training CLI utility written in Go that compares two configuration files (JSON or YAML) and shows the difference between them.  
It supports nested structures and multiple output formats:

- **stylish** (default) - human-friendly tree diff
- **plain** - line-based, machine-readable text diff
- **json** - structured JSON representation of the diff

---

## Tools used

| Tool                                                                          | What it is used for                                             |
| ----------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| [Go](https://go.dev/)                                                         | Language and toolchain.                                         |
| [urfave/cli v3](https://github.com/urfave/cli)                                | Building the CLI interface and flags.                           |
| [golangci-lint](https://golangci-lint.run/)                                   | Fast all-in-one Go linter.                                      |
| [testify](https://github.com/stretchr/testify)                                | Assertions in unit tests.                                       |
| [GitHub Actions](https://docs.github.com/actions)                             | CI for linting, tests and quality checks.                       |
| [Make](https://www.gnu.org/software/make/)                                    | Simple developer tasks: build, test, lint, coverage.            |
| [SonarCloud](https://sonarcloud.io/)                                          | Code quality and coverage metrics.                              |

---

## Installation and local development

### Requirements

- Go (modern version, e.g. 1.21 or newer)
- `make`
- `golangci-lint` installed locally (or via `go install`)

### Clone the repository

```bash
git clone https://github.com/spoddub/GenDiff.git
cd GenDiff
```

### Build

```bash
make build
# builds ./bin/gendiff
```

---

## Usage

Basic examples assuming you are in the project root and built the binary with `make build`:

```bash
# Build the binary
make build

# Default output (stylish) for JSON files
./bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json

# Default output (stylish) for YAML files
./bin/gendiff testdata/fixture/file1.yml testdata/fixture/file2.yml

# Plain format
./bin/gendiff --format plain testdata/fixture/file1.json testdata/fixture/file2.json
# or
./bin/gendiff -f plain testdata/fixture/file1.yml testdata/fixture/file2.yml

# JSON format
./bin/gendiff --format json testdata/fixture/file1.json testdata/fixture/file2.json
```

### CLI help

```text
$ gendiff -h

NAME:
   gendiff - Compares two configuration files and shows a difference.

USAGE:
   gendiff [global options]

GLOBAL OPTIONS:
   --format string, -f string  output format (default: "stylish")
   --help, -h                  show help
```

---

## Asciinema demos

### Stylish format, JSON input

[![asciicast](https://asciinema.org/a/ED1LgcMKgIjeZNoC0kSr4mdvX.svg)](https://asciinema.org/a/ED1LgcMKgIjeZNoC0kSr4mdvX)

### Stylish format, YAML input

[![asciicast](https://asciinema.org/a/lPdBOanNE9YdF9XPowSVOzSml.svg)](https://asciinema.org/a/lPdBOanNE9YdF9XPowSVOzSml)

---

## Tests and quality

Run unit tests:

```bash
make test
# or directly:
# go test ./...
```

Run linters:

```bash
make lint
# or directly:
# golangci-lint run ./...
```