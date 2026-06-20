package main

type manifest struct {
	SchemaVersion          string                 `json:"schema_version"`
	ID                     string                 `json:"id"`
	Title                  string                 `json:"title"`
	RiidoTask              string                 `json:"riido_task"`
	Summary                string                 `json:"summary"`
	GeneratedDoc           string                 `json:"generated_doc"`
	Workflow               string                 `json:"workflow"`
	EvidenceArtifact       string                 `json:"evidence_artifact"`
	DSLFixture             string                 `json:"dsl_fixture"`
	IRFixture              string                 `json:"ir_fixture"`
	OpenAPIFixture         string                 `json:"openapi_fixture"`
	ExpectedOperationCount int                    `json:"expected_operation_count"`
	ExpectedSchemaCount    int                    `json:"expected_schema_count"`
	ExpectedPolicyCount    int                    `json:"expected_policy_count"`
	ExpectedScenarioCount  int                    `json:"expected_scenario_count"`
	RequiredOperations     []operationExpectation `json:"required_operations"`
	RequiredPolicies       []string               `json:"required_policies"`
	RequiredSchemaFields   []schemaExpectation    `json:"required_schema_fields"`
	ForbiddenRequestFields []string               `json:"forbidden_request_fields"`
	NoDiffPathFragments    []string               `json:"no_diff_path_fragments"`
	Invariants             []string               `json:"invariants"`
	Loop                   evidenceLoop           `json:"loop"`
}

type operationExpectation struct {
	OperationID    string `json:"operation_id"`
	Kind           string `json:"kind"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	RequestRef     string `json:"request_ref"`
	ResponseStatus int    `json:"response_status"`
	ResponseRef    string `json:"response_ref"`
}

type schemaExpectation struct {
	Schema string   `json:"schema"`
	Fields []string `json:"fields"`
}
