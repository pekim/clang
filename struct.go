// This is a generated file. DO NOT EDIT.

package clang

type String struct {
	// data => const void *
	// private_flags => unsigned int
}

type StringSet struct {
	// Strings => CXString *
	// Count => unsigned int
}

type VirtualFileOverlayImpl struct{}

type ModuleMapDescriptorImpl struct{}

type FileUniqueID struct {
	// data => unsigned long long[3]
}

type SourceLocation struct {
	// ptr_data => const void *[2]
	// int_data => unsigned int
}

type SourceRange struct {
	// ptr_data => const void *[2]
	// begin_int_data => unsigned int
	// end_int_data => unsigned int
}

type SourceRangeList struct {
	// count => unsigned int
	// ranges => CXSourceRange *
}

type TargetInfoImpl struct{}

type TranslationUnitImpl struct{}

/*
Provides the contents of a file that has not yet been saved to disk.

Each CXUnsavedFile instance provides the name of a file on the system along with the current contents of that file that have not yet been saved to disk.
*/
type UnsavedFile struct {
	// Filename => const char *
	// Contents => const char *
	// Length => unsigned long
}

/*
Describes a version number of the form major.minor.subminor.
*/
type Version struct {
	// Major => int
	// Minor => int
	// Subminor => int
}

/*
Index initialization options.

0 is the default value of each member of this struct except for Size. Initialize the struct in one of the following three ways to avoid adapting code each time a new member is added to it:

or explicitly initialize the first data member and zero-initialize the rest:

or to prevent the -Wmissing-field-initializers warning for the above version:
*/
type IndexOptions struct {
	// Size => unsigned int
	// ThreadBackgroundPriorityForIndexing => unsigned char
	// ThreadBackgroundPriorityForEditing => unsigned char
	// ExcludeDeclarationsFromPCH => unsigned int
	// DisplayDiagnostics => unsigned int
	// StorePreamblesInMemory => unsigned int
	//
	//	=> unsigned int
	//
	// PreambleStoragePath => const char *
	// InvocationEmissionPath => const char *
}

type TUResourceUsageEntry struct {
	// kind => enum CXTUResourceUsageKind
	// amount => unsigned long
}

/*
The memory usage of a CXTranslationUnit, broken into categories.
*/
type TUResourceUsage struct {
	// data => void *
	// numEntries => unsigned int
	// entries => CXTUResourceUsageEntry *
}

/*
A cursor representing some element in the abstract syntax tree for a translation unit.

The cursor abstraction unifies the different kinds of entities in a program--declaration, statements, expressions, references to declarations, etc.--under a single "cursor" abstraction with a common set of operations. Common operation for a cursor include: getting the physical location in a source file where the cursor points, getting the name associated with a cursor, and retrieving cursors for any child nodes of a particular cursor.

Cursors can be produced in two specific ways. clang_getTranslationUnitCursor() produces a cursor for a translation unit, from which one can use clang_visitChildren() to explore the rest of the translation unit. clang_getCursor() maps from a physical source location to the entity that resides at that location, allowing one to map from the source code into the AST.
*/
type Cursor struct {
	// kind => enum CXCursorKind
	// xdata => int
	// data => const void *[3]
}

/*
Describes the availability of a given entity on a particular platform, e.g., a particular class might only be available on Mac OS 10.7 or newer.
*/
type PlatformAvailability struct {
	// Platform => CXString
	// Introduced => CXVersion
	// Deprecated => CXVersion
	// Obsoleted => CXVersion
	// Unavailable => int
	// Message => CXString
}

type CursorSetImpl struct{}

/*
The type of an element in the abstract syntax tree.
*/
type Type struct {
	// kind => enum CXTypeKind
	// data => void *[2]
}

/*
Describes a single preprocessing token.
*/
type Token struct {
	// int_data => unsigned int[4]
	// ptr_data => void *
}

/*
A single result of code completion.
*/
type CompletionResult struct {
	// CursorKind => enum CXCursorKind
	// CompletionString => CXCompletionString
}

/*
Contains the results of code-completion.

This data structure contains the results of code completion, as produced by clang_codeCompleteAt(). Its contents must be freed by clang_disposeCodeCompleteResults.
*/
type CodeCompleteResults struct {
	// Results => CXCompletionResult *
	// NumResults => unsigned int
}

type CursorAndRangeVisitor struct {
	// context => void *
	// visit => enum CXVisitorResult (*)(void *, CXCursor, CXSourceRange)
}

/*
Source location passed to index callbacks.
*/
type IdxLoc struct {
	// ptr_data => void *[2]
	// int_data => unsigned int
}

