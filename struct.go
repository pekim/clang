// This is a generated file. DO NOT EDIT.

package clang

import "structs"

type String_ struct {
	_ structs.HostLayout

	data          uintptr // const void *
	Private_flags uint32
	_             [4]byte
}

type StringSet struct {
	_ structs.HostLayout

	strings uintptr // CXString *
	Count   uint32
	_       [4]byte
}

type VirtualFileOverlayImpl struct {
	_ structs.HostLayout
}

type ModuleMapDescriptorImpl struct {
	_ structs.HostLayout
}

type FileUniqueID struct {
	_ structs.HostLayout

	Data [24]byte // unsigned long long[3]
}

type SourceLocation struct {
	_ structs.HostLayout

	Ptr_data [16]byte // const void *[2]
	Int_data uint32
	_        [4]byte
}

type SourceRange struct {
	_ structs.HostLayout

	Ptr_data       [16]byte // const void *[2]
	Begin_int_data uint32
	End_int_data   uint32
}

type SourceRangeList struct {
	_ structs.HostLayout

	Count  uint32
	_      [4]byte
	ranges uintptr // CXSourceRange *
}

type TargetInfoImpl struct {
	_ structs.HostLayout
}

type TranslationUnitImpl struct {
	_ structs.HostLayout
}

/*
Provides the contents of a file that has not yet been saved to disk.

Each CXUnsavedFile instance provides the name of a file on the system along with the current contents of that file that have not yet been saved to disk.
*/
type UnsavedFile struct {
	_ structs.HostLayout

	/*
	   The file whose contents have not yet been saved.

	   This file must already exist in the file system.
	*/
	filename uintptr // const char *
	// A buffer containing the unsaved contents of this file.
	contents uintptr // const char *
	// The length of the unsaved contents of this buffer.
	Length uint64
}

// Describes a version number of the form major.minor.subminor.
type Version struct {
	_ structs.HostLayout

	// The major version number, e.g., the '10' in '10.7.3'. A negative value indicates that there is no version number at all.
	Major int32
	// The minor version number, e.g., the '7' in '10.7.3'. This value will be negative if no minor version number was provided, e.g., for version '10'.
	Minor int32
	// The subminor version number, e.g., the '3' in '10.7.3'. This value will be negative if no minor or subminor version number was provided, e.g., in version '10' or '10.7'.
	Subminor int32
}

/*
Index initialization options.

0 is the default value of each member of this struct except for Size. Initialize the struct in one of the following three ways to avoid adapting code each time a new member is added to it:

or explicitly initialize the first data member and zero-initialize the rest:

or to prevent the -Wmissing-field-initializers warning for the above version:
*/
type IndexOptions struct {
	_ structs.HostLayout

	/*
	   The size of struct CXIndexOptions used for option versioning.

	   Always initialize this member to sizeof(CXIndexOptions), or assign sizeof(CXIndexOptions) to it right after creating a CXIndexOptions object.
	*/
	Size uint32
	// A CXChoice enumerator that specifies the indexing priority policy.
	ThreadBackgroundPriorityForIndexing byte
	// A CXChoice enumerator that specifies the editing priority policy.
	ThreadBackgroundPriorityForEditing byte
	bitfield_0                         uint16
	/*
	   The path to a directory, in which to store temporary PCH files. If null or empty, the default system temporary directory is used. These PCH files are deleted on clean exit but stay on disk if the program crashes or is killed.

	   This option is ignored if StorePreamblesInMemory is non-zero.

	   Libclang does not create the directory at the specified path in the file system. Therefore it must exist, or storing PCH files will fail.
	*/
	preambleStoragePath uintptr // const char *
	// Specifies a path which will contain log files for certain libclang invocations. A null value implies that libclang invocations are not logged.
	invocationEmissionPath uintptr // const char *
}

func (s *IndexOptions) ExcludeDeclarationsFromPCH() uint {
	return bitfieldGet(s.bitfield_0, 0, 1)
}

