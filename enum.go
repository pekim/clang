// This is a generated file. DO NOT EDIT.

package clang

type CXErrorCode uint32

const (
	CXError_Success          CXErrorCode = 0
	CXError_Failure          CXErrorCode = 1
	CXError_Crashed          CXErrorCode = 2
	CXError_InvalidArguments CXErrorCode = 3
	CXError_ASTReadError     CXErrorCode = 4
)

type CXDiagnosticSeverity uint32

const (
	CXDiagnostic_Ignored CXDiagnosticSeverity = 0
	CXDiagnostic_Note    CXDiagnosticSeverity = 1
	CXDiagnostic_Warning CXDiagnosticSeverity = 2
	CXDiagnostic_Error   CXDiagnosticSeverity = 3
	CXDiagnostic_Fatal   CXDiagnosticSeverity = 4
)

type CXLoadDiag_Error uint32

const (
	CXLoadDiag_None        CXLoadDiag_Error = 0
	CXLoadDiag_Unknown     CXLoadDiag_Error = 1
	CXLoadDiag_CannotLoad  CXLoadDiag_Error = 2
	CXLoadDiag_InvalidFile CXLoadDiag_Error = 3
)

type CXDiagnosticDisplayOptions uint32

const (
	CXDiagnostic_DisplaySourceLocation CXDiagnosticDisplayOptions = 1
	CXDiagnostic_DisplayColumn         CXDiagnosticDisplayOptions = 2
	CXDiagnostic_DisplaySourceRanges   CXDiagnosticDisplayOptions = 4
	CXDiagnostic_DisplayOption         CXDiagnosticDisplayOptions = 8
	CXDiagnostic_DisplayCategoryId     CXDiagnosticDisplayOptions = 16
	CXDiagnostic_DisplayCategoryName   CXDiagnosticDisplayOptions = 32
)

/*
Describes the availability of a particular entity, which indicates whether the use of this entity will result in a warning or error due to it being deprecated or unavailable.
*/
type CXAvailabilityKind uint32

const (
	/*
	   The entity is available.
	*/
	CXAvailability_Available CXAvailabilityKind = 0
	/*
	   The entity is available, but has been deprecated (and its use is not recommended).
	*/
	CXAvailability_Deprecated CXAvailabilityKind = 1
	/*
	   The entity is not available; any use of it will be an error.
	*/
	CXAvailability_NotAvailable CXAvailabilityKind = 2
	/*
	   The entity is available, but not accessible; any use of it will be an error.
	*/
	CXAvailability_NotAccessible CXAvailabilityKind = 3
)

/*
Describes the exception specification of a cursor.

A negative value indicates that the cursor is not a function declaration.
*/
type CXCursor_ExceptionSpecificationKind uint32

const (
	/*
	   The cursor has no exception specification.
	*/
	CXCursor_ExceptionSpecificationKind_None CXCursor_ExceptionSpecificationKind = 0
	/*
	   The cursor has exception specification throw()
	*/
	CXCursor_ExceptionSpecificationKind_DynamicNone CXCursor_ExceptionSpecificationKind = 1
	/*
	   The cursor has exception specification throw(T1, T2)
	*/
	CXCursor_ExceptionSpecificationKind_Dynamic CXCursor_ExceptionSpecificationKind = 2
	/*
	   The cursor has exception specification throw(...).
	*/
	CXCursor_ExceptionSpecificationKind_MSAny CXCursor_ExceptionSpecificationKind = 3
	/*
	   The cursor has exception specification basic noexcept.
	*/
	CXCursor_ExceptionSpecificationKind_BasicNoexcept CXCursor_ExceptionSpecificationKind = 4
	/*
	   The cursor has exception specification computed noexcept.
	*/
	CXCursor_ExceptionSpecificationKind_ComputedNoexcept CXCursor_ExceptionSpecificationKind = 5
	/*
	   The exception specification has not yet been evaluated.
	*/
	CXCursor_ExceptionSpecificationKind_Unevaluated CXCursor_ExceptionSpecificationKind = 6
	/*
	   The exception specification has not yet been instantiated.
	*/
	CXCursor_ExceptionSpecificationKind_Uninstantiated CXCursor_ExceptionSpecificationKind = 7
	/*
	   The exception specification has not been parsed yet.
	*/
	CXCursor_ExceptionSpecificationKind_Unparsed CXCursor_ExceptionSpecificationKind = 8
	/*
	   The cursor has a __declspec(nothrow) exception specification.
	*/
	CXCursor_ExceptionSpecificationKind_NoThrow CXCursor_ExceptionSpecificationKind = 9
)

type CXChoice uint32

const (
	/*
	   Use the default value of an option that may depend on the process environment.
	*/
	CXChoice_Default CXChoice = 0
	/*
	   Enable the option.
	*/
	CXChoice_Enabled CXChoice = 1
	/*
	   Disable the option.
	*/
	CXChoice_Disabled CXChoice = 2
)

type CXGlobalOptFlags uint32

const (
	/*
	   Used to indicate that no special CXIndex options are needed.
	*/
	CXGlobalOpt_None CXGlobalOptFlags = 0
	/*
	   Used to indicate that threads that libclang creates for indexing purposes should use background priority.

	   Affects #clang_indexSourceFile, #clang_indexTranslationUnit, #clang_parseTranslationUnit, #clang_saveTranslationUnit.
	*/
	CXGlobalOpt_ThreadBackgroundPriorityForIndexing CXGlobalOptFlags = 1
	/*
	   Used to indicate that threads that libclang creates for editing purposes should use background priority.

	   Affects #clang_reparseTranslationUnit, #clang_codeCompleteAt, #clang_annotateTokens
	*/
	CXGlobalOpt_ThreadBackgroundPriorityForEditing CXGlobalOptFlags = 2
	/*
	   Used to indicate that all threads that libclang creates should use background priority.
	*/
	CXGlobalOpt_ThreadBackgroundPriorityForAll CXGlobalOptFlags = 3
)

/*
Flags that control the creation of translation units.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when constructing the translation unit.
*/
type CXTranslationUnit_Flags uint32

const (
	/*
	   Used to indicate that no special translation-unit options are needed.
	*/
	CXTranslationUnit_None CXTranslationUnit_Flags = 0
	/*
	   Used to indicate that the parser should construct a "detailed" preprocessing record, including all macro definitions and instantiations.

	   Constructing a detailed preprocessing record requires more memory and time to parse, since the information contained in the record is usually not retained. However, it can be useful for applications that require more detailed information about the behavior of the preprocessor.
	*/
	CXTranslationUnit_DetailedPreprocessingRecord CXTranslationUnit_Flags = 1
	/*
	   Used to indicate that the translation unit is incomplete.

	   When a translation unit is considered "incomplete", semantic analysis that is typically performed at the end of the translation unit will be suppressed. For example, this suppresses the completion of tentative declarations in C and of instantiation of implicitly-instantiation function templates in C++. This option is typically used when parsing a header with the intent of producing a precompiled header.
	*/
	CXTranslationUnit_Incomplete CXTranslationUnit_Flags = 2
	/*
	   Used to indicate that the translation unit should be built with an implicit precompiled header for the preamble.

	   An implicit precompiled header is used as an optimization when a particular translation unit is likely to be reparsed many times when the sources aren't changing that often. In this case, an implicit precompiled header will be built containing all of the initial includes at the top of the main file (what we refer to as the "preamble" of the file). In subsequent parses, if the preamble or the files in it have not changed, clang_reparseTranslationUnit() will re-use the implicit precompiled header to improve parsing performance.
	*/
	CXTranslationUnit_PrecompiledPreamble CXTranslationUnit_Flags = 4
	/*
	   Used to indicate that the translation unit should cache some code-completion results with each reparse of the source file.

	   Caching of code-completion results is a performance optimization that introduces some overhead to reparsing but improves the performance of code-completion operations.
	*/
	CXTranslationUnit_CacheCompletionResults CXTranslationUnit_Flags = 8
	/*
	   Used to indicate that the translation unit will be serialized with clang_saveTranslationUnit.

	   This option is typically used when parsing a header with the intent of producing a precompiled header.
	*/
	CXTranslationUnit_ForSerialization CXTranslationUnit_Flags = 16
	/*
	   DEPRECATED: Enabled chained precompiled preambles in C++.

	   Note: this is a *temporary* option that is available only while we are testing C++ precompiled preamble support. It is deprecated.
	*/
	CXTranslationUnit_CXXChainedPCH CXTranslationUnit_Flags = 32
	/*
	   Used to indicate that function/method bodies should be skipped while parsing.

	   This option can be used to search for declarations/definitions while ignoring the usages.
	*/
	CXTranslationUnit_SkipFunctionBodies CXTranslationUnit_Flags = 64
	/*
	   Used to indicate that brief documentation comments should be included into the set of code completions returned from this translation unit.
	*/
	CXTranslationUnit_IncludeBriefCommentsInCodeCompletion CXTranslationUnit_Flags = 128
	/*
	   Used to indicate that the precompiled preamble should be created on the first parse. Otherwise it will be created on the first reparse. This trades runtime on the first parse (serializing the preamble takes time) for reduced runtime on the second parse (can now reuse the preamble).
	*/
	CXTranslationUnit_CreatePreambleOnFirstParse CXTranslationUnit_Flags = 256
	/*
	   Do not stop processing when fatal errors are encountered.

	   When fatal errors are encountered while parsing a translation unit, semantic analysis is typically stopped early when compiling code. A common source for fatal errors are unresolvable include files. For the purposes of an IDE, this is undesirable behavior and as much information as possible should be reported. Use this flag to enable this behavior.
	*/
	CXTranslationUnit_KeepGoing CXTranslationUnit_Flags = 512
	/*
	   Sets the preprocessor in a mode for parsing a single file only.
	*/
	CXTranslationUnit_SingleFileParse CXTranslationUnit_Flags = 1024
	/*
	   Used in combination with CXTranslationUnit_SkipFunctionBodies to constrain the skipping of function bodies to the preamble.

	   The function bodies of the main file are not skipped.
	*/
	CXTranslationUnit_LimitSkipFunctionBodiesToPreamble CXTranslationUnit_Flags = 2048
	/*
	   Used to indicate that attributed types should be included in CXType.
	*/
	CXTranslationUnit_IncludeAttributedTypes CXTranslationUnit_Flags = 4096
	/*
	   Used to indicate that implicit attributes should be visited.
	*/
	CXTranslationUnit_VisitImplicitAttributes CXTranslationUnit_Flags = 8192
	/*
	   Used to indicate that non-errors from included files should be ignored.

	   If set, clang_getDiagnosticSetFromTU() will not report e.g. warnings from included files anymore. This speeds up clang_getDiagnosticSetFromTU() for the case where these warnings are not of interest, as for an IDE for example, which typically shows only the diagnostics in the main file.
	*/
	CXTranslationUnit_IgnoreNonErrorsFromIncludedFiles CXTranslationUnit_Flags = 16384
	/*
	   Tells the preprocessor not to skip excluded conditional blocks.
	*/
	CXTranslationUnit_RetainExcludedConditionalBlocks CXTranslationUnit_Flags = 32768
)

/*
Flags that control how translation units are saved.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when saving the translation unit.
*/
type CXSaveTranslationUnit_Flags uint32

const (
	/*
	   Used to indicate that no special saving options are needed.
	*/
	CXSaveTranslationUnit_None CXSaveTranslationUnit_Flags = 0
)

/*
Describes the kind of error that occurred (if any) in a call to clang_saveTranslationUnit().
*/
type CXSaveError uint32

const (
	/*
	   Indicates that no error occurred while saving a translation unit.
	*/
	CXSaveError_None CXSaveError = 0
	/*
	   Indicates that an unknown error occurred while attempting to save the file.

	   This error typically indicates that file I/O failed when attempting to write the file.
	*/
	CXSaveError_Unknown CXSaveError = 1
	/*
	   Indicates that errors during translation prevented this attempt to save the translation unit.

	   Errors that prevent the translation unit from being saved can be extracted using clang_getNumDiagnostics() and clang_getDiagnostic().
	*/
	CXSaveError_TranslationErrors CXSaveError = 2
	/*
	   Indicates that the translation unit to be saved was somehow invalid (e.g., NULL).
	*/
	CXSaveError_InvalidTU CXSaveError = 3
)

/*
Flags that control the reparsing of translation units.

The enumerators in this enumeration type are meant to be bitwise ORed together to specify which options should be used when reparsing the translation unit.
*/
type CXReparse_Flags uint32

const (
	/*
	   Used to indicate that no special reparsing options are needed.
	*/
	CXReparse_None CXReparse_Flags = 0
)

/*
Categorizes how memory is being used by a translation unit.
*/
type CXTUResourceUsageKind uint32

