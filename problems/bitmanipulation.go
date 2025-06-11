package problems

import "fmt"

func binary_Add(a uint8, b uint8) uint8 {
	// E.g. 00001010 + 00000110 = 00010000.
	return a + b
}

func binary_AddSigned(a int8, b int8) int8 {
	// E.g. 10001010 + 00010000 = 00000110.
	return a + b
}

func binary_Subtract(a uint8, b uint8) uint8 {
	// E.g. 00010000 - 00001010 = 00000110.
	return a - b
}

func binary_SubtractSigned(a int8, b int8) int8 {
	// E.g. 00000110 - 00010000 = 10001010.
	return a - b
}

func binary_TwosComplement(i int8) string {
	// Negative integers are stored in Two's Complement form: convert to base-2, invert the bits, and add 1.
	// Convert int8 to uint8 to see Two's Complement representation: uint8(-5) gives 11111011.
	return fmt.Sprintf("%08b", uint8(i))
}

func binary_ShiftLeft(i uint8) uint8 {
	// Left shift (<<) doubles an unsigned value using a Logical Shift.
	// Logical shifts move all values, including the "signed" bit.
	return i << 1
}

func binary_ShiftRight(i uint8) uint8 {
	// Right shift (>>) halves an unsigned value using a Logical Shift, truncating the rightmost bit.
	// Logical shifts move all values, including the "signed" bit.
	return i >> 1
}

func binary_ShiftLeftSigned(i int8) int8 {
	// Left shift (<<) doubles a signed value using an Arithmetic Shift.
	// Arithmetic shifts preserve the "signed" bit.
	return i << 1
}

func binary_ShiftRightSigned(i int8) int8 {
	// Right shift (>>) halves an unsigned value using a Logical Shift, truncating the rightmost bit.
	// Arithmetic shifts preserve the "signed" bit.
	return i >> 1
}

func binary_And(a uint8, b uint8) uint8 {
	// AND: & can be used to Get a specific bit, or Clear a specific bit with inverted mask e.g. ^(1<<i).
	return a & b
}

func binary_Clear(a int8, b int8) int8 {
	// AND: & can be used to Get a specific bit, or Clear a specific bit with inverted mask e.g. ^(1<<i).
	// Operand must be signed int8 in order to allow inversions, e.g. ^(1<<2).
	return a & b
}

func binary_Or(a uint8, b uint8) uint8 {
	// OR: | can be used to Set a specific bit.
	return a | b
}

func binary_Xor(a uint8, b uint8) uint8 {
	// XOR: ^ can be used to detect differences.
	return a ^ b
}
