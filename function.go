// This is a generated file. DO NOT EDIT.

package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

// not supported : clang_getCString : return value : const char *

// not supported : clang_disposeString : has params

// not supported : clang_disposeStringSet : has params

// not supported : clang_getBuildSessionTimestamp : return value : unsigned long long

// not supported : clang_VirtualFileOverlay_create : return value : CXVirtualFileOverlay

// not supported : clang_VirtualFileOverlay_addFileMapping : return value : enum CXErrorCode

// not supported : clang_VirtualFileOverlay_setCaseSensitivity : return value : enum CXErrorCode

// not supported : clang_VirtualFileOverlay_writeToBuffer : return value : enum CXErrorCode

// not supported : clang_free : has params

// not supported : clang_VirtualFileOverlay_dispose : has params

// not supported : clang_ModuleMapDescriptor_create : return value : CXModuleMapDescriptor

// not supported : clang_ModuleMapDescriptor_setFrameworkModuleName : return value : enum CXErrorCode

// not supported : clang_ModuleMapDescriptor_setUmbrellaHeader : return value : enum CXErrorCode

// not supported : clang_ModuleMapDescriptor_writeToBuffer : return value : enum CXErrorCode

// not supported : clang_ModuleMapDescriptor_dispose : has params

// not supported : clang_getFileName : has params

// not supported : clang_getFileTime : return value : time_t

// not supported : clang_getFileUniqueID : return value : int

// not supported : clang_File_isEqual : return value : int

// not supported : clang_File_tryGetRealPathName : has params

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

// not supported : clang_getRange : has params

// not supported : clang_equalRanges : return value : unsigned int

// not supported : clang_Range_isNull : return value : int

// not supported : clang_getExpansionLocation : has params

// not supported : clang_getPresumedLocation : has params

// not supported : clang_getInstantiationLocation : has params

// not supported : clang_getSpellingLocation : has params

// not supported : clang_getFileLocation : has params

// not supported : clang_getRangeStart : has params

// not supported : clang_getRangeEnd : has params

// not supported : clang_disposeSourceRangeList : has params

// not supported : clang_getNumDiagnosticsInSet : return value : unsigned int

// not supported : clang_getDiagnosticInSet : return value : CXDiagnostic

// not supported : clang_loadDiagnostics : return value : CXDiagnosticSet

// not supported : clang_disposeDiagnosticSet : has params

// not supported : clang_getChildDiagnostics : return value : CXDiagnosticSet

// not supported : clang_disposeDiagnostic : has params

// not supported : clang_formatDiagnostic : has params

// not supported : clang_defaultDiagnosticDisplayOptions : return value : unsigned int

// not supported : clang_getDiagnosticSeverity : return value : enum CXDiagnosticSeverity

// not supported : clang_getDiagnosticLocation : has params

// not supported : clang_getDiagnosticSpelling : has params

// not supported : clang_getDiagnosticOption : has params

// not supported : clang_getDiagnosticCategory : return value : unsigned int

// not supported : clang_getDiagnosticCategoryName : has params

// not supported : clang_getDiagnosticCategoryText : has params

// not supported : clang_getDiagnosticNumRanges : return value : unsigned int

// not supported : clang_getDiagnosticRange : has params

// not supported : clang_getDiagnosticNumFixIts : return value : unsigned int

// not supported : clang_getDiagnosticFixIt : has params

// not supported : clang_createIndex : return value : CXIndex

// not supported : clang_disposeIndex : has params

// not supported : clang_createIndexWithOptions : return value : CXIndex

// not supported : clang_CXIndex_setGlobalOptions : has params

// not supported : clang_CXIndex_getGlobalOptions : return value : unsigned int

// not supported : clang_CXIndex_setInvocationEmissionPathOption : has params

// not supported : clang_isFileMultipleIncludeGuarded : return value : unsigned int

// not supported : clang_getFile : return value : CXFile

// not supported : clang_getFileContents : return value : const char *

// not supported : clang_getLocation : has params

// not supported : clang_getLocationForOffset : has params

