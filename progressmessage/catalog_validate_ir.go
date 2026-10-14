package progressmessage

import (
	"errors"
	"fmt"
)

func ValidateIR(ir IRDocument) error {
	if ir.SchemaVersion != IRSchemaVersion {
		return fmt.Errorf("progressmessage: unsupported IR schema_version %q", ir.SchemaVersion)
	}
	if ir.SourceSchemaVersion != DSLSchemaVersion {
		return fmt.Errorf("progressmessage: source_schema_version = %q, want %q", ir.SourceSchemaVersion, DSLSchemaVersion)
	}
	if ir.ContractID != ContractID {
		return fmt.Errorf("progressmessage: contract_id = %q, want %q", ir.ContractID, ContractID)
	}
	if !ir.AppendOnly {
		return errors.New("progressmessage: append_only must be true")
	}
	if ir.MaxMessages != MaxMessages {
		return fmt.Errorf("progressmessage: max_messages = %d, want %d", ir.MaxMessages, MaxMessages)
	}
	if !messagesSorted(ir.Messages) {
		return errors.New("progressmessage: IR messages must be sorted by code")
	}
	return validateMessages(ir.Messages)
}
