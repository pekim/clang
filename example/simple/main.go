package main

import (
	"fmt"

	"github.com/pekim/clang"
)

func main() {
	// Initialise the library.
	if err := clang.Init(nil); err != nil {
		panic(err)
	}

	// Parse a header file, creating a TranslationUnit.
	index := clang.CreateIndex(0, 1)
	parseArgs := []string{"-x", "c-header"}
	tu, errorCode := index.ParseTranslationUnit2("internal/clang-c/CXString.h", parseArgs, nil,
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
	)
	if errorCode != clang.Error_Success {
		panic(fmt.Sprintf("failed to create translation unit : %d", errorCode))
	}

	// Traverse the various declarations within the translation unit.
	tuCursor := tu.TranslationUnitCursor()
	tuCursor.VisitChildren(func(cursor clang.Cursor, _parent clang.Cursor) clang.ChildVisitResult {
		fmt.Println(cursor.CursorKind().Spelling(), cursor.Spelling())
		return clang.ChildVisit_Continue
	})
}
