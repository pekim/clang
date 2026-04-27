package clang

import (
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionNoArgsNoReturn(t *testing.T) {
	assert.NoError(t, Init())

	EnableStackTraces()
}

func TestFunctionReturnCXString(t *testing.T) {
	assert.NoError(t, Init())

	cxString := GetClangVersion()
	version := GetCString(cxString)
	assert.True(t, strings.HasPrefix(version, "clang version "))
}

func TestFunctionScalarArgsPointerTypeReturn(t *testing.T) {
	assert.NoError(t, Init())

	index := CreateIndex(0, 0)
	assert.NotNil(t, index)
}

func TestFunctionStringSliceArg(t *testing.T) {
	assert.NoError(t, Init())

	index := CreateIndex(0, 1)
	assert.NotNil(t, index)

	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	tu := CreateTranslationUnitFromSourceFile(index, "test-data/test.h", parseArgs, nil)
	assert.NotNil(t, tu)
	cursor := GetTranslationUnitCursor(tu)
	assert.Equal(t, Cursor_TranslationUnit, GetCursorKind(cursor))
}
