package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

var kindTypes = map[clang.TypeKind]jen.Code{
	clang.Type_Int:  jen.Int32(),
	clang.Type_UInt: jen.Uint32(),
}

type enum struct {
	name       string
	typ        jen.Code
	commentDoc string
}

type enums struct {
	enums []enum
}

func (enums *enums) add(cursor clang.Cursor) {
	kind := cursor.EnumDeclIntegerType().Kind()
	typ, ok := kindTypes[kind]
	if !ok {
		fatalf("unsupported integer type for enum : %s", kind)
	}

	enums.enums = append(enums.enums, enum{
		name:       cursor.Spelling(),
		typ:        typ,
		commentDoc: commentText(cursor.ParsedComment()),
	})
}

func (enums *enums) generate() {
	file := newFile("clang", ".", "enum")
	defer file.save()

	for _, enum := range enums.enums {
		file.Comment(enum.commentDoc)
		file.Type().Id(enum.name).Add(enum.typ)
		file.Line()
	}
}
