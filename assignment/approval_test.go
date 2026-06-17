package assignment

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAssignmentContractToolApprovalWireShape(t *testing.T) {
	at := time.Date(2026, 6, 17, 10, 0, 0, 0, time.UTC)
	request := ToolApprovalRequest{
		ApprovalID:        "approval-1",
		AssignmentID:      "asn-1",
		TaskID:            "task-1",
		AgentID:           "agent-1",
		RuntimeID:         "runtime-1",
		ToolID:            "tool-1",
		ToolKind:          "patch_apply",
		ToolName:          "apply_patch",
		ProviderRequestID: "provider-approval-1",
		Reason:            "protected path write",
		Status:            ApprovalPending,
		Metadata:          map[string]string{"surface": "protected_path_write"},
		RequestedAt:       at,
		ExpiresAt:         at.Add(5 * time.Minute),
	}
	assertApprovalJSON(t, request, `{"approval_id":"approval-1","assignment_id":"asn-1","task_id":"task-1","agent_id":"agent-1","runtime_id":"runtime-1","tool_id":"tool-1","tool_kind":"patch_apply","tool_name":"apply_patch","provider_request_id":"provider-approval-1","reason":"protected path write","status":"pending","metadata":{"surface":"protected_path_write"},"requested_at":"2026-06-17T10:00:00Z","expires_at":"2026-06-17T10:05:00Z"}`)

	decision := ToolApprovalDecision{
		ApprovalID:   "approval-1",
		AssignmentID: "asn-1",
		Decision:     ApprovalDecisionApprove,
		DecidedBy:    "user-1",
		Reason:       "reviewed in web",
		DecidedAt:    at.Add(time.Minute),
	}
	assertApprovalJSON(t, decision, `{"approval_id":"approval-1","assignment_id":"asn-1","decision":"approve","decided_by":"user-1","reason":"reviewed in web","decided_at":"2026-06-17T10:01:00Z"}`)
}

func TestAssignmentContractToolApprovalEnumRoundTrip(t *testing.T) {
	for _, status := range AllApprovalStatuses() {
		code := status.Code()
		if !code.IsKnown() || code.ApprovalStatus() != status {
			t.Fatalf("approval status round trip failed for %q", status)
		}
	}
	if !ApprovalTimedOut.IsTerminal() || !ApprovalApproved.IsTerminal() || !ApprovalDenied.IsTerminal() {
		t.Fatal("resolved approval statuses must be terminal")
	}
	if ApprovalPending.IsTerminal() {
		t.Fatal("pending approval must not be terminal")
	}
	for _, decision := range AllApprovalDecisions() {
		code := decision.Code()
		if !code.IsKnown() || code.ApprovalDecision() != decision {
			t.Fatalf("approval decision round trip failed for %q", decision)
		}
	}
}

func assertApprovalJSON(t *testing.T, value any, want string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if got := string(data); got != want {
		t.Fatalf("json = %s, want %s", got, want)
	}
}
