package assignment

import "time"

type ToolApprovalRequest struct {
	ApprovalID        string            `json:"approval_id"`
	AssignmentID      string            `json:"assignment_id"`
	TaskID            string            `json:"task_id"`
	AgentID           string            `json:"agent_id,omitempty"`
	RuntimeID         string            `json:"runtime_id,omitempty"`
	ToolID            string            `json:"tool_id"`
	ToolKind          string            `json:"tool_kind,omitempty"`
	ToolName          string            `json:"tool_name,omitempty"`
	ProviderRequestID string            `json:"provider_request_id,omitempty"`
	Reason            string            `json:"reason,omitempty"`
	Status            ApprovalStatus    `json:"status"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	RequestedAt       time.Time         `json:"requested_at"`
	ExpiresAt         time.Time         `json:"expires_at"`
}
