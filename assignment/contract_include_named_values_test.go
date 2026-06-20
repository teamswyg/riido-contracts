package assignment

import "testing"

func loadContractNamedValueIncludes(t *testing.T, contract *executableContract) {
	t.Helper()
	loadContractValues(t, contract.PollActionFiles, contractPollActionSchema, &contract.PollActions)
	loadContractValues(t, contract.TaskEventFiles, contractTaskEventSchema, &contract.TaskEvents)
}

func loadContractValues(t *testing.T, files []string, schema string, out *[]contractValue) {
	t.Helper()
	for _, file := range files {
		var doc contractNamedValueDocument
		loadContractInclude(t, file, schema, &doc)
		*out = append(*out, doc.Value)
	}
}
