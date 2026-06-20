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
	PolicyRuleCount  int          `json:"policy_rule_count"`
	ScenarioCount    int          `json:"scenario_count"`
	DSLIRMatch       bool         `json:"dsl_ir_match"`
	OpenAPIMatch     bool         `json:"openapi_match"`
	MapShapePass     bool         `json:"map_shape_pass"`
	Loop             evidenceLoop `json:"loop"`
}
