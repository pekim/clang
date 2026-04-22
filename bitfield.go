package clang

import (
	"golang.org/x/exp/constraints"
)

func bitfieldGet[T constraints.Unsigned](field T, offset int, width int) T {
	mask := bitfieldMask[T](offset, width)
	return (field & mask) >> T(offset)
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

func bitfieldMask[T constraints.Unsigned](offset int, width int) T {
	return ^(^T(0) << uint(width)) << T(offset)
}
