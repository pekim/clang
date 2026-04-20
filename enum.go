// This is a generated file. DO NOT EDIT.

package clang

type ErrorCode uint32

const (
	Error_Success          ErrorCode = 0
	Error_Failure          ErrorCode = 1
	Error_Crashed          ErrorCode = 2
	Error_InvalidArguments ErrorCode = 3
	Error_ASTReadError     ErrorCode = 4
)

type DiagnosticSeverity uint32

const (
	Diagnostic_Ignored DiagnosticSeverity = 0
	Diagnostic_Note    DiagnosticSeverity = 1
	Diagnostic_Warning DiagnosticSeverity = 2
	Diagnostic_Error   DiagnosticSeverity = 3
	Diagnostic_Fatal   DiagnosticSeverity = 4
)

type LoadDiag_Error uint32

const (
	LoadDiag_None        LoadDiag_Error = 0
	LoadDiag_Unknown     LoadDiag_Error = 1
	LoadDiag_CannotLoad  LoadDiag_Error = 2
	LoadDiag_InvalidFile LoadDiag_Error = 3
)

type DiagnosticDisplayOptions uint32

const (
	Diagnostic_DisplaySourceLocation DiagnosticDisplayOptions = 1
	Diagnostic_DisplayColumn         DiagnosticDisplayOptions = 2
	Diagnostic_DisplaySourceRanges   DiagnosticDisplayOptions = 4
	Diagnostic_DisplayOption         DiagnosticDisplayOptions = 8
	Diagnostic_DisplayCategoryId     DiagnosticDisplayOptions = 16
	Diagnostic_DisplayCategoryName   DiagnosticDisplayOptions = 32
)

/*
Describes the availability of a particular entity, which indicates whether the use of this entity will result in a warning or error due to it being deprecated or unavailable.
*/
type AvailabilityKind uint32

const (
	/*
	   The entity is available.
	*/
	Availability_Available AvailabilityKind = 0
	/*
	   The entity is available, but has been deprecated (and its use is not recommended).
	*/
	Availability_Deprecated AvailabilityKind = 1
	/*
	   The entity is not available; any use of it will be an error.
	*/
	Availability_NotAvailable AvailabilityKind = 2
	/*
	   The entity is available, but not accessible; any use of it will be an error.
	*/
	Availability_NotAccessible AvailabilityKind = 3
)

/*
Describes the exception specification of a cursor.

A negative value indicates that the cursor is not a function declaration.
*/
type Cursor_ExceptionSpecificationKind uint32

const (
	/*
	   The cursor has no exception specification.
	*/
	Cursor_ExceptionSpecificationKind_None Cursor_ExceptionSpecificationKind = 0
	/*
	   The cursor has exception specification throw()
	*/
	Cursor_ExceptionSpecificationKind_DynamicNone Cursor_ExceptionSpecificationKind = 1
	/*
	   The cursor has exception specification throw(T1, T2)
	*/
	Cursor_ExceptionSpecificationKind_Dynamic Cursor_ExceptionSpecificationKind = 2
	/*
	   The cursor has exception specification throw(...).
	*/
	Cursor_ExceptionSpecificationKind_MSAny Cursor_ExceptionSpecificationKind = 3
	/*
	   The cursor has exception specification basic noexcept.
	*/
	Cursor_ExceptionSpecificationKind_BasicNoexcept Cursor_ExceptionSpecificationKind = 4
	/*
	   The cursor has exception specification computed noexcept.
	*/
	Cursor_ExceptionSpecificationKind_ComputedNoexcept Cursor_ExceptionSpecificationKind = 5
	/*
	   The exception specification has not yet been evaluated.
	*/
	Cursor_ExceptionSpecificationKind_Unevaluated Cursor_ExceptionSpecificationKind = 6
	/*
	   The exception specification has not yet been instantiated.
	*/
	Cursor_ExceptionSpecificationKind_Uninstantiated Cursor_ExceptionSpecificationKind = 7
	/*
	   The exception specification has not been parsed yet.
	*/
	Cursor_ExceptionSpecificationKind_Unparsed Cursor_ExceptionSpecificationKind = 8
	/*
	   The cursor has a __declspec(nothrow) exception specification.
	*/
	Cursor_ExceptionSpecificationKind_NoThrow Cursor_ExceptionSpecificationKind = 9
)

type Choice uint32

const (
	/*
	   Use the default value of an option that may depend on the process environment.
	*/
	Choice_Default Choice = 0
	/*
	   Enable the option.
	*/
	Choice_Enabled Choice = 1
	/*
	   Disable the option.
	*/
	Choice_Disabled Choice = 2
)

type GlobalOptFlags uint32

const (
	/*
	   Used to indicate that no special CXIndex options are needed.
	*/
	GlobalOpt_None GlobalOptFlags = 0
	/*
	   Used to indicate that threads that libclang creates for indexing purposes should use background priority.

	   Affects #clang_indexSourceFile, #clang_indexTranslationUnit, #clang_parseTranslationUnit, #clang_saveTranslationUnit.
	*/
	GlobalOpt_ThreadBackgroundPriorityForIndexing GlobalOptFlags = 1
	/*
	   Used to indicate that threads that libclang creates for editing purposes should use background priority.

	   Affects #clang_reparseTranslationUnit, #clang_codeCompleteAt, #clang_annotateTokens
	*/
	GlobalOpt_ThreadBackgroundPriorityForEditing GlobalOptFlags = 2
	/*
	   Used to indicate that all threads that libclang creates should use background priority.
	*/
	GlobalOpt_ThreadBackgroundPriorityForAll GlobalOptFlags = 3
)

/*
Flags that control the creation of translation units.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when constructing the translation unit.
*/
type TranslationUnit_Flags uint32

const (
	/*
	   Used to indicate that no special translation-unit options are needed.
	*/
	TranslationUnit_None TranslationUnit_Flags = 0
	/*
	   Used to indicate that the parser should construct a "detailed" preprocessing record, including all macro definitions and instantiations.

	   Constructing a detailed preprocessing record requires more memory and time to parse, since the information contained in the record is usually not retained. However, it can be useful for applications that require more detailed information about the behavior of the preprocessor.
	*/
	TranslationUnit_DetailedPreprocessingRecord TranslationUnit_Flags = 1
	/*
	   Used to indicate that the translation unit is incomplete.

	   When a translation unit is considered "incomplete", semantic analysis that is typically performed at the end of the translation unit will be suppressed. For example, this suppresses the completion of tentative declarations in C and of instantiation of implicitly-instantiation function templates in C++. This option is typically used when parsing a header with the intent of producing a precompiled header.
	*/
	TranslationUnit_Incomplete TranslationUnit_Flags = 2
	/*
	   Used to indicate that the translation unit should be built with an implicit precompiled header for the preamble.

	   An implicit precompiled header is used as an optimization when a particular translation unit is likely to be reparsed many times when the sources aren't changing that often. In this case, an implicit precompiled header will be built containing all of the initial includes at the top of the main file (what we refer to as the "preamble" of the file). In subsequent parses, if the preamble or the files in it have not changed, clang_reparseTranslationUnit() will re-use the implicit precompiled header to improve parsing performance.
	*/
	TranslationUnit_PrecompiledPreamble TranslationUnit_Flags = 4
	/*
	   Used to indicate that the translation unit should cache some code-completion results with each reparse of the source file.

	   Caching of code-completion results is a performance optimization that introduces some overhead to reparsing but improves the performance of code-completion operations.
	*/
	TranslationUnit_CacheCompletionResults TranslationUnit_Flags = 8
	/*
	   Used to indicate that the translation unit will be serialized with clang_saveTranslationUnit.

	   This option is typically used when parsing a header with the intent of producing a precompiled header.
	*/
	TranslationUnit_ForSerialization TranslationUnit_Flags = 16
	/*
	   DEPRECATED: Enabled chained precompiled preambles in C++.

	   Note: this is a *temporary* option that is available only while we are testing C++ precompiled preamble support. It is deprecated.
	*/
	TranslationUnit_CXXChainedPCH TranslationUnit_Flags = 32
	/*
	   Used to indicate that function/method bodies should be skipped while parsing.

	   This option can be used to search for declarations/definitions while ignoring the usages.
	*/
	TranslationUnit_SkipFunctionBodies TranslationUnit_Flags = 64
	/*
	   Used to indicate that brief documentation comments should be included into the set of code completions returned from this translation unit.
	*/
	TranslationUnit_IncludeBriefCommentsInCodeCompletion TranslationUnit_Flags = 128
	/*
	   Used to indicate that the precompiled preamble should be created on the first parse. Otherwise it will be created on the first reparse. This trades runtime on the first parse (serializing the preamble takes time) for reduced runtime on the second parse (can now reuse the preamble).
	*/
	TranslationUnit_CreatePreambleOnFirstParse TranslationUnit_Flags = 256
	/*
	   Do not stop processing when fatal errors are encountered.

	   When fatal errors are encountered while parsing a translation unit, semantic analysis is typically stopped early when compiling code. A common source for fatal errors are unresolvable include files. For the purposes of an IDE, this is undesirable behavior and as much information as possible should be reported. Use this flag to enable this behavior.
	*/
	TranslationUnit_KeepGoing TranslationUnit_Flags = 512
	/*
	   Sets the preprocessor in a mode for parsing a single file only.
	*/
	TranslationUnit_SingleFileParse TranslationUnit_Flags = 1024
	/*
	   Used in combination with CXTranslationUnit_SkipFunctionBodies to constrain the skipping of function bodies to the preamble.

	   The function bodies of the main file are not skipped.
	*/
	TranslationUnit_LimitSkipFunctionBodiesToPreamble TranslationUnit_Flags = 2048
	/*
	   Used to indicate that attributed types should be included in CXType.
	*/
	TranslationUnit_IncludeAttributedTypes TranslationUnit_Flags = 4096
	/*
	   Used to indicate that implicit attributes should be visited.
	*/
	TranslationUnit_VisitImplicitAttributes TranslationUnit_Flags = 8192
	/*
	   Used to indicate that non-errors from included files should be ignored.

	   If set, clang_getDiagnosticSetFromTU() will not report e.g. warnings from included files anymore. This speeds up clang_getDiagnosticSetFromTU() for the case where these warnings are not of interest, as for an IDE for example, which typically shows only the diagnostics in the main file.
	*/
	TranslationUnit_IgnoreNonErrorsFromIncludedFiles TranslationUnit_Flags = 16384
	/*
	   Tells the preprocessor not to skip excluded conditional blocks.
	*/
	TranslationUnit_RetainExcludedConditionalBlocks TranslationUnit_Flags = 32768
)

/*
Flags that control how translation units are saved.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when saving the translation unit.
*/
type SaveTranslationUnit_Flags uint32

const (
	/*
	   Used to indicate that no special saving options are needed.
	*/
	SaveTranslationUnit_None SaveTranslationUnit_Flags = 0
)

/*
Describes the kind of error that occurred (if any) in a call to clang_saveTranslationUnit().
*/
type SaveError uint32

const (
	/*
	   Indicates that no error occurred while saving a translation unit.
	*/
	SaveError_None SaveError = 0
	/*
	   Indicates that an unknown error occurred while attempting to save the file.

	   This error typically indicates that file I/O failed when attempting to write the file.
	*/
	SaveError_Unknown SaveError = 1
	/*
	   Indicates that errors during translation prevented this attempt to save the translation unit.

	   Errors that prevent the translation unit from being saved can be extracted using clang_getNumDiagnostics() and clang_getDiagnostic().
	*/
	SaveError_TranslationErrors SaveError = 2
	/*
	   Indicates that the translation unit to be saved was somehow invalid (e.g., NULL).
	*/
	SaveError_InvalidTU SaveError = 3
)

/*
Flags that control the reparsing of translation units.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when reparsing the translation unit.
*/
type Reparse_Flags uint32

