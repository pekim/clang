// This is a generated file. DO NOT EDIT.

package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

// not supported : clang_getCString : return value : const char *

// not supported : clang_disposeString : param string : CXString

// not supported : clang_disposeStringSet : param set : CXStringSet *

// not supported : clang_getBuildSessionTimestamp : return value : unsigned long long

func VirtualFileOverlay_create(Options uint32) VirtualFileOverlay {
	c_Options := Options

	var retC VirtualFileOverlay
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_Options),
	}

	err := ffi.CallFunction(cif_clang_VirtualFileOverlay_create, ptr_clang_VirtualFileOverlay_create, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_create", err))
	}

	ret := retC
	return ret
}

// not supported : clang_VirtualFileOverlay_addFileMapping : return value : enum CXErrorCode

// not supported : clang_VirtualFileOverlay_setCaseSensitivity : return value : enum CXErrorCode

// not supported : clang_VirtualFileOverlay_writeToBuffer : return value : enum CXErrorCode

// not supported : clang_free : param buffer : void *

// not supported : clang_VirtualFileOverlay_dispose : param  : CXVirtualFileOverlay

func ModuleMapDescriptor_create(Options uint32) ModuleMapDescriptor {
	c_Options := Options

	var retC ModuleMapDescriptor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_Options),
	}

	err := ffi.CallFunction(cif_clang_ModuleMapDescriptor_create, ptr_clang_ModuleMapDescriptor_create, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_create", err))
	}

	ret := retC
	return ret
}

// not supported : clang_ModuleMapDescriptor_setFrameworkModuleName : return value : enum CXErrorCode

// not supported : clang_ModuleMapDescriptor_setUmbrellaHeader : return value : enum CXErrorCode

// not supported : clang_ModuleMapDescriptor_writeToBuffer : return value : enum CXErrorCode

// not supported : clang_ModuleMapDescriptor_dispose : param  : CXModuleMapDescriptor

// not supported : clang_getFileName : param SFile : CXFile

// not supported : clang_getFileTime : return value : time_t

// not supported : clang_getFileUniqueID : return value : int

// not supported : clang_File_isEqual : return value : int

// not supported : clang_File_tryGetRealPathName : param file : CXFile

func GetNullLocation() SourceLocation {
	var retC SourceLocation
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_getNullLocation, ptr_clang_getNullLocation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNullLocation", err))
	}

	ret := retC
	return ret
}

// not supported : clang_equalLocations : return value : unsigned int

// not supported : clang_Location_isInSystemHeader : return value : int

// not supported : clang_Location_isFromMainFile : return value : int

