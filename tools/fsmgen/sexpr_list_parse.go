package main

import "errors"

func parseListNode(tokens []string, index *int) (node, error) {
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
}
