// This is a generated file. DO NOT EDIT.

package clang

import (
	"fmt"
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

// not supported : clang_getCString : return value : const char *

func DisposeString(string_ String_) {
	c_string_ := string_

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_string_),
	}

	err := ffi.CallFunction(cif_clang_disposeString, ptr_clang_disposeString, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeString", err))
	}
}

// not supported : clang_disposeStringSet : param set : CXStringSet *

// not supported : clang_getBuildSessionTimestamp : return value : unsigned long long

func VirtualFileOverlay_create(options uint32) VirtualFileOverlay {
	c_options := options

	var retC VirtualFileOverlay
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(cif_clang_VirtualFileOverlay_create, ptr_clang_VirtualFileOverlay_create, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_create", err))
	}

	ret := retC
	return ret
}

// not supported : clang_VirtualFileOverlay_addFileMapping : param virtualPath : const char *

func VirtualFileOverlay_setCaseSensitivity(p0 VirtualFileOverlay, caseSensitive int32) ErrorCode {
	c_p0 := p0
	c_caseSensitive := caseSensitive

	var retC ErrorCode
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_caseSensitive),
	}

	err := ffi.CallFunction(cif_clang_VirtualFileOverlay_setCaseSensitivity, ptr_clang_VirtualFileOverlay_setCaseSensitivity, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_setCaseSensitivity", err))
	}

	ret := retC
	return ret
}

// not supported : clang_VirtualFileOverlay_writeToBuffer : param out_buffer_ptr : char **

// not supported : clang_free : param buffer : void *

func VirtualFileOverlay_dispose(p0 VirtualFileOverlay) {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_VirtualFileOverlay_dispose, ptr_clang_VirtualFileOverlay_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_VirtualFileOverlay_dispose", err))
	}
}

func ModuleMapDescriptor_create(options uint32) ModuleMapDescriptor {
	c_options := options

	var retC ModuleMapDescriptor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(cif_clang_ModuleMapDescriptor_create, ptr_clang_ModuleMapDescriptor_create, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_create", err))
	}

	ret := retC
	return ret
}

// not supported : clang_ModuleMapDescriptor_setFrameworkModuleName : param name : const char *

// not supported : clang_ModuleMapDescriptor_setUmbrellaHeader : param name : const char *

// not supported : clang_ModuleMapDescriptor_writeToBuffer : param out_buffer_ptr : char **

func ModuleMapDescriptor_dispose(p0 ModuleMapDescriptor) {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_ModuleMapDescriptor_dispose, ptr_clang_ModuleMapDescriptor_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_ModuleMapDescriptor_dispose", err))
	}
}

func GetFileName(sFile File) String_ {
	c_sFile := sFile

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_sFile),
	}

	err := ffi.CallFunction(cif_clang_getFileName, ptr_clang_getFileName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFileName", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getFileTime : return value : time_t

// not supported : clang_getFileUniqueID : param outID : CXFileUniqueID *

func File_isEqual(file1 File, file2 File) int32 {
	c_file1 := file1
	c_file2 := file2

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_file1),
		unsafe.Pointer(&c_file2),
	}

	err := ffi.CallFunction(cif_clang_File_isEqual, ptr_clang_File_isEqual, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_File_isEqual", err))
	}

	ret := retC
	return ret
}

func File_tryGetRealPathName(file File) String_ {
	c_file := file

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_file),
	}

	err := ffi.CallFunction(cif_clang_File_tryGetRealPathName, ptr_clang_File_tryGetRealPathName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_File_tryGetRealPathName", err))
	}

	ret := retC
	return ret
}

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

func EqualLocations(loc1 SourceLocation, loc2 SourceLocation) uint32 {
	c_loc1 := loc1
	c_loc2 := loc2

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc1),
		unsafe.Pointer(&c_loc2),
	}

	err := ffi.CallFunction(cif_clang_equalLocations, ptr_clang_equalLocations, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalLocations", err))
	}

	ret := retC
	return ret
}

func IsBeforeInTranslationUnit(loc1 SourceLocation, loc2 SourceLocation) uint32 {
	c_loc1 := loc1
	c_loc2 := loc2

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc1),
		unsafe.Pointer(&c_loc2),
	}

	err := ffi.CallFunction(cif_clang_isBeforeInTranslationUnit, ptr_clang_isBeforeInTranslationUnit, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isBeforeInTranslationUnit", err))
	}

	ret := retC
	return ret
}

func Location_isInSystemHeader(location SourceLocation) int32 {
	c_location := location

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
	}

	err := ffi.CallFunction(cif_clang_Location_isInSystemHeader, ptr_clang_Location_isInSystemHeader, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Location_isInSystemHeader", err))
	}

	ret := retC
	return ret
}

func Location_isFromMainFile(location SourceLocation) int32 {
	c_location := location

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_location),
	}

	err := ffi.CallFunction(cif_clang_Location_isFromMainFile, ptr_clang_Location_isFromMainFile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Location_isFromMainFile", err))
	}

	ret := retC
	return ret
}

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

func GetRange(begin SourceLocation, end SourceLocation) SourceRange {
	c_begin := begin
	c_end := end

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_begin),
		unsafe.Pointer(&c_end),
	}

	err := ffi.CallFunction(cif_clang_getRange, ptr_clang_getRange, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRange", err))
	}

	ret := retC
	return ret
}

func EqualRanges(range1 SourceRange, range2 SourceRange) uint32 {
	c_range1 := range1
	c_range2 := range2

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range1),
		unsafe.Pointer(&c_range2),
	}

	err := ffi.CallFunction(cif_clang_equalRanges, ptr_clang_equalRanges, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalRanges", err))
	}

	ret := retC
	return ret
}

func Range_isNull(range_ SourceRange) int32 {
	c_range_ := range_

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(cif_clang_Range_isNull, ptr_clang_Range_isNull, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Range_isNull", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getExpansionLocation : param file : CXFile *

// not supported : clang_getPresumedLocation : param filename : CXString *

// not supported : clang_getInstantiationLocation : param file : CXFile *

// not supported : clang_getSpellingLocation : param file : CXFile *

// not supported : clang_getFileLocation : param file : CXFile *

func GetRangeStart(range_ SourceRange) SourceLocation {
	c_range_ := range_

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(cif_clang_getRangeStart, ptr_clang_getRangeStart, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRangeStart", err))
	}

	ret := retC
	return ret
}

func GetRangeEnd(range_ SourceRange) SourceLocation {
	c_range_ := range_

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(cif_clang_getRangeEnd, ptr_clang_getRangeEnd, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getRangeEnd", err))
	}

	ret := retC
	return ret
}

// not supported : clang_disposeSourceRangeList : param ranges : CXSourceRangeList *

func GetNumDiagnosticsInSet(diags DiagnosticSet) uint32 {
	c_diags := diags

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diags),
	}

	err := ffi.CallFunction(cif_clang_getNumDiagnosticsInSet, ptr_clang_getNumDiagnosticsInSet, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumDiagnosticsInSet", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticInSet(diags DiagnosticSet, index uint32) Diagnostic {
	c_diags := diags
	c_index := index

	var retC Diagnostic
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diags),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticInSet, ptr_clang_getDiagnosticInSet, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticInSet", err))
	}

	ret := retC
	return ret
}

// not supported : clang_loadDiagnostics : param file : const char *

func DisposeDiagnosticSet(diags DiagnosticSet) {
	c_diags := diags

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diags),
	}

	err := ffi.CallFunction(cif_clang_disposeDiagnosticSet, ptr_clang_disposeDiagnosticSet, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeDiagnosticSet", err))
	}
}

func GetChildDiagnostics(d Diagnostic) DiagnosticSet {
	c_d := d

	var retC DiagnosticSet
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_d),
	}

	err := ffi.CallFunction(cif_clang_getChildDiagnostics, ptr_clang_getChildDiagnostics, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getChildDiagnostics", err))
	}

	ret := retC
	return ret
}

func DisposeDiagnostic(diagnostic Diagnostic) {
	c_diagnostic := diagnostic

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic),
	}

	err := ffi.CallFunction(cif_clang_disposeDiagnostic, ptr_clang_disposeDiagnostic, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeDiagnostic", err))
	}
}

func FormatDiagnostic(diagnostic Diagnostic, options uint32) String_ {
	c_diagnostic := diagnostic
	c_options := options

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(cif_clang_formatDiagnostic, ptr_clang_formatDiagnostic, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_formatDiagnostic", err))
	}

	ret := retC
	return ret
}

func DefaultDiagnosticDisplayOptions() uint32 {
	var retC uint32
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_defaultDiagnosticDisplayOptions, ptr_clang_defaultDiagnosticDisplayOptions, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultDiagnosticDisplayOptions", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticSeverity(p0 Diagnostic) DiagnosticSeverity {
	c_p0 := p0

	var retC DiagnosticSeverity
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticSeverity, ptr_clang_getDiagnosticSeverity, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticSeverity", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticLocation(p0 Diagnostic) SourceLocation {
	c_p0 := p0

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticLocation, ptr_clang_getDiagnosticLocation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticLocation", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticSpelling(p0 Diagnostic) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticSpelling, ptr_clang_getDiagnosticSpelling, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticSpelling", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getDiagnosticOption : param Disable : CXString *

func GetDiagnosticCategory(p0 Diagnostic) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticCategory, ptr_clang_getDiagnosticCategory, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategory", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticCategoryName(category uint32) String_ {
	c_category := category

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_category),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticCategoryName, ptr_clang_getDiagnosticCategoryName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategoryName", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticCategoryText(p0 Diagnostic) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticCategoryText, ptr_clang_getDiagnosticCategoryText, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticCategoryText", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticNumRanges(p0 Diagnostic) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticNumRanges, ptr_clang_getDiagnosticNumRanges, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticNumRanges", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticRange(diagnostic Diagnostic, range_ uint32) SourceRange {
	c_diagnostic := diagnostic
	c_range_ := range_

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic),
		unsafe.Pointer(&c_range_),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticRange, ptr_clang_getDiagnosticRange, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticRange", err))
	}

	ret := retC
	return ret
}

func GetDiagnosticNumFixIts(diagnostic Diagnostic) uint32 {
	c_diagnostic := diagnostic

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_diagnostic),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticNumFixIts, ptr_clang_getDiagnosticNumFixIts, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticNumFixIts", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getDiagnosticFixIt : param ReplacementRange : CXSourceRange *

/*
Provides a shared context for creating translation units.

It provides two options:

- excludeDeclarationsFromPCH: When non-zero, allows enumeration of "local" declarations (when loading any new translation units). A "local" declaration is one that belongs in the translation unit itself and not in a precompiled header that was used by the translation unit. If zero, all declarations will be enumerated.

Here is an example:

This process of creating the 'pch', loading it separately, and using it (via -include-pch) allows 'excludeDeclsFromPCH' to remove redundant callbacks (which gives the indexer the same performance benefit as the compiler).
*/
func CreateIndex(excludeDeclarationsFromPCH int32, displayDiagnostics int32) Index {
	c_excludeDeclarationsFromPCH := excludeDeclarationsFromPCH
	c_displayDiagnostics := displayDiagnostics

	var retC Index
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_excludeDeclarationsFromPCH),
		unsafe.Pointer(&c_displayDiagnostics),
	}

	err := ffi.CallFunction(cif_clang_createIndex, ptr_clang_createIndex, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_createIndex", err))
	}

	ret := retC
	return ret
}

/*
Destroy the given index.

The index must not be destroyed until all of the translation units created within that index have been destroyed.
*/
func DisposeIndex(index Index) {
	c_index := index

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(cif_clang_disposeIndex, ptr_clang_disposeIndex, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeIndex", err))
	}
}

// not supported : clang_createIndexWithOptions : param options : const CXIndexOptions *

/*
Sets general options associated with a CXIndex.

This function is DEPRECATED. Set CXIndexOptions::ThreadBackgroundPriorityForIndexing and/or CXIndexOptions::ThreadBackgroundPriorityForEditing and call clang_createIndexWithOptions() instead.

For example:
*/
func Index_setGlobalOptions(p0 Index, options uint32) {
	c_p0 := p0
	c_options := options

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(cif_clang_CXIndex_setGlobalOptions, ptr_clang_CXIndex_setGlobalOptions, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXIndex_setGlobalOptions", err))
	}
}

