package clang

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var clangResourceDir = sync.OnceValue(func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		panic(err)
	}

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		panic("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		panic(fmt.Sprintf("expected clang resource dir to start with '/', but it %s", resDir))
	}

	return resDir
})

func TestVisitChildren(t *testing.T) {
	assert.NoError(t, Init())

	index := CreateIndex(0, 1)
	assert.NotNil(t, index)

	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	var tu TranslationUnit
	errorCode := ParseTranslationUnit2(index, "test-data/test.h", parseArgs, nil,
		// errorCode := ParseTranslationUnit2(index, "../clang-go-generate/clang-c/Index.h", parseArgs, nil,
		uint32(TranslationUnit_SkipFunctionBodies|TranslationUnit_DetailedPreprocessingRecord),
		&tu,
	)
	assert.Equal(t, Error_Success, errorCode)

	tuCursor := GetTranslationUnitCursor(tu)
	ok := VisitChildren(tuCursor, func(cursor Cursor, _parent Cursor) ChildVisitResult {
		// fmt.Println(
		// 	GetCString(GetCursorKindSpelling(GetCursorKind(cursor))),
		// 	GetCString(GetCursorSpelling(cursor)),
		// 	client_data,
		// )

		kind := GetCursorKind(cursor)
		if kind == Cursor_FunctionDecl {
			assert.Equal(t, "some_function", GetCString(GetCursorSpelling(cursor)))
		}
		assert.True(t, kind == Cursor_MacroDefinition || kind == Cursor_FunctionDecl)

		return ChildVisit_Continue
	})
	assert.Zero(t, ok)
}