const (
	/*
	   Used to indicate that no special reparsing options are needed.
	*/
	Reparse_None Reparse_Flags = 0
)

/*
Categorizes how memory is being used by a translation unit.
*/
type TUResourceUsageKind uint32

const (
	TUResourceUsage_AST                                TUResourceUsageKind = 1
	TUResourceUsage_Identifiers                        TUResourceUsageKind = 2
	TUResourceUsage_Selectors                          TUResourceUsageKind = 3
	TUResourceUsage_GlobalCompletionResults            TUResourceUsageKind = 4
	TUResourceUsage_SourceManagerContentCache          TUResourceUsageKind = 5
	TUResourceUsage_AST_SideTables                     TUResourceUsageKind = 6
	TUResourceUsage_SourceManager_Membuffer_Malloc     TUResourceUsageKind = 7
	TUResourceUsage_SourceManager_Membuffer_MMap       TUResourceUsageKind = 8
	TUResourceUsage_ExternalASTSource_Membuffer_Malloc TUResourceUsageKind = 9
	TUResourceUsage_ExternalASTSource_Membuffer_MMap   TUResourceUsageKind = 10
	TUResourceUsage_Preprocessor                       TUResourceUsageKind = 11
	TUResourceUsage_PreprocessingRecord                TUResourceUsageKind = 12
	TUResourceUsage_SourceManager_DataStructures       TUResourceUsageKind = 13
	TUResourceUsage_Preprocessor_HeaderSearch          TUResourceUsageKind = 14
	TUResourceUsage_MEMORY_IN_BYTES_BEGIN              TUResourceUsageKind = 1
	TUResourceUsage_MEMORY_IN_BYTES_END                TUResourceUsageKind = 14
	TUResourceUsage_First                              TUResourceUsageKind = 1
	TUResourceUsage_Last                               TUResourceUsageKind = 14
)

/*
Describes the kind of entity that a cursor refers to.
*/
type CursorKind uint32

