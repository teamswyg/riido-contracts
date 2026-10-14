package main

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

func skipSExprComment(source string, start int) int {
	for start < len(source) && source[start] != '\n' {
		start++
	}
	return start
}

func readAtomToken(source string, start int) (string, int) {
	i := start
	for i < len(source) {
		r, width := utf8.DecodeRuneInString(source[i:])
		if unicode.IsSpace(r) || r == '(' || r == ')' || r == ';' {
			break
		}
		i += width
	}
	return source[start:i], i
}

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
