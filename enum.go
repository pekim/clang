// This is a generated file. DO NOT EDIT.

package clang

type CXErrorCode uint32

type CXDiagnosticSeverity uint32

type CXLoadDiag_Error uint32

type CXDiagnosticDisplayOptions uint32

/*
Describes the availability of a particular entity, which indicates whether the use of this entity will result in a warning or error due to it being deprecated or unavailable.
*/
type CXAvailabilityKind uint32

/*
Describes the exception specification of a cursor.

A negative value indicates that the cursor is not a function declaration.
*/
type CXCursor_ExceptionSpecificationKind uint32

type CXChoice uint32

type CXGlobalOptFlags uint32

/*
Flags that control the creation of translation units.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when constructing the translation unit.
*/
type CXTranslationUnit_Flags uint32

/*
Flags that control how translation units are saved.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when saving the translation unit.
*/
type CXSaveTranslationUnit_Flags uint32

/*
Describes the kind of error that occurred (if any) in a call to clang_saveTranslationUnit().
*/
type CXSaveError uint32

/*
Flags that control the reparsing of translation units.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when reparsing the translation unit.
*/
type CXReparse_Flags uint32

/*
Categorizes how memory is being used by a translation unit.
*/
type CXTUResourceUsageKind uint32

/*
Describes the kind of entity that a cursor refers to.
*/
type CXCursorKind uint32

/*
Describe the linkage of the entity referred to by a cursor.
*/
type CXLinkageKind uint32

type CXVisibilityKind uint32

/*
Describe the "language" of the entity referred to by a cursor.
*/
type CXLanguageKind uint32

/*
Describe the "thread-local storage (TLS) kind" of the declaration referred to by a cursor.
*/
type CXTLSKind uint32

/*
Describes the kind of type
*/
type CXTypeKind uint32

/*
Describes the calling convention of a function type
*/
type CXCallingConv uint32

/*
Describes the kind of a template argument.

See the definition of llvm::clang::TemplateArgument::ArgKind for full element descriptions.
*/
type CXTemplateArgumentKind uint32

type CXTypeNullabilityKind uint32

/*
List the possible error codes for clang_Type_getSizeOf,   clang_Type_getAlignOf, clang_Type_getOffsetOf,   clang_Cursor_getOffsetOf, and clang_getOffsetOfBase.

A value of this enumeration type can be returned if the target type is not a valid argument to sizeof, alignof or offsetof.
*/
type CXTypeLayoutError int32

type CXRefQualifierKind uint32

/*
Represents the C++ access control level to a base class for a cursor with kind CX_CXXBaseSpecifier.
*/
type CX_CXXAccessSpecifier uint32

/*
Represents the storage classes as declared in the source. CX_SC_Invalid was added for the case that the passed cursor in not a declaration.
*/
type CX_StorageClass uint32

/*
Represents a specific kind of binary operator which can appear at a cursor.
*/
type CX_BinaryOperatorKind uint32

/*
Describes how the traversal of the children of a particular cursor should proceed after visiting a particular child cursor.

A value of this enumeration type should be returned by each CXCursorVisitor to indicate how clang_visitChildren() proceed.
*/
type CXChildVisitResult uint32

/*
Properties for the printing policy.

See clang::PrintingPolicy for more information.
*/
type CXPrintingPolicyProperty uint32

/*
Property attributes for a CXCursor_ObjCPropertyDecl.
*/
type CXObjCPropertyAttrKind uint32

/*
'Qualifiers' written next to the return and parameter types in Objective-C method declarations.
*/
type CXObjCDeclQualifierKind uint32

type CXNameRefFlags uint32

/*
Describes a kind of token.
*/
type CXTokenKind uint32

/*
Describes a single piece of text within a code-completion string.

Each "chunk" within a code-completion string (CXCompletionString) is either a piece of text with a specific "kind" that describes how that text should be interpreted by the client or is another completion string.
*/
type CXCompletionChunkKind uint32

/*
Flags that can be passed to clang_codeCompleteAt() to modify its behavior.

The enumerators in this enumeration can be bitwise-OR'd together to provide multiple options to clang_codeCompleteAt().
*/
type CXCodeComplete_Flags uint32

/*
Bits that represent the context under which completion is occurring.

The enumerators in this enumeration may be bitwise-OR'd together if multiple contexts are occurring simultaneously.
*/
type CXCompletionContext uint32

type CXEvalResultKind uint32

/*
@{
*/
type CXVisitorResult uint32

type CXResult uint32

type CXIdxEntityKind uint32

type CXIdxEntityLanguage uint32

/*
Extra C++ template information for an entity. This can apply to: CXIdxEntity_Function CXIdxEntity_CXXClass CXIdxEntity_CXXStaticMethod CXIdxEntity_CXXInstanceMethod CXIdxEntity_CXXConstructor CXIdxEntity_CXXConversionFunction CXIdxEntity_CXXTypeAlias
*/
type CXIdxEntityCXXTemplateKind uint32

type CXIdxAttrKind uint32

type CXIdxDeclInfoFlags uint32

type CXIdxObjCContainerKind uint32

/*
Data for IndexerCallbacks#indexEntityReference.

This may be deprecated in a future version as this duplicates the CXSymbolRole_Implicit bit in CXSymbolRole.
*/
type CXIdxEntityRefKind uint32

/*
Roles that are attributed to symbol occurrences.

Internal: this currently mirrors low 9 bits of clang::index::SymbolRole with higher bits zeroed. These high bits may be exposed in the future.
*/
type CXSymbolRole uint32

type CXIndexOptFlags uint32

/*
Describes the kind of binary operators.
*/
type CXBinaryOperatorKind uint32

/*
Describes the kind of unary operators.
*/
type CXUnaryOperatorKind uint32
