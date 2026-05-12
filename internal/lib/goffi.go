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

func LoadLibrary(paths LibraryPaths) (unsafe.Pointer, error) {
	var libPath string
	switch runtime.GOOS {
	case "darwin":
		libPath = paths.Darwin
	case "linux":
		libPath = paths.Linux
	default:
		return nil, fmt.Errorf("GOOS=%s is not supported", runtime.GOOS)
	}

	handle, err := ffi.LoadLibrary(libPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open library : %w", err)
	}
	return handle, nil
}
