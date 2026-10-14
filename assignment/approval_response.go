package assignment

type ToolApprovalCreateResponse struct {
	SchemaVersion string              `json:"schema_version"`
	Approval      ToolApprovalRequest `json:"approval"`
}

type ToolApprovalListResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Approvals     []ToolApprovalRequest `json:"approvals"`
}

type ToolApprovalWaitRequest struct {
	AssignmentID string `json:"assignment_id"`
	WaitMs       int    `json:"wait_ms,omitempty"`
}

type ToolApprovalWaitResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Result        ToolApprovalResult    `json:"result"`
	Decision      *ToolApprovalDecision `json:"decision,omitempty"`
}

type ToolApprovalDecisionResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Result        ToolApprovalResult    `json:"result"`
	Decision      *ToolApprovalDecision `json:"decision,omitempty"`
}
