package stringbuilder

import "testing"

func TestStringBuilder(t *testing.T) {
	sb := NewStringBuilder()
	sb.Append([]string{"one", "two", "three"})
	sb.Append([]string{"four", "five", "six"})
	if sb.BuildString() != "one two three four five six" {
		t.Fail()
	}
}
