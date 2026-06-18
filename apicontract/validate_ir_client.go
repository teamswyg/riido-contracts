package apicontract

import (
	"fmt"
	"strings"
)

func validateIRClientMetadata(modules []ClientModule, operations []IROperation) error {
	if len(modules) == 0 {
		return nil
	}
	clientModules, err := validateClientModules(modules)
	if err != nil {
		return err
	}
	cacheTags := map[string]string{}
	for _, op := range operations {
		if err := validateIRClientOperation(op, clientModules, cacheTags); err != nil {
			return err
		}
	}
	return validateIRClientInvalidations(operations, cacheTags)
}

func validateIRClientOperation(op IROperation, clientModules map[string]struct{}, cacheTags map[string]string) error {
	if op.Client == nil {
		return fmt.Errorf("apicontract: IR operation %q missing client metadata", op.OperationID)
	}
	if err := validateClientMeta(op.OperationID, op.Method, *op.Client, clientModules); err != nil {
		return err
	}
	if strings.EqualFold(op.Method, "GET") {
		if prev, exists := cacheTags[op.Client.CacheTag]; exists {
			return fmt.Errorf("apicontract: duplicate IR client cache_tag %q on %s and %s", op.Client.CacheTag, prev, op.OperationID)
		}
		cacheTags[op.Client.CacheTag] = op.OperationID
	}
	return nil
}

func validateIRClientInvalidations(operations []IROperation, cacheTags map[string]string) error {
	for _, op := range operations {
		for _, tag := range op.Client.Invalidates {
			if _, ok := cacheTags[tag]; !ok {
				return fmt.Errorf("apicontract: IR operation %q invalidates unknown client cache_tag %q", op.OperationID, tag)
			}
		}
	}
	return nil
}
