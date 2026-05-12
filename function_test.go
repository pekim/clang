package clang

import (
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionNoArgsNoReturn(t *testing.T) {
	assert.NoError(t, Init(nil))

	EnableStackTraces()
}

func TestFunctionReturnCXString(t *testing.T) {
	assert.NoError(t, Init(nil))

	assert.True(t, strings.Contains(GetClangVersion(), "clang version "))
}

func TestFunctionScalarArgsPointerTypeReturn(t *testing.T) {
	assert.NoError(t, Init(nil))

	index := CreateIndex(0, 0)
	assert.NotNil(t, index)
}

func TestFunctionStringSliceArg(t *testing.T) {
	assert.NoError(t, Init(nil))

	index := CreateIndex(0, 1)
	assert.NotNil(t, index)

	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	tu := index.CreateTranslationUnitFromSourceFile("test-data/test.h", parseArgs, nil)
	assert.NotNil(t, tu)
	cursor := tu.TranslationUnitCursor()
	assert.Equal(t, Cursor_TranslationUnit, cursor.CursorKind())
}
