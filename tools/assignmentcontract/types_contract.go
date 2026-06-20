package main

type contract struct {
	SchemaVersion           string           `json:"schema_version"`
	ServiceSchemaVersion    string           `json:"service_schema_version"`
	AssignmentStateFiles    []string         `json:"assignment_state_files,omitempty"`
	PollActionFiles         []string         `json:"poll_action_files,omitempty"`
	TaskEventFiles          []string         `json:"task_event_files,omitempty"`
	ExecutionIdentityFile   string           `json:"execution_identity_file,omitempty"`
	ApprovalContractFile    string           `json:"approval_contract_file,omitempty"`
	PayloadFieldFiles       []string         `json:"assignment_payload_field_files,omitempty"`
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
