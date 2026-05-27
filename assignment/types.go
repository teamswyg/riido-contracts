package assignment

import "time"

const (
	// ContractSchemaVersion identifies the executable assignment contract
	// fixture in assignment_contract.riido.json.
	ContractSchemaVersion = "riido-ai-server-contract.v1"

	// SchemaVersion is the C10 SaaS assignment API schema version shared by
	// daemon poll/event clients and the control-plane API surface.
	SchemaVersion = "riido-ai-server.v1"
)

type AssignmentState string

const (
	AssignmentQueued     AssignmentState = "queued"
	AssignmentLeased     AssignmentState = "leased"
	AssignmentReady      AssignmentState = "ready"
	AssignmentRunning    AssignmentState = "running"
	AssignmentCancelling AssignmentState = "cancelling"
	AssignmentCancelled  AssignmentState = "cancelled"
	AssignmentCompleted  AssignmentState = "completed"
	AssignmentFailed     AssignmentState = "failed"
)

func AllAssignmentStates() []AssignmentState {
	return []AssignmentState{
		AssignmentQueued,
		AssignmentLeased,
		AssignmentReady,
		AssignmentRunning,
		AssignmentCancelling,
		AssignmentCancelled,
		AssignmentCompleted,
		AssignmentFailed,
	}
}

func (s AssignmentState) Valid() bool {
	switch s {
	case AssignmentQueued,
		AssignmentLeased,
		AssignmentReady,
		AssignmentRunning,
		AssignmentCancelling,
		AssignmentCancelled,
		AssignmentCompleted,
		AssignmentFailed:
		return true
	default:
		return false
	}
}

type PollAction string

const (
	PollNone   PollAction = "none"
	PollStart  PollAction = "start"
	PollCancel PollAction = "cancel"
	PollActive PollAction = "active"
)

func AllPollActions() []PollAction {
	return []PollAction{
		PollNone,
		PollStart,
		PollCancel,
		PollActive,
	}
}

func (a PollAction) Valid() bool {
	switch a {
	case PollNone,
		PollStart,
		PollCancel,
		PollActive:
		return true
	default:
		return false
	}
}

const (
	EventAssignmentQueued       = "assignment_queued"
	EventAssignmentLeased       = "assignment_leased"
	EventAssignmentReady        = "assignment_ready"
	EventAssignmentRunning      = "assignment_running"
	EventAssignmentCancelling   = "assignment_cancelling"
	EventAssignmentCancelled    = "assignment_cancelled"
	EventAssignmentCompleted    = "assignment_completed"
	EventAssignmentFailed       = "assignment_failed"
	EventAssignmentStateUpdated = "assignment_state_updated"
	EventRiidoLog               = "riido_log"
	EventProviderLog            = "provider_log"
	EventProviderWarning        = "provider_warning"
	EventProviderError          = "provider_error"
)

func AllTaskEventTypes() []string {
	return []string{
		EventAssignmentQueued,
		EventAssignmentLeased,
		EventAssignmentReady,
		EventAssignmentRunning,
		EventAssignmentCancelling,
		EventAssignmentCancelled,
		EventAssignmentCompleted,
		EventAssignmentFailed,
		EventAssignmentStateUpdated,
		EventRiidoLog,
		EventProviderLog,
		EventProviderWarning,
		EventProviderError,
	}
}

type AssignRequest struct {
	ComponentID     string `json:"component_id"`
	AgentID         string `json:"agent_id"`
	RuntimeProvider string `json:"runtime_provider"`
	Prompt          string `json:"prompt"`
	CreatedBy       string `json:"created_by,omitempty"`
}

type Assignment struct {
	ID                    string          `json:"assignment_id"`
	TaskID                string          `json:"task_id"`
	ComponentID           string          `json:"component_id"`
	AgentID               string          `json:"agent_id"`
	RuntimeProvider       string          `json:"runtime_provider"`
	Prompt                string          `json:"prompt"`
	State                 AssignmentState `json:"state"`
	LeaseToken            string          `json:"lease_token,omitempty"`
	ReplacesAssignmentID  string          `json:"replaces_assignment_id,omitempty"`
	BlockedByAssignmentID string          `json:"blocked_by_assignment_id,omitempty"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
}

type PollRequest struct {
	DaemonID  string `json:"daemon_id"`
	DeviceID  string `json:"device_id"`
	RuntimeID string `json:"runtime_id"`
}

type PollResponse struct {
	SchemaVersion string      `json:"schema_version"`
	Action        PollAction  `json:"action"`
	Assignment    *Assignment `json:"assignment,omitempty"`
}

type AgentHeartbeatRequest struct {
	DaemonID            string   `json:"daemon_id"`
	DeviceID            string   `json:"device_id"`
	RuntimeID           string   `json:"runtime_id"`
	RunningTaskIDs      []string `json:"running_task_ids,omitempty"`
	ActiveAssignmentIDs []string `json:"active_assignment_ids,omitempty"`
}

type AgentHeartbeatResponse struct {
	SchemaVersion        string       `json:"schema_version"`
	RefreshedAssignments []Assignment `json:"refreshed_assignments,omitempty"`
}

type AgentEventRequest struct {
	AssignmentID string            `json:"assignment_id"`
	TaskID       string            `json:"task_id"`
	DaemonID     string            `json:"daemon_id,omitempty"`
	DeviceID     string            `json:"device_id,omitempty"`
	RuntimeID    string            `json:"runtime_id,omitempty"`
	State        AssignmentState   `json:"state,omitempty"`
	EventType    string            `json:"event_type,omitempty"`
	Message      string            `json:"message,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type AgentEventResponse struct {
	SchemaVersion string      `json:"schema_version"`
	Assignment    *Assignment `json:"assignment,omitempty"`
	Event         TaskEvent   `json:"event"`
}

type TaskEvent struct {
	Seq          int64             `json:"seq"`
	TaskID       string            `json:"task_id"`
	AssignmentID string            `json:"assignment_id"`
	AgentID      string            `json:"agent_id"`
	Type         string            `json:"type"`
	State        AssignmentState   `json:"state,omitempty"`
	Message      string            `json:"message,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	At           time.Time         `json:"at"`
}

// AgentRuntimeBinding is the shared DTO that binds a SaaS agent identity to a
// customer daemon runtime slot. Registry storage and validation rules remain in
// the control-plane package that owns assignment routing behavior.
type AgentRuntimeBinding struct {
	AgentID         string `json:"agent_id"`
	DaemonID        string `json:"daemon_id"`
	DeviceID        string `json:"device_id,omitempty"`
	RuntimeID       string `json:"runtime_id"`
	RuntimeProvider string `json:"runtime_provider"`
}
