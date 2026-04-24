package clang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionNoArgsNoReturn(t *testing.T) {
	assert.NotPanics(t, func() {
		EnableStackTraces()
	})
}

// func TestFunctionReturnCXString(t *testing.T) {
// 	GetClangVersion()
// }
