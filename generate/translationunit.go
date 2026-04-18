package generate

import (
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

type translationUnit struct {
	headerFile string
	clang.TranslationUnit
}

var clangResourceDir = sync.OnceValue(func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	fatalOnError(err)

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		fatal("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		fatalf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir
})

func newTranslationUnit(headerFile string) translationUnit {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	transUnit := translationUnit{
		headerFile: headerFile,
	}
	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(headerFile, parseArgs, nil,
		clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord,
		&transUnit.TranslationUnit)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}
	return transUnit
}
