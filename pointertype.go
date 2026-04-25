// This is a generated file. DO NOT EDIT.

package clang

import "unsafe"

type VirtualFileOverlay unsafe.Pointer

type ModuleMapDescriptor unsafe.Pointer

type File unsafe.Pointer

type Diagnostic unsafe.Pointer

type DiagnosticSet unsafe.Pointer

// An "index" that consists of a set of translation units that would typically be linked together into an executable or library.
type Index unsafe.Pointer

// An opaque type representing target information for a given translation unit.
type TargetInfo unsafe.Pointer

// A single translation unit, which resides in an index.
type TranslationUnit unsafe.Pointer

// Opaque pointer representing client data that will be passed through to various callbacks and visitors.
type ClientData unsafe.Pointer

// A fast container representing a set of CXCursors.
type CursorSet unsafe.Pointer

/*
Visitor invoked for each cursor found by a traversal.

This visitor function will be invoked for each cursor found by clang_visitCursorChildren(). Its first argument is the cursor being visited, its second argument is the parent visitor for that cursor, and its third argument is the client data provided to clang_visitCursorChildren().

The visitor should return one of the CXChildVisitResult values to direct clang_visitCursorChildren().
*/
type CursorVisitor unsafe.Pointer

type CursorVisitorBlock unsafe.Pointer

// Opaque pointer representing a policy that controls pretty printing for clang_getCursorPrettyPrinted.
type PrintingPolicy unsafe.Pointer

/*
The functions in this group provide access to information about modules.

@{
*/
type Module unsafe.Pointer

/*
A semantic string that describes a code-completion result.

A semantic string that describes the formatting of a code-completion result as a single "template" of text that should be inserted into the source buffer when a particular code-completion result is selected. Each semantic string is made up of some number of "chunks", each of which contains some text along with a description of what that text means, e.g., the name of the entity being referenced, whether the text chunk is part of the template, or whether it is a "placeholder" that the user should replace with actual code,of a specific kind. See CXCompletionChunkKind for a description of the different kinds of chunks.
*/
type CompletionString unsafe.Pointer

/*
Visitor invoked for each file in a translation unit        (used with clang_getInclusions()).

This visitor function will be invoked by clang_getInclusions() for each file included (either at the top-level or by #include directives) within a translation unit.  The first argument is the file being included, and the second and third arguments provide the inclusion stack.  The array is sorted in order of immediate inclusion.  For example, the first element refers to the location that included 'included_file'.
*/
type InclusionVisitor unsafe.Pointer

// Evaluation result of a cursor
type EvalResult unsafe.Pointer

type CursorAndRangeVisitorBlock unsafe.Pointer

// The client's data object that is associated with a CXFile.
type IdxClientFile unsafe.Pointer

// The client's data object that is associated with a semantic entity.
type IdxClientEntity unsafe.Pointer

// The client's data object that is associated with a semantic container of entities.
type IdxClientContainer unsafe.Pointer

// The client's data object that is associated with an AST file (PCH or module).
type IdxClientASTFile unsafe.Pointer

// An indexing action/session, to be applied to one or multiple translation units.
type IndexAction unsafe.Pointer

/*
Visitor invoked for each field found by a traversal.

This visitor function will be invoked for each field found by clang_Type_visitFields. Its first argument is the cursor being visited, its second argument is the client data provided to clang_Type_visitFields.

The visitor should return one of the CXVisitorResult values to direct clang_Type_visitFields.
*/
type FieldVisitor unsafe.Pointer

// @}
type Remapping unsafe.Pointer
