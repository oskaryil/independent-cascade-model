# Independent Cascade Model in Go

![Go](https://github.com/github.com/oskaryil/independent-cascade-model/workflows/Go/badge.svg)

This repository contains code that implements the independent cascade model based on a temporal network.

## Installation

Make sure you have [Go](https://golang.org/) installed (minimum version 1.15).

## Running

The program is implemented and based on the formatting of the contents inside [android.csv](./android.csv), so if you want to use another data set, please adhere to that same formatting.

### Options for running the program:

#### Building from source

You can run the code with the command:

**When running it's necessary to specify a file path to the input data file**

```bash
$ go run main.go -f <relative-file-path>
```

To build a binary:
Run the following command from the root of the repository:

```bash
$ go build
```

#### Using the pre-built binary

```bash
$ ./icm -f <relative-file-path>
```

Example (run from the root of the repository):

```bash
$ ./icm -f ./android.csv
```

## Tests

**Prerequisites:** Go version >= v 1.15

Running tests: `$ make test` alternatively `$ go test ./...`

### Unit and integration test:

```bash
$ go test ./...
```

### Benchmark tests

```bash
$ go test -bench=.
```
