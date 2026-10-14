package assignment

import "strings"

// ExecutionIdentity is the shared assignment execution vocabulary. It is not a
// new wire payload yet; it fixes how existing assignment fields are interpreted
// across daemon/control-plane consumers.
type ExecutionIdentity struct {
	AssignmentID      string `json:"assignment_id"`
	TaskID            string `json:"task_id"`
	ComponentID       string `json:"component_id,omitempty"`
	AgentID           string `json:"agent_id"`
	RuntimeID         string `json:"runtime_id,omitempty"`
	RunID             string `json:"run_id,omitempty"`
	ProviderSessionID string `json:"provider_session_id,omitempty"`
}

func IdentityFromAssignment(assignment Assignment, runtimeID string) ExecutionIdentity {
	runID := ExecutionIDFromAssignment(assignment)
	return ExecutionIdentity{
		AssignmentID:      strings.TrimSpace(assignment.ID),
		TaskID:            strings.TrimSpace(assignment.TaskID),
		ComponentID:       strings.TrimSpace(assignment.ComponentID),
		AgentID:           strings.TrimSpace(assignment.AgentID),
		RuntimeID:         strings.TrimSpace(runtimeID),
		RunID:             runID,
		ProviderSessionID: ResumeSessionIDForAssignment(assignment),
	}
}

func ExecutionIDFromAssignment(assignment Assignment) string {
	return firstNonEmptyTrimmed(assignment.ID, assignment.TaskID)
}

func ResumeSessionIDForAssignment(assignment Assignment) string {
	return firstNonEmptyTrimmed(assignment.ProviderSessionID, assignment.ResumeSessionID)
}

func (identity ExecutionIdentity) ExecutionID() string {
	return firstNonEmptyTrimmed(identity.AssignmentID, identity.RunID, identity.TaskID)
}

func (identity ExecutionIdentity) Valid() bool {
	return identity.ExecutionID() != "" &&
		strings.TrimSpace(identity.TaskID) != "" &&
		strings.TrimSpace(identity.AgentID) != ""
}

func firstNonEmptyTrimmed(values ...string) string {
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			return trimmed
		}
	}
	return ""
}