func GetNullRange() SourceRange {
	var retC SourceRange
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_getNullRange, ptr_clang_getNullRange, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNullRange", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getRange : param begin : CXSourceLocation

// not supported : clang_equalRanges : return value : unsigned int

// not supported : clang_Range_isNull : return value : int

// not supported : clang_getExpansionLocation : param location : CXSourceLocation

// not supported : clang_getPresumedLocation : param location : CXSourceLocation

// not supported : clang_getInstantiationLocation : param location : CXSourceLocation

// not supported : clang_getSpellingLocation : param location : CXSourceLocation

// not supported : clang_getFileLocation : param location : CXSourceLocation

// not supported : clang_getRangeStart : param range : CXSourceRange

// not supported : clang_getRangeEnd : param range : CXSourceRange

// not supported : clang_disposeSourceRangeList : param ranges : CXSourceRangeList *

// not supported : clang_getNumDiagnosticsInSet : return value : unsigned int

// not supported : clang_getDiagnosticInSet : param Diags : CXDiagnosticSet

// not supported : clang_loadDiagnostics : param file : const char *

// not supported : clang_disposeDiagnosticSet : param Diags : CXDiagnosticSet

// not supported : clang_getChildDiagnostics : param D : CXDiagnostic

// not supported : clang_disposeDiagnostic : param Diagnostic : CXDiagnostic

// not supported : clang_formatDiagnostic : param Diagnostic : CXDiagnostic

// not supported : clang_defaultDiagnosticDisplayOptions : return value : unsigned int

// not supported : clang_getDiagnosticSeverity : return value : enum CXDiagnosticSeverity

// not supported : clang_getDiagnosticLocation : param  : CXDiagnostic

// not supported : clang_getDiagnosticSpelling : param  : CXDiagnostic

// not supported : clang_getDiagnosticOption : param Diag : CXDiagnostic

// not supported : clang_getDiagnosticCategory : return value : unsigned int

func GetDiagnosticCategoryName(Category uint32) String_ {
	c_Category := Category

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_Category),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticCategoryName, ptr_clang_getDiagnosticCategoryName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategoryName", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getDiagnosticCategoryText : param  : CXDiagnostic

// not supported : clang_getDiagnosticNumRanges : return value : unsigned int

// not supported : clang_getDiagnosticRange : param Diagnostic : CXDiagnostic

// not supported : clang_getDiagnosticNumFixIts : return value : unsigned int

// not supported : clang_getDiagnosticFixIt : param Diagnostic : CXDiagnostic

/*
Provides a shared context for creating translation units.

It provides two options:

- excludeDeclarationsFromPCH: When non-zero, allows enumeration of "local" declarations (when loading any new translation units). A "local" declaration is one that belongs in the translation unit itself and not in a precompiled header that was used by the translation unit. If zero, all declarations will be enumerated.

Here is an example:

This process of creating the 'pch', loading it separately, and using it (via -include-pch) allows 'excludeDeclsFromPCH' to remove redundant callbacks (which gives the indexer the same performance benefit as the compiler).
*/
func CreateIndex(ExcludeDeclarationsFromPCH int32, DisplayDiagnostics int32) Index {
	c_ExcludeDeclarationsFromPCH := ExcludeDeclarationsFromPCH
	c_DisplayDiagnostics := DisplayDiagnostics

	var retC Index
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_ExcludeDeclarationsFromPCH),
		unsafe.Pointer(&c_DisplayDiagnostics),
	}

	err := ffi.CallFunction(cif_clang_createIndex, ptr_clang_createIndex, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createIndex", err))
	}

	ret := retC
	return ret
}

// not supported : clang_disposeIndex : param index : CXIndex

// not supported : clang_createIndexWithOptions : param options : const CXIndexOptions *

// not supported : clang_CXIndex_setGlobalOptions : param  : CXIndex

// not supported : clang_CXIndex_getGlobalOptions : return value : unsigned int

// not supported : clang_CXIndex_setInvocationEmissionPathOption : param  : CXIndex

// not supported : clang_isFileMultipleIncludeGuarded : return value : unsigned int

// not supported : clang_getFile : param tu : CXTranslationUnit

// not supported : clang_getFileContents : return value : const char *

// not supported : clang_getLocation : param tu : CXTranslationUnit

// not supported : clang_getLocationForOffset : param tu : CXTranslationUnit

// not supported : clang_getSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getAllSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getNumDiagnostics : return value : unsigned int

// not supported : clang_getDiagnostic : param Unit : CXTranslationUnit

// not supported : clang_getDiagnosticSetFromTU : param Unit : CXTranslationUnit

// not supported : clang_getTranslationUnitSpelling : param CTUnit : CXTranslationUnit

// not supported : clang_createTranslationUnitFromSourceFile : param CIdx : CXIndex

// not supported : clang_createTranslationUnit : param CIdx : CXIndex

// not supported : clang_createTranslationUnit2 : return value : enum CXErrorCode

// not supported : clang_defaultEditingTranslationUnitOptions : return value : unsigned int

// not supported : clang_parseTranslationUnit : param CIdx : CXIndex

// not supported : clang_parseTranslationUnit2 : return value : enum CXErrorCode

// not supported : clang_parseTranslationUnit2FullArgv : return value : enum CXErrorCode

// not supported : clang_defaultSaveOptions : return value : unsigned int

// not supported : clang_saveTranslationUnit : return value : int

// not supported : clang_suspendTranslationUnit : return value : unsigned int

// not supported : clang_disposeTranslationUnit : param  : CXTranslationUnit

// not supported : clang_defaultReparseOptions : return value : unsigned int

// not supported : clang_reparseTranslationUnit : return value : int

// not supported : clang_getTUResourceUsageName : return value : const char *

// not supported : clang_getCXTUResourceUsage : param TU : CXTranslationUnit

// not supported : clang_disposeCXTUResourceUsage : param usage : CXTUResourceUsage