/*
Gets the general options associated with a CXIndex.

This function allows to obtain the final option values used by libclang after specifying the option policies via CXChoice enumerators.
*/
func Index_getGlobalOptions(p0 Index) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_CXIndex_getGlobalOptions, ptr_clang_CXIndex_getGlobalOptions, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXIndex_getGlobalOptions", err))
	}

	ret := retC
	return ret
}

// not supported : clang_CXIndex_setInvocationEmissionPathOption : param Path : const char *

// Determine whether the given header is guarded against multiple inclusions, either with the conventional #ifndef/#define/#endif macro guards or with #pragma once.
func IsFileMultipleIncludeGuarded(tu TranslationUnit, file File) uint32 {
	c_tu := tu
	c_file := file

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu),
		unsafe.Pointer(&c_file),
	}

	err := ffi.CallFunction(cif_clang_isFileMultipleIncludeGuarded, ptr_clang_isFileMultipleIncludeGuarded, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isFileMultipleIncludeGuarded", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getFile : param file_name : const char *

// not supported : clang_getFileContents : return value : const char *

// Retrieves the source location associated with a given file/line/column in a particular translation unit.
func GetLocation(tu TranslationUnit, file File, line uint32, column uint32) SourceLocation {
	c_tu := tu
	c_file := file
	c_line := line
	c_column := column

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_line),
		unsafe.Pointer(&c_column),
	}

	err := ffi.CallFunction(cif_clang_getLocation, ptr_clang_getLocation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getLocation", err))
	}

	ret := retC
	return ret
}

// Retrieves the source location associated with a given character offset in a particular translation unit.
func GetLocationForOffset(tu TranslationUnit, file File, offset uint32) SourceLocation {
	c_tu := tu
	c_file := file
	c_offset := offset

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_offset),
	}

	err := ffi.CallFunction(cif_clang_getLocationForOffset, ptr_clang_getLocationForOffset, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getLocationForOffset", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getSkippedRanges : return value : CXSourceRangeList *

// not supported : clang_getAllSkippedRanges : return value : CXSourceRangeList *

// Determine the number of diagnostics produced for the given translation unit.
func GetNumDiagnostics(unit TranslationUnit) uint32 {
	c_unit := unit

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_unit),
	}

	err := ffi.CallFunction(cif_clang_getNumDiagnostics, ptr_clang_getNumDiagnostics, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumDiagnostics", err))
	}

	ret := retC
	return ret
}

// Retrieve a diagnostic associated with the given translation unit.
func GetDiagnostic(unit TranslationUnit, index uint32) Diagnostic {
	c_unit := unit
	c_index := index

	var retC Diagnostic
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_unit),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(cif_clang_getDiagnostic, ptr_clang_getDiagnostic, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnostic", err))
	}

	ret := retC
	return ret
}

// Retrieve the complete set of diagnostics associated with a        translation unit.
func GetDiagnosticSetFromTU(unit TranslationUnit) DiagnosticSet {
	c_unit := unit

	var retC DiagnosticSet
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_unit),
	}

	err := ffi.CallFunction(cif_clang_getDiagnosticSetFromTU, ptr_clang_getDiagnosticSetFromTU, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDiagnosticSetFromTU", err))
	}

	ret := retC
	return ret
}

// Get the original translation unit source file name.
func GetTranslationUnitSpelling(cTUnit TranslationUnit) String_ {
	c_cTUnit := cTUnit

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cTUnit),
	}

	err := ffi.CallFunction(cif_clang_getTranslationUnitSpelling, ptr_clang_getTranslationUnitSpelling, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTranslationUnitSpelling", err))
	}

	ret := retC
	return ret
}

// not supported : clang_createTranslationUnitFromSourceFile : param source_filename : const char *

// not supported : clang_createTranslationUnit : param ast_filename : const char *

// not supported : clang_createTranslationUnit2 : param ast_filename : const char *

/*
Returns the set of flags that is suitable for parsing a translation unit that is being edited.

The set of flags returned provide options for clang_parseTranslationUnit() to indicate that the translation unit is likely to be reparsed many times, either explicitly (via clang_reparseTranslationUnit()) or implicitly (e.g., by code completion (clang_codeCompletionAt())). The returned flag set contains an unspecified set of optimizations (e.g., the precompiled preamble) geared toward improving the performance of these routines. The set of optimizations enabled may change from one version to the next.
*/
func DefaultEditingTranslationUnitOptions() uint32 {
	var retC uint32
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_defaultEditingTranslationUnitOptions, ptr_clang_defaultEditingTranslationUnitOptions, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultEditingTranslationUnitOptions", err))
	}

	ret := retC
	return ret
}

// not supported : clang_parseTranslationUnit : param source_filename : const char *

// not supported : clang_parseTranslationUnit2 : param source_filename : const char *

// not supported : clang_parseTranslationUnit2FullArgv : param source_filename : const char *

/*
Returns the set of flags that is suitable for saving a translation unit.

The set of flags returned provide options for clang_saveTranslationUnit() by default. The returned flag set contains an unspecified set of options that save translation units with the most commonly-requested data.
*/
func DefaultSaveOptions(tU TranslationUnit) uint32 {
	c_tU := tU

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU),
	}

	err := ffi.CallFunction(cif_clang_defaultSaveOptions, ptr_clang_defaultSaveOptions, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultSaveOptions", err))
	}

	ret := retC
	return ret
}

// not supported : clang_saveTranslationUnit : param FileName : const char *

/*
Suspend a translation unit in order to free memory associated with it.

A suspended translation unit uses significantly less memory but on the other side does not support any other calls than clang_reparseTranslationUnit to resume it or clang_disposeTranslationUnit to dispose it completely.
*/
func SuspendTranslationUnit(p0 TranslationUnit) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_suspendTranslationUnit, ptr_clang_suspendTranslationUnit, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_suspendTranslationUnit", err))
	}

	ret := retC
	return ret
}

// Destroy the specified CXTranslationUnit object.
func DisposeTranslationUnit(p0 TranslationUnit) {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_disposeTranslationUnit, ptr_clang_disposeTranslationUnit, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeTranslationUnit", err))
	}
}

/*
Returns the set of flags that is suitable for reparsing a translation unit.

The set of flags returned provide options for clang_reparseTranslationUnit() by default. The returned flag set contains an unspecified set of optimizations geared toward common uses of reparsing. The set of optimizations enabled may change from one version to the next.
*/
func DefaultReparseOptions(tU TranslationUnit) uint32 {
	c_tU := tU

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU),
	}

	err := ffi.CallFunction(cif_clang_defaultReparseOptions, ptr_clang_defaultReparseOptions, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultReparseOptions", err))
	}

	ret := retC
	return ret
}

// not supported : clang_reparseTranslationUnit : param unsaved_files : struct CXUnsavedFile *

// not supported : clang_getTUResourceUsageName : return value : const char *

// Return the memory usage of a translation unit.  This object  should be released with clang_disposeCXTUResourceUsage().
func GetCXTUResourceUsage(tU TranslationUnit) TUResourceUsage {
	c_tU := tU

	var retC TUResourceUsage
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU),
	}

	err := ffi.CallFunction(cif_clang_getCXTUResourceUsage, ptr_clang_getCXTUResourceUsage, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCXTUResourceUsage", err))
	}

	ret := retC
	return ret
}

func DisposeCXTUResourceUsage(usage TUResourceUsage) {
	c_usage := usage

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_usage),
	}

	err := ffi.CallFunction(cif_clang_disposeCXTUResourceUsage, ptr_clang_disposeCXTUResourceUsage, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeCXTUResourceUsage", err))
	}
}

/*
Get target information for this translation unit.

The CXTargetInfo object cannot outlive the CXTranslationUnit object.
*/
func GetTranslationUnitTargetInfo(cTUnit TranslationUnit) TargetInfo {
	c_cTUnit := cTUnit

	var retC TargetInfo
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cTUnit),
	}

	err := ffi.CallFunction(cif_clang_getTranslationUnitTargetInfo, ptr_clang_getTranslationUnitTargetInfo, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTranslationUnitTargetInfo", err))
	}

	ret := retC
	return ret
}

// Destroy the CXTargetInfo object.
func TargetInfo_dispose(info TargetInfo) {
	c_info := info

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_info),
	}

	err := ffi.CallFunction(cif_clang_TargetInfo_dispose, ptr_clang_TargetInfo_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TargetInfo_dispose", err))
	}
}

/*
Get the normalized target triple as a string.

Returns the empty string in case of any error.
*/
func TargetInfo_getTriple(info TargetInfo) String_ {
	c_info := info

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_info),
	}

	err := ffi.CallFunction(cif_clang_TargetInfo_getTriple, ptr_clang_TargetInfo_getTriple, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TargetInfo_getTriple", err))
	}

	ret := retC
	return ret
}

/*
Get the pointer width of the target in bits.

Returns -1 in case of error.
*/
func TargetInfo_getPointerWidth(info TargetInfo) int32 {
	c_info := info

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_info),
	}

	err := ffi.CallFunction(cif_clang_TargetInfo_getPointerWidth, ptr_clang_TargetInfo_getPointerWidth, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_TargetInfo_getPointerWidth", err))
	}

	ret := retC
	return ret
}

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

/*
Retrieve the cursor that represents the given translation unit.

The translation unit cursor can be used to start traversing the various declarations within the given translation unit.
*/
func GetTranslationUnitCursor(p0 TranslationUnit) Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getTranslationUnitCursor, ptr_clang_getTranslationUnitCursor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTranslationUnitCursor", err))
	}

	ret := retC
	return ret
}

// Determine whether two cursors are equivalent.
func EqualCursors(p0 Cursor, p1 Cursor) uint32 {
	c_p0 := p0
	c_p1 := p1

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(cif_clang_equalCursors, ptr_clang_equalCursors, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalCursors", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if cursor is null.
func Cursor_isNull(cursor Cursor) int32 {
	c_cursor := cursor

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isNull, ptr_clang_Cursor_isNull, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isNull", err))
	}

	ret := retC
	return ret
}

// Compute a hash value for the given cursor.
func HashCursor(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_hashCursor, ptr_clang_hashCursor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_hashCursor", err))
	}

	ret := retC
	return ret
}

// Retrieve the kind of the given cursor.
func GetCursorKind(p0 Cursor) CursorKind {
	c_p0 := p0

	var retC CursorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorKind, ptr_clang_getCursorKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorKind", err))
	}

	ret := retC
	return ret
}

// not supported : clang_isDeclaration : param p0 : enum CXCursorKind

/*
Determine whether the given declaration is invalid.

A declaration is invalid if it could not be parsed successfully.
*/
func IsInvalidDeclaration(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_isInvalidDeclaration, ptr_clang_isInvalidDeclaration, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isInvalidDeclaration", err))
	}

	ret := retC
	return ret
}

// not supported : clang_isReference : param p0 : enum CXCursorKind

// not supported : clang_isExpression : param p0 : enum CXCursorKind

// not supported : clang_isStatement : param p0 : enum CXCursorKind

// not supported : clang_isAttribute : param p0 : enum CXCursorKind

// Determine whether the given cursor has any attributes.
func Cursor_hasAttrs(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_hasAttrs, ptr_clang_Cursor_hasAttrs, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_hasAttrs", err))
	}

	ret := retC
	return ret
}

// not supported : clang_isInvalid : param p0 : enum CXCursorKind

// not supported : clang_isTranslationUnit : param p0 : enum CXCursorKind

// not supported : clang_isPreprocessing : param p0 : enum CXCursorKind

// not supported : clang_isUnexposed : param p0 : enum CXCursorKind

// Determine the linkage of the entity referred to by a given cursor.
func GetCursorLinkage(cursor Cursor) LinkageKind {
	c_cursor := cursor

	var retC LinkageKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorLinkage, ptr_clang_getCursorLinkage, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLinkage", err))
	}

	ret := retC
	return ret
}