const (
	CXTUResourceUsage_AST                                CXTUResourceUsageKind = 1
	CXTUResourceUsage_Identifiers                        CXTUResourceUsageKind = 2
	CXTUResourceUsage_Selectors                          CXTUResourceUsageKind = 3
	CXTUResourceUsage_GlobalCompletionResults            CXTUResourceUsageKind = 4
	CXTUResourceUsage_SourceManagerContentCache          CXTUResourceUsageKind = 5
	CXTUResourceUsage_AST_SideTables                     CXTUResourceUsageKind = 6
	CXTUResourceUsage_SourceManager_Membuffer_Malloc     CXTUResourceUsageKind = 7
	CXTUResourceUsage_SourceManager_Membuffer_MMap       CXTUResourceUsageKind = 8
	CXTUResourceUsage_ExternalASTSource_Membuffer_Malloc CXTUResourceUsageKind = 9
	CXTUResourceUsage_ExternalASTSource_Membuffer_MMap   CXTUResourceUsageKind = 10
	CXTUResourceUsage_Preprocessor                       CXTUResourceUsageKind = 11
	CXTUResourceUsage_PreprocessingRecord                CXTUResourceUsageKind = 12
	CXTUResourceUsage_SourceManager_DataStructures       CXTUResourceUsageKind = 13
	CXTUResourceUsage_Preprocessor_HeaderSearch          CXTUResourceUsageKind = 14
	CXTUResourceUsage_MEMORY_IN_BYTES_BEGIN              CXTUResourceUsageKind = 1
	CXTUResourceUsage_MEMORY_IN_BYTES_END                CXTUResourceUsageKind = 14
	CXTUResourceUsage_First                              CXTUResourceUsageKind = 1
	CXTUResourceUsage_Last                               CXTUResourceUsageKind = 14
)

/*
Describes the kind of entity that a cursor refers to.
*/
type CXCursorKind uint32

