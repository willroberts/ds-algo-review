package hashtable

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := NewHashTable(1024)
	ht.Insert("foo", "bar")
	ht.Insert("baz", "bat")

	val, err := ht.Get("foo")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if val != "bar" {
		t.Log("invalid value; expected 'bar'")
		t.Fail()
	}

	val, err = ht.Get("baz")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if val != "bat" {
		t.Log("invalid value; expected 'bat'")
		t.Fail()
	}
}
