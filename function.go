// This is a generated file. DO NOT EDIT.

package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	libc "github.com/pekim/clang/internal/libc"
)

//
/*
Wraps the C function clang_BlockCommandComment_getArgText.
*/
func (comment Comment) BlockCommandCommentArgText(argIdx uint32) string {
	c_comment := comment
	c_argIdx := argIdx

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
		unsafe.Pointer(&c_argIdx),
	}

	err := ffi.CallFunction(
		cif_clang_BlockCommandComment_getArgText,
		ptr_clang_BlockCommandComment_getArgText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_BlockCommandComment_getArgText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_BlockCommandComment_getCommandName.
*/
func (comment Comment) BlockCommandCommentCommandName() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_BlockCommandComment_getCommandName,
		ptr_clang_BlockCommandComment_getCommandName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_BlockCommandComment_getCommandName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_BlockCommandComment_getNumArgs.
*/
func (comment Comment) BlockCommandCommentNumArgs() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_BlockCommandComment_getNumArgs,
		ptr_clang_BlockCommandComment_getNumArgs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_BlockCommandComment_getNumArgs", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_BlockCommandComment_getParagraph.
*/
func (comment Comment) BlockCommandCommentParagraph() Comment {
	c_comment := comment

	var retC Comment
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_BlockCommandComment_getParagraph,
		ptr_clang_BlockCommandComment_getParagraph,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_BlockCommandComment_getParagraph", err))
	}

	ret := retC
	return ret
}

// Queries a CXCursorSet to see if it contains a specific CXCursor.
/*
Wraps the C function clang_CXCursorSet_contains.
*/
func (cset CursorSet) Contains(cursor Cursor) uint32 {
	c_cset := cset
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cset.ptr),
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_CXCursorSet_contains,
		ptr_clang_CXCursorSet_contains,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXCursorSet_contains", err))
	}

	ret := retC
	return ret
}

// Inserts a CXCursor into a CXCursorSet.
/*
Wraps the C function clang_CXCursorSet_insert.
*/
func (cset CursorSet) Insert(cursor Cursor) uint32 {
	c_cset := cset
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cset.ptr),
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_CXCursorSet_insert,
		ptr_clang_CXCursorSet_insert,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXCursorSet_insert", err))
	}

	ret := retC
	return ret
}

/*
Gets the general options associated with a CXIndex.

This function allows to obtain the final option values used by libclang after specifying the option policies via CXChoice enumerators.
*/
/*
Wraps the C function clang_CXIndex_getGlobalOptions.
*/
func (p0 Index) GlobalOptions() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_CXIndex_getGlobalOptions,
		ptr_clang_CXIndex_getGlobalOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXIndex_getGlobalOptions", err))
	}

	ret := retC
	return ret
}

/*
Sets general options associated with a CXIndex.

This function is DEPRECATED. Set CXIndexOptions::ThreadBackgroundPriorityForIndexing and/or CXIndexOptions::ThreadBackgroundPriorityForEditing and call clang_createIndexWithOptions() instead.

For example:
*/
/*
Wraps the C function clang_CXIndex_setGlobalOptions.
*/
func (p0 Index) SetGlobalOptions(options uint32) {
	c_p0 := p0
	c_options := options

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_CXIndex_setGlobalOptions,
		ptr_clang_CXIndex_setGlobalOptions,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXIndex_setGlobalOptions", err))
	}
}

/*
Sets the invocation emission path option in a CXIndex.

This function is DEPRECATED. Set CXIndexOptions::InvocationEmissionPath and call clang_createIndexWithOptions() instead.

The invocation emission path specifies a path which will contain log files for certain libclang invocations. A null value (default) implies that libclang invocations are not logged..
*/
/*
Wraps the C function clang_CXIndex_setInvocationEmissionPathOption.
*/
func (p0 Index) SetInvocationEmissionPathOption(path string) {
	c_p0 := p0
	c_path, free_c_path := libc.CString(path)
	defer free_c_path()

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_path),
	}

	err := ffi.CallFunction(
		cif_clang_CXIndex_setInvocationEmissionPathOption,
		ptr_clang_CXIndex_setInvocationEmissionPathOption,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXIndex_setInvocationEmissionPathOption", err))
	}
}

// Determine if a C++ constructor is a converting constructor.
/*
Wraps the C function clang_CXXConstructor_isConvertingConstructor.
*/
func (c Cursor) XConstructor_isConvertingConstructor() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXConstructor_isConvertingConstructor,
		ptr_clang_CXXConstructor_isConvertingConstructor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isConvertingConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is a copy constructor.
/*
Wraps the C function clang_CXXConstructor_isCopyConstructor.
*/
func (c Cursor) XConstructor_isCopyConstructor() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXConstructor_isCopyConstructor,
		ptr_clang_CXXConstructor_isCopyConstructor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isCopyConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is the default constructor.
/*
Wraps the C function clang_CXXConstructor_isDefaultConstructor.
*/
func (c Cursor) XConstructor_isDefaultConstructor() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXConstructor_isDefaultConstructor,
		ptr_clang_CXXConstructor_isDefaultConstructor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isDefaultConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is a move constructor.
/*
Wraps the C function clang_CXXConstructor_isMoveConstructor.
*/
func (c Cursor) XConstructor_isMoveConstructor() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXConstructor_isMoveConstructor,
		ptr_clang_CXXConstructor_isMoveConstructor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isMoveConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ field is declared 'mutable'.
/*
Wraps the C function clang_CXXField_isMutable.
*/
func (c Cursor) XField_isMutable() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXField_isMutable,
		ptr_clang_CXXField_isMutable,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXField_isMutable", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is declared 'const'.
/*
Wraps the C function clang_CXXMethod_isConst.
*/
func (c Cursor) XMethod_isConst() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isConst,
		ptr_clang_CXXMethod_isConst,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isConst", err))
	}

	ret := retC
	return ret
}

/*
Determine if a C++ member function is a copy-assignment operator, returning 1 if such is the case and 0 otherwise.

> A copy-assignment operator `X::operator=` is a non-static, > non-template member function of _class_ `X` with exactly one > parameter of type `X`, `X&`, `const X&`, `volatile X&` or `const > volatile X&`.

That is, for example, the `operator=` in:

class Foo {        bool operator=(const volatile Foo&);    };

Is a copy-assignment operator, while the `operator=` in:

class Bar {        bool operator=(const int&);    };

Is not.
*/
/*
Wraps the C function clang_CXXMethod_isCopyAssignmentOperator.
*/
func (c Cursor) XMethod_isCopyAssignmentOperator() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isCopyAssignmentOperator,
		ptr_clang_CXXMethod_isCopyAssignmentOperator,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isCopyAssignmentOperator", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ method is declared '= default'.
/*
Wraps the C function clang_CXXMethod_isDefaulted.
*/
func (c Cursor) XMethod_isDefaulted() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isDefaulted,
		ptr_clang_CXXMethod_isDefaulted,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isDefaulted", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ method is declared '= delete'.
/*
Wraps the C function clang_CXXMethod_isDeleted.
*/
func (c Cursor) XMethod_isDeleted() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isDeleted,
		ptr_clang_CXXMethod_isDeleted,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isDeleted", err))
	}

	ret := retC
	return ret
}

/*
Determines if a C++ constructor or conversion function was declared explicit, returning 1 if such is the case and 0 otherwise.

Constructors or conversion functions are declared explicit through the use of the explicit specifier.

For example, the following constructor and conversion function are not explicit as they lack the explicit specifier:

class Foo {         Foo();         operator int();     };

While the following constructor and conversion function are explicit as they are declared with the explicit specifier.

class Foo {         explicit Foo();         explicit operator int();     };

This function will return 0 when given a cursor pointing to one of the former declarations and it will return 1 for a cursor pointing to the latter declarations.

The explicit specifier allows the user to specify a conditional compile-time expression whose value decides whether the marked element is explicit or not.

For example:

constexpr bool foo(int i) { return i % 2 == 0; }

class Foo {          explicit(foo(1)) Foo();          explicit(foo(2)) operator int();     }

This function will return 0 for the constructor and 1 for the conversion function.
*/
/*
Wraps the C function clang_CXXMethod_isExplicit.
*/
func (c Cursor) XMethod_isExplicit() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isExplicit,
		ptr_clang_CXXMethod_isExplicit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isExplicit", err))
	}

	ret := retC
	return ret
}

/*
Determine if a C++ member function is a move-assignment operator, returning 1 if such is the case and 0 otherwise.

> A move-assignment operator `X::operator=` is a non-static, > non-template member function of _class_ `X` with exactly one > parameter of type `X&&`, `const X&&`, `volatile X&&` or `const > volatile X&&`.

That is, for example, the `operator=` in:

class Foo {        bool operator=(const volatile Foo&&);    };

Is a move-assignment operator, while the `operator=` in:

class Bar {        bool operator=(const int&&);    };

Is not.
*/
/*
Wraps the C function clang_CXXMethod_isMoveAssignmentOperator.
*/
func (c Cursor) XMethod_isMoveAssignmentOperator() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isMoveAssignmentOperator,
		ptr_clang_CXXMethod_isMoveAssignmentOperator,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isMoveAssignmentOperator", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is pure virtual.
/*
Wraps the C function clang_CXXMethod_isPureVirtual.
*/
func (c Cursor) XMethod_isPureVirtual() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isPureVirtual,
		ptr_clang_CXXMethod_isPureVirtual,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isPureVirtual", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is declared 'static'.
/*
Wraps the C function clang_CXXMethod_isStatic.
*/
func (c Cursor) XMethod_isStatic() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isStatic,
		ptr_clang_CXXMethod_isStatic,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isStatic", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is explicitly declared 'virtual' or if it overrides a virtual method from one of the base classes.
/*
Wraps the C function clang_CXXMethod_isVirtual.
*/
func (c Cursor) XMethod_isVirtual() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXMethod_isVirtual,
		ptr_clang_CXXMethod_isVirtual,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isVirtual", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ record is abstract, i.e. whether a class or struct has a pure virtual member function.
/*
Wraps the C function clang_CXXRecord_isAbstract.
*/
func (c Cursor) XRecord_isAbstract() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_CXXRecord_isAbstract,
		ptr_clang_CXXRecord_isAbstract,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXRecord_isAbstract", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Comment_getChild.
*/
func (comment Comment) Child(childIdx uint32) Comment {
	c_comment := comment
	c_childIdx := childIdx

	var retC Comment
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
		unsafe.Pointer(&c_childIdx),
	}

	err := ffi.CallFunction(
		cif_clang_Comment_getChild,
		ptr_clang_Comment_getChild,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Comment_getChild", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Comment_getKind.
*/
func (comment Comment) Kind() CommentKind {
	c_comment := comment

	var retC CommentKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_Comment_getKind,
		ptr_clang_Comment_getKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Comment_getKind", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Comment_getNumChildren.
*/
func (comment Comment) NumChildren() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_Comment_getNumChildren,
		ptr_clang_Comment_getNumChildren,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Comment_getNumChildren", err))
	}

	ret := retC
	return ret
}

/*
A CXComment_Paragraph node is considered whitespace if it contains only CXComment_Text nodes that are empty or whitespace.

Other AST nodes (except CXComment_Paragraph and CXComment_Text) are never considered whitespace.
*/
/*
Wraps the C function clang_Comment_isWhitespace.
*/
func (comment Comment) IsWhitespace() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_Comment_isWhitespace,
		ptr_clang_Comment_isWhitespace,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Comment_isWhitespace", err))
	}

	ret := retC
	return ret
}

// If cursor is a statement declaration tries to evaluate the statement and if its variable, tries to evaluate its initializer, into its corresponding type. If it's an expression, tries to evaluate the expression.
/*
Wraps the C function clang_Cursor_Evaluate.
*/
func (c Cursor) Evaluate() EvalResult {
	c_c := c

	var retC EvalResult
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_Evaluate,
		ptr_clang_Cursor_Evaluate,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_Evaluate", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the argument cursor of a function or method.

The argument cursor can be determined for calls as well as for declarations of functions or methods. For other cursors and for invalid indices, an invalid cursor is returned.
*/
/*
Wraps the C function clang_Cursor_getArgument.
*/
func (c Cursor) Argument(i uint32) Cursor {
	c_c := c
	c_i := i

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getArgument,
		ptr_clang_Cursor_getArgument,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getArgument", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Cursor_getBinaryOpcode.
*/
func (c Cursor) BinaryOpcode() BinaryOperatorKind_ {
	c_c := c

	var retC BinaryOperatorKind_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getBinaryOpcode,
		ptr_clang_Cursor_getBinaryOpcode,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getBinaryOpcode", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Cursor_getBinaryOpcodeStr.
*/
func (op BinaryOperatorKind_) CursorBinaryOpcodeStr() string {
	c_op := op

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_op),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getBinaryOpcodeStr,
		ptr_clang_Cursor_getBinaryOpcodeStr,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getBinaryOpcodeStr", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Given a cursor that represents a documentable entity (e.g., declaration), return the associated

first paragraph.
*/
/*
Wraps the C function clang_Cursor_getBriefCommentText.
*/
func (c Cursor) BriefCommentText() string {
	c_c := c

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getBriefCommentText,
		ptr_clang_Cursor_getBriefCommentText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getBriefCommentText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Retrieve the CXStrings representing the mangled symbols of the C++ constructor or destructor at the cursor.
/*
Wraps the C function clang_Cursor_getCXXManglings.
*/
func (p0 Cursor) XManglings() *StringSet {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getCXXManglings,
		ptr_clang_Cursor_getCXXManglings,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getCXXManglings", err))
	}

	ret := (*StringSet)(retC)
	return ret
}

// Given a cursor that represents a declaration, return the associated comment's source range.  The range may include multiple consecutive comments with whitespace in between.
/*
Wraps the C function clang_Cursor_getCommentRange.
*/
func (c Cursor) CommentRange() SourceRange {
	c_c := c

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getCommentRange,
		ptr_clang_Cursor_getCommentRange,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getCommentRange", err))
	}

	ret := retC
	return ret
}

/*
Given a CXCursor_GCCAsmStmt cursor, get the Index-th clobber of it. This function returns a valid empty string if the cursor does not point at a GCC inline assembly block or `Index` is out of bounds.

Users are responsible for releasing the allocation of returned string via clang_disposeString.
*/
/*
Wraps the C function clang_Cursor_getGCCAssemblyClobber.
*/
func (cursor Cursor) GCCAssemblyClobber(index uint32) string {
	c_cursor := cursor
	c_index := index

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyClobber,
		ptr_clang_Cursor_getGCCAssemblyClobber,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyClobber", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Given a CXCursor_GCCAsmStmt cursor, get the constraint and expression cursor to the Index-th input. This function returns 1 when the cursor points at a GCC inline assembly statement, `Index` is within bounds and both the `Constraint` and `Expr` are not NULL. Otherwise, this function returns 0 but leaves `Constraint` and `Expr` intact.

Users are responsible for releasing the allocation of `Constraint` via clang_disposeString.
*/
/*
Wraps the C function clang_Cursor_getGCCAssemblyInput.
*/
func (cursor Cursor) GCCAssemblyInput(index uint32) (string, Cursor, uint32) {
	c_cursor := cursor
	c_index := index
	var constraint string_
	c_constraint := &constraint
	var expr Cursor
	c_expr := &expr

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_index),
		unsafe.Pointer(&c_constraint),
		unsafe.Pointer(&c_expr),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyInput,
		ptr_clang_Cursor_getGCCAssemblyInput,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyInput", err))
	}

	ret := retC
	return constraint.CString(), expr, ret
}

// Given a CXCursor_GCCAsmStmt cursor, count the clobbers in it. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
/*
Wraps the C function clang_Cursor_getGCCAssemblyNumClobbers.
*/
func (cursor Cursor) GCCAssemblyNumClobbers() uint32 {
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyNumClobbers,
		ptr_clang_Cursor_getGCCAssemblyNumClobbers,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyNumClobbers", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, count the number of inputs. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
/*
Wraps the C function clang_Cursor_getGCCAssemblyNumInputs.
*/
func (p0 Cursor) GCCAssemblyNumInputs() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyNumInputs,
		ptr_clang_Cursor_getGCCAssemblyNumInputs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyNumInputs", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, count the number of outputs. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
/*
Wraps the C function clang_Cursor_getGCCAssemblyNumOutputs.
*/
func (p0 Cursor) GCCAssemblyNumOutputs() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyNumOutputs,
		ptr_clang_Cursor_getGCCAssemblyNumOutputs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyNumOutputs", err))
	}

	ret := retC
	return ret
}

/*
Given a CXCursor_GCCAsmStmt cursor, get the constraint and expression cursor to the Index-th output. This function returns 1 when the cursor points at a GCC inline assembly statement, `Index` is within bounds and both the `Constraint` and `Expr` are not NULL. Otherwise, this function returns 0 but leaves `Constraint` and `Expr` intact.

Users are responsible for releasing the allocation of `Constraint` via clang_disposeString.
*/
/*
Wraps the C function clang_Cursor_getGCCAssemblyOutput.
*/
func (cursor Cursor) GCCAssemblyOutput(index uint32) (string, Cursor, uint32) {
	c_cursor := cursor
	c_index := index
	var constraint string_
	c_constraint := &constraint
	var expr Cursor
	c_expr := &expr

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_index),
		unsafe.Pointer(&c_constraint),
		unsafe.Pointer(&c_expr),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyOutput,
		ptr_clang_Cursor_getGCCAssemblyOutput,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyOutput", err))
	}

	ret := retC
	return constraint.CString(), expr, ret
}

/*
Given a CXCursor_GCCAsmStmt cursor, return the assembly template string. As per LLVM IR Assembly Template language, template placeholders for inputs and outputs are either of the form $N where N is a decimal number as an index into the input-output specification, or ${N:M} where N is a decimal number also as an index into the input-output specification and M is the template argument modifier. The index N in both cases points into the the total inputs and outputs, or more specifically, into the list of outputs followed by the inputs, starting from index 0 as the first available template argument.

This function also returns a valid empty string if the cursor does not point at a GCC inline assembly block.

Users are responsible for releasing the allocation of returned string via clang_disposeString.
*/
/*
Wraps the C function clang_Cursor_getGCCAssemblyTemplate.
*/
func (p0 Cursor) GCCAssemblyTemplate() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getGCCAssemblyTemplate,
		ptr_clang_Cursor_getGCCAssemblyTemplate,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyTemplate", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Retrieve the CXString representing the mangled name of the cursor.
/*
Wraps the C function clang_Cursor_getMangling.
*/
func (p0 Cursor) Mangling() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getMangling,
		ptr_clang_Cursor_getMangling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getMangling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Given a CXCursor_ModuleImportDecl cursor, return the associated module.
/*
Wraps the C function clang_Cursor_getModule.
*/
func (c Cursor) Module() Module {
	c_c := c

	var retC Module
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getModule,
		ptr_clang_Cursor_getModule,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getModule", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of non-variadic arguments associated with a given cursor.

The number of arguments can be determined for calls as well as for declarations of functions or methods. For other cursors -1 is returned.
*/
/*
Wraps the C function clang_Cursor_getNumArguments.
*/
func (c Cursor) NumArguments() int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getNumArguments,
		ptr_clang_Cursor_getNumArguments,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getNumArguments", err))
	}

	ret := retC
	return ret
}

/*
Returns the number of template args of a function, struct, or class decl representing a template specialization.

If the argument cursor cannot be converted into a template function declaration, -1 is returned.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

The value 3 would be returned from this call.
*/
/*
Wraps the C function clang_Cursor_getNumTemplateArguments.
*/
func (c Cursor) NumTemplateArguments() int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getNumTemplateArguments,
		ptr_clang_Cursor_getNumTemplateArguments,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getNumTemplateArguments", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents an Objective-C method or parameter declaration, return the associated Objective-C qualifiers for the return type or the parameter respectively. The bits are formed from CXObjCDeclQualifierKind.
/*
Wraps the C function clang_Cursor_getObjCDeclQualifiers.
*/
func (c Cursor) ObjCDeclQualifiers() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getObjCDeclQualifiers,
		ptr_clang_Cursor_getObjCDeclQualifiers,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCDeclQualifiers", err))
	}

	ret := retC
	return ret
}

// Retrieve the CXStrings representing the mangled symbols of the ObjC class interface or implementation at the cursor.
/*
Wraps the C function clang_Cursor_getObjCManglings.
*/
func (p0 Cursor) ObjCManglings() *StringSet {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getObjCManglings,
		ptr_clang_Cursor_getObjCManglings,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCManglings", err))
	}

	ret := (*StringSet)(retC)
	return ret
}

// Given a cursor that represents a property declaration, return the associated property attributes. The bits are formed from CXObjCPropertyAttrKind.
/*
Wraps the C function clang_Cursor_getObjCPropertyAttributes.
*/
func (c Cursor) ObjCPropertyAttributes(reserved uint32) uint32 {
	c_c := c
	c_reserved := reserved

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_reserved),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getObjCPropertyAttributes,
		ptr_clang_Cursor_getObjCPropertyAttributes,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCPropertyAttributes", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a property declaration, return the name of the method that implements the getter.
