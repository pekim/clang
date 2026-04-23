package clang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructsSizes(t *testing.T) {
	for _, struct_ := range structTestVars {
		t.Run(struct_.name, func(t *testing.T) {
			assert.Equal(t, struct_.cSize, struct_.goSize)
		})
	}
}

func TestStringStructField(t *testing.T) {
	var indexOptions IndexOptions
	assert.Empty(t, indexOptions.PreambleStoragePath())
	free := indexOptions.SetPreambleStoragePath("qwerty 42")
	assert.Equal(t, "qwerty 42", indexOptions.PreambleStoragePath())
	free()
}
