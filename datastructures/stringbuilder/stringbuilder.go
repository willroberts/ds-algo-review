package stringbuilder

import "strings"

// StringBuilder represents a string as an underlying slice of components.
// This reduces time complexity of string concatenation by avoiding copy operations.
type StringBuilder struct {
	parts []string
}

// NewStringBuilder instantiates a StringBuilder and preallocates its internal storage.
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{parts: make([]string, 0)}
}

// Append adds the given words to the StringBuilder.
func (sb *StringBuilder) Append(words []string) {
	sb.parts = append(sb.parts, words...)
}

// BuildString serializes the StringBuilder's internal storage.
func (sb *StringBuilder) BuildString() string {
	return strings.Join(sb.parts, " ")
}
