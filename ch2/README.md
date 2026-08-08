# Chapter 2: Go workspace and modules

This chapter introduces the structure behind Go projects and explains the modern module workflow that is used in most current Go applications.

## Learning objectives

By the end of this chapter, you should understand:

- how Go projects are organized
- the difference between older workspace style and modern modules
- how to initialize a module with `go mod init`
- how `go install` differs from `go run` and `go build`

## Key concepts

### Go workspace

Older Go projects often used a workspace layout based on `GOPATH` with folders such as:

- `src/` for source files
- `pkg/` for compiled packages
- `bin/` for installed programs

This is useful for understanding how Go historically arranged its files.

### Go modules

Modern Go uses modules. A module is a folder containing a `go.mod` file.
You can create one with:

```bash
go mod init example.com/myapp
```

This tells Go that the folder is a module and allows dependencies to be managed in a consistent way.

### Go install

`go install` builds and installs a package or program into the Go binary directory.

```bash
go install ./...
```

Use it when you want a program to be available from your shell.

## Useful commands

```bash
# create a new module
go mod init example.com/myapp

# run a program
go run .

# build a binary
go build

# install a program
go install ./...

# inspect Go paths
go env GOPATH
go env GOBIN
go env GOMOD
```

## Files in this chapter

- `gowork/` : example folder showing the older workspace-style idea
- `test/` : place to try commands safely
- `docs/` : notes about modules and workspace concepts

## Summary

- Use `go mod init` to begin a module-based project.
- Use `go run` for quick testing.
- Use `go build` to create a binary.
- Use `go install` to place a tool into your bin directory.
