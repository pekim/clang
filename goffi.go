// This is a generated file. DO NOT EDIT.

package clang

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
	lib "github.com/pekim/clang/internal/lib"
)

var cif_clang_disposeString = &types.CallInterface{}
var cif_clang_VirtualFileOverlay_create = &types.CallInterface{}
var cif_clang_VirtualFileOverlay_setCaseSensitivity = &types.CallInterface{}
var cif_clang_VirtualFileOverlay_dispose = &types.CallInterface{}
var cif_clang_ModuleMapDescriptor_create = &types.CallInterface{}
var cif_clang_ModuleMapDescriptor_dispose = &types.CallInterface{}
var cif_clang_getFileName = &types.CallInterface{}
var cif_clang_File_isEqual = &types.CallInterface{}
var cif_clang_File_tryGetRealPathName = &types.CallInterface{}
var cif_clang_getNullLocation = &types.CallInterface{}
var cif_clang_equalLocations = &types.CallInterface{}
var cif_clang_isBeforeInTranslationUnit = &types.CallInterface{}
var cif_clang_Location_isInSystemHeader = &types.CallInterface{}
var cif_clang_Location_isFromMainFile = &types.CallInterface{}
var cif_clang_getNullRange = &types.CallInterface{}
var cif_clang_getRange = &types.CallInterface{}
var cif_clang_equalRanges = &types.CallInterface{}
var cif_clang_Range_isNull = &types.CallInterface{}
var cif_clang_getRangeStart = &types.CallInterface{}
var cif_clang_getRangeEnd = &types.CallInterface{}
var cif_clang_getNumDiagnosticsInSet = &types.CallInterface{}
var cif_clang_getDiagnosticInSet = &types.CallInterface{}
var cif_clang_disposeDiagnosticSet = &types.CallInterface{}
var cif_clang_getChildDiagnostics = &types.CallInterface{}
var cif_clang_disposeDiagnostic = &types.CallInterface{}
var cif_clang_formatDiagnostic = &types.CallInterface{}
var cif_clang_defaultDiagnosticDisplayOptions = &types.CallInterface{}
var cif_clang_getDiagnosticSeverity = &types.CallInterface{}
var cif_clang_getDiagnosticLocation = &types.CallInterface{}
var cif_clang_getDiagnosticSpelling = &types.CallInterface{}
var cif_clang_getDiagnosticCategory = &types.CallInterface{}
var cif_clang_getDiagnosticCategoryName = &types.CallInterface{}
var cif_clang_getDiagnosticCategoryText = &types.CallInterface{}
var cif_clang_getDiagnosticNumRanges = &types.CallInterface{}
var cif_clang_getDiagnosticRange = &types.CallInterface{}
var cif_clang_getDiagnosticNumFixIts = &types.CallInterface{}
var cif_clang_createIndex = &types.CallInterface{}
var cif_clang_disposeIndex = &types.CallInterface{}
var cif_clang_CXIndex_setGlobalOptions = &types.CallInterface{}
var cif_clang_CXIndex_getGlobalOptions = &types.CallInterface{}
var cif_clang_isFileMultipleIncludeGuarded = &types.CallInterface{}
var cif_clang_getLocation = &types.CallInterface{}
var cif_clang_getLocationForOffset = &types.CallInterface{}
var cif_clang_getNumDiagnostics = &types.CallInterface{}
var cif_clang_getDiagnostic = &types.CallInterface{}
var cif_clang_getDiagnosticSetFromTU = &types.CallInterface{}
var cif_clang_getTranslationUnitSpelling = &types.CallInterface{}
var cif_clang_defaultEditingTranslationUnitOptions = &types.CallInterface{}
var cif_clang_defaultSaveOptions = &types.CallInterface{}
var cif_clang_suspendTranslationUnit = &types.CallInterface{}
var cif_clang_disposeTranslationUnit = &types.CallInterface{}
var cif_clang_defaultReparseOptions = &types.CallInterface{}
var cif_clang_getCXTUResourceUsage = &types.CallInterface{}
var cif_clang_disposeCXTUResourceUsage = &types.CallInterface{}
var cif_clang_getTranslationUnitTargetInfo = &types.CallInterface{}
var cif_clang_TargetInfo_dispose = &types.CallInterface{}
var cif_clang_TargetInfo_getTriple = &types.CallInterface{}
var cif_clang_TargetInfo_getPointerWidth = &types.CallInterface{}
var cif_clang_getNullCursor = &types.CallInterface{}
var cif_clang_getTranslationUnitCursor = &types.CallInterface{}
var cif_clang_equalCursors = &types.CallInterface{}
var cif_clang_Cursor_isNull = &types.CallInterface{}
var cif_clang_hashCursor = &types.CallInterface{}
var cif_clang_getCursorKind = &types.CallInterface{}
var cif_clang_isInvalidDeclaration = &types.CallInterface{}
var cif_clang_Cursor_hasAttrs = &types.CallInterface{}
var cif_clang_getCursorLinkage = &types.CallInterface{}
var cif_clang_getCursorVisibility = &types.CallInterface{}
var cif_clang_getCursorAvailability = &types.CallInterface{}
var cif_clang_Cursor_getVarDeclInitializer = &types.CallInterface{}
var cif_clang_Cursor_hasVarDeclGlobalStorage = &types.CallInterface{}
var cif_clang_Cursor_hasVarDeclExternalStorage = &types.CallInterface{}
var cif_clang_getCursorLanguage = &types.CallInterface{}
var cif_clang_getCursorTLSKind = &types.CallInterface{}
var cif_clang_Cursor_getTranslationUnit = &types.CallInterface{}
var cif_clang_createCXCursorSet = &types.CallInterface{}
var cif_clang_disposeCXCursorSet = &types.CallInterface{}
var cif_clang_CXCursorSet_contains = &types.CallInterface{}
var cif_clang_CXCursorSet_insert = &types.CallInterface{}
var cif_clang_getCursorSemanticParent = &types.CallInterface{}
var cif_clang_getCursorLexicalParent = &types.CallInterface{}
var cif_clang_getIncludedFile = &types.CallInterface{}
var cif_clang_getCursor = &types.CallInterface{}
var cif_clang_getCursorLocation = &types.CallInterface{}
var cif_clang_getCursorExtent = &types.CallInterface{}
var cif_clang_getCursorType = &types.CallInterface{}
var cif_clang_getTypeSpelling = &types.CallInterface{}
var cif_clang_getTypedefDeclUnderlyingType = &types.CallInterface{}
var cif_clang_getEnumDeclIntegerType = &types.CallInterface{}
var cif_clang_Cursor_isBitField = &types.CallInterface{}
var cif_clang_getFieldDeclBitWidth = &types.CallInterface{}
var cif_clang_Cursor_getNumArguments = &types.CallInterface{}
var cif_clang_Cursor_getArgument = &types.CallInterface{}
var cif_clang_Cursor_getNumTemplateArguments = &types.CallInterface{}
var cif_clang_Cursor_getTemplateArgumentKind = &types.CallInterface{}
var cif_clang_Cursor_getTemplateArgumentType = &types.CallInterface{}
var cif_clang_equalTypes = &types.CallInterface{}
var cif_clang_getCanonicalType = &types.CallInterface{}
var cif_clang_isConstQualifiedType = &types.CallInterface{}
var cif_clang_Cursor_isMacroFunctionLike = &types.CallInterface{}
var cif_clang_Cursor_isMacroBuiltin = &types.CallInterface{}
var cif_clang_Cursor_isFunctionInlined = &types.CallInterface{}
var cif_clang_isVolatileQualifiedType = &types.CallInterface{}
var cif_clang_isRestrictQualifiedType = &types.CallInterface{}
var cif_clang_getAddressSpace = &types.CallInterface{}
var cif_clang_getTypedefName = &types.CallInterface{}
var cif_clang_getPointeeType = &types.CallInterface{}
var cif_clang_getUnqualifiedType = &types.CallInterface{}
var cif_clang_getNonReferenceType = &types.CallInterface{}
var cif_clang_getTypeDeclaration = &types.CallInterface{}
var cif_clang_getDeclObjCTypeEncoding = &types.CallInterface{}
var cif_clang_Type_getObjCEncoding = &types.CallInterface{}
var cif_clang_getFunctionTypeCallingConv = &types.CallInterface{}
var cif_clang_getResultType = &types.CallInterface{}
var cif_clang_getExceptionSpecificationType = &types.CallInterface{}
var cif_clang_getNumArgTypes = &types.CallInterface{}
var cif_clang_getArgType = &types.CallInterface{}
var cif_clang_Type_getObjCObjectBaseType = &types.CallInterface{}
var cif_clang_Type_getNumObjCProtocolRefs = &types.CallInterface{}
var cif_clang_Type_getObjCProtocolDecl = &types.CallInterface{}
var cif_clang_Type_getNumObjCTypeArgs = &types.CallInterface{}
var cif_clang_Type_getObjCTypeArg = &types.CallInterface{}
var cif_clang_isFunctionTypeVariadic = &types.CallInterface{}
var cif_clang_getCursorResultType = &types.CallInterface{}
var cif_clang_getCursorExceptionSpecificationType = &types.CallInterface{}
var cif_clang_isPODType = &types.CallInterface{}
var cif_clang_getElementType = &types.CallInterface{}
var cif_clang_getArrayElementType = &types.CallInterface{}
var cif_clang_Type_getNamedType = &types.CallInterface{}
var cif_clang_Type_isTransparentTagTypedef = &types.CallInterface{}
var cif_clang_Type_getNullability = &types.CallInterface{}
var cif_clang_Type_getClassType = &types.CallInterface{}
var cif_clang_Type_getModifiedType = &types.CallInterface{}
var cif_clang_Type_getValueType = &types.CallInterface{}
var cif_clang_Cursor_isAnonymous = &types.CallInterface{}
var cif_clang_Cursor_isAnonymousRecordDecl = &types.CallInterface{}
var cif_clang_Cursor_isInlineNamespace = &types.CallInterface{}
var cif_clang_Type_getNumTemplateArguments = &types.CallInterface{}
var cif_clang_Type_getTemplateArgumentAsType = &types.CallInterface{}
var cif_clang_Type_getCXXRefQualifier = &types.CallInterface{}
var cif_clang_isVirtualBase = &types.CallInterface{}
var cif_clang_getCXXAccessSpecifier = &types.CallInterface{}
var cif_clang_Cursor_getBinaryOpcode = &types.CallInterface{}
var cif_clang_Cursor_getStorageClass = &types.CallInterface{}
var cif_clang_getNumOverloadedDecls = &types.CallInterface{}
var cif_clang_getOverloadedDecl = &types.CallInterface{}
var cif_clang_getIBOutletCollectionType = &types.CallInterface{}
var cif_clang_visitChildren = &types.CallInterface{}
var cif_clang_visitChildrenWithBlock = &types.CallInterface{}
var cif_clang_getCursorUSR = &types.CallInterface{}
var cif_clang_getCursorSpelling = &types.CallInterface{}
var cif_clang_Cursor_getSpellingNameRange = &types.CallInterface{}
var cif_clang_getCursorPrintingPolicy = &types.CallInterface{}
var cif_clang_PrintingPolicy_dispose = &types.CallInterface{}
var cif_clang_getCursorPrettyPrinted = &types.CallInterface{}
var cif_clang_getTypePrettyPrinted = &types.CallInterface{}
var cif_clang_getFullyQualifiedName = &types.CallInterface{}
var cif_clang_getCursorDisplayName = &types.CallInterface{}
var cif_clang_getCursorReferenced = &types.CallInterface{}
var cif_clang_getCursorDefinition = &types.CallInterface{}
var cif_clang_isCursorDefinition = &types.CallInterface{}
var cif_clang_getCanonicalCursor = &types.CallInterface{}
var cif_clang_Cursor_getObjCSelectorIndex = &types.CallInterface{}
var cif_clang_Cursor_isDynamicCall = &types.CallInterface{}
var cif_clang_Cursor_getReceiverType = &types.CallInterface{}
var cif_clang_Cursor_getObjCPropertyAttributes = &types.CallInterface{}
var cif_clang_Cursor_getObjCPropertyGetterName = &types.CallInterface{}
var cif_clang_Cursor_getObjCPropertySetterName = &types.CallInterface{}
var cif_clang_Cursor_getObjCDeclQualifiers = &types.CallInterface{}
var cif_clang_Cursor_isObjCOptional = &types.CallInterface{}
var cif_clang_Cursor_isVariadic = &types.CallInterface{}
var cif_clang_Cursor_getCommentRange = &types.CallInterface{}
var cif_clang_Cursor_getRawCommentText = &types.CallInterface{}
var cif_clang_Cursor_getBriefCommentText = &types.CallInterface{}
var cif_clang_Cursor_getMangling = &types.CallInterface{}
var cif_clang_Cursor_getGCCAssemblyTemplate = &types.CallInterface{}
var cif_clang_Cursor_isGCCAssemblyHasGoto = &types.CallInterface{}
var cif_clang_Cursor_getGCCAssemblyNumOutputs = &types.CallInterface{}
var cif_clang_Cursor_getGCCAssemblyNumInputs = &types.CallInterface{}
var cif_clang_Cursor_getGCCAssemblyNumClobbers = &types.CallInterface{}
var cif_clang_Cursor_getGCCAssemblyClobber = &types.CallInterface{}
var cif_clang_Cursor_isGCCAssemblyVolatile = &types.CallInterface{}
var cif_clang_Cursor_getModule = &types.CallInterface{}
var cif_clang_getModuleForFile = &types.CallInterface{}
var cif_clang_Module_getASTFile = &types.CallInterface{}
var cif_clang_Module_getParent = &types.CallInterface{}
var cif_clang_Module_getName = &types.CallInterface{}
var cif_clang_Module_getFullName = &types.CallInterface{}
var cif_clang_Module_isSystem = &types.CallInterface{}
var cif_clang_Module_getNumTopLevelHeaders = &types.CallInterface{}
var cif_clang_Module_getTopLevelHeader = &types.CallInterface{}
var cif_clang_CXXConstructor_isConvertingConstructor = &types.CallInterface{}
var cif_clang_CXXConstructor_isCopyConstructor = &types.CallInterface{}
var cif_clang_CXXConstructor_isDefaultConstructor = &types.CallInterface{}
var cif_clang_CXXConstructor_isMoveConstructor = &types.CallInterface{}
var cif_clang_CXXField_isMutable = &types.CallInterface{}
var cif_clang_CXXMethod_isDefaulted = &types.CallInterface{}
var cif_clang_CXXMethod_isDeleted = &types.CallInterface{}
var cif_clang_CXXMethod_isPureVirtual = &types.CallInterface{}
var cif_clang_CXXMethod_isStatic = &types.CallInterface{}
var cif_clang_CXXMethod_isVirtual = &types.CallInterface{}
var cif_clang_CXXMethod_isCopyAssignmentOperator = &types.CallInterface{}
var cif_clang_CXXMethod_isMoveAssignmentOperator = &types.CallInterface{}
var cif_clang_CXXMethod_isExplicit = &types.CallInterface{}
var cif_clang_CXXRecord_isAbstract = &types.CallInterface{}
var cif_clang_EnumDecl_isScoped = &types.CallInterface{}
var cif_clang_CXXMethod_isConst = &types.CallInterface{}
var cif_clang_getTemplateCursorKind = &types.CallInterface{}
var cif_clang_getSpecializedCursorTemplate = &types.CallInterface{}
var cif_clang_getCursorReferenceNameRange = &types.CallInterface{}
var cif_clang_getTokenKind = &types.CallInterface{}
var cif_clang_getTokenSpelling = &types.CallInterface{}
var cif_clang_getTokenLocation = &types.CallInterface{}
var cif_clang_getTokenExtent = &types.CallInterface{}
var cif_clang_enableStackTraces = &types.CallInterface{}
var cif_clang_getCompletionChunkKind = &types.CallInterface{}
var cif_clang_getCompletionChunkText = &types.CallInterface{}
var cif_clang_getCompletionChunkCompletionString = &types.CallInterface{}
var cif_clang_getNumCompletionChunks = &types.CallInterface{}
var cif_clang_getCompletionPriority = &types.CallInterface{}
var cif_clang_getCompletionAvailability = &types.CallInterface{}
var cif_clang_getCompletionNumAnnotations = &types.CallInterface{}
var cif_clang_getCompletionAnnotation = &types.CallInterface{}
var cif_clang_getCompletionBriefComment = &types.CallInterface{}
var cif_clang_getCursorCompletionString = &types.CallInterface{}
var cif_clang_defaultCodeCompleteOptions = &types.CallInterface{}
var cif_clang_getClangVersion = &types.CallInterface{}
var cif_clang_toggleCrashRecovery = &types.CallInterface{}
var cif_clang_getInclusions = &types.CallInterface{}
var cif_clang_Cursor_Evaluate = &types.CallInterface{}
var cif_clang_EvalResult_getKind = &types.CallInterface{}
var cif_clang_EvalResult_getAsInt = &types.CallInterface{}
var cif_clang_EvalResult_isUnsignedInt = &types.CallInterface{}
var cif_clang_EvalResult_dispose = &types.CallInterface{}
var cif_clang_findReferencesInFile = &types.CallInterface{}
var cif_clang_findIncludesInFile = &types.CallInterface{}
var cif_clang_findReferencesInFileWithBlock = &types.CallInterface{}
var cif_clang_findIncludesInFileWithBlock = &types.CallInterface{}
var cif_clang_IndexAction_create = &types.CallInterface{}
var cif_clang_IndexAction_dispose = &types.CallInterface{}
var cif_clang_indexLoc_getCXSourceLocation = &types.CallInterface{}
var cif_clang_Type_visitFields = &types.CallInterface{}
var cif_clang_visitCXXBaseClasses = &types.CallInterface{}
var cif_clang_visitCXXMethods = &types.CallInterface{}
var cif_clang_getCursorBinaryOperatorKind = &types.CallInterface{}
var cif_clang_getCursorUnaryOperatorKind = &types.CallInterface{}
var cif_clang_remap_getNumFiles = &types.CallInterface{}
var cif_clang_remap_dispose = &types.CallInterface{}

