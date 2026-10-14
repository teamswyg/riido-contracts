package main

import "encoding/json"

type evidence struct {
	SchemaVersion        string       `json:"schema_version"`
	ID                   string       `json:"id"`
	Status               string       `json:"status"`
	GeneratedDoc         string       `json:"generated_doc"`
	SectionCount         int          `json:"section_count"`
	PolicyAssertionCount int          `json:"policy_assertion_count"`
	TermCount            int          `json:"term_count"`
	FigmaNodeRefCount    int          `json:"figma_node_ref_count"`
	APIPathRefCount      int          `json:"api_path_ref_count"`
	GeneratedReaderCount int          `json:"generated_reader_count"`
	EvidenceArtifact     string       `json:"evidence_artifact"`
	Loop                 evidenceLoop `json:"loop"`
}

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:        evidenceSchemaVersion,
		ID:                   m.ID,
		Status:               "verified",
		GeneratedDoc:         m.GeneratedDoc,
		SectionCount:         len(m.Sections),
		PolicyAssertionCount: model.PolicyAssertionCount,
		TermCount:            model.TermCount,
		FigmaNodeRefCount:    model.FigmaNodeRefCount,
		APIPathRefCount:      model.APIPathRefCount,
		GeneratedReaderCount: model.GeneratedReaderCount,
		EvidenceArtifact:     m.EvidenceArtifact,
		Loop:                 m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
