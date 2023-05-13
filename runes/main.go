package runes

import "unicode"

// BEGIN (write your solution here)
func isASCII(s string) bool {
	for _, ch := range s {
		if ch > unicode.MaxASCII {
			return false
		}
	}
	return true
}