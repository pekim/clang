package generate

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type structField struct {
	name          string
	comment       string
	_tempSpelling string
}

type struct_ struct {
	name    string
	comment string
	fields  []structField
}

type structs struct {
	structs_ []struct_
}

func (ss *structs) add(cursor clang.Cursor) {
	// Structs with a name that start with a "_" do not appear to be referenced
	// anywhere in the API. So do not generate them.
	if strings.HasPrefix(cursor.Spelling(), "_") {
		return
	}

	struct_ := struct_{
		name:    goName(cursor.Spelling()),
		comment: commentText(cursor.ParsedComment()),
	}

	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_FieldDecl {
			return clang.ChildVisit_Continue
		}

		field := structField{
			name:          goName(cursor.Spelling()),
			comment:       commentText(cursor.ParsedComment()),
			_tempSpelling: cursor.Type().Spelling(),
		}
		struct_.fields = append(struct_.fields, field)

		return clang.ChildVisit_Continue
	})

	ss.structs_ = append(ss.structs_, struct_)
}

func (ss *structs) generate() {
	file := newFile("clang", ".", "struct")
	defer file.save()

	for _, struct_ := range ss.structs_ {
		file.Comment(struct_.comment)
		file.Type().Id(struct_.name).StructFunc(func(g *jen.Group) {
			for _, field := range struct_.fields {
				g.Commentf("%s => %s", field.name, field._tempSpelling)
			}
		})
	}
}
