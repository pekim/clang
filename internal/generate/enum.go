package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

type enum struct {
	cName   string
	goName  string
	scalar  scalar
	comment string
	members []enumMember
}

type enumMember struct {
	name    string
	value   int
	comment string
}

type enums []enum

func (gen *gen) addEnum(cursor clang.Cursor) {
	kind := cursor.EnumDeclIntegerType().Kind
	scalar, isScalar := scalarTypes[kind]
	if !isScalar {
		fatalf("unsupported integer type for enum : %s", kind.TypeKindSpelling())
	}

	name := exportedGoName(cursor.CursorSpelling())
	if cursor.CursorSpelling() == "CX_BinaryOperatorKind" {
		// avoid a name clash between CXBinaryOperatorKind and CX_BinaryOperatorKind.
		name = "BinaryOperatorKind_"
	}

	enum := enum{
		cName:   cursor.CursorSpelling(),
		goName:  name,
		scalar:  *scalar,
		comment: commentText(cursor.ParsedComment()),
	}

	cursor.VisitChildren(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind != clang.Cursor_EnumConstantDecl {
			return clang.ChildVisit_Continue
		}

		enum.members = append(enum.members, enumMember{
			name:    exportedGoName(cursor.CursorSpelling()),
			value:   int(cursor.EnumConstantDeclValue()),
			comment: commentText(cursor.ParsedComment()),
		})

		return clang.ChildVisit_Continue
	})

	gen.enums = append(gen.enums, enum)
}

func (enums enums) generate() {
	file := newFile("clang", ".", "enum")
	defer file.save()

	for _, enum := range enums {
		file.Comment(enum.comment)
		file.Commentf("Represents the C enum %s.\n", enum.cName)

		file.Type().Id(enum.goName).Add(enum.scalar.code)
		file.Line()

		file.Const().DefsFunc(func(g *jen.Group) {
			for _, member := range enum.members {
				g.Comment(member.comment)
				g.Id(member.name).Id(enum.goName).Op("=").Lit(member.value)
			}
		})
		file.Line()
	}
}

func (enums enums) find(cName string) *enum {
	for _, enum := range enums {
		if enum.cName == cName {
			return &enum
		}
	}
	return nil
}
