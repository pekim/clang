package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type gen struct {
	headerFile string
	clang.TranslationUnit

	enums        enums
	pointerTypes pointerTypes
	structs      structs
}

func Generate() {
	gen := gen{
		headerFile: "clang-c/Index.h",
	}
	gen.TranslationUnit = parseHeaderFile(gen.headerFile)

	gen.findEntities()
	gen.generate()
}

func (gen *gen) findEntities() {
	gen.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are not from clang-c header files
		filename, _, _, _ := cursor.Location().FileLocation()
		if !strings.Contains(filename.Name(), "clang-c/") {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {

		case clang.Cursor_EnumDecl:
			gen.enums.add(cursor)

		case clang.Cursor_StructDecl:
			gen.structs.add(cursor, gen.enums)

		case clang.Cursor_TypedefDecl:
			if cursor.TypedefDeclUnderlyingType().Kind() == clang.Type_Pointer {
				gen.pointerTypes.add(cursor)
			}
		}

		return clang.ChildVisit_Continue
	})
}

func (gen *gen) generate() {
	gen.enums.generate()
	gen.pointerTypes.generate()
	gen.structs.generate()
}
