package main

import "encoding/json"

type evidence struct {
	SchemaVersion string       `json:"schema_version"`
	ID            string       `json:"id"`
	Status        string       `json:"status"`
	GeneratedDoc  string       `json:"generated_doc"`
	SourceDSL     string       `json:"source_dsl"`
	Workflow      string       `json:"workflow"`
	Artifact      string       `json:"evidence_artifact"`
	FSMSchema     int          `json:"fsm_schema_version"`
	States        int          `json:"state_count"`
	StartStates   int          `json:"start_state_count"`
	Terminals     int          `json:"terminal_state_count"`
	Transitions   int          `json:"transition_count"`
	Invariants    int          `json:"invariant_count"`
	Loop          evidenceLoop `json:"loop"`
}

func writeEvidence(path string, model model) error {
	body, err := json.MarshalIndent(newEvidence(model), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
