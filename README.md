# octo-go-learning

Small, runnable Go lessons that grow from `hello world` to slices, pointers,
dates, parsing, and classic workspace layouts.

This is a hands-on notebook rather than a framework: each chapter keeps the
code intentionally small so you can read it, run it, change it, and observe
the result.

## Start here

### Requirements

- Go 1.21 or newer
- A terminal and editor

### Run a lesson

From the repository root:

```bash
go run ./ch1/hello.go
go run ./ch9/arrays_with_slice.go
```

Most examples are standalone programs. Examples that read from standard input
will pause for a value when you run them.

### Check your changes

```bash
gofmt -w path/to/example.go
go test path/to/example.go
```

The repository's GitHub Actions workflow checks formatting and compiles each
hand-written example independently.

## Learning path

| Chapter | Focus | Start with |
| --- | --- | --- |
| [1](ch1/README.md) | Program structure and strings | `ch1/hello.go` |
| [2](ch2/README.md) | Modules and historical workspaces | `ch2/gowork/` |
| [3](ch3/README.md) | Formatting and string length | `ch3/formatting.go` |
| [4](ch4/README.md) | Input and parsing | `ch4/parse_float.go` |
| [5](ch5/README.md) | Arithmetic and precision | `ch5/arithmetic.go` |
| [6](ch6/README.md) | Dates and time | `ch6/date_calculation.go` |
| [7](ch7/README.md) | Pointers | `ch7/pointer.go` |
| 8 | Arrays | `ch8/arrays.go` |
| 9 | Slices and capacity | `ch9/arrays_with_slice.go` |

After the chapters, use [`practice on go/`](practice%20on%20go/) for short
experiments and challenges. The [`docs/`](docs/) folder contains the study
guide and reference notes.

## A useful study loop

1. Read one example without running it.
2. Run it and compare the output with your prediction.
3. Change one value, statement, or type.
4. Run `gofmt`, then run the file again.
5. Write down the rule you discovered in your own words.

## Repository map

```text
ch1/ ... ch9/       Guided lessons
practice on go/     Extra exercises and experiments
docs/               Study guide and reference notes
notes/              Quick reminders
.github/workflows/  Automated formatting and compile checks
```

## Contributing

Small improvements are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for the
expected workflow and style.
