package main

import "fmt"

func verifyOperations(model model) error {
	for _, exp := range model.Manifest.RequiredOperations {
		op := findOperation(model.Operations, exp.OperationID)
		if op.OperationID == "" {
			return fmt.Errorf("operation %s missing", exp.OperationID)
		}
		if err := verifyOperation(op, exp); err != nil {
			return err
		}
	}
	return nil
}

func verifyOperation(op operation, exp operationExpectation) error {
	if op.Kind != exp.Kind || op.Method != exp.Method || op.Path != exp.Path {
		return fmt.Errorf("operation %s tuple mismatch", exp.OperationID)
	}
	if exp.RequestRef != "" && (op.Request == nil || op.Request.Ref != exp.RequestRef) {
		return fmt.Errorf("operation %s request ref mismatch", exp.OperationID)
	}
	if op.Response == nil || op.Response.Ref != exp.ResponseRef {
		return fmt.Errorf("operation %s response ref mismatch", exp.OperationID)
	}
	if op.Response.Status != exp.ResponseStatus {
		return fmt.Errorf("operation %s response status mismatch", exp.OperationID)
	}
	return nil
}
