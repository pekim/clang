package libc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMallocFree(t *testing.T) {
	ptr := malloc(4)
	assert.NotNil(t, ptr)
	free(ptr)
}
