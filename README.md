# think-money-technical-test

## Setup
Either run the following commands:
- Manual:
    - `go mod download`
- Automatic: `make setup`

## Usage
| Command | Description | Flags |
| --- | --- |
| `make setup` | Installs all dependencies for the code | |
| `make run_tests` | Runs all tests within the code repository | |
| `make run_test_coverage` | Runs all tests, creates test_coverage.out in the root of the project, opens your browser of choice with results of coverage | |
| `make run_benchmarks` | Runs defined benchmarks | |
| `make swag_generate` | Generates the swagger files for the passed main file (relational to cmd/...) | name={folder name of where you want to build swagger files for} |

## Installed packages
- [testify](https://github.com/stretchr/testify)