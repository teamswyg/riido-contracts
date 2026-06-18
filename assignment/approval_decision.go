package assignment

import "time"

type ToolApprovalDecision struct {
	ApprovalID   string           `json:"approval_id"`
	AssignmentID string           `json:"assignment_id"`
	Decision     ApprovalDecision `json:"decision"`
	DecidedBy    string           `json:"decided_by,omitempty"`
	Reason       string           `json:"reason,omitempty"`
	DecidedAt    time.Time        `json:"decided_at"`
}