const (
	/*
	   A declaration whose specific kind is not exposed via this interface.

	   Unexposed declarations have the same operations as any other kind of declaration; one can extract their location information, spelling, find their definitions, etc. However, the specific kind of the declaration is not reported.
	*/
	Cursor_UnexposedDecl CursorKind = 1
	/*
	   A C or C++ struct.
	*/
	Cursor_StructDecl CursorKind = 2
	/*
	   A C or C++ union.
	*/
	Cursor_UnionDecl CursorKind = 3
	/*
	   A C++ class.
	*/
	Cursor_ClassDecl CursorKind = 4
	/*
	   An enumeration.
	*/
	Cursor_EnumDecl CursorKind = 5
	/*
	   A field (in C) or non-static data member (in C++) in a struct, union, or C++ class.
	*/
	Cursor_FieldDecl CursorKind = 6
	/*
	   An enumerator constant.
	*/
	Cursor_EnumConstantDecl CursorKind = 7
	/*
	   A function.
	*/
	Cursor_FunctionDecl CursorKind = 8
	/*
	   A variable.
	*/
	Cursor_VarDecl CursorKind = 9
	/*
	   A function or method parameter.
	*/
	Cursor_ParmDecl CursorKind = 10
	/*
	   An Objective-C @interface.
	*/
	Cursor_ObjCInterfaceDecl CursorKind = 11
	/*
	   An Objective-C @interface for a category.
	*/
	Cursor_ObjCCategoryDecl CursorKind = 12
	/*
	   An Objective-C @protocol declaration.
	*/
	Cursor_ObjCProtocolDecl CursorKind = 13
	/*
	   An Objective-C @property declaration.
	*/
	Cursor_ObjCPropertyDecl CursorKind = 14
	/*
	   An Objective-C instance variable.
	*/
	Cursor_ObjCIvarDecl CursorKind = 15
	/*
	   An Objective-C instance method.
	*/
	Cursor_ObjCInstanceMethodDecl CursorKind = 16
	/*
	   An Objective-C class method.
	*/
	Cursor_ObjCClassMethodDecl CursorKind = 17
	/*
	   An Objective-C @implementation.
	*/
	Cursor_ObjCImplementationDecl CursorKind = 18
	/*
	   An Objective-C @implementation for a category.
	*/
	Cursor_ObjCCategoryImplDecl CursorKind = 19
	/*
	   A typedef.
	*/
	Cursor_TypedefDecl CursorKind = 20
	/*
	   A C++ class method.
	*/
	Cursor_CXXMethod CursorKind = 21
	/*
	   A C++ namespace.
	*/
	Cursor_Namespace CursorKind = 22
	/*
	   A linkage specification, e.g. 'extern "C"'.
	*/
	Cursor_LinkageSpec CursorKind = 23
	/*
	   A C++ constructor.
	*/
	Cursor_Constructor CursorKind = 24
	/*
	   A C++ destructor.
	*/
	Cursor_Destructor CursorKind = 25
	/*
	   A C++ conversion function.
	*/
	Cursor_ConversionFunction CursorKind = 26
	/*
	   A C++ template type parameter.
	*/
	Cursor_TemplateTypeParameter CursorKind = 27
	/*
	   A C++ non-type template parameter.
	*/
	Cursor_NonTypeTemplateParameter CursorKind = 28
	/*
	   A C++ template template parameter.
	*/
	Cursor_TemplateTemplateParameter CursorKind = 29
	/*
	   A C++ function template.
	*/
	Cursor_FunctionTemplate CursorKind = 30
	/*
	   A C++ class template.
	*/
	Cursor_ClassTemplate CursorKind = 31
	/*
	   A C++ class template partial specialization.
	*/
	Cursor_ClassTemplatePartialSpecialization CursorKind = 32
	/*
	   A C++ namespace alias declaration.
	*/
	Cursor_NamespaceAlias CursorKind = 33
	/*
	   A C++ using directive.
	*/
	Cursor_UsingDirective CursorKind = 34
	/*
	   A C++ using declaration.
	*/
	Cursor_UsingDeclaration CursorKind = 35
	/*
	   A C++ alias declaration
	*/
	Cursor_TypeAliasDecl CursorKind = 36
	/*
	   An Objective-C @synthesize definition.
	*/
	Cursor_ObjCSynthesizeDecl CursorKind = 37
	/*
	   An Objective-C @dynamic definition.
	*/
	Cursor_ObjCDynamicDecl CursorKind = 38
	/*
	   An access specifier.
	*/
	Cursor_CXXAccessSpecifier CursorKind = 39
	/*
	   An access specifier.
	*/
	Cursor_FirstDecl CursorKind = 1
	/*
	   An access specifier.
	*/
	Cursor_LastDecl CursorKind = 39
	/*
	   An access specifier.
	*/
	Cursor_FirstRef CursorKind = 40
	/*
	   An access specifier.
	*/
	Cursor_ObjCSuperClassRef CursorKind = 40
	/*
	   An access specifier.
	*/
	Cursor_ObjCProtocolRef CursorKind = 41
	/*
	   An access specifier.
	*/
	Cursor_ObjCClassRef CursorKind = 42
	/*
	   A reference to a type declaration.

	   A type reference occurs anywhere where a type is named but not declared. For example, given:

	   The typedef is a declaration of size_type (CXCursor_TypedefDecl), while the type of the variable "size" is referenced. The cursor referenced by the type of size is the typedef for size_type.
	*/
	Cursor_TypeRef CursorKind = 43
	/*
	   A reference to a type declaration.

	   A type reference occurs anywhere where a type is named but not declared. For example, given:

	   The typedef is a declaration of size_type (CXCursor_TypedefDecl), while the type of the variable "size" is referenced. The cursor referenced by the type of size is the typedef for size_type.
	*/
	Cursor_CXXBaseSpecifier CursorKind = 44
	/*
	   A reference to a class template, function template, template template parameter, or class template partial specialization.
	*/
	Cursor_TemplateRef CursorKind = 45
	/*
	   A reference to a namespace or namespace alias.
	*/
	Cursor_NamespaceRef CursorKind = 46
	/*
	   A reference to a member of a struct, union, or class that occurs in some non-expression context, e.g., a designated initializer.
	*/
	Cursor_MemberRef CursorKind = 47
	/*
	   A reference to a labeled statement.

	   This cursor kind is used to describe the jump to "start_over" in the goto statement in the following example:

	   A label reference cursor refers to a label statement.
	*/
	Cursor_LabelRef CursorKind = 48
	/*
	   A reference to a set of overloaded functions or function templates that has not yet been resolved to a specific function or function template.

	   An overloaded declaration reference cursor occurs in C++ templates where a dependent name refers to a function. For example:

	   Here, the identifier "swap" is associated with an overloaded declaration reference. In the template definition, "swap" refers to either of the two "swap" functions declared above, so both results will be available. At instantiation time, "swap" may also refer to other functions found via argument-dependent lookup (e.g., the "swap" function at the end of the example).

	   The functions clang_getNumOverloadedDecls() and clang_getOverloadedDecl() can be used to retrieve the definitions referenced by this cursor.
	*/
	Cursor_OverloadedDeclRef CursorKind = 49
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_VariableRef CursorKind = 50
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_LastRef CursorKind = 50
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_FirstInvalid CursorKind = 70
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_InvalidFile CursorKind = 70
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_NoDeclFound CursorKind = 71
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_NotImplemented CursorKind = 72
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_InvalidCode CursorKind = 73
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_LastInvalid CursorKind = 73
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	Cursor_FirstExpr CursorKind = 100
	/*
	   An expression whose specific kind is not exposed via this interface.

	   Unexposed expressions have the same operations as any other kind of expression; one can extract their location information, spelling, children, etc. However, the specific kind of the expression is not reported.
	*/
	Cursor_UnexposedExpr CursorKind = 100
	/*
	   An expression that refers to some value declaration, such as a function, variable, or enumerator.
	*/
	Cursor_DeclRefExpr CursorKind = 101
	/*
	   An expression that refers to a member of a struct, union, class, Objective-C class, etc.
	*/
	Cursor_MemberRefExpr CursorKind = 102
	/*
	   An expression that calls a function.
	*/
	Cursor_CallExpr CursorKind = 103
	/*
	   An expression that sends a message to an Objective-C   object or class.
	*/
	Cursor_ObjCMessageExpr CursorKind = 104
	/*
	   An expression that represents a block literal.
	*/
	Cursor_BlockExpr CursorKind = 105
	/*
	   An integer literal.
	*/
	Cursor_IntegerLiteral CursorKind = 106
	/*
	   A floating point number literal.
	*/
	Cursor_FloatingLiteral CursorKind = 107
	/*
	   An imaginary number literal.
	*/
	Cursor_ImaginaryLiteral CursorKind = 108
	/*
	   A string literal.
	*/
	Cursor_StringLiteral CursorKind = 109
	/*
	   A character literal.
	*/
	Cursor_CharacterLiteral CursorKind = 110
	/*
	   A parenthesized expression, e.g. "(1)".

	   This AST node is only formed if full location information is requested.
	*/
	Cursor_ParenExpr CursorKind = 111
	/*
	   This represents the unary-expression's (except sizeof and alignof).
	*/
	Cursor_UnaryOperator CursorKind = 112
	/*
	   [C99 6.5.2.1] Array Subscripting.
	*/
	Cursor_ArraySubscriptExpr CursorKind = 113
	/*
	   A builtin binary operation expression such as "x + y" or "x <= y".
	*/
	Cursor_BinaryOperator CursorKind = 114
	/*
	   Compound assignment such as "+=".
	*/
	Cursor_CompoundAssignOperator CursorKind = 115
	/*
	   The ?: ternary operator.
	*/
	Cursor_ConditionalOperator CursorKind = 116
	/*
	   An explicit cast in C (C99 6.5.4) or a C-style cast in C++ (C++ [expr.cast]), which uses the syntax (Type)expr.

	   For example: (int)f.
	*/
	Cursor_CStyleCastExpr CursorKind = 117
	/*
	   [C99 6.5.2.5]
	*/
	Cursor_CompoundLiteralExpr CursorKind = 118
	/*
	   Describes an C or C++ initializer list.
	*/
	Cursor_InitListExpr CursorKind = 119
	/*
	   The GNU address of label extension, representing &&label.
	*/
	Cursor_AddrLabelExpr CursorKind = 120
	/*
	   This is the GNU Statement Expression extension: ({int X=4; X;})
	*/
	Cursor_StmtExpr CursorKind = 121
	/*
	   Represents a C11 generic selection.
	*/
	Cursor_GenericSelectionExpr CursorKind = 122
	/*
	   Implements the GNU __null extension, which is a name for a null pointer constant that has integral type (e.g., int or long) and is the same size and alignment as a pointer.

	   The __null extension is typically only used by system headers, which define NULL as __null in C++ rather than using 0 (which is an integer that may not match the size of a pointer).
	*/
	Cursor_GNUNullExpr CursorKind = 123
	/*
	   C++'s static_cast<> expression.
	*/
	Cursor_CXXStaticCastExpr CursorKind = 124
	/*
	   C++'s dynamic_cast<> expression.
	*/
	Cursor_CXXDynamicCastExpr CursorKind = 125
	/*
	   C++'s reinterpret_cast<> expression.
	*/
	Cursor_CXXReinterpretCastExpr CursorKind = 126
	/*
	   C++'s const_cast<> expression.
	*/
	Cursor_CXXConstCastExpr CursorKind = 127
	/*
	   Represents an explicit C++ type conversion that uses "functional" notion (C++ [expr.type.conv]).

	   Example:
	*/
	Cursor_CXXFunctionalCastExpr CursorKind = 128
	/*
	   A C++ typeid expression (C++ [expr.typeid]).
	*/
	Cursor_CXXTypeidExpr CursorKind = 129
	/*
	   [C++ 2.13.5] C++ Boolean Literal.
	*/
	Cursor_CXXBoolLiteralExpr CursorKind = 130
	/*
	   [C++0x 2.14.7] C++ Pointer Literal.
	*/
	Cursor_CXXNullPtrLiteralExpr CursorKind = 131
	/*
	   Represents the "this" expression in C++
	*/
	Cursor_CXXThisExpr CursorKind = 132
	/*
	   [C++ 15] C++ Throw Expression.

	   This handles 'throw' and 'throw' assignment-expression. When assignment-expression isn't present, Op will be null.
	*/
	Cursor_CXXThrowExpr CursorKind = 133
	/*
	   A new expression for memory allocation and constructor calls, e.g: "new CXXNewExpr(foo)".
	*/
	Cursor_CXXNewExpr CursorKind = 134
	/*
	   A delete expression for memory deallocation and destructor calls, e.g. "delete[] pArray".
	*/
	Cursor_CXXDeleteExpr CursorKind = 135
	/*
	   A unary expression. (noexcept, sizeof, or other traits)
	*/
	Cursor_UnaryExpr CursorKind = 136
	/*
	   An Objective-C string literal i.e. "foo".
	*/
	Cursor_ObjCStringLiteral CursorKind = 137
	/*
	   An Objective-C @encode expression.
	*/
	Cursor_ObjCEncodeExpr CursorKind = 138
	/*
	   An Objective-C @selector expression.
	*/
	Cursor_ObjCSelectorExpr CursorKind = 139
	/*
	   An Objective-C @protocol expression.
	*/
	Cursor_ObjCProtocolExpr CursorKind = 140
	/*
	   An Objective-C "bridged" cast expression, which casts between Objective-C pointers and C pointers, transferring ownership in the process.
	*/
	Cursor_ObjCBridgedCastExpr CursorKind = 141
	/*
	   Represents a C++0x pack expansion that produces a sequence of expressions.

	   A pack expansion expression contains a pattern (which itself is an expression) followed by an ellipsis. For example:
	*/
	Cursor_PackExpansionExpr CursorKind = 142
	/*
	   Represents an expression that computes the length of a parameter pack.
	*/
	Cursor_SizeOfPackExpr CursorKind = 143
	Cursor_LambdaExpr     CursorKind = 144
	/*
	   Objective-c Boolean Literal.
	*/
	Cursor_ObjCBoolLiteralExpr CursorKind = 145
	/*
	   Represents the "self" expression in an Objective-C method.
	*/
	Cursor_ObjCSelfExpr CursorKind = 146
	/*
	   OpenMP 5.0 [2.1.5, Array Section]. OpenACC 3.3 [2.7.1, Data Specification for Data Clauses (Sub Arrays)]
	*/
	Cursor_ArraySectionExpr CursorKind = 147
	/*
	   Represents an (...) check.
	*/
	Cursor_ObjCAvailabilityCheckExpr CursorKind = 148
	/*
	   Fixed point literal
	*/
	Cursor_FixedPointLiteral CursorKind = 149
	/*
	   OpenMP 5.0 [2.1.4, Array Shaping].
	*/
	Cursor_OMPArrayShapingExpr CursorKind = 150
	/*
	   OpenMP 5.0 [2.1.6 Iterators]
	*/
	Cursor_OMPIteratorExpr CursorKind = 151
	/*
	   OpenCL's addrspace_cast<> expression.
	*/
	Cursor_CXXAddrspaceCastExpr CursorKind = 152
	/*
	   Expression that references a C++20 concept.
	*/
	Cursor_ConceptSpecializationExpr CursorKind = 153
	/*
	   Expression that references a C++20 requires expression.
	*/
	Cursor_RequiresExpr CursorKind = 154
	/*
	   Expression that references a C++20 parenthesized list aggregate initializer.
	*/
	Cursor_CXXParenListInitExpr CursorKind = 155
	/*
	   Represents a C++26 pack indexing expression.
	*/
	Cursor_PackIndexingExpr CursorKind = 156
	/*
	   Represents a C++26 pack indexing expression.
	*/
	Cursor_LastExpr CursorKind = 156
	/*
	   Represents a C++26 pack indexing expression.
	*/
	Cursor_FirstStmt CursorKind = 200
	/*
	   A statement whose specific kind is not exposed via this interface.

	   Unexposed statements have the same operations as any other kind of statement; one can extract their location information, spelling, children, etc. However, the specific kind of the statement is not reported.
	*/
	Cursor_UnexposedStmt CursorKind = 200
	/*
	   A labelled statement in a function.

	   This cursor kind is used to describe the "start_over:" label statement in the following example:
	*/
	Cursor_LabelStmt CursorKind = 201
	/*
	   A group of statements like { stmt stmt }.

	   This cursor kind is used to describe compound statements, e.g. function bodies.
	*/
	Cursor_CompoundStmt CursorKind = 202
	/*
	   A case statement.
	*/
	Cursor_CaseStmt CursorKind = 203
	/*
	   A default statement.
	*/
	Cursor_DefaultStmt CursorKind = 204
	/*
	   An if statement
	*/
	Cursor_IfStmt CursorKind = 205
	/*
	   A switch statement.
	*/
	Cursor_SwitchStmt CursorKind = 206
	/*
	   A while statement.
	*/
	Cursor_WhileStmt CursorKind = 207
	/*
	   A do statement.
	*/
	Cursor_DoStmt CursorKind = 208
	/*
	   A for statement.
	*/
	Cursor_ForStmt CursorKind = 209
	/*
	   A goto statement.
	*/
	Cursor_GotoStmt CursorKind = 210
	/*
	   An indirect goto statement.
	*/
	Cursor_IndirectGotoStmt CursorKind = 211
	/*
	   A continue statement.
	*/
	Cursor_ContinueStmt CursorKind = 212
	/*
	   A break statement.
	*/
	Cursor_BreakStmt CursorKind = 213
	/*
	   A return statement.
	*/
	Cursor_ReturnStmt CursorKind = 214
	/*
	   A GCC inline assembly statement extension.
	*/
	Cursor_GCCAsmStmt CursorKind = 215
	/*
	   A GCC inline assembly statement extension.
	*/
	Cursor_AsmStmt CursorKind = 215
	/*
	   Objective-C's overall @try-@catch-@finally statement.
	*/
	Cursor_ObjCAtTryStmt CursorKind = 216
	/*
	   Objective-C's @catch statement.
	*/
	Cursor_ObjCAtCatchStmt CursorKind = 217
	/*
	   Objective-C's @finally statement.
	*/
	Cursor_ObjCAtFinallyStmt CursorKind = 218
	/*
	   Objective-C's @throw statement.
	*/
	Cursor_ObjCAtThrowStmt CursorKind = 219
	/*
	   Objective-C's @synchronized statement.
	*/
	Cursor_ObjCAtSynchronizedStmt CursorKind = 220
	/*
	   Objective-C's autorelease pool statement.
	*/
	Cursor_ObjCAutoreleasePoolStmt CursorKind = 221
	/*
	   Objective-C's collection statement.
	*/
	Cursor_ObjCForCollectionStmt CursorKind = 222
	/*
	   C++'s catch statement.
	*/
	Cursor_CXXCatchStmt CursorKind = 223
	/*
	   C++'s try statement.
	*/
	Cursor_CXXTryStmt CursorKind = 224
	/*
	   C++'s for (* : *) statement.
	*/
	Cursor_CXXForRangeStmt CursorKind = 225
	/*
	   Windows Structured Exception Handling's try statement.
	*/
	Cursor_SEHTryStmt CursorKind = 226
	/*
	   Windows Structured Exception Handling's except statement.
	*/
	Cursor_SEHExceptStmt CursorKind = 227
	/*
	   Windows Structured Exception Handling's finally statement.
	*/
	Cursor_SEHFinallyStmt CursorKind = 228
	/*
	   A MS inline assembly statement extension.
	*/
	Cursor_MSAsmStmt CursorKind = 229
	/*
	   The null statement ";": C99 6.8.3p3.

	   This cursor kind is used to describe the null statement.
	*/
	Cursor_NullStmt CursorKind = 230
	/*
	   Adaptor class for mixing declarations with statements and expressions.
	*/
	Cursor_DeclStmt CursorKind = 231
	/*
	   OpenMP parallel directive.
	*/
	Cursor_OMPParallelDirective CursorKind = 232
	/*
	   OpenMP SIMD directive.
	*/
	Cursor_OMPSimdDirective CursorKind = 233
	/*
	   OpenMP for directive.
	*/
	Cursor_OMPForDirective CursorKind = 234
	/*
	   OpenMP sections directive.
	*/
	Cursor_OMPSectionsDirective CursorKind = 235
	/*
	   OpenMP section directive.
	*/
	Cursor_OMPSectionDirective CursorKind = 236
	/*
	   OpenMP single directive.
	*/
	Cursor_OMPSingleDirective CursorKind = 237
	/*
	   OpenMP parallel for directive.
	*/
	Cursor_OMPParallelForDirective CursorKind = 238
	/*
	   OpenMP parallel sections directive.
	*/
	Cursor_OMPParallelSectionsDirective CursorKind = 239
	/*
	   OpenMP task directive.
	*/
	Cursor_OMPTaskDirective CursorKind = 240
	/*
	   OpenMP master directive.
	*/
	Cursor_OMPMasterDirective CursorKind = 241
	/*
	   OpenMP critical directive.
	*/
	Cursor_OMPCriticalDirective CursorKind = 242
	/*
	   OpenMP taskyield directive.
	*/
	Cursor_OMPTaskyieldDirective CursorKind = 243
	/*
	   OpenMP barrier directive.
	*/
	Cursor_OMPBarrierDirective CursorKind = 244
	/*
	   OpenMP taskwait directive.
	*/
	Cursor_OMPTaskwaitDirective CursorKind = 245
	/*
	   OpenMP flush directive.
	*/
	Cursor_OMPFlushDirective CursorKind = 246
	/*
	   Windows Structured Exception Handling's leave statement.
	*/
	Cursor_SEHLeaveStmt CursorKind = 247
	/*
	   OpenMP ordered directive.
	*/
	Cursor_OMPOrderedDirective CursorKind = 248
	/*
	   OpenMP atomic directive.
	*/
	Cursor_OMPAtomicDirective CursorKind = 249
	/*
	   OpenMP for SIMD directive.
	*/
	Cursor_OMPForSimdDirective CursorKind = 250
	/*
	   OpenMP parallel for SIMD directive.
	*/
	Cursor_OMPParallelForSimdDirective CursorKind = 251
	/*
	   OpenMP target directive.
	*/
	Cursor_OMPTargetDirective CursorKind = 252
	/*
	   OpenMP teams directive.
	*/
	Cursor_OMPTeamsDirective CursorKind = 253
	/*
	   OpenMP taskgroup directive.
	*/
	Cursor_OMPTaskgroupDirective CursorKind = 254
	/*
	   OpenMP cancellation point directive.
	*/
	Cursor_OMPCancellationPointDirective CursorKind = 255
	/*
	   OpenMP cancel directive.
	*/
	Cursor_OMPCancelDirective CursorKind = 256
	/*
	   OpenMP target data directive.
	*/
	Cursor_OMPTargetDataDirective CursorKind = 257
	/*
	   OpenMP taskloop directive.
	*/
	Cursor_OMPTaskLoopDirective CursorKind = 258
	/*
	   OpenMP taskloop simd directive.
	*/
	Cursor_OMPTaskLoopSimdDirective CursorKind = 259
	/*
	   OpenMP distribute directive.
	*/
	Cursor_OMPDistributeDirective CursorKind = 260
	/*
	   OpenMP target enter data directive.
	*/
	Cursor_OMPTargetEnterDataDirective CursorKind = 261
	/*
	   OpenMP target exit data directive.
	*/
	Cursor_OMPTargetExitDataDirective CursorKind = 262
	/*
	   OpenMP target parallel directive.
	*/
	Cursor_OMPTargetParallelDirective CursorKind = 263
	/*
	   OpenMP target parallel for directive.
	*/
	Cursor_OMPTargetParallelForDirective CursorKind = 264
	/*
	   OpenMP target update directive.
	*/
	Cursor_OMPTargetUpdateDirective CursorKind = 265
	/*
	   OpenMP distribute parallel for directive.
	*/
	Cursor_OMPDistributeParallelForDirective CursorKind = 266
	/*
	   OpenMP distribute parallel for simd directive.
	*/
	Cursor_OMPDistributeParallelForSimdDirective CursorKind = 267
	/*
	   OpenMP distribute simd directive.
	*/
	Cursor_OMPDistributeSimdDirective CursorKind = 268
	/*
	   OpenMP target parallel for simd directive.
	*/
	Cursor_OMPTargetParallelForSimdDirective CursorKind = 269
	/*
	   OpenMP target simd directive.
	*/
	Cursor_OMPTargetSimdDirective CursorKind = 270
	/*
	   OpenMP teams distribute directive.
	*/
	Cursor_OMPTeamsDistributeDirective CursorKind = 271
	/*
	   OpenMP teams distribute simd directive.
	*/
	Cursor_OMPTeamsDistributeSimdDirective CursorKind = 272
	/*
	   OpenMP teams distribute parallel for simd directive.
	*/
	Cursor_OMPTeamsDistributeParallelForSimdDirective CursorKind = 273
	/*
	   OpenMP teams distribute parallel for directive.
	*/
	Cursor_OMPTeamsDistributeParallelForDirective CursorKind = 274
	/*
	   OpenMP target teams directive.
	*/
	Cursor_OMPTargetTeamsDirective CursorKind = 275
	/*
	   OpenMP target teams distribute directive.
	*/
	Cursor_OMPTargetTeamsDistributeDirective CursorKind = 276
	/*
	   OpenMP target teams distribute parallel for directive.
	*/
	Cursor_OMPTargetTeamsDistributeParallelForDirective CursorKind = 277
	/*
	   OpenMP target teams distribute parallel for simd directive.
	*/
	Cursor_OMPTargetTeamsDistributeParallelForSimdDirective CursorKind = 278
	/*
	   OpenMP target teams distribute simd directive.
	*/
	Cursor_OMPTargetTeamsDistributeSimdDirective CursorKind = 279
	/*
	   C++2a std::bit_cast expression.
	*/
	Cursor_BuiltinBitCastExpr CursorKind = 280
	/*
	   OpenMP master taskloop directive.
	*/
	Cursor_OMPMasterTaskLoopDirective CursorKind = 281
	/*
	   OpenMP parallel master taskloop directive.
	*/
	Cursor_OMPParallelMasterTaskLoopDirective CursorKind = 282
	/*
	   OpenMP master taskloop simd directive.
	*/
	Cursor_OMPMasterTaskLoopSimdDirective CursorKind = 283
	/*
	   OpenMP parallel master taskloop simd directive.
	*/
	Cursor_OMPParallelMasterTaskLoopSimdDirective CursorKind = 284
	/*
	   OpenMP parallel master directive.
	*/
	Cursor_OMPParallelMasterDirective CursorKind = 285
	/*
	   OpenMP depobj directive.
	*/
	Cursor_OMPDepobjDirective CursorKind = 286
	/*
	   OpenMP scan directive.
	*/
	Cursor_OMPScanDirective CursorKind = 287
	/*
	   OpenMP tile directive.
	*/
	Cursor_OMPTileDirective CursorKind = 288
	/*
	   OpenMP canonical loop.
	*/
	Cursor_OMPCanonicalLoop CursorKind = 289
	/*
	   OpenMP interop directive.
	*/
	Cursor_OMPInteropDirective CursorKind = 290
	/*
	   OpenMP dispatch directive.
	*/
	Cursor_OMPDispatchDirective CursorKind = 291
	/*
	   OpenMP masked directive.
	*/
	Cursor_OMPMaskedDirective CursorKind = 292
	/*
	   OpenMP unroll directive.
	*/
	Cursor_OMPUnrollDirective CursorKind = 293
	/*
	   OpenMP metadirective directive.
	*/
	Cursor_OMPMetaDirective CursorKind = 294
	/*
	   OpenMP loop directive.
	*/
	Cursor_OMPGenericLoopDirective CursorKind = 295
	/*
	   OpenMP teams loop directive.
	*/
	Cursor_OMPTeamsGenericLoopDirective CursorKind = 296
	/*
	   OpenMP target teams loop directive.
	*/
	Cursor_OMPTargetTeamsGenericLoopDirective CursorKind = 297
	/*
	   OpenMP parallel loop directive.
	*/
	Cursor_OMPParallelGenericLoopDirective CursorKind = 298
	/*
	   OpenMP target parallel loop directive.
	*/
	Cursor_OMPTargetParallelGenericLoopDirective CursorKind = 299
	/*
	   OpenMP parallel masked directive.
	*/
	Cursor_OMPParallelMaskedDirective CursorKind = 300
	/*
	   OpenMP masked taskloop directive.
	*/
	Cursor_OMPMaskedTaskLoopDirective CursorKind = 301
	/*
	   OpenMP masked taskloop simd directive.
	*/
	Cursor_OMPMaskedTaskLoopSimdDirective CursorKind = 302
	/*
	   OpenMP parallel masked taskloop directive.
	*/
	Cursor_OMPParallelMaskedTaskLoopDirective CursorKind = 303
	/*
	   OpenMP parallel masked taskloop simd directive.
	*/
	Cursor_OMPParallelMaskedTaskLoopSimdDirective CursorKind = 304
	/*
	   OpenMP error directive.
	*/
	Cursor_OMPErrorDirective CursorKind = 305
	/*
	   OpenMP scope directive.
	*/
	Cursor_OMPScopeDirective CursorKind = 306
	/*
	   OpenMP reverse directive.
	*/
	Cursor_OMPReverseDirective CursorKind = 307
	/*
	   OpenMP interchange directive.
	*/
	Cursor_OMPInterchangeDirective CursorKind = 308
	/*
	   OpenMP assume directive.
	*/
	Cursor_OMPAssumeDirective CursorKind = 309
	/*
	   OpenMP assume directive.
	*/
	Cursor_OMPStripeDirective CursorKind = 310
	/*
	   OpenMP fuse directive
	*/
	Cursor_OMPFuseDirective CursorKind = 311
	/*
	   OpenMP split directive.
	*/
	Cursor_OMPSplitDirective CursorKind = 312
	/*
	   OpenACC Compute Construct.
	*/
	Cursor_OpenACCComputeConstruct CursorKind = 320
	/*
	   OpenACC Loop Construct.
	*/
	Cursor_OpenACCLoopConstruct CursorKind = 321
	/*
	   OpenACC Combined Constructs.
	*/
	Cursor_OpenACCCombinedConstruct CursorKind = 322
	/*
	   OpenACC data Construct.
	*/
	Cursor_OpenACCDataConstruct CursorKind = 323
	/*
	   OpenACC enter data Construct.
	*/
	Cursor_OpenACCEnterDataConstruct CursorKind = 324
	/*
	   OpenACC exit data Construct.
	*/
	Cursor_OpenACCExitDataConstruct CursorKind = 325
	/*
	   OpenACC host_data Construct.
	*/
	Cursor_OpenACCHostDataConstruct CursorKind = 326
	/*
	   OpenACC wait Construct.
	*/
	Cursor_OpenACCWaitConstruct CursorKind = 327
	/*
	   OpenACC init Construct.
	*/
	Cursor_OpenACCInitConstruct CursorKind = 328
	/*
	   OpenACC shutdown Construct.
	*/
	Cursor_OpenACCShutdownConstruct CursorKind = 329
	/*
	   OpenACC set Construct.
	*/
	Cursor_OpenACCSetConstruct CursorKind = 330
	/*
	   OpenACC update Construct.
	*/
	Cursor_OpenACCUpdateConstruct CursorKind = 331
	/*
	   OpenACC atomic Construct.
	*/
	Cursor_OpenACCAtomicConstruct CursorKind = 332
	/*
	   OpenACC cache Construct.
	*/
	Cursor_OpenACCCacheConstruct CursorKind = 333
	/*
	   OpenACC cache Construct.
	*/
	Cursor_LastStmt CursorKind = 333
	/*
	   Cursor that represents the translation unit itself.

	   The translation unit cursor exists primarily to act as the root cursor for traversing the contents of a translation unit.
	*/
	Cursor_TranslationUnit CursorKind = 350
	/*
	   Cursor that represents the translation unit itself.

	   The translation unit cursor exists primarily to act as the root cursor for traversing the contents of a translation unit.
	*/
	Cursor_FirstAttr CursorKind = 400
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_UnexposedAttr CursorKind = 400
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_IBActionAttr CursorKind = 401
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_IBOutletAttr CursorKind = 402
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_IBOutletCollectionAttr CursorKind = 403
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CXXFinalAttr CursorKind = 404
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CXXOverrideAttr CursorKind = 405
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_AnnotateAttr CursorKind = 406
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_AsmLabelAttr CursorKind = 407
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_PackedAttr CursorKind = 408
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_PureAttr CursorKind = 409
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ConstAttr CursorKind = 410
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_NoDuplicateAttr CursorKind = 411
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CUDAConstantAttr CursorKind = 412
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CUDADeviceAttr CursorKind = 413
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CUDAGlobalAttr CursorKind = 414
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CUDAHostAttr CursorKind = 415
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_CUDASharedAttr CursorKind = 416
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_VisibilityAttr CursorKind = 417
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_DLLExport CursorKind = 418
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_DLLImport CursorKind = 419
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_NSReturnsRetained CursorKind = 420
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_NSReturnsNotRetained CursorKind = 421
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_NSReturnsAutoreleased CursorKind = 422
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_NSConsumesSelf CursorKind = 423
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_NSConsumed CursorKind = 424
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCException CursorKind = 425
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCNSObject CursorKind = 426
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCIndependentClass CursorKind = 427
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCPreciseLifetime CursorKind = 428
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCReturnsInnerPointer CursorKind = 429
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCRequiresSuper CursorKind = 430
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCRootClass CursorKind = 431
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCSubclassingRestricted CursorKind = 432
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCExplicitProtocolImpl CursorKind = 433
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCDesignatedInitializer CursorKind = 434
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCRuntimeVisible CursorKind = 435
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ObjCBoxable CursorKind = 436
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_FlagEnum CursorKind = 437
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_ConvergentAttr CursorKind = 438
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_WarnUnusedAttr CursorKind = 439
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_WarnUnusedResultAttr CursorKind = 440
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_AlignedAttr CursorKind = 441
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_LastAttr CursorKind = 441
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_PreprocessingDirective CursorKind = 500
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_MacroDefinition CursorKind = 501
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_MacroExpansion CursorKind = 502
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_MacroInstantiation CursorKind = 502
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_InclusionDirective CursorKind = 503
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_FirstPreprocessing CursorKind = 500
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	Cursor_LastPreprocessing CursorKind = 503
	/*
	   A module import declaration.
	*/
	Cursor_ModuleImportDecl CursorKind = 600
	/*
	   A module import declaration.
	*/
	Cursor_TypeAliasTemplateDecl CursorKind = 601
	/*
	   A static_assert or _Static_assert node
	*/
	Cursor_StaticAssert CursorKind = 602
	/*
	   a friend declaration.
	*/
	Cursor_FriendDecl CursorKind = 603
	/*
	   a concept declaration.
	*/
	Cursor_ConceptDecl CursorKind = 604
	/*
	   a concept declaration.
	*/
	Cursor_FirstExtraDecl CursorKind = 600
	/*
	   a concept declaration.
	*/
	Cursor_LastExtraDecl CursorKind = 604
	/*
	   A code completion overload candidate.
	*/
	Cursor_OverloadCandidate CursorKind = 700
)

