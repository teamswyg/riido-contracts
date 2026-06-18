package apicontract

import (
	"errors"
	"fmt"
	"strings"
)

func validateDSLOperationIdentity(op DSLOperation, ops map[string]struct{}) error {
	if strings.TrimSpace(op.OperationID) == "" {
		return errors.New("apicontract: operation_id is required")
	}
	if _, exists := ops[op.OperationID]; exists {
		return fmt.Errorf("apicontract: duplicate operation %q", op.OperationID)
	}
	ops[op.OperationID] = struct{}{}
	if op.Kind != "query" && op.Kind != "command" {
		return fmt.Errorf("apicontract: operation %q has unsupported kind %q", op.OperationID, op.Kind)
	}
	if !methodAllowed(op.Method) {
		return fmt.Errorf("apicontract: operation %q has unsupported method %q", op.OperationID, op.Method)
	}
	if !strings.HasPrefix(op.Path, "/") {
		return fmt.Errorf("apicontract: operation %q path must start with /", op.OperationID)
	}
	return validateAuth(op.OperationID, op.Auth)
}
