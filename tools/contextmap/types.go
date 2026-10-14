package main

type manifest struct {
	SchemaVersion    string            `json:"schema_version"`
	ID               string            `json:"id"`
	RiidoTask        string            `json:"riido_task"`
	GeneratedDoc     string            `json:"generated_doc"`
	Module           string            `json:"module"`
	Role             string            `json:"role"`
	BoundaryRule     string            `json:"boundary_rule"`
	OwnedContexts    []ownedContext    `json:"owned_contexts"`
	NonOwnedContexts []nonOwnedContext `json:"non_owned_contexts"`
	DirectionRules   []string          `json:"direction_rules"`
	SSOTLinks        []ssotLink        `json:"ssot_links"`
	Workflow         string            `json:"workflow"`
	EvidenceArtifact string            `json:"evidence_artifact"`
	Loop             evidenceLoop      `json:"loop"`
}

type ownedContext struct {
	Context        string `json:"context"`
	Package        string `json:"package"`
	Responsibility string `json:"responsibility"`
}

type nonOwnedContext struct {
	Context  string `json:"context"`
	Owner    string `json:"owner"`
	Boundary string `json:"boundary"`
}

type ssotLink struct {
	Label string `json:"label"`
	Path  string `json:"path"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
