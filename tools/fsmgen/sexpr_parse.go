package main

import "fmt"

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