/*
Describe the linkage of the entity referred to by a cursor.
*/
type LinkageKind uint32

const (
	/*
	   This value indicates that no linkage information is available for a provided CXCursor.
	*/
	Linkage_Invalid LinkageKind = 0
	/*
	   This is the linkage for variables, parameters, and so on that  have automatic storage.  This covers normal (non-extern) local variables.
	*/
	Linkage_NoLinkage LinkageKind = 1
	/*
	   This is the linkage for static variables and static functions.
	*/
	Linkage_Internal LinkageKind = 2
	/*
	   This is the linkage for entities with external linkage that live in C++ anonymous namespaces.
	*/
	Linkage_UniqueExternal LinkageKind = 3
	/*
	   This is the linkage for entities with true, external linkage.
	*/
	Linkage_External LinkageKind = 4
)

type VisibilityKind uint32

const (
	/*
	   This value indicates that no visibility information is available for a provided CXCursor.
	*/
	Visibility_Invalid VisibilityKind = 0
	/*
	   Symbol not seen by the linker.
	*/
	Visibility_Hidden VisibilityKind = 1
	/*
	   Symbol seen by the linker but resolves to a symbol inside this object.
	*/
	Visibility_Protected VisibilityKind = 2
	/*
	   Symbol seen by the linker and acts like a normal symbol.
	*/
	Visibility_Default VisibilityKind = 3
)

