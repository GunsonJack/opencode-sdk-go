## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/build
```

This will install all the required dependencies and build the SDK.

You can also [install go 1.22+ manually](https://go.dev/doc/install).

## About this fork

This is a manually maintained fork. It does **not** rely on the Stainless code generator.
All code changes are made by hand and reviewed through normal pull requests. For details
on how the fork is maintained, see [`docs/MAINTAINING.md`](./docs/MAINTAINING.md).

## Modifying/Adding code

Because this fork is manually maintained, you can freely edit any file in the repository.
Please follow Go conventions and ensure `./scripts/lint` passes before submitting a PR.

## Adding and running examples

All files in the `examples/` directory can be freely edited or added to.

```go
# add an example to examples/<your-example>/main.go

package main

func main() {
  // ...
}
```

```sh
$ go run ./examples/<your-example>
```

## Using the repository from source

To use a local version of this library from source in another project, edit the `go.mod` with a replace
directive. This can be done through the CLI with the following:

```sh
$ go mod edit -replace github.com/GunsonJack/opencode-sdk-go=/path/to/opencode-sdk-go
```

## Running tests

Most tests require you to [set up a mock server](https://github.com/stoplightio/prism) against the vendored OpenAPI spec to run the tests.

```sh
# you will need npm installed
$ npx prism mock openapi/openapi.json
```

```sh
$ ./scripts/test
```

## Formatting

This library uses the standard gofmt code formatter:

```sh
$ ./scripts/format
```