var ptr_clang_disposeString unsafe.Pointer
var ptr_clang_VirtualFileOverlay_create unsafe.Pointer
var ptr_clang_VirtualFileOverlay_setCaseSensitivity unsafe.Pointer
var ptr_clang_VirtualFileOverlay_dispose unsafe.Pointer
var ptr_clang_ModuleMapDescriptor_create unsafe.Pointer
var ptr_clang_ModuleMapDescriptor_dispose unsafe.Pointer
var ptr_clang_getFileName unsafe.Pointer
var ptr_clang_File_isEqual unsafe.Pointer
var ptr_clang_File_tryGetRealPathName unsafe.Pointer
var ptr_clang_getNullLocation unsafe.Pointer
var ptr_clang_equalLocations unsafe.Pointer
var ptr_clang_isBeforeInTranslationUnit unsafe.Pointer
var ptr_clang_Location_isInSystemHeader unsafe.Pointer
var ptr_clang_Location_isFromMainFile unsafe.Pointer
var ptr_clang_getNullRange unsafe.Pointer
var ptr_clang_getRange unsafe.Pointer
var ptr_clang_equalRanges unsafe.Pointer
var ptr_clang_Range_isNull unsafe.Pointer
var ptr_clang_getRangeStart unsafe.Pointer
var ptr_clang_getRangeEnd unsafe.Pointer
var ptr_clang_getNumDiagnosticsInSet unsafe.Pointer
var ptr_clang_getDiagnosticInSet unsafe.Pointer
var ptr_clang_disposeDiagnosticSet unsafe.Pointer
var ptr_clang_getChildDiagnostics unsafe.Pointer
var ptr_clang_disposeDiagnostic unsafe.Pointer
var ptr_clang_formatDiagnostic unsafe.Pointer
var ptr_clang_defaultDiagnosticDisplayOptions unsafe.Pointer
var ptr_clang_getDiagnosticSeverity unsafe.Pointer
var ptr_clang_getDiagnosticLocation unsafe.Pointer
var ptr_clang_getDiagnosticSpelling unsafe.Pointer
var ptr_clang_getDiagnosticCategory unsafe.Pointer
var ptr_clang_getDiagnosticCategoryName unsafe.Pointer
var ptr_clang_getDiagnosticCategoryText unsafe.Pointer
var ptr_clang_getDiagnosticNumRanges unsafe.Pointer
var ptr_clang_getDiagnosticRange unsafe.Pointer
var ptr_clang_getDiagnosticNumFixIts unsafe.Pointer
var ptr_clang_createIndex unsafe.Pointer
var ptr_clang_disposeIndex unsafe.Pointer
var ptr_clang_CXIndex_setGlobalOptions unsafe.Pointer
var ptr_clang_CXIndex_getGlobalOptions unsafe.Pointer
var ptr_clang_isFileMultipleIncludeGuarded unsafe.Pointer
var ptr_clang_getLocation unsafe.Pointer
var ptr_clang_getLocationForOffset unsafe.Pointer
var ptr_clang_getNumDiagnostics unsafe.Pointer
var ptr_clang_getDiagnostic unsafe.Pointer
var ptr_clang_getDiagnosticSetFromTU unsafe.Pointer
var ptr_clang_getTranslationUnitSpelling unsafe.Pointer
var ptr_clang_defaultEditingTranslationUnitOptions unsafe.Pointer
var ptr_clang_defaultSaveOptions unsafe.Pointer
var ptr_clang_suspendTranslationUnit unsafe.Pointer
var ptr_clang_disposeTranslationUnit unsafe.Pointer
var ptr_clang_defaultReparseOptions unsafe.Pointer
var ptr_clang_getCXTUResourceUsage unsafe.Pointer
var ptr_clang_disposeCXTUResourceUsage unsafe.Pointer
var ptr_clang_getTranslationUnitTargetInfo unsafe.Pointer
var ptr_clang_TargetInfo_dispose unsafe.Pointer
var ptr_clang_TargetInfo_getTriple unsafe.Pointer
var ptr_clang_TargetInfo_getPointerWidth unsafe.Pointer
var ptr_clang_getNullCursor unsafe.Pointer
var ptr_clang_getTranslationUnitCursor unsafe.Pointer
var ptr_clang_equalCursors unsafe.Pointer
var ptr_clang_Cursor_isNull unsafe.Pointer
var ptr_clang_hashCursor unsafe.Pointer
var ptr_clang_getCursorKind unsafe.Pointer
var ptr_clang_isInvalidDeclaration unsafe.Pointer
var ptr_clang_Cursor_hasAttrs unsafe.Pointer
var ptr_clang_getCursorLinkage unsafe.Pointer
var ptr_clang_getCursorVisibility unsafe.Pointer
var ptr_clang_getCursorAvailability unsafe.Pointer
var ptr_clang_Cursor_getVarDeclInitializer unsafe.Pointer
var ptr_clang_Cursor_hasVarDeclGlobalStorage unsafe.Pointer
var ptr_clang_Cursor_hasVarDeclExternalStorage unsafe.Pointer
var ptr_clang_getCursorLanguage unsafe.Pointer
var ptr_clang_getCursorTLSKind unsafe.Pointer
var ptr_clang_Cursor_getTranslationUnit unsafe.Pointer
var ptr_clang_createCXCursorSet unsafe.Pointer
var ptr_clang_disposeCXCursorSet unsafe.Pointer
var ptr_clang_CXCursorSet_contains unsafe.Pointer
var ptr_clang_CXCursorSet_insert unsafe.Pointer
var ptr_clang_getCursorSemanticParent unsafe.Pointer
var ptr_clang_getCursorLexicalParent unsafe.Pointer
var ptr_clang_getIncludedFile unsafe.Pointer
var ptr_clang_getCursor unsafe.Pointer
var ptr_clang_getCursorLocation unsafe.Pointer
var ptr_clang_getCursorExtent unsafe.Pointer
var ptr_clang_getCursorType unsafe.Pointer
var ptr_clang_getTypeSpelling unsafe.Pointer
var ptr_clang_getTypedefDeclUnderlyingType unsafe.Pointer
var ptr_clang_getEnumDeclIntegerType unsafe.Pointer
var ptr_clang_Cursor_isBitField unsafe.Pointer
var ptr_clang_getFieldDeclBitWidth unsafe.Pointer
var ptr_clang_Cursor_getNumArguments unsafe.Pointer
var ptr_clang_Cursor_getArgument unsafe.Pointer
var ptr_clang_Cursor_getNumTemplateArguments unsafe.Pointer
var ptr_clang_Cursor_getTemplateArgumentKind unsafe.Pointer
var ptr_clang_Cursor_getTemplateArgumentType unsafe.Pointer
var ptr_clang_equalTypes unsafe.Pointer
var ptr_clang_getCanonicalType unsafe.Pointer
var ptr_clang_isConstQualifiedType unsafe.Pointer
var ptr_clang_Cursor_isMacroFunctionLike unsafe.Pointer
var ptr_clang_Cursor_isMacroBuiltin unsafe.Pointer
var ptr_clang_Cursor_isFunctionInlined unsafe.Pointer
var ptr_clang_isVolatileQualifiedType unsafe.Pointer
var ptr_clang_isRestrictQualifiedType unsafe.Pointer
var ptr_clang_getAddressSpace unsafe.Pointer
var ptr_clang_getTypedefName unsafe.Pointer
var ptr_clang_getPointeeType unsafe.Pointer
var ptr_clang_getUnqualifiedType unsafe.Pointer
var ptr_clang_getNonReferenceType unsafe.Pointer
var ptr_clang_getTypeDeclaration unsafe.Pointer
var ptr_clang_getDeclObjCTypeEncoding unsafe.Pointer
var ptr_clang_Type_getObjCEncoding unsafe.Pointer
var ptr_clang_getFunctionTypeCallingConv unsafe.Pointer
var ptr_clang_getResultType unsafe.Pointer
var ptr_clang_getExceptionSpecificationType unsafe.Pointer
var ptr_clang_getNumArgTypes unsafe.Pointer
var ptr_clang_getArgType unsafe.Pointer
var ptr_clang_Type_getObjCObjectBaseType unsafe.Pointer
var ptr_clang_Type_getNumObjCProtocolRefs unsafe.Pointer
var ptr_clang_Type_getObjCProtocolDecl unsafe.Pointer
var ptr_clang_Type_getNumObjCTypeArgs unsafe.Pointer
var ptr_clang_Type_getObjCTypeArg unsafe.Pointer
var ptr_clang_isFunctionTypeVariadic unsafe.Pointer
var ptr_clang_getCursorResultType unsafe.Pointer
var ptr_clang_getCursorExceptionSpecificationType unsafe.Pointer
var ptr_clang_isPODType unsafe.Pointer
var ptr_clang_getElementType unsafe.Pointer
var ptr_clang_getArrayElementType unsafe.Pointer
var ptr_clang_Type_getNamedType unsafe.Pointer
var ptr_clang_Type_isTransparentTagTypedef unsafe.Pointer
var ptr_clang_Type_getNullability unsafe.Pointer
var ptr_clang_Type_getClassType unsafe.Pointer
var ptr_clang_Type_getModifiedType unsafe.Pointer
var ptr_clang_Type_getValueType unsafe.Pointer
var ptr_clang_Cursor_isAnonymous unsafe.Pointer
var ptr_clang_Cursor_isAnonymousRecordDecl unsafe.Pointer
var ptr_clang_Cursor_isInlineNamespace unsafe.Pointer
var ptr_clang_Type_getNumTemplateArguments unsafe.Pointer
var ptr_clang_Type_getTemplateArgumentAsType unsafe.Pointer
var ptr_clang_Type_getCXXRefQualifier unsafe.Pointer
var ptr_clang_isVirtualBase unsafe.Pointer
var ptr_clang_getCXXAccessSpecifier unsafe.Pointer
var ptr_clang_Cursor_getBinaryOpcode unsafe.Pointer
var ptr_clang_Cursor_getStorageClass unsafe.Pointer
var ptr_clang_getNumOverloadedDecls unsafe.Pointer
var ptr_clang_getOverloadedDecl unsafe.Pointer
var ptr_clang_getIBOutletCollectionType unsafe.Pointer
var ptr_clang_visitChildren unsafe.Pointer
var ptr_clang_visitChildrenWithBlock unsafe.Pointer
var ptr_clang_getCursorUSR unsafe.Pointer
var ptr_clang_getCursorSpelling unsafe.Pointer
var ptr_clang_Cursor_getSpellingNameRange unsafe.Pointer
var ptr_clang_getCursorPrintingPolicy unsafe.Pointer
var ptr_clang_PrintingPolicy_dispose unsafe.Pointer
var ptr_clang_getCursorPrettyPrinted unsafe.Pointer
var ptr_clang_getTypePrettyPrinted unsafe.Pointer
var ptr_clang_getFullyQualifiedName unsafe.Pointer
var ptr_clang_getCursorDisplayName unsafe.Pointer
var ptr_clang_getCursorReferenced unsafe.Pointer
var ptr_clang_getCursorDefinition unsafe.Pointer
var ptr_clang_isCursorDefinition unsafe.Pointer
var ptr_clang_getCanonicalCursor unsafe.Pointer
var ptr_clang_Cursor_getObjCSelectorIndex unsafe.Pointer
var ptr_clang_Cursor_isDynamicCall unsafe.Pointer
var ptr_clang_Cursor_getReceiverType unsafe.Pointer
var ptr_clang_Cursor_getObjCPropertyAttributes unsafe.Pointer
var ptr_clang_Cursor_getObjCPropertyGetterName unsafe.Pointer
var ptr_clang_Cursor_getObjCPropertySetterName unsafe.Pointer
var ptr_clang_Cursor_getObjCDeclQualifiers unsafe.Pointer
var ptr_clang_Cursor_isObjCOptional unsafe.Pointer
var ptr_clang_Cursor_isVariadic unsafe.Pointer
var ptr_clang_Cursor_getCommentRange unsafe.Pointer
var ptr_clang_Cursor_getRawCommentText unsafe.Pointer
var ptr_clang_Cursor_getBriefCommentText unsafe.Pointer
var ptr_clang_Cursor_getMangling unsafe.Pointer
var ptr_clang_Cursor_getGCCAssemblyTemplate unsafe.Pointer
var ptr_clang_Cursor_isGCCAssemblyHasGoto unsafe.Pointer
var ptr_clang_Cursor_getGCCAssemblyNumOutputs unsafe.Pointer
var ptr_clang_Cursor_getGCCAssemblyNumInputs unsafe.Pointer
var ptr_clang_Cursor_getGCCAssemblyNumClobbers unsafe.Pointer
var ptr_clang_Cursor_getGCCAssemblyClobber unsafe.Pointer
var ptr_clang_Cursor_isGCCAssemblyVolatile unsafe.Pointer
var ptr_clang_Cursor_getModule unsafe.Pointer
var ptr_clang_getModuleForFile unsafe.Pointer
var ptr_clang_Module_getASTFile unsafe.Pointer
var ptr_clang_Module_getParent unsafe.Pointer
var ptr_clang_Module_getName unsafe.Pointer
var ptr_clang_Module_getFullName unsafe.Pointer
var ptr_clang_Module_isSystem unsafe.Pointer
var ptr_clang_Module_getNumTopLevelHeaders unsafe.Pointer
var ptr_clang_Module_getTopLevelHeader unsafe.Pointer
var ptr_clang_CXXConstructor_isConvertingConstructor unsafe.Pointer
var ptr_clang_CXXConstructor_isCopyConstructor unsafe.Pointer
var ptr_clang_CXXConstructor_isDefaultConstructor unsafe.Pointer
var ptr_clang_CXXConstructor_isMoveConstructor unsafe.Pointer
var ptr_clang_CXXField_isMutable unsafe.Pointer
var ptr_clang_CXXMethod_isDefaulted unsafe.Pointer
var ptr_clang_CXXMethod_isDeleted unsafe.Pointer
var ptr_clang_CXXMethod_isPureVirtual unsafe.Pointer
var ptr_clang_CXXMethod_isStatic unsafe.Pointer
var ptr_clang_CXXMethod_isVirtual unsafe.Pointer
var ptr_clang_CXXMethod_isCopyAssignmentOperator unsafe.Pointer
var ptr_clang_CXXMethod_isMoveAssignmentOperator unsafe.Pointer
var ptr_clang_CXXMethod_isExplicit unsafe.Pointer
var ptr_clang_CXXRecord_isAbstract unsafe.Pointer
var ptr_clang_EnumDecl_isScoped unsafe.Pointer
var ptr_clang_CXXMethod_isConst unsafe.Pointer
var ptr_clang_getTemplateCursorKind unsafe.Pointer
var ptr_clang_getSpecializedCursorTemplate unsafe.Pointer
var ptr_clang_getCursorReferenceNameRange unsafe.Pointer
var ptr_clang_getTokenKind unsafe.Pointer
var ptr_clang_getTokenSpelling unsafe.Pointer
var ptr_clang_getTokenLocation unsafe.Pointer
var ptr_clang_getTokenExtent unsafe.Pointer
var ptr_clang_enableStackTraces unsafe.Pointer
var ptr_clang_getCompletionChunkKind unsafe.Pointer
var ptr_clang_getCompletionChunkText unsafe.Pointer
var ptr_clang_getCompletionChunkCompletionString unsafe.Pointer
var ptr_clang_getNumCompletionChunks unsafe.Pointer
var ptr_clang_getCompletionPriority unsafe.Pointer
var ptr_clang_getCompletionAvailability unsafe.Pointer
var ptr_clang_getCompletionNumAnnotations unsafe.Pointer
var ptr_clang_getCompletionAnnotation unsafe.Pointer
var ptr_clang_getCompletionBriefComment unsafe.Pointer
var ptr_clang_getCursorCompletionString unsafe.Pointer
var ptr_clang_defaultCodeCompleteOptions unsafe.Pointer
var ptr_clang_getClangVersion unsafe.Pointer
var ptr_clang_toggleCrashRecovery unsafe.Pointer
var ptr_clang_getInclusions unsafe.Pointer
var ptr_clang_Cursor_Evaluate unsafe.Pointer
var ptr_clang_EvalResult_getKind unsafe.Pointer
var ptr_clang_EvalResult_getAsInt unsafe.Pointer
var ptr_clang_EvalResult_isUnsignedInt unsafe.Pointer
var ptr_clang_EvalResult_dispose unsafe.Pointer
var ptr_clang_findReferencesInFile unsafe.Pointer
var ptr_clang_findIncludesInFile unsafe.Pointer
var ptr_clang_findReferencesInFileWithBlock unsafe.Pointer
var ptr_clang_findIncludesInFileWithBlock unsafe.Pointer
var ptr_clang_IndexAction_create unsafe.Pointer
var ptr_clang_IndexAction_dispose unsafe.Pointer
var ptr_clang_indexLoc_getCXSourceLocation unsafe.Pointer
var ptr_clang_Type_visitFields unsafe.Pointer
var ptr_clang_visitCXXBaseClasses unsafe.Pointer
var ptr_clang_visitCXXMethods unsafe.Pointer
var ptr_clang_getCursorBinaryOperatorKind unsafe.Pointer
var ptr_clang_getCursorUnaryOperatorKind unsafe.Pointer
var ptr_clang_remap_getNumFiles unsafe.Pointer
var ptr_clang_remap_dispose unsafe.Pointer

