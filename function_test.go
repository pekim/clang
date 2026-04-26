package clang

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionNoArgsNoReturn(t *testing.T) {
	assert.NotPanics(t, func() {
		EnableStackTraces()
	})
}

func TestFunctionReturnCXString(t *testing.T) {
	cxString := GetClangVersion()
	version := GetCString(cxString)
	assert.True(t, strings.HasPrefix(version, "clang version "))
}

func TestFunctionScalarArgsPointerTypeReturn(t *testing.T) {
	index := CreateIndex(0, 0)
	assert.NotNil(t, index)
}
