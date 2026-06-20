package assignment

import "testing"

func loadContractPayloadIncludes(t *testing.T, contract *executableContract) {
	t.Helper()
	for _, file := range contract.PayloadFieldFiles {
		var doc contractPayloadFieldDocument
		loadContractInclude(t, file, contractPayloadSchema, &doc)
		contract.AssignmentPayloadFields = append(contract.AssignmentPayloadFields, doc.Field)
	}
}
