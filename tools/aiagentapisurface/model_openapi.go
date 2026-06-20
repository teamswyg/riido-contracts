package main

import "strings"

func openAPITuples(doc openAPIDoc) []operationTuple {
	var out []operationTuple
	for path, methods := range doc.Paths {
		for method, op := range methods {
			if isHTTPMethod(method) {
				out = append(out, operationTuple{
					Method:      strings.ToUpper(method),
					Path:        path,
					OperationID: op.OperationID,
				})
			}
		}
	}
	sortTuples(out)
	return out
}

func isHTTPMethod(method string) bool {
	switch method {
	case "get", "post", "put", "patch", "delete":
		return true
	default:
		return false
	}
}
