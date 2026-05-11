package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

type structField struct {
	typ
	goName       string
	comment      string
	alignment    int
	offset       int
	size         int
	typeSpelling string

	isBitfield            bool
	bitWidth              int
	bitfieldOffset        int
	bitfieldDataFieldName string
}

type struct_ struct {
	cursor             clang.Cursor
	cName              string
	goName             string
	typeDescriptorName string
	comment            string
	alignment          int
	size               int
	fields             []structField
}

type structs []struct_

func (gen *gen) addStruct(cursor clang.Cursor) {
	// Structs with a name that start with a "_" do not appear to be referenced
	// anywhere in the API. So do not generate them.
	if strings.HasPrefix(cursor.CursorSpelling(), "_") {
		return
	}

	struct_ := struct_{
		cursor:             cursor,
		cName:              cursor.CursorSpelling(),
		goName:             exportedGoName(cursor.CursorSpelling()),
		typeDescriptorName: goName(cursor.CursorSpelling()) + "TypeDescriptor",
		comment:            commentText(cursor.ParsedComment()),
		alignment:          int(cursor.CursorType().AlignOf()),
		size:               int(cursor.CursorType().SizeOf()),
	}
	if struct_.alignment < 1 {
		struct_.alignment = 8
	}

	gen.structs = append(gen.structs, struct_)
}

func (ss structs) enrich(gen *gen) {
	for i := range ss {
		(&ss[i]).enrich(gen)
	}
}

func (struct_ *struct_) enrich(gen *gen) {
	nameSuffix := 0
	bitfieldOffset := 0
	var prevField *structField

	struct_.cursor.VisitChildren(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind != clang.Cursor_FieldDecl {
			return clang.ChildVisit_Continue
		}

		typ := newTyp(cursor.CursorType(), gen)

		name := exportedGoName(cursor.CursorSpelling())
		if name == "" && cursor.IsBitField() == 0 {
			name = fmt.Sprintf("_%d", nameSuffix)
			nameSuffix++
		}
		if typ.isCXString() {
			name = goName(cursor.CursorSpelling())
		}
		if typ.isPointer {
			name = goName(cursor.CursorSpelling())
		}

		field := structField{
			typ:          typ,
			goName:       name,
			comment:      commentText(cursor.ParsedComment()),
			alignment:    int(cursor.CursorType().AlignOf()),
			offset:       int(cursor.OffsetOfField()),
			size:         int(cursor.CursorType().SizeOf()),
			typeSpelling: cursor.CursorType().TypeSpelling(),
		}

		field.isBitfield = cursor.IsBitField() != 0
		field.bitWidth = int(cursor.FieldDeclBitWidth())
		if field.isBitfield {
			field.bitfieldDataFieldName = fmt.Sprintf("bitfield_%d", nameSuffix)
			field.bitfieldOffset = bitfieldOffset
			bitfieldOffset += field.bitWidth
		} else if prevField != nil && prevField.isBitfield {
			nameSuffix++
		}

		struct_.fields = append(struct_.fields, field)
		prevField = &field

		return clang.ChildVisit_Continue
	})
}

func (struct_ struct_) generate(file file) {
	struct_.generateStruct(file)
	struct_.generateAccessors(file)
}

func (struct_ struct_) generateStruct(file file) {
	file.Comment(struct_.comment)
	file.Commentf("Represents the C struct %s.\n", struct_.cName)

	file.Type().Id(struct_.goName).StructFunc(func(g *jen.Group) {
		g.Id("_").Qual("structs", "HostLayout")
		g.Line()

		var firstBitfield *structField
		var prevField *structField
		for _, field := range struct_.fields {
			struct_.generateFieldPadding(g, field.offset, prevField)

			if !field.isBitfield && prevField != nil && prevField.isBitfield {
				// generate a field to hold one or more bitfields
				if (field.offset-firstBitfield.offset)/8 != 2 {
					// for now, the only bitfields total 16 bits
					fatalf("expected bitfields for %s.%s to be 16 bits", struct_.goName, firstBitfield.goName)
				}
				g.Id(prevField.bitfieldDataFieldName).Uint16()
			}

			if !field.isBitfield {
				g.Comment(field.comment)

				if field.isEnum() {
					g.Id(field.goName).Id(field.enum.goName)
				} else if field.isScalar() {
					g.Id(field.goName).Add(field.scalar.code)
				} else if field.isStruct() {
					g.Id(field.goName).Id(field.struct_.goName)
				} else if field.isPointer {
					g.Id(field.goName).Qual("unsafe", "Pointer").Comment(field.typeSpelling)
				} else {
					g.Id(field.goName).Index(jen.Lit(field.size)).Byte().Comment(field.typeSpelling)
				}

				firstBitfield = nil
			} else {
				if firstBitfield == nil {
					firstBitfield = &field
				}
			}

			prevField = &field
		}

		struct_.generateFieldPadding(g, struct_.size*8, prevField)
	})
}

