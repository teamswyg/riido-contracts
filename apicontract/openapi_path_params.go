package apicontract

import (
	"regexp"
	"sort"
)

var pathParamPattern = regexp.MustCompile(`\{([^}/]+)\}`)

func pathParameters(path string) []OpenAPIParameter {
	matches := pathParamPattern.FindAllStringSubmatch(path, -1)
	params := make([]OpenAPIParameter, 0, len(matches))
	for _, match := range matches {
		params = append(params, OpenAPIParameter{
			Name:     match[1],
			In:       "path",
			Required: true,
			Schema:   map[string]any{"type": "string"},
		})
	}
	sort.Slice(params, func(i, j int) bool { return params[i].Name < params[j].Name })
	return params
}
