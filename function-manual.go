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
func (parent Cursor) VisitChildren(c_visitor uintptr) uint32 {
	c_parent := parent
	var c_client_data ClientData

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_parent),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	_, err := ffi.CallFunction(
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

func MakeChildVisitor(
	visitor func(
		cursor Cursor, parent Cursor,
	) ChildVisitResult,
) uintptr {
	return ffi.NewCallback(func(cursor Cursor, parent Cursor, _client_data unsafe.Pointer) ChildVisitResult {
		return visitor(cursor, parent)
	})
}

/*
Tokenize the source code described by the given range into raw
lexical tokens.

range is the source range in which text should be tokenized. All of the
tokens produced by tokenization will fall within this source range,
*/
func (tu TranslationUnit) Tokenize(range_ SourceRange) []Token {
	var tokens *Token
	c_tokens := &tokens
	var numTokens uint32
	c_numTokens := &numTokens

	args := []unsafe.Pointer{
		unsafe.Pointer(&tu),
		unsafe.Pointer(&range_),
		unsafe.Pointer(&c_tokens),
		unsafe.Pointer(&c_numTokens),
	}

	_, err := ffi.CallFunction(
		cif_clang_tokenize,
		ptr_clang_tokenize,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitChildren", err))
	}

	return unsafe.Slice(tokens, numTokens)
}
