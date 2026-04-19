package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
}

type enums struct {
	enums []enum
}

func (enums *enums) add(cursor clang.Cursor) {
	enums.enums = append(enums.enums, enum{})
}
