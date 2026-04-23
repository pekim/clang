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
