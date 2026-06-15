package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type node struct {
	atom string
	list []node
}

func (n node) isAtom() bool {
	return n.list == nil
}

func parseSExpr(source string) (node, error) {
	tokens, err := lex(source)
	if err != nil {
		return node{}, err
	}
	index := 0
	root, err := parseNode(tokens, &index)
	if err != nil {
		return node{}, err
	}
	if index != len(tokens) {
		return node{}, fmt.Errorf("unexpected trailing token %q", tokens[index])
	}
	return root, nil
}

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
			start := i
			for i < len(source) {
				r, width = utf8.DecodeRuneInString(source[i:])
				if unicode.IsSpace(r) || r == '(' || r == ')' || r == ';' {
					break
				}
				i += width
			}
			out = append(out, source[start:i])
		}
	}
	return out, nil
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

func parseNode(tokens []string, index *int) (node, error) {
	if *index >= len(tokens) {
		return node{}, errors.New("unexpected end of input")
	}
	token := tokens[*index]
	*index++
	switch token {
	case "(":
		var list []node
		for {
			if *index >= len(tokens) {
				return node{}, errors.New("unterminated list")
			}
			if tokens[*index] == ")" {
				*index++
				return node{list: list}, nil
			}
			child, err := parseNode(tokens, index)
			if err != nil {
				return node{}, err
			}
			list = append(list, child)
		}
	case ")":
		return node{}, errors.New("unexpected )")
	default:
		if strings.HasPrefix(token, "\"") {
			value, err := strconv.Unquote(token)
			if err != nil {
				return node{}, fmt.Errorf("decode string %s: %w", token, err)
			}
			return node{atom: value}, nil
		}
		return node{atom: token}, nil
	}
}

func atom(n node) string {
	if !n.isAtom() {
		return ""
	}
	return n.atom
}

func atomList(n node) []string {
	if n.isAtom() {
		value := atom(n)
		if value == "" {
			return nil
		}
		return []string{value}
	}
	out := make([]string, 0, len(n.list))
	for _, item := range n.list {
		value := atom(item)
		if value != "" {
			out = append(out, value)
		}
	}
	return out
}
