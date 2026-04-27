package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

/*
Visit the children of a particular cursor.

This function visits all the direct children of the given cursor,
invoking the given visitor function with the cursors of each visited child.
The traversal may be recursive, if the visitor returns CXChildVisit_Recurse.
The traversal may also be ended prematurely, if the visitor returns CXChildVisit_Break.
*/
func VisitChildren(parent Cursor, visitor func() ChildVisitResult, client_data ClientData) uint32 {
	c_parent := parent
	c_visitor := ffi.NewCallback(func() ChildVisitResult {
		fmt.Println("visit callback")
		return visitor()
	})
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_parent),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(
		cif_clang_visitChildren,
		ptr_clang_visitChildren,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitChildren", err))
	}

	ret := retC
	return ret
}
