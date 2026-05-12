package generate

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pekim/clang"
)

const ffiPackage = "github.com/go-webgpu/goffi/ffi"
const typesPackage = "github.com/go-webgpu/goffi/types"

type gen struct {
	// headerFile string
	clang.TranslationUnit

	enums        enums
	functions    functions
	pointerTypes pointerTypes
	structs      structs
}

func Generate() {
	err := clang.Init(nil)
	fatalOnError(err)

	gen := gen{}
	gen.parseHeaderFile("internal/clang-c/BuildSystem.h")
	gen.parseHeaderFile("internal/clang-c/CXDiagnostic.h")
	gen.parseHeaderFile("internal/clang-c/Documentation.h")
	gen.parseHeaderFile("internal/clang-c/CXErrorCode.h")
	gen.parseHeaderFile("internal/clang-c/CXFile.h")
	gen.parseHeaderFile("internal/clang-c/CXSourceLocation.h")
	gen.parseHeaderFile("internal/clang-c/CXString.h")
	gen.parseHeaderFile("internal/clang-c/ExternC.h")
	gen.parseHeaderFile("internal/clang-c/Platform.h")
	gen.parseHeaderFile("internal/clang-c/Index.h")
	gen.sortEntities()
	gen.enrich()
	gen.generate()
	gen.printStats()
}

func (gen *gen) findEntities(tu clang.TranslationUnit, headerFile string) {
	tu.TranslationUnitCursor().VisitChildren(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are from diffent files
		file, _, _, _ := cursor.CursorLocation().FileLocation()

		if file.FileName() != headerFile {
			return clang.ChildVisit_Continue
		}

		switch cursor.CursorKind() {

		case clang.Cursor_EnumDecl:
			gen.addEnum(cursor)

		case clang.Cursor_FunctionDecl:
			gen.addFunction(cursor)

		case clang.Cursor_StructDecl:
			gen.addStruct(cursor)

		case clang.Cursor_TypedefDecl:
			if cursor.TypedefDeclUnderlyingType().Kind == clang.Type_Pointer {
				gen.addPointerType(cursor)
			}
		}

		return clang.ChildVisit_Continue
	})
}

func (gen *gen) sortEntities() {
	slices.SortFunc(gen.enums, func(a enum, b enum) int {
		return strings.Compare(a.cName, b.cName)
	})
	slices.SortFunc(gen.functions, func(a function, b function) int {
		return strings.Compare(a.cName, b.cName)
	})
	slices.SortFunc(gen.pointerTypes, func(a pointerType, b pointerType) int {
		return strings.Compare(a.cName, b.cName)
	})
	slices.SortFunc(gen.structs, func(a struct_, b struct_) int {
		return strings.Compare(a.cName, b.cName)
	})
}

func (gen *gen) enrich() {
	gen.functions.enrich(gen)
	gen.structs.enrich(gen)
}

func (gen *gen) generate() {
	gen.enums.generate()
	gen.functions.generate()
	gen.pointerTypes.generate()
	gen.structs.generate()
}

func (gen *gen) printStats() {
	functionsSupported, functionsCount := gen.functions.stats()
	fmt.Printf("functions supported : %d/%d  %.1f%%\n",
		functionsSupported, functionsCount,
		float64(functionsSupported)/float64(functionsCount)*100,
	)

	fmt.Println()
}
