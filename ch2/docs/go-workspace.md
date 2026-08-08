# Go workspace and module guide

This note explains the difference between older Go workspace layout and the modern module system.

## 1. Older Go workspace style

Before Go modules became standard, Go projects were commonly organized around a `GOPATH` workspace.
A typical structure looked like this:

```text
GOPATH/
  bin/
  pkg/
  src/
```

- `src/` contained source code
- `pkg/` contained compiled package files
- `bin/` contained installed commands

This style is still useful for understanding how Go builds and installs programs.

## 2. Modern Go modules

Today, the recommended approach is to use modules.
A module is started with:

```bash
go mod init example.com/myapp
```

This creates a `go.mod` file and makes the folder a proper Go module.

### Why modules matter

Modules make dependency management easier.
They help Go know which version of each dependency to use.

## 3. `go install` explained

`go install` is used to build and install a package or binary.
It is different from `go run`:

- `go run` : executes the code immediately
- `go build` : creates a binary file
- `go install` : places the binary into the Go bin directory

Example:

```bash
go install ./...
```

If the current directory contains a `main` package, the result can be executed from your shell.

## 4. Useful environment commands

```bash
go env GOPATH
go env GOBIN
go env GOMOD
```

These commands help you see where Go stores installed binaries and where the current module is located.

## 5. Simple summary

- Use `go mod init` to start a module.
- Use `go run` to test code quickly.
- Use `go install` to install a program.
- Use `go work` when working with multiple modules together.

If you are learning Go from scratch, focus on modules first because they are the standard approach today.
