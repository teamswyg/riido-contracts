package main

type manifest struct {
	SchemaVersion    string       `json:"schema_version"`
	ID               string       `json:"id"`
	Title            string       `json:"title"`
	RiidoTask        string       `json:"riido_task"`
	Summary          string       `json:"summary"`
	GeneratedDoc     string       `json:"generated_doc"`
	Workflow         string       `json:"workflow"`
	EvidenceArtifact string       `json:"evidence_artifact"`
	Contract         string       `json:"contract"`
	Responsibilities []string     `json:"responsibilities"`
	Boundaries       []string     `json:"boundaries"`
	Invariants       []string     `json:"invariants"`
	Handoff          []string     `json:"handoff"`
	Loop             evidenceLoop `json:"loop"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
