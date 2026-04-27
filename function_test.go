package clang

import (
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

	index := CreateIndex(0, 0)
	assert.NotNil(t, index)

	tu := CreateTranslationUnitFromSourceFile(index, "test-data/test.h",
		[]string{"qaz", "qwerty"}, nil)
	assert.NotNil(t, tu)
}
