package main

import "strings"

func normalizedSet(ops []operationTuple) map[string]bool {
	out := make(map[string]bool, len(ops))
	for _, op := range ops {
		out[normalizedKey(op)] = true
	}
	return out
}

func normalizedKey(op operationTuple) string {
	path := strings.TrimPrefix(op.Path, v1Prefix)
	path = strings.TrimPrefix(path, v2Prefix)
	return op.Method + " " + path
}