// not supported : clang_getSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getAllSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getNumDiagnostics : return value : unsigned int

// not supported : clang_getDiagnostic : return value : CXDiagnostic

// not supported : clang_getDiagnosticSetFromTU : return value : CXDiagnosticSet

// not supported : clang_getTranslationUnitSpelling : has params

// not supported : clang_createTranslationUnitFromSourceFile : return value : CXTranslationUnit

// not supported : clang_createTranslationUnit : return value : CXTranslationUnit

// not supported : clang_createTranslationUnit2 : return value : enum CXErrorCode

// not supported : clang_defaultEditingTranslationUnitOptions : return value : unsigned int

// not supported : clang_parseTranslationUnit : return value : CXTranslationUnit

// not supported : clang_parseTranslationUnit2 : return value : enum CXErrorCode

// not supported : clang_parseTranslationUnit2FullArgv : return value : enum CXErrorCode

// not supported : clang_defaultSaveOptions : return value : unsigned int

// not supported : clang_saveTranslationUnit : return value : int

// not supported : clang_suspendTranslationUnit : return value : unsigned int

// not supported : clang_disposeTranslationUnit : has params

// not supported : clang_defaultReparseOptions : return value : unsigned int

// not supported : clang_reparseTranslationUnit : return value : int

// not supported : clang_getTUResourceUsageName : return value : const char *

// not supported : clang_getCXTUResourceUsage : has params

// not supported : clang_disposeCXTUResourceUsage : has params

// not supported : clang_getTranslationUnitTargetInfo : return value : CXTargetInfo

// not supported : clang_TargetInfo_dispose : has params

// not supported : clang_TargetInfo_getTriple : has params

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

// not supported : clang_getTranslationUnitCursor : has params

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

// not supported : clang_disposeCXPlatformAvailability : has params

// not supported : clang_Cursor_getVarDeclInitializer : has params

// not supported : clang_Cursor_hasVarDeclGlobalStorage : return value : int

// not supported : clang_Cursor_hasVarDeclExternalStorage : return value : int

// not supported : clang_getCursorLanguage : return value : enum CXLanguageKind

// not supported : clang_getCursorTLSKind : return value : enum CXTLSKind

// not supported : clang_Cursor_getTranslationUnit : return value : CXTranslationUnit

// not supported : clang_createCXCursorSet : return value : CXCursorSet

// not supported : clang_disposeCXCursorSet : has params

// not supported : clang_CXCursorSet_contains : return value : unsigned int

// not supported : clang_CXCursorSet_insert : return value : unsigned int

// not supported : clang_getCursorSemanticParent : has params

// not supported : clang_getCursorLexicalParent : has params

// not supported : clang_getOverriddenCursors : has params

// not supported : clang_disposeOverriddenCursors : has params

// not supported : clang_getIncludedFile : return value : CXFile

// not supported : clang_getCursor : has params

// not supported : clang_getCursorLocation : has params

// not supported : clang_getCursorExtent : has params

// not supported : clang_getCursorType : has params

// not supported : clang_getTypeSpelling : has params

// not supported : clang_getTypedefDeclUnderlyingType : has params

// not supported : clang_getEnumDeclIntegerType : has params

// not supported : clang_getEnumConstantDeclValue : return value : long long

// not supported : clang_getEnumConstantDeclUnsignedValue : return value : unsigned long long

// not supported : clang_Cursor_isBitField : return value : unsigned int

// not supported : clang_getFieldDeclBitWidth : return value : int

// not supported : clang_Cursor_getNumArguments : return value : int

// not supported : clang_Cursor_getArgument : has params

// not supported : clang_Cursor_getNumTemplateArguments : return value : int

// not supported : clang_Cursor_getTemplateArgumentKind : return value : enum CXTemplateArgumentKind

// not supported : clang_Cursor_getTemplateArgumentType : has params

// not supported : clang_Cursor_getTemplateArgumentValue : return value : long long

// not supported : clang_Cursor_getTemplateArgumentUnsignedValue : return value : unsigned long long

