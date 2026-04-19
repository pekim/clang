package generate

import (
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

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

func parseHeaderFile(headerFile string) clang.TranslationUnit {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-I", "/home/mike/projects/clang-go/clang-c",
		"-x", "c-header",
	}
	var tu clang.TranslationUnit

	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(
		headerFile, parseArgs, nil,
		clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord,
		&tu,
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}

	return tu
}
