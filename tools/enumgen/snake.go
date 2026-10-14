package main

import (
	"strings"
	"unicode"
)

func snake(value string) string {
	var b strings.Builder
	var prevLower bool
	for index, r := range value {
		if index > 0 && unicode.IsUpper(r) && prevLower {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
		prevLower = unicode.IsLower(r) || unicode.IsDigit(r)
	}
	return b.String()
}
