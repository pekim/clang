package clang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitfield(t *testing.T) {
	var v uint16

	v = 0b0001_1000
	v = bitfieldSet(v, 0, 3, 0b0101)
	assert.Equal(t, uint16(0b0001_1101), v)
	assert.Equal(t, uint16(0b0101), bitfieldGet(v, 0, 3))

	v = 0b0001_1100
	v = bitfieldSet(v, 0, 3, 0b0101)
	assert.Equal(t, uint16(0b0001_1101), v)
	assert.Equal(t, uint16(0b0101), bitfieldGet(v, 0, 3))

	v = 0b0000_0010
	v = bitfieldSet(v, 4, 3, 0b0101)
	assert.Equal(t, uint16(0b0101_0010), v)
	assert.Equal(t, uint16(0b0101), bitfieldGet(v, 4, 3))

	v = 0b0101_0010
	v = bitfieldSet(v, 4, 3, 0b0011)
	assert.Equal(t, uint16(0b0011_0010), v)
	assert.Equal(t, uint16(0b0011), bitfieldGet(v, 4, 3))
}