const (
	/*
	   A declaration whose specific kind is not exposed via this interface.

	   Unexposed declarations have the same operations as any other kind of declaration; one can extract their location information, spelling, find their definitions, etc. However, the specific kind of the declaration is not reported.
	*/
	CXCursor_UnexposedDecl CXCursorKind = 1
	/*
	   A C or C++ struct.
	*/
	CXCursor_StructDecl CXCursorKind = 2
	/*
	   A C or C++ union.
	*/
	CXCursor_UnionDecl CXCursorKind = 3
	/*
	   A C++ class.
	*/
	CXCursor_ClassDecl CXCursorKind = 4
	/*
	   An enumeration.
	*/
	CXCursor_EnumDecl CXCursorKind = 5
	/*
	   A field (in C) or non-static data member (in C++) in a struct, union, or C++ class.
	*/
	CXCursor_FieldDecl CXCursorKind = 6
	/*
	   An enumerator constant.
	*/
	CXCursor_EnumConstantDecl CXCursorKind = 7
	/*
	   A function.
	*/
	CXCursor_FunctionDecl CXCursorKind = 8
	/*
	   A variable.
	*/
	CXCursor_VarDecl CXCursorKind = 9
	/*
	   A function or method parameter.
	*/
	CXCursor_ParmDecl CXCursorKind = 10
	/*
	   An Objective-C @interface.
	*/
	CXCursor_ObjCInterfaceDecl CXCursorKind = 11
	/*
	   An Objective-C @interface for a category.
	*/
	CXCursor_ObjCCategoryDecl CXCursorKind = 12
	/*
	   An Objective-C @protocol declaration.
	*/
	CXCursor_ObjCProtocolDecl CXCursorKind = 13
	/*
	   An Objective-C @property declaration.
	*/
	CXCursor_ObjCPropertyDecl CXCursorKind = 14
	/*
	   An Objective-C instance variable.
	*/
	CXCursor_ObjCIvarDecl CXCursorKind = 15
	/*
	   An Objective-C instance method.
	*/
	CXCursor_ObjCInstanceMethodDecl CXCursorKind = 16
	/*
	   An Objective-C class method.
	*/
	CXCursor_ObjCClassMethodDecl CXCursorKind = 17
	/*
	   An Objective-C @implementation.
	*/
	CXCursor_ObjCImplementationDecl CXCursorKind = 18
	/*
	   An Objective-C @implementation for a category.
	*/
	CXCursor_ObjCCategoryImplDecl CXCursorKind = 19
	/*
	   A typedef.
	*/
	CXCursor_TypedefDecl CXCursorKind = 20
	/*
	   A C++ class method.
	*/
	CXCursor_CXXMethod CXCursorKind = 21
	/*
	   A C++ namespace.
	*/
	CXCursor_Namespace CXCursorKind = 22
	/*
	   A linkage specification, e.g. 'extern "C"'.
	*/
	CXCursor_LinkageSpec CXCursorKind = 23
	/*
	   A C++ constructor.
	*/
	CXCursor_Constructor CXCursorKind = 24
	/*
	   A C++ destructor.
	*/
	CXCursor_Destructor CXCursorKind = 25
	/*
	   A C++ conversion function.
	*/
	CXCursor_ConversionFunction CXCursorKind = 26
	/*
	   A C++ template type parameter.
	*/
	CXCursor_TemplateTypeParameter CXCursorKind = 27
	/*
	   A C++ non-type template parameter.
	*/
	CXCursor_NonTypeTemplateParameter CXCursorKind = 28
	/*
	   A C++ template template parameter.
	*/
	CXCursor_TemplateTemplateParameter CXCursorKind = 29
	/*
	   A C++ function template.
	*/
	CXCursor_FunctionTemplate CXCursorKind = 30
	/*
	   A C++ class template.
	*/
	CXCursor_ClassTemplate CXCursorKind = 31
	/*
	   A C++ class template partial specialization.
	*/
	CXCursor_ClassTemplatePartialSpecialization CXCursorKind = 32
	/*
	   A C++ namespace alias declaration.
	*/
	CXCursor_NamespaceAlias CXCursorKind = 33
	/*
	   A C++ using directive.
	*/
	CXCursor_UsingDirective CXCursorKind = 34
	/*
	   A C++ using declaration.
	*/
	CXCursor_UsingDeclaration CXCursorKind = 35
	/*
	   A C++ alias declaration
	*/
	CXCursor_TypeAliasDecl CXCursorKind = 36
	/*
	   An Objective-C @synthesize definition.
	*/
	CXCursor_ObjCSynthesizeDecl CXCursorKind = 37
	/*
	   An Objective-C @dynamic definition.
	*/
	CXCursor_ObjCDynamicDecl CXCursorKind = 38
	/*
	   An access specifier.
	*/
	CXCursor_CXXAccessSpecifier CXCursorKind = 39
	/*
	   An access specifier.
	*/
	CXCursor_FirstDecl CXCursorKind = 1
	/*
	   An access specifier.
	*/
	CXCursor_LastDecl CXCursorKind = 39
	/*
	   An access specifier.
	*/
	CXCursor_FirstRef CXCursorKind = 40
	/*
	   An access specifier.
	*/
	CXCursor_ObjCSuperClassRef CXCursorKind = 40
	/*
	   An access specifier.
	*/
	CXCursor_ObjCProtocolRef CXCursorKind = 41
	/*
	   An access specifier.
	*/
	CXCursor_ObjCClassRef CXCursorKind = 42
	/*
	   A reference to a type declaration.

	   A type reference occurs anywhere where a type is named but not declared. For example, given:

	   The typedef is a declaration of size_type (CXCursor_TypedefDecl), while the type of the variable "size" is referenced. The cursor referenced by the type of size is the typedef for size_type.
	*/
	CXCursor_TypeRef CXCursorKind = 43
	/*
	   A reference to a type declaration.

	   A type reference occurs anywhere where a type is named but not declared. For example, given:

	   The typedef is a declaration of size_type (CXCursor_TypedefDecl), while the type of the variable "size" is referenced. The cursor referenced by the type of size is the typedef for size_type.
	*/
	CXCursor_CXXBaseSpecifier CXCursorKind = 44
	/*
	   A reference to a class template, function template, template template parameter, or class template partial specialization.
	*/
	CXCursor_TemplateRef CXCursorKind = 45
	/*
	   A reference to a namespace or namespace alias.
	*/
	CXCursor_NamespaceRef CXCursorKind = 46
	/*
	   A reference to a member of a struct, union, or class that occurs in some non-expression context, e.g., a designated initializer.
	*/
	CXCursor_MemberRef CXCursorKind = 47
	/*
	   A reference to a labeled statement.

	   This cursor kind is used to describe the jump to "start_over" in the goto statement in the following example:

	   A label reference cursor refers to a label statement.
	*/
	CXCursor_LabelRef CXCursorKind = 48
	/*
	   A reference to a set of overloaded functions or function templates that has not yet been resolved to a specific function or function template.

	   An overloaded declaration reference cursor occurs in C++ templates where a dependent name refers to a function. For example:

	   Here, the identifier "swap" is associated with an overloaded declaration reference. In the template definition, "swap" refers to either of the two "swap" functions declared above, so both results will be available. At instantiation time, "swap" may also refer to other functions found via argument-dependent lookup (e.g., the "swap" function at the end of the example).

	   The functions clang_getNumOverloadedDecls() and clang_getOverloadedDecl() can be used to retrieve the definitions referenced by this cursor.
	*/
	CXCursor_OverloadedDeclRef CXCursorKind = 49
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_VariableRef CXCursorKind = 50
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_LastRef CXCursorKind = 50
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_FirstInvalid CXCursorKind = 70
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_InvalidFile CXCursorKind = 70
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_NoDeclFound CXCursorKind = 71
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_NotImplemented CXCursorKind = 72
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_InvalidCode CXCursorKind = 73
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_LastInvalid CXCursorKind = 73
	/*
	   A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	*/
	CXCursor_FirstExpr CXCursorKind = 100
	/*
	   An expression whose specific kind is not exposed via this interface.

	   Unexposed expressions have the same operations as any other kind of expression; one can extract their location information, spelling, children, etc. However, the specific kind of the expression is not reported.
	*/
	CXCursor_UnexposedExpr CXCursorKind = 100
	/*
	   An expression that refers to some value declaration, such as a function, variable, or enumerator.
	*/
	CXCursor_DeclRefExpr CXCursorKind = 101
	/*
	   An expression that refers to a member of a struct, union, class, Objective-C class, etc.
	*/
	CXCursor_MemberRefExpr CXCursorKind = 102
	/*
	   An expression that calls a function.
	*/
	CXCursor_CallExpr CXCursorKind = 103
	/*
	   An expression that sends a message to an Objective-C   object or class.
	*/
	CXCursor_ObjCMessageExpr CXCursorKind = 104
	/*
	   An expression that represents a block literal.
	*/
	CXCursor_BlockExpr CXCursorKind = 105
	/*
	   An integer literal.
	*/
	CXCursor_IntegerLiteral CXCursorKind = 106
	/*
	   A floating point number literal.
	*/
	CXCursor_FloatingLiteral CXCursorKind = 107
	/*
	   An imaginary number literal.
	*/
	CXCursor_ImaginaryLiteral CXCursorKind = 108
	/*
	   A string literal.
	*/
	CXCursor_StringLiteral CXCursorKind = 109
	/*
	   A character literal.
	*/
	CXCursor_CharacterLiteral CXCursorKind = 110
	/*
	   A parenthesized expression, e.g. "(1)".

	   This AST node is only formed if full location information is requested.
	*/
	CXCursor_ParenExpr CXCursorKind = 111
	/*
	   This represents the unary-expression's (except sizeof and alignof).
	*/
	CXCursor_UnaryOperator CXCursorKind = 112
	/*
	   [C99 6.5.2.1] Array Subscripting.
	*/
	CXCursor_ArraySubscriptExpr CXCursorKind = 113
	/*
	   A builtin binary operation expression such as "x + y" or "x <= y".
	*/
	CXCursor_BinaryOperator CXCursorKind = 114
	/*
	   Compound assignment such as "+=".
	*/
	CXCursor_CompoundAssignOperator CXCursorKind = 115
	/*
	   The ?: ternary operator.
	*/
	CXCursor_ConditionalOperator CXCursorKind = 116
	/*
	   An explicit cast in C (C99 6.5.4) or a C-style cast in C++ (C++ [expr.cast]), which uses the syntax (Type)expr.

	   For example: (int)f.
	*/
	CXCursor_CStyleCastExpr CXCursorKind = 117
	/*
	   [C99 6.5.2.5]
	*/
	CXCursor_CompoundLiteralExpr CXCursorKind = 118
	/*
	   Describes an C or C++ initializer list.
	*/
	CXCursor_InitListExpr CXCursorKind = 119
	/*
	   The GNU address of label extension, representing &&label.
	*/
	CXCursor_AddrLabelExpr CXCursorKind = 120
	/*
	   This is the GNU Statement Expression extension: ({int X=4; X;})
	*/
	CXCursor_StmtExpr CXCursorKind = 121
	/*
	   Represents a C11 generic selection.
	*/
	CXCursor_GenericSelectionExpr CXCursorKind = 122
	/*
	   Implements the GNU __null extension, which is a name for a null pointer constant that has integral type (e.g., int or long) and is the same size and alignment as a pointer.

	   The __null extension is typically only used by system headers, which define NULL as __null in C++ rather than using 0 (which is an integer that may not match the size of a pointer).
	*/
	CXCursor_GNUNullExpr CXCursorKind = 123
	/*
	   C++'s static_cast<> expression.
	*/
	CXCursor_CXXStaticCastExpr CXCursorKind = 124
	/*
	   C++'s dynamic_cast<> expression.
	*/
	CXCursor_CXXDynamicCastExpr CXCursorKind = 125
	/*
	   C++'s reinterpret_cast<> expression.
	*/
	CXCursor_CXXReinterpretCastExpr CXCursorKind = 126
	/*
	   C++'s const_cast<> expression.
	*/
	CXCursor_CXXConstCastExpr CXCursorKind = 127
	/*
	   Represents an explicit C++ type conversion that uses "functional" notion (C++ [expr.type.conv]).

	   Example:
	*/
	CXCursor_CXXFunctionalCastExpr CXCursorKind = 128
	/*
	   A C++ typeid expression (C++ [expr.typeid]).
	*/
	CXCursor_CXXTypeidExpr CXCursorKind = 129
	/*
	   [C++ 2.13.5] C++ Boolean Literal.
	*/
	CXCursor_CXXBoolLiteralExpr CXCursorKind = 130
	/*
	   [C++0x 2.14.7] C++ Pointer Literal.
	*/
	CXCursor_CXXNullPtrLiteralExpr CXCursorKind = 131
	/*
	   Represents the "this" expression in C++
	*/
	CXCursor_CXXThisExpr CXCursorKind = 132
	/*
	   [C++ 15] C++ Throw Expression.

	   This handles 'throw' and 'throw' assignment-expression. When assignment-expression isn't present, Op will be null.
	*/
	CXCursor_CXXThrowExpr CXCursorKind = 133
	/*
	   A new expression for memory allocation and constructor calls, e.g: "new CXXNewExpr(foo)".
	*/
	CXCursor_CXXNewExpr CXCursorKind = 134
	/*
	   A delete expression for memory deallocation and destructor calls, e.g. "delete[] pArray".
	*/
	CXCursor_CXXDeleteExpr CXCursorKind = 135
	/*
	   A unary expression. (noexcept, sizeof, or other traits)
	*/
	CXCursor_UnaryExpr CXCursorKind = 136
	/*
	   An Objective-C string literal i.e. "foo".
	*/
	CXCursor_ObjCStringLiteral CXCursorKind = 137
	/*
	   An Objective-C @encode expression.
	*/
	CXCursor_ObjCEncodeExpr CXCursorKind = 138
	/*
	   An Objective-C @selector expression.
	*/
	CXCursor_ObjCSelectorExpr CXCursorKind = 139
	/*
	   An Objective-C @protocol expression.
	*/
	CXCursor_ObjCProtocolExpr CXCursorKind = 140
	/*
	   An Objective-C "bridged" cast expression, which casts between Objective-C pointers and C pointers, transferring ownership in the process.
	*/
	CXCursor_ObjCBridgedCastExpr CXCursorKind = 141
	/*
	   Represents a C++0x pack expansion that produces a sequence of expressions.

	   A pack expansion expression contains a pattern (which itself is an expression) followed by an ellipsis. For example:
	*/
	CXCursor_PackExpansionExpr CXCursorKind = 142
	/*
	   Represents an expression that computes the length of a parameter pack.
	*/
	CXCursor_SizeOfPackExpr CXCursorKind = 143
	CXCursor_LambdaExpr     CXCursorKind = 144
	/*
	   Objective-c Boolean Literal.
	*/
	CXCursor_ObjCBoolLiteralExpr CXCursorKind = 145
	/*
	   Represents the "self" expression in an Objective-C method.
	*/
	CXCursor_ObjCSelfExpr CXCursorKind = 146
	/*
	   OpenMP 5.0 [2.1.5, Array Section]. OpenACC 3.3 [2.7.1, Data Specification for Data Clauses (Sub Arrays)]
	*/
	CXCursor_ArraySectionExpr CXCursorKind = 147
	/*
	   Represents an (...) check.
	*/
	CXCursor_ObjCAvailabilityCheckExpr CXCursorKind = 148
	/*
	   Fixed point literal
	*/
	CXCursor_FixedPointLiteral CXCursorKind = 149
	/*
	   OpenMP 5.0 [2.1.4, Array Shaping].
	*/
	CXCursor_OMPArrayShapingExpr CXCursorKind = 150
	/*
	   OpenMP 5.0 [2.1.6 Iterators]
	*/
	CXCursor_OMPIteratorExpr CXCursorKind = 151
	/*
	   OpenCL's addrspace_cast<> expression.
	*/
	CXCursor_CXXAddrspaceCastExpr CXCursorKind = 152
	/*
	   Expression that references a C++20 concept.
	*/
	CXCursor_ConceptSpecializationExpr CXCursorKind = 153
	/*
	   Expression that references a C++20 requires expression.
	*/
	CXCursor_RequiresExpr CXCursorKind = 154
	/*
	   Expression that references a C++20 parenthesized list aggregate initializer.
	*/
	CXCursor_CXXParenListInitExpr CXCursorKind = 155
	/*
	   Represents a C++26 pack indexing expression.
	*/
	CXCursor_PackIndexingExpr CXCursorKind = 156
	/*
	   Represents a C++26 pack indexing expression.
	*/
	CXCursor_LastExpr CXCursorKind = 156
	/*
	   Represents a C++26 pack indexing expression.
	*/
	CXCursor_FirstStmt CXCursorKind = 200
	/*
	   A statement whose specific kind is not exposed via this interface.

	   Unexposed statements have the same operations as any other kind of statement; one can extract their location information, spelling, children, etc. However, the specific kind of the statement is not reported.
	*/
	CXCursor_UnexposedStmt CXCursorKind = 200
	/*
	   A labelled statement in a function.

	   This cursor kind is used to describe the "start_over:" label statement in the following example:
	*/
	CXCursor_LabelStmt CXCursorKind = 201
	/*
	   A group of statements like { stmt stmt }.

	   This cursor kind is used to describe compound statements, e.g. function bodies.
	*/
	CXCursor_CompoundStmt CXCursorKind = 202
	/*
	   A case statement.
	*/
	CXCursor_CaseStmt CXCursorKind = 203
	/*
	   A default statement.
	*/
	CXCursor_DefaultStmt CXCursorKind = 204
	/*
	   An if statement
	*/
	CXCursor_IfStmt CXCursorKind = 205
	/*
	   A switch statement.
	*/
	CXCursor_SwitchStmt CXCursorKind = 206
	/*
	   A while statement.
	*/
	CXCursor_WhileStmt CXCursorKind = 207
	/*
	   A do statement.
	*/
	CXCursor_DoStmt CXCursorKind = 208
	/*
	   A for statement.
	*/
	CXCursor_ForStmt CXCursorKind = 209
	/*
	   A goto statement.
	*/
	CXCursor_GotoStmt CXCursorKind = 210
	/*
	   An indirect goto statement.
	*/
	CXCursor_IndirectGotoStmt CXCursorKind = 211
	/*
	   A continue statement.
	*/
	CXCursor_ContinueStmt CXCursorKind = 212
	/*
	   A break statement.
	*/
	CXCursor_BreakStmt CXCursorKind = 213
	/*
	   A return statement.
	*/
	CXCursor_ReturnStmt CXCursorKind = 214
	/*
	   A GCC inline assembly statement extension.
	*/
	CXCursor_GCCAsmStmt CXCursorKind = 215
	/*
	   A GCC inline assembly statement extension.
	*/
	CXCursor_AsmStmt CXCursorKind = 215
	/*
	   Objective-C's overall @try-@catch-@finally statement.
	*/
	CXCursor_ObjCAtTryStmt CXCursorKind = 216
	/*
	   Objective-C's @catch statement.
	*/
	CXCursor_ObjCAtCatchStmt CXCursorKind = 217
	/*
	   Objective-C's @finally statement.
	*/
	CXCursor_ObjCAtFinallyStmt CXCursorKind = 218
	/*
	   Objective-C's @throw statement.
	*/
	CXCursor_ObjCAtThrowStmt CXCursorKind = 219
	/*
	   Objective-C's @synchronized statement.
	*/
	CXCursor_ObjCAtSynchronizedStmt CXCursorKind = 220
	/*
	   Objective-C's autorelease pool statement.
	*/
	CXCursor_ObjCAutoreleasePoolStmt CXCursorKind = 221
	/*
	   Objective-C's collection statement.
	*/
	CXCursor_ObjCForCollectionStmt CXCursorKind = 222
	/*
	   C++'s catch statement.
	*/
	CXCursor_CXXCatchStmt CXCursorKind = 223
	/*
	   C++'s try statement.
	*/
	CXCursor_CXXTryStmt CXCursorKind = 224
	/*
	   C++'s for (* : *) statement.
	*/
	CXCursor_CXXForRangeStmt CXCursorKind = 225
	/*
	   Windows Structured Exception Handling's try statement.
	*/
	CXCursor_SEHTryStmt CXCursorKind = 226
	/*
	   Windows Structured Exception Handling's except statement.
	*/
	CXCursor_SEHExceptStmt CXCursorKind = 227
	/*
	   Windows Structured Exception Handling's finally statement.
	*/
	CXCursor_SEHFinallyStmt CXCursorKind = 228
	/*
	   A MS inline assembly statement extension.
	*/
	CXCursor_MSAsmStmt CXCursorKind = 229
	/*
	   The null statement ";": C99 6.8.3p3.

	   This cursor kind is used to describe the null statement.
	*/
	CXCursor_NullStmt CXCursorKind = 230
	/*
	   Adaptor class for mixing declarations with statements and expressions.
	*/
	CXCursor_DeclStmt CXCursorKind = 231
	/*
	   OpenMP parallel directive.
	*/
	CXCursor_OMPParallelDirective CXCursorKind = 232
	/*
	   OpenMP SIMD directive.
	*/
	CXCursor_OMPSimdDirective CXCursorKind = 233
	/*
	   OpenMP for directive.
	*/
	CXCursor_OMPForDirective CXCursorKind = 234
	/*
	   OpenMP sections directive.
	*/
	CXCursor_OMPSectionsDirective CXCursorKind = 235
	/*
	   OpenMP section directive.
	*/
	CXCursor_OMPSectionDirective CXCursorKind = 236
	/*
	   OpenMP single directive.
	*/
	CXCursor_OMPSingleDirective CXCursorKind = 237
	/*
	   OpenMP parallel for directive.
	*/
	CXCursor_OMPParallelForDirective CXCursorKind = 238
	/*
	   OpenMP parallel sections directive.
	*/
	CXCursor_OMPParallelSectionsDirective CXCursorKind = 239
	/*
	   OpenMP task directive.
	*/
	CXCursor_OMPTaskDirective CXCursorKind = 240
	/*
	   OpenMP master directive.
	*/
	CXCursor_OMPMasterDirective CXCursorKind = 241
	/*
	   OpenMP critical directive.
	*/
	CXCursor_OMPCriticalDirective CXCursorKind = 242
	/*
	   OpenMP taskyield directive.
	*/
	CXCursor_OMPTaskyieldDirective CXCursorKind = 243
	/*
	   OpenMP barrier directive.
	*/
	CXCursor_OMPBarrierDirective CXCursorKind = 244
	/*
	   OpenMP taskwait directive.
	*/
	CXCursor_OMPTaskwaitDirective CXCursorKind = 245
	/*
	   OpenMP flush directive.
	*/
	CXCursor_OMPFlushDirective CXCursorKind = 246
	/*
	   Windows Structured Exception Handling's leave statement.
	*/
	CXCursor_SEHLeaveStmt CXCursorKind = 247
	/*
	   OpenMP ordered directive.
	*/
	CXCursor_OMPOrderedDirective CXCursorKind = 248
	/*
	   OpenMP atomic directive.
	*/
	CXCursor_OMPAtomicDirective CXCursorKind = 249
	/*
	   OpenMP for SIMD directive.
	*/
	CXCursor_OMPForSimdDirective CXCursorKind = 250
	/*
	   OpenMP parallel for SIMD directive.
	*/
	CXCursor_OMPParallelForSimdDirective CXCursorKind = 251
	/*
	   OpenMP target directive.
	*/
	CXCursor_OMPTargetDirective CXCursorKind = 252
	/*
	   OpenMP teams directive.
	*/
	CXCursor_OMPTeamsDirective CXCursorKind = 253
	/*
	   OpenMP taskgroup directive.
	*/
	CXCursor_OMPTaskgroupDirective CXCursorKind = 254
	/*
	   OpenMP cancellation point directive.
	*/
	CXCursor_OMPCancellationPointDirective CXCursorKind = 255
	/*
	   OpenMP cancel directive.
	*/
	CXCursor_OMPCancelDirective CXCursorKind = 256
	/*
	   OpenMP target data directive.
	*/
	CXCursor_OMPTargetDataDirective CXCursorKind = 257
	/*
	   OpenMP taskloop directive.
	*/
	CXCursor_OMPTaskLoopDirective CXCursorKind = 258
	/*
	   OpenMP taskloop simd directive.
	*/
	CXCursor_OMPTaskLoopSimdDirective CXCursorKind = 259
	/*
	   OpenMP distribute directive.
	*/
	CXCursor_OMPDistributeDirective CXCursorKind = 260
	/*
	   OpenMP target enter data directive.
	*/
	CXCursor_OMPTargetEnterDataDirective CXCursorKind = 261
	/*
	   OpenMP target exit data directive.
	*/
	CXCursor_OMPTargetExitDataDirective CXCursorKind = 262
	/*
	   OpenMP target parallel directive.
	*/
	CXCursor_OMPTargetParallelDirective CXCursorKind = 263
	/*
	   OpenMP target parallel for directive.
	*/
	CXCursor_OMPTargetParallelForDirective CXCursorKind = 264
	/*
	   OpenMP target update directive.
	*/
	CXCursor_OMPTargetUpdateDirective CXCursorKind = 265
	/*
	   OpenMP distribute parallel for directive.
	*/
	CXCursor_OMPDistributeParallelForDirective CXCursorKind = 266
	/*
	   OpenMP distribute parallel for simd directive.
	*/
	CXCursor_OMPDistributeParallelForSimdDirective CXCursorKind = 267
	/*
	   OpenMP distribute simd directive.
	*/
	CXCursor_OMPDistributeSimdDirective CXCursorKind = 268
	/*
	   OpenMP target parallel for simd directive.
	*/
	CXCursor_OMPTargetParallelForSimdDirective CXCursorKind = 269
	/*
	   OpenMP target simd directive.
	*/
	CXCursor_OMPTargetSimdDirective CXCursorKind = 270
	/*
	   OpenMP teams distribute directive.
	*/
	CXCursor_OMPTeamsDistributeDirective CXCursorKind = 271
	/*
	   OpenMP teams distribute simd directive.
	*/
	CXCursor_OMPTeamsDistributeSimdDirective CXCursorKind = 272
	/*
	   OpenMP teams distribute parallel for simd directive.
	*/
	CXCursor_OMPTeamsDistributeParallelForSimdDirective CXCursorKind = 273
	/*
	   OpenMP teams distribute parallel for directive.
	*/
	CXCursor_OMPTeamsDistributeParallelForDirective CXCursorKind = 274
	/*
	   OpenMP target teams directive.
	*/
	CXCursor_OMPTargetTeamsDirective CXCursorKind = 275
	/*
	   OpenMP target teams distribute directive.
	*/
	CXCursor_OMPTargetTeamsDistributeDirective CXCursorKind = 276
	/*
	   OpenMP target teams distribute parallel for directive.
	*/
	CXCursor_OMPTargetTeamsDistributeParallelForDirective CXCursorKind = 277
	/*
	   OpenMP target teams distribute parallel for simd directive.
	*/
	CXCursor_OMPTargetTeamsDistributeParallelForSimdDirective CXCursorKind = 278
	/*
	   OpenMP target teams distribute simd directive.
	*/
	CXCursor_OMPTargetTeamsDistributeSimdDirective CXCursorKind = 279
	/*
	   C++2a std::bit_cast expression.
	*/
	CXCursor_BuiltinBitCastExpr CXCursorKind = 280
	/*
	   OpenMP master taskloop directive.
	*/
	CXCursor_OMPMasterTaskLoopDirective CXCursorKind = 281
	/*
	   OpenMP parallel master taskloop directive.
	*/
	CXCursor_OMPParallelMasterTaskLoopDirective CXCursorKind = 282
	/*
	   OpenMP master taskloop simd directive.
	*/
	CXCursor_OMPMasterTaskLoopSimdDirective CXCursorKind = 283
	/*
	   OpenMP parallel master taskloop simd directive.
	*/
	CXCursor_OMPParallelMasterTaskLoopSimdDirective CXCursorKind = 284
	/*
	   OpenMP parallel master directive.
	*/
	CXCursor_OMPParallelMasterDirective CXCursorKind = 285
	/*
	   OpenMP depobj directive.
	*/
	CXCursor_OMPDepobjDirective CXCursorKind = 286
	/*
	   OpenMP scan directive.
	*/
	CXCursor_OMPScanDirective CXCursorKind = 287
	/*
	   OpenMP tile directive.
	*/
	CXCursor_OMPTileDirective CXCursorKind = 288
	/*
	   OpenMP canonical loop.
	*/
	CXCursor_OMPCanonicalLoop CXCursorKind = 289
	/*
	   OpenMP interop directive.
	*/
	CXCursor_OMPInteropDirective CXCursorKind = 290
	/*
	   OpenMP dispatch directive.
	*/
	CXCursor_OMPDispatchDirective CXCursorKind = 291
	/*
	   OpenMP masked directive.
	*/
	CXCursor_OMPMaskedDirective CXCursorKind = 292
	/*
	   OpenMP unroll directive.
	*/
	CXCursor_OMPUnrollDirective CXCursorKind = 293
	/*
	   OpenMP metadirective directive.
	*/
	CXCursor_OMPMetaDirective CXCursorKind = 294
	/*
	   OpenMP loop directive.
	*/
	CXCursor_OMPGenericLoopDirective CXCursorKind = 295
	/*
	   OpenMP teams loop directive.
	*/
	CXCursor_OMPTeamsGenericLoopDirective CXCursorKind = 296
	/*
	   OpenMP target teams loop directive.
	*/
	CXCursor_OMPTargetTeamsGenericLoopDirective CXCursorKind = 297
	/*
	   OpenMP parallel loop directive.
	*/
	CXCursor_OMPParallelGenericLoopDirective CXCursorKind = 298
	/*
	   OpenMP target parallel loop directive.
	*/
	CXCursor_OMPTargetParallelGenericLoopDirective CXCursorKind = 299
	/*
	   OpenMP parallel masked directive.
	*/
	CXCursor_OMPParallelMaskedDirective CXCursorKind = 300
	/*
	   OpenMP masked taskloop directive.
	*/
	CXCursor_OMPMaskedTaskLoopDirective CXCursorKind = 301
	/*
	   OpenMP masked taskloop simd directive.
	*/
	CXCursor_OMPMaskedTaskLoopSimdDirective CXCursorKind = 302
	/*
	   OpenMP parallel masked taskloop directive.
	*/
	CXCursor_OMPParallelMaskedTaskLoopDirective CXCursorKind = 303
	/*
	   OpenMP parallel masked taskloop simd directive.
	*/
	CXCursor_OMPParallelMaskedTaskLoopSimdDirective CXCursorKind = 304
	/*
	   OpenMP error directive.
	*/
	CXCursor_OMPErrorDirective CXCursorKind = 305
	/*
	   OpenMP scope directive.
	*/
	CXCursor_OMPScopeDirective CXCursorKind = 306
	/*
	   OpenMP reverse directive.
	*/
	CXCursor_OMPReverseDirective CXCursorKind = 307
	/*
	   OpenMP interchange directive.
	*/
	CXCursor_OMPInterchangeDirective CXCursorKind = 308
	/*
	   OpenMP assume directive.
	*/
	CXCursor_OMPAssumeDirective CXCursorKind = 309
	/*
	   OpenMP assume directive.
	*/
	CXCursor_OMPStripeDirective CXCursorKind = 310
	/*
	   OpenMP fuse directive
	*/
	CXCursor_OMPFuseDirective CXCursorKind = 311
	/*
	   OpenMP split directive.
	*/
	CXCursor_OMPSplitDirective CXCursorKind = 312
	/*
	   OpenACC Compute Construct.
	*/
	CXCursor_OpenACCComputeConstruct CXCursorKind = 320
	/*
	   OpenACC Loop Construct.
	*/
	CXCursor_OpenACCLoopConstruct CXCursorKind = 321
	/*
	   OpenACC Combined Constructs.
	*/
	CXCursor_OpenACCCombinedConstruct CXCursorKind = 322
	/*
	   OpenACC data Construct.
	*/
	CXCursor_OpenACCDataConstruct CXCursorKind = 323
	/*
	   OpenACC enter data Construct.
	*/
	CXCursor_OpenACCEnterDataConstruct CXCursorKind = 324
	/*
	   OpenACC exit data Construct.
	*/
	CXCursor_OpenACCExitDataConstruct CXCursorKind = 325
	/*
	   OpenACC host_data Construct.
	*/
	CXCursor_OpenACCHostDataConstruct CXCursorKind = 326
	/*
	   OpenACC wait Construct.
	*/
	CXCursor_OpenACCWaitConstruct CXCursorKind = 327
	/*
	   OpenACC init Construct.
	*/
	CXCursor_OpenACCInitConstruct CXCursorKind = 328
	/*
	   OpenACC shutdown Construct.
	*/
	CXCursor_OpenACCShutdownConstruct CXCursorKind = 329
	/*
	   OpenACC set Construct.
	*/
	CXCursor_OpenACCSetConstruct CXCursorKind = 330
	/*
	   OpenACC update Construct.
	*/
	CXCursor_OpenACCUpdateConstruct CXCursorKind = 331
	/*
	   OpenACC atomic Construct.
	*/
	CXCursor_OpenACCAtomicConstruct CXCursorKind = 332
	/*
	   OpenACC cache Construct.
	*/
	CXCursor_OpenACCCacheConstruct CXCursorKind = 333
	/*
	   OpenACC cache Construct.
	*/
	CXCursor_LastStmt CXCursorKind = 333
	/*
	   Cursor that represents the translation unit itself.

	   The translation unit cursor exists primarily to act as the root cursor for traversing the contents of a translation unit.
	*/
	CXCursor_TranslationUnit CXCursorKind = 350
	/*
	   Cursor that represents the translation unit itself.

	   The translation unit cursor exists primarily to act as the root cursor for traversing the contents of a translation unit.
	*/
	CXCursor_FirstAttr CXCursorKind = 400
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_UnexposedAttr CXCursorKind = 400
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_IBActionAttr CXCursorKind = 401
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_IBOutletAttr CXCursorKind = 402
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_IBOutletCollectionAttr CXCursorKind = 403
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CXXFinalAttr CXCursorKind = 404
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CXXOverrideAttr CXCursorKind = 405
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_AnnotateAttr CXCursorKind = 406
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_AsmLabelAttr CXCursorKind = 407
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_PackedAttr CXCursorKind = 408
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_PureAttr CXCursorKind = 409
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ConstAttr CXCursorKind = 410
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_NoDuplicateAttr CXCursorKind = 411
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CUDAConstantAttr CXCursorKind = 412
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CUDADeviceAttr CXCursorKind = 413
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CUDAGlobalAttr CXCursorKind = 414
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CUDAHostAttr CXCursorKind = 415
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_CUDASharedAttr CXCursorKind = 416
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_VisibilityAttr CXCursorKind = 417
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_DLLExport CXCursorKind = 418
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_DLLImport CXCursorKind = 419
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_NSReturnsRetained CXCursorKind = 420
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_NSReturnsNotRetained CXCursorKind = 421
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_NSReturnsAutoreleased CXCursorKind = 422
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_NSConsumesSelf CXCursorKind = 423
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_NSConsumed CXCursorKind = 424
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCException CXCursorKind = 425
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCNSObject CXCursorKind = 426
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCIndependentClass CXCursorKind = 427
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCPreciseLifetime CXCursorKind = 428
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCReturnsInnerPointer CXCursorKind = 429
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCRequiresSuper CXCursorKind = 430
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCRootClass CXCursorKind = 431
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCSubclassingRestricted CXCursorKind = 432
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCExplicitProtocolImpl CXCursorKind = 433
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCDesignatedInitializer CXCursorKind = 434
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCRuntimeVisible CXCursorKind = 435
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ObjCBoxable CXCursorKind = 436
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_FlagEnum CXCursorKind = 437
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_ConvergentAttr CXCursorKind = 438
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_WarnUnusedAttr CXCursorKind = 439
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_WarnUnusedResultAttr CXCursorKind = 440
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_AlignedAttr CXCursorKind = 441
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_LastAttr CXCursorKind = 441
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_PreprocessingDirective CXCursorKind = 500
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_MacroDefinition CXCursorKind = 501
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_MacroExpansion CXCursorKind = 502
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_MacroInstantiation CXCursorKind = 502
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_InclusionDirective CXCursorKind = 503
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_FirstPreprocessing CXCursorKind = 500
	/*
	   An attribute whose specific kind is not exposed via this interface.
	*/
	CXCursor_LastPreprocessing CXCursorKind = 503
	/*
	   A module import declaration.
	*/
	CXCursor_ModuleImportDecl CXCursorKind = 600
	/*
	   A module import declaration.
	*/
	CXCursor_TypeAliasTemplateDecl CXCursorKind = 601
	/*
	   A static_assert or _Static_assert node
	*/
	CXCursor_StaticAssert CXCursorKind = 602
	/*
	   a friend declaration.
	*/
	CXCursor_FriendDecl CXCursorKind = 603
	/*
	   a concept declaration.
	*/
	CXCursor_ConceptDecl CXCursorKind = 604
	/*
	   a concept declaration.
	*/
	CXCursor_FirstExtraDecl CXCursorKind = 600
	/*
	   a concept declaration.
	*/
	CXCursor_LastExtraDecl CXCursorKind = 604
	/*
	   A code completion overload candidate.
	*/
	CXCursor_OverloadCandidate CXCursorKind = 700
)

