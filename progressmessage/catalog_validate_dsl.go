package progressmessage

import (
	"errors"
	"fmt"
)

func ValidateDSL(dsl DSLDocument) error {
	if dsl.SchemaVersion != DSLSchemaVersion {
		return fmt.Errorf("progressmessage: unsupported DSL schema_version %q", dsl.SchemaVersion)
	}
	if dsl.ContractID != ContractID {
		return fmt.Errorf("progressmessage: contract_id = %q, want %q", dsl.ContractID, ContractID)
	}
	if !dsl.AppendOnly {
		return errors.New("progressmessage: append_only must be true")
	}
	if dsl.MaxMessages != MaxMessages {
		return fmt.Errorf("progressmessage: max_messages = %d, want %d", dsl.MaxMessages, MaxMessages)
	}
	return validateMessages(dsl.Messages)
}
