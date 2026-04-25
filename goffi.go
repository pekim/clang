// This is a generated file. DO NOT EDIT.

package clang

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
	lib "github.com/pekim/clang/internal/lib"
)

var cif_clang_VirtualFileOverlay_create = &types.CallInterface{}
var cif_clang_ModuleMapDescriptor_create = &types.CallInterface{}
var cif_clang_getNullLocation = &types.CallInterface{}
var cif_clang_getNullRange = &types.CallInterface{}
var cif_clang_defaultDiagnosticDisplayOptions = &types.CallInterface{}
var cif_clang_getDiagnosticCategoryName = &types.CallInterface{}
var cif_clang_createIndex = &types.CallInterface{}
var cif_clang_defaultEditingTranslationUnitOptions = &types.CallInterface{}
var cif_clang_getNullCursor = &types.CallInterface{}
var cif_clang_createCXCursorSet = &types.CallInterface{}
var cif_clang_enableStackTraces = &types.CallInterface{}
var cif_clang_defaultCodeCompleteOptions = &types.CallInterface{}
var cif_clang_getClangVersion = &types.CallInterface{}
var cif_clang_toggleCrashRecovery = &types.CallInterface{}

var ptr_clang_VirtualFileOverlay_create unsafe.Pointer
var ptr_clang_ModuleMapDescriptor_create unsafe.Pointer
var ptr_clang_getNullLocation unsafe.Pointer
var ptr_clang_getNullRange unsafe.Pointer
var ptr_clang_defaultDiagnosticDisplayOptions unsafe.Pointer
var ptr_clang_getDiagnosticCategoryName unsafe.Pointer
var ptr_clang_createIndex unsafe.Pointer
var ptr_clang_defaultEditingTranslationUnitOptions unsafe.Pointer
var ptr_clang_getNullCursor unsafe.Pointer
var ptr_clang_createCXCursorSet unsafe.Pointer
var ptr_clang_enableStackTraces unsafe.Pointer
var ptr_clang_defaultCodeCompleteOptions unsafe.Pointer
var ptr_clang_getClangVersion unsafe.Pointer
var ptr_clang_toggleCrashRecovery unsafe.Pointer

func init() {
	library := lib.LoadLibrary(lib.LibraryPaths{
		Darwin: "libclang.dylib",
		Linux:  "libclang.so",
	})

	var err error

	{
		ptr_clang_VirtualFileOverlay_create, err = ffi.GetSymbol(library, "clang_VirtualFileOverlay_create")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_VirtualFileOverlay_create, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_ModuleMapDescriptor_create, err = ffi.GetSymbol(library, "clang_ModuleMapDescriptor_create")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_ModuleMapDescriptor_create, types.DefaultCall, returnType, argTypes)
		}
	}

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
		ptr_clang_defaultDiagnosticDisplayOptions, err = ffi.GetSymbol(library, "clang_defaultDiagnosticDisplayOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_defaultDiagnosticDisplayOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticCategoryName, err = ffi.GetSymbol(library, "clang_getDiagnosticCategoryName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticCategoryName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_createIndex, err = ffi.GetSymbol(library, "clang_createIndex")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.SInt32TypeDescriptor,
				types.SInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_createIndex, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_defaultEditingTranslationUnitOptions, err = ffi.GetSymbol(library, "clang_defaultEditingTranslationUnitOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_defaultEditingTranslationUnitOptions, types.DefaultCall, returnType, argTypes)
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
		ptr_clang_createCXCursorSet, err = ffi.GetSymbol(library, "clang_createCXCursorSet")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_createCXCursorSet, types.DefaultCall, returnType, argTypes)
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
		ptr_clang_defaultCodeCompleteOptions, err = ffi.GetSymbol(library, "clang_defaultCodeCompleteOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_defaultCodeCompleteOptions, types.DefaultCall, returnType, argTypes)
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

	{
		ptr_clang_toggleCrashRecovery, err = ffi.GetSymbol(library, "clang_toggleCrashRecovery")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_toggleCrashRecovery, types.DefaultCall, returnType, argTypes)
		}
	}

}
