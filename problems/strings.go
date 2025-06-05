package problems

import (
	"reflect"
	"strconv"
)

func Strings_IsUnique(s string) bool {
	seen := make(map[rune]struct{})
	for _, c := range s {
		if _, ok := seen[c]; ok {
			return false
		}
		seen[c] = struct{}{}
	}
	return true
}

func Strings_CheckPermutation(a, b string) bool {
	aChars := make(map[rune]int)
	for _, c := range a {
		aChars[c] += 1
	}
	bChars := make(map[rune]int)
	for _, c := range b {
		bChars[c] += 1
	}
	return reflect.DeepEqual(aChars, bChars)
}

func Strings_URLify(s string) string {
	result := make([]rune, 0)
	for _, r := range s {
		if r == ' ' {
			result = append(result, '%')
			result = append(result, '2')
			result = append(result, '0')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// Checks if a is "one edit away" from b by counting differences.
func Strings_OneAway(a, b string) bool {
	lenA := len(a)
	lenB := len(b)
	if lenB-lenA > 1 || lenA-lenB > 1 {
		return false
	}
	if lenA == lenB {
		return oneAwayReplaced(a, b)
	}
	if lenA > lenB {
		return oneAwayInsertedOrDeleted(a, b)
	}
	return oneAwayInsertedOrDeleted(b, a)
}

func oneAwayReplaced(a, b string) bool {
	diffs := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs <= 1
}

func oneAwayInsertedOrDeleted(long, short string) bool {
	i := 0
	j := 0
	diffs := 0
	for i < len(long) && j < len(short) {
		if long[i] != short[i] {
			diffs++
			if diffs > 1 {
				return false
			}
			i++
		} else {
			i++
			j++
		}
	}
	return true
}

func Strings_StringCompression(s string) string {
	compressed := make([]rune, 0)
	last := '�'
	var count int
	for _, r := range s {
		if r == last {
			count++
		} else {
			if last != '�' {
				compressed = append(compressed, rune(strconv.Itoa(count)[0]))
			}
			count = 1
			compressed = append(compressed, r)
			last = r
		}
	}
	compressed = append(compressed, rune(strconv.Itoa(count)[0]))
	if len(compressed) >= len(s) {
		return s
	}
	return string(compressed)
}
