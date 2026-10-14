package main

import "fmt"

func verifyOperations(model model) error {
	for _, exp := range model.Manifest.RequiredOperations {
		op := findOperation(model.Operations, exp.OperationID)
		if op.OperationID == "" {
			return fmt.Errorf("operation %s missing", exp.OperationID)
		}
		if err := verifyOperation(op, exp, model.Manifest.PolicyID); err != nil {
			return err
		}
	}
	return nil
}

func verifyOperation(op operation, exp operationExpectation, policyID string) error {
	if op.Kind != exp.Kind || op.Method != exp.Method || op.Path != exp.Path {
		return fmt.Errorf("operation %s tuple mismatch", exp.OperationID)
	}
	if responseName(op) != exp.ResponseRef || responseStatus(op) != exp.ResponseStatus {
		return fmt.Errorf("operation %s response mismatch", exp.OperationID)
	}
	if op.RBACPolicy != policyID {
		return fmt.Errorf("operation %s policy = %s, want %s", exp.OperationID, op.RBACPolicy, policyID)
	}
	return nil
}