/*
Describe the "language" of the entity referred to by a cursor.
*/
type LanguageKind uint32

const (
	Language_Invalid   LanguageKind = 0
	Language_C         LanguageKind = 1
	Language_ObjC      LanguageKind = 2
	Language_CPlusPlus LanguageKind = 3
)

/*
Describe the "thread-local storage (TLS) kind" of the declaration referred to by a cursor.
*/
type TLSKind uint32

const (
	TLS_None    TLSKind = 0
	TLS_Dynamic TLSKind = 1
	TLS_Static  TLSKind = 2
)

/*
Describes the kind of type
*/
type TypeKind uint32

const (
	/*
	   Represents an invalid type (e.g., where no type is available).
	*/
	Type_Invalid TypeKind = 0
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Unexposed TypeKind = 1
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Void TypeKind = 2
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Bool TypeKind = 3
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Char_U TypeKind = 4
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_UChar TypeKind = 5
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Char16 TypeKind = 6
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Char32 TypeKind = 7
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_UShort TypeKind = 8
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_UInt TypeKind = 9
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ULong TypeKind = 10
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ULongLong TypeKind = 11
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_UInt128 TypeKind = 12
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Char_S TypeKind = 13
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_SChar TypeKind = 14
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_WChar TypeKind = 15
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Short TypeKind = 16
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Int TypeKind = 17
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Long TypeKind = 18
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_LongLong TypeKind = 19
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Int128 TypeKind = 20
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Float TypeKind = 21
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Double TypeKind = 22
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_LongDouble TypeKind = 23
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_NullPtr TypeKind = 24
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Overload TypeKind = 25
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Dependent TypeKind = 26
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ObjCId TypeKind = 27
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ObjCClass TypeKind = 28
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ObjCSel TypeKind = 29
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Float128 TypeKind = 30
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Half TypeKind = 31
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Float16 TypeKind = 32
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ShortAccum TypeKind = 33
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Accum TypeKind = 34
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_LongAccum TypeKind = 35
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_UShortAccum TypeKind = 36
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_UAccum TypeKind = 37
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ULongAccum TypeKind = 38
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_BFloat16 TypeKind = 39
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Ibm128 TypeKind = 40
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_FirstBuiltin TypeKind = 2
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_LastBuiltin TypeKind = 40
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Complex TypeKind = 100
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Pointer TypeKind = 101
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_BlockPointer TypeKind = 102
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_LValueReference TypeKind = 103
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_RValueReference TypeKind = 104
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Record TypeKind = 105
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Enum TypeKind = 106
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Typedef TypeKind = 107
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ObjCInterface TypeKind = 108
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ObjCObjectPointer TypeKind = 109
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_FunctionNoProto TypeKind = 110
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_FunctionProto TypeKind = 111
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_ConstantArray TypeKind = 112
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Vector TypeKind = 113
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_IncompleteArray TypeKind = 114
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_VariableArray TypeKind = 115
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_DependentSizedArray TypeKind = 116
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_MemberPointer TypeKind = 117
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	Type_Auto TypeKind = 118
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Elaborated TypeKind = 119
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Pipe TypeKind = 120
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dRO TypeKind = 121
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dArrayRO TypeKind = 122
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dBufferRO TypeKind = 123
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dRO TypeKind = 124
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayRO TypeKind = 125
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dDepthRO TypeKind = 126
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayDepthRO TypeKind = 127
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAARO TypeKind = 128
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAARO TypeKind = 129
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAADepthRO TypeKind = 130
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAADepthRO TypeKind = 131
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage3dRO TypeKind = 132
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dWO TypeKind = 133
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dArrayWO TypeKind = 134
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dBufferWO TypeKind = 135
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dWO TypeKind = 136
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayWO TypeKind = 137
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dDepthWO TypeKind = 138
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayDepthWO TypeKind = 139
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAAWO TypeKind = 140
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAAWO TypeKind = 141
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAADepthWO TypeKind = 142
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAADepthWO TypeKind = 143
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage3dWO TypeKind = 144
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dRW TypeKind = 145
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dArrayRW TypeKind = 146
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dBufferRW TypeKind = 147
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dRW TypeKind = 148
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayRW TypeKind = 149
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dDepthRW TypeKind = 150
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayDepthRW TypeKind = 151
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAARW TypeKind = 152
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAARW TypeKind = 153
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAADepthRW TypeKind = 154
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAADepthRW TypeKind = 155
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage3dRW TypeKind = 156
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLSampler TypeKind = 157
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLEvent TypeKind = 158
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLQueue TypeKind = 159
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLReserveID TypeKind = 160
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_ObjCObject TypeKind = 161
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_ObjCTypeParam TypeKind = 162
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Attributed TypeKind = 163
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCMcePayload TypeKind = 164
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImePayload TypeKind = 165
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCRefPayload TypeKind = 166
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCSicPayload TypeKind = 167
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCMceResult TypeKind = 168
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResult TypeKind = 169
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCRefResult TypeKind = 170
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCSicResult TypeKind = 171
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResultSingleReferenceStreamout TypeKind = 172
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResultDualReferenceStreamout TypeKind = 173
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeSingleReferenceStreamin TypeKind = 174
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeDualReferenceStreamin TypeKind = 175
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResultSingleRefStreamout TypeKind = 172
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResultDualRefStreamout TypeKind = 173
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeSingleRefStreamin TypeKind = 174
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeDualRefStreamin TypeKind = 175
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_ExtVector TypeKind = 176
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Atomic TypeKind = 177
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_BTFTagAttributed TypeKind = 178
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_HLSLResource TypeKind = 179
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_HLSLAttributedResource TypeKind = 180
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_HLSLInlineSpirv TypeKind = 181
)

/*
Describes the calling convention of a function type
*/
type CallingConv uint32

const (
	CallingConv_Default            CallingConv = 0
	CallingConv_C                  CallingConv = 1
	CallingConv_X86StdCall         CallingConv = 2
	CallingConv_X86FastCall        CallingConv = 3
	CallingConv_X86ThisCall        CallingConv = 4
	CallingConv_X86Pascal          CallingConv = 5
	CallingConv_AAPCS              CallingConv = 6
	CallingConv_AAPCS_VFP          CallingConv = 7
	CallingConv_X86RegCall         CallingConv = 8
	CallingConv_IntelOclBicc       CallingConv = 9
	CallingConv_Win64              CallingConv = 10
	CallingConv_X86_64Win64        CallingConv = 10
	CallingConv_X86_64SysV         CallingConv = 11
	CallingConv_X86VectorCall      CallingConv = 12
	CallingConv_Swift              CallingConv = 13
	CallingConv_PreserveMost       CallingConv = 14
	CallingConv_PreserveAll        CallingConv = 15
	CallingConv_AArch64VectorCall  CallingConv = 16
	CallingConv_SwiftAsync         CallingConv = 17
	CallingConv_AArch64SVEPCS      CallingConv = 18
	CallingConv_M68kRTD            CallingConv = 19
	CallingConv_PreserveNone       CallingConv = 20
	CallingConv_RISCVVectorCall    CallingConv = 21
	CallingConv_RISCVVLSCall_32    CallingConv = 22
	CallingConv_RISCVVLSCall_64    CallingConv = 23
	CallingConv_RISCVVLSCall_128   CallingConv = 24
	CallingConv_RISCVVLSCall_256   CallingConv = 25
	CallingConv_RISCVVLSCall_512   CallingConv = 26
	CallingConv_RISCVVLSCall_1024  CallingConv = 27
	CallingConv_RISCVVLSCall_2048  CallingConv = 28
	CallingConv_RISCVVLSCall_4096  CallingConv = 29
	CallingConv_RISCVVLSCall_8192  CallingConv = 30
	CallingConv_RISCVVLSCall_16384 CallingConv = 31
	CallingConv_RISCVVLSCall_32768 CallingConv = 32
	CallingConv_RISCVVLSCall_65536 CallingConv = 33
	CallingConv_Invalid            CallingConv = 100
	CallingConv_Unexposed          CallingConv = 200
)

/*
Describes the kind of a template argument.

See the definition of llvm::clang::TemplateArgument::ArgKind for full element descriptions.
*/
type TemplateArgumentKind uint32

const (
	TemplateArgumentKind_Null              TemplateArgumentKind = 0
	TemplateArgumentKind_Type              TemplateArgumentKind = 1
	TemplateArgumentKind_Declaration       TemplateArgumentKind = 2
	TemplateArgumentKind_NullPtr           TemplateArgumentKind = 3
	TemplateArgumentKind_Integral          TemplateArgumentKind = 4
	TemplateArgumentKind_Template          TemplateArgumentKind = 5
	TemplateArgumentKind_TemplateExpansion TemplateArgumentKind = 6
	TemplateArgumentKind_Expression        TemplateArgumentKind = 7
	TemplateArgumentKind_Pack              TemplateArgumentKind = 8
	TemplateArgumentKind_Invalid           TemplateArgumentKind = 9
)

