package main

type manifest struct {
	SchemaVersion                  string             `json:"schema_version"`
	ID                             string             `json:"id"`
	Title                          string             `json:"title"`
	RiidoTask                      string             `json:"riido_task"`
	Summary                        string             `json:"summary"`
	GeneratedDoc                   string             `json:"generated_doc"`
	Workflow                       string             `json:"workflow"`
	EvidenceArtifact               string             `json:"evidence_artifact"`
	SourceDSL                      string             `json:"source_dsl"`
	Package                        string             `json:"package"`
	ExpectedEventCount             int                `json:"expected_event_count"`
	ExpectedTransitionCount        int                `json:"expected_transition_count"`
	ExpectedNonTransitionCount     int                `json:"expected_non_transition_count"`
	ExpectedTaskFSMTransitionCount int                `json:"expected_task_fsm_transition_count"`
	ExpectedTaskFSMTriggerCount    int                `json:"expected_task_fsm_trigger_count"`
	ExpectedCanonicalEventFields   int                `json:"expected_canonical_event_fields"`
	ExpectedReduceResultFields     int                `json:"expected_reduce_result_fields"`
	ExpectedNativeConfigCounts     nativeConfigCounts `json:"expected_native_config_counts"`
	Responsibilities               []string           `json:"responsibilities"`
	Invariants                     []string           `json:"invariants"`
	AdjacentContracts              []string           `json:"adjacent_contracts"`
	Loop                           evidenceLoop       `json:"loop"`
}
