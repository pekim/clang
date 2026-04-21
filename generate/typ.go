package generate

import (
	"unsafe"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type scalar struct {
	code jen.Code
	size int
}

var scalarTypes = map[clang.TypeKind]scalar{
	clang.Type_Int: {
		code: jen.Int32(),
		size: int(unsafe.Sizeof(*new(int32))),
	},
	clang.Type_UInt: {
		code: jen.Uint32(),
		size: int(unsafe.Sizeof(*new(uint32))),
	},
	clang.Type_UChar: {
		code: jen.Byte(),
		size: int(unsafe.Sizeof(*new(byte))),
	},
	clang.Type_ULong: {
		code: jen.Uint64(),
		size: int(unsafe.Sizeof(*new(uint64))),
	},
}
