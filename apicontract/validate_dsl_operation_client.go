package apicontract

import (
	"fmt"
	"strings"
)

func validateDSLOperationClient(op DSLOperation, clientModules map[string]struct{}, cacheTags map[string]string) error {
	if len(clientModules) == 0 {
		if op.Client != nil {
			return fmt.Errorf("apicontract: operation %q declares client metadata without client_modules", op.OperationID)
		}
		return nil
	}
	if op.Client == nil {
		return fmt.Errorf("apicontract: operation %q missing client metadata", op.OperationID)
	}
	if err := validateClientMeta(op.OperationID, strings.ToUpper(op.Method), *op.Client, clientModules); err != nil {
		return err
	}
	if strings.EqualFold(op.Method, "GET") {
		if prev, exists := cacheTags[op.Client.CacheTag]; exists {
			return fmt.Errorf("apicontract: duplicate client cache_tag %q on %s and %s", op.Client.CacheTag, prev, op.OperationID)
		}
		cacheTags[op.Client.CacheTag] = op.OperationID
	}
	return nil
}

func validateDSLOperationInvalidations(operations []DSLOperation, clientModules map[string]struct{}, cacheTags map[string]string) error {
	if len(clientModules) == 0 {
		return nil
	}
	for _, op := range operations {
		if op.Client == nil {
			continue
		}
		for _, tag := range op.Client.Invalidates {
			if _, ok := cacheTags[tag]; !ok {
				return fmt.Errorf("apicontract: operation %q invalidates unknown client cache_tag %q", op.OperationID, tag)
			}
		}
	}
	return nil
}
