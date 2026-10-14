package main

func minimalManifest() manifest {
	return manifest{
		SchemaVersion: schemaVersion,
		ID:            "domain-docs",
		RiidoTask:     "RIID-TEST",
		GeneratedDoc:  "docs/20-domain/README.md",
		Summary:       "domain docs",
		ArchitectureLinks: []namedLink{{
			Label: "Context map", Path: "context-map.md",
		}},
		Changes: []changeEntry{{
			Task: "RIID-TEST", Verb: "adds", Items: []string{"test item"},
		}},
		ExternalBoundaries: []string{"runtime code remains outside this repository"},
		OpenQuestions:      namedLink{Label: "Open questions", Path: "../50-roadmap/open-questions.md"},
		Workflow:           ".github/workflows/architecture-docs.yml",
		EvidenceArtifact:   "domain-docs-index-evidence",
		Loop: evidenceLoop{
			Observation:   "observation",
			Hypothesis:    "hypothesis",
			Execute:       "execute",
			Evaluate:      "evaluate",
			Retrospective: "retrospective",
		},
	}
}
