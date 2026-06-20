package main

import "fmt"

func verifyOperation(model model) error {
	op := model.Operation
	exp := model.Manifest.RequiredOperation
	if op.OperationID != exp.OperationID {
		return fmt.Errorf("operation %s missing", exp.OperationID)
	}
	if op.Kind != exp.Kind || op.Method != exp.Method || op.Path != exp.Path {
		return fmt.Errorf("operation %s tuple mismatch", exp.OperationID)
	}
	if responseName(op) != exp.ResponseRef || responseStatus(op) != exp.ResponseStatus {
		return fmt.Errorf("operation %s response mismatch", exp.OperationID)
	}
	if op.RBACPolicy != model.Manifest.PolicyID {
		return fmt.Errorf("operation %s policy mismatch", exp.OperationID)
	}
	if clientRoute(op) != exp.GeneratedPath {
		return fmt.Errorf("operation %s generated path mismatch", exp.OperationID)
	}
	return nil
}
