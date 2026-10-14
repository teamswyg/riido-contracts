package assignment

import "testing"

func loadContractSingleIncludes(t *testing.T, contract *executableContract) {
	t.Helper()
	if contract.ExecutionIdentityFile != "" {
		var doc contractExecutionIdentityDocument
		loadContractInclude(t, contract.ExecutionIdentityFile, contractExecutionIDSchema, &doc)
		contract.ExecutionIdentity = doc.ExecutionID
	}
	if contract.ApprovalContractFile != "" {
		var doc contractApprovalDocument
		loadContractInclude(t, contract.ApprovalContractFile, contractApprovalSchema, &doc)
		contract.ApprovalContract = doc.Approval
	}
}
