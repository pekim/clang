package libc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCString(t *testing.T) {
	cString, free := CString("qwerty")
	assert.NotNil(t, cString)
	free()
}
