package assignment

import "testing"

func assertApprovalContract(t *testing.T, contract contractApproval) {
	t.Helper()
	if contract.Owner != ApprovalContractOwner {
		t.Fatalf("approval owner drifted: %q", contract.Owner)
	}
	if contract.TimeoutTerminalStatus != string(ApprovalTimeoutTerminalStatus) {
		t.Fatalf("approval timeout terminal status drifted: %q", contract.TimeoutTerminalStatus)
	}
	for _, status := range contract.Statuses {
		value := ApprovalStatus(status.Value)
		if !value.Code().IsKnown() || value.IsTerminal() != status.Terminal {
			t.Fatalf("approval status drifted: %#v", status)
		}
	}
	for _, decision := range contract.Decisions {
		value := ApprovalDecision(decision.Value)
		if !value.Code().IsKnown() {
			t.Fatalf("approval decision drifted: %#v", decision)
		}
	}
}
