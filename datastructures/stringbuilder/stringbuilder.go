package stringbuilder

import "strings"

// StringBuilder represents a string as an underlying slice of components.
// This reduces time complexity of string concatenation by avoiding copy operations.
type StringBuilder struct {
	parts []string
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{parts: make([]string, 0)}
}

func (sb *StringBuilder) Append(words []string) {
	sb.parts = append(sb.parts, words...)
}

func (sb *StringBuilder) BuildString() string {
	return strings.Join(sb.parts, " ")
}
