package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

type pointerType struct {
	cName   string
	goName  string
	comment string
}

type pointerTypes []pointerType

func (gen *gen) addPointerType(cursor clang.Cursor) {
	gen.pointerTypes = append(gen.pointerTypes, pointerType{
		cName:   cursor.CursorSpelling(),
		goName:  exportedGoName(cursor.CursorSpelling()),
		comment: commentText(cursor.ParsedComment()),
	})
}

func (types pointerTypes) generate() {
	file := newFile("clang", ".", "pointertype")
	defer file.save()

	for _, typ := range types {
		file.Comment(typ.comment)
		file.Type().Id(typ.goName).Struct(
			jen.Id("ptr").Qual("unsafe", "Pointer"),
		)
		file.Line()
	}
}

func (types pointerTypes) find(cName string) *pointerType {
	for _, type_ := range types {
		if type_.cName == cName {
			return &type_
		}
	}
	return nil
}
