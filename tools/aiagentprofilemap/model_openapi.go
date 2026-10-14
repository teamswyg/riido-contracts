package main

import (
	"fmt"
	"strings"
)

func sameOpenAPIOperation(doc openAPIDoc, exp operationExpectation) bool {
	methods, ok := doc.Paths[exp.Path]
	if !ok {
		return false
	}
	op, ok := methods[strings.ToLower(exp.Method)]
	if !ok || op.OperationID != exp.OperationID {
		return false
	}
	if op.Client.GeneratedPath != exp.GeneratedPath {
		return false
	}
	return openAPIResponseRef(op, exp.ResponseStatus) == exp.ResponseRef
}

func openAPIResponseRef(op openAPIOperation, status int) string {
	resp, ok := op.Responses[fmt.Sprintf("%d", status)]
	if !ok {
		return ""
	}
	media, ok := resp.Content["application/json"]
	if !ok {
		return ""
	}
	return strings.TrimPrefix(media.Schema.Ref, "#/components/schemas/")
}