/*
Describe the linkage of the entity referred to by a cursor.
*/
type CXLinkageKind uint32

const (
	/*
	   This value indicates that no linkage information is available for a provided CXCursor.
	*/
	CXLinkage_Invalid CXLinkageKind = 0
	/*
	   This is the linkage for variables, parameters, and so on that  have automatic storage.  This covers normal (non-extern) local variables.
	*/
	CXLinkage_NoLinkage CXLinkageKind = 1
	/*
	   This is the linkage for static variables and static functions.
	*/
	CXLinkage_Internal CXLinkageKind = 2
	/*
	   This is the linkage for entities with external linkage that live in C++ anonymous namespaces.
	*/
	CXLinkage_UniqueExternal CXLinkageKind = 3
	/*
	   This is the linkage for entities with true, external linkage.
	*/
	CXLinkage_External CXLinkageKind = 4
)

type CXVisibilityKind uint32

const (
	/*
	   This value indicates that no visibility information is available for a provided CXCursor.
	*/
	CXVisibility_Invalid CXVisibilityKind = 0
	/*
	   Symbol not seen by the linker.
	*/
	CXVisibility_Hidden CXVisibilityKind = 1
	/*
	   Symbol seen by the linker but resolves to a symbol inside this object.
	*/
	CXVisibility_Protected CXVisibilityKind = 2
	/*
	   Symbol seen by the linker and acts like a normal symbol.
	*/
	CXVisibility_Default CXVisibilityKind = 3
)

