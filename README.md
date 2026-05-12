# clang

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/clang)](https://pkg.go.dev/github.com/pekim/clang)
[![golangci-lint](https://github.com/pekim/clang/actions/workflows/ci.yml/badge.svg)](https://github.com/pekim/clang/actions/workflows/ci.yml)

Generated Go bindings for libclang's
[C API](https://clang.llvm.org/doxygen/group__CINDEX.html).

The bindings are pure Go, using no cgo.
This is achieved using
[goffi](https://github.com/go-webgpu/goffi).

## API

Most, but not all, of the C API is provided.
It's enough to at least be able to parse header files,
walk the AST, and extract some details.

The [documentation](https://pkg.go.dev/github.com/pekim/clang)
for the Go bindings on pkg.go.dev comes from the Doxygen comments in
libclang's header files,
so it reflects the C API.
Refer to the
[C API](https://clang.llvm.org/doxygen/group__CINDEX.html)
document for the definitive word on the C API.

Most output parameters do not appear as parameters in the Go API,
but instead are return values.

The `CX` prefix is removed from C type names,
and the `clang_` prefix is removed from function names.

Many functions are provided as struct types' methods.

## Requirements

The `libclang` library must be available to dynamically load.

## Platforms

The only architecture supported is `amd64`.
This is because the callback parameter for
[Cursor.VisitChildren](http://localhost:8080/github.com/pekim/clang#Cursor.VisitChildren)
has struct parameters,
and `goffi` currently only supports callback struct arguments on `amd64`.

CI runs tests on `linux/amd64` and `macos/amd64`

## License

clang is licensed under the terms of the MIT license.

## AI use

No AI was used in the creation of this library.

## Development

To regenerate the bindings run `go generate`.
This generates code & runs tests.

### pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
