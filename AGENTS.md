# Repository Guidelines

## Project Structure & Module Organization

- `internal/lib/eventstore/` contains the core domain interfaces and repository logic for the event store.
- `go.mod` defines the module (`yoshiyoshifujii/go-eventstore`) and Go version.
- There are no binaries or entrypoints yet (no `cmd/`), so usage is intended as a library.

## Build, Test, and Development Commands

- `go test ./...` runs all package tests (and builds packages as part of the test step).
- `go test ./internal/lib/eventstore -run TestName` runs a targeted test.
- `go vet ./...` performs static analysis on the codebase.
- `go list ./...` verifies module graph and package discovery.

## Coding Style & Naming Conventions

- Indentation: Go standard tabs; follow `gofmt` output for spacing and alignment.
- Naming: Go conventions (CamelCase for exported types like `EventStore`, lowerCamelCase for locals).
- Interfaces live in `internal/lib/eventstore` and follow clear domain naming: `Aggregate`, `Command`, `Event`.
- Prefer small, focused methods with explicit error returns; avoid panics except for invariant checks.

## Testing Guidelines

- Use Go’s standard testing package; `github.com/stretchr/testify` is available for assertions if desired.
- Place tests alongside code as `*_test.go` within the same package.
- Name tests with `TestXxx` and keep setup minimal; favor table-driven tests for variations.
- There is no stated coverage target; add tests for new behavior and edge cases.

## Commit & Pull Request Guidelines

- History is minimal (`first commit`, `add README`), so no strict convention is established.
- Use short, imperative commit messages (e.g., “add repository tests”).
- PRs should include a brief description, rationale for changes, and notes on tests run.
- Link related issues if applicable; include examples or usage notes when changing public interfaces.