/*
Describe the visibility of the entity referred to by a cursor.

This returns the default visibility if not explicitly specified by a visibility attribute. The default visibility may be changed by commandline arguments.
*/
func GetCursorVisibility(cursor Cursor) VisibilityKind {
	c_cursor := cursor

	var retC VisibilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorVisibility, ptr_clang_getCursorVisibility, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorVisibility", err))
	}

	ret := retC
	return ret
}

// Determine the availability of the entity that this cursor refers to, taking the current target platform into account.
func GetCursorAvailability(cursor Cursor) AvailabilityKind {
	c_cursor := cursor

	var retC AvailabilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorAvailability, ptr_clang_getCursorAvailability, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorAvailability", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getCursorPlatformAvailability : param always_deprecated : int *

// not supported : clang_disposeCXPlatformAvailability : param availability : CXPlatformAvailability *

// If cursor refers to a variable declaration and it has initializer returns cursor referring to the initializer otherwise return null cursor.
func Cursor_getVarDeclInitializer(cursor Cursor) Cursor {
	c_cursor := cursor

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getVarDeclInitializer, ptr_clang_Cursor_getVarDeclInitializer, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getVarDeclInitializer", err))
	}

	ret := retC
	return ret
}

// If cursor refers to a variable declaration that has global storage returns 1. If cursor refers to a variable declaration that doesn't have global storage returns 0. Otherwise returns -1.
func Cursor_hasVarDeclGlobalStorage(cursor Cursor) int32 {
	c_cursor := cursor

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_Cursor_hasVarDeclGlobalStorage, ptr_clang_Cursor_hasVarDeclGlobalStorage, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_hasVarDeclGlobalStorage", err))
	}

	ret := retC
	return ret
}

// If cursor refers to a variable declaration that has external storage returns 1. If cursor refers to a variable declaration that doesn't have external storage returns 0. Otherwise returns -1.
func Cursor_hasVarDeclExternalStorage(cursor Cursor) int32 {
	c_cursor := cursor

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_Cursor_hasVarDeclExternalStorage, ptr_clang_Cursor_hasVarDeclExternalStorage, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_hasVarDeclExternalStorage", err))
	}

	ret := retC
	return ret
}

// Determine the "language" of the entity referred to by a given cursor.
func GetCursorLanguage(cursor Cursor) LanguageKind {
	c_cursor := cursor

	var retC LanguageKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorLanguage, ptr_clang_getCursorLanguage, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLanguage", err))
	}

	ret := retC
	return ret
}

// Determine the "thread-local storage (TLS) kind" of the declaration referred to by a cursor.
func GetCursorTLSKind(cursor Cursor) TLSKind {
	c_cursor := cursor

	var retC TLSKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorTLSKind, ptr_clang_getCursorTLSKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorTLSKind", err))
	}

	ret := retC
	return ret
}

// Returns the translation unit that a cursor originated from.
func Cursor_getTranslationUnit(p0 Cursor) TranslationUnit {
	c_p0 := p0

	var retC TranslationUnit
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getTranslationUnit, ptr_clang_Cursor_getTranslationUnit, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTranslationUnit", err))
	}

	ret := retC
	return ret
}

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

// Disposes a CXCursorSet and releases its associated memory.
func DisposeCXCursorSet(cset CursorSet) {
	c_cset := cset

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cset),
	}

	err := ffi.CallFunction(cif_clang_disposeCXCursorSet, ptr_clang_disposeCXCursorSet, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_disposeCXCursorSet", err))
	}
}

// Queries a CXCursorSet to see if it contains a specific CXCursor.
func CursorSet_contains(cset CursorSet, cursor Cursor) uint32 {
	c_cset := cset
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cset),
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_CXCursorSet_contains, ptr_clang_CXCursorSet_contains, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXCursorSet_contains", err))
	}

	ret := retC
	return ret
}

// Inserts a CXCursor into a CXCursorSet.
func CursorSet_insert(cset CursorSet, cursor Cursor) uint32 {
	c_cset := cset
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cset),
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_CXCursorSet_insert, ptr_clang_CXCursorSet_insert, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXCursorSet_insert", err))
	}

	ret := retC
	return ret
}

/*
Determine the semantic parent of the given cursor.

The semantic parent of a cursor is the cursor that semantically contains the given cursor. For many declarations, the lexical and semantic parents are equivalent (the lexical parent is returned by clang_getCursorLexicalParent()). They diverge when declarations or definitions are provided out-of-line. For example:

In the out-of-line definition of C::f, the semantic parent is the class C, of which this function is a member. The lexical parent is the place where the declaration actually occurs in the source code; in this case, the definition occurs in the translation unit. In general, the lexical parent for a given entity can change without affecting the semantics of the program, and the lexical parent of different declarations of the same entity may be different. Changing the semantic parent of a declaration, on the other hand, can have a major impact on semantics, and redeclarations of a particular entity should all have the same semantic context.

In the example above, both declarations of C::f have C as their semantic context, while the lexical context of the first C::f is C and the lexical context of the second C::f is the translation unit.

For global declarations, the semantic parent is the translation unit.
*/
func GetCursorSemanticParent(cursor Cursor) Cursor {
	c_cursor := cursor

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorSemanticParent, ptr_clang_getCursorSemanticParent, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorSemanticParent", err))
	}

	ret := retC
	return ret
}

/*
Determine the lexical parent of the given cursor.

The lexical parent of a cursor is the cursor in which the given cursor was actually written. For many declarations, the lexical and semantic parents are equivalent (the semantic parent is returned by clang_getCursorSemanticParent()). They diverge when declarations or definitions are provided out-of-line. For example:

In the out-of-line definition of C::f, the semantic parent is the class C, of which this function is a member. The lexical parent is the place where the declaration actually occurs in the source code; in this case, the definition occurs in the translation unit. In general, the lexical parent for a given entity can change without affecting the semantics of the program, and the lexical parent of different declarations of the same entity may be different. Changing the semantic parent of a declaration, on the other hand, can have a major impact on semantics, and redeclarations of a particular entity should all have the same semantic context.

In the example above, both declarations of C::f have C as their semantic context, while the lexical context of the first C::f is C and the lexical context of the second C::f is the translation unit.

For declarations written in the global scope, the lexical parent is the translation unit.
*/
func GetCursorLexicalParent(cursor Cursor) Cursor {
	c_cursor := cursor

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorLexicalParent, ptr_clang_getCursorLexicalParent, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLexicalParent", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getOverriddenCursors : param overridden : CXCursor **

// not supported : clang_disposeOverriddenCursors : param overridden : CXCursor *

// Retrieve the file that is included by the given inclusion directive cursor.
func GetIncludedFile(cursor Cursor) File {
	c_cursor := cursor

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getIncludedFile, ptr_clang_getIncludedFile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getIncludedFile", err))
	}

	ret := retC
	return ret
}

/*
Map a source location to the cursor that describes the entity at that location in the source code.

clang_getCursor() maps an arbitrary source location within a translation unit down to the most specific cursor that describes the entity at that location. For example, given an expression x + y, invoking clang_getCursor() with a source location pointing to "x" will return the cursor for "x"; similarly for "y". If the cursor points anywhere between "x" or "y" (e.g., on the + or the whitespace around it), clang_getCursor() will return a cursor referring to the "+" expression.
*/
func GetCursor(p0 TranslationUnit, p1 SourceLocation) Cursor {
	c_p0 := p0
	c_p1 := p1

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(cif_clang_getCursor, ptr_clang_getCursor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursor", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the physical location of the source constructor referenced by the given cursor.

The location of a declaration is typically the location of the name of that declaration, where the name of that declaration would occur if it is unnamed, or some keyword that introduces that particular declaration. The location of a reference is where that reference occurs within the source code.
*/
func GetCursorLocation(p0 Cursor) SourceLocation {
	c_p0 := p0

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorLocation, ptr_clang_getCursorLocation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorLocation", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the physical extent of the source construct referenced by the given cursor.

The extent of a cursor starts with the file/line/column pointing at the first character within the source construct that the cursor refers to and ends with the last character within that source construct. For a declaration, the extent covers the declaration itself. For a reference, the extent covers the location of the reference (e.g., where the referenced entity was actually used).
*/
func GetCursorExtent(p0 Cursor) SourceRange {
	c_p0 := p0

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorExtent, ptr_clang_getCursorExtent, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorExtent", err))
	}

	ret := retC
	return ret
}

// Retrieve the type of a CXCursor (if any).
func GetCursorType(c Cursor) Type_ {
	c_c := c

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getCursorType, ptr_clang_getCursorType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorType", err))
	}

	ret := retC
	return ret
}

/*
Pretty-print the underlying type using the rules of the language of the translation unit from which it came.

If the type is invalid, an empty string is returned.
*/
func GetTypeSpelling(cT Type_) String_ {
	c_cT := cT

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(cif_clang_getTypeSpelling, ptr_clang_getTypeSpelling, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypeSpelling", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the underlying type of a typedef declaration.

If the cursor does not reference a typedef declaration, an invalid type is returned.
*/
func GetTypedefDeclUnderlyingType(c Cursor) Type_ {
	c_c := c

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getTypedefDeclUnderlyingType, ptr_clang_getTypedefDeclUnderlyingType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypedefDeclUnderlyingType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the integer type of an enum declaration.

If the cursor does not reference an enum declaration, an invalid type is returned.
*/
func GetEnumDeclIntegerType(c Cursor) Type_ {
	c_c := c

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getEnumDeclIntegerType, ptr_clang_getEnumDeclIntegerType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getEnumDeclIntegerType", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getEnumConstantDeclValue : return value : long long

// not supported : clang_getEnumConstantDeclUnsignedValue : return value : unsigned long long

// Returns non-zero if the cursor specifies a Record member that is a bit-field.
func Cursor_isBitField(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isBitField, ptr_clang_Cursor_isBitField, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isBitField", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the bit width of a bit-field declaration as an integer.

If the cursor does not reference a bit-field, or if the bit-field's width expression cannot be evaluated, -1 is returned.

For example:
*/
func GetFieldDeclBitWidth(c Cursor) int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getFieldDeclBitWidth, ptr_clang_getFieldDeclBitWidth, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFieldDeclBitWidth", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of non-variadic arguments associated with a given cursor.

The number of arguments can be determined for calls as well as for declarations of functions or methods. For other cursors -1 is returned.
*/
func Cursor_getNumArguments(c Cursor) int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getNumArguments, ptr_clang_Cursor_getNumArguments, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getNumArguments", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the argument cursor of a function or method.

The argument cursor can be determined for calls as well as for declarations of functions or methods. For other cursors and for invalid indices, an invalid cursor is returned.
*/
func Cursor_getArgument(c Cursor, i uint32) Cursor {
	c_c := c
	c_i := i

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getArgument, ptr_clang_Cursor_getArgument, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getArgument", err))
	}

	ret := retC
	return ret
}

/*
Returns the number of template args of a function, struct, or class decl representing a template specialization.

If the argument cursor cannot be converted into a template function declaration, -1 is returned.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

The value 3 would be returned from this call.
*/
func Cursor_getNumTemplateArguments(c Cursor) int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getNumTemplateArguments, ptr_clang_Cursor_getNumTemplateArguments, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getNumTemplateArguments", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the kind of the I'th template argument of the CXCursor C.

If the argument CXCursor does not represent a FunctionDecl, StructDecl, or ClassTemplatePartialSpecialization, an invalid template argument kind is returned.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

For I = 0, 1, and 2, Type, Integral, and Integral will be returned, respectively.
*/
func Cursor_getTemplateArgumentKind(c Cursor, i uint32) TemplateArgumentKind {
	c_c := c
	c_i := i

	var retC TemplateArgumentKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getTemplateArgumentKind, ptr_clang_Cursor_getTemplateArgumentKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTemplateArgumentKind", err))
	}

	ret := retC
	return ret
}

/*
Retrieve a CXType representing the type of a TemplateArgument of a  function decl representing a template specialization.

If the argument CXCursor does not represent a FunctionDecl, StructDecl, ClassDecl or ClassTemplatePartialSpecialization whose I'th template argument has a kind of CXTemplateArgKind_Integral, an invalid type is returned.

For example, for the following declaration and specialization:   template <typename T, int kInt, bool kBool>   void foo() { ... }

template <>   void foo<float, -7, true>();

If called with I = 0, "float", will be returned. Invalid types will be returned for I == 1 or 2.
*/
func Cursor_getTemplateArgumentType(c Cursor, i uint32) Type_ {
	c_c := c
	c_i := i

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getTemplateArgumentType, ptr_clang_Cursor_getTemplateArgumentType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getTemplateArgumentType", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Cursor_getTemplateArgumentValue : return value : long long

// not supported : clang_Cursor_getTemplateArgumentUnsignedValue : return value : unsigned long long

// Determine whether two CXTypes represent the same type.
func EqualTypes(a Type_, b Type_) uint32 {
	c_a := a
	c_b := b

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_a),
		unsafe.Pointer(&c_b),
	}

	err := ffi.CallFunction(cif_clang_equalTypes, ptr_clang_equalTypes, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_equalTypes", err))
	}

	ret := retC
	return ret
}

/*
Return the canonical type for a CXType.

Clang's type system explicitly models typedefs and all the ways a specific type can be represented.  The canonical type is the underlying type with all the "sugar" removed.  For example, if 'T' is a typedef for 'int', the canonical type for 'T' would be 'int'.
*/
func GetCanonicalType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getCanonicalType, ptr_clang_getCanonicalType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCanonicalType", err))
	}

	ret := retC
	return ret
}

// Determine whether a CXType has the "const" qualifier set, without looking through typedefs that may have added "const" at a different level.
func IsConstQualifiedType(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_isConstQualifiedType, ptr_clang_isConstQualifiedType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isConstQualifiedType", err))
	}

	ret := retC
	return ret
}

