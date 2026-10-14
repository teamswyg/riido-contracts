package main

type manifest struct {
	SchemaVersion     string       `json:"schema_version"`
	ID                string       `json:"id"`
	Title             string       `json:"title"`
	RiidoTask         string       `json:"riido_task"`
	GeneratedDoc      string       `json:"generated_doc"`
	Workflow          string       `json:"workflow"`
	EvidenceArtifact  string       `json:"evidence_artifact"`
	ModulePath        string       `json:"module_path"`
	License           string       `json:"license"`
	LoopSource        string       `json:"loop_source"`
	Fragments         []string     `json:"fragments"`
	Summary           []string     `json:"summary"`
	Owns              []string     `json:"owns"`
	DoesNotOwn        []string     `json:"does_not_own"`
	Rationale         []string     `json:"rationale"`
	DocLinks          []docLink    `json:"doc_links"`
	Packages          []packageRef `json:"packages"`
	FSM               fsmDoc       `json:"fsm"`
	Decisions         []string     `json:"decisions"`
	Verification      []string     `json:"verification"`
	Rules             []string     `json:"rules"`
	RequiredMarkers   []string     `json:"required_markers"`
	ForbiddenLiterals []string     `json:"forbidden_literals"`
	Loop              evidenceLoop `json:"loop"`
}

type docLink struct {
	Topic string `json:"topic"`
	Path  string `json:"path"`
}

type packageRef struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
