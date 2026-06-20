package main

type evidence struct {
	SchemaVersion     string           `json:"schema_version"`
	ID                string           `json:"id"`
	Status            string           `json:"status"`
	GeneratedDoc      string           `json:"generated_doc"`
	Workflow          string           `json:"workflow"`
	EvidenceArtifact  string           `json:"evidence_artifact"`
	OperationCount    int              `json:"operation_count"`
	V1OperationCount  int              `json:"v1_operation_count"`
	V2OperationCount  int              `json:"v2_operation_count"`
	V2OnlyOperations  []operationTuple `json:"v2_only_operations"`
	OpenAPIPathCount  int              `json:"openapi_path_count"`
	OpenAPIOpCount    int              `json:"openapi_operation_count"`
	StreamVariants    []string         `json:"client_stream_variants"`
	DSLIRMatch        bool             `json:"dsl_ir_match"`
	IROpenAPIMatch    bool             `json:"ir_openapi_match"`
	V2CoversV1        bool             `json:"v2_covers_v1"`
	StreamVariantPass bool             `json:"stream_variant_pass"`
	Loop              evidenceLoop     `json:"loop"`
}