type TypeNullabilityKind uint32

const (
	/*
	   Values of this type can never be null.
	*/
	TypeNullability_NonNull TypeNullabilityKind = 0
	/*
	   Values of this type can be null.
	*/
	TypeNullability_Nullable TypeNullabilityKind = 1
	/*
	   Whether values of this type can be null is (explicitly) unspecified. This captures a (fairly rare) case where we can't conclude anything about the nullability of the type even though it has been considered.
	*/
	TypeNullability_Unspecified TypeNullabilityKind = 2
	/*
	   Nullability is not applicable to this type.
	*/
	TypeNullability_Invalid TypeNullabilityKind = 3
	/*
	   Generally behaves like Nullable, except when used in a block parameter that was imported into a swift async method. There, swift will assume that the parameter can get null even if no error occurred. _Nullable parameters are assumed to only get null on error.
	*/
	TypeNullability_NullableResult TypeNullabilityKind = 4
)

/*
List the possible error codes for clang_Type_getSizeOf,   clang_Type_getAlignOf, clang_Type_getOffsetOf,   clang_Cursor_getOffsetOf, and clang_getOffsetOfBase.

A value of this enumeration type can be returned if the target type is not a valid argument to sizeof, alignof or offsetof.
*/
type TypeLayoutError int32

const (
	/*
	   Type is of kind CXType_Invalid.
	*/
	TypeLayoutError_Invalid TypeLayoutError = -1
	/*
	   The type is an incomplete Type.
	*/
	TypeLayoutError_Incomplete TypeLayoutError = -2
	/*
	   The type is a dependent Type.
	*/
	TypeLayoutError_Dependent TypeLayoutError = -3
	/*
	   The type is not a constant size type.
	*/
	TypeLayoutError_NotConstantSize TypeLayoutError = -4
	/*
	   The Field name is not valid for this record.
	*/
	TypeLayoutError_InvalidFieldName TypeLayoutError = -5
	/*
	   The type is undeduced.
	*/
	TypeLayoutError_Undeduced TypeLayoutError = -6
)

type RefQualifierKind uint32

const (
	/*
	   No ref-qualifier was provided.
	*/
	RefQualifier_None RefQualifierKind = 0
	/*
	   An lvalue ref-qualifier was provided (&).
	*/
	RefQualifier_LValue RefQualifierKind = 1
	/*
	   An rvalue ref-qualifier was provided (&&).
	*/
	RefQualifier_RValue RefQualifierKind = 2
)

/*
Represents the C++ access control level to a base class for a cursor with kind CX_CXXBaseSpecifier.
*/
type CXXAccessSpecifier uint32

const (
	CXXInvalidAccessSpecifier CXXAccessSpecifier = 0
	CXXPublic                 CXXAccessSpecifier = 1
	CXXProtected              CXXAccessSpecifier = 2
	CXXPrivate                CXXAccessSpecifier = 3
)

/*
Represents the storage classes as declared in the source. CX_SC_Invalid was added for the case that the passed cursor in not a declaration.
*/
type StorageClass uint32

const (
	SC_Invalid              StorageClass = 0
	SC_None                 StorageClass = 1
	SC_Extern               StorageClass = 2
	SC_Static               StorageClass = 3
	SC_PrivateExtern        StorageClass = 4
	SC_OpenCLWorkGroupLocal StorageClass = 5
	SC_Auto                 StorageClass = 6
	SC_Register             StorageClass = 7
)

/*
Represents a specific kind of binary operator which can appear at a cursor.
*/
type BinaryOperatorKind_ uint32

const (
	BO_Invalid   BinaryOperatorKind_ = 0
	BO_PtrMemD   BinaryOperatorKind_ = 1
	BO_PtrMemI   BinaryOperatorKind_ = 2
	BO_Mul       BinaryOperatorKind_ = 3
	BO_Div       BinaryOperatorKind_ = 4
	BO_Rem       BinaryOperatorKind_ = 5
	BO_Add       BinaryOperatorKind_ = 6
	BO_Sub       BinaryOperatorKind_ = 7
	BO_Shl       BinaryOperatorKind_ = 8
	BO_Shr       BinaryOperatorKind_ = 9
	BO_Cmp       BinaryOperatorKind_ = 10
	BO_LT        BinaryOperatorKind_ = 11
	BO_GT        BinaryOperatorKind_ = 12
	BO_LE        BinaryOperatorKind_ = 13
	BO_GE        BinaryOperatorKind_ = 14
	BO_EQ        BinaryOperatorKind_ = 15
	BO_NE        BinaryOperatorKind_ = 16
	BO_And       BinaryOperatorKind_ = 17
	BO_Xor       BinaryOperatorKind_ = 18
	BO_Or        BinaryOperatorKind_ = 19
	BO_LAnd      BinaryOperatorKind_ = 20
	BO_LOr       BinaryOperatorKind_ = 21
	BO_Assign    BinaryOperatorKind_ = 22
	BO_MulAssign BinaryOperatorKind_ = 23
	BO_DivAssign BinaryOperatorKind_ = 24
	BO_RemAssign BinaryOperatorKind_ = 25
	BO_AddAssign BinaryOperatorKind_ = 26
	BO_SubAssign BinaryOperatorKind_ = 27
	BO_ShlAssign BinaryOperatorKind_ = 28
	BO_ShrAssign BinaryOperatorKind_ = 29
	BO_AndAssign BinaryOperatorKind_ = 30
	BO_XorAssign BinaryOperatorKind_ = 31
	BO_OrAssign  BinaryOperatorKind_ = 32
	BO_Comma     BinaryOperatorKind_ = 33
	BO_LAST      BinaryOperatorKind_ = 33
)

/*
Describes how the traversal of the children of a particular cursor should proceed after visiting a particular child cursor.

A value of this enumeration type should be returned by each CXCursorVisitor to indicate how clang_visitChildren() proceed.
*/
type ChildVisitResult uint32

const (
	/*
	   Terminates the cursor traversal.
	*/
	ChildVisit_Break ChildVisitResult = 0
	/*
	   Continues the cursor traversal with the next sibling of the cursor just visited, without visiting its children.
	*/
	ChildVisit_Continue ChildVisitResult = 1
	/*
	   Recursively traverse the children of this cursor, using the same visitor and client data.
	*/
	ChildVisit_Recurse ChildVisitResult = 2
)

/*
Properties for the printing policy.

See clang::PrintingPolicy for more information.
*/
type PrintingPolicyProperty uint32

const (
	PrintingPolicy_Indentation                           PrintingPolicyProperty = 0
	PrintingPolicy_SuppressSpecifiers                    PrintingPolicyProperty = 1
	PrintingPolicy_SuppressTagKeyword                    PrintingPolicyProperty = 2
	PrintingPolicy_IncludeTagDefinition                  PrintingPolicyProperty = 3
	PrintingPolicy_SuppressScope                         PrintingPolicyProperty = 4
	PrintingPolicy_SuppressUnwrittenScope                PrintingPolicyProperty = 5
	PrintingPolicy_SuppressInitializers                  PrintingPolicyProperty = 6
	PrintingPolicy_ConstantArraySizeAsWritten            PrintingPolicyProperty = 7
	PrintingPolicy_AnonymousTagLocations                 PrintingPolicyProperty = 8
	PrintingPolicy_SuppressStrongLifetime                PrintingPolicyProperty = 9
	PrintingPolicy_SuppressLifetimeQualifiers            PrintingPolicyProperty = 10
	PrintingPolicy_SuppressTemplateArgsInCXXConstructors PrintingPolicyProperty = 11
	PrintingPolicy_Bool                                  PrintingPolicyProperty = 12
	PrintingPolicy_Restrict                              PrintingPolicyProperty = 13
	PrintingPolicy_Alignof                               PrintingPolicyProperty = 14
	PrintingPolicy_UnderscoreAlignof                     PrintingPolicyProperty = 15
	PrintingPolicy_UseVoidForZeroParams                  PrintingPolicyProperty = 16
	PrintingPolicy_TerseOutput                           PrintingPolicyProperty = 17
	PrintingPolicy_PolishForDeclaration                  PrintingPolicyProperty = 18
	PrintingPolicy_Half                                  PrintingPolicyProperty = 19
	PrintingPolicy_MSWChar                               PrintingPolicyProperty = 20
	PrintingPolicy_IncludeNewlines                       PrintingPolicyProperty = 21
	PrintingPolicy_MSVCFormatting                        PrintingPolicyProperty = 22
	PrintingPolicy_ConstantsAsWritten                    PrintingPolicyProperty = 23
	PrintingPolicy_SuppressImplicitBase                  PrintingPolicyProperty = 24
	PrintingPolicy_FullyQualifiedName                    PrintingPolicyProperty = 25
	PrintingPolicy_LastProperty                          PrintingPolicyProperty = 25
)

/*
Property attributes for a CXCursor_ObjCPropertyDecl.
*/
type ObjCPropertyAttrKind uint32

const (
	ObjCPropertyAttr_noattr            ObjCPropertyAttrKind = 0
	ObjCPropertyAttr_readonly          ObjCPropertyAttrKind = 1
	ObjCPropertyAttr_getter            ObjCPropertyAttrKind = 2
	ObjCPropertyAttr_assign            ObjCPropertyAttrKind = 4
	ObjCPropertyAttr_readwrite         ObjCPropertyAttrKind = 8
	ObjCPropertyAttr_retain            ObjCPropertyAttrKind = 16
	ObjCPropertyAttr_copy              ObjCPropertyAttrKind = 32
	ObjCPropertyAttr_nonatomic         ObjCPropertyAttrKind = 64
	ObjCPropertyAttr_setter            ObjCPropertyAttrKind = 128
	ObjCPropertyAttr_atomic            ObjCPropertyAttrKind = 256
	ObjCPropertyAttr_weak              ObjCPropertyAttrKind = 512
	ObjCPropertyAttr_strong            ObjCPropertyAttrKind = 1024
	ObjCPropertyAttr_unsafe_unretained ObjCPropertyAttrKind = 2048
	ObjCPropertyAttr_class             ObjCPropertyAttrKind = 4096
)

/*
'Qualifiers' written next to the return and parameter types in Objective-C method declarations.
*/
type ObjCDeclQualifierKind uint32

const (
	ObjCDeclQualifier_None   ObjCDeclQualifierKind = 0
	ObjCDeclQualifier_In     ObjCDeclQualifierKind = 1
	ObjCDeclQualifier_Inout  ObjCDeclQualifierKind = 2
	ObjCDeclQualifier_Out    ObjCDeclQualifierKind = 4
	ObjCDeclQualifier_Bycopy ObjCDeclQualifierKind = 8
	ObjCDeclQualifier_Byref  ObjCDeclQualifierKind = 16
	ObjCDeclQualifier_Oneway ObjCDeclQualifierKind = 32
)

type NameRefFlags uint32

const (
	/*
	   Include the nested-name-specifier, e.g. Foo:: in x.Foo::y, in the range.
	*/
	NameRange_WantQualifier NameRefFlags = 1
	/*
	   Include the explicit template arguments, e.g. <int> in x.f<int>, in the range.
	*/
	NameRange_WantTemplateArgs NameRefFlags = 2
	/*
	   If the name is non-contiguous, return the full spanning range.

	   Non-contiguous names occur in Objective-C when a selector with two or more parameters is used, or in C++ when using an operator:
	*/
	NameRange_WantSinglePiece NameRefFlags = 4
)

/*
Describes a kind of token.
*/
type TokenKind uint32