// not supported : clang_equalTypes : return value : unsigned int

// not supported : clang_getCanonicalType : has params

// not supported : clang_isConstQualifiedType : return value : unsigned int

// not supported : clang_Cursor_isMacroFunctionLike : return value : unsigned int

// not supported : clang_Cursor_isMacroBuiltin : return value : unsigned int

// not supported : clang_Cursor_isFunctionInlined : return value : unsigned int

// not supported : clang_isVolatileQualifiedType : return value : unsigned int

// not supported : clang_isRestrictQualifiedType : return value : unsigned int

// not supported : clang_getAddressSpace : return value : unsigned int

// not supported : clang_getTypedefName : has params

// not supported : clang_getPointeeType : has params

// not supported : clang_getUnqualifiedType : has params

// not supported : clang_getNonReferenceType : has params

// not supported : clang_getTypeDeclaration : has params

// not supported : clang_getDeclObjCTypeEncoding : has params

// not supported : clang_Type_getObjCEncoding : has params

// not supported : clang_getTypeKindSpelling : has params

// not supported : clang_getFunctionTypeCallingConv : return value : enum CXCallingConv

// not supported : clang_getResultType : has params

// not supported : clang_getExceptionSpecificationType : return value : int

// not supported : clang_getNumArgTypes : return value : int

// not supported : clang_getArgType : has params

// not supported : clang_Type_getObjCObjectBaseType : has params

// not supported : clang_Type_getNumObjCProtocolRefs : return value : unsigned int

// not supported : clang_Type_getObjCProtocolDecl : has params

// not supported : clang_Type_getNumObjCTypeArgs : return value : unsigned int

// not supported : clang_Type_getObjCTypeArg : has params

// not supported : clang_isFunctionTypeVariadic : return value : unsigned int

// not supported : clang_getCursorResultType : has params

// not supported : clang_getCursorExceptionSpecificationType : return value : int

// not supported : clang_isPODType : return value : unsigned int

// not supported : clang_getElementType : has params

// not supported : clang_getNumElements : return value : long long

// not supported : clang_getArrayElementType : has params

// not supported : clang_getArraySize : return value : long long

// not supported : clang_Type_getNamedType : has params

// not supported : clang_Type_isTransparentTagTypedef : return value : unsigned int

// not supported : clang_Type_getNullability : return value : enum CXTypeNullabilityKind

// not supported : clang_Type_getAlignOf : return value : long long

// not supported : clang_Type_getClassType : has params

// not supported : clang_Type_getSizeOf : return value : long long

// not supported : clang_Type_getOffsetOf : return value : long long

// not supported : clang_Type_getModifiedType : has params

// not supported : clang_Type_getValueType : has params

// not supported : clang_Cursor_getOffsetOfField : return value : long long

// not supported : clang_Cursor_isAnonymous : return value : unsigned int

// not supported : clang_Cursor_isAnonymousRecordDecl : return value : unsigned int

// not supported : clang_Cursor_isInlineNamespace : return value : unsigned int

// not supported : clang_Type_getNumTemplateArguments : return value : int

// not supported : clang_Type_getTemplateArgumentAsType : has params

// not supported : clang_Type_getCXXRefQualifier : return value : enum CXRefQualifierKind

// not supported : clang_isVirtualBase : return value : unsigned int

// not supported : clang_getOffsetOfBase : return value : long long

// not supported : clang_getCXXAccessSpecifier : return value : enum CX_CXXAccessSpecifier

// not supported : clang_Cursor_getBinaryOpcode : return value : enum CX_BinaryOperatorKind

// not supported : clang_Cursor_getBinaryOpcodeStr : has params

// not supported : clang_Cursor_getStorageClass : return value : enum CX_StorageClass

// not supported : clang_getNumOverloadedDecls : return value : unsigned int

// not supported : clang_getOverloadedDecl : has params

// not supported : clang_getIBOutletCollectionType : has params

// not supported : clang_visitChildren : return value : unsigned int

// not supported : clang_visitChildrenWithBlock : return value : unsigned int

