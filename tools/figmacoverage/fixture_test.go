package main

func minimalManifest() manifest {
	return manifest{
		SchemaVersion:    schemaVersion,
		ID:               "figma-test",
		RiidoTask:        "RIID-TEST",
		Workflow:         ".github/workflows/architecture-docs.yml",
		EvidenceArtifact: "architecture-docs-evidence",
		Loop: evidenceLoop{
			Observation:   "observe",
			Hypothesis:    "hypothesis",
			Execute:       "execute",
			Evaluate:      "evaluate",
			Retrospective: "retro",
		},
		HumanDoc:     "docs/30-architecture/figma-ai-agent-coverage.md",
		StabilizedBy: []string{"teamswyg/riido-contracts#1"},
		Figma: figmaSource{
			FileKey: "file-key", FileName: "AI Agent", PageID: "1:1", PageName: "UI",
			InspectedAt: "2026-06-02", InspectionSource: "test",
		},
		InspectionMethod:      inspectionMethod{Authority: "test"},
		CoveragePolicy:        coveragePolicy{Summary: "summary", TopDown: "top", BottomUp: "bottom"},
		ExpectedPages:         []page{{NodeID: "1:1", Name: "UI", ChildCount: 1}},
		ExpectedTopLevelNodes: []node{{NodeID: "2:1", Name: "Section"}},
		Entries:               []coverageEntry{minimalEntry()},
		APIAnnotationContentPolicy: annotationContentPolicy{
			CategoryID: "700:0", CategoryLabel: "API Generated", LabelFormat: []string{"a", "b", "c"},
			Rule: "rule", LiveInspection: liveInspection{TotalAPIGeneratedAnnotations: 1},
		},
		APIGeneratedAnnotationInventory: []annotationInventory{minimalAnnotation()},
		VerifiedEvidenceNodes:           []node{{NodeID: "2:1", Name: "Section"}},
	}
}

func minimalEntry() coverageEntry {
	return coverageEntry{
		NodeID: "2:1", Name: "Section", CoverageStatus: "covered", EvidenceKind: "figma_section",
		SSOTDocs: []string{"docs/20-domain/ai-agent-policy.md"}, OwnerRepos: []string{"riido-contracts"},
		CoveredFacts: []string{"fact"}, DirectionLoop: directionLoop{TopDown: "top", BottomUp: "bottom"},
	}
}

func minimalAnnotation() annotationInventory {
	return annotationInventory{
		UIArea: "Area", CategoryID: "700:0", CategoryLabel: "API Generated",
		FigmaGeneratedPath: "riido.v2.aiAgent.tasks", CanonicalGeneratedPath: "aiAgent.tasks",
		OperationKind: "Query", Background: "background", AnnotationCount: 1,
		Sources: []annotationSource{{PageID: "1:1", TopLevelNodeID: "2:1", CoverageEntryNodeID: "2:1"}},
	}
}
