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

// not supported : clang_getFileName : return value : CXString

// not supported : clang_getFileTime : return value : time_t

// not supported : clang_getFileUniqueID : return value : int

// not supported : clang_File_isEqual : return value : int

// not supported : clang_File_tryGetRealPathName : return value : CXString

// not supported : clang_getNullLocation : return value : CXSourceLocation

// not supported : clang_equalLocations : return value : unsigned int

// not supported : clang_isBeforeInTranslationUnit : return value : unsigned int

// not supported : clang_Location_isInSystemHeader : return value : int

// not supported : clang_Location_isFromMainFile : return value : int

// not supported : clang_getNullRange : return value : CXSourceRange

// not supported : clang_getRange : return value : CXSourceRange

// not supported : clang_equalRanges : return value : unsigned int

// not supported : clang_Range_isNull : return value : int

// not supported : clang_getExpansionLocation : has params

// not supported : clang_getPresumedLocation : has params

// not supported : clang_getInstantiationLocation : has params

// not supported : clang_getSpellingLocation : has params

// not supported : clang_getFileLocation : has params

// not supported : clang_getRangeStart : return value : CXSourceLocation

// not supported : clang_getRangeEnd : return value : CXSourceLocation

// not supported : clang_disposeSourceRangeList : has params

// not supported : clang_getNumDiagnosticsInSet : return value : unsigned int

// not supported : clang_getDiagnosticInSet : return value : CXDiagnostic

// not supported : clang_loadDiagnostics : return value : CXDiagnosticSet

// not supported : clang_disposeDiagnosticSet : has params

// not supported : clang_getChildDiagnostics : return value : CXDiagnosticSet

// not supported : clang_disposeDiagnostic : has params

// not supported : clang_formatDiagnostic : return value : CXString

// not supported : clang_defaultDiagnosticDisplayOptions : return value : unsigned int

// not supported : clang_getDiagnosticSeverity : return value : enum CXDiagnosticSeverity

// not supported : clang_getDiagnosticLocation : return value : CXSourceLocation

// not supported : clang_getDiagnosticSpelling : return value : CXString

// not supported : clang_getDiagnosticOption : return value : CXString

// not supported : clang_getDiagnosticCategory : return value : unsigned int

// not supported : clang_getDiagnosticCategoryName : return value : CXString

// not supported : clang_getDiagnosticCategoryText : return value : CXString

// not supported : clang_getDiagnosticNumRanges : return value : unsigned int

// not supported : clang_getDiagnosticRange : return value : CXSourceRange

// not supported : clang_getDiagnosticNumFixIts : return value : unsigned int

// not supported : clang_getDiagnosticFixIt : return value : CXString

// not supported : clang_createIndex : return value : CXIndex

// not supported : clang_disposeIndex : has params

// not supported : clang_createIndexWithOptions : return value : CXIndex

// not supported : clang_CXIndex_setGlobalOptions : has params

// not supported : clang_CXIndex_getGlobalOptions : return value : unsigned int

// not supported : clang_CXIndex_setInvocationEmissionPathOption : has params

// not supported : clang_isFileMultipleIncludeGuarded : return value : unsigned int

// not supported : clang_getFile : return value : CXFile

// not supported : clang_getFileContents : return value : const char *

// not supported : clang_getLocation : return value : CXSourceLocation

// not supported : clang_getLocationForOffset : return value : CXSourceLocation

// not supported : clang_getSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getAllSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getNumDiagnostics : return value : unsigned int

// not supported : clang_getDiagnostic : return value : CXDiagnostic

// not supported : clang_getDiagnosticSetFromTU : return value : CXDiagnosticSet

// not supported : clang_getTranslationUnitSpelling : return value : CXString

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

// not supported : clang_getCXTUResourceUsage : return value : CXTUResourceUsage

// not supported : clang_disposeCXTUResourceUsage : has params

// not supported : clang_getTranslationUnitTargetInfo : return value : CXTargetInfo

// not supported : clang_TargetInfo_dispose : has params

// not supported : clang_TargetInfo_getTriple : return value : CXString

// not supported : clang_TargetInfo_getPointerWidth : return value : int

// not supported : clang_getNullCursor : return value : CXCursor

// not supported : clang_getTranslationUnitCursor : return value : CXCursor

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

// not supported : clang_Cursor_getVarDeclInitializer : return value : CXCursor

// not supported : clang_Cursor_hasVarDeclGlobalStorage : return value : int

