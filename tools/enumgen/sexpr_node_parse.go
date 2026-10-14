package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseNode(tokens []string, index *int) (node, error) {
	if *index >= len(tokens) {
		return node{}, errors.New("unexpected end of input")
	}
	token := tokens[*index]
	*index++
	switch token {
	case "(":
		return parseList(tokens, index)
	case ")":
		return node{}, errors.New("unexpected )")
	default:
		return parseAtom(token)
	}
}

func parseAtom(token string) (node, error) {
	if !strings.HasPrefix(token, "\"") {
		return node{atom: token}, nil
	}
	value, err := strconv.Unquote(token)
	if err != nil {
		return node{}, fmt.Errorf("decode string %s: %w", token, err)
	}
	return node{atom: value}, nil
}
