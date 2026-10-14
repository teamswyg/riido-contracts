package main

type manifest struct {
	SchemaVersion      string        `json:"schema_version"`
	ID                 string        `json:"id"`
	RiidoTask          string        `json:"riido_task"`
	GeneratedDoc       string        `json:"generated_doc"`
	Summary            string        `json:"summary"`
	ArchitectureLinks  []namedLink   `json:"architecture_links"`
	Changes            []changeEntry `json:"changes"`
	ExternalBoundaries []string      `json:"external_boundaries"`
	OpenQuestions      namedLink     `json:"open_questions"`
	Workflow           string        `json:"workflow"`
	EvidenceArtifact   string        `json:"evidence_artifact"`
	Loop               evidenceLoop  `json:"loop"`
}

type namedLink struct {
	Label string `json:"label"`
	Path  string `json:"path"`
}

type changeEntry struct {
	Task  string   `json:"task"`
	Verb  string   `json:"verb"`
	Items []string `json:"items"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
