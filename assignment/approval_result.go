package assignment

import "time"

type ToolApprovalResult struct {
	ApprovalID   string         `json:"approval_id"`
	AssignmentID string         `json:"assignment_id"`
	Status       ApprovalStatus `json:"status"`
	ResolvedAt   time.Time      `json:"resolved_at"`
}