/*
Wraps the C function clang_Cursor_getObjCPropertyGetterName.
*/
func (c Cursor) ObjCPropertyGetterName() string {
	c_c := c

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getObjCPropertyGetterName,
		ptr_clang_Cursor_getObjCPropertyGetterName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCPropertyGetterName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Given a cursor that represents a property declaration, return the name of the method that implements the setter, if any.
/*
Wraps the C function clang_Cursor_getObjCPropertySetterName.
*/
func (c Cursor) ObjCPropertySetterName() string {
	c_c := c

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getObjCPropertySetterName,
		ptr_clang_Cursor_getObjCPropertySetterName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCPropertySetterName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
If the cursor points to a selector identifier in an Objective-C method or message expression, this returns the selector index.

After getting a cursor with #clang_getCursor, this can be called to determine if the location points to a selector identifier.
*/
/*
Wraps the C function clang_Cursor_getObjCSelectorIndex.
*/
func (p0 Cursor) ObjCSelectorIndex() int32 {
	c_p0 := p0

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getObjCSelectorIndex,
		ptr_clang_Cursor_getObjCSelectorIndex,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCSelectorIndex", err))
	}

	ret := retC
	return ret
}

/*
Return the offset of the field represented by the Cursor.

If the cursor is not a field declaration, -1 is returned. If the cursor semantic parent is not a record field declaration,   CXTypeLayoutError_Invalid is returned. If the field's type declaration is an incomplete type,   CXTypeLayoutError_Incomplete is returned. If the field's type declaration is a dependent type,   CXTypeLayoutError_Dependent is returned. If the field's name S is not found,   CXTypeLayoutError_InvalidFieldName is returned.
*/
/*
Wraps the C function clang_Cursor_getOffsetOfField.
*/
func (c Cursor) OffsetOfField() int64 {
	c_c := c

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getOffsetOfField,
		ptr_clang_Cursor_getOffsetOfField,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getOffsetOfField", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a documentable entity (e.g., declaration), return the associated parsed comment as a CXComment_FullComment AST node.
/*
Wraps the C function clang_Cursor_getParsedComment.
*/
func (c Cursor) ParsedComment() Comment {
	c_c := c

	var retC Comment
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getParsedComment,
		ptr_clang_Cursor_getParsedComment,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getParsedComment", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a declaration, return the associated comment text, including comment markers.
/*
Wraps the C function clang_Cursor_getRawCommentText.
*/
func (c Cursor) RawCommentText() string {
	c_c := c

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getRawCommentText,
		ptr_clang_Cursor_getRawCommentText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getRawCommentText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Given a cursor pointing to an Objective-C message or property reference, or C++ method call, returns the CXType of the receiver.
/*
Wraps the C function clang_Cursor_getReceiverType.
*/
func (c Cursor) ReceiverType() Type {
	c_c := c

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getReceiverType,
		ptr_clang_Cursor_getReceiverType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getReceiverType", err))
	}

	ret := retC
	return ret
}

// Retrieve a range for a piece that forms the cursors spelling name. Most of the times there is only one range for the complete spelling but for Objective-C methods and Objective-C message expressions, there are multiple pieces for each selector identifier.
/*
Wraps the C function clang_Cursor_getSpellingNameRange.
*/
func (p0 Cursor) SpellingNameRange(pieceIndex uint32, options uint32) SourceRange {
	c_p0 := p0
	c_pieceIndex := pieceIndex
	c_options := options

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_pieceIndex),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getSpellingNameRange,
		ptr_clang_Cursor_getSpellingNameRange,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getSpellingNameRange", err))
	}

	ret := retC
	return ret
}

/*
Returns the storage class for a function or variable declaration.

If the passed in Cursor is not a function or variable declaration, CX_SC_Invalid is returned else the storage class.
*/
/*
Wraps the C function clang_Cursor_getStorageClass.
*/
func (p0 Cursor) StorageClass() StorageClass {
	c_p0 := p0

	var retC StorageClass
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getStorageClass,
		ptr_clang_Cursor_getStorageClass,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getStorageClass", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the kind of the I'th template argument of the CXCursor C.

If the argument CXCursor does not represent a FunctionDecl, StructDecl, or ClassTemplatePartialSpecialization, an invalid template argument kind is returned.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

For I = 0, 1, and 2, Type, Integral, and Integral will be returned, respectively.
*/
/*
Wraps the C function clang_Cursor_getTemplateArgumentKind.
*/
func (c Cursor) TemplateArgumentKind(i uint32) TemplateArgumentKind {
	c_c := c
	c_i := i

	var retC TemplateArgumentKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getTemplateArgumentKind,
		ptr_clang_Cursor_getTemplateArgumentKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTemplateArgumentKind", err))
	}

	ret := retC
	return ret
}

/*
Retrieve a CXType representing the type of a TemplateArgument of a  function decl representing a template specialization.

If the argument CXCursor does not represent a FunctionDecl, StructDecl, ClassDecl or ClassTemplatePartialSpecialization whose I'th template argument has a kind of CXTemplateArgKind_Integral, an invalid type is returned.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

If called with I = 0, "float", will be returned. Invalid types will be returned for I == 1 or 2.
*/
/*
Wraps the C function clang_Cursor_getTemplateArgumentType.
*/
func (c Cursor) TemplateArgumentType(i uint32) Type {
	c_c := c
	c_i := i

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getTemplateArgumentType,
		ptr_clang_Cursor_getTemplateArgumentType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTemplateArgumentType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the value of an Integral TemplateArgument (of a function  decl representing a template specialization) as an unsigned long long.

It is undefined to call this function on a CXCursor that does not represent a FunctionDecl, StructDecl, ClassDecl or ClassTemplatePartialSpecialization or whose I'th template argument is not an integral value.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, 2147483649, true>();

If called with I = 1 or 2, 2147483649 or true will be returned, respectively. For I == 0, this function's behavior is undefined.
*/
/*
Wraps the C function clang_Cursor_getTemplateArgumentUnsignedValue.
*/
func (c Cursor) TemplateArgumentUnsignedValue(i uint32) uint64 {
	c_c := c
	c_i := i

	var retC uint64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getTemplateArgumentUnsignedValue,
		ptr_clang_Cursor_getTemplateArgumentUnsignedValue,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTemplateArgumentUnsignedValue", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the value of an Integral TemplateArgument (of a function  decl representing a template specialization) as a signed long long.

It is undefined to call this function on a CXCursor that does not represent a FunctionDecl, StructDecl, ClassDecl or ClassTemplatePartialSpecialization whose I'th template argument is not an integral value.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

If called with I = 1 or 2, -7 or true will be returned, respectively. For I == 0, this function's behavior is undefined.
*/
/*
Wraps the C function clang_Cursor_getTemplateArgumentValue.
*/
func (c Cursor) TemplateArgumentValue(i uint32) int64 {
	c_c := c
	c_i := i

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getTemplateArgumentValue,
		ptr_clang_Cursor_getTemplateArgumentValue,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTemplateArgumentValue", err))
	}

	ret := retC
	return ret
}

// Returns the translation unit that a cursor originated from.
/*
Wraps the C function clang_Cursor_getTranslationUnit.
*/
func (p0 Cursor) TranslationUnit() TranslationUnit {
	c_p0 := p0

	var retC TranslationUnit
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getTranslationUnit,
		ptr_clang_Cursor_getTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTranslationUnit", err))
	}

	ret := retC
	return ret
}

// If cursor refers to a variable declaration and it has initializer returns cursor referring to the initializer otherwise return null cursor.
/*
Wraps the C function clang_Cursor_getVarDeclInitializer.
*/
func (cursor Cursor) VarDeclInitializer() Cursor {
	c_cursor := cursor

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_getVarDeclInitializer,
		ptr_clang_Cursor_getVarDeclInitializer,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getVarDeclInitializer", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor has any attributes.
/*
Wraps the C function clang_Cursor_hasAttrs.
*/
func (c Cursor) HasAttrs() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_hasAttrs,
		ptr_clang_Cursor_hasAttrs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_hasAttrs", err))
	}

	ret := retC
	return ret
}

// If cursor refers to a variable declaration that has external storage returns 1. If cursor refers to a variable declaration that doesn't have external storage returns 0. Otherwise returns -1.
/*
Wraps the C function clang_Cursor_hasVarDeclExternalStorage.
*/
func (cursor Cursor) HasVarDeclExternalStorage() int32 {
	c_cursor := cursor

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_hasVarDeclExternalStorage,
		ptr_clang_Cursor_hasVarDeclExternalStorage,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_hasVarDeclExternalStorage", err))
	}

	ret := retC
	return ret
}

// If cursor refers to a variable declaration that has global storage returns 1. If cursor refers to a variable declaration that doesn't have global storage returns 0. Otherwise returns -1.
/*
Wraps the C function clang_Cursor_hasVarDeclGlobalStorage.
*/
func (cursor Cursor) HasVarDeclGlobalStorage() int32 {
	c_cursor := cursor

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_hasVarDeclGlobalStorage,
		ptr_clang_Cursor_hasVarDeclGlobalStorage,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_hasVarDeclGlobalStorage", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor represents an anonymous tag or namespace
/*
Wraps the C function clang_Cursor_isAnonymous.
*/
func (c Cursor) IsAnonymous() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isAnonymous,
		ptr_clang_Cursor_isAnonymous,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isAnonymous", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor represents an anonymous record declaration.
/*
Wraps the C function clang_Cursor_isAnonymousRecordDecl.
*/
func (c Cursor) IsAnonymousRecordDecl() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isAnonymousRecordDecl,
		ptr_clang_Cursor_isAnonymousRecordDecl,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isAnonymousRecordDecl", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the cursor specifies a Record member that is a bit-field.
/*
Wraps the C function clang_Cursor_isBitField.
*/
func (c Cursor) IsBitField() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isBitField,
		ptr_clang_Cursor_isBitField,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isBitField", err))
	}

	ret := retC
	return ret
}

/*
Given a cursor pointing to a C++ method call or an Objective-C message, returns non-zero if the method/message is "dynamic", meaning:

For a C++ method: the call is virtual. For an Objective-C message: the receiver is an object instance, not 'super' or a specific class.

If the method/message is "static" or the cursor does not point to a method/message, it will return zero.
*/
/*
Wraps the C function clang_Cursor_isDynamicCall.
*/
func (c Cursor) IsDynamicCall() int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isDynamicCall,
		ptr_clang_Cursor_isDynamicCall,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isDynamicCall", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the given cursor points to a symbol marked with external_source_symbol attribute.
/*
Wraps the C function clang_Cursor_isExternalSymbol.
*/
func (c Cursor) IsExternalSymbol() (string, string, uint32, uint32) {
	c_c := c
	var language string_
	c_language := &language
	var definedIn string_
	c_definedIn := &definedIn
	var isGenerated uint32
	c_isGenerated := &isGenerated

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_language),
		unsafe.Pointer(&c_definedIn),
		unsafe.Pointer(&c_isGenerated),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isExternalSymbol,
		ptr_clang_Cursor_isExternalSymbol,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isExternalSymbol", err))
	}

	ret := retC
	return language.CString(), definedIn.CString(), isGenerated, ret
}

// Determine whether a  CXCursor that is a function declaration, is an inline declaration.
/*
Wraps the C function clang_Cursor_isFunctionInlined.
*/
func (c Cursor) IsFunctionInlined() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isFunctionInlined,
		ptr_clang_Cursor_isFunctionInlined,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isFunctionInlined", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, check if the assembly block has goto labels. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
/*
Wraps the C function clang_Cursor_isGCCAssemblyHasGoto.
*/
func (p0 Cursor) IsGCCAssemblyHasGoto() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isGCCAssemblyHasGoto,
		ptr_clang_Cursor_isGCCAssemblyHasGoto,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isGCCAssemblyHasGoto", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, check if the inline assembly is `volatile`. This function returns 0 if the cursor does not point at a GCC inline assembly block.
/*
Wraps the C function clang_Cursor_isGCCAssemblyVolatile.
*/
func (cursor Cursor) IsGCCAssemblyVolatile() uint32 {
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isGCCAssemblyVolatile,
		ptr_clang_Cursor_isGCCAssemblyVolatile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isGCCAssemblyVolatile", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor represents an inline namespace declaration.
/*
Wraps the C function clang_Cursor_isInlineNamespace.
*/
func (c Cursor) IsInlineNamespace() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isInlineNamespace,
		ptr_clang_Cursor_isInlineNamespace,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isInlineNamespace", err))
	}

	ret := retC
	return ret
}

// Determine whether a  CXCursor that is a macro, is a builtin one.
/*
Wraps the C function clang_Cursor_isMacroBuiltin.
*/
func (c Cursor) IsMacroBuiltin() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isMacroBuiltin,
		ptr_clang_Cursor_isMacroBuiltin,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isMacroBuiltin", err))
	}

	ret := retC
	return ret
}

// Determine whether a  CXCursor that is a macro, is function like.
/*
Wraps the C function clang_Cursor_isMacroFunctionLike.
*/
func (c Cursor) IsMacroFunctionLike() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isMacroFunctionLike,
		ptr_clang_Cursor_isMacroFunctionLike,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isMacroFunctionLike", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if cursor is null.
/*
Wraps the C function clang_Cursor_isNull.
*/
func (cursor Cursor) IsNull() int32 {
	c_cursor := cursor

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isNull,
		ptr_clang_Cursor_isNull,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isNull", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents an Objective-C method or property declaration, return non-zero if the declaration was affected by "\@optional". Returns zero if the cursor is not such a declaration or it is "\@required".
/*
Wraps the C function clang_Cursor_isObjCOptional.
*/
func (c Cursor) IsObjCOptional() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isObjCOptional,
		ptr_clang_Cursor_isObjCOptional,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isObjCOptional", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the given cursor is a variadic function or method.
/*
Wraps the C function clang_Cursor_isVariadic.
*/
func (c Cursor) IsVariadic() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_Cursor_isVariadic,
		ptr_clang_Cursor_isVariadic,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isVariadic", err))
	}

	ret := retC
	return ret
}

// Determine if an enum declaration refers to a scoped enum.
/*
Wraps the C function clang_EnumDecl_isScoped.
*/
func (c Cursor) EnumDecl_isScoped() uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_EnumDecl_isScoped,
		ptr_clang_EnumDecl_isScoped,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EnumDecl_isScoped", err))
	}

	ret := retC
	return ret
}

// Disposes the created Eval memory.
/*
Wraps the C function clang_EvalResult_dispose.
*/
func (e EvalResult) Dispose() {
	c_e := e

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_dispose,
		ptr_clang_EvalResult_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_dispose", err))
	}
}

// Returns the evaluation result as double if the kind is double.
/*
Wraps the C function clang_EvalResult_getAsDouble.
*/
func (e EvalResult) AsDouble() float64 {
	c_e := e

	var retC float64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_getAsDouble,
		ptr_clang_EvalResult_getAsDouble,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getAsDouble", err))
	}

	ret := retC
	return ret
}

// Returns the evaluation result as integer if the kind is Int.
/*
Wraps the C function clang_EvalResult_getAsInt.
*/
func (e EvalResult) AsInt() int32 {
	c_e := e

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_getAsInt,
		ptr_clang_EvalResult_getAsInt,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getAsInt", err))
	}

	ret := retC
	return ret
}

// Returns the evaluation result as a long long integer if the kind is Int. This prevents overflows that may happen if the result is returned with clang_EvalResult_getAsInt.
/*
Wraps the C function clang_EvalResult_getAsLongLong.
*/
func (e EvalResult) AsLongLong() int64 {
	c_e := e

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_getAsLongLong,
		ptr_clang_EvalResult_getAsLongLong,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getAsLongLong", err))
	}

	ret := retC
	return ret
}

// Returns the evaluation result as a constant string if the kind is other than Int or float. User must not free this pointer, instead call clang_EvalResult_dispose on the CXEvalResult returned by clang_Cursor_Evaluate.
/*
Wraps the C function clang_EvalResult_getAsStr.
*/
func (e EvalResult) AsStr() string {
	c_e := e

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_getAsStr,
		ptr_clang_EvalResult_getAsStr,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getAsStr", err))
	}

	ret := libc.GoString(retC)
	return ret
}

// Returns the evaluation result as an unsigned integer if the kind is Int and clang_EvalResult_isUnsignedInt is non-zero.
/*
Wraps the C function clang_EvalResult_getAsUnsigned.
*/
func (e EvalResult) AsUnsigned() uint64 {
	c_e := e

	var retC uint64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_getAsUnsigned,
		ptr_clang_EvalResult_getAsUnsigned,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getAsUnsigned", err))
	}

	ret := retC
	return ret
}

// Returns the kind of the evaluated result.
/*
Wraps the C function clang_EvalResult_getKind.
*/
func (e EvalResult) Kind() EvalResultKind {
	c_e := e

	var retC EvalResultKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_getKind,
		ptr_clang_EvalResult_getKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getKind", err))
	}

	ret := retC
	return ret
}

// Returns a non-zero value if the kind is Int and the evaluation result resulted in an unsigned integer.
/*
Wraps the C function clang_EvalResult_isUnsignedInt.
*/
func (e EvalResult) IsUnsignedInt() uint32 {
	c_e := e

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_EvalResult_isUnsignedInt,
		ptr_clang_EvalResult_isUnsignedInt,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_isUnsignedInt", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the file1 and file2 point to the same file, or they are both NULL.
/*
Wraps the C function clang_File_isEqual.
*/
func (file1 File) IsEqual(file2 File) int32 {
	c_file1 := file1
	c_file2 := file2

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_file1.ptr),
		unsafe.Pointer(&c_file2),
	}

	err := ffi.CallFunction(
		cif_clang_File_isEqual,
		ptr_clang_File_isEqual,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_File_isEqual", err))
	}

	ret := retC
	return ret
}

/*
Returns the real path name of file.

An empty string may be returned. Use clang_getFileName() in that case.
*/
/*
Wraps the C function clang_File_tryGetRealPathName.
*/
func (file File) TryGetRealPathName() string {
	c_file := file

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_file.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_File_tryGetRealPathName,
		ptr_clang_File_tryGetRealPathName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_File_tryGetRealPathName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Convert a given full parsed comment to an HTML fragment.

Specific details of HTML layout are subject to change.  Don't try to parse this HTML back into an AST, use other APIs instead.

Currently the following CSS classes are used:



Function argument documentation is rendered as a <dl> list with arguments sorted in function prototype order.  CSS classes used:

Template parameter documentation is rendered as a <dl> list with parameters sorted in template parameter list order.  CSS classes used:
*/
/*
Wraps the C function clang_FullComment_getAsHTML.
*/
func (comment Comment) FullCommentAsHTML() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_FullComment_getAsHTML,
		ptr_clang_FullComment_getAsHTML,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_FullComment_getAsHTML", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Convert a given full parsed comment to an XML document.

A Relax NG schema for the XML can be found in comment-xml-schema.rng file inside clang source tree.
*/
/*
Wraps the C function clang_FullComment_getAsXML.
*/
func (comment Comment) FullCommentAsXML() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_FullComment_getAsXML,
		ptr_clang_FullComment_getAsXML,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_FullComment_getAsXML", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_HTMLStartTagComment_isSelfClosing.
*/
func (comment Comment) HTMLStartTagComment_isSelfClosing() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_HTMLStartTagComment_isSelfClosing,
		ptr_clang_HTMLStartTagComment_isSelfClosing,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_HTMLStartTagComment_isSelfClosing", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_HTMLStartTag_getAttrName.
*/
func (comment Comment) HTMLStartTagAttrName(attrIdx uint32) string {
	c_comment := comment
	c_attrIdx := attrIdx

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
		unsafe.Pointer(&c_attrIdx),
	}

	err := ffi.CallFunction(
		cif_clang_HTMLStartTag_getAttrName,
		ptr_clang_HTMLStartTag_getAttrName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_HTMLStartTag_getAttrName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_HTMLStartTag_getAttrValue.
*/
func (comment Comment) HTMLStartTagAttrValue(attrIdx uint32) string {
	c_comment := comment
	c_attrIdx := attrIdx

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
		unsafe.Pointer(&c_attrIdx),
	}

	err := ffi.CallFunction(
		cif_clang_HTMLStartTag_getAttrValue,
		ptr_clang_HTMLStartTag_getAttrValue,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_HTMLStartTag_getAttrValue", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_HTMLStartTag_getNumAttrs.
*/
func (comment Comment) HTMLStartTagNumAttrs() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_HTMLStartTag_getNumAttrs,
		ptr_clang_HTMLStartTag_getNumAttrs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_HTMLStartTag_getNumAttrs", err))
	}

	ret := retC
	return ret
}

// Convert an HTML tag AST node to string.
/*
Wraps the C function clang_HTMLTagComment_getAsString.
*/
func (comment Comment) HTMLTagCommentAsString() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_HTMLTagComment_getAsString,
		ptr_clang_HTMLTagComment_getAsString,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_HTMLTagComment_getAsString", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_HTMLTagComment_getTagName.
*/
func (comment Comment) HTMLTagCommentTagName() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_HTMLTagComment_getTagName,
		ptr_clang_HTMLTagComment_getTagName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_HTMLTagComment_getTagName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// An indexing action/session, to be applied to one or multiple translation units.
/*
Wraps the C function clang_IndexAction_create.
*/
func (cIdx Index) Action_create() IndexAction {
	c_cIdx := cIdx

	var retC IndexAction
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_IndexAction_create,
		ptr_clang_IndexAction_create,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_IndexAction_create", err))
	}

	ret := retC
	return ret
}

/*
Destroy the given index action.

The index action must not be destroyed until all of the translation units created within that index action have been destroyed.
*/
/*
Wraps the C function clang_IndexAction_dispose.
*/
func (p0 IndexAction) Dispose() {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_IndexAction_dispose,
		ptr_clang_IndexAction_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_IndexAction_dispose", err))
	}
}

//
/*
Wraps the C function clang_InlineCommandComment_getArgText.
*/
func (comment Comment) InlineCommandCommentArgText(argIdx uint32) string {
	c_comment := comment
	c_argIdx := argIdx

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
		unsafe.Pointer(&c_argIdx),
	}

	err := ffi.CallFunction(
		cif_clang_InlineCommandComment_getArgText,
		ptr_clang_InlineCommandComment_getArgText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_InlineCommandComment_getArgText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_InlineCommandComment_getCommandName.
*/
func (comment Comment) InlineCommandCommentCommandName() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_InlineCommandComment_getCommandName,
		ptr_clang_InlineCommandComment_getCommandName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_InlineCommandComment_getCommandName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_InlineCommandComment_getNumArgs.
*/
func (comment Comment) InlineCommandCommentNumArgs() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_InlineCommandComment_getNumArgs,
		ptr_clang_InlineCommandComment_getNumArgs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_InlineCommandComment_getNumArgs", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_InlineCommandComment_getRenderKind.
*/
func (comment Comment) InlineCommandCommentRenderKind() CommentInlineCommandRenderKind {
	c_comment := comment

	var retC CommentInlineCommandRenderKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_InlineCommandComment_getRenderKind,
		ptr_clang_InlineCommandComment_getRenderKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_InlineCommandComment_getRenderKind", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_InlineContentComment_hasTrailingNewline.
*/
func (comment Comment) InlineContentComment_hasTrailingNewline() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_InlineContentComment_hasTrailingNewline,
		ptr_clang_InlineContentComment_hasTrailingNewline,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_InlineContentComment_hasTrailingNewline", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the given source location is in the main file of the corresponding translation unit.
/*
Wraps the C function clang_Location_isFromMainFile.
*/
func (location SourceLocation) Location_isFromMainFile() int32 {
	c_location := location

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
	}

	err := ffi.CallFunction(
		cif_clang_Location_isFromMainFile,
		ptr_clang_Location_isFromMainFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Location_isFromMainFile", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the given source location is in a system header.
/*
Wraps the C function clang_Location_isInSystemHeader.
*/
func (location SourceLocation) Location_isInSystemHeader() int32 {
	c_location := location

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
	}

	err := ffi.CallFunction(
		cif_clang_Location_isInSystemHeader,
		ptr_clang_Location_isInSystemHeader,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Location_isInSystemHeader", err))
	}

	ret := retC
	return ret
}

// not supported : clang_ModuleCache_prune : param PruneInterval : time_t

// Create a CXModuleMapDescriptor object. Must be disposed with clang_ModuleMapDescriptor_dispose().
/*
Wraps the C function clang_ModuleMapDescriptor_create.
*/
func ModuleMapDescriptor_create(options uint32) ModuleMapDescriptor {
	c_options := options

	var retC ModuleMapDescriptor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_ModuleMapDescriptor_create,
		ptr_clang_ModuleMapDescriptor_create,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_create", err))
	}

	ret := retC
	return ret
}

// Dispose a CXModuleMapDescriptor object.
/*
Wraps the C function clang_ModuleMapDescriptor_dispose.
*/
func (p0 ModuleMapDescriptor) Dispose() {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_ModuleMapDescriptor_dispose,
		ptr_clang_ModuleMapDescriptor_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_dispose", err))
	}
}

