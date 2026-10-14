package main

type stateDocument struct {
	SchemaVersion string `json:"schema_version"`
	State         state  `json:"state"`
}

type namedValueDocument struct {
	SchemaVersion string     `json:"schema_version"`
	Value         namedValue `json:"value"`
}

type executionIDDocument struct {
	SchemaVersion string      `json:"schema_version"`
	ExecutionID   executionID `json:"execution_identity"`
}

type approvalContractDocument struct {
	SchemaVersion string           `json:"schema_version"`
	Approval      approvalContract `json:"approval_contract"`
}

type payloadFieldDocument struct {
	SchemaVersion string       `json:"schema_version"`
	Field         payloadField `json:"field"`
}
