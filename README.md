# clang

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/clang)](https://pkg.go.dev/github.com/pekim/clang)
[![golangci-lint](https://github.com/pekim/clang/actions/workflows/ci.yml/badge.svg)](https://github.com/pekim/clang/actions/workflows/ci.yml)

Generated Go bindings for libclang's
[C API](https://clang.llvm.org/doxygen/group__CINDEX.html).

The bindings are pure Go, using no cgo.
This is achieved using
[goffi](https://github.com/go-webgpu/goffi).

## API

Most, but not all, of the C API is supported.
It's enough to at least be able to parse header files,
walk the AST, and extract some details.
This is evidenced by this repository's use of its own bindings
in its code generation code.

The [documentation](https://pkg.go.dev/github.com/pekim/clang)
for the Go bindings on pkg.go.dev comes from the Doxygen comments in
libclang's header files,
so it reflects the C API.
Refer to the
[C API](https://clang.llvm.org/doxygen/group__CINDEX.html)
documentation for the definitive word on the C API.

Most output parameters do not appear as function parameters in the Go API,
but instead as return values.

The `CX` prefix is removed from C type names,
and the `clang_` prefix is removed from function names.

Many C functions are provided as Go methods on struct types.

## Requirements

The `libclang` library must be available to dynamically load.
The default library paths are OS-specific.

| OS    | filename         |
| ----- | ---------------- |
| linux | `libclang.so`    |
| macos | `libclang.dylib` |

The default library paths may be overridden with an
argument to the
[Init](http://localhost:8080/github.com/pekim/clang#Init)
function.

## Platforms

The only architecture supported is `amd64`.
This is because the callback parameter for
[Cursor.VisitChildren](http://localhost:8080/github.com/pekim/clang#Cursor.VisitChildren)
has struct parameters,
and `goffi` currently only supports callback struct arguments on `amd64`.

The CI worklow runs tests on `linux/amd64` and `macos/amd64`

## Example

```go
package main

import (
	"fmt"

	"github.com/pekim/clang"
)

func main() {
	// Initialise the library.
	if err := clang.Init(nil); err != nil {
		panic(err)
	}

	// Parse a header file, creating a TranslationUnit.
	index := clang.CreateIndex(0, 1)
	parseArgs := []string{"-x", "c-header"}
	tu, errorCode := index.ParseTranslationUnit2("internal/clang-c/CXString.h", parseArgs, nil,
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
	)
	if errorCode != clang.Error_Success {
		panic(fmt.Sprintf("failed to create translation unit : %d", errorCode))
	}

	// Traverse the various declarations within the translation unit.
	tuCursor := tu.TranslationUnitCursor()
	tuCursor.VisitChildren(func(cursor clang.Cursor, _parent clang.Cursor) clang.ChildVisitResult {
		fmt.Println(cursor.CursorKind().CursorKindSpelling(), cursor.CursorSpelling())
		return clang.ChildVisit_Continue
	})
}
```

## License

clang is licensed under the terms of the MIT license.

## AI use

No AI was used in the creation of this library.

## Development

To regenerate the bindings run `go generate`.
This generates code and then runs tests.

### pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`

```

```