// Sets the framework module name that the module.modulemap describes.
/*
Wraps the C function clang_ModuleMapDescriptor_setFrameworkModuleName.
*/
func (p0 ModuleMapDescriptor) SetFrameworkModuleName(name string) ErrorCode {
	c_p0 := p0
	c_name, free_c_name := libc.CString(name)
	defer free_c_name()

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_name),
	}

	err := ffi.CallFunction(
		cif_clang_ModuleMapDescriptor_setFrameworkModuleName,
		ptr_clang_ModuleMapDescriptor_setFrameworkModuleName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_setFrameworkModuleName", err))
	}

	ret := retC
	return ret
}

// Sets the umbrella header name that the module.modulemap describes.
/*
Wraps the C function clang_ModuleMapDescriptor_setUmbrellaHeader.
*/
func (p0 ModuleMapDescriptor) SetUmbrellaHeader(name string) ErrorCode {
	c_p0 := p0
	c_name, free_c_name := libc.CString(name)
	defer free_c_name()

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_name),
	}

	err := ffi.CallFunction(
		cif_clang_ModuleMapDescriptor_setUmbrellaHeader,
		ptr_clang_ModuleMapDescriptor_setUmbrellaHeader,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_setUmbrellaHeader", err))
	}

	ret := retC
	return ret
}

// not supported : clang_ModuleMapDescriptor_writeToBuffer : param out_buffer_ptr : char **

//
/*
Wraps the C function clang_Module_getASTFile.
*/
func (module Module) ASTFile() File {
	c_module := module

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_Module_getASTFile,
		ptr_clang_Module_getASTFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getASTFile", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Module_getFullName.
*/
func (module Module) FullName() string {
	c_module := module

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_Module_getFullName,
		ptr_clang_Module_getFullName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getFullName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_Module_getName.
*/
func (module Module) Name() string {
	c_module := module

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_Module_getName,
		ptr_clang_Module_getName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_Module_getNumTopLevelHeaders.
*/
func (p0 TranslationUnit) ModuleNumTopLevelHeaders(module Module) uint32 {
	c_p0 := p0
	c_module := module

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(
		cif_clang_Module_getNumTopLevelHeaders,
		ptr_clang_Module_getNumTopLevelHeaders,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getNumTopLevelHeaders", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Module_getParent.
*/
func (module Module) Parent() Module {
	c_module := module

	var retC Module
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_Module_getParent,
		ptr_clang_Module_getParent,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getParent", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Module_getTopLevelHeader.
*/
func (p0 TranslationUnit) ModuleTopLevelHeader(module Module, index uint32) File {
	c_p0 := p0
	c_module := module
	c_index := index

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_module),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(
		cif_clang_Module_getTopLevelHeader,
		ptr_clang_Module_getTopLevelHeader,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getTopLevelHeader", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_Module_isSystem.
*/
func (module Module) IsSystem() int32 {
	c_module := module

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_Module_isSystem,
		ptr_clang_Module_isSystem,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_isSystem", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_ParamCommandComment_getDirection.
*/
func (comment Comment) ParamCommandCommentDirection() CommentParamPassDirection {
	c_comment := comment

	var retC CommentParamPassDirection
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_ParamCommandComment_getDirection,
		ptr_clang_ParamCommandComment_getDirection,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ParamCommandComment_getDirection", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_ParamCommandComment_getParamIndex.
*/
func (comment Comment) ParamCommandCommentParamIndex() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_ParamCommandComment_getParamIndex,
		ptr_clang_ParamCommandComment_getParamIndex,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ParamCommandComment_getParamIndex", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_ParamCommandComment_getParamName.
*/
func (comment Comment) ParamCommandCommentParamName() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_ParamCommandComment_getParamName,
		ptr_clang_ParamCommandComment_getParamName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ParamCommandComment_getParamName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_ParamCommandComment_isDirectionExplicit.
*/
func (comment Comment) ParamCommandComment_isDirectionExplicit() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_ParamCommandComment_isDirectionExplicit,
		ptr_clang_ParamCommandComment_isDirectionExplicit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ParamCommandComment_isDirectionExplicit", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_ParamCommandComment_isParamIndexValid.
*/
func (comment Comment) ParamCommandComment_isParamIndexValid() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_ParamCommandComment_isParamIndexValid,
		ptr_clang_ParamCommandComment_isParamIndexValid,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ParamCommandComment_isParamIndexValid", err))
	}

	ret := retC
	return ret
}

// Release a printing policy.
/*
Wraps the C function clang_PrintingPolicy_dispose.
*/
func (policy PrintingPolicy) Dispose() {
	c_policy := policy

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_policy.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_PrintingPolicy_dispose,
		ptr_clang_PrintingPolicy_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_PrintingPolicy_dispose", err))
	}
}

// Get a property value for the given printing policy.
/*
Wraps the C function clang_PrintingPolicy_getProperty.
*/
func (policy PrintingPolicy) Property(property PrintingPolicyProperty) uint32 {
	c_policy := policy
	c_property := property

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_policy.ptr),
		unsafe.Pointer(&c_property),
	}

	err := ffi.CallFunction(
		cif_clang_PrintingPolicy_getProperty,
		ptr_clang_PrintingPolicy_getProperty,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_PrintingPolicy_getProperty", err))
	}

	ret := retC
	return ret
}

// Set a property value for the given printing policy.
/*
Wraps the C function clang_PrintingPolicy_setProperty.
*/
func (policy PrintingPolicy) SetProperty(property PrintingPolicyProperty, value uint32) {
	c_policy := policy
	c_property := property
	c_value := value

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_policy.ptr),
		unsafe.Pointer(&c_property),
		unsafe.Pointer(&c_value),
	}

	err := ffi.CallFunction(
		cif_clang_PrintingPolicy_setProperty,
		ptr_clang_PrintingPolicy_setProperty,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_PrintingPolicy_setProperty", err))
	}
}

// Returns non-zero if range is null.
/*
Wraps the C function clang_Range_isNull.
*/
func (range_ SourceRange) Range_isNull() int32 {
	c_range_ := range_

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(
		cif_clang_Range_isNull,
		ptr_clang_Range_isNull,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Range_isNull", err))
	}

	ret := retC
	return ret
}

/*
For example,

for C and TT nesting depth is 0, for T nesting depth is 1.
*/
/*
Wraps the C function clang_TParamCommandComment_getDepth.
*/
func (comment Comment) TParamCommandCommentDepth() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_TParamCommandComment_getDepth,
		ptr_clang_TParamCommandComment_getDepth,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TParamCommandComment_getDepth", err))
	}

	ret := retC
	return ret
}

/*
For example,

for C and TT nesting depth is 0, so we can ask for index at depth 0: at depth 0 C's index is 0, TT's index is 1.

For T nesting depth is 1, so we can ask for index at depth 0 and 1: at depth 0 T's index is 1 (same as TT's), at depth 1 T's index is 0.
*/
/*
Wraps the C function clang_TParamCommandComment_getIndex.
*/
func (comment Comment) TParamCommandCommentIndex(depth uint32) uint32 {
	c_comment := comment
	c_depth := depth

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
		unsafe.Pointer(&c_depth),
	}

	err := ffi.CallFunction(
		cif_clang_TParamCommandComment_getIndex,
		ptr_clang_TParamCommandComment_getIndex,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TParamCommandComment_getIndex", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_TParamCommandComment_getParamName.
*/
func (comment Comment) TParamCommandCommentParamName() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_TParamCommandComment_getParamName,
		ptr_clang_TParamCommandComment_getParamName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TParamCommandComment_getParamName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_TParamCommandComment_isParamPositionValid.
*/
func (comment Comment) TParamCommandComment_isParamPositionValid() uint32 {
	c_comment := comment

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_TParamCommandComment_isParamPositionValid,
		ptr_clang_TParamCommandComment_isParamPositionValid,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TParamCommandComment_isParamPositionValid", err))
	}

	ret := retC
	return ret
}

// Destroy the CXTargetInfo object.
/*
Wraps the C function clang_TargetInfo_dispose.
*/
func (info TargetInfo) Dispose() {
	c_info := info

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_info.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_TargetInfo_dispose,
		ptr_clang_TargetInfo_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TargetInfo_dispose", err))
	}
}

/*
Get the pointer width of the target in bits.

Returns -1 in case of error.
*/
/*
Wraps the C function clang_TargetInfo_getPointerWidth.
*/
func (info TargetInfo) PointerWidth() int32 {
	c_info := info

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_info.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_TargetInfo_getPointerWidth,
		ptr_clang_TargetInfo_getPointerWidth,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TargetInfo_getPointerWidth", err))
	}

	ret := retC
	return ret
}

/*
Get the normalized target triple as a string.

Returns the empty string in case of any error.
*/
/*
Wraps the C function clang_TargetInfo_getTriple.
*/
func (info TargetInfo) Triple() string {
	c_info := info

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_info.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_TargetInfo_getTriple,
		ptr_clang_TargetInfo_getTriple,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TargetInfo_getTriple", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_TextComment_getText.
*/
func (comment Comment) TextCommentText() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_TextComment_getText,
		ptr_clang_TextComment_getText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TextComment_getText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Return the alignment of a type in bytes as per C++[expr.alignof]   standard.

If the type declaration is invalid, CXTypeLayoutError_Invalid is returned. If the type declaration is an incomplete type, CXTypeLayoutError_Incomplete   is returned. If the type declaration is a dependent type, CXTypeLayoutError_Dependent is   returned. If the type declaration is not a constant size type,   CXTypeLayoutError_NotConstantSize is returned.
*/
/*
Wraps the C function clang_Type_getAlignOf.
*/
func (t Type) AlignOf() int64 {
	c_t := t

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getAlignOf,
		ptr_clang_Type_getAlignOf,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getAlignOf", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the ref-qualifier kind of a function or method.

The ref-qualifier is returned for C++ functions or methods. For other types or non-C++ declarations, CXRefQualifier_None is returned.
*/
/*
Wraps the C function clang_Type_getCXXRefQualifier.
*/
func (t Type) XRefQualifier() RefQualifierKind {
	c_t := t

	var retC RefQualifierKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getCXXRefQualifier,
		ptr_clang_Type_getCXXRefQualifier,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getCXXRefQualifier", err))
	}

	ret := retC
	return ret
}

/*
Return the class type of an member pointer type.

If a non-member-pointer type is passed in, an invalid type is returned.
*/
/*
Wraps the C function clang_Type_getClassType.
*/
func (t Type) ClassType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getClassType,
		ptr_clang_Type_getClassType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getClassType", err))
	}

	ret := retC
	return ret
}

/*
Return the type that was modified by this attributed type.

If the type is not an attributed type, an invalid type is returned.
*/
/*
Wraps the C function clang_Type_getModifiedType.
*/
func (t Type) ModifiedType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getModifiedType,
		ptr_clang_Type_getModifiedType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getModifiedType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the type named by the qualified-id.

If a non-elaborated type is passed in, an invalid type is returned.
*/
/*
Wraps the C function clang_Type_getNamedType.
*/
func (t Type) NamedType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getNamedType,
		ptr_clang_Type_getNamedType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNamedType", err))
	}

	ret := retC
	return ret
}

// Retrieve the nullability kind of a pointer type.
/*
Wraps the C function clang_Type_getNullability.
*/
func (t Type) Nullability() TypeNullabilityKind {
	c_t := t

	var retC TypeNullabilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getNullability,
		ptr_clang_Type_getNullability,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNullability", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of protocol references associated with an ObjC object/id.

If the type is not an ObjC object, 0 is returned.
*/
/*
Wraps the C function clang_Type_getNumObjCProtocolRefs.
*/
func (t Type) NumObjCProtocolRefs() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getNumObjCProtocolRefs,
		ptr_clang_Type_getNumObjCProtocolRefs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNumObjCProtocolRefs", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of type arguments associated with an ObjC object.

If the type is not an ObjC object, 0 is returned.
*/
/*
Wraps the C function clang_Type_getNumObjCTypeArgs.
*/
func (t Type) NumObjCTypeArgs() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getNumObjCTypeArgs,
		ptr_clang_Type_getNumObjCTypeArgs,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNumObjCTypeArgs", err))
	}

	ret := retC
	return ret
}

// Returns the number of template arguments for given template specialization, or -1 if type T is not a template specialization.
/*
Wraps the C function clang_Type_getNumTemplateArguments.
*/
func (t Type) NumTemplateArguments() int32 {
	c_t := t

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getNumTemplateArguments,
		ptr_clang_Type_getNumTemplateArguments,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNumTemplateArguments", err))
	}

	ret := retC
	return ret
}

// Returns the Objective-C type encoding for the specified CXType.
/*
Wraps the C function clang_Type_getObjCEncoding.
*/
func (type_ Type) ObjCEncoding() string {
	c_type_ := type_

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_type_),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getObjCEncoding,
		ptr_clang_Type_getObjCEncoding,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCEncoding", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieves the base type of the ObjCObjectType.

If the type is not an ObjC object, an invalid type is returned.
*/
/*
Wraps the C function clang_Type_getObjCObjectBaseType.
*/
func (t Type) ObjCObjectBaseType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getObjCObjectBaseType,
		ptr_clang_Type_getObjCObjectBaseType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCObjectBaseType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the decl for a protocol reference for an ObjC object/id.

If the type is not an ObjC object or there are not enough protocol references, an invalid cursor is returned.
*/
/*
Wraps the C function clang_Type_getObjCProtocolDecl.
*/
func (t Type) ObjCProtocolDecl(i uint32) Cursor {
	c_t := t
	c_i := i

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getObjCProtocolDecl,
		ptr_clang_Type_getObjCProtocolDecl,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCProtocolDecl", err))
	}

	ret := retC
	return ret
}

/*
Retrieve a type argument associated with an ObjC object.

If the type is not an ObjC or the index is not valid, an invalid type is returned.
*/
/*
Wraps the C function clang_Type_getObjCTypeArg.
*/
func (t Type) ObjCTypeArg(i uint32) Type {
	c_t := t
	c_i := i

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getObjCTypeArg,
		ptr_clang_Type_getObjCTypeArg,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCTypeArg", err))
	}

	ret := retC
	return ret
}

/*
Return the offset of a field named S in a record of type T in bits   as it would be returned by __offsetof__ as per C++11[18.2p4]

If the cursor is not a record field declaration, CXTypeLayoutError_Invalid   is returned. If the field's type declaration is an incomplete type,   CXTypeLayoutError_Incomplete is returned. If the field's type declaration is a dependent type,   CXTypeLayoutError_Dependent is returned. If the field's name S is not found,   CXTypeLayoutError_InvalidFieldName is returned.
*/
/*
Wraps the C function clang_Type_getOffsetOf.
*/
func (t Type) OffsetOf(s string) int64 {
	c_t := t
	c_s, free_c_s := libc.CString(s)
	defer free_c_s()

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_s),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getOffsetOf,
		ptr_clang_Type_getOffsetOf,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getOffsetOf", err))
	}

	ret := retC
	return ret
}

/*
Return the size of a type in bytes as per C++[expr.sizeof] standard.

If the type declaration is invalid, CXTypeLayoutError_Invalid is returned. If the type declaration is an incomplete type, CXTypeLayoutError_Incomplete   is returned. If the type declaration is a dependent type, CXTypeLayoutError_Dependent is   returned.
*/
/*
Wraps the C function clang_Type_getSizeOf.
*/
func (t Type) SizeOf() int64 {
	c_t := t

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getSizeOf,
		ptr_clang_Type_getSizeOf,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getSizeOf", err))
	}

	ret := retC
	return ret
}

/*
Returns the type template argument of a template class specialization at given index.

This function only returns template type arguments and does not handle template template arguments or variadic packs.
*/
/*
Wraps the C function clang_Type_getTemplateArgumentAsType.
*/
func (t Type) TemplateArgumentAsType(i uint32) Type {
	c_t := t
	c_i := i

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getTemplateArgumentAsType,
		ptr_clang_Type_getTemplateArgumentAsType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getTemplateArgumentAsType", err))
	}

	ret := retC
	return ret
}

/*
Gets the type contained by this atomic type.

If a non-atomic type is passed in, an invalid type is returned.
*/
/*
Wraps the C function clang_Type_getValueType.
*/
func (cT Type) ValueType() Type {
	c_cT := cT

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(
		cif_clang_Type_getValueType,
		ptr_clang_Type_getValueType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getValueType", err))
	}

	ret := retC
	return ret
}

/*
Determine if a typedef is 'transparent' tag.

A typedef is considered 'transparent' if it shares a name and spelling location with its underlying tag type, as is the case with the NS_ENUM macro.
*/
/*
Wraps the C function clang_Type_isTransparentTagTypedef.
*/
func (t Type) IsTransparentTagTypedef() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_Type_isTransparentTagTypedef,
		ptr_clang_Type_isTransparentTagTypedef,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_isTransparentTagTypedef", err))
	}

	ret := retC
	return ret
}

/*
Visit the fields of a particular type.

This function visits all the direct fields of the given cursor, invoking the given visitor function with the cursors of each visited field. The traversal may be ended prematurely, if the visitor returns CXFieldVisit_Break.
*/
/*
Wraps the C function clang_Type_visitFields.
*/
func (t Type) VisitFields(visitor FieldVisitor, client_data ClientData) uint32 {
	c_t := t
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(
		cif_clang_Type_visitFields,
		ptr_clang_Type_visitFields,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_visitFields", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_VerbatimBlockLineComment_getText.
*/
func (comment Comment) VerbatimBlockLineCommentText() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_VerbatimBlockLineComment_getText,
		ptr_clang_VerbatimBlockLineComment_getText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VerbatimBlockLineComment_getText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

//
/*
Wraps the C function clang_VerbatimLineComment_getText.
*/
func (comment Comment) VerbatimLineCommentText() string {
	c_comment := comment

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_comment),
	}

	err := ffi.CallFunction(
		cif_clang_VerbatimLineComment_getText,
		ptr_clang_VerbatimLineComment_getText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VerbatimLineComment_getText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Map an absolute virtual file path to an absolute real one. The virtual path must be canonicalized (not contain "."/"..").
/*
Wraps the C function clang_VirtualFileOverlay_addFileMapping.
*/
func (p0 VirtualFileOverlay) AddFileMapping(virtualPath string, realPath string) ErrorCode {
	c_p0 := p0
	c_virtualPath, free_c_virtualPath := libc.CString(virtualPath)
	defer free_c_virtualPath()
	c_realPath, free_c_realPath := libc.CString(realPath)
	defer free_c_realPath()

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_virtualPath),
		unsafe.Pointer(&c_realPath),
	}

	err := ffi.CallFunction(
		cif_clang_VirtualFileOverlay_addFileMapping,
		ptr_clang_VirtualFileOverlay_addFileMapping,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_addFileMapping", err))
	}

	ret := retC
	return ret
}

// Create a CXVirtualFileOverlay object. Must be disposed with clang_VirtualFileOverlay_dispose().
/*
Wraps the C function clang_VirtualFileOverlay_create.
*/
func VirtualFileOverlay_create(options uint32) VirtualFileOverlay {
	c_options := options

	var retC VirtualFileOverlay
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_VirtualFileOverlay_create,
		ptr_clang_VirtualFileOverlay_create,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_create", err))
	}

	ret := retC
	return ret
}

// Dispose a CXVirtualFileOverlay object.
/*
Wraps the C function clang_VirtualFileOverlay_dispose.
*/
func (p0 VirtualFileOverlay) Dispose() {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_VirtualFileOverlay_dispose,
		ptr_clang_VirtualFileOverlay_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_dispose", err))
	}
}

// Set the case sensitivity for the CXVirtualFileOverlay object. The CXVirtualFileOverlay object is case-sensitive by default, this option can be used to override the default.
/*
Wraps the C function clang_VirtualFileOverlay_setCaseSensitivity.
*/
func (p0 VirtualFileOverlay) SetCaseSensitivity(caseSensitive int32) ErrorCode {
	c_p0 := p0
	c_caseSensitive := caseSensitive

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_caseSensitive),
	}

	err := ffi.CallFunction(
		cif_clang_VirtualFileOverlay_setCaseSensitivity,
		ptr_clang_VirtualFileOverlay_setCaseSensitivity,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_setCaseSensitivity", err))
	}

	ret := retC
	return ret
}

// not supported : clang_VirtualFileOverlay_writeToBuffer : param out_buffer_ptr : char **

/*
Annotate the given set of tokens by providing cursors for each token that can be mapped to a specific entity within the abstract syntax tree.

This token-annotation routine is equivalent to invoking clang_getCursor() for the source locations of each of the tokens. The cursors provided are filtered, so that only those cursors that have a direct correspondence to the token are accepted. For example, given a function call f(x), clang_getCursor() would provide the following cursors:

* when the cursor is over the 'f', a DeclRefExpr cursor referring to 'f'.   * when the cursor is over the '(' or the ')', a CallExpr referring to 'f'.   * when the cursor is over the 'x', a DeclRefExpr cursor referring to 'x'.

Only the first and last of these cursors will occur within the annotate, since the tokens "f" and "x' directly refer to a function and a variable, respectively, but the parentheses are just a small part of the full syntax of the function call expression, which is not provided as an annotation.
*/
/*
Wraps the C function clang_annotateTokens.
*/
func (tU TranslationUnit) AnnotateTokens(numTokens uint32) (Token, Cursor) {
	c_tU := tU
	var tokens Token
	c_tokens := &tokens
	c_numTokens := numTokens
	var cursors Cursor
	c_cursors := &cursors

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_tokens),
		unsafe.Pointer(&c_numTokens),
		unsafe.Pointer(&c_cursors),
	}

	err := ffi.CallFunction(
		cif_clang_annotateTokens,
		ptr_clang_annotateTokens,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_annotateTokens", err))
	}
	return tokens, cursors
}

/*
Perform code completion at a given location in a translation unit.

This function performs code completion at a particular file, line, and column within source code, providing results that suggest potential code snippets based on the context of the completion. The basic model for code completion is that Clang will parse a complete source file, performing syntax checking up to the location where code-completion has been requested. At that point, a special code-completion token is passed to the parser, which recognizes this token and determines, based on the current location in the C/Objective-C/C++ grammar and the state of semantic analysis, what completions to provide. These completions are returned via a new CXCodeCompleteResults structure.

Code completion itself is meant to be triggered by the client when the user types punctuation characters or whitespace, at which point the code-completion location will coincide with the cursor. For example, if p is a pointer, code-completion might be triggered after the "-" and then after the ">" in p->. When the code-completion location is after the ">", the completion results will provide, e.g., the members of the struct that "p" points to. The client is responsible for placing the cursor at the beginning of the token currently being typed, then filtering the results based on the contents of the token. For example, when code-completing for the expression p->get, the client should provide the location just after the ">" (e.g., pointing at the "g") to this code-completion hook. Then, the client can filter the results based on the current token text ("get"), only showing those results that start with "get". The intent of this interface is to separate the relatively high-latency acquisition of code-completion results from the filtering of results on a per-character basis, which must have a lower latency.
*/
/*
Wraps the C function clang_codeCompleteAt.
*/
func (tU TranslationUnit) CodeCompleteAt(complete_filename string, complete_line uint32, complete_column uint32, unsaved_files []UnsavedFile, options uint32) *CodeCompleteResults {
	c_tU := tU
	c_complete_filename, free_c_complete_filename := libc.CString(complete_filename)
	defer free_c_complete_filename()
	c_complete_line := complete_line
	c_complete_column := complete_column
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_num_unsaved_files := len(unsaved_files)
	c_options := options

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_complete_filename),
		unsafe.Pointer(&c_complete_line),
		unsafe.Pointer(&c_complete_column),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteAt,
		ptr_clang_codeCompleteAt,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteAt", err))
	}

	ret := (*CodeCompleteResults)(retC)
	return ret
}

