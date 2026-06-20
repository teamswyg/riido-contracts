package main

type contract struct {
	SchemaVersion           string           `json:"schema_version"`
	ServiceSchemaVersion    string           `json:"service_schema_version"`
	AssignmentStates        []state          `json:"assignment_states"`
	PollActions             []namedValue     `json:"poll_actions"`
	TaskEvents              []namedValue     `json:"task_events"`
	ExecutionIdentity       executionID      `json:"execution_identity"`
	ApprovalContract        approvalContract `json:"approval_contract"`
	AssignmentPayloadFields []payloadField   `json:"assignment_payload_fields"`
}

type state struct {
	Name        string   `json:"name"`
	Value       string   `json:"value"`
	AgentActive bool     `json:"agent_active"`
	Terminal    bool     `json:"terminal"`
	Transitions []string `json:"transitions"`
}

type namedValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type executionID struct {
	ExecutionKeyOrder     []string `json:"execution_key_order"`
	ResumeSessionKeyOrder []string `json:"resume_session_key_order"`
	RunIDDefaultSource    string   `json:"run_id_default_source"`
}
