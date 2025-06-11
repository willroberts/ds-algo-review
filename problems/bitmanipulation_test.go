package problems

import (
	"fmt"
	"testing"
)

func TestBitManipulation(t *testing.T) {
	var unsigned uint8 = 10             // Binary: 00001010.
	var addend uint8 = 6                // Binary: 00000110.
	sum := binary_Add(unsigned, addend) // Binary: 00010000.
	if fmt.Sprintf("%08b", sum) != "00010000" {
		t.Log("expected", "00010000", "got", fmt.Sprintf("%08b", sum))
		t.Fail()
	}

	var signed int8 = -10                              // Binary: 10001010.
	var subtrahend int8 = 16                           // Binary: 00010000.
	difference := binary_AddSigned(signed, subtrahend) // Binary: 00000110.
	if fmt.Sprintf("%08b", difference) != "00000110" {
		t.Log("expected", "00000110", "got", fmt.Sprintf("%08b", difference))
		t.Fail()
	}

	if binary_Subtract(8, 4) != 4 {
		// 00001000 - 00000100 = 00000100
		t.Log("expected 4, got", binary_Subtract(8, 4))
		t.Fail()
	}
	if binary_SubtractSigned(-8, 4) != -12 {
		// 10001000 - 00000100 = 10001100
		t.Log("expected -12, got", binary_SubtractSigned(-8, 4))
		t.Fail()
	}

	if binary_TwosComplement(-5) != "11111011" {
		t.Log("expected", "11111011", "got", binary_TwosComplement(-5))
		t.Fail()
	}

	if binary_ShiftLeft(8) != 16 {
		t.Log("expected", 16, "got", binary_ShiftLeft(8))
		t.Fail()
	}
	if binary_ShiftRight(8) != 4 {
		t.Log("expected", 4, "got", binary_ShiftRight(8))
		t.Fail()
	}
	if binary_ShiftLeftSigned(-8) != -16 {
		t.Log("expected", -16, "got", binary_ShiftLeftSigned(-8))
		t.Fail()
	}
	if binary_ShiftRightSigned(-8) != -4 {
		t.Log("expected", -4, "got", binary_ShiftRightSigned(-8))
		t.Fail()
	}

	var i uint8 = 85 // Binary: 01010101.
	// Get the 3rd bit (00000100).
	if binary_And(i, 1<<2) != 4 {
		t.Log("expected 4, got", binary_And(i, 1<<2))
		t.Fail()
	}
	// Set the 2nd bit to 1 (01010111).
	if binary_Or(i, 1<<1) != 87 {
		t.Log("expected 87, got", binary_Or(i, 1<<1))
		t.Fail()
	}
	// Clear the 3rd bit (01010001).
	if binary_Clear(int8(i), ^(1<<2)) != 81 {
		t.Log("expected 81, got", binary_Clear(int8(i), ^(1<<2)))
		t.Fail()
	}
	// Toggle the 4rd bit (01011101).
	if binary_Xor(i, 1<<3) != 93 {
		t.Log("expected 93, got", binary_Xor(i, 1<<3))
		t.Fail()
	}
}
