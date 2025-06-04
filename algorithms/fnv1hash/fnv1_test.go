package fnv1hash

import "testing"

func TestFNV1(t *testing.T) {
	input := "testing123"
	hash := FNV1([]byte(input))
	if hash != 8404323582830432641 {
		t.Fail()
	}
}

func TestFNV1a(t *testing.T) {
	input := "testing123"
	hash := FNV1a([]byte(input))
	if hash != 4517894324711253781 {
		t.Fail()
	}
}