// not supported : clang_getTranslationUnitTargetInfo : param CTUnit : CXTranslationUnit

// not supported : clang_TargetInfo_dispose : param Info : CXTargetInfo

// not supported : clang_TargetInfo_getTriple : param Info : CXTargetInfo

// not supported : clang_TargetInfo_getPointerWidth : return value : int

// Retrieve the NULL cursor, which represents no entity.
func GetNullCursor() Cursor {
	var retC Cursor
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_getNullCursor, ptr_clang_getNullCursor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNullCursor", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getTranslationUnitCursor : param  : CXTranslationUnit

// not supported : clang_equalCursors : return value : unsigned int

// not supported : clang_Cursor_isNull : return value : int

// not supported : clang_hashCursor : return value : unsigned int

// not supported : clang_getCursorKind : return value : enum CXCursorKind

// not supported : clang_isDeclaration : return value : unsigned int

// not supported : clang_isInvalidDeclaration : return value : unsigned int

// not supported : clang_isReference : return value : unsigned int

// not supported : clang_isExpression : return value : unsigned int

// not supported : clang_isStatement : return value : unsigned int

// not supported : clang_isAttribute : return value : unsigned int

// not supported : clang_Cursor_hasAttrs : return value : unsigned int

// not supported : clang_isInvalid : return value : unsigned int

// not supported : clang_isTranslationUnit : return value : unsigned int

// not supported : clang_isPreprocessing : return value : unsigned int

// not supported : clang_isUnexposed : return value : unsigned int

// not supported : clang_getCursorLinkage : return value : enum CXLinkageKind

// not supported : clang_getCursorVisibility : return value : enum CXVisibilityKind

// not supported : clang_getCursorAvailability : return value : enum CXAvailabilityKind

// not supported : clang_getCursorPlatformAvailability : return value : int

// not supported : clang_disposeCXPlatformAvailability : param availability : CXPlatformAvailability *

// not supported : clang_Cursor_getVarDeclInitializer : param cursor : CXCursor

// not supported : clang_Cursor_hasVarDeclGlobalStorage : return value : int

// not supported : clang_Cursor_hasVarDeclExternalStorage : return value : int

// not supported : clang_getCursorLanguage : return value : enum CXLanguageKind

// not supported : clang_getCursorTLSKind : return value : enum CXTLSKind

// not supported : clang_Cursor_getTranslationUnit : param  : CXCursor

// Creates an empty CXCursorSet.
func CreateCXCursorSet() CursorSet {
	var retC CursorSet
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_createCXCursorSet, ptr_clang_createCXCursorSet, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createCXCursorSet", err))
	}

	ret := retC
	return ret
}

// not supported : clang_disposeCXCursorSet : param cset : CXCursorSet

// not supported : clang_CXCursorSet_contains : return value : unsigned int

// not supported : clang_CXCursorSet_insert : return value : unsigned int

// not supported : clang_getCursorSemanticParent : param cursor : CXCursor

// not supported : clang_getCursorLexicalParent : param cursor : CXCursor

// not supported : clang_getOverriddenCursors : param cursor : CXCursor

// not supported : clang_disposeOverriddenCursors : param overridden : CXCursor *

// not supported : clang_getIncludedFile : param cursor : CXCursor

// not supported : clang_getCursor : param  : CXTranslationUnit

// not supported : clang_getCursorLocation : param  : CXCursor

// not supported : clang_getCursorExtent : param  : CXCursor

// not supported : clang_getCursorType : param C : CXCursor

// not supported : clang_getTypeSpelling : param CT : CXType

// not supported : clang_getTypedefDeclUnderlyingType : param C : CXCursor

// not supported : clang_getEnumDeclIntegerType : param C : CXCursor

// not supported : clang_getEnumConstantDeclValue : return value : long long

// not supported : clang_getEnumConstantDeclUnsignedValue : return value : unsigned long long

// not supported : clang_Cursor_isBitField : return value : unsigned int

// not supported : clang_getFieldDeclBitWidth : return value : int

// not supported : clang_Cursor_getNumArguments : return value : int

// not supported : clang_Cursor_getArgument : param C : CXCursor

// not supported : clang_Cursor_getNumTemplateArguments : return value : int

// not supported : clang_Cursor_getTemplateArgumentKind : return value : enum CXTemplateArgumentKind

