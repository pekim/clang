package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type gen struct {
	headerFile string
	clang.TranslationUnit

	enums enums
}

func Generate() {
	gen := gen{
		headerFile: "clang-c/Index.h",
	}
	gen.TranslationUnit = parseHeaderFile(gen.headerFile)

	gen.generate()
}

func (gen *gen) generate() {
	gen.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are not from clang-c header files
		filename, _, _, _ := cursor.Location().FileLocation()
		if !strings.Contains(filename.Name(), "clang-c/") {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {

		case clang.Cursor_EnumDecl:
			gen.enums.add(cursor)
		}

		return clang.ChildVisit_Continue
	})

	fmt.Println(len(gen.enums.enums))
}
