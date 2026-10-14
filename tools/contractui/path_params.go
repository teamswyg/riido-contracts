package main

import "strings"

func pathParams(path string) []string {
	parts := strings.Split(path, "/")
	out := []string{}
	for _, part := range parts {
		if !strings.HasPrefix(part, "{") || !strings.HasSuffix(part, "}") {
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(part, "{"), "}")
		if name != "" {
			out = append(out, name)
		}
	}
	return out
}