// Determine whether a  CXCursor that is a macro, is function like.
func Cursor_isMacroFunctionLike(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isMacroFunctionLike, ptr_clang_Cursor_isMacroFunctionLike, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isMacroFunctionLike", err))
	}

	ret := retC
	return ret
}

// Determine whether a  CXCursor that is a macro, is a builtin one.
func Cursor_isMacroBuiltin(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isMacroBuiltin, ptr_clang_Cursor_isMacroBuiltin, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isMacroBuiltin", err))
	}

	ret := retC
	return ret
}

// Determine whether a  CXCursor that is a function declaration, is an inline declaration.
func Cursor_isFunctionInlined(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isFunctionInlined, ptr_clang_Cursor_isFunctionInlined, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isFunctionInlined", err))
	}

	ret := retC
	return ret
}

// Determine whether a CXType has the "volatile" qualifier set, without looking through typedefs that may have added "volatile" at a different level.
func IsVolatileQualifiedType(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_isVolatileQualifiedType, ptr_clang_isVolatileQualifiedType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isVolatileQualifiedType", err))
	}

	ret := retC
	return ret
}

// Determine whether a CXType has the "restrict" qualifier set, without looking through typedefs that may have added "restrict" at a different level.
func IsRestrictQualifiedType(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_isRestrictQualifiedType, ptr_clang_isRestrictQualifiedType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isRestrictQualifiedType", err))
	}

	ret := retC
	return ret
}

// Returns the address space of the given type.
func GetAddressSpace(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getAddressSpace, ptr_clang_getAddressSpace, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getAddressSpace", err))
	}

	ret := retC
	return ret
}

// Returns the typedef name of the given type.
func GetTypedefName(cT Type_) String_ {
	c_cT := cT

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(cif_clang_getTypedefName, ptr_clang_getTypedefName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypedefName", err))
	}

	ret := retC
	return ret
}

// For pointer types, returns the type of the pointee.
func GetPointeeType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getPointeeType, ptr_clang_getPointeeType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getPointeeType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the unqualified variant of the given type, removing as little sugar as possible.

For example, given the following series of typedefs:

Executing clang_getUnqualifiedType() on a CXType that represents DifferenceType, will desugar to a type representing Integer, that has no qualifiers.

And, executing clang_getUnqualifiedType() on the type of the first argument of the following function declaration:

Will return a type representing int, removing the const qualifier.

Sugar over array types is not desugared.

A type can be checked for qualifiers with clang_isConstQualifiedType(), clang_isVolatileQualifiedType() and clang_isRestrictQualifiedType().

A type that resulted from a call to clang_getUnqualifiedType will return false for all of the above calls.
*/
func GetUnqualifiedType(cT Type_) Type_ {
	c_cT := cT

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(cif_clang_getUnqualifiedType, ptr_clang_getUnqualifiedType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getUnqualifiedType", err))
	}

	ret := retC
	return ret
}

/*
For reference types (e.g., "const int&"), returns the type that the reference refers to (e.g "const int").

Otherwise, returns the type itself.

A type that has kind CXType_LValueReference or CXType_RValueReference is a reference type.
*/
func GetNonReferenceType(cT Type_) Type_ {
	c_cT := cT

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(cif_clang_getNonReferenceType, ptr_clang_getNonReferenceType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNonReferenceType", err))
	}

	ret := retC
	return ret
}

// Return the cursor for the declaration of the given type.
func GetTypeDeclaration(t Type_) Cursor {
	c_t := t

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getTypeDeclaration, ptr_clang_getTypeDeclaration, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypeDeclaration", err))
	}

	ret := retC
	return ret
}

// Returns the Objective-C type encoding for the specified declaration.
func GetDeclObjCTypeEncoding(c Cursor) String_ {
	c_c := c

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getDeclObjCTypeEncoding, ptr_clang_getDeclObjCTypeEncoding, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getDeclObjCTypeEncoding", err))
	}

	ret := retC
	return ret
}

// Returns the Objective-C type encoding for the specified CXType.
func Type_getObjCEncoding(type_ Type_) String_ {
	c_type_ := type_

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_type_),
	}

	err := ffi.CallFunction(cif_clang_Type_getObjCEncoding, ptr_clang_Type_getObjCEncoding, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCEncoding", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getTypeKindSpelling : param K : enum CXTypeKind

/*
Retrieve the calling convention associated with a function type.

If a non-function type is passed in, CXCallingConv_Invalid is returned.
*/
func GetFunctionTypeCallingConv(t Type_) CallingConv {
	c_t := t

	var retC CallingConv
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getFunctionTypeCallingConv, ptr_clang_getFunctionTypeCallingConv, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFunctionTypeCallingConv", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the return type associated with a function type.

If a non-function type is passed in, an invalid type is returned.
*/
func GetResultType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getResultType, ptr_clang_getResultType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getResultType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the exception specification type associated with a function type. This is a value of type CXCursor_ExceptionSpecificationKind.

If a non-function type is passed in, an error code of -1 is returned.
*/
func GetExceptionSpecificationType(t Type_) int32 {
	c_t := t

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getExceptionSpecificationType, ptr_clang_getExceptionSpecificationType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getExceptionSpecificationType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of non-variadic parameters associated with a function type.

If a non-function type is passed in, -1 is returned.
*/
func GetNumArgTypes(t Type_) int32 {
	c_t := t

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getNumArgTypes, ptr_clang_getNumArgTypes, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumArgTypes", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the type of a parameter of a function type.

If a non-function type is passed in or the function does not have enough parameters, an invalid type is returned.
*/
func GetArgType(t Type_, i uint32) Type_ {
	c_t := t
	c_i := i

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_getArgType, ptr_clang_getArgType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getArgType", err))
	}

	ret := retC
	return ret
}

/*
Retrieves the base type of the ObjCObjectType.

If the type is not an ObjC object, an invalid type is returned.
*/
func Type_getObjCObjectBaseType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getObjCObjectBaseType, ptr_clang_Type_getObjCObjectBaseType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCObjectBaseType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of protocol references associated with an ObjC object/id.

If the type is not an ObjC object, 0 is returned.
*/
func Type_getNumObjCProtocolRefs(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getNumObjCProtocolRefs, ptr_clang_Type_getNumObjCProtocolRefs, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNumObjCProtocolRefs", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the decl for a protocol reference for an ObjC object/id.

If the type is not an ObjC object or there are not enough protocol references, an invalid cursor is returned.
*/
func Type_getObjCProtocolDecl(t Type_, i uint32) Cursor {
	c_t := t
	c_i := i

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_Type_getObjCProtocolDecl, ptr_clang_Type_getObjCProtocolDecl, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCProtocolDecl", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the number of type arguments associated with an ObjC object.

If the type is not an ObjC object, 0 is returned.
*/
func Type_getNumObjCTypeArgs(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getNumObjCTypeArgs, ptr_clang_Type_getNumObjCTypeArgs, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNumObjCTypeArgs", err))
	}

	ret := retC
	return ret
}

/*
Retrieve a type argument associated with an ObjC object.

If the type is not an ObjC or the index is not valid, an invalid type is returned.
*/
func Type_getObjCTypeArg(t Type_, i uint32) Type_ {
	c_t := t
	c_i := i

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_Type_getObjCTypeArg, ptr_clang_Type_getObjCTypeArg, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getObjCTypeArg", err))
	}

	ret := retC
	return ret
}

// Return 1 if the CXType is a variadic function type, and 0 otherwise.
func IsFunctionTypeVariadic(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_isFunctionTypeVariadic, ptr_clang_isFunctionTypeVariadic, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isFunctionTypeVariadic", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the return type associated with a given cursor.

This only returns a valid type if the cursor refers to a function or method.
*/
func GetCursorResultType(c Cursor) Type_ {
	c_c := c

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getCursorResultType, ptr_clang_getCursorResultType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorResultType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the exception specification type associated with a given cursor. This is a value of type CXCursor_ExceptionSpecificationKind.

This only returns a valid result if the cursor refers to a function or method.
*/
func GetCursorExceptionSpecificationType(c Cursor) int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getCursorExceptionSpecificationType, ptr_clang_getCursorExceptionSpecificationType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorExceptionSpecificationType", err))
	}

	ret := retC
	return ret
}

// Return 1 if the CXType is a POD (plain old data) type, and 0  otherwise.
func IsPODType(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_isPODType, ptr_clang_isPODType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isPODType", err))
	}

	ret := retC
	return ret
}

/*
Return the element type of an array, complex, or vector type.

If a type is passed in that is not an array, complex, or vector type, an invalid type is returned.
*/
func GetElementType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getElementType, ptr_clang_getElementType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getElementType", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getNumElements : return value : long long

/*
Return the element type of an array type.

If a non-array type is passed in, an invalid type is returned.
*/
func GetArrayElementType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_getArrayElementType, ptr_clang_getArrayElementType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getArrayElementType", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getArraySize : return value : long long

/*
Retrieve the type named by the qualified-id.

If a non-elaborated type is passed in, an invalid type is returned.
*/
func Type_getNamedType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getNamedType, ptr_clang_Type_getNamedType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNamedType", err))
	}

	ret := retC
	return ret
}

/*
Determine if a typedef is 'transparent' tag.

A typedef is considered 'transparent' if it shares a name and spelling location with its underlying tag type, as is the case with the NS_ENUM macro.
*/
func Type_isTransparentTagTypedef(t Type_) uint32 {
	c_t := t

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_isTransparentTagTypedef, ptr_clang_Type_isTransparentTagTypedef, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_isTransparentTagTypedef", err))
	}

	ret := retC
	return ret
}

// Retrieve the nullability kind of a pointer type.
func Type_getNullability(t Type_) TypeNullabilityKind {
	c_t := t

	var retC TypeNullabilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getNullability, ptr_clang_Type_getNullability, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNullability", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Type_getAlignOf : return value : long long

/*
Return the class type of an member pointer type.

If a non-member-pointer type is passed in, an invalid type is returned.
*/
func Type_getClassType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getClassType, ptr_clang_Type_getClassType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getClassType", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Type_getSizeOf : return value : long long

// not supported : clang_Type_getOffsetOf : return value : long long

/*
Return the type that was modified by this attributed type.

If the type is not an attributed type, an invalid type is returned.
*/
func Type_getModifiedType(t Type_) Type_ {
	c_t := t

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getModifiedType, ptr_clang_Type_getModifiedType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getModifiedType", err))
	}

	ret := retC
	return ret
}

