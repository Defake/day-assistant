package strings

import (
	"strings"
)

func SplitLines(s string) []string {
	return strings.FieldsFunc(s, func(c rune) bool { 
		return c == '\n' || c == '\r' 
	})
}

func SplitBy(s string, splitBy rune) []string {
	return strings.FieldsFunc(s, func(c rune) bool { 
		return c == splitBy
	})
}
