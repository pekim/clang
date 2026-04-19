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
	name    string
	typ     jen.Code
	comment string
	members []enumMember
}

type enumMember struct {
	name    string
	value   int
	comment string
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

	enum := enum{
		name:    goName(cursor.Spelling()),
		typ:     typ,
		comment: commentText(cursor.ParsedComment()),
	}

	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_EnumConstantDecl {
			return clang.ChildVisit_Continue
		}

		enum.members = append(enum.members, enumMember{
			name:    goName(cursor.Spelling()),
			value:   int(cursor.EnumConstantDeclValue()),
			comment: commentText(cursor.ParsedComment()),
		})

		return clang.ChildVisit_Continue
	})

	enums.enums = append(enums.enums, enum)
}

func (enums *enums) generate() {
	file := newFile("clang", ".", "enum")
	defer file.save()

	for _, enum := range enums.enums {
		file.Comment(enum.comment)
		file.Type().Id(enum.name).Add(enum.typ)
		file.Line()

		file.Const().DefsFunc(func(g *jen.Group) {
			for _, member := range enum.members {
				g.Comment(member.comment)
				g.Id(member.name).Id(enum.name).Op("=").Lit(member.value)
			}
		})
		file.Line()
	}
}