const (
	/*
	   A token that contains some kind of punctuation.
	*/
	Token_Punctuation TokenKind = 0
	/*
	   A language keyword.
	*/
	Token_Keyword TokenKind = 1
	/*
	   An identifier (that is not a keyword).
	*/
	Token_Identifier TokenKind = 2
	/*
	   A numeric, string, or character literal.
	*/
	Token_Literal TokenKind = 3
	/*
	   A comment.
	*/
	Token_Comment TokenKind = 4
)

/*
Describes a single piece of text within a code-completion string.

Each "chunk" within a code-completion string (CXCompletionString) is either a piece of text with a specific "kind" that describes how that text should be interpreted by the client or is another completion string.
*/
type CompletionChunkKind uint32

const (
	/*
	   A code-completion string that describes "optional" text that could be a part of the template (but is not required).

	   The Optional chunk is the only kind of chunk that has a code-completion string for its representation, which is accessible via clang_getCompletionChunkCompletionString(). The code-completion string describes an additional part of the template that is completely optional. For example, optional chunks can be used to describe the placeholders for arguments that match up with defaulted function parameters, e.g. given:

	   The code-completion string for this function would contain:   - a TypedText chunk for "f".   - a LeftParen chunk for "(".   - a Placeholder chunk for "int x"   - an Optional chunk containing the remaining defaulted arguments, e.g.,       - a Comma chunk for ","       - a Placeholder chunk for "float y"       - an Optional chunk containing the last defaulted argument:           - a Comma chunk for ","           - a Placeholder chunk for "double z"   - a RightParen chunk for ")"

	   There are many ways to handle Optional chunks. Two simple approaches are:   - Completely ignore optional chunks, in which case the template for the     function "f" would only include the first parameter ("int x").   - Fully expand all optional chunks, in which case the template for the     function "f" would have all of the parameters.
	*/
	CompletionChunk_Optional CompletionChunkKind = 0
	/*
	   Text that a user would be expected to type to get this code-completion result.

	   There will be exactly one "typed text" chunk in a semantic string, which will typically provide the spelling of a keyword or the name of a declaration that could be used at the current code point. Clients are expected to filter the code-completion results based on the text in this chunk.
	*/
	CompletionChunk_TypedText CompletionChunkKind = 1
	/*
	   Text that should be inserted as part of a code-completion result.

	   A "text" chunk represents text that is part of the template to be inserted into user code should this particular code-completion result be selected.
	*/
	CompletionChunk_Text CompletionChunkKind = 2
	/*
	   Placeholder text that should be replaced by the user.

	   A "placeholder" chunk marks a place where the user should insert text into the code-completion template. For example, placeholders might mark the function parameters for a function declaration, to indicate that the user should provide arguments for each of those parameters. The actual text in a placeholder is a suggestion for the text to display before the user replaces the placeholder with real code.
	*/
	CompletionChunk_Placeholder CompletionChunkKind = 3
	/*
	   Informative text that should be displayed but never inserted as part of the template.

	   An "informative" chunk contains annotations that can be displayed to help the user decide whether a particular code-completion result is the right option, but which is not part of the actual template to be inserted by code completion.
	*/
	CompletionChunk_Informative CompletionChunkKind = 4
	/*
	   Text that describes the current parameter when code-completion is referring to function call, message send, or template specialization.

	   A "current parameter" chunk occurs when code-completion is providing information about a parameter corresponding to the argument at the code-completion point. For example, given a function

	   and the source code add(, where the code-completion point is after the "(", the code-completion string will contain a "current parameter" chunk for "int x", indicating that the current argument will initialize that parameter. After typing further, to add(17, (where the code-completion point is after the ","), the code-completion string will contain a "current parameter" chunk to "int y".
	*/
	CompletionChunk_CurrentParameter CompletionChunkKind = 5
	/*
	   A left parenthesis ('('), used to initiate a function call or signal the beginning of a function parameter list.
	*/
	CompletionChunk_LeftParen CompletionChunkKind = 6
	/*
	   A right parenthesis (')'), used to finish a function call or signal the end of a function parameter list.
	*/
	CompletionChunk_RightParen CompletionChunkKind = 7
	/*
	   A left bracket ('[').
	*/
	CompletionChunk_LeftBracket CompletionChunkKind = 8
	/*
	   A right bracket (']').
	*/
	CompletionChunk_RightBracket CompletionChunkKind = 9
	/*
	   A left brace ('{').
	*/
	CompletionChunk_LeftBrace CompletionChunkKind = 10
	/*
	   A right brace ('}').
	*/
	CompletionChunk_RightBrace CompletionChunkKind = 11
	/*
	   A left angle bracket ('<').
	*/
	CompletionChunk_LeftAngle CompletionChunkKind = 12
	/*
	   A right angle bracket ('>').
	*/
	CompletionChunk_RightAngle CompletionChunkKind = 13
	/*
	   A comma separator (',').
	*/
	CompletionChunk_Comma CompletionChunkKind = 14
	/*
	   Text that specifies the result type of a given result.

	   This special kind of informative chunk is not meant to be inserted into the text buffer. Rather, it is meant to illustrate the type that an expression using the given completion string would have.
	*/
	CompletionChunk_ResultType CompletionChunkKind = 15
	/*
	   A colon (':').
	*/
	CompletionChunk_Colon CompletionChunkKind = 16
	/*
	   A semicolon (';').
	*/
	CompletionChunk_SemiColon CompletionChunkKind = 17
	/*
	   An '=' sign.
	*/
	CompletionChunk_Equal CompletionChunkKind = 18
	/*
	   Horizontal space (' ').
	*/
	CompletionChunk_HorizontalSpace CompletionChunkKind = 19
	/*
	   Vertical space ('\n'), after which it is generally a good idea to perform indentation.
	*/
	CompletionChunk_VerticalSpace CompletionChunkKind = 20
)

/*
Flags that can be passed to clang_codeCompleteAt() to modify its behavior.

The enumerators in this enumeration can be bitwise-OR'd together to provide multiple options to clang_codeCompleteAt().
*/
type CodeComplete_Flags uint32

const (
	/*
	   Whether to include macros within the set of code completions returned.
	*/
	CodeComplete_IncludeMacros CodeComplete_Flags = 1
	/*
	   Whether to include code patterns for language constructs within the set of code completions, e.g., for loops.
	*/
	CodeComplete_IncludeCodePatterns CodeComplete_Flags = 2
	/*
	   Whether to include brief documentation within the set of code completions returned.
	*/
	CodeComplete_IncludeBriefComments CodeComplete_Flags = 4
	/*
	   Whether to speed up completion by omitting top- or namespace-level entities defined in the preamble. There's no guarantee any particular entity is omitted. This may be useful if the headers are indexed externally.
	*/
	CodeComplete_SkipPreamble CodeComplete_Flags = 8
	/*
	   Whether to include completions with small fix-its, e.g. change '.' to '->' on member access, etc.
	*/
	CodeComplete_IncludeCompletionsWithFixIts CodeComplete_Flags = 16
)

/*
Bits that represent the context under which completion is occurring.

The enumerators in this enumeration may be bitwise-OR'd together if multiple contexts are occurring simultaneously.
*/
type CompletionContext uint32

const (
	/*
	   The context for completions is unexposed, as only Clang results should be included. (This is equivalent to having no context bits set.)
	*/
	CompletionContext_Unexposed CompletionContext = 0
	/*
	   Completions for any possible type should be included in the results.
	*/
	CompletionContext_AnyType CompletionContext = 1
	/*
	   Completions for any possible value (variables, function calls, etc.) should be included in the results.
	*/
	CompletionContext_AnyValue CompletionContext = 2
	/*
	   Completions for values that resolve to an Objective-C object should be included in the results.
	*/
	CompletionContext_ObjCObjectValue CompletionContext = 4
	/*
	   Completions for values that resolve to an Objective-C selector should be included in the results.
	*/
	CompletionContext_ObjCSelectorValue CompletionContext = 8
	/*
	   Completions for values that resolve to a C++ class type should be included in the results.
	*/
	CompletionContext_CXXClassTypeValue CompletionContext = 16
	/*
	   Completions for fields of the member being accessed using the dot operator should be included in the results.
	*/
	CompletionContext_DotMemberAccess CompletionContext = 32
	/*
	   Completions for fields of the member being accessed using the arrow operator should be included in the results.
	*/
	CompletionContext_ArrowMemberAccess CompletionContext = 64
	/*
	   Completions for properties of the Objective-C object being accessed using the dot operator should be included in the results.
	*/
	CompletionContext_ObjCPropertyAccess CompletionContext = 128
	/*
	   Completions for enum tags should be included in the results.
	*/
	CompletionContext_EnumTag CompletionContext = 256
	/*
	   Completions for union tags should be included in the results.
	*/
	CompletionContext_UnionTag CompletionContext = 512
	/*
	   Completions for struct tags should be included in the results.
	*/
	CompletionContext_StructTag CompletionContext = 1024
	/*
	   Completions for C++ class names should be included in the results.
	*/
	CompletionContext_ClassTag CompletionContext = 2048
	/*
	   Completions for C++ namespaces and namespace aliases should be included in the results.
	*/
	CompletionContext_Namespace CompletionContext = 4096
	/*
	   Completions for C++ nested name specifiers should be included in the results.
	*/
	CompletionContext_NestedNameSpecifier CompletionContext = 8192
	/*
	   Completions for Objective-C interfaces (classes) should be included in the results.
	*/
	CompletionContext_ObjCInterface CompletionContext = 16384
	/*
	   Completions for Objective-C protocols should be included in the results.
	*/
	CompletionContext_ObjCProtocol CompletionContext = 32768
	/*
	   Completions for Objective-C categories should be included in the results.
	*/
	CompletionContext_ObjCCategory CompletionContext = 65536
	/*
	   Completions for Objective-C instance messages should be included in the results.
	*/
	CompletionContext_ObjCInstanceMessage CompletionContext = 131072
	/*
	   Completions for Objective-C class messages should be included in the results.
	*/
	CompletionContext_ObjCClassMessage CompletionContext = 262144
	/*
	   Completions for Objective-C selector names should be included in the results.
	*/
	CompletionContext_ObjCSelectorName CompletionContext = 524288
	/*
	   Completions for preprocessor macro names should be included in the results.
	*/
	CompletionContext_MacroName CompletionContext = 1048576
	/*
	   Natural language completions should be included in the results.
	*/
	CompletionContext_NaturalLanguage CompletionContext = 2097152
	/*
	   #include file completions should be included in the results.
	*/
	CompletionContext_IncludedFile CompletionContext = 4194304
	/*
	   The current context is unknown, so set all contexts.
	*/
	CompletionContext_Unknown CompletionContext = 8388607
)

type EvalResultKind uint32

const (
	Eval_Int            EvalResultKind = 1
	Eval_Float          EvalResultKind = 2
	Eval_ObjCStrLiteral EvalResultKind = 3
	Eval_StrLiteral     EvalResultKind = 4
	Eval_CFStr          EvalResultKind = 5
	Eval_Other          EvalResultKind = 6
	Eval_UnExposed      EvalResultKind = 0
)

/*
@{
*/
type VisitorResult uint32

const (
	Visit_Break    VisitorResult = 0
	Visit_Continue VisitorResult = 1
)

type Result uint32

const (
	/*
	   Function returned successfully.
	*/
	Result_Success Result = 0
	/*
	   One of the parameters was invalid for the function.
	*/
	Result_Invalid Result = 1
	/*
	   The function was terminated by a callback (e.g. it returned CXVisit_Break)
	*/
	Result_VisitBreak Result = 2
)

type IdxEntityKind uint32

