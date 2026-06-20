package main

type evidence struct {
	SchemaVersion           string       `json:"schema_version"`
	ID                      string       `json:"id"`
	Status                  string       `json:"status"`
	GeneratedDoc            string       `json:"generated_doc"`
	Workflow                string       `json:"workflow"`
	EvidenceArtifact        string       `json:"evidence_artifact"`
	OperationCount          int          `json:"operation_count"`
	OnboardingOpCount       int          `json:"onboarding_operation_count"`
	DirectCreateOpCount     int          `json:"direct_create_operation_count"`
	FixtureFieldCount       int          `json:"fixture_field_count"`
	CreateRequestFieldCount int          `json:"create_request_field_count"`
	ScenarioCount           int          `json:"scenario_count"`
	FixtureRows             []fixtureRow `json:"fixture_rows"`
	DSLIRMatch              bool         `json:"dsl_ir_match"`
	OpenAPIMatch            bool         `json:"openapi_match"`
	NoDiffPathsClean        bool         `json:"no_diff_paths_clean"`
	Loop                    evidenceLoop `json:"loop"`
}
