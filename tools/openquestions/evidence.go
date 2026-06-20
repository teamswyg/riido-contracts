package main

import "encoding/json"

type evidence struct {
	SchemaVersion    string       `json:"schema_version"`
	ID               string       `json:"id"`
	Status           string       `json:"status"`
	OpenCount        int          `json:"open_count"`
	ResolvedCount    int          `json:"resolved_count"`
	QuestionCount    int          `json:"question_count"`
	EvidenceArtifact string       `json:"evidence_artifact"`
	Workflow         string       `json:"workflow"`
	Questions        []question   `json:"questions"`
	Loop             evidenceLoop `json:"loop"`
}

func writeEvidence(path string, m manifest) error {
	body, err := json.MarshalIndent(newEvidence(m), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}

func newEvidence(m manifest) evidence {
	return evidence{
		SchemaVersion: evidenceVersion, ID: m.ID, Status: status(m),
		OpenCount: countStatus(m, "open"), ResolvedCount: countStatus(m, "resolved"),
		QuestionCount: len(m.Questions), EvidenceArtifact: m.EvidenceArtifact,
		Workflow: m.Workflow, Questions: m.Questions, Loop: m.Loop,
	}
}
