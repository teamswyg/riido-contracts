package main

type manifest struct {
	SchemaVersion          string       `json:"schema_version"`
	ID                     string       `json:"id"`
	Title                  string       `json:"title"`
	RiidoTask              string       `json:"riido_task"`
	Summary                string       `json:"summary"`
	GeneratedDoc           string       `json:"generated_doc"`
	Workflow               string       `json:"workflow"`
	EvidenceArtifact       string       `json:"evidence_artifact"`
	Fixtures               []fixtureRef `json:"fixtures"`
	Invariants             []string     `json:"invariants"`
	DeliveryRules          []string     `json:"delivery_rules"`
	Boundaries             []string     `json:"boundaries"`
	RequiredGeneratedPaths []string     `json:"required_generated_paths"`
	Loop                   evidenceLoop `json:"loop"`
}

type fixtureRef struct {
	ContractID string `json:"contract_id"`
	DSL        string `json:"dsl"`
	IR         string `json:"ir"`
	OpenAPI    string `json:"openapi"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
