package main

type manifest struct {
	SchemaVersion            string          `json:"schema_version"`
	ID                       string          `json:"id"`
	RiidoTask                string          `json:"riido_task"`
	GeneratedDoc             string          `json:"generated_doc"`
	Title                    string          `json:"title"`
	Intro                    []string        `json:"intro"`
	Sections                 []policySection `json:"sections"`
	RequiredSectionOrder     []string        `json:"required_section_order"`
	RequiredPolicyAssertions []string        `json:"required_policy_assertions"`
	RequiredTerms            []string        `json:"required_terms"`
	RequiredGeneratedReaders []string        `json:"required_generated_readers"`
	Workflow                 string          `json:"workflow"`
	EvidenceArtifact         string          `json:"evidence_artifact"`
	Loop                     evidenceLoop    `json:"loop"`
}

type policySection struct {
	Title       string   `json:"title"`
	Body        []string `json:"body"`
	Subsections []string `json:"subsections"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
