package generate

import (
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/pekim/clang"
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

func (gen *gen) parseHeaderFile(headerFile string) {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}
	// var tu clang.TranslationUnit

	index := clang.CreateIndex(0, 1)
	tu, errCode := index.ParseTranslationUnit2(
		headerFile, parseArgs, nil,
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}

	gen.findEntities(tu, headerFile)
}