func (s *IndexOptions) SetExcludeDeclarationsFromPCH(value uint) {
	bitfieldSet(s.bitfield_0, 0, 1, value)
}

func (s *IndexOptions) DisplayDiagnostics() uint {
	return bitfieldGet(s.bitfield_0, 1, 1)
}

func (s *IndexOptions) SetDisplayDiagnostics(value uint) {
	bitfieldSet(s.bitfield_0, 1, 1, value)
}

// Store PCH in memory. If zero, PCH are stored in temporary files.
func (s *IndexOptions) StorePreamblesInMemory() uint {
	return bitfieldGet(s.bitfield_0, 2, 1)
}

// Store PCH in memory. If zero, PCH are stored in temporary files.
func (s *IndexOptions) SetStorePreamblesInMemory(value uint) {
	bitfieldSet(s.bitfield_0, 2, 1, value)
}

type TUResourceUsageEntry struct {
	_ structs.HostLayout

	Kind   [4]byte // enum CXTUResourceUsageKind
	_      [4]byte
	Amount uint64
}

// The memory usage of a CXTranslationUnit, broken into categories.
type TUResourceUsage struct {
	_ structs.HostLayout

	data       uintptr // void *
	NumEntries uint32
	_          [4]byte
	entries    uintptr // CXTUResourceUsageEntry *
}

/*
A cursor representing some element in the abstract syntax tree for a translation unit.

The cursor abstraction unifies the different kinds of entities in a program--declaration, statements, expressions, references to declarations, etc.--under a single "cursor" abstraction with a common set of operations. Common operation for a cursor include: getting the physical location in a source file where the cursor points, getting the name associated with a cursor, and retrieving cursors for any child nodes of a particular cursor.

Cursors can be produced in two specific ways. clang_getTranslationUnitCursor() produces a cursor for a translation unit, from which one can use clang_visitChildren() to explore the rest of the translation unit. clang_getCursor() maps from a physical source location to the entity that resides at that location, allowing one to map from the source code into the AST.
*/
type Cursor struct {
	_ structs.HostLayout

	Kind  [4]byte // enum CXCursorKind
	Xdata int32
	Data  [24]byte // const void *[3]
}

// Describes the availability of a given entity on a particular platform, e.g., a particular class might only be available on Mac OS 10.7 or newer.
type PlatformAvailability struct {
	_ structs.HostLayout

	/*
	   A string that describes the platform for which this structure provides availability information.

	   Possible values are "ios" or "macos".
	*/
	Platform String_
	// The version number in which this entity was introduced.
	Introduced Version
	// The version number in which this entity was deprecated (but is still available).
	Deprecated Version
	// The version number in which this entity was obsoleted, and therefore is no longer available.
	Obsoleted Version
	// Whether the entity is unconditionally unavailable on this platform.
	Unavailable int32
	// An optional message to provide to a user of this API, e.g., to suggest replacement APIs.
	Message String_
}

type CursorSetImpl struct {
	_ structs.HostLayout
}

// The type of an element in the abstract syntax tree.
type Type_ struct {
	_ structs.HostLayout

	Kind [4]byte // enum CXTypeKind
	_    [4]byte
	Data [16]byte // void *[2]
}

// Describes a single preprocessing token.
type Token struct {
	_ structs.HostLayout

	Int_data [16]byte // unsigned int[4]
	ptr_data uintptr  // void *
}

// A single result of code completion.
type CompletionResult struct {
	_ structs.HostLayout

	/*
	   The kind of entity that this completion refers to.

	   The cursor kind will be a macro, keyword, or a declaration (one of the *Decl cursor kinds), describing the entity that the completion is referring to.
	*/
	CursorKind [4]byte // enum CXCursorKind
	_          [4]byte
	// The code-completion string that describes how to insert this code-completion result into the editing buffer.
	CompletionString [8]byte // CXCompletionString
}

