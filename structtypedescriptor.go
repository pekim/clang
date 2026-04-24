// This is a generated file. DO NOT EDIT.

package clang

import (
	"unsafe"

	types "github.com/go-webgpu/goffi/types"
)

var string_TypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(String_{}),
}

var stringSetTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(StringSet{}),
}

var virtualFileOverlayImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(VirtualFileOverlayImpl{}),
}

var moduleMapDescriptorImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(ModuleMapDescriptorImpl{}),
}

var fileUniqueIDTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      24,
		},
	},
	Size: unsafe.Sizeof(FileUniqueID{}),
}

var sourceLocationTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(SourceLocation{}),
}

var sourceRangeTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.UInt32TypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(SourceRange{}),
}

var sourceRangeListTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(SourceRangeList{}),
}

var targetInfoImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(TargetInfoImpl{}),
}

var translationUnitImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(TranslationUnitImpl{}),
}

var unsavedFileTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt64TypeDescriptor,
	},
	Size: unsafe.Sizeof(UnsavedFile{}),
}

var versionTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(Version{}),
}

var indexOptionsTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
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

var tUResourceUsageEntryTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		types.UInt64TypeDescriptor,
	},
	Size: unsafe.Sizeof(TUResourceUsageEntry{}),
}

var tUResourceUsageTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(TUResourceUsage{}),
}

var cursorTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		types.SInt32TypeDescriptor,
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      24,
		},
	},
	Size: unsafe.Sizeof(Cursor{}),
}

var platformAvailabilityTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
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

var cursorSetImplTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members:   []*types.TypeDescriptor{},
	Size:      unsafe.Sizeof(CursorSetImpl{}),
}

var type_TypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      16,
		},
	},
	Size: unsafe.Sizeof(Type_{}),
}

var tokenTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(Token{}),
}

var completionResultTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      8,
		},
	},
	Size: unsafe.Sizeof(CompletionResult{}),
}

var codeCompleteResultsTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(CodeCompleteResults{}),
}

var cursorAndRangeVisitorTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(CursorAndRangeVisitor{}),
}

var idxLocTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      16,
		},
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxLoc{}),
}

var idxIncludedFileInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		idxLocTypeDescriptor,
		types.PointerTypeDescriptor,
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      8,
		},
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
		types.SInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxIncludedFileInfo{}),
}

var idxImportedASTFileInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      8,
		},
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      8,
		},
		idxLocTypeDescriptor,
		types.SInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxImportedASTFileInfo{}),
}

var idxAttrInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxAttrInfo{}),
}

var idxEntityInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxEntityInfo{}),
}

var idxContainerInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		cursorTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxContainerInfo{}),
}

var idxIBOutletCollectionAttrInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxIBOutletCollectionAttrInfo{}),
}

var idxDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
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

var idxObjCContainerDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
	},
	Size: unsafe.Sizeof(IdxObjCContainerDeclInfo{}),
}

var idxBaseClassInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxBaseClassInfo{}),
}

var idxObjCProtocolRefInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCProtocolRefInfo{}),
}

var idxObjCProtocolRefListInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCProtocolRefListInfo{}),
}

var idxObjCInterfaceDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCInterfaceDeclInfo{}),
}

var idxObjCCategoryDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
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

var idxObjCPropertyDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxObjCPropertyDeclInfo{}),
}

var idxCXXClassDeclInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.UInt32TypeDescriptor,
	},
	Size: unsafe.Sizeof(IdxCXXClassDeclInfo{}),
}

var idxEntityRefInfoTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
	Kind:      types.StructType,
	Members: []*types.TypeDescriptor{
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
		cursorTypeDescriptor,
		idxLocTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		types.PointerTypeDescriptor,
		{
			Alignment: 0,
			Kind:      types.UInt8Type,
			Size:      4,
		},
	},
	Size: unsafe.Sizeof(IdxEntityRefInfo{}),
}

var indexerCallbacksTypeDescriptor = &types.TypeDescriptor{
	Alignment: 0,
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