// not supported : clang_Cursor_hasVarDeclExternalStorage : return value : int

// not supported : clang_getCursorLanguage : return value : enum CXLanguageKind

// not supported : clang_getCursorTLSKind : return value : enum CXTLSKind

// not supported : clang_Cursor_getTranslationUnit : return value : CXTranslationUnit

// not supported : clang_createCXCursorSet : return value : CXCursorSet

// not supported : clang_disposeCXCursorSet : has params

// not supported : clang_CXCursorSet_contains : return value : unsigned int

// not supported : clang_CXCursorSet_insert : return value : unsigned int

// not supported : clang_getCursorSemanticParent : return value : CXCursor

// not supported : clang_getCursorLexicalParent : return value : CXCursor

// not supported : clang_getOverriddenCursors : has params

// not supported : clang_disposeOverriddenCursors : has params

// not supported : clang_getIncludedFile : return value : CXFile

// not supported : clang_getCursor : return value : CXCursor

// not supported : clang_getCursorLocation : return value : CXSourceLocation

// not supported : clang_getCursorExtent : return value : CXSourceRange

// not supported : clang_getCursorType : return value : CXType

// not supported : clang_getTypeSpelling : return value : CXString

// not supported : clang_getTypedefDeclUnderlyingType : return value : CXType

// not supported : clang_getEnumDeclIntegerType : return value : CXType

// not supported : clang_getEnumConstantDeclValue : return value : long long

// not supported : clang_getEnumConstantDeclUnsignedValue : return value : unsigned long long

// not supported : clang_Cursor_isBitField : return value : unsigned int

// not supported : clang_getFieldDeclBitWidth : return value : int

// not supported : clang_Cursor_getNumArguments : return value : int

// not supported : clang_Cursor_getArgument : return value : CXCursor

// not supported : clang_Cursor_getNumTemplateArguments : return value : int

// not supported : clang_Cursor_getTemplateArgumentKind : return value : enum CXTemplateArgumentKind

// not supported : clang_Cursor_getTemplateArgumentType : return value : CXType

// not supported : clang_Cursor_getTemplateArgumentValue : return value : long long

// not supported : clang_Cursor_getTemplateArgumentUnsignedValue : return value : unsigned long long

// not supported : clang_equalTypes : return value : unsigned int

// not supported : clang_getCanonicalType : return value : CXType

// not supported : clang_isConstQualifiedType : return value : unsigned int

// not supported : clang_Cursor_isMacroFunctionLike : return value : unsigned int

// not supported : clang_Cursor_isMacroBuiltin : return value : unsigned int

// not supported : clang_Cursor_isFunctionInlined : return value : unsigned int

// not supported : clang_isVolatileQualifiedType : return value : unsigned int

// not supported : clang_isRestrictQualifiedType : return value : unsigned int

// not supported : clang_getAddressSpace : return value : unsigned int

// not supported : clang_getTypedefName : return value : CXString

// not supported : clang_getPointeeType : return value : CXType

// not supported : clang_getUnqualifiedType : return value : CXType

// not supported : clang_getNonReferenceType : return value : CXType

// not supported : clang_getTypeDeclaration : return value : CXCursor

// not supported : clang_getDeclObjCTypeEncoding : return value : CXString

// not supported : clang_Type_getObjCEncoding : return value : CXString

// not supported : clang_getTypeKindSpelling : return value : CXString

// not supported : clang_getFunctionTypeCallingConv : return value : enum CXCallingConv

// not supported : clang_getResultType : return value : CXType

// not supported : clang_getExceptionSpecificationType : return value : int

// not supported : clang_getNumArgTypes : return value : int

// not supported : clang_getArgType : return value : CXType

// not supported : clang_Type_getObjCObjectBaseType : return value : CXType

// not supported : clang_Type_getNumObjCProtocolRefs : return value : unsigned int

// not supported : clang_Type_getObjCProtocolDecl : return value : CXCursor

// not supported : clang_Type_getNumObjCTypeArgs : return value : unsigned int

// not supported : clang_Type_getObjCTypeArg : return value : CXType

// not supported : clang_isFunctionTypeVariadic : return value : unsigned int

// not supported : clang_getCursorResultType : return value : CXType

// not supported : clang_getCursorExceptionSpecificationType : return value : int

// not supported : clang_isPODType : return value : unsigned int

// not supported : clang_getElementType : return value : CXType

// not supported : clang_getNumElements : return value : long long

