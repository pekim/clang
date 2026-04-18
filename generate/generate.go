package generate

import "fmt"

func Generate() {
	tu := newTranslationUnit("clang-c/Index.h")
	fmt.Println(tu)
}
