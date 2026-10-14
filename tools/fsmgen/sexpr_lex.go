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
			i = skipSExprComment(source, i)
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
			value, next := readAtomToken(source, i)
			out = append(out, value)
			i = next
		}
	}
	return out, nil
}
