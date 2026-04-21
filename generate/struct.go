package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type structField struct {
	goName       string
	comment      string
	enum         *enum
	isEnum       bool
	scalar       scalar
	isScalar     bool
	isPointer    bool
	offset       int
	size         int
	typeSpelling string
}

type struct_ struct {
	cName   string
	goName  string
	comment string
	size    int
	fields  []structField
}

type structs struct {
	structs_ []struct_
}

func (ss *structs) add(cursor clang.Cursor, enums enums) {
	// Structs with a name that start with a "_" do not appear to be referenced
	// anywhere in the API. So do not generate them.
	if strings.HasPrefix(cursor.Spelling(), "_") {
		return
	}

	struct_ := struct_{
		cName:   cursor.Spelling(),
		goName:  exportedGoName(cursor.Spelling()),
		comment: commentText(cursor.ParsedComment()),
		size:    int(cursor.Type().SizeOf()),
	}

	unnamed := 1
	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_FieldDecl {
			return clang.ChildVisit_Continue
		}

		isPointer := cursor.Type().Kind() == clang.Type_Pointer

		name := exportedGoName(cursor.Spelling())
		if name == "" {
			name = fmt.Sprintf("_%d", unnamed)
			unnamed++
		}
		if isPointer {
			name = goName(cursor.Spelling())
		}

		field := structField{
			goName:       name,
			comment:      commentText(cursor.ParsedComment()),
			isPointer:    isPointer,
			offset:       int(cursor.OffsetOfField()),
			size:         int(cursor.Type().SizeOf()),
			typeSpelling: cursor.Type().Spelling(),
		}
		field.enum, field.isEnum = enums.find(cursor.Type().Spelling())
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
		file.Type().Id(struct_.goName).StructFunc(func(g *jen.Group) {
			g.Id("_").Qual("structs", "HostLayout")
			g.Line()

			var prevField *structField
			for _, field := range struct_.fields {
				generateFieldPadding(g, field.offset, prevField)

				g.Comment(field.comment)
				// g.Commentf("OFFSET, SIZE : %d, %d", field.offset, field.size)

				if field.isEnum {
					g.Id(field.goName).Id(field.enum.goName)
				} else if field.isScalar {
					g.Id(field.goName).Add(field.scalar.code)
				} else if field.isPointer {
					g.Id(field.goName).Uintptr().Comment(field.typeSpelling)
				} else {
					g.Id(field.goName).Index(jen.Lit(field.size)).Byte().Comment(field.typeSpelling)
				}

				prevField = &field
			}

			generateFieldPadding(g, struct_.size*8, prevField)
		})
	}

	ss.generateTestVars()
}

func generateFieldPadding(g *jen.Group, targetOffset int, prevField *structField) {
	if prevField == nil {
		return
	}

	prevFieldEndOffset := (prevField.offset / 8) + prevField.size
	padding := (targetOffset / 8) - prevFieldEndOffset
	if padding > 0 {
		g.Id("_").Index(jen.Lit(padding)).Byte()
	}
}

func (ss *structs) generateTestVars() {
	file := newFile("clang", ".", "struct_testvars")
	defer file.save()

	file.
		Var().Id("structTestVars").
		Op("=").
		Index().Struct(
		jen.Id("name").String(),
		jen.Id("cSize").Uintptr(),
		jen.Id("goSize").Uintptr(),
	).
		ValuesFunc(func(g *jen.Group) {
			for _, struct_ := range ss.structs_ {
				if struct_.size == clang.TypeLayoutError_Incomplete {
					continue
				}

				g.Line().Values(
					jen.Lit(struct_.cName),
					jen.Lit(struct_.size),
					jen.Qual("unsafe", "Sizeof").Call(jen.Op("*").New(jen.Id(struct_.goName))),
				)
			}
			g.Line()
		})
}
