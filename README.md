# wtfunc

Find what Go function is on a given line of a file. This is intended to be a building block for other tools, e.g. an editor plugin that runs the test at the current cursor position.

## Installation

```
go get -u github.com/jim/wtfunc
```

## Usage

### Find function on a specific line using `--line`

Pass the path to a valid Go file on the command line or the contents of such a file via standard in. Pass the line number to check for a function body with `-line`. This program uses Go's standard flag parsing and requires that the `-line` flag comes before the filename.

```
$ wtfunc -line 21 main.go
run
```

### Get the list of all functions in a file using `-list`:

```
$ wtfunc -list main.go
main
run
```

### Limit results to test funcs using `-test`

This works with either `-list` or `-line` and includes only functions whose names begin with "Test":

```
$ wtfunc -test -list main_test.go
TestCLI
```

```
$ wtfunc -test -line 25 main_test.go
TestCLI
```