// Returns the cursor kind for the container for the current code completion context. The container is only guaranteed to be set for contexts where a container exists (i.e. member accesses or Objective-C message sends); if there is not a container, this function will return CXCursor_InvalidCode.
/*
Wraps the C function clang_codeCompleteGetContainerKind.
*/
func (results *CodeCompleteResults) CodeCompleteGetContainerKind() (uint32, CursorKind) {
	c_results := results
	var isIncomplete uint32
	c_isIncomplete := &isIncomplete

	var retC CursorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
		unsafe.Pointer(&c_isIncomplete),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteGetContainerKind,
		ptr_clang_codeCompleteGetContainerKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteGetContainerKind", err))
	}

	ret := retC
	return isIncomplete, ret
}

// Returns the USR for the container for the current code completion context. If there is not a container for the current context, this function will return the empty string.
/*
Wraps the C function clang_codeCompleteGetContainerUSR.
*/
func (results *CodeCompleteResults) CodeCompleteGetContainerUSR() string {
	c_results := results

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteGetContainerUSR,
		ptr_clang_codeCompleteGetContainerUSR,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteGetContainerUSR", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Determines what completions are appropriate for the context the given code completion.
/*
Wraps the C function clang_codeCompleteGetContexts.
*/
func (results *CodeCompleteResults) CodeCompleteGetContexts() uint64 {
	c_results := results

	var retC uint64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteGetContexts,
		ptr_clang_codeCompleteGetContexts,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteGetContexts", err))
	}

	ret := retC
	return ret
}

// Retrieve a diagnostic associated with the given code completion.
/*
Wraps the C function clang_codeCompleteGetDiagnostic.
*/
func (results *CodeCompleteResults) CodeCompleteGetDiagnostic(index uint32) Diagnostic {
	c_results := results
	c_index := index

	var retC Diagnostic
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteGetDiagnostic,
		ptr_clang_codeCompleteGetDiagnostic,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteGetDiagnostic", err))
	}

	ret := retC
	return ret
}

// Determine the number of diagnostics produced prior to the location where code completion was performed.
/*
Wraps the C function clang_codeCompleteGetNumDiagnostics.
*/
func (results *CodeCompleteResults) CodeCompleteGetNumDiagnostics() uint32 {
	c_results := results

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteGetNumDiagnostics,
		ptr_clang_codeCompleteGetNumDiagnostics,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteGetNumDiagnostics", err))
	}

	ret := retC
	return ret
}

// Returns the currently-entered selector for an Objective-C message send, formatted like "initWithFoo:bar:". Only guaranteed to return a non-empty string for CXCompletionContext_ObjCInstanceMessage and CXCompletionContext_ObjCClassMessage.
/*
Wraps the C function clang_codeCompleteGetObjCSelector.
*/
func (results *CodeCompleteResults) CodeCompleteGetObjCSelector() string {
	c_results := results

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
	}

	err := ffi.CallFunction(
		cif_clang_codeCompleteGetObjCSelector,
		ptr_clang_codeCompleteGetObjCSelector,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_codeCompleteGetObjCSelector", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Construct a USR for a specified Objective-C category.
/*
Wraps the C function clang_constructUSR_ObjCCategory.
*/
func ConstructUSR_ObjCCategory(class_name string, category_name string) string {
	c_class_name, free_c_class_name := libc.CString(class_name)
	defer free_c_class_name()
	c_category_name, free_c_category_name := libc.CString(category_name)
	defer free_c_category_name()

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_class_name),
		unsafe.Pointer(&c_category_name),
	}

	err := ffi.CallFunction(
		cif_clang_constructUSR_ObjCCategory,
		ptr_clang_constructUSR_ObjCCategory,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_constructUSR_ObjCCategory", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Construct a USR for a specified Objective-C class.
/*
Wraps the C function clang_constructUSR_ObjCClass.
*/
func ConstructUSR_ObjCClass(class_name string) string {
	c_class_name, free_c_class_name := libc.CString(class_name)
	defer free_c_class_name()

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_class_name),
	}

	err := ffi.CallFunction(
		cif_clang_constructUSR_ObjCClass,
		ptr_clang_constructUSR_ObjCClass,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_constructUSR_ObjCClass", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Construct a USR for a specified Objective-C instance variable and   the USR for its containing class.
/*
Wraps the C function clang_constructUSR_ObjCIvar.
*/
func ConstructUSR_ObjCIvar(name string, classUSR string_) string {
	c_name, free_c_name := libc.CString(name)
	defer free_c_name()
	c_classUSR := classUSR

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_name),
		unsafe.Pointer(&c_classUSR),
	}

	err := ffi.CallFunction(
		cif_clang_constructUSR_ObjCIvar,
		ptr_clang_constructUSR_ObjCIvar,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_constructUSR_ObjCIvar", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Construct a USR for a specified Objective-C method and   the USR for its containing class.
/*
Wraps the C function clang_constructUSR_ObjCMethod.
*/
func ConstructUSR_ObjCMethod(name string, isInstanceMethod uint32, classUSR string_) string {
	c_name, free_c_name := libc.CString(name)
	defer free_c_name()
	c_isInstanceMethod := isInstanceMethod
	c_classUSR := classUSR

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_name),
		unsafe.Pointer(&c_isInstanceMethod),
		unsafe.Pointer(&c_classUSR),
	}

	err := ffi.CallFunction(
		cif_clang_constructUSR_ObjCMethod,
		ptr_clang_constructUSR_ObjCMethod,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_constructUSR_ObjCMethod", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Construct a USR for a specified Objective-C property and the USR  for its containing class.
/*
Wraps the C function clang_constructUSR_ObjCProperty.
*/
func ConstructUSR_ObjCProperty(property string, classUSR string_) string {
	c_property, free_c_property := libc.CString(property)
	defer free_c_property()
	c_classUSR := classUSR

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_property),
		unsafe.Pointer(&c_classUSR),
	}

	err := ffi.CallFunction(
		cif_clang_constructUSR_ObjCProperty,
		ptr_clang_constructUSR_ObjCProperty,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_constructUSR_ObjCProperty", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Construct a USR for a specified Objective-C protocol.
/*
Wraps the C function clang_constructUSR_ObjCProtocol.
*/
func ConstructUSR_ObjCProtocol(protocol_name string) string {
	c_protocol_name, free_c_protocol_name := libc.CString(protocol_name)
	defer free_c_protocol_name()

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_protocol_name),
	}

	err := ffi.CallFunction(
		cif_clang_constructUSR_ObjCProtocol,
		ptr_clang_constructUSR_ObjCProtocol,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_constructUSR_ObjCProtocol", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Traverses the translation unit to create a CXAPISet.
/*
Wraps the C function clang_createAPISet.
*/
func (tu TranslationUnit) CreateAPISet() (APISet, ErrorCode) {
	c_tu := tu
	var out_api APISet
	c_out_api := &out_api

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_out_api),
	}

	err := ffi.CallFunction(
		cif_clang_createAPISet,
		ptr_clang_createAPISet,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createAPISet", err))
	}

	ret := retC
	return out_api, ret
}

// Creates an empty CXCursorSet.
/*
Wraps the C function clang_createCXCursorSet.
*/
func CreateCXCursorSet() CursorSet {
	var retC CursorSet
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_createCXCursorSet,
		ptr_clang_createCXCursorSet,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createCXCursorSet", err))
	}

	ret := retC
	return ret
}

/*
Provides a shared context for creating translation units.

It provides two options:

- excludeDeclarationsFromPCH: When non-zero, allows enumeration of "local" declarations (when loading any new translation units). A "local" declaration is one that belongs in the translation unit itself and not in a precompiled header that was used by the translation unit. If zero, all declarations will be enumerated.

Here is an example:



This process of creating the 'pch', loading it separately, and using it (via -include-pch) allows 'excludeDeclsFromPCH' to remove redundant callbacks (which gives the indexer the same performance benefit as the compiler).
*/
/*
Wraps the C function clang_createIndex.
*/
func CreateIndex(excludeDeclarationsFromPCH int32, displayDiagnostics int32) Index {
	c_excludeDeclarationsFromPCH := excludeDeclarationsFromPCH
	c_displayDiagnostics := displayDiagnostics

	var retC Index
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_excludeDeclarationsFromPCH),
		unsafe.Pointer(&c_displayDiagnostics),
	}

	err := ffi.CallFunction(
		cif_clang_createIndex,
		ptr_clang_createIndex,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createIndex", err))
	}

	ret := retC
	return ret
}

/*
Provides a shared context for creating translation units.

Call this function instead of clang_createIndex() if you need to configure the additional options in CXIndexOptions.



For example:
*/
/*
Wraps the C function clang_createIndexWithOptions.
*/
func (options *IndexOptions) CreateIndexWithOptions() Index {
	c_options := options

	var retC Index
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_createIndexWithOptions,
		ptr_clang_createIndexWithOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createIndexWithOptions", err))
	}

	ret := retC
	return ret
}

// Same as clang_createTranslationUnit2, but returns the CXTranslationUnit instead of an error code.  In case of an error this routine returns a NULL CXTranslationUnit, without further detailed error codes.
/*
Wraps the C function clang_createTranslationUnit.
*/
func (cIdx Index) CreateTranslationUnit(ast_filename string) TranslationUnit {
	c_cIdx := cIdx
	c_ast_filename, free_c_ast_filename := libc.CString(ast_filename)
	defer free_c_ast_filename()

	var retC TranslationUnit
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
		unsafe.Pointer(&c_ast_filename),
	}

	err := ffi.CallFunction(
		cif_clang_createTranslationUnit,
		ptr_clang_createTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createTranslationUnit", err))
	}

	ret := retC
	return ret
}

// Create a translation unit from an AST file (-emit-ast).
/*
Wraps the C function clang_createTranslationUnit2.
*/
func (cIdx Index) CreateTranslationUnit2(ast_filename string) (TranslationUnit, ErrorCode) {
	c_cIdx := cIdx
	c_ast_filename, free_c_ast_filename := libc.CString(ast_filename)
	defer free_c_ast_filename()
	var out_TU TranslationUnit
	c_out_TU := &out_TU

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
		unsafe.Pointer(&c_ast_filename),
		unsafe.Pointer(&c_out_TU),
	}

	err := ffi.CallFunction(
		cif_clang_createTranslationUnit2,
		ptr_clang_createTranslationUnit2,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createTranslationUnit2", err))
	}

	ret := retC
	return out_TU, ret
}

/*
Return the CXTranslationUnit for a given source file and the provided command line arguments one would pass to the compiler.

Note: The 'source_filename' argument is optional.  If the caller provides a NULL pointer, the name of the source file is expected to reside in the specified command line arguments.

Note: When encountered in 'clang_command_line_args', the following options are ignored:

'-c'   '-emit-ast'   '-fsyntax-only'   '-o <output file>'  (both '-o' and '<output file>' are ignored)
*/
/*
Wraps the C function clang_createTranslationUnitFromSourceFile.
*/
func (cIdx Index) CreateTranslationUnitFromSourceFile(source_filename string, clang_command_line_args []string, unsaved_files []UnsavedFile) TranslationUnit {
	c_cIdx := cIdx
	c_source_filename, free_c_source_filename := libc.CString(source_filename)
	defer free_c_source_filename()
	c_num_clang_command_line_args := len(clang_command_line_args)
	c_clang_command_line_args, free_c_clang_command_line_args := libc.CStrings(clang_command_line_args)
	defer free_c_clang_command_line_args()
	c_num_unsaved_files := len(unsaved_files)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}

	var retC TranslationUnit
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
		unsafe.Pointer(&c_source_filename),
		unsafe.Pointer(&c_num_clang_command_line_args),
		unsafe.Pointer(&c_clang_command_line_args),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_unsaved_files),
	}

	err := ffi.CallFunction(
		cif_clang_createTranslationUnitFromSourceFile,
		ptr_clang_createTranslationUnitFromSourceFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createTranslationUnitFromSourceFile", err))
	}

	ret := retC
	return ret
}

// Returns a default set of code-completion options that can be passed toclang_codeCompleteAt().
/*
Wraps the C function clang_defaultCodeCompleteOptions.
*/
func DefaultCodeCompleteOptions() uint32 {
	var retC uint32
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_defaultCodeCompleteOptions,
		ptr_clang_defaultCodeCompleteOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultCodeCompleteOptions", err))
	}

	ret := retC
	return ret
}

// Retrieve the set of display options most similar to the default behavior of the clang compiler.
/*
Wraps the C function clang_defaultDiagnosticDisplayOptions.
*/
func DefaultDiagnosticDisplayOptions() uint32 {
	var retC uint32
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_defaultDiagnosticDisplayOptions,
		ptr_clang_defaultDiagnosticDisplayOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultDiagnosticDisplayOptions", err))
	}

	ret := retC
	return ret
}

/*
Returns the set of flags that is suitable for parsing a translation unit that is being edited.

The set of flags returned provide options for clang_parseTranslationUnit() to indicate that the translation unit is likely to be reparsed many times, either explicitly (via clang_reparseTranslationUnit()) or implicitly (e.g., by code completion (clang_codeCompletionAt())). The returned flag set contains an unspecified set of optimizations (e.g., the precompiled preamble) geared toward improving the performance of these routines. The set of optimizations enabled may change from one version to the next.
*/
/*
Wraps the C function clang_defaultEditingTranslationUnitOptions.
*/
func DefaultEditingTranslationUnitOptions() uint32 {
	var retC uint32
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_defaultEditingTranslationUnitOptions,
		ptr_clang_defaultEditingTranslationUnitOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultEditingTranslationUnitOptions", err))
	}

	ret := retC
	return ret
}

/*
Returns the set of flags that is suitable for reparsing a translation unit.

The set of flags returned provide options for clang_reparseTranslationUnit() by default. The returned flag set contains an unspecified set of optimizations geared toward common uses of reparsing. The set of optimizations enabled may change from one version to the next.
*/
/*
Wraps the C function clang_defaultReparseOptions.
*/
func (tU TranslationUnit) DefaultReparseOptions() uint32 {
	c_tU := tU

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_defaultReparseOptions,
		ptr_clang_defaultReparseOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultReparseOptions", err))
	}

	ret := retC
	return ret
}

/*
Returns the set of flags that is suitable for saving a translation unit.

The set of flags returned provide options for clang_saveTranslationUnit() by default. The returned flag set contains an unspecified set of options that save translation units with the most commonly-requested data.
*/
/*
Wraps the C function clang_defaultSaveOptions.
*/
func (tU TranslationUnit) DefaultSaveOptions() uint32 {
	c_tU := tU

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_defaultSaveOptions,
		ptr_clang_defaultSaveOptions,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultSaveOptions", err))
	}

	ret := retC
	return ret
}

/*
Dispose of an APISet.

The provided CXAPISet can not be used after this function is called.
*/
/*
Wraps the C function clang_disposeAPISet.
*/
func (api APISet) DisposeAPISet() {
	c_api := api

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_api.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_disposeAPISet,
		ptr_clang_disposeAPISet,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeAPISet", err))
	}
}

// Disposes a CXCursorSet and releases its associated memory.
/*
Wraps the C function clang_disposeCXCursorSet.
*/
func (cset CursorSet) DisposeCXCursorSet() {
	c_cset := cset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cset.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_disposeCXCursorSet,
		ptr_clang_disposeCXCursorSet,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeCXCursorSet", err))
	}
}

// Free the memory associated with a CXPlatformAvailability structure.
/*
Wraps the C function clang_disposeCXPlatformAvailability.
*/
func (availability *PlatformAvailability) DisposeCXPlatformAvailability() {
	c_availability := availability

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_availability),
	}

	err := ffi.CallFunction(
		cif_clang_disposeCXPlatformAvailability,
		ptr_clang_disposeCXPlatformAvailability,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeCXPlatformAvailability", err))
	}
}

//
/*
Wraps the C function clang_disposeCXTUResourceUsage.
*/
func (usage TUResourceUsage) DisposeCXTUResourceUsage() {
	c_usage := usage

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_usage),
	}

	err := ffi.CallFunction(
		cif_clang_disposeCXTUResourceUsage,
		ptr_clang_disposeCXTUResourceUsage,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeCXTUResourceUsage", err))
	}
}

// Free the given set of code-completion results.
/*
Wraps the C function clang_disposeCodeCompleteResults.
*/
func (results *CodeCompleteResults) DisposeCodeCompleteResults() {
	c_results := results

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
	}

	err := ffi.CallFunction(
		cif_clang_disposeCodeCompleteResults,
		ptr_clang_disposeCodeCompleteResults,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeCodeCompleteResults", err))
	}
}

// Destroy a diagnostic.
/*
Wraps the C function clang_disposeDiagnostic.
*/
func (diagnostic Diagnostic) DisposeDiagnostic() {
	c_diagnostic := diagnostic

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_disposeDiagnostic,
		ptr_clang_disposeDiagnostic,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeDiagnostic", err))
	}
}

// Release a CXDiagnosticSet and all of its contained diagnostics.
/*
Wraps the C function clang_disposeDiagnosticSet.
*/
func (diags DiagnosticSet) DisposeDiagnosticSet() {
	c_diags := diags

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diags.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_disposeDiagnosticSet,
		ptr_clang_disposeDiagnosticSet,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeDiagnosticSet", err))
	}
}

/*
Destroy the given index.

The index must not be destroyed until all of the translation units created within that index have been destroyed.
*/
/*
Wraps the C function clang_disposeIndex.
*/
func (index Index) DisposeIndex() {
	c_index := index

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_index.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_disposeIndex,
		ptr_clang_disposeIndex,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeIndex", err))
	}
}

// Free the set of overridden cursors returned by clang_getOverriddenCursors().
/*
Wraps the C function clang_disposeOverriddenCursors.
*/
func (overridden *Cursor) DisposeOverriddenCursors() {
	c_overridden := overridden

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_overridden),
	}

	err := ffi.CallFunction(
		cif_clang_disposeOverriddenCursors,
		ptr_clang_disposeOverriddenCursors,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeOverriddenCursors", err))
	}
}

// Destroy the given CXSourceRangeList.
/*
Wraps the C function clang_disposeSourceRangeList.
*/
func (ranges *SourceRangeList) DisposeSourceRangeList() {
	c_ranges := ranges

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_ranges),
	}

	err := ffi.CallFunction(
		cif_clang_disposeSourceRangeList,
		ptr_clang_disposeSourceRangeList,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeSourceRangeList", err))
	}
}

// Free the given string.
/*
Wraps the C function clang_disposeString.
*/
func (string_ string_) DisposeString() {
	c_string_ := string_

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_string_),
	}

	err := ffi.CallFunction(
		cif_clang_disposeString,
		ptr_clang_disposeString,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeString", err))
	}
}

// Free the given string set.
/*
Wraps the C function clang_disposeStringSet.
*/
func (set *StringSet) DisposeStringSet() {
	c_set := set

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_set),
	}

	err := ffi.CallFunction(
		cif_clang_disposeStringSet,
		ptr_clang_disposeStringSet,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeStringSet", err))
	}
}

// Free the given set of tokens.
/*
Wraps the C function clang_disposeTokens.
*/
func (tU TranslationUnit) DisposeTokens(numTokens uint32) Token {
	c_tU := tU
	var tokens Token
	c_tokens := &tokens
	c_numTokens := numTokens

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_tokens),
		unsafe.Pointer(&c_numTokens),
	}

	err := ffi.CallFunction(
		cif_clang_disposeTokens,
		ptr_clang_disposeTokens,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeTokens", err))
	}
	return tokens
}

// Destroy the specified CXTranslationUnit object.
/*
Wraps the C function clang_disposeTranslationUnit.
*/
func (p0 TranslationUnit) DisposeTranslationUnit() {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_disposeTranslationUnit,
		ptr_clang_disposeTranslationUnit,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeTranslationUnit", err))
	}
}

//
/*
Wraps the C function clang_enableStackTraces.
*/
func EnableStackTraces() {
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_enableStackTraces,
		ptr_clang_enableStackTraces,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_enableStackTraces", err))
	}
}

// Determine whether two cursors are equivalent.
/*
Wraps the C function clang_equalCursors.
*/
func (p0 Cursor) EqualCursors(p1 Cursor) uint32 {
	c_p0 := p0
	c_p1 := p1

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_equalCursors,
		ptr_clang_equalCursors,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalCursors", err))
	}

	ret := retC
	return ret
}

// Determine whether two source locations, which must refer into the same translation unit, refer to exactly the same point in the source code.
/*
Wraps the C function clang_equalLocations.
*/
func (loc1 SourceLocation) EqualLocations(loc2 SourceLocation) uint32 {
	c_loc1 := loc1
	c_loc2 := loc2

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc1),
		unsafe.Pointer(&c_loc2),
	}

	err := ffi.CallFunction(
		cif_clang_equalLocations,
		ptr_clang_equalLocations,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalLocations", err))
	}

	ret := retC
	return ret
}

// Determine whether two ranges are equivalent.
/*
Wraps the C function clang_equalRanges.
*/
func (range1 SourceRange) EqualRanges(range2 SourceRange) uint32 {
	c_range1 := range1
	c_range2 := range2

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range1),
		unsafe.Pointer(&c_range2),
	}

	err := ffi.CallFunction(
		cif_clang_equalRanges,
		ptr_clang_equalRanges,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalRanges", err))
	}

	ret := retC
	return ret
}

// Determine whether two CXTypes represent the same type.
/*
Wraps the C function clang_equalTypes.
*/
func (a Type) EqualTypes(b Type) uint32 {
	c_a := a
	c_b := b

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_a),
		unsafe.Pointer(&c_b),
	}

	err := ffi.CallFunction(
		cif_clang_equalTypes,
		ptr_clang_equalTypes,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalTypes", err))
	}

	ret := retC
	return ret
}

// not supported : clang_executeOnThread : param fn : void (*)(void *)

// Find #import/#include directives in a specific file.
/*
Wraps the C function clang_findIncludesInFile.
*/
func (tU TranslationUnit) FindIncludesInFile(file File, visitor CursorAndRangeVisitor) Result {
	c_tU := tU
	c_file := file
	c_visitor := visitor

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_visitor),
	}

	err := ffi.CallFunction(
		cif_clang_findIncludesInFile,
		ptr_clang_findIncludesInFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findIncludesInFile", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_findIncludesInFileWithBlock.
*/
func (p0 TranslationUnit) FindIncludesInFileWithBlock(p1 File, p2 CursorAndRangeVisitorBlock) Result {
	c_p0 := p0
	c_p1 := p1
	c_p2 := p2

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
		unsafe.Pointer(&c_p2),
	}

	err := ffi.CallFunction(
		cif_clang_findIncludesInFileWithBlock,
		ptr_clang_findIncludesInFileWithBlock,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findIncludesInFileWithBlock", err))
	}

	ret := retC
	return ret
}

