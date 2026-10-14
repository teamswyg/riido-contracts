package apicontract

import (
	"errors"
	"fmt"
	"strings"
)

func validateIRHeader(ir IRDocument) error {
	if ir.SchemaVersion != IRSchemaVersion {
		return fmt.Errorf("apicontract: unsupported IR schema_version %q", ir.SchemaVersion)
	}
	if strings.TrimSpace(ir.ContractID) == "" || strings.TrimSpace(ir.Context) == "" {
		return errors.New("apicontract: IR contract_id and context are required")
	}
	return nil
}
