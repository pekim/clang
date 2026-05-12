package clang

import (
	"unsafe"

	lib "github.com/pekim/clang/internal/lib"
)

type LibraryPaths struct {
	Darwin string
	Linux  string
}

func loadLibrary(userPaths *LibraryPaths) (unsafe.Pointer, error) {
	paths := lib.LibraryPaths{
		Darwin: "libclang.dylib",
		Linux:  "libclang.so",
	}
	if userPaths != nil {
		if userPaths.Darwin != "" {
			paths.Darwin = userPaths.Darwin
		}
		if userPaths.Linux != "" {
			paths.Linux = userPaths.Linux
		}
	}

	return lib.LoadLibrary(paths)
}
