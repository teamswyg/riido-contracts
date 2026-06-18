package apicontract

import "fmt"

func validateIROperationRefs(operations []IROperation, components map[string]struct{}) error {
	for _, op := range operations {
		if err := validateAuth(op.OperationID, op.Auth); err != nil {
			return err
		}
		if _, ok := components[op.Response.Ref]; !ok {
			return fmt.Errorf("apicontract: IR operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
		}
		if op.Request != nil {
			if _, ok := components[op.Request.Ref]; !ok {
				return fmt.Errorf("apicontract: IR operation %q request schema %q is missing", op.OperationID, op.Request.Ref)
			}
		}
	}
	return nil
}