/*
Data for ppIncludedFile callback.
*/
type IdxIncludedFileInfo struct {
	// hashLoc => CXIdxLoc
	// filename => const char *
	// file => CXFile
	// isImport => int
	// isAngled => int
	// isModuleImport => int
}

/*
Data for IndexerCallbacks#importedASTFile.
*/
type IdxImportedASTFileInfo struct {
	// file => CXFile
	// module => CXModule
	// loc => CXIdxLoc
	// isImplicit => int
}

type IdxAttrInfo struct {
	// kind => CXIdxAttrKind
	// cursor => CXCursor
	// loc => CXIdxLoc
}

type IdxEntityInfo struct {
	// kind => CXIdxEntityKind
	// templateKind => CXIdxEntityCXXTemplateKind
	// lang => CXIdxEntityLanguage
	// name => const char *
	// USR => const char *
	// cursor => CXCursor
	// attributes => const CXIdxAttrInfo *const *
	// numAttributes => unsigned int
}

type IdxContainerInfo struct {
	// cursor => CXCursor
}

type IdxIBOutletCollectionAttrInfo struct {
	// attrInfo => const CXIdxAttrInfo *
	// objcClass => const CXIdxEntityInfo *
	// classCursor => CXCursor
	// classLoc => CXIdxLoc
}

type IdxDeclInfo struct {
	// entityInfo => const CXIdxEntityInfo *
	// cursor => CXCursor
	// loc => CXIdxLoc
	// semanticContainer => const CXIdxContainerInfo *
	// lexicalContainer => const CXIdxContainerInfo *
	// isRedeclaration => int
	// isDefinition => int
	// isContainer => int
	// declAsContainer => const CXIdxContainerInfo *
	// isImplicit => int
	// attributes => const CXIdxAttrInfo *const *
	// numAttributes => unsigned int
	// flags => unsigned int
}

type IdxObjCContainerDeclInfo struct {
	// declInfo => const CXIdxDeclInfo *
	// kind => CXIdxObjCContainerKind
}

type IdxBaseClassInfo struct {
	// base => const CXIdxEntityInfo *
	// cursor => CXCursor
	// loc => CXIdxLoc
}

type IdxObjCProtocolRefInfo struct {
	// protocol => const CXIdxEntityInfo *
	// cursor => CXCursor
	// loc => CXIdxLoc
}

type IdxObjCProtocolRefListInfo struct {
	// protocols => const CXIdxObjCProtocolRefInfo *const *
	// numProtocols => unsigned int
}

type IdxObjCInterfaceDeclInfo struct {
	// containerInfo => const CXIdxObjCContainerDeclInfo *
	// superInfo => const CXIdxBaseClassInfo *
	// protocols => const CXIdxObjCProtocolRefListInfo *
}

type IdxObjCCategoryDeclInfo struct {
	// containerInfo => const CXIdxObjCContainerDeclInfo *
	// objcClass => const CXIdxEntityInfo *
	// classCursor => CXCursor
	// classLoc => CXIdxLoc
	// protocols => const CXIdxObjCProtocolRefListInfo *
}

type IdxObjCPropertyDeclInfo struct {
	// declInfo => const CXIdxDeclInfo *
	// getter => const CXIdxEntityInfo *
	// setter => const CXIdxEntityInfo *
}

type IdxCXXClassDeclInfo struct {
	// declInfo => const CXIdxDeclInfo *
	// bases => const CXIdxBaseClassInfo *const *
	// numBases => unsigned int
}

/*
Data for IndexerCallbacks#indexEntityReference.
*/
type IdxEntityRefInfo struct {
	// kind => CXIdxEntityRefKind
	// cursor => CXCursor
	// loc => CXIdxLoc
	// referencedEntity => const CXIdxEntityInfo *
	// parentEntity => const CXIdxEntityInfo *
	// container => const CXIdxContainerInfo *
	// role => CXSymbolRole
}

/*
A group of callbacks used by #clang_indexSourceFile and #clang_indexTranslationUnit.
*/
type IndexerCallbacks struct {
	// abortQuery => int (*)(CXClientData, void *)
	// diagnostic => void (*)(CXClientData, CXDiagnosticSet, void *)
	// enteredMainFile => CXIdxClientFile (*)(CXClientData, CXFile, void *)
	// ppIncludedFile => CXIdxClientFile (*)(CXClientData, const CXIdxIncludedFileInfo *)
	// importedASTFile => CXIdxClientASTFile (*)(CXClientData, const CXIdxImportedASTFileInfo *)
	// startedTranslationUnit => CXIdxClientContainer (*)(CXClientData, void *)
	// indexDeclaration => void (*)(CXClientData, const CXIdxDeclInfo *)
	// indexEntityReference => void (*)(CXClientData, const CXIdxEntityRefInfo *)
}
