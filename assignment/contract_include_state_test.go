package assignment

import "testing"

func loadContractStateIncludes(t *testing.T, contract *executableContract) {
	t.Helper()
	for _, file := range contract.AssignmentStateFiles {
		var doc contractStateDocument
		loadContractInclude(t, file, contractStateSchema, &doc)
		contract.AssignmentStates = append(contract.AssignmentStates, doc.State)
	}
}
