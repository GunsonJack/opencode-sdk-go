# Maintaining the opencode-sdk-go fork

This document describes how the fork is structured and how to keep it healthy.

## Vendored OpenAPI spec

The canonical API specification lives at:

```
openapi/openapi.json
```

All mock-server tests and any future code-generation tooling should reference this file.
When the upstream API changes, update this file and adjust the SDK code to match.

## Running the mock server

The repository includes a helper script that starts [Prism](https://github.com/stoplightio/prism)
in daemon mode against the vendored spec:

```sh
./scripts/mock --daemon
```

This is used by the test suite to validate requests and responses against the spec.

## Validating the SDK locally

Run the full local validation pipeline:

```sh
./scripts/bootstrap   # install dependencies
./scripts/lint        # run linters
go test ./...         # run all tests (requires the mock server)
```

You can also run `./scripts/test` which combines the mock server startup and test execution.

## Compatibility wrappers

When the upstream API deprecates or removes a route:

1. **Prefer additive wrappers.** Add new types/methods alongside existing ones rather than
   replacing them, so that consumers are not broken by a minor release.
2. **Document deprecated routes.** Mark old methods with a `// Deprecated:` comment and
   note the replacement in the godoc string.
3. **Remove in a major version.** Only remove deprecated symbols in a new major version
   bump, giving consumers time to migrate.

## General guidelines

- This fork is **manually maintained** — there is no automated code generator.
- All changes go through normal pull requests and code review.
- Keep the vendored spec (`openapi/openapi.json`) as the single source of truth for the API surface.
