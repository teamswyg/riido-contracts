package assignment

type executableContract struct {
	SchemaVersion           string                           `json:"schema_version"`
	ServiceSchemaVersion    string                           `json:"service_schema_version"`
	AssignmentStateFiles    []string                         `json:"assignment_state_files,omitempty"`
	PollActionFiles         []string                         `json:"poll_action_files,omitempty"`
	TaskEventFiles          []string                         `json:"task_event_files,omitempty"`
	ExecutionIdentityFile   string                           `json:"execution_identity_file,omitempty"`
	ApprovalContractFile    string                           `json:"approval_contract_file,omitempty"`
	PayloadFieldFiles       []string                         `json:"assignment_payload_field_files,omitempty"`
	AssignmentStates        []contractState                  `json:"assignment_states"`
	PollActions             []contractValue                  `json:"poll_actions"`
	TaskEvents              []contractValue                  `json:"task_events"`
	ExecutionIdentity       contractExecutionIdentity        `json:"execution_identity"`
	ApprovalContract        contractApproval                 `json:"approval_contract"`
	AssignmentPayloadFields []contractAssignmentPayloadField `json:"assignment_payload_fields"`
}

type contractState struct {
	Name        string   `json:"name"`
	Value       string   `json:"value"`
	AgentActive bool     `json:"agent_active"`
	Terminal    bool     `json:"terminal"`
	Transitions []string `json:"transitions"`
}

type contractValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type contractExecutionIdentity struct {
	ExecutionKeyOrder     []string `json:"execution_key_order"`
	ResumeSessionKeyOrder []string `json:"resume_session_key_order"`
	RunIDDefaultSource    string   `json:"run_id_default_source"`
}

type contractApproval struct {
	Owner                 string                `json:"owner"`
	TimeoutTerminalStatus string                `json:"timeout_terminal_status"`
	Statuses              []contractStatusValue `json:"statuses"`
	Decisions             []contractValue       `json:"decisions"`
}

type contractStatusValue struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Terminal bool   `json:"terminal"`
}

type contractAssignmentPayloadField struct {
	Name      string `json:"name"`
	Source    string `json:"source"`
	MaxLength int    `json:"max_length"`
	Required  bool   `json:"required"`
	Snapshot  string `json:"snapshot"`
	Consumer  string `json:"consumer"`
}
