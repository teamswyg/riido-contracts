package assignment

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