const (
	IdxEntity_Unexposed             IdxEntityKind = 0
	IdxEntity_Typedef               IdxEntityKind = 1
	IdxEntity_Function              IdxEntityKind = 2
	IdxEntity_Variable              IdxEntityKind = 3
	IdxEntity_Field                 IdxEntityKind = 4
	IdxEntity_EnumConstant          IdxEntityKind = 5
	IdxEntity_ObjCClass             IdxEntityKind = 6
	IdxEntity_ObjCProtocol          IdxEntityKind = 7
	IdxEntity_ObjCCategory          IdxEntityKind = 8
	IdxEntity_ObjCInstanceMethod    IdxEntityKind = 9
	IdxEntity_ObjCClassMethod       IdxEntityKind = 10
	IdxEntity_ObjCProperty          IdxEntityKind = 11
	IdxEntity_ObjCIvar              IdxEntityKind = 12
	IdxEntity_Enum                  IdxEntityKind = 13
	IdxEntity_Struct                IdxEntityKind = 14
	IdxEntity_Union                 IdxEntityKind = 15
	IdxEntity_CXXClass              IdxEntityKind = 16
	IdxEntity_CXXNamespace          IdxEntityKind = 17
	IdxEntity_CXXNamespaceAlias     IdxEntityKind = 18
	IdxEntity_CXXStaticVariable     IdxEntityKind = 19
	IdxEntity_CXXStaticMethod       IdxEntityKind = 20
	IdxEntity_CXXInstanceMethod     IdxEntityKind = 21
	IdxEntity_CXXConstructor        IdxEntityKind = 22
	IdxEntity_CXXDestructor         IdxEntityKind = 23
	IdxEntity_CXXConversionFunction IdxEntityKind = 24
	IdxEntity_CXXTypeAlias          IdxEntityKind = 25
	IdxEntity_CXXInterface          IdxEntityKind = 26
	IdxEntity_CXXConcept            IdxEntityKind = 27
)

type IdxEntityLanguage uint32

const (
	IdxEntityLang_None  IdxEntityLanguage = 0
	IdxEntityLang_C     IdxEntityLanguage = 1
	IdxEntityLang_ObjC  IdxEntityLanguage = 2
	IdxEntityLang_CXX   IdxEntityLanguage = 3
	IdxEntityLang_Swift IdxEntityLanguage = 4
)

/*
Extra C++ template information for an entity. This can apply to: CXIdxEntity_Function CXIdxEntity_CXXClass CXIdxEntity_CXXStaticMethod CXIdxEntity_CXXInstanceMethod CXIdxEntity_CXXConstructor CXIdxEntity_CXXConversionFunction CXIdxEntity_CXXTypeAlias
*/
type IdxEntityCXXTemplateKind uint32

const (
	IdxEntity_NonTemplate                   IdxEntityCXXTemplateKind = 0
	IdxEntity_Template                      IdxEntityCXXTemplateKind = 1
	IdxEntity_TemplatePartialSpecialization IdxEntityCXXTemplateKind = 2
	IdxEntity_TemplateSpecialization        IdxEntityCXXTemplateKind = 3
)

type IdxAttrKind uint32

const (
	IdxAttr_Unexposed          IdxAttrKind = 0
	IdxAttr_IBAction           IdxAttrKind = 1
	IdxAttr_IBOutlet           IdxAttrKind = 2
	IdxAttr_IBOutletCollection IdxAttrKind = 3
)

type IdxDeclInfoFlags uint32

const (
	IdxDeclFlag_Skipped IdxDeclInfoFlags = 1
)

type IdxObjCContainerKind uint32

const (
	IdxObjCContainer_ForwardRef     IdxObjCContainerKind = 0
	IdxObjCContainer_Interface      IdxObjCContainerKind = 1
	IdxObjCContainer_Implementation IdxObjCContainerKind = 2
)

/*
Data for IndexerCallbacks#indexEntityReference.

This may be deprecated in a future version as this duplicates the CXSymbolRole_Implicit bit in CXSymbolRole.
*/
type IdxEntityRefKind uint32

const (
	/*
	   The entity is referenced directly in user's code.
	*/
	IdxEntityRef_Direct IdxEntityRefKind = 1
	/*
	   An implicit reference, e.g. a reference of an Objective-C method via the dot syntax.
	*/
	IdxEntityRef_Implicit IdxEntityRefKind = 2
)

/*
Roles that are attributed to symbol occurrences.

Internal: this currently mirrors low 9 bits of clang::index::SymbolRole with higher bits zeroed. These high bits may be exposed in the future.
*/
type SymbolRole uint32

const (
	SymbolRole_None        SymbolRole = 0
	SymbolRole_Declaration SymbolRole = 1
	SymbolRole_Definition  SymbolRole = 2
	SymbolRole_Reference   SymbolRole = 4
	SymbolRole_Read        SymbolRole = 8
	SymbolRole_Write       SymbolRole = 16
	SymbolRole_Call        SymbolRole = 32
	SymbolRole_Dynamic     SymbolRole = 64
	SymbolRole_AddressOf   SymbolRole = 128
	SymbolRole_Implicit    SymbolRole = 256
)

type IndexOptFlags uint32

const (
	/*
	   Used to indicate that no special indexing options are needed.
	*/
	IndexOpt_None IndexOptFlags = 0
	/*
	   Used to indicate that IndexerCallbacks#indexEntityReference should be invoked for only one reference of an entity per source file that does not also include a declaration/definition of the entity.
	*/
	IndexOpt_SuppressRedundantRefs IndexOptFlags = 1
	/*
	   Function-local symbols should be indexed. If this is not set function-local symbols will be ignored.
	*/
	IndexOpt_IndexFunctionLocalSymbols IndexOptFlags = 2
	/*
	   Implicit function/class template instantiations should be indexed. If this is not set, implicit instantiations will be ignored.
	*/
	IndexOpt_IndexImplicitTemplateInstantiations IndexOptFlags = 4
	/*
	   Suppress all compiler warnings when parsing for indexing.
	*/
	IndexOpt_SuppressWarnings IndexOptFlags = 8
	/*
	   Skip a function/method body that was already parsed during an indexing session associated with a CXIndexAction object. Bodies in system headers are always skipped.
	*/
	IndexOpt_SkipParsedBodiesInSession IndexOptFlags = 16
)

/*
Describes the kind of binary operators.
*/
type BinaryOperatorKind uint32

const (
	/*
	   This value describes cursors which are not binary operators.
	*/
	BinaryOperator_Invalid BinaryOperatorKind = 0
	/*
	   C++ Pointer - to - member operator.
	*/
	BinaryOperator_PtrMemD BinaryOperatorKind = 1
	/*
	   C++ Pointer - to - member operator.
	*/
	BinaryOperator_PtrMemI BinaryOperatorKind = 2
	/*
	   Multiplication operator.
	*/
	BinaryOperator_Mul BinaryOperatorKind = 3
	/*
	   Division operator.
	*/
	BinaryOperator_Div BinaryOperatorKind = 4
	/*
	   Remainder operator.
	*/
	BinaryOperator_Rem BinaryOperatorKind = 5
	/*
	   Addition operator.
	*/
	BinaryOperator_Add BinaryOperatorKind = 6
	/*
	   Subtraction operator.
	*/
	BinaryOperator_Sub BinaryOperatorKind = 7
	/*
	   Bitwise shift left operator.
	*/
	BinaryOperator_Shl BinaryOperatorKind = 8
	/*
	   Bitwise shift right operator.
	*/
	BinaryOperator_Shr BinaryOperatorKind = 9
	/*
	   C++ three-way comparison (spaceship) operator.
	*/
	BinaryOperator_Cmp BinaryOperatorKind = 10
	/*
	   Less than operator.
	*/
	BinaryOperator_LT BinaryOperatorKind = 11
	/*
	   Greater than operator.
	*/
	BinaryOperator_GT BinaryOperatorKind = 12
	/*
	   Less or equal operator.
	*/
	BinaryOperator_LE BinaryOperatorKind = 13
	/*
	   Greater or equal operator.
	*/
	BinaryOperator_GE BinaryOperatorKind = 14
	/*
	   Equal operator.
	*/
	BinaryOperator_EQ BinaryOperatorKind = 15
	/*
	   Not equal operator.
	*/
	BinaryOperator_NE BinaryOperatorKind = 16
	/*
	   Bitwise AND operator.
	*/
	BinaryOperator_And BinaryOperatorKind = 17
	/*
	   Bitwise XOR operator.
	*/
	BinaryOperator_Xor BinaryOperatorKind = 18
	/*
	   Bitwise OR operator.
	*/
	BinaryOperator_Or BinaryOperatorKind = 19
	/*
	   Logical AND operator.
	*/
	BinaryOperator_LAnd BinaryOperatorKind = 20
	/*
	   Logical OR operator.
	*/
	BinaryOperator_LOr BinaryOperatorKind = 21
	/*
	   Assignment operator.
	*/
	BinaryOperator_Assign BinaryOperatorKind = 22
	/*
	   Multiplication assignment operator.
	*/
	BinaryOperator_MulAssign BinaryOperatorKind = 23
	/*
	   Division assignment operator.
	*/
	BinaryOperator_DivAssign BinaryOperatorKind = 24
	/*
	   Remainder assignment operator.
	*/
	BinaryOperator_RemAssign BinaryOperatorKind = 25
	/*
	   Addition assignment operator.
	*/
	BinaryOperator_AddAssign BinaryOperatorKind = 26
	/*
	   Subtraction assignment operator.
	*/
	BinaryOperator_SubAssign BinaryOperatorKind = 27
	/*
	   Bitwise shift left assignment operator.
	*/
	BinaryOperator_ShlAssign BinaryOperatorKind = 28
	/*
	   Bitwise shift right assignment operator.
	*/
	BinaryOperator_ShrAssign BinaryOperatorKind = 29
	/*
	   Bitwise AND assignment operator.
	*/
	BinaryOperator_AndAssign BinaryOperatorKind = 30
	/*
	   Bitwise XOR assignment operator.
	*/
	BinaryOperator_XorAssign BinaryOperatorKind = 31
	/*
	   Bitwise OR assignment operator.
	*/
	BinaryOperator_OrAssign BinaryOperatorKind = 32
	/*
	   Comma operator.
	*/
	BinaryOperator_Comma BinaryOperatorKind = 33
	/*
	   Comma operator.
	*/
	BinaryOperator_Last BinaryOperatorKind = 33
)

/*
Describes the kind of unary operators.
*/
type UnaryOperatorKind uint32

const (
	/*
	   This value describes cursors which are not unary operators.
	*/
	UnaryOperator_Invalid UnaryOperatorKind = 0
	/*
	   Postfix increment operator.
	*/
	UnaryOperator_PostInc UnaryOperatorKind = 1
	/*
	   Postfix decrement operator.
	*/
	UnaryOperator_PostDec UnaryOperatorKind = 2
	/*
	   Prefix increment operator.
	*/
	UnaryOperator_PreInc UnaryOperatorKind = 3
	/*
	   Prefix decrement operator.
	*/
	UnaryOperator_PreDec UnaryOperatorKind = 4
	/*
	   Address of operator.
	*/
	UnaryOperator_AddrOf UnaryOperatorKind = 5
	/*
	   Dereference operator.
	*/
	UnaryOperator_Deref UnaryOperatorKind = 6
	/*
	   Plus operator.
	*/
	UnaryOperator_Plus UnaryOperatorKind = 7
	/*
	   Minus operator.
	*/
	UnaryOperator_Minus UnaryOperatorKind = 8
	/*
	   Not operator.
	*/
	UnaryOperator_Not UnaryOperatorKind = 9
	/*
	   LNot operator.
	*/
	UnaryOperator_LNot UnaryOperatorKind = 10
	/*
	   "__real expr" operator.
	*/
	UnaryOperator_Real UnaryOperatorKind = 11
	/*
	   "__imag expr" operator.
	*/
	UnaryOperator_Imag UnaryOperatorKind = 12
	/*
	   __extension__ marker operator.
	*/
	UnaryOperator_Extension UnaryOperatorKind = 13
	/*
	   C++ co_await operator.
	*/
	UnaryOperator_Coawait UnaryOperatorKind = 14
	/*
	   C++ co_await operator.
	*/
	UnaryOperator_Last UnaryOperatorKind = 14
)
