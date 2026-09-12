# Contributing

Thanks for helping improve these Go lessons.

## Before opening a change

1. Keep examples small and focused on one concept.
2. Use lowercase, descriptive filenames with underscores when needed.
3. Format changed Go files with `gofmt`.
4. Compile the changed example with `go test path/to/file.go`.
5. Update the nearest README when a lesson or filename changes.

## Commit guidance

Use a short, descriptive commit message, for example:

```text
feat(ch9): explain slice capacity growth
```

Avoid committing generated binaries, `bin/`, `pkg/`, editor settings, or local
workspace files.
