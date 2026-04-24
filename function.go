// This is a generated file. DO NOT EDIT.

package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

func EnableStackTraces() {
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_enableStackTraces, ptr_clang_enableStackTraces, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_enableStackTraces", err))
	}
}
