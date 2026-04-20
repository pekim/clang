package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

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
	scalar, isScalar := scalarTypes[kind]
	if !isScalar {
		fatalf("unsupported integer type for enum : %s", kind)
	}

	name := exportedGoName(cursor.Spelling())
	if cursor.Spelling() == "CX_BinaryOperatorKind" {
		// avoid a name clash between CXBinaryOperatorKind and CX_BinaryOperatorKind.
		name = "BinaryOperatorKind_"
	}

	enum := enum{
		name:    name,
		typ:     scalar.code,
		comment: commentText(cursor.ParsedComment()),
	}

	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_EnumConstantDecl {
			return clang.ChildVisit_Continue
		}

		enum.members = append(enum.members, enumMember{
			name:    exportedGoName(cursor.Spelling()),
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
