package problems

import (
	"fmt"
	"testing"
)

func TestBitManipulation(t *testing.T) {
	var unsigned uint8 = 10             // Binary: 00001010.
	var addend uint8 = 6                // Binary: 00000110.
	sum := Binary_Add(unsigned, addend) // Binary: 00010000.
	if fmt.Sprintf("%08b", sum) != "00010000" {
		t.Log("expected", "00010000", "got", fmt.Sprintf("%08b", sum))
		t.Fail()
	}

	var signed int8 = -10                              // Binary: 10001010.
	var subtrahend int8 = 16                           // Binary: 00010000.
	difference := Binary_AddSigned(signed, subtrahend) // Binary: 00000110.
	if fmt.Sprintf("%08b", difference) != "00000110" {
		t.Log("expected", "00000110", "got", fmt.Sprintf("%08b", difference))
		t.Fail()
	}

	if Binary_TwosComplement(-5) != "11111011" {
		t.Log("expected", "11111011", "got", Binary_TwosComplement(-5))
		t.Fail()
	}

	if Binary_ShiftLeft(8) != 16 {
		t.Log("expected", 16, "got", Binary_ShiftLeft(8))
		t.Fail()
	}
	if Binary_ShiftRightSigned(-8) != -4 {
		t.Log("expected", -4, "got", Binary_ShiftRightSigned(-8))
		t.Fail()
	}

	var i uint8 = 85 // Binary: 01010101.
	// Get the 3rd bit (00000100).
	if Binary_And(i, 1<<2) != 4 {
		t.Log("expected 4, got", Binary_And(i, 1<<2))
		t.Fail()
	}
	// Set the 2nd bit to 1 (01010111).
	if Binary_Or(i, 1<<1) != 87 {
		t.Log("expected 87, got", Binary_Or(i, 1<<1))
		t.Fail()
	}
	// Clear the 3rd bit (01010001).
	if Binary_Clear(int8(i), ^(1<<2)) != 81 {
		t.Log("expected 81, got", Binary_Clear(int8(i), ^(1<<2)))
		t.Fail()
	}
	// Toggle the 4rd bit (01011101).
	if Binary_Xor(i, 1<<3) != 93 {
		t.Log("expected 93, got", Binary_Xor(i, 1<<3))
		t.Fail()
	}
}
