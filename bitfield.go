package clang

import (
	"golang.org/x/exp/constraints"
)

func bitfieldGet[T constraints.Unsigned](field T, offset int, width int) uint {
	mask := bitfieldMask[T](offset, width)
	return uint((field & mask) >> T(offset))
}

func bitfieldSet[T constraints.Unsigned](field T, offset int, width int, value uint) T {
	// unset the bits first
	mask := ^bitfieldMask[T](offset, width)
	field = field & mask

	// set the necessary bits
	value <<= offset
	field |= T(value)

	return field
}

// bitfieldMask creates a mask of 1s for the bits at offset through
// to offset+width.
//
// For example if offset=3, width=2 and T=uint16, the mask produced would
// be 0b0000_0000_0001_1000
func bitfieldMask[T constraints.Unsigned](offset int, width int) T {
	return ^(^T(0) << uint(width)) << T(offset)
}
