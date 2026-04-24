package clang

import (
	"testing"
	"unsafe"

	"github.com/go-webgpu/goffi/ffi"
	"github.com/stretchr/testify/assert"
)

func TestCallFunctionThroughGoffi(t *testing.T) {
	err := ffi.CallFunction(cif_clang_enableStackTraces, ptr_clang_enableStackTraces, nil, []unsafe.Pointer{})
	assert.NoError(t, err)
}
