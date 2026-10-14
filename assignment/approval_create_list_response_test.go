package assignment

import (
	"testing"
	"time"
)

func TestAssignmentContractToolApprovalCreateResponseWireShape(t *testing.T) {
	request := approvalResponseFixture()
	assertApprovalJSON(t, ToolApprovalCreateResponse{
		SchemaVersion: SchemaVersion,
		Approval:      request,
	}, `{"schema_version":"riido-ai-server.v1","approval":{"approval_id":"approval-1","assignment_id":"asn-1","task_id":"task-1","tool_id":"tool-1","status":"pending","requested_at":"2026-06-17T10:00:00Z","expires_at":"2026-06-17T10:05:00Z"}}`)
}

func TestAssignmentContractToolApprovalListResponseWireShape(t *testing.T) {
	request := approvalResponseFixture()
	assertApprovalJSON(t, ToolApprovalListResponse{
		SchemaVersion: SchemaVersion,
		Approvals:     []ToolApprovalRequest{request},
	}, `{"schema_version":"riido-ai-server.v1","approvals":[{"approval_id":"approval-1","assignment_id":"asn-1","task_id":"task-1","tool_id":"tool-1","status":"pending","requested_at":"2026-06-17T10:00:00Z","expires_at":"2026-06-17T10:05:00Z"}]}`)
}

func approvalResponseFixture() ToolApprovalRequest {
	at := approvalWireTime()
	return ToolApprovalRequest{
		ApprovalID:   "approval-1",
		AssignmentID: "asn-1",
		TaskID:       "task-1",
		ToolID:       "tool-1",
		Status:       ApprovalPending,
		RequestedAt:  at,
		ExpiresAt:    at.Add(5 * time.Minute),
	}
}