// Find references of a declaration in a specific file.
/*
Wraps the C function clang_findReferencesInFile.
*/
func (cursor Cursor) FindReferencesInFile(file File, visitor CursorAndRangeVisitor) Result {
	c_cursor := cursor
	c_file := file
	c_visitor := visitor

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_visitor),
	}

	err := ffi.CallFunction(
		cif_clang_findReferencesInFile,
		ptr_clang_findReferencesInFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findReferencesInFile", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_findReferencesInFileWithBlock.
*/
func (p0 Cursor) FindReferencesInFileWithBlock(p1 File, p2 CursorAndRangeVisitorBlock) Result {
	c_p0 := p0
	c_p1 := p1
	c_p2 := p2

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
		unsafe.Pointer(&c_p2),
	}

	err := ffi.CallFunction(
		cif_clang_findReferencesInFileWithBlock,
		ptr_clang_findReferencesInFileWithBlock,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findReferencesInFileWithBlock", err))
	}

	ret := retC
	return ret
}

/*
Format the given diagnostic in a manner that is suitable for display.

This routine will format the given diagnostic to a string, rendering the diagnostic according to the various options given. The clang_defaultDiagnosticDisplayOptions() function returns the set of options that most closely mimics the behavior of the clang compiler.
*/
/*
Wraps the C function clang_formatDiagnostic.
*/
func (diagnostic Diagnostic) FormatDiagnostic(options uint32) string {
	c_diagnostic := diagnostic
	c_options := options

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic.ptr),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_formatDiagnostic,
		ptr_clang_formatDiagnostic,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_formatDiagnostic", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// free memory allocated by libclang, such as the buffer returned by CXVirtualFileOverlay() or clang_ModuleMapDescriptor_writeToBuffer().
/*
Wraps the C function clang_free.
*/
func Free(buffer unsafe.Pointer) {
	c_buffer := buffer

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_buffer),
	}

	err := ffi.CallFunction(
		cif_clang_free,
		ptr_clang_free,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_free", err))
	}
}

// Returns the address space of the given type.
/*
Wraps the C function clang_getAddressSpace.
*/
func (t Type) AddressSpace() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getAddressSpace,
		ptr_clang_getAddressSpace,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getAddressSpace", err))
	}

	ret := retC
	return ret
}

/*
Retrieve all ranges from all files that were skipped by the preprocessor.

The preprocessor will skip lines when they are surrounded by an if/ifdef/ifndef directive whose condition does not evaluate to true.
*/
/*
Wraps the C function clang_getAllSkippedRanges.
*/
func (tu TranslationUnit) AllSkippedRanges() *SourceRangeList {
	c_tu := tu

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getAllSkippedRanges,
		ptr_clang_getAllSkippedRanges,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getAllSkippedRanges", err))
	}

	ret := (*SourceRangeList)(retC)
	return ret
}

/*
Retrieve the type of a parameter of a function type.

If a non-function type is passed in or the function does not have enough parameters, an invalid type is returned.
*/
/*
Wraps the C function clang_getArgType.
*/
func (t Type) ArgType(i uint32) Type {
	c_t := t
	c_i := i

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(
		cif_clang_getArgType,
		ptr_clang_getArgType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getArgType", err))
	}

	ret := retC
	return ret
}

/*
Return the element type of an array type.

If a non-array type is passed in, an invalid type is returned.
*/
/*
Wraps the C function clang_getArrayElementType.
*/
func (t Type) ArrayElementType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getArrayElementType,
		ptr_clang_getArrayElementType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getArrayElementType", err))
	}

	ret := retC
	return ret
}

/*
Return the array size of a constant array.

If a non-array type is passed in, -1 is returned.
*/
/*
Wraps the C function clang_getArraySize.
*/
func (t Type) ArraySize() int64 {
	c_t := t

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getArraySize,
		ptr_clang_getArraySize,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getArraySize", err))
	}

	ret := retC
	return ret
}

// Retrieve the spelling of a given CXBinaryOperatorKind.
/*
Wraps the C function clang_getBinaryOperatorKindSpelling.
*/
func (kind BinaryOperatorKind) BinaryOperatorKindSpelling() string {
	c_kind := kind

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_kind),
	}

	err := ffi.CallFunction(
		cif_clang_getBinaryOperatorKindSpelling,
		ptr_clang_getBinaryOperatorKindSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getBinaryOperatorKindSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Return the timestamp for use with Clang's -fbuild-session-timestamp= option.
/*
Wraps the C function clang_getBuildSessionTimestamp.
*/
func GetBuildSessionTimestamp() uint64 {
	var retC uint64
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_getBuildSessionTimestamp,
		ptr_clang_getBuildSessionTimestamp,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getBuildSessionTimestamp", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the character data associated with the given string.

The returned data is a reference and not owned by the user. This data is only valid while the `CXString` is valid. This function is similar to `std::string::c_str()`.
*/
/*
Wraps the C function clang_getCString.
*/
func (string_ string_) CString() string {
	c_string_ := string_

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_string_),
	}

	err := ffi.CallFunction(
		cif_clang_getCString,
		ptr_clang_getCString,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCString", err))
	}

	ret := libc.GoString(retC)
	return ret
}

// Return the memory usage of a translation unit.  This object  should be released with clang_disposeCXTUResourceUsage().
/*
Wraps the C function clang_getCXTUResourceUsage.
*/
func (tU TranslationUnit) TUResourceUsage() TUResourceUsage {
	c_tU := tU

	var retC TUResourceUsage
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getCXTUResourceUsage,
		ptr_clang_getCXTUResourceUsage,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCXTUResourceUsage", err))
	}

	ret := retC
	return ret
}

/*
Returns the access control level for the referenced object.

If the cursor refers to a C++ declaration, its access control level within its parent scope is returned. Otherwise, if the cursor refers to a base specifier or access specifier, the specifier itself is returned.
*/
/*
Wraps the C function clang_getCXXAccessSpecifier.
*/
func (p0 Cursor) XAccessSpecifier() CXXAccessSpecifier {
	c_p0 := p0

	var retC CXXAccessSpecifier
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCXXAccessSpecifier,
		ptr_clang_getCXXAccessSpecifier,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCXXAccessSpecifier", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the canonical cursor corresponding to the given cursor.

In the C family of languages, many kinds of entities can be declared several times within a single translation unit. For example, a structure type can be forward-declared (possibly multiple times) and later defined:



The declarations and the definition of X are represented by three different cursors, all of which are declarations of the same underlying entity. One of these cursor is considered the "canonical" cursor, which is effectively the representative for the underlying entity. One can determine if two cursors are declarations of the same underlying entity by comparing their canonical cursors.
*/
/*
Wraps the C function clang_getCanonicalCursor.
*/
func (p0 Cursor) CanonicalCursor() Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCanonicalCursor,
		ptr_clang_getCanonicalCursor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCanonicalCursor", err))
	}

	ret := retC
	return ret
}

/*
Return the canonical type for a CXType.

Clang's type system explicitly models typedefs and all the ways a specific type can be represented.  The canonical type is the underlying type with all the "sugar" removed.  For example, if 'T' is a typedef for 'int', the canonical type for 'T' would be 'int'.
*/
/*
Wraps the C function clang_getCanonicalType.
*/
func (t Type) CanonicalType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getCanonicalType,
		ptr_clang_getCanonicalType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCanonicalType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the child diagnostics of a CXDiagnostic.

This CXDiagnosticSet does not need to be released by clang_disposeDiagnosticSet.
*/
/*
Wraps the C function clang_getChildDiagnostics.
*/
func (d Diagnostic) ChildDiagnostics() DiagnosticSet {
	c_d := d

	var retC DiagnosticSet
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_d.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getChildDiagnostics,
		ptr_clang_getChildDiagnostics,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getChildDiagnostics", err))
	}

	ret := retC
	return ret
}

// Return a version string, suitable for showing to a user, but not        intended to be parsed (the format is not guaranteed to be stable).
/*
Wraps the C function clang_getClangVersion.
*/
func GetClangVersion() string {
	var retC string_
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_getClangVersion,
		ptr_clang_getClangVersion,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getClangVersion", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Retrieve the annotation associated with the given completion string.
/*
Wraps the C function clang_getCompletionAnnotation.
*/
func (completion_string CompletionString) CompletionAnnotation(annotation_number uint32) string {
	c_completion_string := completion_string
	c_annotation_number := annotation_number

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
		unsafe.Pointer(&c_annotation_number),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionAnnotation,
		ptr_clang_getCompletionAnnotation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionAnnotation", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Determine the availability of the entity that this code-completion string refers to.
/*
Wraps the C function clang_getCompletionAvailability.
*/
func (completion_string CompletionString) CompletionAvailability() AvailabilityKind {
	c_completion_string := completion_string

	var retC AvailabilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionAvailability,
		ptr_clang_getCompletionAvailability,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionAvailability", err))
	}

	ret := retC
	return ret
}

// Retrieve the brief documentation comment attached to the declaration that corresponds to the given completion string.
/*
Wraps the C function clang_getCompletionBriefComment.
*/
func (completion_string CompletionString) CompletionBriefComment() string {
	c_completion_string := completion_string

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionBriefComment,
		ptr_clang_getCompletionBriefComment,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionBriefComment", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Retrieve the completion string associated with a particular chunk within a completion string.
/*
Wraps the C function clang_getCompletionChunkCompletionString.
*/
func (completion_string CompletionString) CompletionChunkCompletionString(chunk_number uint32) CompletionString {
	c_completion_string := completion_string
	c_chunk_number := chunk_number

	var retC CompletionString
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
		unsafe.Pointer(&c_chunk_number),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionChunkCompletionString,
		ptr_clang_getCompletionChunkCompletionString,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionChunkCompletionString", err))
	}

	ret := retC
	return ret
}

// Determine the kind of a particular chunk within a completion string.
/*
Wraps the C function clang_getCompletionChunkKind.
*/
func (completion_string CompletionString) CompletionChunkKind(chunk_number uint32) CompletionChunkKind {
	c_completion_string := completion_string
	c_chunk_number := chunk_number

	var retC CompletionChunkKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
		unsafe.Pointer(&c_chunk_number),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionChunkKind,
		ptr_clang_getCompletionChunkKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionChunkKind", err))
	}

	ret := retC
	return ret
}

// Retrieve the text associated with a particular chunk within a completion string.
/*
Wraps the C function clang_getCompletionChunkText.
*/
func (completion_string CompletionString) CompletionChunkText(chunk_number uint32) string {
	c_completion_string := completion_string
	c_chunk_number := chunk_number

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
		unsafe.Pointer(&c_chunk_number),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionChunkText,
		ptr_clang_getCompletionChunkText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionChunkText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Fix-its that *must* be applied before inserting the text for the corresponding completion.

By default, clang_codeCompleteAt() only returns completions with empty fix-its. Extra completions with non-empty fix-its should be explicitly requested by setting CXCodeComplete_IncludeCompletionsWithFixIts.

For the clients to be able to compute position of the cursor after applying fix-its, the following conditions are guaranteed to hold for replacement_range of the stored fix-its:  - Ranges in the fix-its are guaranteed to never contain the completion  point (or identifier under completion point, if any) inside them, except  at the start or at the end of the range.  - If a fix-it range starts or ends with completion point (or starts or  ends after the identifier under completion point), it will contain at  least one character. It allows to unambiguously recompute completion  point after applying the fix-it.

The intuition is that provided fix-its change code around the identifier we complete, but are not allowed to touch the identifier itself or the completion point. One example of completions with corrections are the ones replacing '.' with '->' and vice versa:

std::unique_ptr<std::vector<int>> vec_ptr; In 'vec_ptr.^', one of the completions is 'push_back', it requires replacing '.' with '->'. In 'vec_ptr->^', one of the completions is 'release', it requires replacing '->' with '.'.
*/
/*
Wraps the C function clang_getCompletionFixIt.
*/
func (results *CodeCompleteResults) CompletionFixIt(completion_index uint32, fixit_index uint32) (SourceRange, string) {
	c_results := results
	c_completion_index := completion_index
	c_fixit_index := fixit_index
	var replacement_range SourceRange
	c_replacement_range := &replacement_range

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
		unsafe.Pointer(&c_completion_index),
		unsafe.Pointer(&c_fixit_index),
		unsafe.Pointer(&c_replacement_range),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionFixIt,
		ptr_clang_getCompletionFixIt,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionFixIt", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return replacement_range, ret
}

// Retrieve the number of annotations associated with the given completion string.
/*
Wraps the C function clang_getCompletionNumAnnotations.
*/
func (completion_string CompletionString) CompletionNumAnnotations() uint32 {
	c_completion_string := completion_string

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionNumAnnotations,
		ptr_clang_getCompletionNumAnnotations,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionNumAnnotations", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of fix-its for the given completion index.

Calling this makes sense only if CXCodeComplete_IncludeCompletionsWithFixIts option was set.
*/
/*
Wraps the C function clang_getCompletionNumFixIts.
*/
func (results *CodeCompleteResults) CompletionNumFixIts(completion_index uint32) uint32 {
	c_results := results
	c_completion_index := completion_index

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
		unsafe.Pointer(&c_completion_index),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionNumFixIts,
		ptr_clang_getCompletionNumFixIts,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionNumFixIts", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the parent context of the given completion string.

The parent context of a completion string is the semantic parent of the declaration (if any) that the code completion represents. For example, a code completion for an Objective-C method would have the method's class or protocol as its context.
*/
/*
Wraps the C function clang_getCompletionParent.
*/
func (completion_string CompletionString) CompletionParent() (CursorKind, string) {
	c_completion_string := completion_string
	var kind CursorKind
	c_kind := &kind

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
		unsafe.Pointer(&c_kind),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionParent,
		ptr_clang_getCompletionParent,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionParent", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return kind, ret
}

/*
Determine the priority of this code completion.

The priority of a code completion indicates how likely it is that this particular completion is the completion that the user will select. The priority is selected by various internal heuristics.
*/
/*
Wraps the C function clang_getCompletionPriority.
*/
func (completion_string CompletionString) CompletionPriority() uint32 {
	c_completion_string := completion_string

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getCompletionPriority,
		ptr_clang_getCompletionPriority,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionPriority", err))
	}

	ret := retC
	return ret
}

/*
Map a source location to the cursor that describes the entity at that location in the source code.

clang_getCursor() maps an arbitrary source location within a translation unit down to the most specific cursor that describes the entity at that location. For example, given an expression x + y, invoking clang_getCursor() with a source location pointing to "x" will return the cursor for "x"; similarly for "y". If the cursor points anywhere between "x" or "y" (e.g., on the + or the whitespace around it), clang_getCursor() will return a cursor referring to the "+" expression.
*/
/*
Wraps the C function clang_getCursor.
*/
func (p0 TranslationUnit) Cursor(p1 SourceLocation) Cursor {
	c_p0 := p0
	c_p1 := p1

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_getCursor,
		ptr_clang_getCursor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursor", err))
	}

	ret := retC
	return ret
}

// Determine the availability of the entity that this cursor refers to, taking the current target platform into account.
/*
Wraps the C function clang_getCursorAvailability.
*/
func (cursor Cursor) CursorAvailability() AvailabilityKind {
	c_cursor := cursor

	var retC AvailabilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorAvailability,
		ptr_clang_getCursorAvailability,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorAvailability", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the binary operator kind of this cursor.

If this cursor is not a binary operator then returns Invalid.
*/
/*
Wraps the C function clang_getCursorBinaryOperatorKind.
*/
func (cursor Cursor) CursorBinaryOperatorKind() BinaryOperatorKind {
	c_cursor := cursor

	var retC BinaryOperatorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorBinaryOperatorKind,
		ptr_clang_getCursorBinaryOperatorKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorBinaryOperatorKind", err))
	}

	ret := retC
	return ret
}

// Retrieve a completion string for an arbitrary declaration or macro definition cursor.
/*
Wraps the C function clang_getCursorCompletionString.
*/
func (cursor Cursor) CursorCompletionString() CompletionString {
	c_cursor := cursor

	var retC CompletionString
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorCompletionString,
		ptr_clang_getCursorCompletionString,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorCompletionString", err))
	}

	ret := retC
	return ret
}

/*
For a cursor that is either a reference to or a declaration  of some entity, retrieve a cursor that describes the definition of  that entity.

Some entities can be declared multiple times within a translation  unit, but only one of those declarations can also be a  definition. For example, given:



there are three declarations of the function "f", but only the  second one is a definition. The clang_getCursorDefinition()  function will take any cursor pointing to a declaration of "f"  (the first or fourth lines of the example) or a cursor referenced  that uses "f" (the call to "f' inside "g") and will return a  declaration cursor pointing to the definition (the second "f"  declaration).

If given a cursor for which there is no corresponding definition,  e.g., because there is no definition of that entity within this  translation unit, returns a NULL cursor.
*/
/*
Wraps the C function clang_getCursorDefinition.
*/
func (p0 Cursor) CursorDefinition() Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorDefinition,
		ptr_clang_getCursorDefinition,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorDefinition", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the display name for the entity referenced by this cursor.

The display name contains extra information that helps identify the cursor, such as the parameters of a function or template or the arguments of a class template specialization.
*/
/*
Wraps the C function clang_getCursorDisplayName.
*/
func (p0 Cursor) CursorDisplayName() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorDisplayName,
		ptr_clang_getCursorDisplayName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorDisplayName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the exception specification type associated with a given cursor. This is a value of type CXCursor_ExceptionSpecificationKind.

This only returns a valid result if the cursor refers to a function or method.
*/
/*
Wraps the C function clang_getCursorExceptionSpecificationType.
*/
func (c Cursor) CursorExceptionSpecificationType() int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorExceptionSpecificationType,
		ptr_clang_getCursorExceptionSpecificationType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorExceptionSpecificationType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the physical extent of the source construct referenced by the given cursor.

The extent of a cursor starts with the file/line/column pointing at the first character within the source construct that the cursor refers to and ends with the last character within that source construct. For a declaration, the extent covers the declaration itself. For a reference, the extent covers the location of the reference (e.g., where the referenced entity was actually used).
*/
/*
Wraps the C function clang_getCursorExtent.
*/
func (p0 Cursor) CursorExtent() SourceRange {
	c_p0 := p0

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorExtent,
		ptr_clang_getCursorExtent,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorExtent", err))
	}

	ret := retC
	return ret
}

// Retrieve the kind of the given cursor.
/*
Wraps the C function clang_getCursorKind.
*/
func (p0 Cursor) CursorKind() CursorKind {
	c_p0 := p0

	var retC CursorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorKind,
		ptr_clang_getCursorKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorKind", err))
	}

	ret := retC
	return ret
}

/*
These routines are used for testing and debugging, only, and should not be relied upon.

@{
*/
/*
Wraps the C function clang_getCursorKindSpelling.
*/
func (kind CursorKind) CursorKindSpelling() string {
	c_kind := kind

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_kind),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorKindSpelling,
		ptr_clang_getCursorKindSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorKindSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Determine the "language" of the entity referred to by a given cursor.
/*
Wraps the C function clang_getCursorLanguage.
*/
func (cursor Cursor) CursorLanguage() LanguageKind {
	c_cursor := cursor

	var retC LanguageKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorLanguage,
		ptr_clang_getCursorLanguage,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLanguage", err))
	}

	ret := retC
	return ret
}

/*
Determine the lexical parent of the given cursor.

The lexical parent of a cursor is the cursor in which the given cursor was actually written. For many declarations, the lexical and semantic parents are equivalent (the semantic parent is returned by clang_getCursorSemanticParent()). They diverge when declarations or definitions are provided out-of-line. For example:



In the out-of-line definition of C::f, the semantic parent is the class C, of which this function is a member. The lexical parent is the place where the declaration actually occurs in the source code; in this case, the definition occurs in the translation unit. In general, the lexical parent for a given entity can change without affecting the semantics of the program, and the lexical parent of different declarations of the same entity may be different. Changing the semantic parent of a declaration, on the other hand, can have a major impact on semantics, and redeclarations of a particular entity should all have the same semantic context.

In the example above, both declarations of C::f have C as their semantic context, while the lexical context of the first C::f is C and the lexical context of the second C::f is the translation unit.

For declarations written in the global scope, the lexical parent is the translation unit.
*/
/*
Wraps the C function clang_getCursorLexicalParent.
*/
func (cursor Cursor) CursorLexicalParent() Cursor {
	c_cursor := cursor

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorLexicalParent,
		ptr_clang_getCursorLexicalParent,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLexicalParent", err))
	}

	ret := retC
	return ret
}

// Determine the linkage of the entity referred to by a given cursor.
/*
Wraps the C function clang_getCursorLinkage.
*/
func (cursor Cursor) CursorLinkage() LinkageKind {
	c_cursor := cursor

	var retC LinkageKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorLinkage,
		ptr_clang_getCursorLinkage,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLinkage", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the physical location of the source constructor referenced by the given cursor.

The location of a declaration is typically the location of the name of that declaration, where the name of that declaration would occur if it is unnamed, or some keyword that introduces that particular declaration. The location of a reference is where that reference occurs within the source code.
*/
/*
Wraps the C function clang_getCursorLocation.
*/
func (p0 Cursor) CursorLocation() SourceLocation {
	c_p0 := p0

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorLocation,
		ptr_clang_getCursorLocation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLocation", err))
	}

	ret := retC
	return ret
}

/*
Determine the availability of the entity that this cursor refers to on any platforms for which availability information is known.

















Note that the client is responsible for calling clang_disposeCXPlatformAvailability to free each of the platform-availability structures returned. There are min(N, availability_size) such structures.
*/
/*
Wraps the C function clang_getCursorPlatformAvailability.
*/
func (cursor Cursor) CursorPlatformAvailability(availability_size int32) (int32, string, int32, string, PlatformAvailability, int32) {
	c_cursor := cursor
	var always_deprecated int32
	c_always_deprecated := &always_deprecated
	var deprecated_message string_
	c_deprecated_message := &deprecated_message
	var always_unavailable int32
	c_always_unavailable := &always_unavailable
	var unavailable_message string_
	c_unavailable_message := &unavailable_message
	var availability PlatformAvailability
	c_availability := &availability
	c_availability_size := availability_size

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_always_deprecated),
		unsafe.Pointer(&c_deprecated_message),
		unsafe.Pointer(&c_always_unavailable),
		unsafe.Pointer(&c_unavailable_message),
		unsafe.Pointer(&c_availability),
		unsafe.Pointer(&c_availability_size),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorPlatformAvailability,
		ptr_clang_getCursorPlatformAvailability,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorPlatformAvailability", err))
	}

	ret := retC
	return always_deprecated, deprecated_message.CString(), always_unavailable, unavailable_message.CString(), availability, ret
}

// Pretty print declarations.
/*
Wraps the C function clang_getCursorPrettyPrinted.
*/
func (cursor Cursor) CursorPrettyPrinted(policy PrintingPolicy) string {
	c_cursor := cursor
	c_policy := policy

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_policy),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorPrettyPrinted,
		ptr_clang_getCursorPrettyPrinted,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorPrettyPrinted", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the default policy for the cursor.