func (struct_ struct_) generateAccessors(file file) {
	for _, field := range struct_.fields {
		if field.isBitfield {
			struct_.generateBitfieldAccessors(file, field)
		}
		if field.isCXString() {
			struct_.generateCXStringAccessors(file, field)
		}
		if field.isString {
			struct_.generateStringAccessors(file, field)
		}
	}
}

func (struct_ struct_) generateBitfieldAccessors(file file, field structField) {
	if field.goName == "" {
		return
	}

	// getter
	file.Comment(field.comment)
	file.
		Func().
		Params(jen.Id("s").Op("*").Id(struct_.goName)).
		Id(field.goName).Params().Uint().Block(
		jen.Return().Id("bitfieldGet").Call(
			jen.Id("s").Dot(field.bitfieldDataFieldName),
			jen.Lit(field.bitfieldOffset),
			jen.Lit(field.bitWidth),
		),
	)
	file.Line()

	// setter
	file.Comment(field.comment)
	file.
		Func().
		Params(jen.Id("s").Op("*").Id(struct_.goName)).
		Id("Set" + field.goName).Params(jen.Id("value").Uint()).Block(
		jen.Id("bitfieldSet").Call(
			jen.Id("s").Dot(field.bitfieldDataFieldName),
			jen.Lit(field.bitfieldOffset),
			jen.Lit(field.bitWidth),
			jen.Id("value"),
		),
	)
	file.Line()
}

func (struct_ struct_) generateCXStringAccessors(file file, field structField) {
	// getter
	file.Comment(field.comment)
	file.
		Func().
		Params(jen.Id("s").Op("*").Id(struct_.goName)).
		Id(exportedGoName(field.goName)).Params().String().Block(
		jen.Return().Id("s").Dot(field.goName).Dot("CString").Call(),
	)
	file.Line()
}

func (struct_ struct_) generateStringAccessors(file file, field structField) {
	// getter
	file.Comment(field.comment)
	file.
		Func().
		Params(jen.Id("s").Op("*").Id(struct_.goName)).
		Id(exportedGoName(field.goName)).Params().String().Block(
		jen.Return().Qual("github.com/pekim/clang/internal/libc", "GoString").Call(
			jen.Id("s").Dot(field.goName),
		),
	)
	file.Line()

	// setter
	file.Comment(field.comment)
	file.Comment(`
A C string will be allocated on the C heap.
A function is returned that when called will free the C string's memory.
`)
	file.
		Func().
		Params(jen.Id("s").Op("*").Id(struct_.goName)).
		Id("Set"+exportedGoName(field.goName)).
		Params(jen.Id("str").String()).
		Parens(jen.Id("free").Func().Params()).
		Block(
			jen.
				List(jen.Id("cString"), jen.Id("free")).
				Op(":=").
				Qual("github.com/pekim/clang/internal/libc", "CString").Call(jen.Id("str")),
			jen.
				Id("s").Dot(field.goName).
				Op("=").
				Id("cString"),
			jen.Return().Id("free"),
		)
	file.Line()
}

func (struct_ struct_) generateFieldPadding(g *jen.Group, targetOffset int, prevField *structField) {
	if prevField == nil {
		return
	}

	prevFieldEndOffset := (prevField.offset / 8) + prevField.size
	padding := (targetOffset / 8) - prevFieldEndOffset
	if padding > 0 {
		g.Id("_").Index(jen.Lit(padding)).Byte().Comment("padding")
	}
}

func (ss structs) generate() {
	file := newFile("clang", ".", "struct")
	defer file.save()

	for _, struct_ := range ss {
		struct_.generate(file)
	}

	ss.generateTestVars()
	ss.generateTypeDescriptors()
}

func (ss structs) generateTestVars() {
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
			for _, struct_ := range ss {
				if struct_.size == int(clang.TypeLayoutError_Incomplete) {
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

func (ss structs) find(cName string) *struct_ {
	for _, struct_ := range ss {
		if struct_.cName == cName {
			return &struct_
		}
	}
	return nil
}
