package assignment

import "time"

type Assignment struct {
	ID                       string              `json:"assignment_id"`
	TaskID                   string              `json:"task_id"`
	ComponentID              string              `json:"component_id"`
	AgentID                  string              `json:"agent_id"`
	RuntimeProvider          string              `json:"runtime_provider"`
	ModelID                  string              `json:"model_id,omitempty"`
	Prompt                   string              `json:"prompt"`
	AgentInstruction         string              `json:"agent_instruction,omitempty"`
	AllowExperimentalRuntime bool                `json:"allow_experimental_runtime,omitempty"`
	ResumeSessionID          string              `json:"resume_session_id,omitempty"`
	ProviderSessionID        string              `json:"provider_session_id,omitempty"`
	Worktree                 *AssignmentWorktree `json:"worktree,omitempty"`
	State                    AssignmentState     `json:"state"`
	LeaseToken               string              `json:"lease_token,omitempty"`
	ReplacesAssignmentID     string              `json:"replaces_assignment_id,omitempty"`
	BlockedByAssignmentID    string              `json:"blocked_by_assignment_id,omitempty"`
	CreatedAt                time.Time           `json:"created_at"`
	UpdatedAt                time.Time           `json:"updated_at"`
}
