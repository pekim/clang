// This is a generated file. DO NOT EDIT.

package clang

import (
	"unsafe"

	types "github.com/go-webgpu/goffi/types"
)

var aPISetImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(APISetImpl{}),
}

var codeCompleteResultsTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(CodeCompleteResults{}),
}

var commentTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(Comment{}),
}

var completionResultTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(CompletionResult{}),
}

var cursorTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		{
			Alignment: 8,
			Kind:      types.UInt8Type,
			Size:      24,
		},
	},
	Size: unsafe.Sizeof(Cursor{}),
}

var cursorAndRangeVisitorTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(CursorAndRangeVisitor{}),
}

var cursorSetImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(CursorSetImpl{}),
}

var fileUniqueIDTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 8,
			Kind:      types.UInt8Type,
			Size:      24,
		},
	},
	Size: unsafe.Sizeof(FileUniqueID{}),
}

var idxAttrInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxAttrInfo{}),
}

var idxBaseClassInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxBaseClassInfo{}),
}

var idxCXXClassDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxCXXClassDeclInfo{}),
}

var idxContainerInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		cursorTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxContainerInfo{}),
}

var idxDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		types.PointerTypeDescriptor,
		types.SInt32TypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxDeclInfo{}),
}

var idxEntityInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxEntityInfo{}),
}

var idxEntityRefInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxEntityRefInfo{}),
}

var idxIBOutletCollectionAttrInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxIBOutletCollectionAttrInfo{}),
}

var idxImportedASTFileInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		idxLocTypeDescriptor,
		types.SInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxImportedASTFileInfo{}),
}

var idxIncludedFileInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		idxLocTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxIncludedFileInfo{}),
}

var idxLocTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 8,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxLoc{}),
}

var idxObjCCategoryDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCCategoryDeclInfo{}),
}

var idxObjCContainerDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCContainerDeclInfo{}),
}

var idxObjCInterfaceDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCInterfaceDeclInfo{}),
}

var idxObjCPropertyDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCPropertyDeclInfo{}),
}

var idxObjCProtocolRefInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCProtocolRefInfo{}),
}

var idxObjCProtocolRefListInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCProtocolRefListInfo{}),
}

var indexOptionsTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.UInt8TypeDescriptor,
		types.UInt8TypeDescriptor,
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IndexOptions{}),
}

var moduleMapDescriptorImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(ModuleMapDescriptorImpl{}),
}

var platformAvailabilityTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		string_TypeDescriptor,
		versionTypeDescriptor,
		versionTypeDescriptor,
		versionTypeDescriptor,
		types.SInt32TypeDescriptor,
		string_TypeDescriptor,
	},
	Size: unsafe.Sizeof(PlatformAvailability{}),
}

var sourceLocationTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 8,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(SourceLocation{}),
}

var sourceRangeTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 8,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(SourceRange{}),
}

var sourceRangeListTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(SourceRangeList{}),
}

var string_TypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(string_{}),
}

var stringSetTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(StringSet{}),
}

var tUResourceUsageTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(TUResourceUsage{}),
}

var tUResourceUsageEntryTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.UInt64TypeDescriptor,
	},
	Size: unsafe.Sizeof(TUResourceUsageEntry{}),
}

var targetInfoImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(TargetInfoImpl{}),
}

var tokenTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 4,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(Token{}),
}

var translationUnitImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(TranslationUnitImpl{}),
}

var type_TypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		{
			Alignment: 8,
			Kind:      types.UInt8Type,
			Size:      16,
		},
	},
	Size: unsafe.Sizeof(Type{}),
}

var unsavedFileTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt64TypeDescriptor,
	},
	Size: unsafe.Sizeof(UnsavedFile{}),
}

var versionTypeDescriptor = &types.TypeDescriptor{
	Alignment: 4,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(Version{}),
}

var virtualFileOverlayImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(VirtualFileOverlayImpl{}),
}

var indexerCallbacksTypeDescriptor = &types.TypeDescriptor{
	Alignment: 8,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IndexerCallbacks{}),
}
