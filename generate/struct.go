package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type structField struct {
	name          string
	comment       string
	scalar        scalar
	isScalar      bool
	isPointer     bool
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
		name:    exportedGoName(cursor.Spelling()),
		comment: commentText(cursor.ParsedComment()),
	}

	unnamed := 1
	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_FieldDecl {
			return clang.ChildVisit_Continue
		}

		isPointer := strings.Contains(cursor.Type().Spelling(), "*")

		name := exportedGoName(cursor.Spelling())
		if name == "" {
			name = fmt.Sprintf("_%d", unnamed)
			unnamed++
		}
		if isPointer {
			name = goName(cursor.Spelling())
		}

		field := structField{
			name:          name,
			comment:       commentText(cursor.ParsedComment()),
			isPointer:     isPointer,
			_tempSpelling: cursor.Type().Spelling(),
		}
		field.scalar, field.isScalar = scalarTypes[cursor.Type().Kind()]

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
			g.Id("_").Qual("structs", "HostLayout")
			g.Line()

			for _, field := range struct_.fields {
				if field.isScalar {
					g.Comment(field.comment)
					g.Id(field.name).Add(field.scalar.code)
				} else if field.isPointer {
					g.Comment(field.comment)
					g.Id(field.name).Uintptr().Comment(field._tempSpelling)
				} else {
					g.Commentf("%s => %s", field.name, field._tempSpelling)
				}
			}
		})
	}
}
