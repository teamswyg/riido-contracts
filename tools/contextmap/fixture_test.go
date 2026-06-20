package main

func minimalManifest() manifest {
	return manifest{
		SchemaVersion: schemaVersion,
		ID:            "context-map",
		RiidoTask:     "RIID-TEST",
		GeneratedDoc:  "docs/20-domain/context-map.md",
		Module:        "github.com/teamswyg/riido-contracts",
		Role:          "shared facts",
		BoundaryRule:  "no runtime imports",
		OwnedContexts: []ownedContext{{
			Context: "Task", Package: "task", Responsibility: "task states",
		}},
		NonOwnedContexts: []nonOwnedContext{{
			Context: "Runtime", Owner: "riido-daemon", Boundary: "process execution",
		}},
		DirectionRules:   []string{"task may import ir"},
		SSOTLinks:        []ssotLink{{Label: "Policy", Path: "../30-architecture/contract-promotion-policy.md"}},
		Workflow:         ".github/workflows/architecture-docs.yml",
		EvidenceArtifact: "context-map-evidence",
		Loop: evidenceLoop{
			Observation:   "observation",
			Hypothesis:    "hypothesis",
			Execute:       "execute",
			Evaluate:      "evaluate",
			Retrospective: "retrospective",
		},
	}
}
