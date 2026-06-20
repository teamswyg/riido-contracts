package main

import "encoding/json"

type evidence struct {
	SchemaVersion        string            `json:"schema_version"`
	ID                   string            `json:"id"`
	Status               string            `json:"status"`
	Contract             string            `json:"contract"`
	ContractSchema       string            `json:"contract_schema"`
	ServiceSchema        string            `json:"service_schema"`
	AssignmentStateCount int               `json:"assignment_state_count"`
	AssignmentStateFiles int               `json:"assignment_state_files"`
	PollActionCount      int               `json:"poll_action_count"`
	PollActionFiles      int               `json:"poll_action_files"`
	TaskEventCount       int               `json:"task_event_count"`
	TaskEventFiles       int               `json:"task_event_files"`
	PayloadFieldCount    int               `json:"payload_field_count"`
	PayloadFieldFiles    int               `json:"payload_field_files"`
	PayloadFields        []payloadEvidence `json:"payload_fields"`
	EvidenceArtifact     string            `json:"evidence_artifact"`
	Workflow             string            `json:"workflow"`
	Loop                 evidenceLoop      `json:"loop"`
}

type payloadEvidence struct {
	Name      string `json:"name"`
	Source    string `json:"source"`
	MaxLength int    `json:"max_length"`
	Required  bool   `json:"required"`
	Snapshot  string `json:"snapshot"`
	Consumer  string `json:"consumer"`
}

func writeEvidence(path string, m manifest, c contract) error {
	body, err := json.MarshalIndent(newEvidence(m, c), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
