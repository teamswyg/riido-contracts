package apicontract

import "fmt"

func validateDSLOperations(operations []DSLOperation, components, clientModules map[string]struct{}) (map[string]string, error) {
	ops := map[string]struct{}{}
	cacheTags := map[string]string{}
	for _, op := range operations {
		if err := validateDSLOperationIdentity(op, ops); err != nil {
			return nil, err
		}
		if err := validateDSLOperationSchemas(op, components); err != nil {
			return nil, err
		}
		if err := validateDSLOperationClient(op, clientModules, cacheTags); err != nil {
			return nil, err
		}
	}
	return cacheTags, nil
}

func validateDSLOperationSchemas(op DSLOperation, components map[string]struct{}) error {
	if op.Response.Status <= 0 || op.Response.Ref == "" {
		return fmt.Errorf("apicontract: operation %q response is required", op.OperationID)
	}
	if _, ok := components[op.Response.Ref]; !ok {
		return fmt.Errorf("apicontract: operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
	}
	if op.Request != nil {
		if _, ok := components[op.Request.Ref]; !ok {
			return fmt.Errorf("apicontract: operation %q request schema %q is missing", op.OperationID, op.Request.Ref)
		}
	}
	return nil
}