/*
Gets the type contained by this atomic type.

If a non-atomic type is passed in, an invalid type is returned.
*/
func Type_getValueType(cT Type_) Type_ {
	c_cT := cT

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
	}

	err := ffi.CallFunction(cif_clang_Type_getValueType, ptr_clang_Type_getValueType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getValueType", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Cursor_getOffsetOfField : return value : long long

// Determine whether the given cursor represents an anonymous tag or namespace
func Cursor_isAnonymous(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isAnonymous, ptr_clang_Cursor_isAnonymous, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isAnonymous", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor represents an anonymous record declaration.
func Cursor_isAnonymousRecordDecl(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isAnonymousRecordDecl, ptr_clang_Cursor_isAnonymousRecordDecl, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isAnonymousRecordDecl", err))
	}

	ret := retC
	return ret
}

// Determine whether the given cursor represents an inline namespace declaration.
func Cursor_isInlineNamespace(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isInlineNamespace, ptr_clang_Cursor_isInlineNamespace, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isInlineNamespace", err))
	}

	ret := retC
	return ret
}

// Returns the number of template arguments for given template specialization, or -1 if type T is not a template specialization.
func Type_getNumTemplateArguments(t Type_) int32 {
	c_t := t

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getNumTemplateArguments, ptr_clang_Type_getNumTemplateArguments, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getNumTemplateArguments", err))
	}

	ret := retC
	return ret
}

/*
Returns the type template argument of a template class specialization at given index.

This function only returns template type arguments and does not handle template template arguments or variadic packs.
*/
func Type_getTemplateArgumentAsType(t Type_, i uint32) Type_ {
	c_t := t
	c_i := i

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_i),
	}

	err := ffi.CallFunction(cif_clang_Type_getTemplateArgumentAsType, ptr_clang_Type_getTemplateArgumentAsType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getTemplateArgumentAsType", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the ref-qualifier kind of a function or method.

The ref-qualifier is returned for C++ functions or methods. For other types or non-C++ declarations, CXRefQualifier_None is returned.
*/
func Type_getCXXRefQualifier(t Type_) RefQualifierKind {
	c_t := t

	var retC RefQualifierKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
	}

	err := ffi.CallFunction(cif_clang_Type_getCXXRefQualifier, ptr_clang_Type_getCXXRefQualifier, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_getCXXRefQualifier", err))
	}

	ret := retC
	return ret
}

// Returns 1 if the base class specified by the cursor with kind   CX_CXXBaseSpecifier is virtual.
func IsVirtualBase(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_isVirtualBase, ptr_clang_isVirtualBase, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isVirtualBase", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getOffsetOfBase : return value : long long

/*
Returns the access control level for the referenced object.

If the cursor refers to a C++ declaration, its access control level within its parent scope is returned. Otherwise, if the cursor refers to a base specifier or access specifier, the specifier itself is returned.
*/
func GetCXXAccessSpecifier(p0 Cursor) CXXAccessSpecifier {
	c_p0 := p0

	var retC CXXAccessSpecifier
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCXXAccessSpecifier, ptr_clang_getCXXAccessSpecifier, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCXXAccessSpecifier", err))
	}

	ret := retC
	return ret
}

func Cursor_getBinaryOpcode(c Cursor) BinaryOperatorKind_ {
	c_c := c

	var retC BinaryOperatorKind_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getBinaryOpcode, ptr_clang_Cursor_getBinaryOpcode, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getBinaryOpcode", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Cursor_getBinaryOpcodeStr : param Op : enum CX_BinaryOperatorKind

/*
Returns the storage class for a function or variable declaration.

If the passed in Cursor is not a function or variable declaration, CX_SC_Invalid is returned else the storage class.
*/
func Cursor_getStorageClass(p0 Cursor) StorageClass {
	c_p0 := p0

	var retC StorageClass
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getStorageClass, ptr_clang_Cursor_getStorageClass, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getStorageClass", err))
	}

	ret := retC
	return ret
}

// Determine the number of overloaded declarations referenced by a CXCursor_OverloadedDeclRef cursor.
func GetNumOverloadedDecls(cursor Cursor) uint32 {
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getNumOverloadedDecls, ptr_clang_getNumOverloadedDecls, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumOverloadedDecls", err))
	}

	ret := retC
	return ret
}

// Retrieve a cursor for one of the overloaded declarations referenced by a CXCursor_OverloadedDeclRef cursor.
func GetOverloadedDecl(cursor Cursor, index uint32) Cursor {
	c_cursor := cursor
	c_index := index

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(cif_clang_getOverloadedDecl, ptr_clang_getOverloadedDecl, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getOverloadedDecl", err))
	}

	ret := retC
	return ret
}

// For cursors representing an iboutletcollection attribute,  this function returns the collection element type.
func GetIBOutletCollectionType(p0 Cursor) Type_ {
	c_p0 := p0

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getIBOutletCollectionType, ptr_clang_getIBOutletCollectionType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getIBOutletCollectionType", err))
	}

	ret := retC
	return ret
}

/*
Visit the children of a particular cursor.

This function visits all the direct children of the given cursor, invoking the given visitor function with the cursors of each visited child. The traversal may be recursive, if the visitor returns CXChildVisit_Recurse. The traversal may also be ended prematurely, if the visitor returns CXChildVisit_Break.
*/
func VisitChildren(parent Cursor, visitor CursorVisitor, client_data ClientData) uint32 {
	c_parent := parent
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_parent),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(cif_clang_visitChildren, ptr_clang_visitChildren, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitChildren", err))
	}

	ret := retC
	return ret
}

// Visits the children of a cursor using the specified block.  Behaves identically to clang_visitChildren() in all other respects.
func VisitChildrenWithBlock(parent Cursor, block CursorVisitorBlock) uint32 {
	c_parent := parent
	c_block := block

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_parent),
		unsafe.Pointer(&c_block),
	}

	err := ffi.CallFunction(cif_clang_visitChildrenWithBlock, ptr_clang_visitChildrenWithBlock, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitChildrenWithBlock", err))
	}

	ret := retC
	return ret
}

/*
Retrieve a Unified Symbol Resolution (USR) for the entity referenced by the given cursor.

A Unified Symbol Resolution (USR) is a string that identifies a particular entity (function, class, variable, etc.) within a program. USRs can be compared across translation units to determine, e.g., when references in one translation refer to an entity defined in another translation unit.
*/
func GetCursorUSR(p0 Cursor) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorUSR, ptr_clang_getCursorUSR, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorUSR", err))
	}

	ret := retC
	return ret
}

// not supported : clang_constructUSR_ObjCClass : param class_name : const char *

// not supported : clang_constructUSR_ObjCCategory : param class_name : const char *

// not supported : clang_constructUSR_ObjCProtocol : param protocol_name : const char *

// not supported : clang_constructUSR_ObjCIvar : param name : const char *

// not supported : clang_constructUSR_ObjCMethod : param name : const char *

// not supported : clang_constructUSR_ObjCProperty : param property : const char *

// Retrieve a name for the entity referenced by this cursor.
func GetCursorSpelling(p0 Cursor) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorSpelling, ptr_clang_getCursorSpelling, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorSpelling", err))
	}

	ret := retC
	return ret
}

// Retrieve a range for a piece that forms the cursors spelling name. Most of the times there is only one range for the complete spelling but for Objective-C methods and Objective-C message expressions, there are multiple pieces for each selector identifier.
func Cursor_getSpellingNameRange(p0 Cursor, pieceIndex uint32, options uint32) SourceRange {
	c_p0 := p0
	c_pieceIndex := pieceIndex
	c_options := options

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_pieceIndex),
		unsafe.Pointer(&c_options),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getSpellingNameRange, ptr_clang_Cursor_getSpellingNameRange, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getSpellingNameRange", err))
	}

	ret := retC
	return ret
}

// not supported : clang_PrintingPolicy_getProperty : param Property : enum CXPrintingPolicyProperty

// not supported : clang_PrintingPolicy_setProperty : param Property : enum CXPrintingPolicyProperty

/*
Retrieve the default policy for the cursor.

The policy should be released after use with clang_PrintingPolicy_dispose.
*/
func GetCursorPrintingPolicy(p0 Cursor) PrintingPolicy {
	c_p0 := p0

	var retC PrintingPolicy
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorPrintingPolicy, ptr_clang_getCursorPrintingPolicy, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorPrintingPolicy", err))
	}

	ret := retC
	return ret
}

// Release a printing policy.
func PrintingPolicy_dispose(policy PrintingPolicy) {
	c_policy := policy

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_policy),
	}

	err := ffi.CallFunction(cif_clang_PrintingPolicy_dispose, ptr_clang_PrintingPolicy_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_PrintingPolicy_dispose", err))
	}
}

// Pretty print declarations.
func GetCursorPrettyPrinted(cursor Cursor, policy PrintingPolicy) String_ {
	c_cursor := cursor
	c_policy := policy

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_policy),
	}

	err := ffi.CallFunction(cif_clang_getCursorPrettyPrinted, ptr_clang_getCursorPrettyPrinted, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorPrettyPrinted", err))
	}

	ret := retC
	return ret
}

/*
Pretty-print the underlying type using a custom printing policy.

If the type is invalid, an empty string is returned.
*/
func GetTypePrettyPrinted(cT Type_, cxPolicy PrintingPolicy) String_ {
	c_cT := cT
	c_cxPolicy := cxPolicy

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
		unsafe.Pointer(&c_cxPolicy),
	}

	err := ffi.CallFunction(cif_clang_getTypePrettyPrinted, ptr_clang_getTypePrettyPrinted, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTypePrettyPrinted", err))
	}

	ret := retC
	return ret
}

/*
Get the fully qualified name for a type.

This includes full qualification of all template parameters.

Policy - Further refine the type formatting WithGlobalNsPrefix - If non-zero, function will prepend a '::' to qualified names
*/
func GetFullyQualifiedName(cT Type_, policy PrintingPolicy, withGlobalNsPrefix uint32) String_ {
	c_cT := cT
	c_policy := policy
	c_withGlobalNsPrefix := withGlobalNsPrefix

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cT),
		unsafe.Pointer(&c_policy),
		unsafe.Pointer(&c_withGlobalNsPrefix),
	}

	err := ffi.CallFunction(cif_clang_getFullyQualifiedName, ptr_clang_getFullyQualifiedName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getFullyQualifiedName", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the display name for the entity referenced by this cursor.

The display name contains extra information that helps identify the cursor, such as the parameters of a function or template or the arguments of a class template specialization.
*/
func GetCursorDisplayName(p0 Cursor) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorDisplayName, ptr_clang_getCursorDisplayName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorDisplayName", err))
	}

	ret := retC
	return ret
}

/*
For a cursor that is a reference, retrieve a cursor representing the entity that it references.

Reference cursors refer to other entities in the AST. For example, an Objective-C superclass reference cursor refers to an Objective-C class. This function produces the cursor for the Objective-C class from the cursor for the superclass reference. If the input cursor is a declaration or definition, it returns that declaration or definition unchanged. Otherwise, returns the NULL cursor.
*/
func GetCursorReferenced(p0 Cursor) Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorReferenced, ptr_clang_getCursorReferenced, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorReferenced", err))
	}

	ret := retC
	return ret
}

/*
For a cursor that is either a reference to or a declaration  of some entity, retrieve a cursor that describes the definition of  that entity.

Some entities can be declared multiple times within a translation  unit, but only one of those declarations can also be a  definition. For example, given:

there are three declarations of the function "f", but only the  second one is a definition. The clang_getCursorDefinition()  function will take any cursor pointing to a declaration of "f"  (the first or fourth lines of the example) or a cursor referenced  that uses "f" (the call to "f' inside "g") and will return a  declaration cursor pointing to the definition (the second "f"  declaration).

If given a cursor for which there is no corresponding definition,  e.g., because there is no definition of that entity within this  translation unit, returns a NULL cursor.
*/
func GetCursorDefinition(p0 Cursor) Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCursorDefinition, ptr_clang_getCursorDefinition, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorDefinition", err))
	}

	ret := retC
	return ret
}

// Determine whether the declaration pointed to by this cursor is also a definition of that entity.
func IsCursorDefinition(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_isCursorDefinition, ptr_clang_isCursorDefinition, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_isCursorDefinition", err))
	}

	ret := retC
	return ret
}

