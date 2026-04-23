package lib

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/go-webgpu/goffi/ffi"
)

type LibraryPaths struct {
	Darwin string
	Linux  string
}

func LoadLibrary(paths LibraryPaths) unsafe.Pointer {
	var libPath string
	switch runtime.GOOS {
	case "darwin":
		libPath = paths.Darwin
	case "linux":
		libPath = paths.Linux
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}

	handle, err := ffi.LoadLibrary(libPath)
	if err != nil {
		panic(fmt.Errorf("failed to open library : %w", err))
	}
	return handle
}
