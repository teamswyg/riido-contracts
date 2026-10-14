package assignment

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
