# readonly

`readonly` is a Go linter that enforces the immutability of struct fields, preventing modifications from outside their package.

This tool offers a simpler alternative to getter methods.

## Installation

```go install github.com/devinalvaro/readonly/cmd/readonly@latest```

## Usage

```readonly [package]```

## Example

Consider the following struct:

```go
package example

type Config struct {
	Name  string `readonly:"enforce"`
	Value string
}
```

`readonly` will report an error if `Config.Name` is modified from outside the `example` package.

Another example:

```go
package example

type Config struct {
	Name  string `readonly:"enforce_all"`
	Value string
}
```

`readonly` now enforces both `Config.Name` and `Config.Value`.

## Testing

The [test file](test/testdata/src/noflags/main.go) shows all cases that `readonly` covers.