The policy should be released after use with clang_PrintingPolicy_dispose.
*/
/*
Wraps the C function clang_getCursorPrintingPolicy.
*/
func (p0 Cursor) CursorPrintingPolicy() PrintingPolicy {
	c_p0 := p0

	var retC PrintingPolicy
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorPrintingPolicy,
		ptr_clang_getCursorPrintingPolicy,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorPrintingPolicy", err))
	}

	ret := retC
	return ret
}

// Given a cursor that references something else, return the source range covering that reference.
/*
Wraps the C function clang_getCursorReferenceNameRange.
*/
func (c Cursor) CursorReferenceNameRange(nameFlags uint32, pieceIndex uint32) SourceRange {
	c_c := c
	c_nameFlags := nameFlags
	c_pieceIndex := pieceIndex

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_nameFlags),
		unsafe.Pointer(&c_pieceIndex),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorReferenceNameRange,
		ptr_clang_getCursorReferenceNameRange,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorReferenceNameRange", err))
	}

	ret := retC
	return ret
}

/*
For a cursor that is a reference, retrieve a cursor representing the entity that it references.

Reference cursors refer to other entities in the AST. For example, an Objective-C superclass reference cursor refers to an Objective-C class. This function produces the cursor for the Objective-C class from the cursor for the superclass reference. If the input cursor is a declaration or definition, it returns that declaration or definition unchanged. Otherwise, returns the NULL cursor.
*/
/*
Wraps the C function clang_getCursorReferenced.
*/
func (p0 Cursor) CursorReferenced() Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorReferenced,
		ptr_clang_getCursorReferenced,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorReferenced", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the return type associated with a given cursor.

This only returns a valid type if the cursor refers to a function or method.
*/
/*
Wraps the C function clang_getCursorResultType.
*/
func (c Cursor) CursorResultType() Type {
	c_c := c

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorResultType,
		ptr_clang_getCursorResultType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorResultType", err))
	}

	ret := retC
	return ret
}

/*
Determine the semantic parent of the given cursor.

The semantic parent of a cursor is the cursor that semantically contains the given cursor. For many declarations, the lexical and semantic parents are equivalent (the lexical parent is returned by clang_getCursorLexicalParent()). They diverge when declarations or definitions are provided out-of-line. For example:



In the out-of-line definition of C::f, the semantic parent is the class C, of which this function is a member. The lexical parent is the place where the declaration actually occurs in the source code; in this case, the definition occurs in the translation unit. In general, the lexical parent for a given entity can change without affecting the semantics of the program, and the lexical parent of different declarations of the same entity may be different. Changing the semantic parent of a declaration, on the other hand, can have a major impact on semantics, and redeclarations of a particular entity should all have the same semantic context.

In the example above, both declarations of C::f have C as their semantic context, while the lexical context of the first C::f is C and the lexical context of the second C::f is the translation unit.

For global declarations, the semantic parent is the translation unit.
*/
/*
Wraps the C function clang_getCursorSemanticParent.
*/
func (cursor Cursor) CursorSemanticParent() Cursor {
	c_cursor := cursor

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorSemanticParent,
		ptr_clang_getCursorSemanticParent,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorSemanticParent", err))
	}

	ret := retC
	return ret
}

// Retrieve a name for the entity referenced by this cursor.
/*
Wraps the C function clang_getCursorSpelling.
*/
func (p0 Cursor) CursorSpelling() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorSpelling,
		ptr_clang_getCursorSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Determine the "thread-local storage (TLS) kind" of the declaration referred to by a cursor.
/*
Wraps the C function clang_getCursorTLSKind.
*/
func (cursor Cursor) CursorTLSKind() TLSKind {
	c_cursor := cursor

	var retC TLSKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorTLSKind,
		ptr_clang_getCursorTLSKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorTLSKind", err))
	}

	ret := retC
	return ret
}

// Retrieve the type of a CXCursor (if any).
/*
Wraps the C function clang_getCursorType.
*/
func (c Cursor) CursorType() Type {
	c_c := c

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorType,
		ptr_clang_getCursorType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve a Unified Symbol Resolution (USR) for the entity referenced by the given cursor.

A Unified Symbol Resolution (USR) is a string that identifies a particular entity (function, class, variable, etc.) within a program. USRs can be compared across translation units to determine, e.g., when references in one translation refer to an entity defined in another translation unit.
*/
/*
Wraps the C function clang_getCursorUSR.
*/
func (p0 Cursor) CursorUSR() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorUSR,
		ptr_clang_getCursorUSR,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorUSR", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the unary operator kind of this cursor.

If this cursor is not a unary operator then returns Invalid.
*/
/*
Wraps the C function clang_getCursorUnaryOperatorKind.
*/
func (cursor Cursor) CursorUnaryOperatorKind() UnaryOperatorKind {
	c_cursor := cursor

	var retC UnaryOperatorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorUnaryOperatorKind,
		ptr_clang_getCursorUnaryOperatorKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorUnaryOperatorKind", err))
	}

	ret := retC
	return ret
}

/*
Describe the visibility of the entity referred to by a cursor.

This returns the default visibility if not explicitly specified by a visibility attribute. The default visibility may be changed by commandline arguments.
*/
/*
Wraps the C function clang_getCursorVisibility.
*/
func (cursor Cursor) CursorVisibility() VisibilityKind {
	c_cursor := cursor

	var retC VisibilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getCursorVisibility,
		ptr_clang_getCursorVisibility,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorVisibility", err))
	}

	ret := retC
	return ret
}

// Returns the Objective-C type encoding for the specified declaration.
/*
Wraps the C function clang_getDeclObjCTypeEncoding.
*/
func (c Cursor) DeclObjCTypeEncoding() string {
	c_c := c

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getDeclObjCTypeEncoding,
		ptr_clang_getDeclObjCTypeEncoding,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDeclObjCTypeEncoding", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// not supported : clang_getDefinitionSpellingAndExtent : param startBuf : const char **

// Retrieve a diagnostic associated with the given translation unit.
/*
Wraps the C function clang_getDiagnostic.
*/
func (unit TranslationUnit) Diagnostic(index uint32) Diagnostic {
	c_unit := unit
	c_index := index

	var retC Diagnostic
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_unit.ptr),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnostic,
		ptr_clang_getDiagnostic,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnostic", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the category number for this diagnostic.

Diagnostics can be categorized into groups along with other, related diagnostics (e.g., diagnostics under the same warning flag). This routine retrieves the category number for the given diagnostic.
*/
/*
Wraps the C function clang_getDiagnosticCategory.
*/
func (p0 Diagnostic) DiagnosticCategory() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticCategory,
		ptr_clang_getDiagnosticCategory,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategory", err))
	}

	ret := retC
	return ret
}

// Retrieve the name of a particular diagnostic category.  This  is now deprecated.  Use clang_getDiagnosticCategoryText()  instead.
/*
Wraps the C function clang_getDiagnosticCategoryName.
*/
func GetDiagnosticCategoryName(category uint32) string {
	c_category := category

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_category),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticCategoryName,
		ptr_clang_getDiagnosticCategoryName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategoryName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Retrieve the diagnostic category text for a given diagnostic.
/*
Wraps the C function clang_getDiagnosticCategoryText.
*/
func (p0 Diagnostic) DiagnosticCategoryText() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticCategoryText,
		ptr_clang_getDiagnosticCategoryText,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategoryText", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the replacement information for a given fix-it.

Fix-its are described in terms of a source range whose contents should be replaced by a string. This approach generalizes over three kinds of operations: removal of source code (the range covers the code to be removed and the replacement string is empty), replacement of source code (the range covers the code to be replaced and the replacement string provides the new code), and insertion (both the start and end of the range point at the insertion location, and the replacement string provides the text to insert).
*/
/*
Wraps the C function clang_getDiagnosticFixIt.
*/
func (diagnostic Diagnostic) DiagnosticFixIt(fixIt uint32) (SourceRange, string) {
	c_diagnostic := diagnostic
	c_fixIt := fixIt
	var replacementRange SourceRange
	c_replacementRange := &replacementRange

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic.ptr),
		unsafe.Pointer(&c_fixIt),
		unsafe.Pointer(&c_replacementRange),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticFixIt,
		ptr_clang_getDiagnosticFixIt,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticFixIt", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return replacementRange, ret
}

// Retrieve a diagnostic associated with the given CXDiagnosticSet.
/*
Wraps the C function clang_getDiagnosticInSet.
*/
func (diags DiagnosticSet) DiagnosticInSet(index uint32) Diagnostic {
	c_diags := diags
	c_index := index

	var retC Diagnostic
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diags.ptr),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticInSet,
		ptr_clang_getDiagnosticInSet,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticInSet", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the source location of the given diagnostic.

This location is where Clang would print the caret ('^') when displaying the diagnostic on the command line.
*/
/*
Wraps the C function clang_getDiagnosticLocation.
*/
func (p0 Diagnostic) DiagnosticLocation() SourceLocation {
	c_p0 := p0

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticLocation,
		ptr_clang_getDiagnosticLocation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticLocation", err))
	}

	ret := retC
	return ret
}

// Determine the number of fix-it hints associated with the given diagnostic.
/*
Wraps the C function clang_getDiagnosticNumFixIts.
*/
func (diagnostic Diagnostic) DiagnosticNumFixIts() uint32 {
	c_diagnostic := diagnostic

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticNumFixIts,
		ptr_clang_getDiagnosticNumFixIts,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticNumFixIts", err))
	}

	ret := retC
	return ret
}

// Determine the number of source ranges associated with the given diagnostic.
/*
Wraps the C function clang_getDiagnosticNumRanges.
*/
func (p0 Diagnostic) DiagnosticNumRanges() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticNumRanges,
		ptr_clang_getDiagnosticNumRanges,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticNumRanges", err))
	}

	ret := retC
	return ret
}

// Retrieve the name of the command-line option that enabled this diagnostic.
/*
Wraps the C function clang_getDiagnosticOption.
*/
func (diag Diagnostic) DiagnosticOption() (string, string) {
	c_diag := diag
	var disable string_
	c_disable := &disable

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diag.ptr),
		unsafe.Pointer(&c_disable),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticOption,
		ptr_clang_getDiagnosticOption,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticOption", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return disable.CString(), ret
}

/*
Retrieve a source range associated with the diagnostic.

A diagnostic's source ranges highlight important elements in the source code. On the command line, Clang displays source ranges by underlining them with '~' characters.
*/
/*
Wraps the C function clang_getDiagnosticRange.
*/
func (diagnostic Diagnostic) DiagnosticRange(range_ uint32) SourceRange {
	c_diagnostic := diagnostic
	c_range_ := range_

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic.ptr),
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticRange,
		ptr_clang_getDiagnosticRange,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticRange", err))
	}

	ret := retC
	return ret
}

// Retrieve the complete set of diagnostics associated with a        translation unit.
/*
Wraps the C function clang_getDiagnosticSetFromTU.
*/
func (unit TranslationUnit) DiagnosticSetFromTU() DiagnosticSet {
	c_unit := unit

	var retC DiagnosticSet
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_unit.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticSetFromTU,
		ptr_clang_getDiagnosticSetFromTU,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticSetFromTU", err))
	}

	ret := retC
	return ret
}

// Determine the severity of the given diagnostic.
/*
Wraps the C function clang_getDiagnosticSeverity.
*/
func (p0 Diagnostic) DiagnosticSeverity() DiagnosticSeverity {
	c_p0 := p0

	var retC DiagnosticSeverity
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticSeverity,
		ptr_clang_getDiagnosticSeverity,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticSeverity", err))
	}

	ret := retC
	return ret
}

// Retrieve the text of the given diagnostic.
/*
Wraps the C function clang_getDiagnosticSpelling.
*/
func (p0 Diagnostic) DiagnosticSpelling() string {
	c_p0 := p0

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getDiagnosticSpelling,
		ptr_clang_getDiagnosticSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Return the element type of an array, complex, or vector type.

If a type is passed in that is not an array, complex, or vector type, an invalid type is returned.
*/
/*
Wraps the C function clang_getElementType.
*/
func (t Type) ElementType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getElementType,
		ptr_clang_getElementType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getElementType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the integer value of an enum constant declaration as an unsigned  long long.

If the cursor does not reference an enum constant declaration, ULLONG_MAX is returned. Since this is also potentially a valid constant value, the kind of the cursor must be verified before calling this function.
*/
/*
Wraps the C function clang_getEnumConstantDeclUnsignedValue.
*/
func (c Cursor) EnumConstantDeclUnsignedValue() uint64 {
	c_c := c

	var retC uint64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getEnumConstantDeclUnsignedValue,
		ptr_clang_getEnumConstantDeclUnsignedValue,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getEnumConstantDeclUnsignedValue", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the integer value of an enum constant declaration as a signed  long long.

If the cursor does not reference an enum constant declaration, LLONG_MIN is returned. Since this is also potentially a valid constant value, the kind of the cursor must be verified before calling this function.
*/
/*
Wraps the C function clang_getEnumConstantDeclValue.
*/
func (c Cursor) EnumConstantDeclValue() int64 {
	c_c := c

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getEnumConstantDeclValue,
		ptr_clang_getEnumConstantDeclValue,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getEnumConstantDeclValue", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the integer type of an enum declaration.

If the cursor does not reference an enum declaration, an invalid type is returned.
*/
/*
Wraps the C function clang_getEnumDeclIntegerType.
*/
func (c Cursor) EnumDeclIntegerType() Type {
	c_c := c

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getEnumDeclIntegerType,
		ptr_clang_getEnumDeclIntegerType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getEnumDeclIntegerType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the exception specification type associated with a function type. This is a value of type CXCursor_ExceptionSpecificationKind.

If a non-function type is passed in, an error code of -1 is returned.
*/
/*
Wraps the C function clang_getExceptionSpecificationType.
*/
func (t Type) ExceptionSpecificationType() int32 {
	c_t := t

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getExceptionSpecificationType,
		ptr_clang_getExceptionSpecificationType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getExceptionSpecificationType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the file, line, column, and offset represented by the given source location.

If the location refers into a macro expansion, retrieves the location of the macro expansion.
*/
/*
Wraps the C function clang_getExpansionLocation.
*/
func (location SourceLocation) ExpansionLocation() (File, uint32, uint32, uint32) {
	c_location := location
	var file File
	c_file := &file
	var line uint32
	c_line := &line
	var column uint32
	c_column := &column
	var offset uint32
	c_offset := &offset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(
		cif_clang_getExpansionLocation,
		ptr_clang_getExpansionLocation,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getExpansionLocation", err))
	}
	return file, line, column, offset
}

/*
Retrieve the bit width of a bit-field declaration as an integer.

If the cursor does not reference a bit-field, or if the bit-field's width expression cannot be evaluated, -1 is returned.

For example:
*/
/*
Wraps the C function clang_getFieldDeclBitWidth.
*/
func (c Cursor) FieldDeclBitWidth() int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getFieldDeclBitWidth,
		ptr_clang_getFieldDeclBitWidth,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFieldDeclBitWidth", err))
	}

	ret := retC
	return ret
}

// Retrieve a file handle within the given translation unit.
/*
Wraps the C function clang_getFile.
*/
func (tu TranslationUnit) File(file_name string) File {
	c_tu := tu
	c_file_name, free_c_file_name := libc.CString(file_name)
	defer free_c_file_name()

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_file_name),
	}

	err := ffi.CallFunction(
		cif_clang_getFile,
		ptr_clang_getFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFile", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getFileContents : param size : size_t *

/*
Retrieve the file, line, column, and offset represented by the given source location.

If the location refers into a macro expansion, return where the macro was expanded or where the macro argument was written, if the location points at a macro argument.
*/
/*
Wraps the C function clang_getFileLocation.
*/
func (location SourceLocation) FileLocation() (File, uint32, uint32, uint32) {
	c_location := location
	var file File
	c_file := &file
	var line uint32
	c_line := &line
	var column uint32
	c_column := &column
	var offset uint32
	c_offset := &offset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(
		cif_clang_getFileLocation,
		ptr_clang_getFileLocation,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFileLocation", err))
	}
	return file, line, column, offset
}

// Retrieve the complete file and path name of the given file.
/*
Wraps the C function clang_getFileName.
*/
func (sFile File) FileName() string {
	c_sFile := sFile

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_sFile.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getFileName,
		ptr_clang_getFileName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFileName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// not supported : clang_getFileTime : return value : time_t

// Retrieve the unique ID for the given file.
/*
Wraps the C function clang_getFileUniqueID.
*/
func (file File) FileUniqueID() (FileUniqueID, int32) {
	c_file := file
	var outID FileUniqueID
	c_outID := &outID

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_file.ptr),
		unsafe.Pointer(&c_outID),
	}

	err := ffi.CallFunction(
		cif_clang_getFileUniqueID,
		ptr_clang_getFileUniqueID,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFileUniqueID", err))
	}

	ret := retC
	return outID, ret
}

/*
Get the fully qualified name for a type.

This includes full qualification of all template parameters.

Policy - Further refine the type formatting WithGlobalNsPrefix - If non-zero, function will prepend a '::' to qualified names
*/
/*
Wraps the C function clang_getFullyQualifiedName.
*/
func (cT Type) FullyQualifiedName(policy PrintingPolicy, withGlobalNsPrefix uint32) string {
	c_cT := cT
	c_policy := policy
	c_withGlobalNsPrefix := withGlobalNsPrefix

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
		unsafe.Pointer(&c_policy),
		unsafe.Pointer(&c_withGlobalNsPrefix),
	}

	err := ffi.CallFunction(
		cif_clang_getFullyQualifiedName,
		ptr_clang_getFullyQualifiedName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFullyQualifiedName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the calling convention associated with a function type.

If a non-function type is passed in, CXCallingConv_Invalid is returned.
*/
/*
Wraps the C function clang_getFunctionTypeCallingConv.
*/
func (t Type) FunctionTypeCallingConv() CallingConv {
	c_t := t

	var retC CallingConv
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getFunctionTypeCallingConv,
		ptr_clang_getFunctionTypeCallingConv,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFunctionTypeCallingConv", err))
	}

	ret := retC
	return ret
}

// For cursors representing an iboutletcollection attribute,  this function returns the collection element type.
/*
Wraps the C function clang_getIBOutletCollectionType.
*/
func (p0 Cursor) IBOutletCollectionType() Type {
	c_p0 := p0

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getIBOutletCollectionType,
		ptr_clang_getIBOutletCollectionType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getIBOutletCollectionType", err))
	}

	ret := retC
	return ret
}

// Retrieve the file that is included by the given inclusion directive cursor.
/*
Wraps the C function clang_getIncludedFile.
*/
func (cursor Cursor) IncludedFile() File {
	c_cursor := cursor

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getIncludedFile,
		ptr_clang_getIncludedFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getIncludedFile", err))
	}

	ret := retC
	return ret
}

// Visit the set of preprocessor inclusions in a translation unit.   The visitor function is called with the provided data for every included   file.  This does not include headers included by the PCH file (unless one   is inspecting the inclusions in the PCH file itself).
/*
Wraps the C function clang_getInclusions.
*/
func (tu TranslationUnit) Inclusions(visitor InclusionVisitor, client_data ClientData) {
	c_tu := tu
	c_visitor := visitor
	c_client_data := client_data

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(
		cif_clang_getInclusions,
		ptr_clang_getInclusions,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getInclusions", err))
	}
}

/*
Legacy API to retrieve the file, line, column, and offset represented by the given source location.

This interface has been replaced by the newer interface #clang_getExpansionLocation(). See that interface's documentation for details.
*/
/*
Wraps the C function clang_getInstantiationLocation.
*/
func (location SourceLocation) InstantiationLocation() (File, uint32, uint32, uint32) {
	c_location := location
	var file File
	c_file := &file
	var line uint32
	c_line := &line
	var column uint32
	c_column := &column
	var offset uint32
	c_offset := &offset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(
		cif_clang_getInstantiationLocation,
		ptr_clang_getInstantiationLocation,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getInstantiationLocation", err))
	}
	return file, line, column, offset
}

// Retrieves the source location associated with a given file/line/column in a particular translation unit.
/*
Wraps the C function clang_getLocation.
*/
func (tu TranslationUnit) Location(file File, line uint32, column uint32) SourceLocation {
	c_tu := tu
	c_file := file
	c_line := line
	c_column := column

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
	}

	err := ffi.CallFunction(
		cif_clang_getLocation,
		ptr_clang_getLocation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getLocation", err))
	}

	ret := retC
	return ret
}

// Retrieves the source location associated with a given character offset in a particular translation unit.
/*
Wraps the C function clang_getLocationForOffset.
*/
func (tu TranslationUnit) LocationForOffset(file File, offset uint32) SourceLocation {
	c_tu := tu
	c_file := file
	c_offset := offset

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(
		cif_clang_getLocationForOffset,
		ptr_clang_getLocationForOffset,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getLocationForOffset", err))
	}

	ret := retC
	return ret
}

// Given a CXFile header file, return the module that contains it, if one exists.
/*
Wraps the C function clang_getModuleForFile.
*/
func (p0 TranslationUnit) ModuleForFile(p1 File) Module {
	c_p0 := p0
	c_p1 := p1

	var retC Module
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_getModuleForFile,
		ptr_clang_getModuleForFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getModuleForFile", err))
	}

	ret := retC
	return ret
}

/*
For reference types (e.g., "const int&"), returns the type that the reference refers to (e.g "const int").

Otherwise, returns the type itself.

A type that has kind CXType_LValueReference or CXType_RValueReference is a reference type.
*/
/*
Wraps the C function clang_getNonReferenceType.
*/
func (cT Type) NonReferenceType() Type {
	c_cT := cT

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(
		cif_clang_getNonReferenceType,
		ptr_clang_getNonReferenceType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNonReferenceType", err))
	}

	ret := retC
	return ret
}

// Retrieve the NULL cursor, which represents no entity.
/*
Wraps the C function clang_getNullCursor.
*/
func GetNullCursor() Cursor {
	var retC Cursor
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_getNullCursor,
		ptr_clang_getNullCursor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNullCursor", err))
	}

	ret := retC
	return ret
}

// Retrieve a NULL (invalid) source location.
/*
Wraps the C function clang_getNullLocation.
*/
func GetNullLocation() SourceLocation {
	var retC SourceLocation
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_getNullLocation,
		ptr_clang_getNullLocation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNullLocation", err))
	}

	ret := retC
	return ret
}

// Retrieve a NULL (invalid) source range.
/*
Wraps the C function clang_getNullRange.
*/
func GetNullRange() SourceRange {
	var retC SourceRange
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(
		cif_clang_getNullRange,
		ptr_clang_getNullRange,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNullRange", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of non-variadic parameters associated with a function type.

If a non-function type is passed in, -1 is returned.
*/
/*
Wraps the C function clang_getNumArgTypes.
*/
func (t Type) NumArgTypes() int32 {
	c_t := t

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getNumArgTypes,
		ptr_clang_getNumArgTypes,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumArgTypes", err))
	}

	ret := retC
	return ret
}

// Retrieve the number of chunks in the given code-completion string.
/*
Wraps the C function clang_getNumCompletionChunks.
*/
func (completion_string CompletionString) NumCompletionChunks() uint32 {
	c_completion_string := completion_string

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getNumCompletionChunks,
		ptr_clang_getNumCompletionChunks,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumCompletionChunks", err))
	}

	ret := retC
	return ret
}

