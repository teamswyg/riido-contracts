package main

import (
	"encoding/json"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

type evidence struct {
	SchemaVersion    string            `json:"schema_version"`
	ID               string            `json:"id"`
	Status           string            `json:"status"`
	ContractID       string            `json:"contract_id"`
	MessageCount     int               `json:"message_count"`
	MaxMessages      int               `json:"max_messages"`
	UsageCounts      map[string]int    `json:"usage_counts"`
	GeneratedDoc     string            `json:"generated_doc"`
	EvidenceArtifact string            `json:"evidence_artifact"`
	Workflow         string            `json:"workflow"`
	DSL              string            `json:"dsl"`
	IR               string            `json:"ir"`
	Loop             evidenceLoop      `json:"loop"`
	Messages         []evidenceMessage `json:"messages"`
}

type evidenceMessage struct {
	Code     int    `json:"code"`
	Key      string `json:"key"`
	Usage    string `json:"usage"`
	Category string `json:"category"`
}

func writeEvidence(path string, m docManifest, ir progressmessage.IRDocument) error {
	body, err := json.MarshalIndent(newEvidence(m, ir), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
