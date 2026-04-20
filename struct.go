// This is a generated file. DO NOT EDIT.

package clang

type String struct {
	// Data => const void *
	Private_flags uint32
}

type StringSet struct {
	// Strings => CXString *
	Count uint32
}

type VirtualFileOverlayImpl struct{}

type ModuleMapDescriptorImpl struct{}

type FileUniqueID struct {
	// Data => unsigned long long[3]
}

type SourceLocation struct {
	// Ptr_data => const void *[2]
	Int_data uint32
}

type SourceRange struct {
	// Ptr_data => const void *[2]
	Begin_int_data uint32
	End_int_data   uint32
}

type SourceRangeList struct {
	Count uint32
	// Ranges => CXSourceRange *
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
	/*
	   The length of the unsaved contents of this buffer.

	*/
	Length uint64
}

/*
Describes a version number of the form major.minor.subminor.
*/
type Version struct {
	/*
	   The major version number, e.g., the '10' in '10.7.3'. A negative value indicates that there is no version number at all.
	*/
	Major int32
	/*
	   The minor version number, e.g., the '7' in '10.7.3'. This value will be negative if no minor version number was provided, e.g., for version '10'.
	*/
	Minor int32
	/*
	   The subminor version number, e.g., the '3' in '10.7.3'. This value will be negative if no minor or subminor version number was provided, e.g., in version '10' or '10.7'.
	*/
	Subminor int32
}

/*
Index initialization options.

0 is the default value of each member of this struct except for Size. Initialize the struct in one of the following three ways to avoid adapting code each time a new member is added to it:

or explicitly initialize the first data member and zero-initialize the rest:

or to prevent the -Wmissing-field-initializers warning for the above version:
*/
type IndexOptions struct {
	/*
	   The size of struct CXIndexOptions used for option versioning.

	   Always initialize this member to sizeof(CXIndexOptions), or assign sizeof(CXIndexOptions) to it right after creating a CXIndexOptions object.
	*/
	Size uint32
	// ThreadBackgroundPriorityForIndexing => unsigned char
	// ThreadBackgroundPriorityForEditing => unsigned char
	/*


	 */
	ExcludeDeclarationsFromPCH uint32
	/*
	 */
	DisplayDiagnostics uint32
	/*
	   Store PCH in memory. If zero, PCH are stored in temporary files.
	*/
	StorePreamblesInMemory uint32
	_1                     uint32
	// PreambleStoragePath => const char *
	// InvocationEmissionPath => const char *
}

type TUResourceUsageEntry struct {
	// Kind => enum CXTUResourceUsageKind
	Amount uint64
}

/*
The memory usage of a CXTranslationUnit, broken into categories.
*/
type TUResourceUsage struct {
	// Data => void *
	NumEntries uint32
	// Entries => CXTUResourceUsageEntry *
}

/*
A cursor representing some element in the abstract syntax tree for a translation unit.

The cursor abstraction unifies the different kinds of entities in a program--declaration, statements, expressions, references to declarations, etc.--under a single "cursor" abstraction with a common set of operations. Common operation for a cursor include: getting the physical location in a source file where the cursor points, getting the name associated with a cursor, and retrieving cursors for any child nodes of a particular cursor.

Cursors can be produced in two specific ways. clang_getTranslationUnitCursor() produces a cursor for a translation unit, from which one can use clang_visitChildren() to explore the rest of the translation unit. clang_getCursor() maps from a physical source location to the entity that resides at that location, allowing one to map from the source code into the AST.
*/
type Cursor struct {
	// Kind => enum CXCursorKind
	Xdata int32
	// Data => const void *[3]
}

/*
Describes the availability of a given entity on a particular platform, e.g., a particular class might only be available on Mac OS 10.7 or newer.
*/
type PlatformAvailability struct {
	// Platform => CXString
	// Introduced => CXVersion
	// Deprecated => CXVersion
	// Obsoleted => CXVersion
	/*
	   Whether the entity is unconditionally unavailable on this platform.

	*/
	Unavailable int32
	// Message => CXString
}

type CursorSetImpl struct{}

/*
The type of an element in the abstract syntax tree.
*/
type Type struct {
	// Kind => enum CXTypeKind
	// Data => void *[2]
}

/*
Describes a single preprocessing token.
*/
type Token struct {
	// Int_data => unsigned int[4]
	// Ptr_data => void *
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
	/*
	   The number of code-completion results stored in the Results array.

	*/
	NumResults uint32
}

type CursorAndRangeVisitor struct {
	// Context => void *
	// Visit => enum CXVisitorResult (*)(void *, CXCursor, CXSourceRange)
}

/*
Source location passed to index callbacks.
*/
type IdxLoc struct {
	// Ptr_data => void *[2]
	Int_data uint32
}

/*
Data for ppIncludedFile callback.
*/
type IdxIncludedFileInfo struct {
	// HashLoc => CXIdxLoc
	// Filename => const char *
	// File => CXFile
	IsImport int32
	IsAngled int32
	/*
	   Non-zero if the directive was automatically turned into a module import.
	*/
	IsModuleImport int32
}

