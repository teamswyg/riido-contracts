package apicontract

import (
	"fmt"
	"strings"
)

func validateDSLHeader(dsl DSLDocument) error {
	if dsl.SchemaVersion != DSLSchemaVersion {
		return fmt.Errorf("apicontract: unsupported DSL schema_version %q", dsl.SchemaVersion)
	}
	for name, value := range map[string]string{
		"contract_id":            dsl.ContractID,
		"context":                dsl.Context,
		"service.name":           dsl.Service.Name,
		"service.schema_version": dsl.Service.SchemaVersion,
	} {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("apicontract: %s is required", name)
		}
	}
	return nil
}
