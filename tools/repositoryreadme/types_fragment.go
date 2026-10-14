package main

type manifestFragment struct {
	SchemaVersion     string       `json:"schema_version"`
	ID                string       `json:"id"`
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
