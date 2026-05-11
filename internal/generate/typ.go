package generate

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

type scalar struct {
	code           jen.Code
	size           int
	typeDescriptor string
}

var scalarTypes = map[clang.TypeKind]*scalar{
	clang.Type_Double: {
		code:           jen.Float64(),
		size:           int(unsafe.Sizeof(*new(float64))),
		typeDescriptor: "DoubleTypeDescriptor",
	},
	clang.Type_Int: {
		code:           jen.Int32(),
		size:           int(unsafe.Sizeof(*new(int32))),
		typeDescriptor: "SInt32TypeDescriptor",
	},
	clang.Type_LongLong: {
		code:           jen.Int64(),
		size:           int(unsafe.Sizeof(*new(int64))),
		typeDescriptor: "SInt64TypeDescriptor",
	},
	clang.Type_UInt: {
		code:           jen.Uint32(),
		size:           int(unsafe.Sizeof(*new(uint32))),
		typeDescriptor: "UInt32TypeDescriptor",
	},
	clang.Type_UChar: {
		code:           jen.Byte(),
		size:           int(unsafe.Sizeof(*new(byte))),
		typeDescriptor: "UInt8TypeDescriptor",
	},
	clang.Type_ULong: {
		code:           jen.Uint64(),
		size:           int(unsafe.Sizeof(*new(uint64))),
		typeDescriptor: "UInt64TypeDescriptor",
	},
	clang.Type_ULongLong: {
		code:           jen.Uint64(),
		size:           int(unsafe.Sizeof(*new(uint64))),
		typeDescriptor: "UInt64TypeDescriptor",
	},
}

type typ struct {
	description        string
	size               int
	enum               *enum
	enumPointer        *enum
	pointerType        *pointerType
	pointerTypePointer *pointerType
	scalar             *scalar
	struct_            *struct_
	structPointer      *struct_
	isScalarPointer    bool
	isString           bool
	isStringPointer    bool
	isPointer          bool
	isPointerPointer   bool
	isVoid             bool
}

func newTyp(clangType clang.Type, gen *gen) typ {
	isPointer := clangType.Kind == clang.Type_Pointer
	isPointerPointer := isPointer && clangType.PointeeType().Kind == clang.Type_Pointer

	typ := typ{
		description:      clangType.TypeSpelling(),
		size:             int(clangType.SizeOf()),
		isPointer:        isPointer,
		isPointerPointer: isPointerPointer,
		isVoid:           clangType.Kind == clang.Type_Void,
	}
	typ.enum = gen.enums.find(strings.TrimPrefix(clangType.TypeSpelling(), "enum "))
	typ.pointerType = gen.pointerTypes.find(clangType.TypeSpelling())
	typ.scalar = scalarTypes[clangType.Kind]
	typ.struct_ = gen.structs.find(clangType.TypeSpelling())
	if typ.isPointer {
		typ.scalar = scalarTypes[clangType.PointeeType().Kind]
		typ.isVoid = clangType.PointeeType().Kind == clang.Type_Void

		pointeeSpelling := clangType.PointeeType().TypeSpelling()
		pointeeSpelling = strings.TrimPrefix(pointeeSpelling, "const ")
		pointeeSpelling = strings.TrimPrefix(pointeeSpelling, "enum ")
		pointeeSpelling = strings.TrimPrefix(pointeeSpelling, "struct ")
		typ.enumPointer = gen.enums.find(pointeeSpelling)
		typ.pointerTypePointer = gen.pointerTypes.find(pointeeSpelling)
		typ.structPointer = gen.structs.find(pointeeSpelling)
	}
	typ.isScalarPointer = typ.isPointer && typ.isScalar()
	typ.isString = isPointer && clangType.PointeeType().Kind == clang.Type_Char_S
	typ.isStringPointer = isPointerPointer && clangType.PointeeType().PointeeType().Kind == clang.Type_Char_S

	return typ
}

func (typ typ) isEnum() bool {
	return typ.enum != nil
}
func (typ typ) isEnumPointer() bool {
	return typ.enumPointer != nil
}
func (typ typ) isPointerType() bool {
	return typ.pointerType != nil
}
func (typ typ) isPointerTypePointer() bool {
	return typ.pointerTypePointer != nil
}
func (typ typ) isScalar() bool {
	return typ.scalar != nil
}
func (typ typ) isStruct() bool {
	return typ.struct_ != nil
}
func (typ typ) isStructPointer() bool {
	return typ.structPointer != nil
}