// not supported : clang_Cursor_getTemplateArgumentType : param C : CXCursor

// not supported : clang_Cursor_getTemplateArgumentValue : return value : long long

// not supported : clang_Cursor_getTemplateArgumentUnsignedValue : return value : unsigned long long

// not supported : clang_equalTypes : return value : unsigned int

// not supported : clang_getCanonicalType : param T : CXType

// not supported : clang_isConstQualifiedType : return value : unsigned int

// not supported : clang_Cursor_isMacroFunctionLike : return value : unsigned int

// not supported : clang_Cursor_isMacroBuiltin : return value : unsigned int

// not supported : clang_Cursor_isFunctionInlined : return value : unsigned int

// not supported : clang_isVolatileQualifiedType : return value : unsigned int

// not supported : clang_isRestrictQualifiedType : return value : unsigned int

// not supported : clang_getAddressSpace : return value : unsigned int

// not supported : clang_getTypedefName : param CT : CXType

// not supported : clang_getPointeeType : param T : CXType

// not supported : clang_getUnqualifiedType : param CT : CXType

// not supported : clang_getNonReferenceType : param CT : CXType

// not supported : clang_getTypeDeclaration : param T : CXType

// not supported : clang_getDeclObjCTypeEncoding : param C : CXCursor

// not supported : clang_Type_getObjCEncoding : param type : CXType

// not supported : clang_getTypeKindSpelling : param K : enum CXTypeKind

// not supported : clang_getFunctionTypeCallingConv : return value : enum CXCallingConv

// not supported : clang_getResultType : param T : CXType

// not supported : clang_getExceptionSpecificationType : return value : int

// not supported : clang_getNumArgTypes : return value : int

// not supported : clang_getArgType : param T : CXType

// not supported : clang_Type_getObjCObjectBaseType : param T : CXType

// not supported : clang_Type_getNumObjCProtocolRefs : return value : unsigned int

// not supported : clang_Type_getObjCProtocolDecl : param T : CXType

// not supported : clang_Type_getNumObjCTypeArgs : return value : unsigned int

// not supported : clang_Type_getObjCTypeArg : param T : CXType

// not supported : clang_isFunctionTypeVariadic : return value : unsigned int

// not supported : clang_getCursorResultType : param C : CXCursor

// not supported : clang_getCursorExceptionSpecificationType : return value : int

// not supported : clang_isPODType : return value : unsigned int

// not supported : clang_getElementType : param T : CXType

// not supported : clang_getNumElements : return value : long long

// not supported : clang_getArrayElementType : param T : CXType

// not supported : clang_getArraySize : return value : long long

// not supported : clang_Type_getNamedType : param T : CXType

// not supported : clang_Type_isTransparentTagTypedef : return value : unsigned int

// not supported : clang_Type_getNullability : return value : enum CXTypeNullabilityKind

// not supported : clang_Type_getAlignOf : return value : long long

// not supported : clang_Type_getClassType : param T : CXType

// not supported : clang_Type_getSizeOf : return value : long long

// not supported : clang_Type_getOffsetOf : return value : long long

// not supported : clang_Type_getModifiedType : param T : CXType

// not supported : clang_Type_getValueType : param CT : CXType

// not supported : clang_Cursor_getOffsetOfField : return value : long long

// not supported : clang_Cursor_isAnonymous : return value : unsigned int

// not supported : clang_Cursor_isAnonymousRecordDecl : return value : unsigned int

// not supported : clang_Cursor_isInlineNamespace : return value : unsigned int

// not supported : clang_Type_getNumTemplateArguments : return value : int

// not supported : clang_Type_getTemplateArgumentAsType : param T : CXType

// not supported : clang_Type_getCXXRefQualifier : return value : enum CXRefQualifierKind

// not supported : clang_isVirtualBase : return value : unsigned int

// not supported : clang_getOffsetOfBase : return value : long long

// not supported : clang_getCXXAccessSpecifier : return value : enum CX_CXXAccessSpecifier

// not supported : clang_Cursor_getBinaryOpcode : return value : enum CX_BinaryOperatorKind

// not supported : clang_Cursor_getBinaryOpcodeStr : param Op : enum CX_BinaryOperatorKind

// not supported : clang_Cursor_getStorageClass : return value : enum CX_StorageClass

