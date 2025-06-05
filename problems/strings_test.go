package problems

import "testing"

func TestStrings_IsUnique(t *testing.T) {
	if !Strings_IsUnique("taco") {
		t.Fail()
	}
	if Strings_IsUnique("burrito") {
		t.Fail()
	}
}

func TestStrings_CheckPermutation(t *testing.T) {
	if !Strings_CheckPermutation("abcd", "dcba") {
		t.Fail()
	}
	if Strings_CheckPermutation("abcd", "efgh") {
		t.Fail()
	}
}

func TestStrings_URLify(t *testing.T) {
	if Strings_URLify("Hello World") != "Hello%20World" {
		t.Fail()
	}
}

func TestStrings_OneAway(t *testing.T) {
	if Strings_OneAway("pale", "bake") {
		t.Fail()
	}
	if !Strings_OneAway("pale", "bale") {
		t.Fail()
	}
}

func TestStrings_StringCompression(t *testing.T) {
	if Strings_StringCompression("aaabbbccc") != "a3b3c3" {
		t.Fail()
	}
	if Strings_StringCompression("aabbcc") != "aabbcc" {
		t.Fail()
	}
}
