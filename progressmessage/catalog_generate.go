package progressmessage

import "sort"

func GenerateIR(dsl DSLDocument) (IRDocument, error) {
	if err := ValidateDSL(dsl); err != nil {
		return IRDocument{}, err
	}
	messages := append([]MessageDefinition(nil), dsl.Messages...)
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Code < messages[j].Code
	})
	return IRDocument{
		SchemaVersion:       IRSchemaVersion,
		ContractID:          dsl.ContractID,
		SourceSchemaVersion: dsl.SchemaVersion,
		AppendOnly:          dsl.AppendOnly,
		MaxMessages:         dsl.MaxMessages,
		Messages:            messages,
	}, nil
}