/*
Retrieve the canonical cursor corresponding to the given cursor.

In the C family of languages, many kinds of entities can be declared several times within a single translation unit. For example, a structure type can be forward-declared (possibly multiple times) and later defined:

The declarations and the definition of X are represented by three different cursors, all of which are declarations of the same underlying entity. One of these cursor is considered the "canonical" cursor, which is effectively the representative for the underlying entity. One can determine if two cursors are declarations of the same underlying entity by comparing their canonical cursors.
*/
func GetCanonicalCursor(p0 Cursor) Cursor {
	c_p0 := p0

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getCanonicalCursor, ptr_clang_getCanonicalCursor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCanonicalCursor", err))
	}

	ret := retC
	return ret
}

/*
If the cursor points to a selector identifier in an Objective-C method or message expression, this returns the selector index.

After getting a cursor with #clang_getCursor, this can be called to determine if the location points to a selector identifier.
*/
func Cursor_getObjCSelectorIndex(p0 Cursor) int32 {
	c_p0 := p0

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getObjCSelectorIndex, ptr_clang_Cursor_getObjCSelectorIndex, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCSelectorIndex", err))
	}

	ret := retC
	return ret
}

/*
Given a cursor pointing to a C++ method call or an Objective-C message, returns non-zero if the method/message is "dynamic", meaning:

For a C++ method: the call is virtual. For an Objective-C message: the receiver is an object instance, not 'super' or a specific class.

If the method/message is "static" or the cursor does not point to a method/message, it will return zero.
*/
func Cursor_isDynamicCall(c Cursor) int32 {
	c_c := c

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isDynamicCall, ptr_clang_Cursor_isDynamicCall, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isDynamicCall", err))
	}

	ret := retC
	return ret
}

// Given a cursor pointing to an Objective-C message or property reference, or C++ method call, returns the CXType of the receiver.
func Cursor_getReceiverType(c Cursor) Type_ {
	c_c := c

	var retC Type_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getReceiverType, ptr_clang_Cursor_getReceiverType, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getReceiverType", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a property declaration, return the associated property attributes. The bits are formed from CXObjCPropertyAttrKind.
func Cursor_getObjCPropertyAttributes(c Cursor, reserved uint32) uint32 {
	c_c := c
	c_reserved := reserved

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_reserved),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getObjCPropertyAttributes, ptr_clang_Cursor_getObjCPropertyAttributes, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCPropertyAttributes", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a property declaration, return the name of the method that implements the getter.
func Cursor_getObjCPropertyGetterName(c Cursor) String_ {
	c_c := c

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getObjCPropertyGetterName, ptr_clang_Cursor_getObjCPropertyGetterName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCPropertyGetterName", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a property declaration, return the name of the method that implements the setter, if any.
func Cursor_getObjCPropertySetterName(c Cursor) String_ {
	c_c := c

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getObjCPropertySetterName, ptr_clang_Cursor_getObjCPropertySetterName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCPropertySetterName", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents an Objective-C method or parameter declaration, return the associated Objective-C qualifiers for the return type or the parameter respectively. The bits are formed from CXObjCDeclQualifierKind.
func Cursor_getObjCDeclQualifiers(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getObjCDeclQualifiers, ptr_clang_Cursor_getObjCDeclQualifiers, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getObjCDeclQualifiers", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents an Objective-C method or property declaration, return non-zero if the declaration was affected by "\@optional". Returns zero if the cursor is not such a declaration or it is "\@required".
func Cursor_isObjCOptional(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isObjCOptional, ptr_clang_Cursor_isObjCOptional, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isObjCOptional", err))
	}

	ret := retC
	return ret
}

// Returns non-zero if the given cursor is a variadic function or method.
func Cursor_isVariadic(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isVariadic, ptr_clang_Cursor_isVariadic, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isVariadic", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Cursor_isExternalSymbol : param language : CXString *

// Given a cursor that represents a declaration, return the associated comment's source range.  The range may include multiple consecutive comments with whitespace in between.
func Cursor_getCommentRange(c Cursor) SourceRange {
	c_c := c

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getCommentRange, ptr_clang_Cursor_getCommentRange, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getCommentRange", err))
	}

	ret := retC
	return ret
}

// Given a cursor that represents a declaration, return the associated comment text, including comment markers.
func Cursor_getRawCommentText(c Cursor) String_ {
	c_c := c

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getRawCommentText, ptr_clang_Cursor_getRawCommentText, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getRawCommentText", err))
	}

	ret := retC
	return ret
}

/*
Given a cursor that represents a documentable entity (e.g., declaration), return the associated

first paragraph.
*/
func Cursor_getBriefCommentText(c Cursor) String_ {
	c_c := c

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getBriefCommentText, ptr_clang_Cursor_getBriefCommentText, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getBriefCommentText", err))
	}

	ret := retC
	return ret
}

// Retrieve the CXString representing the mangled name of the cursor.
func Cursor_getMangling(p0 Cursor) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getMangling, ptr_clang_Cursor_getMangling, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getMangling", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Cursor_getCXXManglings : return value : CXStringSet *

// not supported : clang_Cursor_getObjCManglings : return value : CXStringSet *

/*
Given a CXCursor_GCCAsmStmt cursor, return the assembly template string. As per LLVM IR Assembly Template language, template placeholders for inputs and outputs are either of the form $N where N is a decimal number as an index into the input-output specification, or ${N:M} where N is a decimal number also as an index into the input-output specification and M is the template argument modifier. The index N in both cases points into the the total inputs and outputs, or more specifically, into the list of outputs followed by the inputs, starting from index 0 as the first available template argument.

This function also returns a valid empty string if the cursor does not point at a GCC inline assembly block.

Users are responsible for releasing the allocation of returned string via clang_disposeString.
*/
func Cursor_getGCCAssemblyTemplate(p0 Cursor) String_ {
	c_p0 := p0

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getGCCAssemblyTemplate, ptr_clang_Cursor_getGCCAssemblyTemplate, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyTemplate", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, check if the assembly block has goto labels. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
func Cursor_isGCCAssemblyHasGoto(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isGCCAssemblyHasGoto, ptr_clang_Cursor_isGCCAssemblyHasGoto, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isGCCAssemblyHasGoto", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, count the number of outputs. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
func Cursor_getGCCAssemblyNumOutputs(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getGCCAssemblyNumOutputs, ptr_clang_Cursor_getGCCAssemblyNumOutputs, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyNumOutputs", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, count the number of inputs. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
func Cursor_getGCCAssemblyNumInputs(p0 Cursor) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getGCCAssemblyNumInputs, ptr_clang_Cursor_getGCCAssemblyNumInputs, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyNumInputs", err))
	}

	ret := retC
	return ret
}

// not supported : clang_Cursor_getGCCAssemblyInput : param Constraint : CXString *

// not supported : clang_Cursor_getGCCAssemblyOutput : param Constraint : CXString *

// Given a CXCursor_GCCAsmStmt cursor, count the clobbers in it. This function also returns 0 if the cursor does not point at a GCC inline assembly block.
func Cursor_getGCCAssemblyNumClobbers(cursor Cursor) uint32 {
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getGCCAssemblyNumClobbers, ptr_clang_Cursor_getGCCAssemblyNumClobbers, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyNumClobbers", err))
	}

	ret := retC
	return ret
}

/*
Given a CXCursor_GCCAsmStmt cursor, get the Index-th clobber of it. This function returns a valid empty string if the cursor does not point at a GCC inline assembly block or `Index` is out of bounds.

Users are responsible for releasing the allocation of returned string via clang_disposeString.
*/
func Cursor_getGCCAssemblyClobber(cursor Cursor, index uint32) String_ {
	c_cursor := cursor
	c_index := index

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getGCCAssemblyClobber, ptr_clang_Cursor_getGCCAssemblyClobber, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getGCCAssemblyClobber", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_GCCAsmStmt cursor, check if the inline assembly is `volatile`. This function returns 0 if the cursor does not point at a GCC inline assembly block.
func Cursor_isGCCAssemblyVolatile(cursor Cursor) uint32 {
	c_cursor := cursor

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_Cursor_isGCCAssemblyVolatile, ptr_clang_Cursor_isGCCAssemblyVolatile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_isGCCAssemblyVolatile", err))
	}

	ret := retC
	return ret
}

// Given a CXCursor_ModuleImportDecl cursor, return the associated module.
func Cursor_getModule(c Cursor) Module {
	c_c := c

	var retC Module
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_getModule, ptr_clang_Cursor_getModule, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_getModule", err))
	}

	ret := retC
	return ret
}

// Given a CXFile header file, return the module that contains it, if one exists.
func GetModuleForFile(p0 TranslationUnit, p1 File) Module {
	c_p0 := p0
	c_p1 := p1

	var retC Module
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(cif_clang_getModuleForFile, ptr_clang_getModuleForFile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getModuleForFile", err))
	}

	ret := retC
	return ret
}

func Module_getASTFile(module Module) File {
	c_module := module

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(cif_clang_Module_getASTFile, ptr_clang_Module_getASTFile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getASTFile", err))
	}

	ret := retC
	return ret
}

func Module_getParent(module Module) Module {
	c_module := module

	var retC Module
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(cif_clang_Module_getParent, ptr_clang_Module_getParent, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getParent", err))
	}

	ret := retC
	return ret
}

func Module_getName(module Module) String_ {
	c_module := module

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(cif_clang_Module_getName, ptr_clang_Module_getName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getName", err))
	}

	ret := retC
	return ret
}

func Module_getFullName(module Module) String_ {
	c_module := module

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(cif_clang_Module_getFullName, ptr_clang_Module_getFullName, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getFullName", err))
	}

	ret := retC
	return ret
}

func Module_isSystem(module Module) int32 {
	c_module := module

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(cif_clang_Module_isSystem, ptr_clang_Module_isSystem, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_isSystem", err))
	}

	ret := retC
	return ret
}

func Module_getNumTopLevelHeaders(p0 TranslationUnit, module Module) uint32 {
	c_p0 := p0
	c_module := module

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_module),
	}

	err := ffi.CallFunction(cif_clang_Module_getNumTopLevelHeaders, ptr_clang_Module_getNumTopLevelHeaders, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getNumTopLevelHeaders", err))
	}

	ret := retC
	return ret
}

