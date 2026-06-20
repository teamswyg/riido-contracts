package main

type evidence struct {
	SchemaVersion          string             `json:"schema_version"`
	ID                     string             `json:"id"`
	Status                 string             `json:"status"`
	GeneratedDoc           string             `json:"generated_doc"`
	Package                string             `json:"package"`
	Workflow               string             `json:"workflow"`
	EvidenceArtifact       string             `json:"evidence_artifact"`
	EventCount             int                `json:"event_count"`
	TransitionCount        int                `json:"transition_count"`
	NonTransitionCount     int                `json:"non_transition_count"`
	TaskFSMTransitionCount int                `json:"task_fsm_transition_count"`
	TaskFSMTriggerCount    int                `json:"task_fsm_trigger_count"`
	CanonicalEventFields   int                `json:"canonical_event_fields"`
	ReduceResultFields     int                `json:"reduce_result_fields"`
	NativeConfigCounts     nativeConfigCounts `json:"native_config_counts"`
	Loop                   evidenceLoop       `json:"loop"`
}
