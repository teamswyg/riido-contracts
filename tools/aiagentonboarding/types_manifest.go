package main

type manifest struct {
	SchemaVersion                   string                 `json:"schema_version"`
	ID                              string                 `json:"id"`
	Title                           string                 `json:"title"`
	RiidoTask                       string                 `json:"riido_task"`
	Summary                         string                 `json:"summary"`
	GeneratedDoc                    string                 `json:"generated_doc"`
	Workflow                        string                 `json:"workflow"`
	EvidenceArtifact                string                 `json:"evidence_artifact"`
	DSLFixture                      string                 `json:"dsl_fixture"`
	IRFixture                       string                 `json:"ir_fixture"`
	OpenAPIFixture                  string                 `json:"openapi_fixture"`
	ExpectedOperationCount          int                    `json:"expected_operation_count"`
	ExpectedOnboardingOpCount       int                    `json:"expected_onboarding_operation_count"`
	ExpectedDirectCreateOpCount     int                    `json:"expected_direct_create_operation_count"`
	ExpectedFixtureFieldCount       int                    `json:"expected_fixture_field_count"`
	ExpectedCreateRequestFieldCount int                    `json:"expected_create_request_field_count"`
	ExpectedScenarioCount           int                    `json:"expected_scenario_count"`
	RequiredOperations              []operationExpectation `json:"required_operations"`
	RequiredFixtureFields           []string               `json:"required_fixture_fields"`
	RequiredCreateRequestFields     []string               `json:"required_create_request_fields"`
	FixtureRows                     []fixtureRow           `json:"fixture_rows"`
	NoDiffPathFragments             []string               `json:"no_diff_path_fragments"`
	Invariants                      []string               `json:"invariants"`
	Loop                            evidenceLoop           `json:"loop"`
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

type fixtureRow struct {
	Name     string `json:"name"`
	TmpColor string `json:"tmp_color"`
}