/*
Contains the results of code-completion.

This data structure contains the results of code completion, as produced by clang_codeCompleteAt(). Its contents must be freed by clang_disposeCodeCompleteResults.
*/
type CodeCompleteResults struct {
	_ structs.HostLayout

	// The code-completion results.
	results uintptr // CXCompletionResult *
	// The number of code-completion results stored in the Results array.
	NumResults uint32
	_          [4]byte
}

type CursorAndRangeVisitor struct {
	_ structs.HostLayout

	context uintptr // void *
	visit   uintptr // enum CXVisitorResult (*)(void *, CXCursor, CXSourceRange)
}

// Source location passed to index callbacks.
type IdxLoc struct {
	_ structs.HostLayout

	Ptr_data [16]byte // void *[2]
	Int_data uint32
	_        [4]byte
}

// Data for ppIncludedFile callback.
type IdxIncludedFileInfo struct {
	_ structs.HostLayout

	// Location of '#' in the #include/#import directive.
	HashLoc IdxLoc
	// Filename as written in the #include/#import directive.
	filename uintptr // const char *
	// The actual file that the #include/#import directive resolved to.
	File     [8]byte // CXFile
	IsImport int32
	IsAngled int32
	// Non-zero if the directive was automatically turned into a module import.
	IsModuleImport int32
	_              [4]byte
}

// Data for IndexerCallbacks#importedASTFile.
type IdxImportedASTFileInfo struct {
	_ structs.HostLayout

	// Top level AST file containing the imported PCH, module or submodule.
	File [8]byte // CXFile
	// The imported module or NULL if the AST file is a PCH.
	Module [8]byte // CXModule
	// Location where the file is imported. Applicable only for modules.
	Loc IdxLoc
	// Non-zero if an inclusion directive was automatically turned into a module import. Applicable only for modules.
	IsImplicit int32
	_          [4]byte
}

type IdxAttrInfo struct {
	_ structs.HostLayout

	Kind   IdxAttrKind
	_      [4]byte
	Cursor Cursor
	Loc    IdxLoc
}

type IdxEntityInfo struct {
	_ structs.HostLayout

	Kind          IdxEntityKind
	TemplateKind  IdxEntityCXXTemplateKind
	Lang          IdxEntityLanguage
	_             [4]byte
	name          uintptr // const char *
	uSR           uintptr // const char *
	Cursor        Cursor
	attributes    uintptr // const CXIdxAttrInfo *const *
	NumAttributes uint32
	_             [4]byte
}

type IdxContainerInfo struct {
	_ structs.HostLayout

	Cursor Cursor
}

type IdxIBOutletCollectionAttrInfo struct {
	_ structs.HostLayout

	attrInfo    uintptr // const CXIdxAttrInfo *
	objcClass   uintptr // const CXIdxEntityInfo *
	ClassCursor Cursor
	ClassLoc    IdxLoc
}

type IdxDeclInfo struct {
	_ structs.HostLayout

	entityInfo        uintptr // const CXIdxEntityInfo *
	Cursor            Cursor
	Loc               IdxLoc
	semanticContainer uintptr // const CXIdxContainerInfo *
	// Generally same as #semanticContainer but can be different in cases like out-of-line C++ member functions.
	lexicalContainer uintptr // const CXIdxContainerInfo *
	IsRedeclaration  int32
	IsDefinition     int32
	IsContainer      int32
	_                [4]byte
	declAsContainer  uintptr // const CXIdxContainerInfo *
	// Whether the declaration exists in code or was created implicitly by the compiler, e.g. implicit Objective-C methods for properties.
	IsImplicit    int32
	_             [4]byte
	attributes    uintptr // const CXIdxAttrInfo *const *
	NumAttributes uint32
	Flags         uint32
}

type IdxObjCContainerDeclInfo struct {
	_ structs.HostLayout

	declInfo uintptr // const CXIdxDeclInfo *
	Kind     IdxObjCContainerKind
	_        [4]byte
}

type IdxBaseClassInfo struct {
	_ structs.HostLayout

	base   uintptr // const CXIdxEntityInfo *
	Cursor Cursor
	Loc    IdxLoc
}