/*
Describe the "language" of the entity referred to by a cursor.
*/
type CXLanguageKind uint32

const (
	CXLanguage_Invalid   CXLanguageKind = 0
	CXLanguage_C         CXLanguageKind = 1
	CXLanguage_ObjC      CXLanguageKind = 2
	CXLanguage_CPlusPlus CXLanguageKind = 3
)

/*
Describe the "thread-local storage (TLS) kind" of the declaration referred to by a cursor.
*/
type CXTLSKind uint32

const (
	CXTLS_None    CXTLSKind = 0
	CXTLS_Dynamic CXTLSKind = 1
	CXTLS_Static  CXTLSKind = 2
)

/*
Describes the kind of type
*/
type CXTypeKind uint32

const (
	/*
	   Represents an invalid type (e.g., where no type is available).
	*/
	CXType_Invalid CXTypeKind = 0
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Unexposed CXTypeKind = 1
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Void CXTypeKind = 2
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Bool CXTypeKind = 3
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Char_U CXTypeKind = 4
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_UChar CXTypeKind = 5
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Char16 CXTypeKind = 6
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Char32 CXTypeKind = 7
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_UShort CXTypeKind = 8
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_UInt CXTypeKind = 9
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ULong CXTypeKind = 10
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ULongLong CXTypeKind = 11
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_UInt128 CXTypeKind = 12
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Char_S CXTypeKind = 13
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_SChar CXTypeKind = 14
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_WChar CXTypeKind = 15
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Short CXTypeKind = 16
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Int CXTypeKind = 17
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Long CXTypeKind = 18
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_LongLong CXTypeKind = 19
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Int128 CXTypeKind = 20
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Float CXTypeKind = 21
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Double CXTypeKind = 22
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_LongDouble CXTypeKind = 23
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_NullPtr CXTypeKind = 24
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Overload CXTypeKind = 25
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Dependent CXTypeKind = 26
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ObjCId CXTypeKind = 27
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ObjCClass CXTypeKind = 28
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ObjCSel CXTypeKind = 29
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Float128 CXTypeKind = 30
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Half CXTypeKind = 31
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Float16 CXTypeKind = 32
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ShortAccum CXTypeKind = 33
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Accum CXTypeKind = 34
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_LongAccum CXTypeKind = 35
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_UShortAccum CXTypeKind = 36
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_UAccum CXTypeKind = 37
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ULongAccum CXTypeKind = 38
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_BFloat16 CXTypeKind = 39
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Ibm128 CXTypeKind = 40
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_FirstBuiltin CXTypeKind = 2
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_LastBuiltin CXTypeKind = 40
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Complex CXTypeKind = 100
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Pointer CXTypeKind = 101
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_BlockPointer CXTypeKind = 102
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_LValueReference CXTypeKind = 103
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_RValueReference CXTypeKind = 104
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Record CXTypeKind = 105
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Enum CXTypeKind = 106
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Typedef CXTypeKind = 107
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ObjCInterface CXTypeKind = 108
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ObjCObjectPointer CXTypeKind = 109
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_FunctionNoProto CXTypeKind = 110
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_FunctionProto CXTypeKind = 111
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_ConstantArray CXTypeKind = 112
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Vector CXTypeKind = 113
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_IncompleteArray CXTypeKind = 114
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_VariableArray CXTypeKind = 115
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_DependentSizedArray CXTypeKind = 116
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_MemberPointer CXTypeKind = 117
	/*
	   A type whose specific kind is not exposed via this interface.
	*/
	CXType_Auto CXTypeKind = 118
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_Elaborated CXTypeKind = 119
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_Pipe CXTypeKind = 120
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dRO CXTypeKind = 121
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dArrayRO CXTypeKind = 122
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dBufferRO CXTypeKind = 123
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dRO CXTypeKind = 124
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayRO CXTypeKind = 125
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dDepthRO CXTypeKind = 126
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayDepthRO CXTypeKind = 127
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dMSAARO CXTypeKind = 128
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayMSAARO CXTypeKind = 129
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dMSAADepthRO CXTypeKind = 130
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayMSAADepthRO CXTypeKind = 131
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage3dRO CXTypeKind = 132
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dWO CXTypeKind = 133
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dArrayWO CXTypeKind = 134
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dBufferWO CXTypeKind = 135
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dWO CXTypeKind = 136
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayWO CXTypeKind = 137
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dDepthWO CXTypeKind = 138
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayDepthWO CXTypeKind = 139
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dMSAAWO CXTypeKind = 140
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayMSAAWO CXTypeKind = 141
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dMSAADepthWO CXTypeKind = 142
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayMSAADepthWO CXTypeKind = 143
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage3dWO CXTypeKind = 144
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dRW CXTypeKind = 145
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dArrayRW CXTypeKind = 146
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage1dBufferRW CXTypeKind = 147
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dRW CXTypeKind = 148
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayRW CXTypeKind = 149
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dDepthRW CXTypeKind = 150
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayDepthRW CXTypeKind = 151
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dMSAARW CXTypeKind = 152
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayMSAARW CXTypeKind = 153
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dMSAADepthRW CXTypeKind = 154
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage2dArrayMSAADepthRW CXTypeKind = 155
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLImage3dRW CXTypeKind = 156
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLSampler CXTypeKind = 157
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLEvent CXTypeKind = 158
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLQueue CXTypeKind = 159
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLReserveID CXTypeKind = 160
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_ObjCObject CXTypeKind = 161
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_ObjCTypeParam CXTypeKind = 162
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_Attributed CXTypeKind = 163
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCMcePayload CXTypeKind = 164
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImePayload CXTypeKind = 165
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCRefPayload CXTypeKind = 166
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCSicPayload CXTypeKind = 167
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCMceResult CXTypeKind = 168
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeResult CXTypeKind = 169
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCRefResult CXTypeKind = 170
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCSicResult CXTypeKind = 171
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeResultSingleReferenceStreamout CXTypeKind = 172
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeResultDualReferenceStreamout CXTypeKind = 173
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeSingleReferenceStreamin CXTypeKind = 174
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeDualReferenceStreamin CXTypeKind = 175
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeResultSingleRefStreamout CXTypeKind = 172
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeResultDualRefStreamout CXTypeKind = 173
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeSingleRefStreamin CXTypeKind = 174
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_OCLIntelSubgroupAVCImeDualRefStreamin CXTypeKind = 175
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_ExtVector CXTypeKind = 176
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_Atomic CXTypeKind = 177
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_BTFTagAttributed CXTypeKind = 178
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_HLSLResource CXTypeKind = 179
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_HLSLAttributedResource CXTypeKind = 180
	/*
	   Represents a type that was referred to using an elaborated type keyword.

	   E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	CXType_HLSLInlineSpirv CXTypeKind = 181
)

/*
Describes the calling convention of a function type
*/
type CXCallingConv uint32

const (
	CXCallingConv_Default            CXCallingConv = 0
	CXCallingConv_C                  CXCallingConv = 1
	CXCallingConv_X86StdCall         CXCallingConv = 2
	CXCallingConv_X86FastCall        CXCallingConv = 3
	CXCallingConv_X86ThisCall        CXCallingConv = 4
	CXCallingConv_X86Pascal          CXCallingConv = 5
	CXCallingConv_AAPCS              CXCallingConv = 6
	CXCallingConv_AAPCS_VFP          CXCallingConv = 7
	CXCallingConv_X86RegCall         CXCallingConv = 8
	CXCallingConv_IntelOclBicc       CXCallingConv = 9
	CXCallingConv_Win64              CXCallingConv = 10
	CXCallingConv_X86_64Win64        CXCallingConv = 10
	CXCallingConv_X86_64SysV         CXCallingConv = 11
	CXCallingConv_X86VectorCall      CXCallingConv = 12
	CXCallingConv_Swift              CXCallingConv = 13
	CXCallingConv_PreserveMost       CXCallingConv = 14
	CXCallingConv_PreserveAll        CXCallingConv = 15
	CXCallingConv_AArch64VectorCall  CXCallingConv = 16
	CXCallingConv_SwiftAsync         CXCallingConv = 17
	CXCallingConv_AArch64SVEPCS      CXCallingConv = 18
	CXCallingConv_M68kRTD            CXCallingConv = 19
	CXCallingConv_PreserveNone       CXCallingConv = 20
	CXCallingConv_RISCVVectorCall    CXCallingConv = 21
	CXCallingConv_RISCVVLSCall_32    CXCallingConv = 22
	CXCallingConv_RISCVVLSCall_64    CXCallingConv = 23
	CXCallingConv_RISCVVLSCall_128   CXCallingConv = 24
	CXCallingConv_RISCVVLSCall_256   CXCallingConv = 25
	CXCallingConv_RISCVVLSCall_512   CXCallingConv = 26
	CXCallingConv_RISCVVLSCall_1024  CXCallingConv = 27
	CXCallingConv_RISCVVLSCall_2048  CXCallingConv = 28
	CXCallingConv_RISCVVLSCall_4096  CXCallingConv = 29
	CXCallingConv_RISCVVLSCall_8192  CXCallingConv = 30
	CXCallingConv_RISCVVLSCall_16384 CXCallingConv = 31
	CXCallingConv_RISCVVLSCall_32768 CXCallingConv = 32
	CXCallingConv_RISCVVLSCall_65536 CXCallingConv = 33
	CXCallingConv_Invalid            CXCallingConv = 100
	CXCallingConv_Unexposed          CXCallingConv = 200
)

/*
Describes the kind of a template argument.

See the definition of llvm::clang::TemplateArgument::ArgKind for full element descriptions.
*/
type CXTemplateArgumentKind uint32

const (
	CXTemplateArgumentKind_Null              CXTemplateArgumentKind = 0
	CXTemplateArgumentKind_Type              CXTemplateArgumentKind = 1
	CXTemplateArgumentKind_Declaration       CXTemplateArgumentKind = 2
	CXTemplateArgumentKind_NullPtr           CXTemplateArgumentKind = 3
	CXTemplateArgumentKind_Integral          CXTemplateArgumentKind = 4
	CXTemplateArgumentKind_Template          CXTemplateArgumentKind = 5
	CXTemplateArgumentKind_TemplateExpansion CXTemplateArgumentKind = 6
	CXTemplateArgumentKind_Expression        CXTemplateArgumentKind = 7
	CXTemplateArgumentKind_Pack              CXTemplateArgumentKind = 8
	CXTemplateArgumentKind_Invalid           CXTemplateArgumentKind = 9
)

type CXTypeNullabilityKind uint32

const (
	/*
	   Values of this type can never be null.
	*/
	CXTypeNullability_NonNull CXTypeNullabilityKind = 0
	/*
	   Values of this type can be null.
	*/
	CXTypeNullability_Nullable CXTypeNullabilityKind = 1
	/*
	   Whether values of this type can be null is (explicitly) unspecified. This captures a (fairly rare) case where we can't conclude anything about the nullability of the type even though it has been considered.
	*/
	CXTypeNullability_Unspecified CXTypeNullabilityKind = 2
	/*
	   Nullability is not applicable to this type.
	*/
	CXTypeNullability_Invalid CXTypeNullabilityKind = 3
	/*
	   Generally behaves like Nullable, except when used in a block parameter that was imported into a swift async method. There, swift will assume that the parameter can get null even if no error occurred. _Nullable parameters are assumed to only get null on error.
	*/
	CXTypeNullability_NullableResult CXTypeNullabilityKind = 4
)

/*
List the possible error codes for clang_Type_getSizeOf,   clang_Type_getAlignOf, clang_Type_getOffsetOf,   clang_Cursor_getOffsetOf, and clang_getOffsetOfBase.

A value of this enumeration type can be returned if the target type is not a valid argument to sizeof, alignof or offsetof.
*/
type CXTypeLayoutError int32

const (
	/*
	   Type is of kind CXType_Invalid.
	*/
	CXTypeLayoutError_Invalid CXTypeLayoutError = -1
	/*
	   The type is an incomplete Type.
	*/
	CXTypeLayoutError_Incomplete CXTypeLayoutError = -2
	/*
	   The type is a dependent Type.
	*/
	CXTypeLayoutError_Dependent CXTypeLayoutError = -3
	/*
	   The type is not a constant size type.
	*/
	CXTypeLayoutError_NotConstantSize CXTypeLayoutError = -4
	/*
	   The Field name is not valid for this record.
	*/
	CXTypeLayoutError_InvalidFieldName CXTypeLayoutError = -5
	/*
	   The type is undeduced.
	*/
	CXTypeLayoutError_Undeduced CXTypeLayoutError = -6
)

type CXRefQualifierKind uint32

const (
	/*
	   No ref-qualifier was provided.
	*/
	CXRefQualifier_None CXRefQualifierKind = 0
	/*
	   An lvalue ref-qualifier was provided (&).
	*/
	CXRefQualifier_LValue CXRefQualifierKind = 1
	/*
	   An rvalue ref-qualifier was provided (&&).
	*/
	CXRefQualifier_RValue CXRefQualifierKind = 2
)

/*
Represents the C++ access control level to a base class for a cursor with kind CX_CXXBaseSpecifier.
*/
type CX_CXXAccessSpecifier uint32

const (
	CX_CXXInvalidAccessSpecifier CX_CXXAccessSpecifier = 0
	CX_CXXPublic                 CX_CXXAccessSpecifier = 1
	CX_CXXProtected              CX_CXXAccessSpecifier = 2
	CX_CXXPrivate                CX_CXXAccessSpecifier = 3
)

/*
Represents the storage classes as declared in the source. CX_SC_Invalid was added for the case that the passed cursor in not a declaration.
*/
type CX_StorageClass uint32

const (
	CX_SC_Invalid              CX_StorageClass = 0
	CX_SC_None                 CX_StorageClass = 1
	CX_SC_Extern               CX_StorageClass = 2
	CX_SC_Static               CX_StorageClass = 3
	CX_SC_PrivateExtern        CX_StorageClass = 4
	CX_SC_OpenCLWorkGroupLocal CX_StorageClass = 5
	CX_SC_Auto                 CX_StorageClass = 6
	CX_SC_Register             CX_StorageClass = 7
)

/*
Represents a specific kind of binary operator which can appear at a cursor.
*/
type CX_BinaryOperatorKind uint32

const (
	CX_BO_Invalid   CX_BinaryOperatorKind = 0
	CX_BO_PtrMemD   CX_BinaryOperatorKind = 1
	CX_BO_PtrMemI   CX_BinaryOperatorKind = 2
	CX_BO_Mul       CX_BinaryOperatorKind = 3
	CX_BO_Div       CX_BinaryOperatorKind = 4
	CX_BO_Rem       CX_BinaryOperatorKind = 5
	CX_BO_Add       CX_BinaryOperatorKind = 6
	CX_BO_Sub       CX_BinaryOperatorKind = 7
	CX_BO_Shl       CX_BinaryOperatorKind = 8
	CX_BO_Shr       CX_BinaryOperatorKind = 9
	CX_BO_Cmp       CX_BinaryOperatorKind = 10
	CX_BO_LT        CX_BinaryOperatorKind = 11
	CX_BO_GT        CX_BinaryOperatorKind = 12
	CX_BO_LE        CX_BinaryOperatorKind = 13
	CX_BO_GE        CX_BinaryOperatorKind = 14
	CX_BO_EQ        CX_BinaryOperatorKind = 15
	CX_BO_NE        CX_BinaryOperatorKind = 16
	CX_BO_And       CX_BinaryOperatorKind = 17
	CX_BO_Xor       CX_BinaryOperatorKind = 18
	CX_BO_Or        CX_BinaryOperatorKind = 19
	CX_BO_LAnd      CX_BinaryOperatorKind = 20
	CX_BO_LOr       CX_BinaryOperatorKind = 21
	CX_BO_Assign    CX_BinaryOperatorKind = 22
	CX_BO_MulAssign CX_BinaryOperatorKind = 23
	CX_BO_DivAssign CX_BinaryOperatorKind = 24
	CX_BO_RemAssign CX_BinaryOperatorKind = 25
	CX_BO_AddAssign CX_BinaryOperatorKind = 26
	CX_BO_SubAssign CX_BinaryOperatorKind = 27
	CX_BO_ShlAssign CX_BinaryOperatorKind = 28
	CX_BO_ShrAssign CX_BinaryOperatorKind = 29
	CX_BO_AndAssign CX_BinaryOperatorKind = 30
	CX_BO_XorAssign CX_BinaryOperatorKind = 31
	CX_BO_OrAssign  CX_BinaryOperatorKind = 32
	CX_BO_Comma     CX_BinaryOperatorKind = 33
	CX_BO_LAST      CX_BinaryOperatorKind = 33
)

/*
Describes how the traversal of the children of a particular cursor should proceed after visiting a particular child cursor.

A value of this enumeration type should be returned by each CXCursorVisitor to indicate how clang_visitChildren() proceed.
*/
type CXChildVisitResult uint32

const (
	/*
	   Terminates the cursor traversal.
	*/
	CXChildVisit_Break CXChildVisitResult = 0
	/*
	   Continues the cursor traversal with the next sibling of the cursor just visited, without visiting its children.
	*/
	CXChildVisit_Continue CXChildVisitResult = 1
	/*
	   Recursively traverse the children of this cursor, using the same visitor and client data.
	*/
	CXChildVisit_Recurse CXChildVisitResult = 2
)

/*
Properties for the printing policy.

See clang::PrintingPolicy for more information.
*/
type CXPrintingPolicyProperty uint32

const (
	CXPrintingPolicy_Indentation                           CXPrintingPolicyProperty = 0
	CXPrintingPolicy_SuppressSpecifiers                    CXPrintingPolicyProperty = 1
	CXPrintingPolicy_SuppressTagKeyword                    CXPrintingPolicyProperty = 2
	CXPrintingPolicy_IncludeTagDefinition                  CXPrintingPolicyProperty = 3
	CXPrintingPolicy_SuppressScope                         CXPrintingPolicyProperty = 4
	CXPrintingPolicy_SuppressUnwrittenScope                CXPrintingPolicyProperty = 5
	CXPrintingPolicy_SuppressInitializers                  CXPrintingPolicyProperty = 6
	CXPrintingPolicy_ConstantArraySizeAsWritten            CXPrintingPolicyProperty = 7
	CXPrintingPolicy_AnonymousTagLocations                 CXPrintingPolicyProperty = 8
	CXPrintingPolicy_SuppressStrongLifetime                CXPrintingPolicyProperty = 9
	CXPrintingPolicy_SuppressLifetimeQualifiers            CXPrintingPolicyProperty = 10
	CXPrintingPolicy_SuppressTemplateArgsInCXXConstructors CXPrintingPolicyProperty = 11
	CXPrintingPolicy_Bool                                  CXPrintingPolicyProperty = 12
	CXPrintingPolicy_Restrict                              CXPrintingPolicyProperty = 13
	CXPrintingPolicy_Alignof                               CXPrintingPolicyProperty = 14
	CXPrintingPolicy_UnderscoreAlignof                     CXPrintingPolicyProperty = 15
	CXPrintingPolicy_UseVoidForZeroParams                  CXPrintingPolicyProperty = 16
	CXPrintingPolicy_TerseOutput                           CXPrintingPolicyProperty = 17
	CXPrintingPolicy_PolishForDeclaration                  CXPrintingPolicyProperty = 18
	CXPrintingPolicy_Half                                  CXPrintingPolicyProperty = 19
	CXPrintingPolicy_MSWChar                               CXPrintingPolicyProperty = 20
	CXPrintingPolicy_IncludeNewlines                       CXPrintingPolicyProperty = 21
	CXPrintingPolicy_MSVCFormatting                        CXPrintingPolicyProperty = 22
	CXPrintingPolicy_ConstantsAsWritten                    CXPrintingPolicyProperty = 23
	CXPrintingPolicy_SuppressImplicitBase                  CXPrintingPolicyProperty = 24
	CXPrintingPolicy_FullyQualifiedName                    CXPrintingPolicyProperty = 25
	CXPrintingPolicy_LastProperty                          CXPrintingPolicyProperty = 25
)

/*
Property attributes for a CXCursor_ObjCPropertyDecl.
*/
type CXObjCPropertyAttrKind uint32

const (
	CXObjCPropertyAttr_noattr            CXObjCPropertyAttrKind = 0
	CXObjCPropertyAttr_readonly          CXObjCPropertyAttrKind = 1
	CXObjCPropertyAttr_getter            CXObjCPropertyAttrKind = 2
	CXObjCPropertyAttr_assign            CXObjCPropertyAttrKind = 4
	CXObjCPropertyAttr_readwrite         CXObjCPropertyAttrKind = 8
	CXObjCPropertyAttr_retain            CXObjCPropertyAttrKind = 16
	CXObjCPropertyAttr_copy              CXObjCPropertyAttrKind = 32
	CXObjCPropertyAttr_nonatomic         CXObjCPropertyAttrKind = 64
	CXObjCPropertyAttr_setter            CXObjCPropertyAttrKind = 128
	CXObjCPropertyAttr_atomic            CXObjCPropertyAttrKind = 256
	CXObjCPropertyAttr_weak              CXObjCPropertyAttrKind = 512
	CXObjCPropertyAttr_strong            CXObjCPropertyAttrKind = 1024
	CXObjCPropertyAttr_unsafe_unretained CXObjCPropertyAttrKind = 2048
	CXObjCPropertyAttr_class             CXObjCPropertyAttrKind = 4096
)

/*
'Qualifiers' written next to the return and parameter types in Objective-C method declarations.
*/
type CXObjCDeclQualifierKind uint32

const (
	CXObjCDeclQualifier_None   CXObjCDeclQualifierKind = 0
	CXObjCDeclQualifier_In     CXObjCDeclQualifierKind = 1
	CXObjCDeclQualifier_Inout  CXObjCDeclQualifierKind = 2
	CXObjCDeclQualifier_Out    CXObjCDeclQualifierKind = 4
	CXObjCDeclQualifier_Bycopy CXObjCDeclQualifierKind = 8
	CXObjCDeclQualifier_Byref  CXObjCDeclQualifierKind = 16
	CXObjCDeclQualifier_Oneway CXObjCDeclQualifierKind = 32
)

type CXNameRefFlags uint32

const (
	/*
	   Include the nested-name-specifier, e.g. Foo:: in x.Foo::y, in the range.
	*/
	CXNameRange_WantQualifier CXNameRefFlags = 1
	/*
	   Include the explicit template arguments, e.g. <int> in x.f<int>, in the range.
	*/
	CXNameRange_WantTemplateArgs CXNameRefFlags = 2
	/*
	   If the name is non-contiguous, return the full spanning range.

	   Non-contiguous names occur in Objective-C when a selector with two or more parameters is used, or in C++ when using an operator:
	*/
	CXNameRange_WantSinglePiece CXNameRefFlags = 4
)

/*
Describes a kind of token.
*/
type CXTokenKind uint32

const (
	/*
	   A token that contains some kind of punctuation.
	*/
	CXToken_Punctuation CXTokenKind = 0
	/*
	   A language keyword.
	*/
	CXToken_Keyword CXTokenKind = 1
	/*
	   An identifier (that is not a keyword).
	*/
	CXToken_Identifier CXTokenKind = 2
	/*
	   A numeric, string, or character literal.
	*/
	CXToken_Literal CXTokenKind = 3
	/*
	   A comment.
	*/
	CXToken_Comment CXTokenKind = 4
)

/*
Describes a single piece of text within a code-completion string.

Each "chunk" within a code-completion string (CXCompletionString) is either a piece of text with a specific "kind" that describes how that text should be interpreted by the client or is another completion string.
*/
type CXCompletionChunkKind uint32

const (
	/*
	   A code-completion string that describes "optional" text that could be a part of the template (but is not required).

	   The Optional chunk is the only kind of chunk that has a code-completion string for its representation, which is accessible via clang_getCompletionChunkCompletionString(). The code-completion string describes an additional part of the template that is completely optional. For example, optional chunks can be used to describe the placeholders for arguments that match up with defaulted function parameters, e.g. given:

	   The code-completion string for this function would contain:   - a TypedText chunk for "f".   - a LeftParen chunk for "(".   - a Placeholder chunk for "int x"   - an Optional chunk containing the remaining defaulted arguments, e.g.,       - a Comma chunk for ","       - a Placeholder chunk for "float y"       - an Optional chunk containing the last defaulted argument:           - a Comma chunk for ","           - a Placeholder chunk for "double z"   - a RightParen chunk for ")"

	   There are many ways to handle Optional chunks. Two simple approaches are:   - Completely ignore optional chunks, in which case the template for the     function "f" would only include the first parameter ("int x").   - Fully expand all optional chunks, in which case the template for the     function "f" would have all of the parameters.
	*/
	CXCompletionChunk_Optional CXCompletionChunkKind = 0
	/*
	   Text that a user would be expected to type to get this code-completion result.

	   There will be exactly one "typed text" chunk in a semantic string, which will typically provide the spelling of a keyword or the name of a declaration that could be used at the current code point. Clients are expected to filter the code-completion results based on the text in this chunk.
	*/
	CXCompletionChunk_TypedText CXCompletionChunkKind = 1
	/*
	   Text that should be inserted as part of a code-completion result.

	   A "text" chunk represents text that is part of the template to be inserted into user code should this particular code-completion result be selected.
	*/
	CXCompletionChunk_Text CXCompletionChunkKind = 2
	/*
	   Placeholder text that should be replaced by the user.

	   A "placeholder" chunk marks a place where the user should insert text into the code-completion template. For example, placeholders might mark the function parameters for a function declaration, to indicate that the user should provide arguments for each of those parameters. The actual text in a placeholder is a suggestion for the text to display before the user replaces the placeholder with real code.
	*/
	CXCompletionChunk_Placeholder CXCompletionChunkKind = 3
	/*
	   Informative text that should be displayed but never inserted as part of the template.

	   An "informative" chunk contains annotations that can be displayed to help the user decide whether a particular code-completion result is the right option, but which is not part of the actual template to be inserted by code completion.
	*/
	CXCompletionChunk_Informative CXCompletionChunkKind = 4
	/*
	   Text that describes the current parameter when code-completion is referring to function call, message send, or template specialization.

	   A "current parameter" chunk occurs when code-completion is providing information about a parameter corresponding to the argument at the code-completion point. For example, given a function

	   and the source code add(, where the code-completion point is after the "(", the code-completion string will contain a "current parameter" chunk for "int x", indicating that the current argument will initialize that parameter. After typing further, to add(17, (where the code-completion point is after the ","), the code-completion string will contain a "current parameter" chunk to "int y".
	*/
	CXCompletionChunk_CurrentParameter CXCompletionChunkKind = 5
	/*
	   A left parenthesis ('('), used to initiate a function call or signal the beginning of a function parameter list.
	*/
	CXCompletionChunk_LeftParen CXCompletionChunkKind = 6
	/*
	   A right parenthesis (')'), used to finish a function call or signal the end of a function parameter list.
	*/
	CXCompletionChunk_RightParen CXCompletionChunkKind = 7
	/*
	   A left bracket ('[').
	*/
	CXCompletionChunk_LeftBracket CXCompletionChunkKind = 8
	/*
	   A right bracket (']').
	*/
	CXCompletionChunk_RightBracket CXCompletionChunkKind = 9
	/*
	   A left brace ('{').
	*/
	CXCompletionChunk_LeftBrace CXCompletionChunkKind = 10
	/*
	   A right brace ('}').
	*/
	CXCompletionChunk_RightBrace CXCompletionChunkKind = 11
	/*
	   A left angle bracket ('<').
	*/
	CXCompletionChunk_LeftAngle CXCompletionChunkKind = 12
	/*
	   A right angle bracket ('>').
	*/
	CXCompletionChunk_RightAngle CXCompletionChunkKind = 13
	/*
	   A comma separator (',').
	*/
	CXCompletionChunk_Comma CXCompletionChunkKind = 14
	/*
	   Text that specifies the result type of a given result.

	   This special kind of informative chunk is not meant to be inserted into the text buffer. Rather, it is meant to illustrate the type that an expression using the given completion string would have.
	*/
	CXCompletionChunk_ResultType CXCompletionChunkKind = 15
	/*
	   A colon (':').
	*/
	CXCompletionChunk_Colon CXCompletionChunkKind = 16
	/*
	   A semicolon (';').
	*/
	CXCompletionChunk_SemiColon CXCompletionChunkKind = 17
	/*
	   An '=' sign.
	*/
	CXCompletionChunk_Equal CXCompletionChunkKind = 18
	/*
	   Horizontal space (' ').
	*/
	CXCompletionChunk_HorizontalSpace CXCompletionChunkKind = 19
	/*
	   Vertical space ('\n'), after which it is generally a good idea to perform indentation.
	*/
	CXCompletionChunk_VerticalSpace CXCompletionChunkKind = 20
)

/*
Flags that can be passed to clang_codeCompleteAt() to modify its behavior.

The enumerators in this enumeration can be bitwise-OR'd together to provide multiple options to clang_codeCompleteAt().
*/
type CXCodeComplete_Flags uint32

const (
	/*
	   Whether to include macros within the set of code completions returned.
	*/
	CXCodeComplete_IncludeMacros CXCodeComplete_Flags = 1
	/*
	   Whether to include code patterns for language constructs within the set of code completions, e.g., for loops.
	*/
	CXCodeComplete_IncludeCodePatterns CXCodeComplete_Flags = 2
	/*
	   Whether to include brief documentation within the set of code completions returned.
	*/
	CXCodeComplete_IncludeBriefComments CXCodeComplete_Flags = 4
	/*
	   Whether to speed up completion by omitting top- or namespace-level entities defined in the preamble. There's no guarantee any particular entity is omitted. This may be useful if the headers are indexed externally.
	*/
	CXCodeComplete_SkipPreamble CXCodeComplete_Flags = 8
	/*
	   Whether to include completions with small fix-its, e.g. change '.' to '->' on member access, etc.
	*/
	CXCodeComplete_IncludeCompletionsWithFixIts CXCodeComplete_Flags = 16
)

/*
Bits that represent the context under which completion is occurring.

The enumerators in this enumeration may be bitwise-OR'd together if multiple contexts are occurring simultaneously.
*/
type CXCompletionContext uint32

const (
	/*
	   The context for completions is unexposed, as only Clang results should be included. (This is equivalent to having no context bits set.)
	*/
	CXCompletionContext_Unexposed CXCompletionContext = 0
	/*
	   Completions for any possible type should be included in the results.
	*/
	CXCompletionContext_AnyType CXCompletionContext = 1
	/*
	   Completions for any possible value (variables, function calls, etc.) should be included in the results.
	*/
	CXCompletionContext_AnyValue CXCompletionContext = 2
	/*
	   Completions for values that resolve to an Objective-C object should be included in the results.
	*/
	CXCompletionContext_ObjCObjectValue CXCompletionContext = 4
	/*
	   Completions for values that resolve to an Objective-C selector should be included in the results.
	*/
	CXCompletionContext_ObjCSelectorValue CXCompletionContext = 8
	/*
	   Completions for values that resolve to a C++ class type should be included in the results.
	*/
	CXCompletionContext_CXXClassTypeValue CXCompletionContext = 16
	/*
	   Completions for fields of the member being accessed using the dot operator should be included in the results.
	*/
	CXCompletionContext_DotMemberAccess CXCompletionContext = 32
	/*
	   Completions for fields of the member being accessed using the arrow operator should be included in the results.
	*/
	CXCompletionContext_ArrowMemberAccess CXCompletionContext = 64
	/*
	   Completions for properties of the Objective-C object being accessed using the dot operator should be included in the results.
	*/
	CXCompletionContext_ObjCPropertyAccess CXCompletionContext = 128
	/*
	   Completions for enum tags should be included in the results.
	*/
	CXCompletionContext_EnumTag CXCompletionContext = 256
	/*
	   Completions for union tags should be included in the results.
	*/
	CXCompletionContext_UnionTag CXCompletionContext = 512
	/*
	   Completions for struct tags should be included in the results.
	*/
	CXCompletionContext_StructTag CXCompletionContext = 1024
	/*
	   Completions for C++ class names should be included in the results.
	*/
	CXCompletionContext_ClassTag CXCompletionContext = 2048
	/*
	   Completions for C++ namespaces and namespace aliases should be included in the results.
	*/
	CXCompletionContext_Namespace CXCompletionContext = 4096
	/*
	   Completions for C++ nested name specifiers should be included in the results.
	*/
	CXCompletionContext_NestedNameSpecifier CXCompletionContext = 8192
	/*
	   Completions for Objective-C interfaces (classes) should be included in the results.
	*/
	CXCompletionContext_ObjCInterface CXCompletionContext = 16384
	/*
	   Completions for Objective-C protocols should be included in the results.
	*/
	CXCompletionContext_ObjCProtocol CXCompletionContext = 32768
	/*
	   Completions for Objective-C categories should be included in the results.
	*/
	CXCompletionContext_ObjCCategory CXCompletionContext = 65536
	/*
	   Completions for Objective-C instance messages should be included in the results.
	*/
	CXCompletionContext_ObjCInstanceMessage CXCompletionContext = 131072
	/*
	   Completions for Objective-C class messages should be included in the results.
	*/
	CXCompletionContext_ObjCClassMessage CXCompletionContext = 262144
	/*
	   Completions for Objective-C selector names should be included in the results.
	*/
	CXCompletionContext_ObjCSelectorName CXCompletionContext = 524288
	/*
	   Completions for preprocessor macro names should be included in the results.
	*/
	CXCompletionContext_MacroName CXCompletionContext = 1048576
	/*
	   Natural language completions should be included in the results.
	*/
	CXCompletionContext_NaturalLanguage CXCompletionContext = 2097152
	/*
	   #include file completions should be included in the results.
	*/
	CXCompletionContext_IncludedFile CXCompletionContext = 4194304
	/*
	   The current context is unknown, so set all contexts.
	*/
	CXCompletionContext_Unknown CXCompletionContext = 8388607
)

type CXEvalResultKind uint32

const (
	CXEval_Int            CXEvalResultKind = 1
	CXEval_Float          CXEvalResultKind = 2
	CXEval_ObjCStrLiteral CXEvalResultKind = 3
	CXEval_StrLiteral     CXEvalResultKind = 4
	CXEval_CFStr          CXEvalResultKind = 5
	CXEval_Other          CXEvalResultKind = 6
	CXEval_UnExposed      CXEvalResultKind = 0
)

/*
@{
*/
type CXVisitorResult uint32

const (
	CXVisit_Break    CXVisitorResult = 0
	CXVisit_Continue CXVisitorResult = 1
)

type CXResult uint32

const (
	/*
	   Function returned successfully.
	*/
	CXResult_Success CXResult = 0
	/*
	   One of the parameters was invalid for the function.
	*/
	CXResult_Invalid CXResult = 1
	/*
	   The function was terminated by a callback (e.g. it returned CXVisit_Break)
	*/
	CXResult_VisitBreak CXResult = 2
)

type CXIdxEntityKind uint32

const (
	CXIdxEntity_Unexposed             CXIdxEntityKind = 0
	CXIdxEntity_Typedef               CXIdxEntityKind = 1
	CXIdxEntity_Function              CXIdxEntityKind = 2
	CXIdxEntity_Variable              CXIdxEntityKind = 3
	CXIdxEntity_Field                 CXIdxEntityKind = 4
	CXIdxEntity_EnumConstant          CXIdxEntityKind = 5
	CXIdxEntity_ObjCClass             CXIdxEntityKind = 6
	CXIdxEntity_ObjCProtocol          CXIdxEntityKind = 7
	CXIdxEntity_ObjCCategory          CXIdxEntityKind = 8
	CXIdxEntity_ObjCInstanceMethod    CXIdxEntityKind = 9
	CXIdxEntity_ObjCClassMethod       CXIdxEntityKind = 10
	CXIdxEntity_ObjCProperty          CXIdxEntityKind = 11
	CXIdxEntity_ObjCIvar              CXIdxEntityKind = 12
	CXIdxEntity_Enum                  CXIdxEntityKind = 13
	CXIdxEntity_Struct                CXIdxEntityKind = 14
	CXIdxEntity_Union                 CXIdxEntityKind = 15
	CXIdxEntity_CXXClass              CXIdxEntityKind = 16
	CXIdxEntity_CXXNamespace          CXIdxEntityKind = 17
	CXIdxEntity_CXXNamespaceAlias     CXIdxEntityKind = 18
	CXIdxEntity_CXXStaticVariable     CXIdxEntityKind = 19
	CXIdxEntity_CXXStaticMethod       CXIdxEntityKind = 20
	CXIdxEntity_CXXInstanceMethod     CXIdxEntityKind = 21
	CXIdxEntity_CXXConstructor        CXIdxEntityKind = 22
	CXIdxEntity_CXXDestructor         CXIdxEntityKind = 23
	CXIdxEntity_CXXConversionFunction CXIdxEntityKind = 24
	CXIdxEntity_CXXTypeAlias          CXIdxEntityKind = 25
	CXIdxEntity_CXXInterface          CXIdxEntityKind = 26
	CXIdxEntity_CXXConcept            CXIdxEntityKind = 27
)

type CXIdxEntityLanguage uint32

const (
	CXIdxEntityLang_None  CXIdxEntityLanguage = 0
	CXIdxEntityLang_C     CXIdxEntityLanguage = 1
	CXIdxEntityLang_ObjC  CXIdxEntityLanguage = 2
	CXIdxEntityLang_CXX   CXIdxEntityLanguage = 3
	CXIdxEntityLang_Swift CXIdxEntityLanguage = 4
)

/*
Extra C++ template information for an entity. This can apply to: CXIdxEntity_Function CXIdxEntity_CXXClass CXIdxEntity_CXXStaticMethod CXIdxEntity_CXXInstanceMethod CXIdxEntity_CXXConstructor CXIdxEntity_CXXConversionFunction CXIdxEntity_CXXTypeAlias
*/
type CXIdxEntityCXXTemplateKind uint32

const (
	CXIdxEntity_NonTemplate                   CXIdxEntityCXXTemplateKind = 0
	CXIdxEntity_Template                      CXIdxEntityCXXTemplateKind = 1
	CXIdxEntity_TemplatePartialSpecialization CXIdxEntityCXXTemplateKind = 2
	CXIdxEntity_TemplateSpecialization        CXIdxEntityCXXTemplateKind = 3
)

type CXIdxAttrKind uint32

const (
	CXIdxAttr_Unexposed          CXIdxAttrKind = 0
	CXIdxAttr_IBAction           CXIdxAttrKind = 1
	CXIdxAttr_IBOutlet           CXIdxAttrKind = 2
	CXIdxAttr_IBOutletCollection CXIdxAttrKind = 3
)

type CXIdxDeclInfoFlags uint32

const (
	CXIdxDeclFlag_Skipped CXIdxDeclInfoFlags = 1
)

type CXIdxObjCContainerKind uint32

const (
	CXIdxObjCContainer_ForwardRef     CXIdxObjCContainerKind = 0
	CXIdxObjCContainer_Interface      CXIdxObjCContainerKind = 1
	CXIdxObjCContainer_Implementation CXIdxObjCContainerKind = 2
)

/*
Data for IndexerCallbacks#indexEntityReference.

This may be deprecated in a future version as this duplicates the CXSymbolRole_Implicit bit in CXSymbolRole.
*/
type CXIdxEntityRefKind uint32

const (
	/*
	   The entity is referenced directly in user's code.
	*/
	CXIdxEntityRef_Direct CXIdxEntityRefKind = 1
	/*
	   An implicit reference, e.g. a reference of an Objective-C method via the dot syntax.
	*/
	CXIdxEntityRef_Implicit CXIdxEntityRefKind = 2
)

/*
Roles that are attributed to symbol occurrences.

Internal: this currently mirrors low 9 bits of clang::index::SymbolRole with higher bits zeroed. These high bits may be exposed in the future.
*/
type CXSymbolRole uint32

const (
	CXSymbolRole_None        CXSymbolRole = 0
	CXSymbolRole_Declaration CXSymbolRole = 1
	CXSymbolRole_Definition  CXSymbolRole = 2
	CXSymbolRole_Reference   CXSymbolRole = 4
	CXSymbolRole_Read        CXSymbolRole = 8
	CXSymbolRole_Write       CXSymbolRole = 16
	CXSymbolRole_Call        CXSymbolRole = 32
	CXSymbolRole_Dynamic     CXSymbolRole = 64
	CXSymbolRole_AddressOf   CXSymbolRole = 128
	CXSymbolRole_Implicit    CXSymbolRole = 256
)

type CXIndexOptFlags uint32

const (
	/*
	   Used to indicate that no special indexing options are needed.
	*/
	CXIndexOpt_None CXIndexOptFlags = 0
	/*
	   Used to indicate that IndexerCallbacks#indexEntityReference should be invoked for only one reference of an entity per source file that does not also include a declaration/definition of the entity.
	*/
	CXIndexOpt_SuppressRedundantRefs CXIndexOptFlags = 1
	/*
	   Function-local symbols should be indexed. If this is not set function-local symbols will be ignored.
	*/
	CXIndexOpt_IndexFunctionLocalSymbols CXIndexOptFlags = 2
	/*
	   Implicit function/class template instantiations should be indexed. If this is not set, implicit instantiations will be ignored.
	*/
	CXIndexOpt_IndexImplicitTemplateInstantiations CXIndexOptFlags = 4
	/*
	   Suppress all compiler warnings when parsing for indexing.
	*/
	CXIndexOpt_SuppressWarnings CXIndexOptFlags = 8
	/*
	   Skip a function/method body that was already parsed during an indexing session associated with a CXIndexAction object. Bodies in system headers are always skipped.
	*/
	CXIndexOpt_SkipParsedBodiesInSession CXIndexOptFlags = 16
)

/*
Describes the kind of binary operators.
*/
type CXBinaryOperatorKind uint32

const (
	/*
	   This value describes cursors which are not binary operators.
	*/
	CXBinaryOperator_Invalid CXBinaryOperatorKind = 0
	/*
	   C++ Pointer - to - member operator.
	*/
	CXBinaryOperator_PtrMemD CXBinaryOperatorKind = 1
	/*
	   C++ Pointer - to - member operator.
	*/
	CXBinaryOperator_PtrMemI CXBinaryOperatorKind = 2
	/*
	   Multiplication operator.
	*/
	CXBinaryOperator_Mul CXBinaryOperatorKind = 3
	/*
	   Division operator.
	*/
	CXBinaryOperator_Div CXBinaryOperatorKind = 4
	/*
	   Remainder operator.
	*/
	CXBinaryOperator_Rem CXBinaryOperatorKind = 5
	/*
	   Addition operator.
	*/
	CXBinaryOperator_Add CXBinaryOperatorKind = 6
	/*
	   Subtraction operator.
	*/
	CXBinaryOperator_Sub CXBinaryOperatorKind = 7
	/*
	   Bitwise shift left operator.
	*/
	CXBinaryOperator_Shl CXBinaryOperatorKind = 8
	/*
	   Bitwise shift right operator.
	*/
	CXBinaryOperator_Shr CXBinaryOperatorKind = 9
	/*
	   C++ three-way comparison (spaceship) operator.
	*/
	CXBinaryOperator_Cmp CXBinaryOperatorKind = 10
	/*
	   Less than operator.
	*/
	CXBinaryOperator_LT CXBinaryOperatorKind = 11
	/*
	   Greater than operator.
	*/
	CXBinaryOperator_GT CXBinaryOperatorKind = 12
	/*
	   Less or equal operator.
	*/
	CXBinaryOperator_LE CXBinaryOperatorKind = 13
	/*
	   Greater or equal operator.
	*/
	CXBinaryOperator_GE CXBinaryOperatorKind = 14
	/*
	   Equal operator.
	*/
	CXBinaryOperator_EQ CXBinaryOperatorKind = 15
	/*
	   Not equal operator.
	*/
	CXBinaryOperator_NE CXBinaryOperatorKind = 16
	/*
	   Bitwise AND operator.
	*/
	CXBinaryOperator_And CXBinaryOperatorKind = 17
	/*
	   Bitwise XOR operator.
	*/
	CXBinaryOperator_Xor CXBinaryOperatorKind = 18
	/*
	   Bitwise OR operator.
	*/
	CXBinaryOperator_Or CXBinaryOperatorKind = 19
	/*
	   Logical AND operator.
	*/
	CXBinaryOperator_LAnd CXBinaryOperatorKind = 20
	/*
	   Logical OR operator.
	*/
	CXBinaryOperator_LOr CXBinaryOperatorKind = 21
	/*
	   Assignment operator.
	*/
	CXBinaryOperator_Assign CXBinaryOperatorKind = 22
	/*
	   Multiplication assignment operator.
	*/
	CXBinaryOperator_MulAssign CXBinaryOperatorKind = 23
	/*
	   Division assignment operator.
	*/
	CXBinaryOperator_DivAssign CXBinaryOperatorKind = 24
	/*
	   Remainder assignment operator.
	*/
	CXBinaryOperator_RemAssign CXBinaryOperatorKind = 25
	/*
	   Addition assignment operator.
	*/
	CXBinaryOperator_AddAssign CXBinaryOperatorKind = 26
	/*
	   Subtraction assignment operator.
	*/
	CXBinaryOperator_SubAssign CXBinaryOperatorKind = 27
	/*
	   Bitwise shift left assignment operator.
	*/
	CXBinaryOperator_ShlAssign CXBinaryOperatorKind = 28
	/*
	   Bitwise shift right assignment operator.
	*/
	CXBinaryOperator_ShrAssign CXBinaryOperatorKind = 29
	/*
	   Bitwise AND assignment operator.
	*/
	CXBinaryOperator_AndAssign CXBinaryOperatorKind = 30
	/*
	   Bitwise XOR assignment operator.
	*/
	CXBinaryOperator_XorAssign CXBinaryOperatorKind = 31
	/*
	   Bitwise OR assignment operator.
	*/
	CXBinaryOperator_OrAssign CXBinaryOperatorKind = 32
	/*
	   Comma operator.
	*/
	CXBinaryOperator_Comma CXBinaryOperatorKind = 33
	/*
	   Comma operator.
	*/
	CXBinaryOperator_Last CXBinaryOperatorKind = 33
)

/*
Describes the kind of unary operators.
*/
type CXUnaryOperatorKind uint32

const (
	/*
	   This value describes cursors which are not unary operators.
	*/
	CXUnaryOperator_Invalid CXUnaryOperatorKind = 0
	/*
	   Postfix increment operator.
	*/
	CXUnaryOperator_PostInc CXUnaryOperatorKind = 1
	/*
	   Postfix decrement operator.
	*/
	CXUnaryOperator_PostDec CXUnaryOperatorKind = 2
	/*
	   Prefix increment operator.
	*/
	CXUnaryOperator_PreInc CXUnaryOperatorKind = 3
	/*
	   Prefix decrement operator.
	*/
	CXUnaryOperator_PreDec CXUnaryOperatorKind = 4
	/*
	   Address of operator.
	*/
	CXUnaryOperator_AddrOf CXUnaryOperatorKind = 5
	/*
	   Dereference operator.
	*/
	CXUnaryOperator_Deref CXUnaryOperatorKind = 6
	/*
	   Plus operator.
	*/
	CXUnaryOperator_Plus CXUnaryOperatorKind = 7
	/*
	   Minus operator.
	*/
	CXUnaryOperator_Minus CXUnaryOperatorKind = 8
	/*
	   Not operator.
	*/
	CXUnaryOperator_Not CXUnaryOperatorKind = 9
	/*
	   LNot operator.
	*/
	CXUnaryOperator_LNot CXUnaryOperatorKind = 10
	/*
	   "__real expr" operator.
	*/
	CXUnaryOperator_Real CXUnaryOperatorKind = 11
	/*
	   "__imag expr" operator.
	*/
	CXUnaryOperator_Imag CXUnaryOperatorKind = 12
	/*
	   __extension__ marker operator.
	*/
	CXUnaryOperator_Extension CXUnaryOperatorKind = 13
	/*
	   C++ co_await operator.
	*/
	CXUnaryOperator_Coawait CXUnaryOperatorKind = 14
	/*
	   C++ co_await operator.
	*/
	CXUnaryOperator_Last CXUnaryOperatorKind = 14
)
