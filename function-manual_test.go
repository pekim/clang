package clang

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var clangResourceDir = sync.OnceValue(func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		panic(err)
	}

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		panic("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		panic(fmt.Sprintf("expected clang resource dir to start with '/', but it %s", resDir))
	}

	return resDir
})

func TestVisitChildren(t *testing.T) {
	assert.NoError(t, Init())

	index := CreateIndex(0, 1)
	assert.NotNil(t, index)

	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	var tu TranslationUnit
	errorCode := ParseTranslationUnit2(index, "test-data/test.h", parseArgs, nil,
		// errorCode := ParseTranslationUnit2(index, "../clang-go-generate/clang-c/Index.h", parseArgs, nil,
		uint32(TranslationUnit_SkipFunctionBodies|TranslationUnit_DetailedPreprocessingRecord),
		// 	0,
		&tu,
	)
	assert.Equal(t, Error_Success, errorCode)

	tuCursor := GetTranslationUnitCursor(tu)
	fmt.Println("cursor spelling", GetCString(GetCursorSpelling(tuCursor)))
	tuKind := GetCursorKind(tuCursor)
	fmt.Println(tuKind)
	// fmt.Println("cursor kind spelling", GetCString(GetCursorKindSpelling(tuKind)))
	ok := VisitChildren(tuCursor, func() ChildVisitResult {
		fmt.Println("!!")
		return ChildVisit_Continue
	}, nil)
	assert.Zero(t, ok)
}
