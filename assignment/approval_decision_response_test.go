package assignment

import "testing"

func TestAssignmentContractToolApprovalWaitRequestWireShape(t *testing.T) {
	assertApprovalJSON(t, ToolApprovalWaitRequest{
		AssignmentID: "asn-1",
		WaitMs:       30000,
	}, `{"assignment_id":"asn-1","wait_ms":30000}`)
}

func TestAssignmentContractToolApprovalWaitResponseWireShape(t *testing.T) {
	result, decision := approvalDecisionResponseFixture()
	assertApprovalJSON(t, ToolApprovalWaitResponse{
		SchemaVersion: SchemaVersion,
		Result:        result,
		Decision:      decision,
	}, `{"schema_version":"riido-ai-server.v1","result":{"approval_id":"approval-1","assignment_id":"asn-1","status":"denied","resolved_at":"2026-06-17T10:00:00Z"},"decision":{"approval_id":"approval-1","assignment_id":"asn-1","decision":"deny","decided_by":"user-1","decided_at":"2026-06-17T10:00:00Z"}}`)
}

func TestAssignmentContractToolApprovalDecisionResponseWireShape(t *testing.T) {
	result, decision := approvalDecisionResponseFixture()
	assertApprovalJSON(t, ToolApprovalDecisionResponse{
		SchemaVersion: SchemaVersion,
		Result:        result,
		Decision:      decision,
	}, `{"schema_version":"riido-ai-server.v1","result":{"approval_id":"approval-1","assignment_id":"asn-1","status":"denied","resolved_at":"2026-06-17T10:00:00Z"},"decision":{"approval_id":"approval-1","assignment_id":"asn-1","decision":"deny","decided_by":"user-1","decided_at":"2026-06-17T10:00:00Z"}}`)
}

func approvalDecisionResponseFixture() (ToolApprovalResult, *ToolApprovalDecision) {
	at := approvalWireTime()
	decision := &ToolApprovalDecision{
		ApprovalID:   "approval-1",
		AssignmentID: "asn-1",
		Decision:     ApprovalDecisionDeny,
		DecidedBy:    "user-1",
		DecidedAt:    at,
	}
	result := ToolApprovalResult{
		ApprovalID:   "approval-1",
		AssignmentID: "asn-1",
		Status:       ApprovalDenied,
		ResolvedAt:   at,
	}
	return result, decision
}
