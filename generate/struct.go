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
	offset       int
	size         int
	typeSpelling string

	isBitfield            bool
	bitWidth              int
	bitfieldOffset        int
	bitfieldDataFieldName string

	enum      *enum
	isEnum    bool
	scalar    scalar
	isScalar  bool
	struct_   *struct_
	isStruct  bool
	isPointer bool
}

type struct_ struct {
	cursor  clang.Cursor
	cName   string
	goName  string
	comment string
	size    int
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
		cursor:  cursor,
		cName:   cursor.Spelling(),
		goName:  exportedGoName(cursor.Spelling()),
		comment: commentText(cursor.ParsedComment()),
		size:    int(cursor.Type().SizeOf()),
	}

	ss.structs_ = append(ss.structs_, struct_)
}

func (ss *structs) enrich(gen *gen) {
	for i := range ss.structs_ {
		(&ss.structs_[i]).enrich(gen)
	}
}

func (struct_ *struct_) enrich(gen *gen) {
	nameSuffix := 0
	bitfieldOffset := 0
	var prevField *structField

	struct_.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_FieldDecl {
			return clang.ChildVisit_Continue
		}

		isPointer := cursor.Type().Kind() == clang.Type_Pointer

		name := exportedGoName(cursor.Spelling())
		if name == "" && !cursor.IsBitField() {
			name = fmt.Sprintf("_%d", nameSuffix)
			nameSuffix++
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
		field.enum, field.isEnum = gen.enums.find(cursor.Type().Spelling())
		field.scalar, field.isScalar = scalarTypes[cursor.Type().Kind()]
		field.struct_, field.isStruct = gen.structs.find(cursor.Type().Spelling())

		field.isBitfield = cursor.IsBitField()
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
	file.Comment(struct_.comment)
	file.Type().Id(struct_.goName).StructFunc(func(g *jen.Group) {
		g.Id("_").Qual("structs", "HostLayout")
		g.Line()

		var firstBitfield *structField
		var prevField *structField
		for _, field := range struct_.fields {
			struct_.generateFieldPadding(g, field.offset, prevField)

			if !field.isBitfield && prevField != nil && prevField.isBitfield {
				// generate a field to hold one or more bitfields
				g.Id(prevField.bitfieldDataFieldName).Index(jen.Lit((field.offset - firstBitfield.offset) / 8)).Byte()
			}

			if !field.isBitfield {
				g.Comment(field.comment)

				if field.isEnum {
					g.Id(field.goName).Id(field.enum.goName)
				} else if field.isScalar {
					g.Id(field.goName).Add(field.scalar.code)
				} else if field.isStruct {
					g.Id(field.goName).Id(field.struct_.goName)
				} else if field.isPointer {
					g.Id(field.goName).Uintptr().Comment(field.typeSpelling)
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

func (struct_ struct_) generateFieldPadding(g *jen.Group, targetOffset int, prevField *structField) {
	if prevField == nil {
		return
	}

	prevFieldEndOffset := (prevField.offset / 8) + prevField.size
	padding := (targetOffset / 8) - prevFieldEndOffset
	if padding > 0 {
		g.Id("_").Index(jen.Lit(padding)).Byte()
	}
}

func (ss *structs) generate() {
	file := newFile("clang", ".", "struct")
	defer file.save()

	for _, struct_ := range ss.structs_ {
		struct_.generate(file)
	}

	ss.generateTestVars()
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

func (ss *structs) find(cName string) (*struct_, bool) {
	for _, struct_ := range ss.structs_ {
		if struct_.cName == cName {
			return &struct_, true
		}
	}
	return nil, false
}