// not supported : clang_getNumOverloadedDecls : return value : unsigned int

// not supported : clang_getOverloadedDecl : param cursor : CXCursor

// not supported : clang_getIBOutletCollectionType : param  : CXCursor

// not supported : clang_visitChildren : return value : unsigned int

// not supported : clang_visitChildrenWithBlock : return value : unsigned int

// not supported : clang_getCursorUSR : param  : CXCursor

// not supported : clang_constructUSR_ObjCClass : param class_name : const char *

// not supported : clang_constructUSR_ObjCCategory : param class_name : const char *

// not supported : clang_constructUSR_ObjCProtocol : param protocol_name : const char *

// not supported : clang_constructUSR_ObjCIvar : param name : const char *

// not supported : clang_constructUSR_ObjCMethod : param name : const char *

// not supported : clang_constructUSR_ObjCProperty : param property : const char *

// not supported : clang_getCursorSpelling : param  : CXCursor

// not supported : clang_Cursor_getSpellingNameRange : param  : CXCursor

// not supported : clang_PrintingPolicy_getProperty : return value : unsigned int

// not supported : clang_PrintingPolicy_setProperty : param Policy : CXPrintingPolicy

// not supported : clang_getCursorPrintingPolicy : param  : CXCursor

// not supported : clang_PrintingPolicy_dispose : param Policy : CXPrintingPolicy

// not supported : clang_getCursorPrettyPrinted : param Cursor : CXCursor

// not supported : clang_getTypePrettyPrinted : param CT : CXType

// not supported : clang_getFullyQualifiedName : param CT : CXType

// not supported : clang_getCursorDisplayName : param  : CXCursor

// not supported : clang_getCursorReferenced : param  : CXCursor

// not supported : clang_getCursorDefinition : param  : CXCursor

// not supported : clang_isCursorDefinition : return value : unsigned int

// not supported : clang_getCanonicalCursor : param  : CXCursor

// not supported : clang_Cursor_getObjCSelectorIndex : return value : int

// not supported : clang_Cursor_isDynamicCall : return value : int

// not supported : clang_Cursor_getReceiverType : param C : CXCursor

// not supported : clang_Cursor_getObjCPropertyAttributes : return value : unsigned int

// not supported : clang_Cursor_getObjCPropertyGetterName : param C : CXCursor

// not supported : clang_Cursor_getObjCPropertySetterName : param C : CXCursor

// not supported : clang_Cursor_getObjCDeclQualifiers : return value : unsigned int

// not supported : clang_Cursor_isObjCOptional : return value : unsigned int

// not supported : clang_Cursor_isVariadic : return value : unsigned int

// not supported : clang_Cursor_isExternalSymbol : return value : unsigned int

// not supported : clang_Cursor_getCommentRange : param C : CXCursor

// not supported : clang_Cursor_getRawCommentText : param C : CXCursor

// not supported : clang_Cursor_getBriefCommentText : param C : CXCursor

// not supported : clang_Cursor_getMangling : param  : CXCursor

// not supported : clang_Cursor_getCXXManglings : return value : CXStringSet *

// not supported : clang_Cursor_getObjCManglings : return value : CXStringSet *

// not supported : clang_Cursor_getGCCAssemblyTemplate : param  : CXCursor

// not supported : clang_Cursor_isGCCAssemblyHasGoto : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumOutputs : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumInputs : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyInput : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyOutput : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumClobbers : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyClobber : param Cursor : CXCursor

// not supported : clang_Cursor_isGCCAssemblyVolatile : return value : unsigned int

// not supported : clang_Cursor_getModule : param C : CXCursor

// not supported : clang_getModuleForFile : param  : CXTranslationUnit

// not supported : clang_Module_getASTFile : param Module : CXModule

// not supported : clang_Module_getParent : param Module : CXModule

// not supported : clang_Module_getName : param Module : CXModule

// not supported : clang_Module_getFullName : param Module : CXModule

// not supported : clang_Module_isSystem : return value : int

// not supported : clang_Module_getNumTopLevelHeaders : return value : unsigned int

// not supported : clang_Module_getTopLevelHeader : param  : CXTranslationUnit

// not supported : clang_CXXConstructor_isConvertingConstructor : return value : unsigned int

// not supported : clang_CXXConstructor_isCopyConstructor : return value : unsigned int