// not supported : clang_getCursorUSR : has params

// not supported : clang_constructUSR_ObjCClass : has params

// not supported : clang_constructUSR_ObjCCategory : has params

// not supported : clang_constructUSR_ObjCProtocol : has params

// not supported : clang_constructUSR_ObjCIvar : has params

// not supported : clang_constructUSR_ObjCMethod : has params

// not supported : clang_constructUSR_ObjCProperty : has params

// not supported : clang_getCursorSpelling : has params

// not supported : clang_Cursor_getSpellingNameRange : has params

// not supported : clang_PrintingPolicy_getProperty : return value : unsigned int

// not supported : clang_PrintingPolicy_setProperty : has params

// not supported : clang_getCursorPrintingPolicy : return value : CXPrintingPolicy

// not supported : clang_PrintingPolicy_dispose : has params

// not supported : clang_getCursorPrettyPrinted : has params

// not supported : clang_getTypePrettyPrinted : has params

// not supported : clang_getFullyQualifiedName : has params

// not supported : clang_getCursorDisplayName : has params

// not supported : clang_getCursorReferenced : has params

// not supported : clang_getCursorDefinition : has params

// not supported : clang_isCursorDefinition : return value : unsigned int

// not supported : clang_getCanonicalCursor : has params

// not supported : clang_Cursor_getObjCSelectorIndex : return value : int

// not supported : clang_Cursor_isDynamicCall : return value : int

// not supported : clang_Cursor_getReceiverType : has params

// not supported : clang_Cursor_getObjCPropertyAttributes : return value : unsigned int

// not supported : clang_Cursor_getObjCPropertyGetterName : has params

// not supported : clang_Cursor_getObjCPropertySetterName : has params

// not supported : clang_Cursor_getObjCDeclQualifiers : return value : unsigned int

// not supported : clang_Cursor_isObjCOptional : return value : unsigned int

// not supported : clang_Cursor_isVariadic : return value : unsigned int

// not supported : clang_Cursor_isExternalSymbol : return value : unsigned int

// not supported : clang_Cursor_getCommentRange : has params

// not supported : clang_Cursor_getRawCommentText : has params

// not supported : clang_Cursor_getBriefCommentText : has params

// not supported : clang_Cursor_getMangling : has params

// not supported : clang_Cursor_getCXXManglings : return value : CXStringSet *

// not supported : clang_Cursor_getObjCManglings : return value : CXStringSet *

// not supported : clang_Cursor_getGCCAssemblyTemplate : has params

// not supported : clang_Cursor_isGCCAssemblyHasGoto : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumOutputs : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumInputs : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyInput : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyOutput : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumClobbers : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyClobber : has params

// not supported : clang_Cursor_isGCCAssemblyVolatile : return value : unsigned int

// not supported : clang_Cursor_getModule : return value : CXModule

// not supported : clang_getModuleForFile : return value : CXModule

// not supported : clang_Module_getASTFile : return value : CXFile

// not supported : clang_Module_getParent : return value : CXModule

// not supported : clang_Module_getName : has params

// not supported : clang_Module_getFullName : has params

// not supported : clang_Module_isSystem : return value : int

// not supported : clang_Module_getNumTopLevelHeaders : return value : unsigned int

// not supported : clang_Module_getTopLevelHeader : return value : CXFile

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

// not supported : clang_getSpecializedCursorTemplate : has params

// not supported : clang_getCursorReferenceNameRange : has params

// not supported : clang_getToken : return value : CXToken *

// not supported : clang_getTokenKind : return value : CXTokenKind

// not supported : clang_getTokenSpelling : has params

// not supported : clang_getTokenLocation : has params

// not supported : clang_getTokenExtent : has params

// not supported : clang_tokenize : has params

// not supported : clang_annotateTokens : has params

// not supported : clang_disposeTokens : has params

// not supported : clang_getCursorKindSpelling : has params

// not supported : clang_getDefinitionSpellingAndExtent : has params

