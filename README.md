# gocrazy

> A collection of (not so) crazy little Golang utils that I use in my projects.

## Install

```bash
go get github.com/danawoodman/gocrazy
```

## Usage

### Env helpers

#### `gocrazy.Getenv(key, fallback)`

Get an environment variable or fallback to a default value.

```go
import "github.com/danawoodman/gocrazy"

var port = gocrazy.Getenv("PORT", "3000")
```

### Url helpers

#### `gocrazy.AppendQueryParams(url, params)`

Append query params to a url.

```go
import "github.com/danawoodman/gocrazy"

params := struct {
	Foo string `url:"foo"` // use url tag to specify param name
	Bar string `json:"bar"` // fallback to json tag
	Baz string `url:"-"` // omit
	Bang string // fallback to field name "Bang"
}{
	Foo: "foo",
	Bar: "bar",
	Baz: "baz",
	Bang: "bang",
}
var url = gocrazy.AppendQueryParams("https://example.com", params)
```

### File paths

#### `gocrazy.ExpandHome(path)`

Expand a path that starts with `~` or `$HOME` to the current user's home directory.

```go
import "github.com/danawoodman/gocrazy"

var path = gocrazy.ExpandHome("~/foo/bar")
```

## Development

Run tests in watch mode (using [gochange](https://github.com/danawoodman/gochange)) using `make` or `make dev`.

Run test suite once with `make test`.

## License

[MIT](./LICENSE.md)
