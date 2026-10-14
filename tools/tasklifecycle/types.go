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
	SourceDSL                  string       `json:"source_dsl"`
	ExpectedFSMSchemaVersion   int          `json:"expected_fsm_schema_version"`
	ExpectedStateCount         int          `json:"expected_state_count"`
	ExpectedStartStateCount    int          `json:"expected_start_state_count"`
	ExpectedTerminalStateCount int          `json:"expected_terminal_state_count"`
	ExpectedTransitionCount    int          `json:"expected_transition_count"`
	Responsibilities           []string     `json:"responsibilities"`
	Invariants                 []string     `json:"invariants"`
	Loop                       evidenceLoop `json:"loop"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
