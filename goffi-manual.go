package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
)

var cif_clang_visitChildren = &types.CallInterface{}
var ptr_clang_visitChildren unsafe.Pointer

func initManual(library unsafe.Pointer) error {
	var err error

	{
		ptr_clang_visitChildren, err = ffi.GetSymbol(library, "clang_visitChildren")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_visitChildren, types.DefaultCall, returnType, argTypes)
			if err != nil {
				return fmt.Errorf("failed to prepare call interface for %s : %w", "clang_visitChildren", err)
			}
		}
	}

	return nil
}