// not supported : clang_CXXConstructor_isDefaultConstructor : return value : unsigned int

// not supported : clang_CXXConstructor_isMoveConstructor : return value : unsigned int

// not supported : clang_CXXField_isMutable : return value : unsigned int

// not supported : clang_CXXMethod_isDefaulted : return value : unsigned int

// not supported : clang_CXXMethod_isDeleted : return value : unsigned int

// not supported : clang_CXXMethod_isPureVirtual : return value : unsigned int

// not supported : clang_CXXMethod_isStatic : return value : unsigned int

// not supported : clang_CXXMethod_isVirtual : return value : unsigned int

// not supported : clang_CXXMethod_isCopyAssignmentOperator : return value : unsigned int

// not supported : clang_CXXMethod_isMoveAssignmentOperator : return value : unsigned int

// not supported : clang_CXXMethod_isExplicit : return value : unsigned int

// not supported : clang_CXXRecord_isAbstract : return value : unsigned int

// not supported : clang_EnumDecl_isScoped : return value : unsigned int

// not supported : clang_CXXMethod_isConst : return value : unsigned int

// not supported : clang_getTemplateCursorKind : return value : enum CXCursorKind

// not supported : clang_getSpecializedCursorTemplate : param C : CXCursor

// not supported : clang_getCursorReferenceNameRange : param C : CXCursor

// not supported : clang_getToken : return value : CXToken *

// not supported : clang_getTokenKind : return value : CXTokenKind

// not supported : clang_getTokenSpelling : param  : CXTranslationUnit

// not supported : clang_getTokenLocation : param  : CXTranslationUnit

// not supported : clang_getTokenExtent : param  : CXTranslationUnit

// not supported : clang_tokenize : param TU : CXTranslationUnit

// not supported : clang_annotateTokens : param TU : CXTranslationUnit

// not supported : clang_disposeTokens : param TU : CXTranslationUnit

// not supported : clang_getCursorKindSpelling : param Kind : enum CXCursorKind

// not supported : clang_getDefinitionSpellingAndExtent : param  : CXCursor

func EnableStackTraces() {
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_enableStackTraces, ptr_clang_enableStackTraces, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_enableStackTraces", err))
	}
}

// not supported : clang_executeOnThread : param fn : void (*)(void *)

// not supported : clang_getCompletionChunkKind : return value : enum CXCompletionChunkKind

// not supported : clang_getCompletionChunkText : param completion_string : CXCompletionString

// not supported : clang_getCompletionChunkCompletionString : param completion_string : CXCompletionString

// not supported : clang_getNumCompletionChunks : return value : unsigned int

// not supported : clang_getCompletionPriority : return value : unsigned int

// not supported : clang_getCompletionAvailability : return value : enum CXAvailabilityKind

// not supported : clang_getCompletionNumAnnotations : return value : unsigned int

// not supported : clang_getCompletionAnnotation : param completion_string : CXCompletionString

// not supported : clang_getCompletionParent : param completion_string : CXCompletionString

// not supported : clang_getCompletionBriefComment : param completion_string : CXCompletionString

// not supported : clang_getCursorCompletionString : param cursor : CXCursor

// not supported : clang_getCompletionNumFixIts : return value : unsigned int

// not supported : clang_getCompletionFixIt : param results : CXCodeCompleteResults *

// not supported : clang_defaultCodeCompleteOptions : return value : unsigned int

// not supported : clang_codeCompleteAt : return value : CXCodeCompleteResults *

// not supported : clang_sortCodeCompletionResults : param Results : CXCompletionResult *

// not supported : clang_disposeCodeCompleteResults : param Results : CXCodeCompleteResults *

// not supported : clang_codeCompleteGetNumDiagnostics : return value : unsigned int

// not supported : clang_codeCompleteGetDiagnostic : param Results : CXCodeCompleteResults *

// not supported : clang_codeCompleteGetContexts : return value : unsigned long long

// not supported : clang_codeCompleteGetContainerKind : return value : enum CXCursorKind

// not supported : clang_codeCompleteGetContainerUSR : param Results : CXCodeCompleteResults *

// not supported : clang_codeCompleteGetObjCSelector : param Results : CXCodeCompleteResults *