func init() {
	library := lib.LoadLibrary(lib.LibraryPaths{
		Darwin: "libclang.dylib",
		Linux:  "libclang.so",
	})

	var err error

	{
		ptr_clang_disposeString, err = ffi.GetSymbol(library, "clang_disposeString")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				string_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeString, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_VirtualFileOverlay_create, err = ffi.GetSymbol(library, "clang_VirtualFileOverlay_create")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_VirtualFileOverlay_create, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_VirtualFileOverlay_setCaseSensitivity, err = ffi.GetSymbol(library, "clang_VirtualFileOverlay_setCaseSensitivity")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.SInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_VirtualFileOverlay_setCaseSensitivity, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_VirtualFileOverlay_dispose, err = ffi.GetSymbol(library, "clang_VirtualFileOverlay_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_VirtualFileOverlay_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_ModuleMapDescriptor_create, err = ffi.GetSymbol(library, "clang_ModuleMapDescriptor_create")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_ModuleMapDescriptor_create, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_ModuleMapDescriptor_dispose, err = ffi.GetSymbol(library, "clang_ModuleMapDescriptor_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_ModuleMapDescriptor_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getFileName, err = ffi.GetSymbol(library, "clang_getFileName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getFileName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_File_isEqual, err = ffi.GetSymbol(library, "clang_File_isEqual")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_File_isEqual, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_File_tryGetRealPathName, err = ffi.GetSymbol(library, "clang_File_tryGetRealPathName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_File_tryGetRealPathName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNullLocation, err = ffi.GetSymbol(library, "clang_getNullLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getNullLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_equalLocations, err = ffi.GetSymbol(library, "clang_equalLocations")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceLocationTypeDescriptor,
				sourceLocationTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_equalLocations, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isBeforeInTranslationUnit, err = ffi.GetSymbol(library, "clang_isBeforeInTranslationUnit")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceLocationTypeDescriptor,
				sourceLocationTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isBeforeInTranslationUnit, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Location_isInSystemHeader, err = ffi.GetSymbol(library, "clang_Location_isInSystemHeader")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceLocationTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Location_isInSystemHeader, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Location_isFromMainFile, err = ffi.GetSymbol(library, "clang_Location_isFromMainFile")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceLocationTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Location_isFromMainFile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNullRange, err = ffi.GetSymbol(library, "clang_getNullRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getNullRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getRange, err = ffi.GetSymbol(library, "clang_getRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceLocationTypeDescriptor,
				sourceLocationTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_equalRanges, err = ffi.GetSymbol(library, "clang_equalRanges")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceRangeTypeDescriptor,
				sourceRangeTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_equalRanges, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Range_isNull, err = ffi.GetSymbol(library, "clang_Range_isNull")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceRangeTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Range_isNull, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getRangeStart, err = ffi.GetSymbol(library, "clang_getRangeStart")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceRangeTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getRangeStart, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getRangeEnd, err = ffi.GetSymbol(library, "clang_getRangeEnd")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				sourceRangeTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getRangeEnd, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNumDiagnosticsInSet, err = ffi.GetSymbol(library, "clang_getNumDiagnosticsInSet")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getNumDiagnosticsInSet, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticInSet, err = ffi.GetSymbol(library, "clang_getDiagnosticInSet")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticInSet, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_disposeDiagnosticSet, err = ffi.GetSymbol(library, "clang_disposeDiagnosticSet")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeDiagnosticSet, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getChildDiagnostics, err = ffi.GetSymbol(library, "clang_getChildDiagnostics")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getChildDiagnostics, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_disposeDiagnostic, err = ffi.GetSymbol(library, "clang_disposeDiagnostic")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeDiagnostic, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_formatDiagnostic, err = ffi.GetSymbol(library, "clang_formatDiagnostic")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_formatDiagnostic, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_defaultDiagnosticDisplayOptions, err = ffi.GetSymbol(library, "clang_defaultDiagnosticDisplayOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_defaultDiagnosticDisplayOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticSeverity, err = ffi.GetSymbol(library, "clang_getDiagnosticSeverity")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticSeverity, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticLocation, err = ffi.GetSymbol(library, "clang_getDiagnosticLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticSpelling, err = ffi.GetSymbol(library, "clang_getDiagnosticSpelling")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticSpelling, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticCategory, err = ffi.GetSymbol(library, "clang_getDiagnosticCategory")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticCategory, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticCategoryName, err = ffi.GetSymbol(library, "clang_getDiagnosticCategoryName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticCategoryName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticCategoryText, err = ffi.GetSymbol(library, "clang_getDiagnosticCategoryText")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticCategoryText, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticNumRanges, err = ffi.GetSymbol(library, "clang_getDiagnosticNumRanges")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticNumRanges, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticRange, err = ffi.GetSymbol(library, "clang_getDiagnosticRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticNumFixIts, err = ffi.GetSymbol(library, "clang_getDiagnosticNumFixIts")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticNumFixIts, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_createIndex, err = ffi.GetSymbol(library, "clang_createIndex")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.SInt32TypeDescriptor,
				types.SInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_createIndex, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_disposeIndex, err = ffi.GetSymbol(library, "clang_disposeIndex")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeIndex, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXIndex_setGlobalOptions, err = ffi.GetSymbol(library, "clang_CXIndex_setGlobalOptions")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXIndex_setGlobalOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXIndex_getGlobalOptions, err = ffi.GetSymbol(library, "clang_CXIndex_getGlobalOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXIndex_getGlobalOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isFileMultipleIncludeGuarded, err = ffi.GetSymbol(library, "clang_isFileMultipleIncludeGuarded")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isFileMultipleIncludeGuarded, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getLocation, err = ffi.GetSymbol(library, "clang_getLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getLocationForOffset, err = ffi.GetSymbol(library, "clang_getLocationForOffset")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getLocationForOffset, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNumDiagnostics, err = ffi.GetSymbol(library, "clang_getNumDiagnostics")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getNumDiagnostics, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnostic, err = ffi.GetSymbol(library, "clang_getDiagnostic")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnostic, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDiagnosticSetFromTU, err = ffi.GetSymbol(library, "clang_getDiagnosticSetFromTU")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDiagnosticSetFromTU, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTranslationUnitSpelling, err = ffi.GetSymbol(library, "clang_getTranslationUnitSpelling")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTranslationUnitSpelling, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_defaultEditingTranslationUnitOptions, err = ffi.GetSymbol(library, "clang_defaultEditingTranslationUnitOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_defaultEditingTranslationUnitOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_defaultSaveOptions, err = ffi.GetSymbol(library, "clang_defaultSaveOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_defaultSaveOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_suspendTranslationUnit, err = ffi.GetSymbol(library, "clang_suspendTranslationUnit")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_suspendTranslationUnit, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_disposeTranslationUnit, err = ffi.GetSymbol(library, "clang_disposeTranslationUnit")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeTranslationUnit, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_defaultReparseOptions, err = ffi.GetSymbol(library, "clang_defaultReparseOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_defaultReparseOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCXTUResourceUsage, err = ffi.GetSymbol(library, "clang_getCXTUResourceUsage")
		if err == nil {
			returnType := tUResourceUsageTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCXTUResourceUsage, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_disposeCXTUResourceUsage, err = ffi.GetSymbol(library, "clang_disposeCXTUResourceUsage")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				tUResourceUsageTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeCXTUResourceUsage, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTranslationUnitTargetInfo, err = ffi.GetSymbol(library, "clang_getTranslationUnitTargetInfo")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTranslationUnitTargetInfo, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_TargetInfo_dispose, err = ffi.GetSymbol(library, "clang_TargetInfo_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_TargetInfo_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_TargetInfo_getTriple, err = ffi.GetSymbol(library, "clang_TargetInfo_getTriple")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_TargetInfo_getTriple, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_TargetInfo_getPointerWidth, err = ffi.GetSymbol(library, "clang_TargetInfo_getPointerWidth")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_TargetInfo_getPointerWidth, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNullCursor, err = ffi.GetSymbol(library, "clang_getNullCursor")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getNullCursor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTranslationUnitCursor, err = ffi.GetSymbol(library, "clang_getTranslationUnitCursor")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTranslationUnitCursor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_equalCursors, err = ffi.GetSymbol(library, "clang_equalCursors")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_equalCursors, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isNull, err = ffi.GetSymbol(library, "clang_Cursor_isNull")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isNull, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_hashCursor, err = ffi.GetSymbol(library, "clang_hashCursor")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_hashCursor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorKind, err = ffi.GetSymbol(library, "clang_getCursorKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isInvalidDeclaration, err = ffi.GetSymbol(library, "clang_isInvalidDeclaration")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isInvalidDeclaration, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_hasAttrs, err = ffi.GetSymbol(library, "clang_Cursor_hasAttrs")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_hasAttrs, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorLinkage, err = ffi.GetSymbol(library, "clang_getCursorLinkage")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorLinkage, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorVisibility, err = ffi.GetSymbol(library, "clang_getCursorVisibility")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorVisibility, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorAvailability, err = ffi.GetSymbol(library, "clang_getCursorAvailability")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorAvailability, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getVarDeclInitializer, err = ffi.GetSymbol(library, "clang_Cursor_getVarDeclInitializer")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getVarDeclInitializer, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_hasVarDeclGlobalStorage, err = ffi.GetSymbol(library, "clang_Cursor_hasVarDeclGlobalStorage")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_hasVarDeclGlobalStorage, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_hasVarDeclExternalStorage, err = ffi.GetSymbol(library, "clang_Cursor_hasVarDeclExternalStorage")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_hasVarDeclExternalStorage, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorLanguage, err = ffi.GetSymbol(library, "clang_getCursorLanguage")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorLanguage, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorTLSKind, err = ffi.GetSymbol(library, "clang_getCursorTLSKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorTLSKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getTranslationUnit, err = ffi.GetSymbol(library, "clang_Cursor_getTranslationUnit")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getTranslationUnit, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_createCXCursorSet, err = ffi.GetSymbol(library, "clang_createCXCursorSet")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_createCXCursorSet, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_disposeCXCursorSet, err = ffi.GetSymbol(library, "clang_disposeCXCursorSet")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_disposeCXCursorSet, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXCursorSet_contains, err = ffi.GetSymbol(library, "clang_CXCursorSet_contains")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXCursorSet_contains, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXCursorSet_insert, err = ffi.GetSymbol(library, "clang_CXCursorSet_insert")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXCursorSet_insert, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorSemanticParent, err = ffi.GetSymbol(library, "clang_getCursorSemanticParent")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorSemanticParent, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorLexicalParent, err = ffi.GetSymbol(library, "clang_getCursorLexicalParent")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorLexicalParent, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getIncludedFile, err = ffi.GetSymbol(library, "clang_getIncludedFile")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getIncludedFile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursor, err = ffi.GetSymbol(library, "clang_getCursor")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				sourceLocationTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorLocation, err = ffi.GetSymbol(library, "clang_getCursorLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorExtent, err = ffi.GetSymbol(library, "clang_getCursorExtent")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorExtent, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorType, err = ffi.GetSymbol(library, "clang_getCursorType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTypeSpelling, err = ffi.GetSymbol(library, "clang_getTypeSpelling")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTypeSpelling, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTypedefDeclUnderlyingType, err = ffi.GetSymbol(library, "clang_getTypedefDeclUnderlyingType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTypedefDeclUnderlyingType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getEnumDeclIntegerType, err = ffi.GetSymbol(library, "clang_getEnumDeclIntegerType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getEnumDeclIntegerType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isBitField, err = ffi.GetSymbol(library, "clang_Cursor_isBitField")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isBitField, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getFieldDeclBitWidth, err = ffi.GetSymbol(library, "clang_getFieldDeclBitWidth")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getFieldDeclBitWidth, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getNumArguments, err = ffi.GetSymbol(library, "clang_Cursor_getNumArguments")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getNumArguments, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getArgument, err = ffi.GetSymbol(library, "clang_Cursor_getArgument")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getArgument, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getNumTemplateArguments, err = ffi.GetSymbol(library, "clang_Cursor_getNumTemplateArguments")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getNumTemplateArguments, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getTemplateArgumentKind, err = ffi.GetSymbol(library, "clang_Cursor_getTemplateArgumentKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getTemplateArgumentKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getTemplateArgumentType, err = ffi.GetSymbol(library, "clang_Cursor_getTemplateArgumentType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getTemplateArgumentType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_equalTypes, err = ffi.GetSymbol(library, "clang_equalTypes")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_equalTypes, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCanonicalType, err = ffi.GetSymbol(library, "clang_getCanonicalType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCanonicalType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isConstQualifiedType, err = ffi.GetSymbol(library, "clang_isConstQualifiedType")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isConstQualifiedType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isMacroFunctionLike, err = ffi.GetSymbol(library, "clang_Cursor_isMacroFunctionLike")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isMacroFunctionLike, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isMacroBuiltin, err = ffi.GetSymbol(library, "clang_Cursor_isMacroBuiltin")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isMacroBuiltin, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isFunctionInlined, err = ffi.GetSymbol(library, "clang_Cursor_isFunctionInlined")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isFunctionInlined, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isVolatileQualifiedType, err = ffi.GetSymbol(library, "clang_isVolatileQualifiedType")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isVolatileQualifiedType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isRestrictQualifiedType, err = ffi.GetSymbol(library, "clang_isRestrictQualifiedType")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isRestrictQualifiedType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getAddressSpace, err = ffi.GetSymbol(library, "clang_getAddressSpace")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getAddressSpace, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTypedefName, err = ffi.GetSymbol(library, "clang_getTypedefName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTypedefName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getPointeeType, err = ffi.GetSymbol(library, "clang_getPointeeType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getPointeeType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getUnqualifiedType, err = ffi.GetSymbol(library, "clang_getUnqualifiedType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getUnqualifiedType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNonReferenceType, err = ffi.GetSymbol(library, "clang_getNonReferenceType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getNonReferenceType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTypeDeclaration, err = ffi.GetSymbol(library, "clang_getTypeDeclaration")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTypeDeclaration, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getDeclObjCTypeEncoding, err = ffi.GetSymbol(library, "clang_getDeclObjCTypeEncoding")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getDeclObjCTypeEncoding, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getObjCEncoding, err = ffi.GetSymbol(library, "clang_Type_getObjCEncoding")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getObjCEncoding, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getFunctionTypeCallingConv, err = ffi.GetSymbol(library, "clang_getFunctionTypeCallingConv")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getFunctionTypeCallingConv, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getResultType, err = ffi.GetSymbol(library, "clang_getResultType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getResultType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getExceptionSpecificationType, err = ffi.GetSymbol(library, "clang_getExceptionSpecificationType")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getExceptionSpecificationType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNumArgTypes, err = ffi.GetSymbol(library, "clang_getNumArgTypes")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getNumArgTypes, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getArgType, err = ffi.GetSymbol(library, "clang_getArgType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getArgType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getObjCObjectBaseType, err = ffi.GetSymbol(library, "clang_Type_getObjCObjectBaseType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getObjCObjectBaseType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getNumObjCProtocolRefs, err = ffi.GetSymbol(library, "clang_Type_getNumObjCProtocolRefs")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getNumObjCProtocolRefs, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getObjCProtocolDecl, err = ffi.GetSymbol(library, "clang_Type_getObjCProtocolDecl")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getObjCProtocolDecl, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getNumObjCTypeArgs, err = ffi.GetSymbol(library, "clang_Type_getNumObjCTypeArgs")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getNumObjCTypeArgs, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getObjCTypeArg, err = ffi.GetSymbol(library, "clang_Type_getObjCTypeArg")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getObjCTypeArg, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isFunctionTypeVariadic, err = ffi.GetSymbol(library, "clang_isFunctionTypeVariadic")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isFunctionTypeVariadic, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorResultType, err = ffi.GetSymbol(library, "clang_getCursorResultType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorResultType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorExceptionSpecificationType, err = ffi.GetSymbol(library, "clang_getCursorExceptionSpecificationType")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorExceptionSpecificationType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isPODType, err = ffi.GetSymbol(library, "clang_isPODType")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isPODType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getElementType, err = ffi.GetSymbol(library, "clang_getElementType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getElementType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getArrayElementType, err = ffi.GetSymbol(library, "clang_getArrayElementType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getArrayElementType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getNamedType, err = ffi.GetSymbol(library, "clang_Type_getNamedType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getNamedType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_isTransparentTagTypedef, err = ffi.GetSymbol(library, "clang_Type_isTransparentTagTypedef")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_isTransparentTagTypedef, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getNullability, err = ffi.GetSymbol(library, "clang_Type_getNullability")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getNullability, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getClassType, err = ffi.GetSymbol(library, "clang_Type_getClassType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getClassType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getModifiedType, err = ffi.GetSymbol(library, "clang_Type_getModifiedType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getModifiedType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getValueType, err = ffi.GetSymbol(library, "clang_Type_getValueType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getValueType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isAnonymous, err = ffi.GetSymbol(library, "clang_Cursor_isAnonymous")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isAnonymous, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isAnonymousRecordDecl, err = ffi.GetSymbol(library, "clang_Cursor_isAnonymousRecordDecl")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isAnonymousRecordDecl, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isInlineNamespace, err = ffi.GetSymbol(library, "clang_Cursor_isInlineNamespace")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isInlineNamespace, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getNumTemplateArguments, err = ffi.GetSymbol(library, "clang_Type_getNumTemplateArguments")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getNumTemplateArguments, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getTemplateArgumentAsType, err = ffi.GetSymbol(library, "clang_Type_getTemplateArgumentAsType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getTemplateArgumentAsType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_getCXXRefQualifier, err = ffi.GetSymbol(library, "clang_Type_getCXXRefQualifier")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_getCXXRefQualifier, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isVirtualBase, err = ffi.GetSymbol(library, "clang_isVirtualBase")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isVirtualBase, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCXXAccessSpecifier, err = ffi.GetSymbol(library, "clang_getCXXAccessSpecifier")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCXXAccessSpecifier, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getBinaryOpcode, err = ffi.GetSymbol(library, "clang_Cursor_getBinaryOpcode")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getBinaryOpcode, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getStorageClass, err = ffi.GetSymbol(library, "clang_Cursor_getStorageClass")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getStorageClass, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNumOverloadedDecls, err = ffi.GetSymbol(library, "clang_getNumOverloadedDecls")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getNumOverloadedDecls, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getOverloadedDecl, err = ffi.GetSymbol(library, "clang_getOverloadedDecl")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getOverloadedDecl, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getIBOutletCollectionType, err = ffi.GetSymbol(library, "clang_getIBOutletCollectionType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getIBOutletCollectionType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_visitChildren, err = ffi.GetSymbol(library, "clang_visitChildren")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_visitChildren, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_visitChildrenWithBlock, err = ffi.GetSymbol(library, "clang_visitChildrenWithBlock")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_visitChildrenWithBlock, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorUSR, err = ffi.GetSymbol(library, "clang_getCursorUSR")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorUSR, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorSpelling, err = ffi.GetSymbol(library, "clang_getCursorSpelling")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorSpelling, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getSpellingNameRange, err = ffi.GetSymbol(library, "clang_Cursor_getSpellingNameRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getSpellingNameRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorPrintingPolicy, err = ffi.GetSymbol(library, "clang_getCursorPrintingPolicy")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorPrintingPolicy, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_PrintingPolicy_dispose, err = ffi.GetSymbol(library, "clang_PrintingPolicy_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_PrintingPolicy_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorPrettyPrinted, err = ffi.GetSymbol(library, "clang_getCursorPrettyPrinted")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorPrettyPrinted, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTypePrettyPrinted, err = ffi.GetSymbol(library, "clang_getTypePrettyPrinted")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTypePrettyPrinted, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getFullyQualifiedName, err = ffi.GetSymbol(library, "clang_getFullyQualifiedName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getFullyQualifiedName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorDisplayName, err = ffi.GetSymbol(library, "clang_getCursorDisplayName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorDisplayName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorReferenced, err = ffi.GetSymbol(library, "clang_getCursorReferenced")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorReferenced, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorDefinition, err = ffi.GetSymbol(library, "clang_getCursorDefinition")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorDefinition, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_isCursorDefinition, err = ffi.GetSymbol(library, "clang_isCursorDefinition")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_isCursorDefinition, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCanonicalCursor, err = ffi.GetSymbol(library, "clang_getCanonicalCursor")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCanonicalCursor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getObjCSelectorIndex, err = ffi.GetSymbol(library, "clang_Cursor_getObjCSelectorIndex")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getObjCSelectorIndex, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isDynamicCall, err = ffi.GetSymbol(library, "clang_Cursor_isDynamicCall")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isDynamicCall, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getReceiverType, err = ffi.GetSymbol(library, "clang_Cursor_getReceiverType")
		if err == nil {
			returnType := type_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getReceiverType, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getObjCPropertyAttributes, err = ffi.GetSymbol(library, "clang_Cursor_getObjCPropertyAttributes")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getObjCPropertyAttributes, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getObjCPropertyGetterName, err = ffi.GetSymbol(library, "clang_Cursor_getObjCPropertyGetterName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getObjCPropertyGetterName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getObjCPropertySetterName, err = ffi.GetSymbol(library, "clang_Cursor_getObjCPropertySetterName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getObjCPropertySetterName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getObjCDeclQualifiers, err = ffi.GetSymbol(library, "clang_Cursor_getObjCDeclQualifiers")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getObjCDeclQualifiers, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isObjCOptional, err = ffi.GetSymbol(library, "clang_Cursor_isObjCOptional")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isObjCOptional, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isVariadic, err = ffi.GetSymbol(library, "clang_Cursor_isVariadic")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isVariadic, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getCommentRange, err = ffi.GetSymbol(library, "clang_Cursor_getCommentRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getCommentRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getRawCommentText, err = ffi.GetSymbol(library, "clang_Cursor_getRawCommentText")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getRawCommentText, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getBriefCommentText, err = ffi.GetSymbol(library, "clang_Cursor_getBriefCommentText")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getBriefCommentText, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getMangling, err = ffi.GetSymbol(library, "clang_Cursor_getMangling")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getMangling, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getGCCAssemblyTemplate, err = ffi.GetSymbol(library, "clang_Cursor_getGCCAssemblyTemplate")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getGCCAssemblyTemplate, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isGCCAssemblyHasGoto, err = ffi.GetSymbol(library, "clang_Cursor_isGCCAssemblyHasGoto")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isGCCAssemblyHasGoto, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getGCCAssemblyNumOutputs, err = ffi.GetSymbol(library, "clang_Cursor_getGCCAssemblyNumOutputs")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getGCCAssemblyNumOutputs, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getGCCAssemblyNumInputs, err = ffi.GetSymbol(library, "clang_Cursor_getGCCAssemblyNumInputs")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getGCCAssemblyNumInputs, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getGCCAssemblyNumClobbers, err = ffi.GetSymbol(library, "clang_Cursor_getGCCAssemblyNumClobbers")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getGCCAssemblyNumClobbers, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getGCCAssemblyClobber, err = ffi.GetSymbol(library, "clang_Cursor_getGCCAssemblyClobber")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getGCCAssemblyClobber, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_isGCCAssemblyVolatile, err = ffi.GetSymbol(library, "clang_Cursor_isGCCAssemblyVolatile")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_isGCCAssemblyVolatile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_getModule, err = ffi.GetSymbol(library, "clang_Cursor_getModule")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_getModule, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getModuleForFile, err = ffi.GetSymbol(library, "clang_getModuleForFile")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getModuleForFile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_getASTFile, err = ffi.GetSymbol(library, "clang_Module_getASTFile")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_getASTFile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_getParent, err = ffi.GetSymbol(library, "clang_Module_getParent")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_getParent, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_getName, err = ffi.GetSymbol(library, "clang_Module_getName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_getName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_getFullName, err = ffi.GetSymbol(library, "clang_Module_getFullName")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_getFullName, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_isSystem, err = ffi.GetSymbol(library, "clang_Module_isSystem")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_isSystem, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_getNumTopLevelHeaders, err = ffi.GetSymbol(library, "clang_Module_getNumTopLevelHeaders")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_getNumTopLevelHeaders, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Module_getTopLevelHeader, err = ffi.GetSymbol(library, "clang_Module_getTopLevelHeader")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Module_getTopLevelHeader, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXConstructor_isConvertingConstructor, err = ffi.GetSymbol(library, "clang_CXXConstructor_isConvertingConstructor")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXConstructor_isConvertingConstructor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXConstructor_isCopyConstructor, err = ffi.GetSymbol(library, "clang_CXXConstructor_isCopyConstructor")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXConstructor_isCopyConstructor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXConstructor_isDefaultConstructor, err = ffi.GetSymbol(library, "clang_CXXConstructor_isDefaultConstructor")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXConstructor_isDefaultConstructor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXConstructor_isMoveConstructor, err = ffi.GetSymbol(library, "clang_CXXConstructor_isMoveConstructor")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXConstructor_isMoveConstructor, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXField_isMutable, err = ffi.GetSymbol(library, "clang_CXXField_isMutable")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXField_isMutable, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isDefaulted, err = ffi.GetSymbol(library, "clang_CXXMethod_isDefaulted")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isDefaulted, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isDeleted, err = ffi.GetSymbol(library, "clang_CXXMethod_isDeleted")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isDeleted, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isPureVirtual, err = ffi.GetSymbol(library, "clang_CXXMethod_isPureVirtual")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isPureVirtual, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isStatic, err = ffi.GetSymbol(library, "clang_CXXMethod_isStatic")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isStatic, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isVirtual, err = ffi.GetSymbol(library, "clang_CXXMethod_isVirtual")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isVirtual, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isCopyAssignmentOperator, err = ffi.GetSymbol(library, "clang_CXXMethod_isCopyAssignmentOperator")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isCopyAssignmentOperator, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isMoveAssignmentOperator, err = ffi.GetSymbol(library, "clang_CXXMethod_isMoveAssignmentOperator")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isMoveAssignmentOperator, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isExplicit, err = ffi.GetSymbol(library, "clang_CXXMethod_isExplicit")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isExplicit, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXRecord_isAbstract, err = ffi.GetSymbol(library, "clang_CXXRecord_isAbstract")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXRecord_isAbstract, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_EnumDecl_isScoped, err = ffi.GetSymbol(library, "clang_EnumDecl_isScoped")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_EnumDecl_isScoped, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_CXXMethod_isConst, err = ffi.GetSymbol(library, "clang_CXXMethod_isConst")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_CXXMethod_isConst, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTemplateCursorKind, err = ffi.GetSymbol(library, "clang_getTemplateCursorKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTemplateCursorKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getSpecializedCursorTemplate, err = ffi.GetSymbol(library, "clang_getSpecializedCursorTemplate")
		if err == nil {
			returnType := cursorTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getSpecializedCursorTemplate, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorReferenceNameRange, err = ffi.GetSymbol(library, "clang_getCursorReferenceNameRange")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.UInt32TypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorReferenceNameRange, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTokenKind, err = ffi.GetSymbol(library, "clang_getTokenKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				tokenTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTokenKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTokenSpelling, err = ffi.GetSymbol(library, "clang_getTokenSpelling")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				tokenTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTokenSpelling, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTokenLocation, err = ffi.GetSymbol(library, "clang_getTokenLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				tokenTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTokenLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getTokenExtent, err = ffi.GetSymbol(library, "clang_getTokenExtent")
		if err == nil {
			returnType := sourceRangeTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				tokenTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getTokenExtent, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_enableStackTraces, err = ffi.GetSymbol(library, "clang_enableStackTraces")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_enableStackTraces, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionChunkKind, err = ffi.GetSymbol(library, "clang_getCompletionChunkKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionChunkKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionChunkText, err = ffi.GetSymbol(library, "clang_getCompletionChunkText")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionChunkText, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionChunkCompletionString, err = ffi.GetSymbol(library, "clang_getCompletionChunkCompletionString")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionChunkCompletionString, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getNumCompletionChunks, err = ffi.GetSymbol(library, "clang_getNumCompletionChunks")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getNumCompletionChunks, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionPriority, err = ffi.GetSymbol(library, "clang_getCompletionPriority")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionPriority, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionAvailability, err = ffi.GetSymbol(library, "clang_getCompletionAvailability")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionAvailability, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionNumAnnotations, err = ffi.GetSymbol(library, "clang_getCompletionNumAnnotations")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionNumAnnotations, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionAnnotation, err = ffi.GetSymbol(library, "clang_getCompletionAnnotation")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionAnnotation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCompletionBriefComment, err = ffi.GetSymbol(library, "clang_getCompletionBriefComment")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCompletionBriefComment, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorCompletionString, err = ffi.GetSymbol(library, "clang_getCursorCompletionString")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorCompletionString, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_defaultCodeCompleteOptions, err = ffi.GetSymbol(library, "clang_defaultCodeCompleteOptions")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_defaultCodeCompleteOptions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getClangVersion, err = ffi.GetSymbol(library, "clang_getClangVersion")
		if err == nil {
			returnType := string_TypeDescriptor
			argTypes := []*types.TypeDescriptor{}
			err = ffi.PrepareCallInterface(cif_clang_getClangVersion, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_toggleCrashRecovery, err = ffi.GetSymbol(library, "clang_toggleCrashRecovery")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.UInt32TypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_toggleCrashRecovery, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getInclusions, err = ffi.GetSymbol(library, "clang_getInclusions")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getInclusions, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Cursor_Evaluate, err = ffi.GetSymbol(library, "clang_Cursor_Evaluate")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Cursor_Evaluate, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_EvalResult_getKind, err = ffi.GetSymbol(library, "clang_EvalResult_getKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_EvalResult_getKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_EvalResult_getAsInt, err = ffi.GetSymbol(library, "clang_EvalResult_getAsInt")
		if err == nil {
			returnType := types.SInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_EvalResult_getAsInt, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_EvalResult_isUnsignedInt, err = ffi.GetSymbol(library, "clang_EvalResult_isUnsignedInt")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_EvalResult_isUnsignedInt, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_EvalResult_dispose, err = ffi.GetSymbol(library, "clang_EvalResult_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_EvalResult_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_findReferencesInFile, err = ffi.GetSymbol(library, "clang_findReferencesInFile")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.PointerTypeDescriptor,
				cursorAndRangeVisitorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_findReferencesInFile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_findIncludesInFile, err = ffi.GetSymbol(library, "clang_findIncludesInFile")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
				cursorAndRangeVisitorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_findIncludesInFile, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_findReferencesInFileWithBlock, err = ffi.GetSymbol(library, "clang_findReferencesInFileWithBlock")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_findReferencesInFileWithBlock, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_findIncludesInFileWithBlock, err = ffi.GetSymbol(library, "clang_findIncludesInFileWithBlock")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_findIncludesInFileWithBlock, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_IndexAction_create, err = ffi.GetSymbol(library, "clang_IndexAction_create")
		if err == nil {
			returnType := types.PointerTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_IndexAction_create, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_IndexAction_dispose, err = ffi.GetSymbol(library, "clang_IndexAction_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_IndexAction_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_indexLoc_getCXSourceLocation, err = ffi.GetSymbol(library, "clang_indexLoc_getCXSourceLocation")
		if err == nil {
			returnType := sourceLocationTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				idxLocTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_indexLoc_getCXSourceLocation, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_Type_visitFields, err = ffi.GetSymbol(library, "clang_Type_visitFields")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_Type_visitFields, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_visitCXXBaseClasses, err = ffi.GetSymbol(library, "clang_visitCXXBaseClasses")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_visitCXXBaseClasses, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_visitCXXMethods, err = ffi.GetSymbol(library, "clang_visitCXXMethods")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				type_TypeDescriptor,
				types.PointerTypeDescriptor,
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_visitCXXMethods, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorBinaryOperatorKind, err = ffi.GetSymbol(library, "clang_getCursorBinaryOperatorKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorBinaryOperatorKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_getCursorUnaryOperatorKind, err = ffi.GetSymbol(library, "clang_getCursorUnaryOperatorKind")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				cursorTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_getCursorUnaryOperatorKind, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_remap_getNumFiles, err = ffi.GetSymbol(library, "clang_remap_getNumFiles")
		if err == nil {
			returnType := types.UInt32TypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_remap_getNumFiles, types.DefaultCall, returnType, argTypes)
		}
	}

	{
		ptr_clang_remap_dispose, err = ffi.GetSymbol(library, "clang_remap_dispose")
		if err == nil {
			returnType := types.VoidTypeDescriptor
			argTypes := []*types.TypeDescriptor{
				types.PointerTypeDescriptor,
			}
			err = ffi.PrepareCallInterface(cif_clang_remap_dispose, types.DefaultCall, returnType, argTypes)
		}
	}

}