func EnableStackTraces() {
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_enableStackTraces, ptr_clang_enableStackTraces, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_enableStackTraces", err))
	}
}

// not supported : clang_executeOnThread : has params

// not supported : clang_getCompletionChunkKind : return value : enum CXCompletionChunkKind

// not supported : clang_getCompletionChunkText : has params

// not supported : clang_getCompletionChunkCompletionString : return value : CXCompletionString

// not supported : clang_getNumCompletionChunks : return value : unsigned int

// not supported : clang_getCompletionPriority : return value : unsigned int

// not supported : clang_getCompletionAvailability : return value : enum CXAvailabilityKind

// not supported : clang_getCompletionNumAnnotations : return value : unsigned int

// not supported : clang_getCompletionAnnotation : has params

// not supported : clang_getCompletionParent : has params

// not supported : clang_getCompletionBriefComment : has params

// not supported : clang_getCursorCompletionString : return value : CXCompletionString

// not supported : clang_getCompletionNumFixIts : return value : unsigned int

// not supported : clang_getCompletionFixIt : has params

// not supported : clang_defaultCodeCompleteOptions : return value : unsigned int

// not supported : clang_codeCompleteAt : return value : CXCodeCompleteResults *

// not supported : clang_sortCodeCompletionResults : has params

// not supported : clang_disposeCodeCompleteResults : has params

// not supported : clang_codeCompleteGetNumDiagnostics : return value : unsigned int

// not supported : clang_codeCompleteGetDiagnostic : return value : CXDiagnostic

// not supported : clang_codeCompleteGetContexts : return value : unsigned long long

// not supported : clang_codeCompleteGetContainerKind : return value : enum CXCursorKind

// not supported : clang_codeCompleteGetContainerUSR : has params

// not supported : clang_codeCompleteGetObjCSelector : has params

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

// not supported : clang_toggleCrashRecovery : has params

// not supported : clang_getInclusions : has params

// not supported : clang_Cursor_Evaluate : return value : CXEvalResult

// not supported : clang_EvalResult_getKind : return value : CXEvalResultKind

// not supported : clang_EvalResult_getAsInt : return value : int

// not supported : clang_EvalResult_getAsLongLong : return value : long long

// not supported : clang_EvalResult_isUnsignedInt : return value : unsigned int

// not supported : clang_EvalResult_getAsUnsigned : return value : unsigned long long

// not supported : clang_EvalResult_getAsDouble : return value : double

// not supported : clang_EvalResult_getAsStr : return value : const char *

// not supported : clang_EvalResult_dispose : has params

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

// not supported : clang_index_getClientContainer : return value : CXIdxClientContainer

// not supported : clang_index_setClientContainer : has params

// not supported : clang_index_getClientEntity : return value : CXIdxClientEntity

// not supported : clang_index_setClientEntity : has params

// not supported : clang_IndexAction_create : return value : CXIndexAction

// not supported : clang_IndexAction_dispose : has params

// not supported : clang_indexSourceFile : return value : int

// not supported : clang_indexSourceFileFullArgv : return value : int

// not supported : clang_indexTranslationUnit : return value : int

// not supported : clang_indexLoc_getFileLocation : has params

// not supported : clang_indexLoc_getCXSourceLocation : has params

// not supported : clang_Type_visitFields : return value : unsigned int

// not supported : clang_visitCXXBaseClasses : return value : unsigned int

// not supported : clang_visitCXXMethods : return value : unsigned int

// not supported : clang_getBinaryOperatorKindSpelling : has params

// not supported : clang_getCursorBinaryOperatorKind : return value : enum CXBinaryOperatorKind

// not supported : clang_getUnaryOperatorKindSpelling : has params

// not supported : clang_getCursorUnaryOperatorKind : return value : enum CXUnaryOperatorKind

// not supported : clang_getRemappings : return value : CXRemapping

// not supported : clang_getRemappingsFromFileList : return value : CXRemapping

// not supported : clang_remap_getNumFiles : return value : unsigned int

// not supported : clang_remap_getFilenames : has params

// not supported : clang_remap_dispose : has params