func Module_getTopLevelHeader(p0 TranslationUnit, module Module, index uint32) File {
	c_p0 := p0
	c_module := module
	c_index := index

	var retC File
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_module),
		unsafe.Pointer(&c_index),
	}

	err := ffi.CallFunction(cif_clang_Module_getTopLevelHeader, ptr_clang_Module_getTopLevelHeader, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Module_getTopLevelHeader", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is a converting constructor.
func XConstructor_isConvertingConstructor(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXConstructor_isConvertingConstructor, ptr_clang_CXXConstructor_isConvertingConstructor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isConvertingConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is a copy constructor.
func XConstructor_isCopyConstructor(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXConstructor_isCopyConstructor, ptr_clang_CXXConstructor_isCopyConstructor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isCopyConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is the default constructor.
func XConstructor_isDefaultConstructor(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXConstructor_isDefaultConstructor, ptr_clang_CXXConstructor_isDefaultConstructor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isDefaultConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ constructor is a move constructor.
func XConstructor_isMoveConstructor(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXConstructor_isMoveConstructor, ptr_clang_CXXConstructor_isMoveConstructor, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXConstructor_isMoveConstructor", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ field is declared 'mutable'.
func XField_isMutable(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXField_isMutable, ptr_clang_CXXField_isMutable, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXField_isMutable", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ method is declared '= default'.
func XMethod_isDefaulted(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isDefaulted, ptr_clang_CXXMethod_isDefaulted, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isDefaulted", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ method is declared '= delete'.
func XMethod_isDeleted(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isDeleted, ptr_clang_CXXMethod_isDeleted, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isDeleted", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is pure virtual.
func XMethod_isPureVirtual(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isPureVirtual, ptr_clang_CXXMethod_isPureVirtual, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isPureVirtual", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is declared 'static'.
func XMethod_isStatic(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isStatic, ptr_clang_CXXMethod_isStatic, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isStatic", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is explicitly declared 'virtual' or if it overrides a virtual method from one of the base classes.
func XMethod_isVirtual(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isVirtual, ptr_clang_CXXMethod_isVirtual, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isVirtual", err))
	}

	ret := retC
	return ret
}

/*
Determine if a C++ member function is a copy-assignment operator, returning 1 if such is the case and 0 otherwise.

> A copy-assignment operator `X::operator=` is a non-static, > non-template member function of _class_ `X` with exactly one > parameter of type `X`, `X&`, `const X&`, `volatile X&` or `const > volatile X&`.

That is, for example, the `operator=` in:

class Foo {        bool operator=(const volatile Foo&);    };

Is a copy-assignment operator, while the `operator=` in:

class Bar {        bool operator=(const int&);    };

Is not.
*/
func XMethod_isCopyAssignmentOperator(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isCopyAssignmentOperator, ptr_clang_CXXMethod_isCopyAssignmentOperator, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isCopyAssignmentOperator", err))
	}

	ret := retC
	return ret
}

/*
Determine if a C++ member function is a move-assignment operator, returning 1 if such is the case and 0 otherwise.

> A move-assignment operator `X::operator=` is a non-static, > non-template member function of _class_ `X` with exactly one > parameter of type `X&&`, `const X&&`, `volatile X&&` or `const > volatile X&&`.

That is, for example, the `operator=` in:

class Foo {        bool operator=(const volatile Foo&&);    };

Is a move-assignment operator, while the `operator=` in:

class Bar {        bool operator=(const int&&);    };

Is not.
*/
func XMethod_isMoveAssignmentOperator(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isMoveAssignmentOperator, ptr_clang_CXXMethod_isMoveAssignmentOperator, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isMoveAssignmentOperator", err))
	}

	ret := retC
	return ret
}

/*
Determines if a C++ constructor or conversion function was declared explicit, returning 1 if such is the case and 0 otherwise.

Constructors or conversion functions are declared explicit through the use of the explicit specifier.

For example, the following constructor and conversion function are not explicit as they lack the explicit specifier:

class Foo {         Foo();         operator int();     };

While the following constructor and conversion function are explicit as they are declared with the explicit specifier.

class Foo {         explicit Foo();         explicit operator int();     };

This function will return 0 when given a cursor pointing to one of the former declarations and it will return 1 for a cursor pointing to the latter declarations.

The explicit specifier allows the user to specify a conditional compile-time expression whose value decides whether the marked element is explicit or not.

For example:

constexpr bool foo(int i) { return i % 2 == 0; }

class Foo {          explicit(foo(1)) Foo();          explicit(foo(2)) operator int();     }

This function will return 0 for the constructor and 1 for the conversion function.
*/
func XMethod_isExplicit(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isExplicit, ptr_clang_CXXMethod_isExplicit, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isExplicit", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ record is abstract, i.e. whether a class or struct has a pure virtual member function.
func XRecord_isAbstract(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXRecord_isAbstract, ptr_clang_CXXRecord_isAbstract, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXRecord_isAbstract", err))
	}

	ret := retC
	return ret
}

// Determine if an enum declaration refers to a scoped enum.
func EnumDecl_isScoped(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_EnumDecl_isScoped, ptr_clang_EnumDecl_isScoped, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EnumDecl_isScoped", err))
	}

	ret := retC
	return ret
}

// Determine if a C++ member function or member function template is declared 'const'.
func XMethod_isConst(c Cursor) uint32 {
	c_c := c

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_CXXMethod_isConst, ptr_clang_CXXMethod_isConst, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_CXXMethod_isConst", err))
	}

	ret := retC
	return ret
}

/*
Given a cursor that represents a template, determine the cursor kind of the specializations would be generated by instantiating the template.

This routine can be used to determine what flavor of function template, class template, or class template partial specialization is stored in the cursor. For example, it can describe whether a class template cursor is declared with "struct", "class" or "union".
*/
func GetTemplateCursorKind(c Cursor) CursorKind {
	c_c := c

	var retC CursorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getTemplateCursorKind, ptr_clang_getTemplateCursorKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTemplateCursorKind", err))
	}

	ret := retC
	return ret
}

/*
Given a cursor that may represent a specialization or instantiation of a template, retrieve the cursor that represents the template that it specializes or from which it was instantiated.

This routine determines the template involved both for explicit specializations of templates and for implicit instantiations of the template, both of which are referred to as "specializations". For a class template specialization (e.g., std::vector<bool>), this routine will return either the primary template (std::vector) or, if the specialization was instantiated from a class template partial specialization, the class template partial specialization. For a class template partial specialization and a function template specialization (including instantiations), this this routine will return the specialized template.

For members of a class template (e.g., member functions, member classes, or static data members), returns the specialized or instantiated member. Although not strictly "templates" in the C++ language, members of class templates have the same notions of specializations and instantiations that templates do, so this routine treats them similarly.
*/
func GetSpecializedCursorTemplate(c Cursor) Cursor {
	c_c := c

	var retC Cursor
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_getSpecializedCursorTemplate, ptr_clang_getSpecializedCursorTemplate, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getSpecializedCursorTemplate", err))
	}

	ret := retC
	return ret
}

// Given a cursor that references something else, return the source range covering that reference.
func GetCursorReferenceNameRange(c Cursor, nameFlags uint32, pieceIndex uint32) SourceRange {
	c_c := c
	c_nameFlags := nameFlags
	c_pieceIndex := pieceIndex

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
		unsafe.Pointer(&c_nameFlags),
		unsafe.Pointer(&c_pieceIndex),
	}

	err := ffi.CallFunction(cif_clang_getCursorReferenceNameRange, ptr_clang_getCursorReferenceNameRange, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorReferenceNameRange", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getToken : return value : CXToken *

// Determine the kind of the given token.
func GetTokenKind(p0 Token) TokenKind {
	c_p0 := p0

	var retC TokenKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_getTokenKind, ptr_clang_getTokenKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenKind", err))
	}

	ret := retC
	return ret
}

/*
Determine the spelling of the given token.

The spelling of a token is the textual representation of that token, e.g., the text of an identifier or keyword.
*/
func GetTokenSpelling(p0 TranslationUnit, p1 Token) String_ {
	c_p0 := p0
	c_p1 := p1

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(cif_clang_getTokenSpelling, ptr_clang_getTokenSpelling, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenSpelling", err))
	}

	ret := retC
	return ret
}

// Retrieve the source location of the given token.
func GetTokenLocation(p0 TranslationUnit, p1 Token) SourceLocation {
	c_p0 := p0
	c_p1 := p1

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(cif_clang_getTokenLocation, ptr_clang_getTokenLocation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenLocation", err))
	}

	ret := retC
	return ret
}

// Retrieve a source range that covers the given token.
func GetTokenExtent(p0 TranslationUnit, p1 Token) SourceRange {
	c_p0 := p0
	c_p1 := p1

	var retC SourceRange
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
	}

	err := ffi.CallFunction(cif_clang_getTokenExtent, ptr_clang_getTokenExtent, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getTokenExtent", err))
	}

	ret := retC
	return ret
}

// not supported : clang_tokenize : param Tokens : CXToken **

// not supported : clang_annotateTokens : param Tokens : CXToken *

// not supported : clang_disposeTokens : param Tokens : CXToken *

// not supported : clang_getCursorKindSpelling : param Kind : enum CXCursorKind

// not supported : clang_getDefinitionSpellingAndExtent : param startBuf : const char **

func EnableStackTraces() {
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_enableStackTraces, ptr_clang_enableStackTraces, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_enableStackTraces", err))
	}
}

// not supported : clang_executeOnThread : param fn : void (*)(void *)

// Determine the kind of a particular chunk within a completion string.
func GetCompletionChunkKind(completion_string CompletionString, chunk_number uint32) CompletionChunkKind {
	c_completion_string := completion_string
	c_chunk_number := chunk_number

	var retC CompletionChunkKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
		unsafe.Pointer(&c_chunk_number),
	}

	err := ffi.CallFunction(cif_clang_getCompletionChunkKind, ptr_clang_getCompletionChunkKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionChunkKind", err))
	}

	ret := retC
	return ret
}

// Retrieve the text associated with a particular chunk within a completion string.
func GetCompletionChunkText(completion_string CompletionString, chunk_number uint32) String_ {
	c_completion_string := completion_string
	c_chunk_number := chunk_number

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
		unsafe.Pointer(&c_chunk_number),
	}

	err := ffi.CallFunction(cif_clang_getCompletionChunkText, ptr_clang_getCompletionChunkText, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionChunkText", err))
	}

	ret := retC
	return ret
}

// Retrieve the completion string associated with a particular chunk within a completion string.
func GetCompletionChunkCompletionString(completion_string CompletionString, chunk_number uint32) CompletionString {
	c_completion_string := completion_string
	c_chunk_number := chunk_number

	var retC CompletionString
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
		unsafe.Pointer(&c_chunk_number),
	}

	err := ffi.CallFunction(cif_clang_getCompletionChunkCompletionString, ptr_clang_getCompletionChunkCompletionString, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionChunkCompletionString", err))
	}

	ret := retC
	return ret
}

// Retrieve the number of chunks in the given code-completion string.
func GetNumCompletionChunks(completion_string CompletionString) uint32 {
	c_completion_string := completion_string

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
	}

	err := ffi.CallFunction(cif_clang_getNumCompletionChunks, ptr_clang_getNumCompletionChunks, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getNumCompletionChunks", err))
	}

	ret := retC
	return ret
}

/*
Determine the priority of this code completion.

The priority of a code completion indicates how likely it is that this particular completion is the completion that the user will select. The priority is selected by various internal heuristics.
*/
func GetCompletionPriority(completion_string CompletionString) uint32 {
	c_completion_string := completion_string

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
	}

	err := ffi.CallFunction(cif_clang_getCompletionPriority, ptr_clang_getCompletionPriority, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionPriority", err))
	}

	ret := retC
	return ret
}

// Determine the availability of the entity that this code-completion string refers to.
func GetCompletionAvailability(completion_string CompletionString) AvailabilityKind {
	c_completion_string := completion_string

	var retC AvailabilityKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
	}

	err := ffi.CallFunction(cif_clang_getCompletionAvailability, ptr_clang_getCompletionAvailability, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionAvailability", err))
	}

	ret := retC
	return ret
}

// Retrieve the number of annotations associated with the given completion string.
func GetCompletionNumAnnotations(completion_string CompletionString) uint32 {
	c_completion_string := completion_string

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
	}

	err := ffi.CallFunction(cif_clang_getCompletionNumAnnotations, ptr_clang_getCompletionNumAnnotations, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionNumAnnotations", err))
	}

	ret := retC
	return ret
}

// Retrieve the annotation associated with the given completion string.
func GetCompletionAnnotation(completion_string CompletionString, annotation_number uint32) String_ {
	c_completion_string := completion_string
	c_annotation_number := annotation_number

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
		unsafe.Pointer(&c_annotation_number),
	}

	err := ffi.CallFunction(cif_clang_getCompletionAnnotation, ptr_clang_getCompletionAnnotation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionAnnotation", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getCompletionParent : param kind : enum CXCursorKind *

// Retrieve the brief documentation comment attached to the declaration that corresponds to the given completion string.
func GetCompletionBriefComment(completion_string CompletionString) String_ {
	c_completion_string := completion_string

	var retC String_
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_completion_string),
	}

	err := ffi.CallFunction(cif_clang_getCompletionBriefComment, ptr_clang_getCompletionBriefComment, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCompletionBriefComment", err))
	}

	ret := retC
	return ret
}

