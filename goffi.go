// This is a generated file. DO NOT EDIT.

package clang

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
	lib "github.com/pekim/clang/internal/lib"
)

var cif_clang_getNullLocation = &types.CallInterface{}
var cif_clang_getNullRange = &types.CallInterface{}
var cif_clang_getNullCursor = &types.CallInterface{}
var cif_clang_enableStackTraces = &types.CallInterface{}
var cif_clang_getClangVersion = &types.CallInterface{}

var ptr_clang_getNullLocation unsafe.Pointer
var ptr_clang_getNullRange unsafe.Pointer
var ptr_clang_getNullCursor unsafe.Pointer
var ptr_clang_enableStackTraces unsafe.Pointer
var ptr_clang_getClangVersion unsafe.Pointer

func init() {
	library := lib.LoadLibrary(lib.LibraryPaths{
		Darwin: "libclang.dylib",
		Linux:  "libclang.so",
	})

	var err error

	{
		ptr_clang_getNullLocation, err = ffi.GetSymbol(library, "clang_getNullLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getNullLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNullRange, err = ffi.GetSymbol(library, "clang_getNullRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getNullRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNullCursor, err = ffi.GetSymbol(library, "clang_getNullCursor")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getNullCursor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_enableStackTraces, err = ffi.GetSymbol(library, "clang_enableStackTraces")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_enableStackTraces, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getClangVersion, err = ffi.GetSymbol(library, "clang_getClangVersion")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getClangVersion, types.DefaultCall, returnType, argTypes)
		}
	}

}
