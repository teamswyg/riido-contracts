package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type evidence struct {
	SchemaVersion        string       `json:"schema_version"`
	ID                   string       `json:"id"`
	Status               string       `json:"status"`
	OwnedContextCount    int          `json:"owned_context_count"`
	NonOwnedContextCount int          `json:"non_owned_context_count"`
	DirectionRuleCount   int          `json:"direction_rule_count"`
	SSOTLinkCount        int          `json:"ssot_link_count"`
	CheckDoc             bool         `json:"check_doc"`
	EvidenceArtifact     string       `json:"evidence_artifact"`
	Loop                 evidenceLoop `json:"loop"`
}

func newEvidence(m manifest, checkDoc bool) evidence {
	return evidence{
		SchemaVersion:        evidenceSchemaVersion,
		ID:                   m.ID,
		Status:               "verified",
		OwnedContextCount:    len(m.OwnedContexts),
		NonOwnedContextCount: len(m.NonOwnedContexts),
		DirectionRuleCount:   len(m.DirectionRules),
		SSOTLinkCount:        len(m.SSOTLinks),
		CheckDoc:             checkDoc,
		EvidenceArtifact:     m.EvidenceArtifact,
		Loop:                 m.Loop,
	}
}

func writeEvidence(path string, value evidence) error {
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o644)
}