type IdxObjCProtocolRefInfo struct {
	_ structs.HostLayout

	protocol uintptr // const CXIdxEntityInfo *
	Cursor   Cursor
	Loc      IdxLoc
}

type IdxObjCProtocolRefListInfo struct {
	_ structs.HostLayout

	protocols    uintptr // const CXIdxObjCProtocolRefInfo *const *
	NumProtocols uint32
	_            [4]byte
}

type IdxObjCInterfaceDeclInfo struct {
	_ structs.HostLayout

	containerInfo uintptr // const CXIdxObjCContainerDeclInfo *
	superInfo     uintptr // const CXIdxBaseClassInfo *
	protocols     uintptr // const CXIdxObjCProtocolRefListInfo *
}

type IdxObjCCategoryDeclInfo struct {
	_ structs.HostLayout

	containerInfo uintptr // const CXIdxObjCContainerDeclInfo *
	objcClass     uintptr // const CXIdxEntityInfo *
	ClassCursor   Cursor
	ClassLoc      IdxLoc
	protocols     uintptr // const CXIdxObjCProtocolRefListInfo *
}

type IdxObjCPropertyDeclInfo struct {
	_ structs.HostLayout

	declInfo uintptr // const CXIdxDeclInfo *
	getter   uintptr // const CXIdxEntityInfo *
	setter   uintptr // const CXIdxEntityInfo *
}

type IdxCXXClassDeclInfo struct {
	_ structs.HostLayout

	declInfo uintptr // const CXIdxDeclInfo *
	bases    uintptr // const CXIdxBaseClassInfo *const *
	NumBases uint32
	_        [4]byte
}

// Data for IndexerCallbacks#indexEntityReference.
type IdxEntityRefInfo struct {
	_ structs.HostLayout

	Kind IdxEntityRefKind
	_    [4]byte
	// Reference cursor.
	Cursor Cursor
	Loc    IdxLoc
	// The entity that gets referenced.
	referencedEntity uintptr // const CXIdxEntityInfo *
	/*
	   Immediate "parent" of the reference. For example:

	   The parent of reference of type 'Foo' is the variable 'var'. For references inside statement bodies of functions/methods, the parentEntity will be the function/method.
	*/
	parentEntity uintptr // const CXIdxEntityInfo *
	// Lexical container context of the reference.
	container uintptr // const CXIdxContainerInfo *
	// Sets of symbol roles of the reference.
	Role SymbolRole
	_    [4]byte
}

// A group of callbacks used by #clang_indexSourceFile and #clang_indexTranslationUnit.
type IndexerCallbacks struct {
	_ structs.HostLayout

	// Called periodically to check whether indexing should be aborted. Should return 0 to continue, and non-zero to abort.
	abortQuery uintptr // int (*)(CXClientData, void *)
	// Called at the end of indexing; passes the complete diagnostic set.
	diagnostic      uintptr // void (*)(CXClientData, CXDiagnosticSet, void *)
	enteredMainFile uintptr // CXIdxClientFile (*)(CXClientData, CXFile, void *)
	// Called when a file gets #included/#imported.
	ppIncludedFile uintptr // CXIdxClientFile (*)(CXClientData, const CXIdxIncludedFileInfo *)
	/*
	   Called when a AST file (PCH or module) gets imported.

	   AST files will not get indexed (there will not be callbacks to index all the entities in an AST file). The recommended action is that, if the AST file is not already indexed, to initiate a new indexing job specific to the AST file.
	*/
	importedASTFile uintptr // CXIdxClientASTFile (*)(CXClientData, const CXIdxImportedASTFileInfo *)
	// Called at the beginning of indexing a translation unit.
	startedTranslationUnit uintptr // CXIdxClientContainer (*)(CXClientData, void *)
	indexDeclaration       uintptr // void (*)(CXClientData, const CXIdxDeclInfo *)
	// Called to index a reference of an entity.
	indexEntityReference uintptr // void (*)(CXClientData, const CXIdxEntityRefInfo *)
}
