package generate

import (
	"unsafe"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

var scalarTypes = map[clang.TypeKind]struct {
	code jen.Code
	size int
}{
	clang.Type_Int: {
		code: jen.Int32(),
		size: int(unsafe.Sizeof(*new(int32))),
	},
	clang.Type_UInt: {
		code: jen.Uint32(),
		size: int(unsafe.Sizeof(*new(uint32))),
	},
	clang.Type_ULong: {
		code: jen.Uint64(),
		size: int(unsafe.Sizeof(*new(uint64))),
	},
}
