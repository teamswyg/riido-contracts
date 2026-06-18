package assignment

import "testing"

func TestAssignmentContractToolApprovalStatusRoundTrip(t *testing.T) {
	for _, status := range AllApprovalStatuses() {
		code := status.Code()
		if !code.IsKnown() || code.ApprovalStatus() != status {
			t.Fatalf("approval status round trip failed for %q", status)
		}
	}
}

func TestAssignmentContractToolApprovalStatusTerminality(t *testing.T) {
	if !ApprovalTimedOut.IsTerminal() || !ApprovalApproved.IsTerminal() || !ApprovalDenied.IsTerminal() {
		t.Fatal("resolved approval statuses must be terminal")
	}
	if ApprovalPending.IsTerminal() {
		t.Fatal("pending approval must not be terminal")
	}
}

func TestAssignmentContractToolApprovalDecisionRoundTrip(t *testing.T) {
	for _, decision := range AllApprovalDecisions() {
		code := decision.Code()
		if !code.IsKnown() || code.ApprovalDecision() != decision {
			t.Fatalf("approval decision round trip failed for %q", decision)
		}
	}
}
