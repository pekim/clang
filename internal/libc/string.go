package libc

import (
	"math"
	"unsafe"
)

func CString(s string) (unsafe.Pointer, func()) {
	cStr := malloc(len(s) + 1)
	cStrSlice := unsafe.Slice(cStr, len(s)+1)
	copy(cStrSlice, []byte(s))
	cStrSlice[len(s)] = 0
	return unsafe.Pointer(cStr), func() { free(cStr) }
}

// GoString returns a Go string for a null terminated C string.
//
// The Go string is a copy of the C bytes.
// The C string memory is freed.
func GoStringFree(cStr unsafe.Pointer) string {
	defer CStringFree(cStr)
	return GoString(cStr)
}

func CStringFree(cStr unsafe.Pointer) {
	free((*byte)(cStr))
}

// goString returns a Go string for a null terminated C string.
//
// The Go string is a copy of the C bytes.
func GoString(cStr unsafe.Pointer) string {
	if cStr == nil {
		return ""
	}

	// Look for terminating 0, to find the length of the string.
	length := 0
	str := unsafe.Slice((*byte)(cStr), math.MaxInt)
	for i, b := range str {
		if b == 0 {
			length = i
			break
		}
	}

	goBytes := make([]byte, length)
	copy(goBytes, str)
	return string(goBytes)
}
