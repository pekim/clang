package generate

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
	packageDir string
	filename   string
}

func newFile(packageName string, packageDir string, filename string) file {
	f := file{
		File:       jen.NewFile(packageName),
		packageDir: packageDir,
		filename:   filename,
	}

	f.HeaderComment("This is a generated file. DO NOT EDIT.")
	f.Line()

	return f
}

func (f file) save() {
	filename := filepath.Join(f.packageDir, f.filename+".go")
	err := f.Save(filename)
	if err != nil {
		fatalOnError(err)
	}

	cmd := exec.Command("gofmt", "-w", filename)
	fmtOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to format %s\n", filename)
		fatalOnError(err)
	}
	if len(fmtOutput) > 0 {
		fmt.Println(string(fmtOutput))
	}
}

// func unsafePointer() *jen.Statement {
// 	return jen.Qual("unsafe", "Pointer")
// }
