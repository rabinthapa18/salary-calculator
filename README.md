# Salary Calculator

This project is a salary calculator application written in Go. It includes functionality to process input data, perform calculations, and validate results.

## Directories

- **api/**: Contains the API implementation.
  - api.go: Defines the API endpoints and handlers.
- **cmd/**: Contains the command-line interface (CLI) and API entry points.
  - **api/**: Entry point for the API server.
    - main.go: Starts the API server.
  - **cli/**: Entry point for the CLI application.
    - main.go: Starts the CLI application.
- **pkg/**: Contains the core logic and utilities.
- **testCases/**: Contains the test cases and logic for running them.
- **tests/**: Contains unit tests for the core logic.

## Usage

### Running the API Server

To start the API server, run following command from root directory:

```
go run cmd/api/main.go
```

> The API server will start on port 8080. You can also access the static website at `http://localhost:8080`.

### Running the CLI Application

To start the CLI application,, run following command from root directory:

```
go run cmd/cli/main.go <input string>
```

> Replace `<input string>` with the input data. For example:

```
go run cmd/cli/main.go 1000 1300 1500 4 0 9 9 17 17 22 22 23
```

### Running the Test Cases

To run the test cases, navigate to the root folder of the project and run:

```
go run testCases/main.go
```

> More details on test cases can be found in the [README](./testCases/README.md) file.

### Running Unit Tests

To run the unit tests, navigate to the root folder of the project and run:

```
go test ./...
```