// Determine the number of diagnostics produced for the given translation unit.
/*
Wraps the C function clang_getNumDiagnostics.
*/
func (unit TranslationUnit) NumDiagnostics() uint32 {
	c_unit := unit

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_unit.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getNumDiagnostics,
		ptr_clang_getNumDiagnostics,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumDiagnostics", err))
	}

	ret := retC
	return ret
}

// Determine the number of diagnostics in a CXDiagnosticSet.
/*
Wraps the C function clang_getNumDiagnosticsInSet.
*/
func (diags DiagnosticSet) NumDiagnosticsInSet() uint32 {
	c_diags := diags

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diags.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getNumDiagnosticsInSet,
		ptr_clang_getNumDiagnosticsInSet,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumDiagnosticsInSet", err))
	}

	ret := retC
	return ret
}

/*
Return the number of elements of an array or vector type.

If a type is passed in that is not an array or vector type, -1 is returned.
*/
/*
Wraps the C function clang_getNumElements.
*/
func (t Type) NumElements() int64 {
	c_t := t

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getNumElements,
		ptr_clang_getNumElements,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumElements", err))
	}

	ret := retC
	return ret
}

// Determine the number of overloaded declarations referenced by a CXCursor_OverloadedDeclRef cursor.
/*
Wraps the C function clang_getNumOverloadedDecls.
*/
func (cursor Cursor) NumOverloadedDecls() uint32 {
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getNumOverloadedDecls,
		ptr_clang_getNumOverloadedDecls,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumOverloadedDecls", err))
	}

	ret := retC
	return ret
}

/*
Returns the offset in bits of a CX_CXXBaseSpecifier relative to the parent class.

Returns a small negative number if the offset cannot be computed. See CXTypeLayoutError for error codes.
*/
/*
Wraps the C function clang_getOffsetOfBase.
*/
func (parent Cursor) OffsetOfBase(base Cursor) int64 {
	c_parent := parent
	c_base := base

	var retC int64
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_parent),
		unsafe.Pointer(&c_base),
	}

	err := ffi.CallFunction(
		cif_clang_getOffsetOfBase,
		ptr_clang_getOffsetOfBase,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getOffsetOfBase", err))
	}

	ret := retC
	return ret
}

// Retrieve a cursor for one of the overloaded declarations referenced by a CXCursor_OverloadedDeclRef cursor.
/*
Wraps the C function clang_getOverloadedDecl.
*/
func (cursor Cursor) OverloadedDecl(index uint32) Cursor {
	c_cursor := cursor
	c_index := index

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(
		cif_clang_getOverloadedDecl,
		ptr_clang_getOverloadedDecl,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getOverloadedDecl", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getOverriddenCursors : param overridden : CXCursor **

// For pointer types, returns the type of the pointee.
/*
Wraps the C function clang_getPointeeType.
*/
func (t Type) PointeeType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getPointeeType,
		ptr_clang_getPointeeType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getPointeeType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the file, line and column represented by the given source location, as specified in a # line directive.

Example: given the following source code in a file somefile.c



the location information returned by this function would be

File: dummy.c Line: 124 Column: 12

whereas clang_getExpansionLocation would have returned

File: somefile.c Line: 3 Column: 12
*/
/*
Wraps the C function clang_getPresumedLocation.
*/
func (location SourceLocation) PresumedLocation() (string, uint32, uint32) {
	c_location := location
	var filename string_
	c_filename := &filename
	var line uint32
	c_line := &line
	var column uint32
	c_column := &column

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
		unsafe.Pointer(&c_filename),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
	}

	err := ffi.CallFunction(
		cif_clang_getPresumedLocation,
		ptr_clang_getPresumedLocation,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getPresumedLocation", err))
	}
	return filename.CString(), line, column
}

// Retrieve a source range given the beginning and ending source locations.
/*
Wraps the C function clang_getRange.
*/
func (begin SourceLocation) Range(end SourceLocation) SourceRange {
	c_begin := begin
	c_end := end

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_begin),
		unsafe.Pointer(&c_end),
	}

	err := ffi.CallFunction(
		cif_clang_getRange,
		ptr_clang_getRange,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRange", err))
	}

	ret := retC
	return ret
}

// Retrieve a source location representing the last character within a source range.
/*
Wraps the C function clang_getRangeEnd.
*/
func (range_ SourceRange) RangeEnd() SourceLocation {
	c_range_ := range_

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(
		cif_clang_getRangeEnd,
		ptr_clang_getRangeEnd,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRangeEnd", err))
	}

	ret := retC
	return ret
}

// Retrieve a source location representing the first character within a source range.
/*
Wraps the C function clang_getRangeStart.
*/
func (range_ SourceRange) RangeStart() SourceLocation {
	c_range_ := range_

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(
		cif_clang_getRangeStart,
		ptr_clang_getRangeStart,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRangeStart", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_getRemappings.
*/
func GetRemappings(p0 string) Remapping {
	c_p0, free_c_p0 := libc.CString(p0)
	defer free_c_p0()

	var retC Remapping
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getRemappings,
		ptr_clang_getRemappings,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRemappings", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getRemappingsFromFileList : param p0 : const char **

/*
Retrieve the return type associated with a function type.

If a non-function type is passed in, an invalid type is returned.
*/
/*
Wraps the C function clang_getResultType.
*/
func (t Type) ResultType() Type {
	c_t := t

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getResultType,
		ptr_clang_getResultType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getResultType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve all ranges that were skipped by the preprocessor.

The preprocessor will skip lines when they are surrounded by an if/ifdef/ifndef directive whose condition does not evaluate to true.
*/
/*
Wraps the C function clang_getSkippedRanges.
*/
func (tu TranslationUnit) SkippedRanges(file File) *SourceRangeList {
	c_tu := tu
	c_file := file

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_file),
	}

	err := ffi.CallFunction(
		cif_clang_getSkippedRanges,
		ptr_clang_getSkippedRanges,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getSkippedRanges", err))
	}

	ret := (*SourceRangeList)(retC)
	return ret
}

/*
Given a cursor that may represent a specialization or instantiation of a template, retrieve the cursor that represents the template that it specializes or from which it was instantiated.

This routine determines the template involved both for explicit specializations of templates and for implicit instantiations of the template, both of which are referred to as "specializations". For a class template specialization (e.g., std::vector<bool>), this routine will return either the primary template (std::vector) or, if the specialization was instantiated from a class template partial specialization, the class template partial specialization. For a class template partial specialization and a function template specialization (including instantiations), this this routine will return the specialized template.

For members of a class template (e.g., member functions, member classes, or static data members), returns the specialized or instantiated member. Although not strictly "templates" in the C++ language, members of class templates have the same notions of specializations and instantiations that templates do, so this routine treats them similarly.
*/
/*
Wraps the C function clang_getSpecializedCursorTemplate.
*/
func (c Cursor) SpecializedCursorTemplate() Cursor {
	c_c := c

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getSpecializedCursorTemplate,
		ptr_clang_getSpecializedCursorTemplate,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getSpecializedCursorTemplate", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the file, line, column, and offset represented by the given source location.

If the location refers into a macro instantiation, return where the location was originally spelled in the source file.
*/
/*
Wraps the C function clang_getSpellingLocation.
*/
func (location SourceLocation) SpellingLocation() (File, uint32, uint32, uint32) {
	c_location := location
	var file File
	c_file := &file
	var line uint32
	c_line := &line
	var column uint32
	c_column := &column
	var offset uint32
	c_offset := &offset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(
		cif_clang_getSpellingLocation,
		ptr_clang_getSpellingLocation,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getSpellingLocation", err))
	}
	return file, line, column, offset
}

/*
Generate a single symbol symbol graph for the declaration at the given cursor. Returns a null string if the AST node for the cursor isn't a declaration.

The output contains the symbol graph as well as some additional information about related symbols.
*/
/*
Wraps the C function clang_getSymbolGraphForCursor.
*/
func (cursor Cursor) SymbolGraphForCursor() string {
	c_cursor := cursor

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(
		cif_clang_getSymbolGraphForCursor,
		ptr_clang_getSymbolGraphForCursor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getSymbolGraphForCursor", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Generate a single symbol symbol graph for the given USR. Returns a null string if the associated symbol can not be found in the provided CXAPISet.

The output contains the symbol graph as well as some additional information about related symbols.
*/
/*
Wraps the C function clang_getSymbolGraphForUSR.
*/
func GetSymbolGraphForUSR(usr string, api APISet) string {
	c_usr, free_c_usr := libc.CString(usr)
	defer free_c_usr()
	c_api := api

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_usr),
		unsafe.Pointer(&c_api),
	}

	err := ffi.CallFunction(
		cif_clang_getSymbolGraphForUSR,
		ptr_clang_getSymbolGraphForUSR,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getSymbolGraphForUSR", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Returns the human-readable null-terminated C string that represents  the name of the memory category.  This string should never be freed.
/*
Wraps the C function clang_getTUResourceUsageName.
*/
func (kind TUResourceUsageKind) TUResourceUsageName() string {
	c_kind := kind

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_kind),
	}

	err := ffi.CallFunction(
		cif_clang_getTUResourceUsageName,
		ptr_clang_getTUResourceUsageName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTUResourceUsageName", err))
	}

	ret := libc.GoString(retC)
	return ret
}

/*
Given a cursor that represents a template, determine the cursor kind of the specializations would be generated by instantiating the template.

This routine can be used to determine what flavor of function template, class template, or class template partial specialization is stored in the cursor. For example, it can describe whether a class template cursor is declared with "struct", "class" or "union".
*/
/*
Wraps the C function clang_getTemplateCursorKind.
*/
func (c Cursor) TemplateCursorKind() CursorKind {
	c_c := c

	var retC CursorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getTemplateCursorKind,
		ptr_clang_getTemplateCursorKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTemplateCursorKind", err))
	}

	ret := retC
	return ret
}

// Get the raw lexical token starting with the given location.
/*
Wraps the C function clang_getToken.
*/
func (tU TranslationUnit) Token(location SourceLocation) *Token {
	c_tU := tU
	c_location := location

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_location),
	}

	err := ffi.CallFunction(
		cif_clang_getToken,
		ptr_clang_getToken,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getToken", err))
	}

	ret := (*Token)(retC)
	return ret
}

// Retrieve a source range that covers the given token.
/*
Wraps the C function clang_getTokenExtent.
*/
func (p0 TranslationUnit) TokenExtent(p1 Token) SourceRange {
	c_p0 := p0
	c_p1 := p1

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_getTokenExtent,
		ptr_clang_getTokenExtent,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenExtent", err))
	}

	ret := retC
	return ret
}

// Determine the kind of the given token.
/*
Wraps the C function clang_getTokenKind.
*/
func (p0 Token) TokenKind() TokenKind {
	c_p0 := p0

	var retC TokenKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_getTokenKind,
		ptr_clang_getTokenKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenKind", err))
	}

	ret := retC
	return ret
}

// Retrieve the source location of the given token.
/*
Wraps the C function clang_getTokenLocation.
*/
func (p0 TranslationUnit) TokenLocation(p1 Token) SourceLocation {
	c_p0 := p0
	c_p1 := p1

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_getTokenLocation,
		ptr_clang_getTokenLocation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenLocation", err))
	}

	ret := retC
	return ret
}

/*
Determine the spelling of the given token.

The spelling of a token is the textual representation of that token, e.g., the text of an identifier or keyword.
*/
/*
Wraps the C function clang_getTokenSpelling.
*/
func (p0 TranslationUnit) TokenSpelling(p1 Token) string {
	c_p0 := p0
	c_p1 := p1

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_getTokenSpelling,
		ptr_clang_getTokenSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the cursor that represents the given translation unit.

The translation unit cursor can be used to start traversing the various declarations within the given translation unit.
*/
/*
Wraps the C function clang_getTranslationUnitCursor.
*/
func (p0 TranslationUnit) TranslationUnitCursor() Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getTranslationUnitCursor,
		ptr_clang_getTranslationUnitCursor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTranslationUnitCursor", err))
	}

	ret := retC
	return ret
}

// Get the original translation unit source file name.
/*
Wraps the C function clang_getTranslationUnitSpelling.
*/
func (cTUnit TranslationUnit) TranslationUnitSpelling() string {
	c_cTUnit := cTUnit

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cTUnit.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getTranslationUnitSpelling,
		ptr_clang_getTranslationUnitSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTranslationUnitSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Get target information for this translation unit.

The CXTargetInfo object cannot outlive the CXTranslationUnit object.
*/
/*
Wraps the C function clang_getTranslationUnitTargetInfo.
*/
func (cTUnit TranslationUnit) TranslationUnitTargetInfo() TargetInfo {
	c_cTUnit := cTUnit

	var retC TargetInfo
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cTUnit.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_getTranslationUnitTargetInfo,
		ptr_clang_getTranslationUnitTargetInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTranslationUnitTargetInfo", err))
	}

	ret := retC
	return ret
}

// Return the cursor for the declaration of the given type.
/*
Wraps the C function clang_getTypeDeclaration.
*/
func (t Type) TypeDeclaration() Cursor {
	c_t := t

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_getTypeDeclaration,
		ptr_clang_getTypeDeclaration,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypeDeclaration", err))
	}

	ret := retC
	return ret
}

// Retrieve the spelling of a given CXTypeKind.
/*
Wraps the C function clang_getTypeKindSpelling.
*/
func (k TypeKind) TypeKindSpelling() string {
	c_k := k

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_k),
	}

	err := ffi.CallFunction(
		cif_clang_getTypeKindSpelling,
		ptr_clang_getTypeKindSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypeKindSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Pretty-print the underlying type using a custom printing policy.

If the type is invalid, an empty string is returned.
*/
/*
Wraps the C function clang_getTypePrettyPrinted.
*/
func (cT Type) TypePrettyPrinted(cxPolicy PrintingPolicy) string {
	c_cT := cT
	c_cxPolicy := cxPolicy

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
		unsafe.Pointer(&c_cxPolicy),
	}

	err := ffi.CallFunction(
		cif_clang_getTypePrettyPrinted,
		ptr_clang_getTypePrettyPrinted,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypePrettyPrinted", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Pretty-print the underlying type using the rules of the language of the translation unit from which it came.

If the type is invalid, an empty string is returned.
*/
/*
Wraps the C function clang_getTypeSpelling.
*/
func (cT Type) TypeSpelling() string {
	c_cT := cT

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(
		cif_clang_getTypeSpelling,
		ptr_clang_getTypeSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypeSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the underlying type of a typedef declaration.

If the cursor does not reference a typedef declaration, an invalid type is returned.
*/
/*
Wraps the C function clang_getTypedefDeclUnderlyingType.
*/
func (c Cursor) TypedefDeclUnderlyingType() Type {
	c_c := c

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(
		cif_clang_getTypedefDeclUnderlyingType,
		ptr_clang_getTypedefDeclUnderlyingType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypedefDeclUnderlyingType", err))
	}

	ret := retC
	return ret
}

// Returns the typedef name of the given type.
/*
Wraps the C function clang_getTypedefName.
*/
func (cT Type) TypedefName() string {
	c_cT := cT

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(
		cif_clang_getTypedefName,
		ptr_clang_getTypedefName,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypedefName", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

// Retrieve the spelling of a given CXUnaryOperatorKind.
/*
Wraps the C function clang_getUnaryOperatorKindSpelling.
*/
func (kind UnaryOperatorKind) UnaryOperatorKindSpelling() string {
	c_kind := kind

	var retC string_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_kind),
	}

	err := ffi.CallFunction(
		cif_clang_getUnaryOperatorKindSpelling,
		ptr_clang_getUnaryOperatorKindSpelling,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getUnaryOperatorKindSpelling", err))
	}

	ret := retC.CString()
	retC.DisposeString()
	return ret
}

/*
Retrieve the unqualified variant of the given type, removing as little sugar as possible.

For example, given the following series of typedefs:



Executing clang_getUnqualifiedType() on a CXType that represents DifferenceType, will desugar to a type representing Integer, that has no qualifiers.

And, executing clang_getUnqualifiedType() on the type of the first argument of the following function declaration:



Will return a type representing int, removing the const qualifier.

Sugar over array types is not desugared.

A type can be checked for qualifiers with clang_isConstQualifiedType(), clang_isVolatileQualifiedType() and clang_isRestrictQualifiedType().

A type that resulted from a call to clang_getUnqualifiedType will return false for all of the above calls.
*/
/*
Wraps the C function clang_getUnqualifiedType.
*/
func (cT Type) UnqualifiedType() Type {
	c_cT := cT

	var retC Type
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(
		cif_clang_getUnqualifiedType,
		ptr_clang_getUnqualifiedType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getUnqualifiedType", err))
	}

	ret := retC
	return ret
}

// Compute a hash value for the given cursor.
/*
Wraps the C function clang_hashCursor.
*/
func (p0 Cursor) HashCursor() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_hashCursor,
		ptr_clang_hashCursor,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_hashCursor", err))
	}

	ret := retC
	return ret
}

// Retrieve the CXSourceLocation represented by the given CXIdxLoc.
/*
Wraps the C function clang_indexLoc_getCXSourceLocation.
*/
func (loc IdxLoc) IndexLocCXSourceLocation() SourceLocation {
	c_loc := loc

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc),
	}

	err := ffi.CallFunction(
		cif_clang_indexLoc_getCXSourceLocation,
		ptr_clang_indexLoc_getCXSourceLocation,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_indexLoc_getCXSourceLocation", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the CXIdxFile, file, line, column, and offset represented by the given CXIdxLoc.

If the location refers into a macro expansion, retrieves the location of the macro expansion and if it refers into a macro argument retrieves the location of the argument.
*/
/*
Wraps the C function clang_indexLoc_getFileLocation.
*/
func (loc IdxLoc) IndexLocFileLocation() (IdxClientFile, File, uint32, uint32, uint32) {
	c_loc := loc
	var indexFile IdxClientFile
	c_indexFile := &indexFile
	var file File
	c_file := &file
	var line uint32
	c_line := &line
	var column uint32
	c_column := &column
	var offset uint32
	c_offset := &offset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc),
		unsafe.Pointer(&c_indexFile),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(
		cif_clang_indexLoc_getFileLocation,
		ptr_clang_indexLoc_getFileLocation,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_indexLoc_getFileLocation", err))
	}
	return indexFile, file, line, column, offset
}

/*
Index the given source file and the translation unit corresponding to that file via callbacks implemented through #IndexerCallbacks.













The rest of the parameters are the same as #clang_parseTranslationUnit.
*/
/*
Wraps the C function clang_indexSourceFile.
*/
func (p0 IndexAction) IndexSourceFile(client_data ClientData, index_callbacks_size uint32, index_options uint32, source_filename string, command_line_args []string, unsaved_files []UnsavedFile, tU_options uint32) (IndexerCallbacks, TranslationUnit, int32) {
	c_p0 := p0
	c_client_data := client_data
	var index_callbacks IndexerCallbacks
	c_index_callbacks := &index_callbacks
	c_index_callbacks_size := index_callbacks_size
	c_index_options := index_options
	c_source_filename, free_c_source_filename := libc.CString(source_filename)
	defer free_c_source_filename()
	c_command_line_args, free_c_command_line_args := libc.CStrings(command_line_args)
	defer free_c_command_line_args()
	c_num_command_line_args := len(command_line_args)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_num_unsaved_files := len(unsaved_files)
	var out_TU TranslationUnit
	c_out_TU := &out_TU
	c_tU_options := tU_options

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_client_data),
		unsafe.Pointer(&c_index_callbacks),
		unsafe.Pointer(&c_index_callbacks_size),
		unsafe.Pointer(&c_index_options),
		unsafe.Pointer(&c_source_filename),
		unsafe.Pointer(&c_command_line_args),
		unsafe.Pointer(&c_num_command_line_args),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_out_TU),
		unsafe.Pointer(&c_tU_options),
	}

	err := ffi.CallFunction(
		cif_clang_indexSourceFile,
		ptr_clang_indexSourceFile,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_indexSourceFile", err))
	}

	ret := retC
	return index_callbacks, out_TU, ret
}

// Same as clang_indexSourceFile but requires a full command line for command_line_args including argv[0]. This is useful if the standard library paths are relative to the binary.
/*
Wraps the C function clang_indexSourceFileFullArgv.
*/
func (p0 IndexAction) IndexSourceFileFullArgv(client_data ClientData, index_callbacks_size uint32, index_options uint32, source_filename string, command_line_args []string, unsaved_files []UnsavedFile, tU_options uint32) (IndexerCallbacks, TranslationUnit, int32) {
	c_p0 := p0
	c_client_data := client_data
	var index_callbacks IndexerCallbacks
	c_index_callbacks := &index_callbacks
	c_index_callbacks_size := index_callbacks_size
	c_index_options := index_options
	c_source_filename, free_c_source_filename := libc.CString(source_filename)
	defer free_c_source_filename()
	c_command_line_args, free_c_command_line_args := libc.CStrings(command_line_args)
	defer free_c_command_line_args()
	c_num_command_line_args := len(command_line_args)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_num_unsaved_files := len(unsaved_files)
	var out_TU TranslationUnit
	c_out_TU := &out_TU
	c_tU_options := tU_options

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_client_data),
		unsafe.Pointer(&c_index_callbacks),
		unsafe.Pointer(&c_index_callbacks_size),
		unsafe.Pointer(&c_index_options),
		unsafe.Pointer(&c_source_filename),
		unsafe.Pointer(&c_command_line_args),
		unsafe.Pointer(&c_num_command_line_args),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_out_TU),
		unsafe.Pointer(&c_tU_options),
	}

	err := ffi.CallFunction(
		cif_clang_indexSourceFileFullArgv,
		ptr_clang_indexSourceFileFullArgv,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_indexSourceFileFullArgv", err))
	}

	ret := retC
	return index_callbacks, out_TU, ret
}

/*
Index the given translation unit via callbacks implemented through #IndexerCallbacks.

The order of callback invocations is not guaranteed to be the same as when indexing a source file. The high level order will be:

-Preprocessor callbacks invocations   -Declaration/reference callbacks invocations   -Diagnostic callback invocations

The parameters are the same as #clang_indexSourceFile.
*/
/*
Wraps the C function clang_indexTranslationUnit.
*/
func (p0 IndexAction) IndexTranslationUnit(client_data ClientData, index_callbacks_size uint32, index_options uint32, p5 TranslationUnit) (IndexerCallbacks, int32) {
	c_p0 := p0
	c_client_data := client_data
	var index_callbacks IndexerCallbacks
	c_index_callbacks := &index_callbacks
	c_index_callbacks_size := index_callbacks_size
	c_index_options := index_options
	c_p5 := p5

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_client_data),
		unsafe.Pointer(&c_index_callbacks),
		unsafe.Pointer(&c_index_callbacks_size),
		unsafe.Pointer(&c_index_options),
		unsafe.Pointer(&c_p5),
	}

	err := ffi.CallFunction(
		cif_clang_indexTranslationUnit,
		ptr_clang_indexTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_indexTranslationUnit", err))
	}

	ret := retC
	return index_callbacks, ret
}

