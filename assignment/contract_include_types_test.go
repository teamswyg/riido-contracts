package assignment

type contractStateDocument struct {
	SchemaVersion string        `json:"schema_version"`
	State         contractState `json:"state"`
}

type contractNamedValueDocument struct {
	SchemaVersion string        `json:"schema_version"`
	Value         contractValue `json:"value"`
}

type contractExecutionIdentityDocument struct {
	SchemaVersion string                    `json:"schema_version"`
	ExecutionID   contractExecutionIdentity `json:"execution_identity"`
}

type contractApprovalDocument struct {
	SchemaVersion string           `json:"schema_version"`
	Approval      contractApproval `json:"approval_contract"`
}

type contractPayloadFieldDocument struct {
	SchemaVersion string                         `json:"schema_version"`
	Field         contractAssignmentPayloadField `json:"field"`
}
