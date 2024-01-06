# gocrazy

> A collection of (not so) crazy little Golang utils that I use in my projects.

## Install

```bash
go get github.com/danawoodman/gocrazy
```

## Usage

```go
import "github.com/danawoodman/gocrazy"

var port = gocrazy.Getenv("PORT", "3000")
```

## Development

Run tests in watch mode (using [gochange](https://github.com/danawoodman/gochange)) using `make` or `make dev`.

Run test suite once with `make test`.

## License

[MIT](./LICENSE.md)