/*
Data for IndexerCallbacks#importedASTFile.
*/
type IdxImportedASTFileInfo struct {
	// File => CXFile
	// Module => CXModule
	// Loc => CXIdxLoc
	/*
	   Non-zero if an inclusion directive was automatically turned into a module import. Applicable only for modules.

	*/
	IsImplicit int32
}

type IdxAttrInfo struct {
	// Kind => CXIdxAttrKind
	// Cursor => CXCursor
	// Loc => CXIdxLoc
}

type IdxEntityInfo struct {
	// Kind => CXIdxEntityKind
	// TemplateKind => CXIdxEntityCXXTemplateKind
	// Lang => CXIdxEntityLanguage
	// Name => const char *
	// USR => const char *
	// Cursor => CXCursor
	// Attributes => const CXIdxAttrInfo *const *
	NumAttributes uint32
}

type IdxContainerInfo struct {
	// Cursor => CXCursor
}

type IdxIBOutletCollectionAttrInfo struct {
	// AttrInfo => const CXIdxAttrInfo *
	// ObjcClass => const CXIdxEntityInfo *
	// ClassCursor => CXCursor
	// ClassLoc => CXIdxLoc
}

type IdxDeclInfo struct {
	// EntityInfo => const CXIdxEntityInfo *
	// Cursor => CXCursor
	// Loc => CXIdxLoc
	// SemanticContainer => const CXIdxContainerInfo *
	// LexicalContainer => const CXIdxContainerInfo *
	IsRedeclaration int32
	IsDefinition    int32
	IsContainer     int32
	// DeclAsContainer => const CXIdxContainerInfo *
	/*
	   Whether the declaration exists in code or was created implicitly by the compiler, e.g. implicit Objective-C methods for properties.

	*/
	IsImplicit int32
	// Attributes => const CXIdxAttrInfo *const *
	NumAttributes uint32
	Flags         uint32
}

type IdxObjCContainerDeclInfo struct {
	// DeclInfo => const CXIdxDeclInfo *
	// Kind => CXIdxObjCContainerKind
}

type IdxBaseClassInfo struct {
	// Base => const CXIdxEntityInfo *
	// Cursor => CXCursor
	// Loc => CXIdxLoc
}

type IdxObjCProtocolRefInfo struct {
	// Protocol => const CXIdxEntityInfo *
	// Cursor => CXCursor
	// Loc => CXIdxLoc
}

type IdxObjCProtocolRefListInfo struct {
	// Protocols => const CXIdxObjCProtocolRefInfo *const *
	NumProtocols uint32
}

type IdxObjCInterfaceDeclInfo struct {
	// ContainerInfo => const CXIdxObjCContainerDeclInfo *
	// SuperInfo => const CXIdxBaseClassInfo *
	// Protocols => const CXIdxObjCProtocolRefListInfo *
}

type IdxObjCCategoryDeclInfo struct {
	// ContainerInfo => const CXIdxObjCContainerDeclInfo *
	// ObjcClass => const CXIdxEntityInfo *
	// ClassCursor => CXCursor
	// ClassLoc => CXIdxLoc
	// Protocols => const CXIdxObjCProtocolRefListInfo *
}

type IdxObjCPropertyDeclInfo struct {
	// DeclInfo => const CXIdxDeclInfo *
	// Getter => const CXIdxEntityInfo *
	// Setter => const CXIdxEntityInfo *
}

type IdxCXXClassDeclInfo struct {
	// DeclInfo => const CXIdxDeclInfo *
	// Bases => const CXIdxBaseClassInfo *const *
	NumBases uint32
}

/*
Data for IndexerCallbacks#indexEntityReference.
*/
type IdxEntityRefInfo struct {
	// Kind => CXIdxEntityRefKind
	// Cursor => CXCursor
	// Loc => CXIdxLoc
	// ReferencedEntity => const CXIdxEntityInfo *
	// ParentEntity => const CXIdxEntityInfo *
	// Container => const CXIdxContainerInfo *
	// Role => CXSymbolRole
}

/*
A group of callbacks used by #clang_indexSourceFile and #clang_indexTranslationUnit.
*/
type IndexerCallbacks struct {
	// AbortQuery => int (*)(CXClientData, void *)
	// Diagnostic => void (*)(CXClientData, CXDiagnosticSet, void *)
	// EnteredMainFile => CXIdxClientFile (*)(CXClientData, CXFile, void *)
	// PpIncludedFile => CXIdxClientFile (*)(CXClientData, const CXIdxIncludedFileInfo *)
	// ImportedASTFile => CXIdxClientASTFile (*)(CXClientData, const CXIdxImportedASTFileInfo *)
	// StartedTranslationUnit => CXIdxClientContainer (*)(CXClientData, void *)
	// IndexDeclaration => void (*)(CXClientData, const CXIdxDeclInfo *)
	// IndexEntityReference => void (*)(CXClientData, const CXIdxEntityRefInfo *)
}
