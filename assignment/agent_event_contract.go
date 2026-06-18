package assignment

type AgentEventRequest struct {
	AssignmentID      string            `json:"assignment_id"`
	TaskID            string            `json:"task_id"`
	DaemonID          string            `json:"daemon_id,omitempty"`
	DeviceID          string            `json:"device_id,omitempty"`
	RuntimeID         string            `json:"runtime_id,omitempty"`
	RuntimeProvider   string            `json:"runtime_provider,omitempty"`
	ModelID           string            `json:"model_id,omitempty"`
	ProviderSessionID string            `json:"provider_session_id,omitempty"`
	State             AssignmentState   `json:"state,omitempty"`
	EventType         string            `json:"event_type,omitempty"`
	Message           string            `json:"message,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
}

type AgentEventResponse struct {
	SchemaVersion string      `json:"schema_version"`
	Assignment    *Assignment `json:"assignment,omitempty"`
	Event         TaskEvent   `json:"event"`
}
