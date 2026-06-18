package assignment

type AssignRequest struct {
	ComponentID              string              `json:"component_id"`
	AgentID                  string              `json:"agent_id"`
	RuntimeProvider          string              `json:"runtime_provider"`
	ModelID                  string              `json:"model_id,omitempty"`
	Prompt                   string              `json:"prompt"`
	AgentInstruction         string              `json:"agent_instruction,omitempty"`
	AllowExperimentalRuntime bool                `json:"allow_experimental_runtime,omitempty"`
	ResumeSessionID          string              `json:"resume_session_id,omitempty"`
	Worktree                 *AssignmentWorktree `json:"worktree,omitempty"`
	CreatedBy                string              `json:"created_by,omitempty"`
}
