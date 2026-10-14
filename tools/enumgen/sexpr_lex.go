package main

import (
	"unicode"
	"unicode/utf8"
)

func lex(source string) ([]string, error) {
	out := []string{}
	for i := 0; i < len(source); {
		r, width := utf8.DecodeRuneInString(source[i:])
		switch {
		case unicode.IsSpace(r):
			i += width
		case r == ';':
			for i < len(source) && source[i] != '\n' {
				i++
			}
		case r == '(' || r == ')':
			out = append(out, string(r))
			i += width
		case r == '"':
			value, next, err := readStringToken(source, i)
			if err != nil {
				return nil, err
			}
			out = append(out, value)
			i = next
		default:
			token, next := readAtomToken(source, i)
			out = append(out, token)
			i = next
		}
	}
	return out, nil
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
