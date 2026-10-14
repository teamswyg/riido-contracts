package main

import (
	"fmt"
	"strings"
)

func verifyScope(model model) error {
	if !model.ForbiddenFieldsAbsent {
		return fmt.Errorf("assignment request schemas contain forbidden request fields")
	}
	if !model.NoDiffPathsAbsent {
		return fmt.Errorf("non-task path fragments reuse task-scoped AI Agent assignment operations")
	}
	for _, op := range model.Operations {
		if !operationHasTaskScope(op) {
			return fmt.Errorf("operation %s is not task scoped", op.OperationID)
		}
	}
	return nil
}

func operationHasTaskScope(op operation) bool {
	return strings.Contains(op.Path, "{task_id}")
}