// Retrieve a completion string for an arbitrary declaration or macro definition cursor.
func GetCursorCompletionString(cursor Cursor) CompletionString {
	c_cursor := cursor

	var retC CompletionString
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorCompletionString, ptr_clang_getCursorCompletionString, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorCompletionString", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getCompletionNumFixIts : param results : CXCodeCompleteResults *

// not supported : clang_getCompletionFixIt : param results : CXCodeCompleteResults *

// Returns a default set of code-completion options that can be passed toclang_codeCompleteAt().
func DefaultCodeCompleteOptions() uint32 {
	var retC uint32
	args := []unsafe.Pointer{}

	err := ffi.CallFunction(cif_clang_defaultCodeCompleteOptions, ptr_clang_defaultCodeCompleteOptions, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_defaultCodeCompleteOptions", err))
	}

	ret := retC
	return ret
}

// not supported : clang_codeCompleteAt : return value : CXCodeCompleteResults *

// not supported : clang_sortCodeCompletionResults : param Results : CXCompletionResult *

// not supported : clang_disposeCodeCompleteResults : param Results : CXCodeCompleteResults *

// not supported : clang_codeCompleteGetNumDiagnostics : param Results : CXCodeCompleteResults *

// not supported : clang_codeCompleteGetDiagnostic : param Results : CXCodeCompleteResults *

// not supported : clang_codeCompleteGetContexts : return value : unsigned long long

// not supported : clang_codeCompleteGetContainerKind : param Results : CXCodeCompleteResults *

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
func ToggleCrashRecovery(isEnabled uint32) {
	c_isEnabled := isEnabled

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_isEnabled),
	}

	err := ffi.CallFunction(cif_clang_toggleCrashRecovery, ptr_clang_toggleCrashRecovery, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_toggleCrashRecovery", err))
	}
}

// Visit the set of preprocessor inclusions in a translation unit.   The visitor function is called with the provided data for every included   file.  This does not include headers included by the PCH file (unless one   is inspecting the inclusions in the PCH file itself).
func GetInclusions(tu TranslationUnit, visitor InclusionVisitor, client_data ClientData) {
	c_tu := tu
	c_visitor := visitor
	c_client_data := client_data

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tu),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(cif_clang_getInclusions, ptr_clang_getInclusions, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getInclusions", err))
	}
}

// If cursor is a statement declaration tries to evaluate the statement and if its variable, tries to evaluate its initializer, into its corresponding type. If it's an expression, tries to evaluate the expression.
func Cursor_Evaluate(c Cursor) EvalResult {
	c_c := c

	var retC EvalResult
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_c),
	}

	err := ffi.CallFunction(cif_clang_Cursor_Evaluate, ptr_clang_Cursor_Evaluate, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Cursor_Evaluate", err))
	}

	ret := retC
	return ret
}

// Returns the kind of the evaluated result.
func EvalResult_getKind(e EvalResult) EvalResultKind {
	c_e := e

	var retC EvalResultKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e),
	}

	err := ffi.CallFunction(cif_clang_EvalResult_getKind, ptr_clang_EvalResult_getKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getKind", err))
	}

	ret := retC
	return ret
}

// Returns the evaluation result as integer if the kind is Int.
func EvalResult_getAsInt(e EvalResult) int32 {
	c_e := e

	var retC int32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e),
	}

	err := ffi.CallFunction(cif_clang_EvalResult_getAsInt, ptr_clang_EvalResult_getAsInt, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_getAsInt", err))
	}

	ret := retC
	return ret
}

// not supported : clang_EvalResult_getAsLongLong : return value : long long

// Returns a non-zero value if the kind is Int and the evaluation result resulted in an unsigned integer.
func EvalResult_isUnsignedInt(e EvalResult) uint32 {
	c_e := e

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e),
	}

	err := ffi.CallFunction(cif_clang_EvalResult_isUnsignedInt, ptr_clang_EvalResult_isUnsignedInt, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_isUnsignedInt", err))
	}

	ret := retC
	return ret
}

// not supported : clang_EvalResult_getAsUnsigned : return value : unsigned long long

// not supported : clang_EvalResult_getAsDouble : return value : double

// not supported : clang_EvalResult_getAsStr : return value : const char *

// Disposes the created Eval memory.
func EvalResult_dispose(e EvalResult) {
	c_e := e

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_e),
	}

	err := ffi.CallFunction(cif_clang_EvalResult_dispose, ptr_clang_EvalResult_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_EvalResult_dispose", err))
	}
}

// Find references of a declaration in a specific file.
func FindReferencesInFile(cursor Cursor, file File, visitor CursorAndRangeVisitor) Result {
	c_cursor := cursor
	c_file := file
	c_visitor := visitor

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_visitor),
	}

	err := ffi.CallFunction(cif_clang_findReferencesInFile, ptr_clang_findReferencesInFile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findReferencesInFile", err))
	}

	ret := retC
	return ret
}

// Find #import/#include directives in a specific file.
func FindIncludesInFile(tU TranslationUnit, file File, visitor CursorAndRangeVisitor) Result {
	c_tU := tU
	c_file := file
	c_visitor := visitor

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_tU),
		unsafe.Pointer(&c_file),
		unsafe.Pointer(&c_visitor),
	}

	err := ffi.CallFunction(cif_clang_findIncludesInFile, ptr_clang_findIncludesInFile, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findIncludesInFile", err))
	}

	ret := retC
	return ret
}

func FindReferencesInFileWithBlock(p0 Cursor, p1 File, p2 CursorAndRangeVisitorBlock) Result {
	c_p0 := p0
	c_p1 := p1
	c_p2 := p2

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
		unsafe.Pointer(&c_p2),
	}

	err := ffi.CallFunction(cif_clang_findReferencesInFileWithBlock, ptr_clang_findReferencesInFileWithBlock, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findReferencesInFileWithBlock", err))
	}

	ret := retC
	return ret
}

func FindIncludesInFileWithBlock(p0 TranslationUnit, p1 File, p2 CursorAndRangeVisitorBlock) Result {
	c_p0 := p0
	c_p1 := p1
	c_p2 := p2

	var retC Result
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
		unsafe.Pointer(&c_p1),
		unsafe.Pointer(&c_p2),
	}

	err := ffi.CallFunction(cif_clang_findIncludesInFileWithBlock, ptr_clang_findIncludesInFileWithBlock, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_findIncludesInFileWithBlock", err))
	}

	ret := retC
	return ret
}

// not supported : clang_index_isEntityObjCContainerKind : param p0 : CXIdxEntityKind

// not supported : clang_index_getObjCContainerDeclInfo : return value : const CXIdxObjCContainerDeclInfo *

// not supported : clang_index_getObjCInterfaceDeclInfo : return value : const CXIdxObjCInterfaceDeclInfo *

// not supported : clang_index_getObjCCategoryDeclInfo : return value : const CXIdxObjCCategoryDeclInfo *

// not supported : clang_index_getObjCProtocolRefListInfo : return value : const CXIdxObjCProtocolRefListInfo *

// not supported : clang_index_getObjCPropertyDeclInfo : return value : const CXIdxObjCPropertyDeclInfo *

// not supported : clang_index_getIBOutletCollectionAttrInfo : return value : const CXIdxIBOutletCollectionAttrInfo *

// not supported : clang_index_getCXXClassDeclInfo : return value : const CXIdxCXXClassDeclInfo *

// not supported : clang_index_getClientContainer : param p0 : const CXIdxContainerInfo *

// not supported : clang_index_setClientContainer : param p0 : const CXIdxContainerInfo *

// not supported : clang_index_getClientEntity : param p0 : const CXIdxEntityInfo *

// not supported : clang_index_setClientEntity : param p0 : const CXIdxEntityInfo *

// An indexing action/session, to be applied to one or multiple translation units.
func IndexAction_create(cIdx Index) IndexAction {
	c_cIdx := cIdx

	var retC IndexAction
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cIdx),
	}

	err := ffi.CallFunction(cif_clang_IndexAction_create, ptr_clang_IndexAction_create, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_IndexAction_create", err))
	}

	ret := retC
	return ret
}

/*
Destroy the given index action.

The index action must not be destroyed until all of the translation units created within that index action have been destroyed.
*/
func IndexAction_dispose(p0 IndexAction) {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_IndexAction_dispose, ptr_clang_IndexAction_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_IndexAction_dispose", err))
	}
}

// not supported : clang_indexSourceFile : param index_callbacks : IndexerCallbacks *

// not supported : clang_indexSourceFileFullArgv : param index_callbacks : IndexerCallbacks *

// not supported : clang_indexTranslationUnit : param index_callbacks : IndexerCallbacks *

// not supported : clang_indexLoc_getFileLocation : param indexFile : CXIdxClientFile *

// Retrieve the CXSourceLocation represented by the given CXIdxLoc.
func IndexLoc_getCXSourceLocation(loc IdxLoc) SourceLocation {
	c_loc := loc

	var retC SourceLocation
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_loc),
	}

	err := ffi.CallFunction(cif_clang_indexLoc_getCXSourceLocation, ptr_clang_indexLoc_getCXSourceLocation, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_indexLoc_getCXSourceLocation", err))
	}

	ret := retC
	return ret
}

/*
Visit the fields of a particular type.

This function visits all the direct fields of the given cursor, invoking the given visitor function with the cursors of each visited field. The traversal may be ended prematurely, if the visitor returns CXFieldVisit_Break.
*/
func Type_visitFields(t Type_, visitor FieldVisitor, client_data ClientData) uint32 {
	c_t := t
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(cif_clang_Type_visitFields, ptr_clang_Type_visitFields, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_Type_visitFields", err))
	}

	ret := retC
	return ret
}

/*
Visit the base classes of a type.

This function visits all the direct base classes of a the given cursor, invoking the given visitor function with the cursors of each visited base. The traversal may be ended prematurely, if the visitor returns CXFieldVisit_Break.
*/
func VisitCXXBaseClasses(t Type_, visitor FieldVisitor, client_data ClientData) uint32 {
	c_t := t
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(cif_clang_visitCXXBaseClasses, ptr_clang_visitCXXBaseClasses, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitCXXBaseClasses", err))
	}

	ret := retC
	return ret
}

/*
Visit the class methods of a type.

This function visits all the methods of the given cursor, invoking the given visitor function with the cursors of each visited method. The traversal may be ended prematurely, if the visitor returns CXFieldVisit_Break.
*/
func VisitCXXMethods(t Type_, visitor FieldVisitor, client_data ClientData) uint32 {
	c_t := t
	c_visitor := visitor
	c_client_data := client_data

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_t),
		unsafe.Pointer(&c_visitor),
		unsafe.Pointer(&c_client_data),
	}

	err := ffi.CallFunction(cif_clang_visitCXXMethods, ptr_clang_visitCXXMethods, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_visitCXXMethods", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getBinaryOperatorKindSpelling : param kind : enum CXBinaryOperatorKind

/*
Retrieve the binary operator kind of this cursor.

If this cursor is not a binary operator then returns Invalid.
*/
func GetCursorBinaryOperatorKind(cursor Cursor) BinaryOperatorKind {
	c_cursor := cursor

	var retC BinaryOperatorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorBinaryOperatorKind, ptr_clang_getCursorBinaryOperatorKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorBinaryOperatorKind", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getUnaryOperatorKindSpelling : param kind : enum CXUnaryOperatorKind

/*
Retrieve the unary operator kind of this cursor.

If this cursor is not a unary operator then returns Invalid.
*/
func GetCursorUnaryOperatorKind(cursor Cursor) UnaryOperatorKind {
	c_cursor := cursor

	var retC UnaryOperatorKind
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_cursor),
	}

	err := ffi.CallFunction(cif_clang_getCursorUnaryOperatorKind, ptr_clang_getCursorUnaryOperatorKind, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_getCursorUnaryOperatorKind", err))
	}

	ret := retC
	return ret
}

// not supported : clang_getRemappings : param p0 : const char *

// not supported : clang_getRemappingsFromFileList : param p0 : const char **

func Remap_getNumFiles(p0 Remapping) uint32 {
	c_p0 := p0

	var retC uint32
	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_remap_getNumFiles, ptr_clang_remap_getNumFiles, unsafe.Pointer(&retC), args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_remap_getNumFiles", err))
	}

	ret := retC
	return ret
}

// not supported : clang_remap_getFilenames : param p2 : CXString *

func Remap_dispose(p0 Remapping) {
	c_p0 := p0

	args := []unsafe.Pointer{
		unsafe.Pointer(&c_p0),
	}

	err := ffi.CallFunction(cif_clang_remap_dispose, ptr_clang_remap_dispose, nil, args)
	if err != nil {
		panic(fmt.Sprintf("failed to call %s : %s", "clang_remap_dispose", err))
	}
}
