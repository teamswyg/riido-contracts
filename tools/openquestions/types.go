package main

type manifest struct {
	SchemaVersion    string       `json:"schema_version"`
	ID               string       `json:"id"`
	Title            string       `json:"title"`
	RiidoTask        string       `json:"riido_task"`
	GeneratedDoc     string       `json:"generated_doc"`
	Workflow         string       `json:"workflow"`
	EvidenceArtifact string       `json:"evidence_artifact"`
	Questions        []question   `json:"questions"`
	Loop             evidenceLoop `json:"loop"`
}

type question struct {
	ID            string `json:"id"`
	Area          string `json:"area"`
	Status        string `json:"status"`
	Question      string `json:"question"`
	CurrentStance string `json:"current_stance"`
	NextArtifact  string `json:"next_artifact"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