// not supported : clang_getArrayElementType : return value : CXType

// not supported : clang_getArraySize : return value : long long

// not supported : clang_Type_getNamedType : return value : CXType

// not supported : clang_Type_isTransparentTagTypedef : return value : unsigned int

// not supported : clang_Type_getNullability : return value : enum CXTypeNullabilityKind

// not supported : clang_Type_getAlignOf : return value : long long

// not supported : clang_Type_getClassType : return value : CXType

// not supported : clang_Type_getSizeOf : return value : long long

// not supported : clang_Type_getOffsetOf : return value : long long

// not supported : clang_Type_getModifiedType : return value : CXType

// not supported : clang_Type_getValueType : return value : CXType

// not supported : clang_Cursor_getOffsetOfField : return value : long long

// not supported : clang_Cursor_isAnonymous : return value : unsigned int

// not supported : clang_Cursor_isAnonymousRecordDecl : return value : unsigned int

// not supported : clang_Cursor_isInlineNamespace : return value : unsigned int

// not supported : clang_Type_getNumTemplateArguments : return value : int

// not supported : clang_Type_getTemplateArgumentAsType : return value : CXType

// not supported : clang_Type_getCXXRefQualifier : return value : enum CXRefQualifierKind

// not supported : clang_isVirtualBase : return value : unsigned int

// not supported : clang_getOffsetOfBase : return value : long long

// not supported : clang_getCXXAccessSpecifier : return value : enum CX_CXXAccessSpecifier

// not supported : clang_Cursor_getBinaryOpcode : return value : enum CX_BinaryOperatorKind

// not supported : clang_Cursor_getBinaryOpcodeStr : return value : CXString

// not supported : clang_Cursor_getStorageClass : return value : enum CX_StorageClass

// not supported : clang_getNumOverloadedDecls : return value : unsigned int

// not supported : clang_getOverloadedDecl : return value : CXCursor

// not supported : clang_getIBOutletCollectionType : return value : CXType

// not supported : clang_visitChildren : return value : unsigned int

// not supported : clang_visitChildrenWithBlock : return value : unsigned int

// not supported : clang_getCursorUSR : return value : CXString

// not supported : clang_constructUSR_ObjCClass : return value : CXString

// not supported : clang_constructUSR_ObjCCategory : return value : CXString

// not supported : clang_constructUSR_ObjCProtocol : return value : CXString

// not supported : clang_constructUSR_ObjCIvar : return value : CXString

// not supported : clang_constructUSR_ObjCMethod : return value : CXString

// not supported : clang_constructUSR_ObjCProperty : return value : CXString

// not supported : clang_getCursorSpelling : return value : CXString

// not supported : clang_Cursor_getSpellingNameRange : return value : CXSourceRange

// not supported : clang_PrintingPolicy_getProperty : return value : unsigned int

// not supported : clang_PrintingPolicy_setProperty : has params

// not supported : clang_getCursorPrintingPolicy : return value : CXPrintingPolicy

// not supported : clang_PrintingPolicy_dispose : has params

// not supported : clang_getCursorPrettyPrinted : return value : CXString

// not supported : clang_getTypePrettyPrinted : return value : CXString

// not supported : clang_getFullyQualifiedName : return value : CXString

// not supported : clang_getCursorDisplayName : return value : CXString

// not supported : clang_getCursorReferenced : return value : CXCursor

// not supported : clang_getCursorDefinition : return value : CXCursor

// not supported : clang_isCursorDefinition : return value : unsigned int

// not supported : clang_getCanonicalCursor : return value : CXCursor

// not supported : clang_Cursor_getObjCSelectorIndex : return value : int

// not supported : clang_Cursor_isDynamicCall : return value : int

// not supported : clang_Cursor_getReceiverType : return value : CXType

// not supported : clang_Cursor_getObjCPropertyAttributes : return value : unsigned int

// not supported : clang_Cursor_getObjCPropertyGetterName : return value : CXString

// not supported : clang_Cursor_getObjCPropertySetterName : return value : CXString

// not supported : clang_Cursor_getObjCDeclQualifiers : return value : unsigned int

// not supported : clang_Cursor_isObjCOptional : return value : unsigned int

// not supported : clang_Cursor_isVariadic : return value : unsigned int

// not supported : clang_Cursor_isExternalSymbol : return value : unsigned int

// not supported : clang_Cursor_getCommentRange : return value : CXSourceRange