// Return a version string, suitable for showing to a user, but not        intended to be parsed (the format is not guaranteed to be stable).
func GetClangVersion() String_ {
	var retC String_
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_getClangVersion, ptr_clang_getClangVersion, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getClangVersion", err))
	}

	ret := retC
	return ret
}

// Enable/disable crash recovery.
func ToggleCrashRecovery(IsEnabled uint32) {
	c_IsEnabled := IsEnabled

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_IsEnabled),
	}

	err := ffi.CallFunction(cif_clang_toggleCrashRecovery, ptr_clang_toggleCrashRecovery, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_toggleCrashRecovery", err))
	}
}

// not supported : clang_getInclusions : param tu : CXTranslationUnit

// not supported : clang_Cursor_Evaluate : param C : CXCursor

// not supported : clang_EvalResult_getKind : return value : CXEvalResultKind

// not supported : clang_EvalResult_getAsInt : return value : int

// not supported : clang_EvalResult_getAsLongLong : return value : long long

// not supported : clang_EvalResult_isUnsignedInt : return value : unsigned int

// not supported : clang_EvalResult_getAsUnsigned : return value : unsigned long long

// not supported : clang_EvalResult_getAsDouble : return value : double

// not supported : clang_EvalResult_getAsStr : return value : const char *

// not supported : clang_EvalResult_dispose : param E : CXEvalResult

// not supported : clang_findReferencesInFile : return value : CXResult

// not supported : clang_findIncludesInFile : return value : CXResult

// not supported : clang_findReferencesInFileWithBlock : return value : CXResult

// not supported : clang_findIncludesInFileWithBlock : return value : CXResult

// not supported : clang_index_isEntityObjCContainerKind : return value : int

// not supported : clang_index_getObjCContainerDeclInfo : return value : const CXIdxObjCContainerDeclInfo *

// not supported : clang_index_getObjCInterfaceDeclInfo : return value : const CXIdxObjCInterfaceDeclInfo *

// not supported : clang_index_getObjCCategoryDeclInfo : return value : const CXIdxObjCCategoryDeclInfo *

// not supported : clang_index_getObjCProtocolRefListInfo : return value : const CXIdxObjCProtocolRefListInfo *

// not supported : clang_index_getObjCPropertyDeclInfo : return value : const CXIdxObjCPropertyDeclInfo *

// not supported : clang_index_getIBOutletCollectionAttrInfo : return value : const CXIdxIBOutletCollectionAttrInfo *

// not supported : clang_index_getCXXClassDeclInfo : return value : const CXIdxCXXClassDeclInfo *

// not supported : clang_index_getClientContainer : param  : const CXIdxContainerInfo *

// not supported : clang_index_setClientContainer : param  : const CXIdxContainerInfo *

// not supported : clang_index_getClientEntity : param  : const CXIdxEntityInfo *

// not supported : clang_index_setClientEntity : param  : const CXIdxEntityInfo *

// not supported : clang_IndexAction_create : param CIdx : CXIndex

// not supported : clang_IndexAction_dispose : param  : CXIndexAction

// not supported : clang_indexSourceFile : return value : int

// not supported : clang_indexSourceFileFullArgv : return value : int

// not supported : clang_indexTranslationUnit : return value : int

// not supported : clang_indexLoc_getFileLocation : param loc : CXIdxLoc

// not supported : clang_indexLoc_getCXSourceLocation : param loc : CXIdxLoc

// not supported : clang_Type_visitFields : return value : unsigned int

// not supported : clang_visitCXXBaseClasses : return value : unsigned int

// not supported : clang_visitCXXMethods : return value : unsigned int

// not supported : clang_getBinaryOperatorKindSpelling : param kind : enum CXBinaryOperatorKind

// not supported : clang_getCursorBinaryOperatorKind : return value : enum CXBinaryOperatorKind

// not supported : clang_getUnaryOperatorKindSpelling : param kind : enum CXUnaryOperatorKind

// not supported : clang_getCursorUnaryOperatorKind : return value : enum CXUnaryOperatorKind

// not supported : clang_getRemappings : param  : const char *

// not supported : clang_getRemappingsFromFileList : param  : const char **

// not supported : clang_remap_getNumFiles : return value : unsigned int

// not supported : clang_remap_getFilenames : param  : CXRemapping

// not supported : clang_remap_dispose : param  : CXRemapping
