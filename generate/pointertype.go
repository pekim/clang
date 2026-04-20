package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type pointerType struct {
	name    string
	comment string
}

type pointerTypes struct {
	types []pointerType
}

func (types *pointerTypes) add(cursor clang.Cursor) {
	types.types = append(types.types, pointerType{
		name:    goName(cursor.Spelling()),
		comment: commentText(cursor.ParsedComment()),
	})
}

func (types *pointerTypes) generate() {
	file := newFile("clang", ".", "pointertype")
	defer file.save()

	for _, typ := range types.types {
		file.Comment(typ.comment)
		file.Type().Id(typ.name).Struct(
			jen.Id("c").Uintptr(),
		)
		file.Line()
	}
}