// not supported : clang_Cursor_getRawCommentText : return value : CXString

// not supported : clang_Cursor_getBriefCommentText : return value : CXString

// not supported : clang_Cursor_getMangling : return value : CXString

// not supported : clang_Cursor_getCXXManglings : return value : CXStringSet *

// not supported : clang_Cursor_getObjCManglings : return value : CXStringSet *

// not supported : clang_Cursor_getGCCAssemblyTemplate : return value : CXString

// not supported : clang_Cursor_isGCCAssemblyHasGoto : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumOutputs : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumInputs : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyInput : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyOutput : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyNumClobbers : return value : unsigned int

// not supported : clang_Cursor_getGCCAssemblyClobber : return value : CXString

// not supported : clang_Cursor_isGCCAssemblyVolatile : return value : unsigned int

// not supported : clang_Cursor_getModule : return value : CXModule

// not supported : clang_getModuleForFile : return value : CXModule

// not supported : clang_Module_getASTFile : return value : CXFile

// not supported : clang_Module_getParent : return value : CXModule

// not supported : clang_Module_getName : return value : CXString

// not supported : clang_Module_getFullName : return value : CXString

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

// not supported : clang_getSpecializedCursorTemplate : return value : CXCursor

// not supported : clang_getCursorReferenceNameRange : return value : CXSourceRange

// not supported : clang_getToken : return value : CXToken *

// not supported : clang_getTokenKind : return value : CXTokenKind

// not supported : clang_getTokenSpelling : return value : CXString

// not supported : clang_getTokenLocation : return value : CXSourceLocation

// not supported : clang_getTokenExtent : return value : CXSourceRange

// not supported : clang_tokenize : has params

// not supported : clang_annotateTokens : has params

// not supported : clang_disposeTokens : has params

// not supported : clang_getCursorKindSpelling : return value : CXString

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

// not supported : clang_getCompletionChunkText : return value : CXString

// not supported : clang_getCompletionChunkCompletionString : return value : CXCompletionString

// not supported : clang_getNumCompletionChunks : return value : unsigned int

// not supported : clang_getCompletionPriority : return value : unsigned int

// not supported : clang_getCompletionAvailability : return value : enum CXAvailabilityKind

// not supported : clang_getCompletionNumAnnotations : return value : unsigned int

// not supported : clang_getCompletionAnnotation : return value : CXString

// not supported : clang_getCompletionParent : return value : CXString

// not supported : clang_getCompletionBriefComment : return value : CXString

// not supported : clang_getCursorCompletionString : return value : CXCompletionString

// not supported : clang_getCompletionNumFixIts : return value : unsigned int

// not supported : clang_getCompletionFixIt : return value : CXString

// not supported : clang_defaultCodeCompleteOptions : return value : unsigned int

// not supported : clang_codeCompleteAt : return value : CXCodeCompleteResults *

// not supported : clang_sortCodeCompletionResults : has params

// not supported : clang_disposeCodeCompleteResults : has params

// not supported : clang_codeCompleteGetNumDiagnostics : return value : unsigned int

// not supported : clang_codeCompleteGetDiagnostic : return value : CXDiagnostic

// not supported : clang_codeCompleteGetContexts : return value : unsigned long long

// not supported : clang_codeCompleteGetContainerKind : return value : enum CXCursorKind

// not supported : clang_codeCompleteGetContainerUSR : return value : CXString

// not supported : clang_codeCompleteGetObjCSelector : return value : CXString

// not supported : clang_getClangVersion : return value : CXString

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

// not supported : clang_indexLoc_getCXSourceLocation : return value : CXSourceLocation

// not supported : clang_Type_visitFields : return value : unsigned int

// not supported : clang_visitCXXBaseClasses : return value : unsigned int

// not supported : clang_visitCXXMethods : return value : unsigned int

// not supported : clang_getBinaryOperatorKindSpelling : return value : CXString

// not supported : clang_getCursorBinaryOperatorKind : return value : enum CXBinaryOperatorKind

// not supported : clang_getUnaryOperatorKindSpelling : return value : CXString

// not supported : clang_getCursorUnaryOperatorKind : return value : enum CXUnaryOperatorKind

// not supported : clang_getRemappings : return value : CXRemapping

// not supported : clang_getRemappingsFromFileList : return value : CXRemapping

// not supported : clang_remap_getNumFiles : return value : unsigned int

// not supported : clang_remap_getFilenames : has params

// not supported : clang_remap_dispose : has params
