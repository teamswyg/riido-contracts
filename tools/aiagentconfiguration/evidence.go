package main

type evidence struct {
	SchemaVersion    string       `json:"schema_version"`
	ID               string       `json:"id"`
	Status           string       `json:"status"`
	GeneratedDoc     string       `json:"generated_doc"`
	Workflow         string       `json:"workflow"`
	EvidenceArtifact string       `json:"evidence_artifact"`
	OperationCount   int          `json:"operation_count"`
	SchemaCount      int          `json:"schema_count"`
	PolicyCount      int          `json:"policy_count"`
	EnumCount        int          `json:"enum_count"`
	ScenarioCount    int          `json:"scenario_count"`
	DSLIRMatch       bool         `json:"dsl_ir_match"`
	OpenAPIMatch     bool         `json:"openapi_match"`
	StreamPass       bool         `json:"stream_pass"`
	Loop             evidenceLoop `json:"loop"`
}
