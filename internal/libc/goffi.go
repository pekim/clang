package libc

import (
	"fmt"
	"unsafe"

	"github.com/go-webgpu/goffi/ffi"
	"github.com/go-webgpu/goffi/types"

	"github.com/pekim/clang/internal/lib"
)

var cif_free = &types.CallInterface{}
var cif_malloc = &types.CallInterface{}

var ptr_free unsafe.Pointer
var ptr_malloc unsafe.Pointer

func init() {
	library := lib.LoadLibrary(lib.LibraryPaths{
		Darwin: "libSystem.dylib",
		Linux:  "libc.so.6",
	})

	var err error

	{
		ptr_free, err = ffi.GetSymbol(library, "free")
		if err != nil {
			panic(fmt.Sprintf("failed to get symbol 'free' : %s", err))
		}
		returnType := types.VoidTypeDescriptor
		argTypes := []*types.TypeDescriptor{
			types.PointerTypeDescriptor,
		}
		_ = ffi.PrepareCallInterface(cif_free, types.DefaultCall, returnType, argTypes)
	}
	{

		ptr_malloc, err = ffi.GetSymbol(library, "malloc")
		if err != nil {
			panic(fmt.Sprintf("failed to get symbol 'malloc' : %s", err))
		}
		returnType := types.PointerTypeDescriptor
		argTypes := []*types.TypeDescriptor{
			types.UInt64TypeDescriptor,
		}
		_ = ffi.PrepareCallInterface(cif_malloc, types.DefaultCall, returnType, argTypes)
	}
}

func free(ptr *byte) {
	args := []unsafe.Pointer{unsafe.Pointer(&ptr)}
	err := ffi.CallFunction(cif_free, ptr_free, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call free : %s", err))
	}
}

func malloc(length int) *byte {
	var retC unsafe.Pointer
	args := []unsafe.Pointer{unsafe.Pointer(&length)}
	err := ffi.CallFunction(cif_malloc, ptr_malloc, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call malloc : %s", err))
	}
	return (*byte)(retC)
}
