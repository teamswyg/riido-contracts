package main

import (
	"errors"
	"strings"
)

func readStringToken(source string, start int) (string, int, error) {
	var b strings.Builder
	b.WriteByte('"')
	escaped := false
	for i := start + 1; i < len(source); i++ {
		ch := source[i]
		b.WriteByte(ch)
		if escaped {
			escaped = false
			continue
		}
		if ch == '\\' {
			escaped = true
			continue
		}
		if ch == '"' {
			return b.String(), i + 1, nil
		}
	}
	return "", 0, errors.New("unterminated string literal")
}
