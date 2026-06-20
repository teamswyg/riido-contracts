package main

import "encoding/json"

type evidence struct {
	SchemaVersion        string       `json:"schema_version"`
	ID                   string       `json:"id"`
	Status               string       `json:"status"`
	Contract             string       `json:"contract"`
	ContractSchema       string       `json:"contract_schema"`
	ServiceSchema        string       `json:"service_schema"`
	AssignmentStateCount int          `json:"assignment_state_count"`
	PollActionCount      int          `json:"poll_action_count"`
	TaskEventCount       int          `json:"task_event_count"`
	PayloadFieldCount    int          `json:"payload_field_count"`
	EvidenceArtifact     string       `json:"evidence_artifact"`
	Workflow             string       `json:"workflow"`
	Loop                 evidenceLoop `json:"loop"`
}

func writeEvidence(path string, m manifest, c contract) error {
	body, err := json.MarshalIndent(newEvidence(m, c), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
