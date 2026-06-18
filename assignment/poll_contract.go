package assignment

type PollRequest struct {
	DaemonID  string `json:"daemon_id"`
	DeviceID  string `json:"device_id"`
	RuntimeID string `json:"runtime_id"`
	// WaitMs is an optional long-poll hint in milliseconds. Zero/omitted
	// preserves the legacy point-in-time short-poll wire shape.
	WaitMs int `json:"wait_ms,omitempty"`
}

type PollResponse struct {
	SchemaVersion string      `json:"schema_version"`
	Action        PollAction  `json:"action"`
	Assignment    *Assignment `json:"assignment,omitempty"`
}
