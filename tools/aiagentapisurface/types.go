package main

type manifest struct {
	SchemaVersion              string       `json:"schema_version"`
	ID                         string       `json:"id"`
	Title                      string       `json:"title"`
	RiidoTask                  string       `json:"riido_task"`
	Summary                    string       `json:"summary"`
	GeneratedDoc               string       `json:"generated_doc"`
	Workflow                   string       `json:"workflow"`
	EvidenceArtifact           string       `json:"evidence_artifact"`
	ContractID                 string       `json:"contract_id"`
	DSLFixture                 string       `json:"dsl_fixture"`
	IRFixture                  string       `json:"ir_fixture"`
	OpenAPIFixture             string       `json:"openapi_fixture"`
	ExpectedOperationCount     int          `json:"expected_operation_count"`
	ExpectedV1Operations       int          `json:"expected_v1_operation_count"`
	ExpectedV2Operations       int          `json:"expected_v2_operation_count"`
	ExpectedV2OnlyOperations   int          `json:"expected_v2_workspace_only_operation_count"`
	ExpectedOpenAPIPaths       int          `json:"expected_openapi_path_count"`
	ExpectedOpenAPIOperations  int          `json:"expected_openapi_operation_count"`
	ExpectedStreamVariantCount int          `json:"expected_client_stream_variant_count"`
	RequiredStreamVariants     []string     `json:"required_stream_variants"`
	CodegenRules               []string     `json:"codegen_rules"`
	Invariants                 []string     `json:"invariants"`
	Loop                       evidenceLoop `json:"loop"`
}
