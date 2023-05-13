package strings

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here) (write your solution here)
func latinLetters(s string) string {
	res_str := &strings.Builder{}
	for _, ch := range s {
		if unicode.Is(unicode.Latin, ch) {
			res_str.WriteRune(ch)
		}
	}
	return res_str.String()
}