//
/*
Wraps the C function clang_index_getCXXClassDeclInfo.
*/
func (p0 *IdxDeclInfo) IndexCXXClassDeclInfo() *IdxCXXClassDeclInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getCXXClassDeclInfo,
		ptr_clang_index_getCXXClassDeclInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getCXXClassDeclInfo", err))
	}

	ret := (*IdxCXXClassDeclInfo)(retC)
	return ret
}

// For retrieving a custom CXIdxClientContainer attached to a container.
/*
Wraps the C function clang_index_getClientContainer.
*/
func (p0 *IdxContainerInfo) IndexClientContainer() IdxClientContainer {
	c_p0 := p0

	var retC IdxClientContainer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getClientContainer,
		ptr_clang_index_getClientContainer,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getClientContainer", err))
	}

	ret := retC
	return ret
}

// For retrieving a custom CXIdxClientEntity attached to an entity.
/*
Wraps the C function clang_index_getClientEntity.
*/
func (p0 *IdxEntityInfo) IndexClientEntity() IdxClientEntity {
	c_p0 := p0

	var retC IdxClientEntity
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getClientEntity,
		ptr_clang_index_getClientEntity,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getClientEntity", err))
	}

	ret := retC
	return ret
}

//
/*
Wraps the C function clang_index_getIBOutletCollectionAttrInfo.
*/
func (p0 *IdxAttrInfo) IndexIBOutletCollectionAttrInfo() *IdxIBOutletCollectionAttrInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getIBOutletCollectionAttrInfo,
		ptr_clang_index_getIBOutletCollectionAttrInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getIBOutletCollectionAttrInfo", err))
	}

	ret := (*IdxIBOutletCollectionAttrInfo)(retC)
	return ret
}

//
/*
Wraps the C function clang_index_getObjCCategoryDeclInfo.
*/
func (p0 *IdxDeclInfo) IndexObjCCategoryDeclInfo() *IdxObjCCategoryDeclInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getObjCCategoryDeclInfo,
		ptr_clang_index_getObjCCategoryDeclInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getObjCCategoryDeclInfo", err))
	}

	ret := (*IdxObjCCategoryDeclInfo)(retC)
	return ret
}

//
/*
Wraps the C function clang_index_getObjCContainerDeclInfo.
*/
func (p0 *IdxDeclInfo) IndexObjCContainerDeclInfo() *IdxObjCContainerDeclInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getObjCContainerDeclInfo,
		ptr_clang_index_getObjCContainerDeclInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getObjCContainerDeclInfo", err))
	}

	ret := (*IdxObjCContainerDeclInfo)(retC)
	return ret
}

//
/*
Wraps the C function clang_index_getObjCInterfaceDeclInfo.
*/
func (p0 *IdxDeclInfo) IndexObjCInterfaceDeclInfo() *IdxObjCInterfaceDeclInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getObjCInterfaceDeclInfo,
		ptr_clang_index_getObjCInterfaceDeclInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getObjCInterfaceDeclInfo", err))
	}

	ret := (*IdxObjCInterfaceDeclInfo)(retC)
	return ret
}

//
/*
Wraps the C function clang_index_getObjCPropertyDeclInfo.
*/
func (p0 *IdxDeclInfo) IndexObjCPropertyDeclInfo() *IdxObjCPropertyDeclInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getObjCPropertyDeclInfo,
		ptr_clang_index_getObjCPropertyDeclInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getObjCPropertyDeclInfo", err))
	}

	ret := (*IdxObjCPropertyDeclInfo)(retC)
	return ret
}

//
/*
Wraps the C function clang_index_getObjCProtocolRefListInfo.
*/
func (p0 *IdxDeclInfo) IndexObjCProtocolRefListInfo() *IdxObjCProtocolRefListInfo {
	c_p0 := p0

	var retC unsafe.Pointer
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_getObjCProtocolRefListInfo,
		ptr_clang_index_getObjCProtocolRefListInfo,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_getObjCProtocolRefListInfo", err))
	}

	ret := (*IdxObjCProtocolRefListInfo)(retC)
	return ret
}

//
/*
Wraps the C function clang_index_isEntityObjCContainerKind.
*/
func (p0 IdxEntityKind) Index_isEntityObjCContainerKind() int32 {
	c_p0 := p0

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_index_isEntityObjCContainerKind,
		ptr_clang_index_isEntityObjCContainerKind,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_isEntityObjCContainerKind", err))
	}

	ret := retC
	return ret
}

// For setting a custom CXIdxClientContainer attached to a container.
/*
Wraps the C function clang_index_setClientContainer.
*/
func (p0 *IdxContainerInfo) Index_setClientContainer(p1 IdxClientContainer) {
	c_p0 := p0
	c_p1 := p1

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_index_setClientContainer,
		ptr_clang_index_setClientContainer,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_setClientContainer", err))
	}
}

// For setting a custom CXIdxClientEntity attached to an entity.
/*
Wraps the C function clang_index_setClientEntity.
*/
func (p0 *IdxEntityInfo) Index_setClientEntity(p1 IdxClientEntity) {
	c_p0 := p0
	c_p1 := p1

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(
		cif_clang_index_setClientEntity,
		ptr_clang_index_setClientEntity,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_index_setClientEntity", err))
	}
}

// Determine whether the given cursor kind represents an attribute.
/*
Wraps the C function clang_isAttribute.
*/
func (p0 CursorKind) IsAttribute() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isAttribute,
		ptr_clang_isAttribute,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isAttribute", err))
	}

	ret := retC
	return ret
}

// Determine for two source locations if the first comes strictly before the second one in the source code.
/*
Wraps the C function clang_isBeforeInTranslationUnit.
*/
func (loc1 SourceLocation) IsBeforeInTranslationUnit(loc2 SourceLocation) uint32 {
	c_loc1 := loc1
	c_loc2 := loc2

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc1),
		unsafe.Pointer(&c_loc2),
	}

	err := ffi.CallFunction(
		cif_clang_isBeforeInTranslationUnit,
		ptr_clang_isBeforeInTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isBeforeInTranslationUnit", err))
	}

	ret := retC
	return ret
}

// Determine whether a CXType has the "const" qualifier set, without looking through typedefs that may have added "const" at a different level.
/*
Wraps the C function clang_isConstQualifiedType.
*/
func (t Type) IsConstQualifiedType() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_isConstQualifiedType,
		ptr_clang_isConstQualifiedType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isConstQualifiedType", err))
	}

	ret := retC
	return ret
}

// Determine whether the declaration pointed to by this cursor is also a definition of that entity.
/*
Wraps the C function clang_isCursorDefinition.
*/
func (p0 Cursor) IsCursorDefinition() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isCursorDefinition,
		ptr_clang_isCursorDefinition,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isCursorDefinition", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor kind represents a declaration.
/*
Wraps the C function clang_isDeclaration.
*/
func (p0 CursorKind) IsDeclaration() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isDeclaration,
		ptr_clang_isDeclaration,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isDeclaration", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor kind represents an expression.
/*
Wraps the C function clang_isExpression.
*/
func (p0 CursorKind) IsExpression() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isExpression,
		ptr_clang_isExpression,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isExpression", err))
	}

	ret := retC
	return ret
}

// Determine whether the given header is guarded against multiple inclusions, either with the conventional #ifndef/#define/#endif macro guards or with #pragma once.
/*
Wraps the C function clang_isFileMultipleIncludeGuarded.
*/
func (tu TranslationUnit) IsFileMultipleIncludeGuarded(file File) uint32 {
	c_tu := tu
	c_file := file

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu.ptr),
		unsafe.Pointer(&c_file),
	}

	err := ffi.CallFunction(
		cif_clang_isFileMultipleIncludeGuarded,
		ptr_clang_isFileMultipleIncludeGuarded,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isFileMultipleIncludeGuarded", err))
	}

	ret := retC
	return ret
}

// Return 1 if the CXType is a variadic function type, and 0 otherwise.
/*
Wraps the C function clang_isFunctionTypeVariadic.
*/
func (t Type) IsFunctionTypeVariadic() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_isFunctionTypeVariadic,
		ptr_clang_isFunctionTypeVariadic,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isFunctionTypeVariadic", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor kind represents an invalid cursor.
/*
Wraps the C function clang_isInvalid.
*/
func (p0 CursorKind) IsInvalid() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isInvalid,
		ptr_clang_isInvalid,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isInvalid", err))
	}

	ret := retC
	return ret
}

/*
Determine whether the given declaration is invalid.

A declaration is invalid if it could not be parsed successfully.
*/
/*
Wraps the C function clang_isInvalidDeclaration.
*/
func (p0 Cursor) IsInvalidDeclaration() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isInvalidDeclaration,
		ptr_clang_isInvalidDeclaration,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isInvalidDeclaration", err))
	}

	ret := retC
	return ret
}

// Return 1 if the CXType is a POD (plain old data) type, and 0  otherwise.
/*
Wraps the C function clang_isPODType.
*/
func (t Type) IsPODType() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_isPODType,
		ptr_clang_isPODType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isPODType", err))
	}

	ret := retC
	return ret
}

// * Determine whether the given cursor represents a preprocessing element, such as a preprocessor directive or macro instantiation.
/*
Wraps the C function clang_isPreprocessing.
*/
func (p0 CursorKind) IsPreprocessing() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isPreprocessing,
		ptr_clang_isPreprocessing,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isPreprocessing", err))
	}

	ret := retC
	return ret
}

/*
Determine whether the given cursor kind represents a simple reference.

Note that other kinds of cursors (such as expressions) can also refer to other cursors. Use clang_getCursorReferenced() to determine whether a particular cursor refers to another entity.
*/
/*
Wraps the C function clang_isReference.
*/
func (p0 CursorKind) IsReference() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isReference,
		ptr_clang_isReference,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isReference", err))
	}

	ret := retC
	return ret
}

// Determine whether a CXType has the "restrict" qualifier set, without looking through typedefs that may have added "restrict" at a different level.
/*
Wraps the C function clang_isRestrictQualifiedType.
*/
func (t Type) IsRestrictQualifiedType() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_isRestrictQualifiedType,
		ptr_clang_isRestrictQualifiedType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isRestrictQualifiedType", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor kind represents a statement.
/*
Wraps the C function clang_isStatement.
*/
func (p0 CursorKind) IsStatement() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isStatement,
		ptr_clang_isStatement,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isStatement", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor kind represents a translation unit.
/*
Wraps the C function clang_isTranslationUnit.
*/
func (p0 CursorKind) IsTranslationUnit() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isTranslationUnit,
		ptr_clang_isTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isTranslationUnit", err))
	}

	ret := retC
	return ret
}

// * Determine whether the given cursor represents a currently  unexposed piece of the AST (e.g., CXCursor_UnexposedStmt).
/*
Wraps the C function clang_isUnexposed.
*/
func (p0 CursorKind) IsUnexposed() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isUnexposed,
		ptr_clang_isUnexposed,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isUnexposed", err))
	}

	ret := retC
	return ret
}

// Returns 1 if the base class specified by the cursor with kind   CX_CXXBaseSpecifier is virtual.
/*
Wraps the C function clang_isVirtualBase.
*/
func (p0 Cursor) IsVirtualBase() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(
		cif_clang_isVirtualBase,
		ptr_clang_isVirtualBase,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isVirtualBase", err))
	}

	ret := retC
	return ret
}

// Determine whether a CXType has the "volatile" qualifier set, without looking through typedefs that may have added "volatile" at a different level.
/*
Wraps the C function clang_isVolatileQualifiedType.
*/
func (t Type) IsVolatileQualifiedType() uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(
		cif_clang_isVolatileQualifiedType,
		ptr_clang_isVolatileQualifiedType,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isVolatileQualifiedType", err))
	}

	ret := retC
	return ret
}

// Deserialize a set of diagnostics from a Clang diagnostics bitcode file.
/*
Wraps the C function clang_loadDiagnostics.
*/
func LoadDiagnostics(file string) (LoadDiag_Error, string, DiagnosticSet) {
	c_file, free_c_file := libc.CString(file)
	defer free_c_file()
	var error LoadDiag_Error
	c_error := &error
	var errorString string_
	c_errorString := &errorString

	var retC DiagnosticSet
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_error),
		unsafe.Pointer(&c_errorString),
	}

	err := ffi.CallFunction(
		cif_clang_loadDiagnostics,
		ptr_clang_loadDiagnostics,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_loadDiagnostics", err))
	}

	ret := retC
	return error, errorString.CString(), ret
}

// Same as clang_parseTranslationUnit2, but returns the CXTranslationUnit instead of an error code.  In case of an error this routine returns a NULL CXTranslationUnit, without further detailed error codes.
/*
Wraps the C function clang_parseTranslationUnit.
*/
func (cIdx Index) ParseTranslationUnit(source_filename string, command_line_args []string, unsaved_files []UnsavedFile, options uint32) TranslationUnit {
	c_cIdx := cIdx
	c_source_filename, free_c_source_filename := libc.CString(source_filename)
	defer free_c_source_filename()
	c_command_line_args, free_c_command_line_args := libc.CStrings(command_line_args)
	defer free_c_command_line_args()
	c_num_command_line_args := len(command_line_args)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_num_unsaved_files := len(unsaved_files)
	c_options := options

	var retC TranslationUnit
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
		unsafe.Pointer(&c_source_filename),
		unsafe.Pointer(&c_command_line_args),
		unsafe.Pointer(&c_num_command_line_args),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_parseTranslationUnit,
		ptr_clang_parseTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_parseTranslationUnit", err))
	}

	ret := retC
	return ret
}

/*
Parse the given source file and the translation unit corresponding to that file.

This routine is the main entry point for the Clang C API, providing the ability to parse a source file into a translation unit that can then be queried by other functions in the API. This routine accepts a set of command-line arguments so that the compilation can be configured in the same way that the compiler is configured on the command line.
*/
/*
Wraps the C function clang_parseTranslationUnit2.
*/
func (cIdx Index) ParseTranslationUnit2(source_filename string, command_line_args []string, unsaved_files []UnsavedFile, options uint32) (TranslationUnit, ErrorCode) {
	c_cIdx := cIdx
	c_source_filename, free_c_source_filename := libc.CString(source_filename)
	defer free_c_source_filename()
	c_command_line_args, free_c_command_line_args := libc.CStrings(command_line_args)
	defer free_c_command_line_args()
	c_num_command_line_args := len(command_line_args)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_num_unsaved_files := len(unsaved_files)
	c_options := options
	var out_TU TranslationUnit
	c_out_TU := &out_TU

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
		unsafe.Pointer(&c_source_filename),
		unsafe.Pointer(&c_command_line_args),
		unsafe.Pointer(&c_num_command_line_args),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_options),
		unsafe.Pointer(&c_out_TU),
	}

	err := ffi.CallFunction(
		cif_clang_parseTranslationUnit2,
		ptr_clang_parseTranslationUnit2,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_parseTranslationUnit2", err))
	}

	ret := retC
	return out_TU, ret
}

// Same as clang_parseTranslationUnit2 but requires a full command line for command_line_args including argv[0]. This is useful if the standard library paths are relative to the binary.
/*
Wraps the C function clang_parseTranslationUnit2FullArgv.
*/
func (cIdx Index) ParseTranslationUnit2FullArgv(source_filename string, command_line_args []string, unsaved_files []UnsavedFile, options uint32) (TranslationUnit, ErrorCode) {
	c_cIdx := cIdx
	c_source_filename, free_c_source_filename := libc.CString(source_filename)
	defer free_c_source_filename()
	c_command_line_args, free_c_command_line_args := libc.CStrings(command_line_args)
	defer free_c_command_line_args()
	c_num_command_line_args := len(command_line_args)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_num_unsaved_files := len(unsaved_files)
	c_options := options
	var out_TU TranslationUnit
	c_out_TU := &out_TU

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx.ptr),
		unsafe.Pointer(&c_source_filename),
		unsafe.Pointer(&c_command_line_args),
		unsafe.Pointer(&c_num_command_line_args),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_options),
		unsafe.Pointer(&c_out_TU),
	}

	err := ffi.CallFunction(
		cif_clang_parseTranslationUnit2FullArgv,
		ptr_clang_parseTranslationUnit2FullArgv,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_parseTranslationUnit2FullArgv", err))
	}

	ret := retC
	return out_TU, ret
}

//
/*
Wraps the C function clang_remap_dispose.
*/
func (p0 Remapping) Remap_dispose() {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_remap_dispose,
		ptr_clang_remap_dispose,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_remap_dispose", err))
	}
}

//
/*
Wraps the C function clang_remap_getFilenames.
*/
func (p0 Remapping) RemapFilenames(p1 uint32) (string, string) {
	c_p0 := p0
	c_p1 := p1
	var p2 string_
	c_p2 := &p2
	var p3 string_
	c_p3 := &p3

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
		unsafe.Pointer(&c_p1),
		unsafe.Pointer(&c_p2),
		unsafe.Pointer(&c_p3),
	}

	err := ffi.CallFunction(
		cif_clang_remap_getFilenames,
		ptr_clang_remap_getFilenames,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_remap_getFilenames", err))
	}
	return p2.CString(), p3.CString()
}

//
/*
Wraps the C function clang_remap_getNumFiles.
*/
func (p0 Remapping) RemapNumFiles() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_remap_getNumFiles,
		ptr_clang_remap_getNumFiles,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_remap_getNumFiles", err))
	}

	ret := retC
	return ret
}

/*
Reparse the source files that produced this translation unit.

This routine can be used to re-parse the source files that originally created the given translation unit, for example because those source files have changed (either on disk or as passed via unsaved_files). The source code will be reparsed with the same command-line options as it was originally parsed.

Reparsing a translation unit invalidates all cursors and source locations that refer into that translation unit. This makes reparsing a translation unit semantically equivalent to destroying the translation unit and then creating a new translation unit with the same command-line arguments. However, it may be more efficient to reparse a translation unit using this routine.
*/
/*
Wraps the C function clang_reparseTranslationUnit.
*/
func (tU TranslationUnit) ReparseTranslationUnit(unsaved_files []UnsavedFile, options uint32) int32 {
	c_tU := tU
	c_num_unsaved_files := len(unsaved_files)
	var c_unsaved_files unsafe.Pointer
	if len(unsaved_files) > 0 {
		c_unsaved_files = unsafe.Pointer(&unsaved_files[0])
	}
	c_options := options

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_num_unsaved_files),
		unsafe.Pointer(&c_unsaved_files),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_reparseTranslationUnit,
		ptr_clang_reparseTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_reparseTranslationUnit", err))
	}

	ret := retC
	return ret
}

/*
Saves a translation unit into a serialized representation of that translation unit on disk.

Any translation unit that was parsed without error can be saved into a file. The translation unit can then be deserialized into a new CXTranslationUnit with clang_createTranslationUnit() or, if it is an incomplete translation unit that corresponds to a header, used as a precompiled header when parsing other translation units.
*/
/*
Wraps the C function clang_saveTranslationUnit.
*/
func (tU TranslationUnit) SaveTranslationUnit(fileName string, options uint32) int32 {
	c_tU := tU
	c_fileName, free_c_fileName := libc.CString(fileName)
	defer free_c_fileName()
	c_options := options

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU.ptr),
		unsafe.Pointer(&c_fileName),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(
		cif_clang_saveTranslationUnit,
		ptr_clang_saveTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_saveTranslationUnit", err))
	}

	ret := retC
	return ret
}

// Sort the code-completion results in case-insensitive alphabetical order.
/*
Wraps the C function clang_sortCodeCompletionResults.
*/
func (results *CompletionResult) SortCodeCompletionResults(numResults uint32) {
	c_results := results
	c_numResults := numResults

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_results),
		unsafe.Pointer(&c_numResults),
	}

	err := ffi.CallFunction(
		cif_clang_sortCodeCompletionResults,
		ptr_clang_sortCodeCompletionResults,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_sortCodeCompletionResults", err))
	}
}

/*
Suspend a translation unit in order to free memory associated with it.

A suspended translation unit uses significantly less memory but on the other side does not support any other calls than clang_reparseTranslationUnit to resume it or clang_disposeTranslationUnit to dispose it completely.
*/
/*
Wraps the C function clang_suspendTranslationUnit.
*/
func (p0 TranslationUnit) SuspendTranslationUnit() uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0.ptr),
	}

	err := ffi.CallFunction(
		cif_clang_suspendTranslationUnit,
		ptr_clang_suspendTranslationUnit,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_suspendTranslationUnit", err))
	}

	ret := retC
	return ret
}

// Enable/disable crash recovery.
/*
Wraps the C function clang_toggleCrashRecovery.
*/
func ToggleCrashRecovery(isEnabled uint32) {
	c_isEnabled := isEnabled

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_isEnabled),
	}

	err := ffi.CallFunction(
		cif_clang_toggleCrashRecovery,
		ptr_clang_toggleCrashRecovery,
		nil,
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_toggleCrashRecovery", err))
	}
}

// not supported : clang_tokenize : param Tokens : CXToken **

/*
Visit the base classes of a type.

This function visits all the direct base classes of a the given cursor, invoking the given visitor function with the cursors of each visited base. The traversal may be ended prematurely, if the visitor returns CXFieldVisit_Break.
*/
/*
Wraps the C function clang_visitCXXBaseClasses.
*/
func (t Type) VisitCXXBaseClasses(visitor FieldVisitor, client_data ClientData) uint32 {
	c_t := t
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(
		cif_clang_visitCXXBaseClasses,
		ptr_clang_visitCXXBaseClasses,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitCXXBaseClasses", err))
	}

	ret := retC
	return ret
}

/*
Visit the class methods of a type.

This function visits all the methods of the given cursor, invoking the given visitor function with the cursors of each visited method. The traversal may be ended prematurely, if the visitor returns CXFieldVisit_Break.
*/
/*
Wraps the C function clang_visitCXXMethods.
*/
func (t Type) VisitCXXMethods(visitor FieldVisitor, client_data ClientData) uint32 {
	c_t := t
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(
		cif_clang_visitCXXMethods,
		ptr_clang_visitCXXMethods,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitCXXMethods", err))
	}

	ret := retC
	return ret
}

// Visits the children of a cursor using the specified block.  Behaves identically to clang_visitChildren() in all other respects.
/*
Wraps the C function clang_visitChildrenWithBlock.
*/
func (parent Cursor) VisitChildrenWithBlock(block CursorVisitorBlock) uint32 {
	c_parent := parent
	c_block := block

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_parent),
		unsafe.Pointer(&c_block),
	}

	err := ffi.CallFunction(
		cif_clang_visitChildrenWithBlock,
		ptr_clang_visitChildrenWithBlock,
		unsafe.Pointer(&retC),
		args,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitChildrenWithBlock", err))
	}

	ret := retC
	return ret
}