func (typ typ) goffiTypeDescriptor() (*jen.Statement, bool) {
	if typ.isVoid {
		if typ.isPointer {
			return jen.Qual(typesPackage, "PointerTypeDescriptor"), true
		}
		return jen.Qual(typesPackage, "VoidTypeDescriptor"), true
	}
	if typ.isEnum() {
		return jen.Qual(typesPackage, typ.enum.scalar.typeDescriptor), true
	}
	if typ.isPointer {
		return jen.Qual(typesPackage, "PointerTypeDescriptor"), true
	}
	if typ.isPointerType() {
		return jen.Qual(typesPackage, "PointerTypeDescriptor"), true
	}
	if typ.isScalar() {
		return jen.Qual(typesPackage, typ.scalar.typeDescriptor), true
	}
	if typ.isStruct() {
		return jen.Id(typ.struct_.typeDescriptorName), true
	}

	// panic(fmt.Sprintf("unsupported typ : %s", typ))
	return jen.Null(), false
}

func (typ typ) String() string {
	return typ.description
}

func (typ typ) cDecl() jen.Code {
	if typ.isVoid {
		return jen.Null()
	}
	if typ.isString {
		return jen.Qual("unsafe", "Pointer")
	}
	if typ.isEnum() {
		return jen.Id(typ.enum.goName)
	}
	if typ.isPointer {
		return jen.Qual("unsafe", "Pointer")
	}
	if typ.isPointerType() {
		return jen.Id(typ.pointerType.goName)
	}
	if typ.isScalar() {
		return typ.scalar.code
	}
	if typ.isStruct() {
		return jen.Id(typ.struct_.goName)
	}

	panic(fmt.Sprintf("unsupported typ : %s", typ))
}

func (typ typ) goDecl() jen.Code {
	if typ.isVoid {
		if typ.isPointer {
			return jen.Qual("unsafe", "Pointer")
		}
		return jen.Null()
	}
	if typ.isString {
		return jen.String()
	}
	if typ.isEnum() {
		return jen.Id(typ.enum.goName)
	}
	if typ.isEnumPointer() {
		return jen.Op("*").Id(typ.enumPointer.goName)
	}
	if typ.isPointerType() {
		return jen.Id(typ.pointerType.goName)
	}
	if typ.isPointerTypePointer() {
		return jen.Op("*").Id(typ.pointerTypePointer.goName)
	}
	if typ.isScalar() {
		if typ.isPointer {
			return jen.Op("*").Add(typ.scalar.code)
		}
		return typ.scalar.code
	}
	if typ.isStruct() {
		return jen.Id(typ.struct_.goName)
	}
	if typ.isStructPointer() {
		return jen.Op("*").Id(typ.structPointer.goName)
	}
	if typ.isPointer {
		return jen.Qual("unsafe", "Pointer")
	}

	panic(fmt.Sprintf("unsupported typ : %s", typ))
}

func (typ typ) goOutReturnDecl() jen.Code {
	if typ.isEnumPointer() {
		return jen.Id(typ.enumPointer.goName)
	}
	if typ.isPointerTypePointer() {
		return jen.Id(typ.pointerTypePointer.goName)
	}
	if typ.isScalarPointer {
		return jen.Add(typ.scalar.code)
	}
	if typ.isCXStringPointer() {
		return jen.String()
	}
	if typ.isStructPointer() {
		return jen.Id(typ.structPointer.goName)
	}

	panic(fmt.Sprintf("unsupported typ : %s", typ))
}

func (typ typ) goOutVarDecl() jen.Code {
	if typ.isEnumPointer() {
		return jen.Id(typ.enumPointer.goName)
	}
	if typ.isPointerTypePointer() {
		return jen.Id(typ.pointerTypePointer.goName)
	}
	if typ.isScalarPointer {
		return jen.Add(typ.scalar.code)
	}
	if typ.isCXStringPointer() {
		return jen.Id(typ.structPointer.goName)
	}
	if typ.isStructPointer() {
		return jen.Id(typ.structPointer.goName)
	}

	panic(fmt.Sprintf("unsupported typ : %s", typ))
}

func (typ typ) isCXString() bool {
	return typ.isStruct() && typ.struct_.cName == "CXString"
}

func (typ typ) isCXStringPointer() bool {
	return typ.isStructPointer() && typ.structPointer.cName == "CXString"
}
