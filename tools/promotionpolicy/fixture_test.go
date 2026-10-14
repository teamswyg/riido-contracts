package main

func minimalManifest() manifest {
	return manifest{
		SchemaVersion: schemaVersion, ID: "promotion", RiidoTask: "RIID-TEST",
		Workflow:         ".github/workflows/architecture-docs.yml",
		EvidenceArtifact: "architecture-docs-evidence",
		Loop: evidenceLoop{
			Observation:   "observe",
			Hypothesis:    "hypothesis",
			Execute:       "execute",
			Evaluate:      "evaluate",
			Retrospective: "retro",
		},
		GeneratedDoc: "docs/30-architecture/contract-promotion-policy.md",
		Summary:      "policy summary",
		PromotionConditions: []string{
			"two repos agree",
			"no runtime imports",
			"versioned independently",
			"docs/migration/contracts.md updated",
			"downstream import gate",
		},
		SingleRuntimeRule: "single runtime stays local",
		SchemaVersionAxes: []versionAxis{{Axis: "task.FSMSchemaVersion", Rule: "changes with FSM"}},
		ModuleTagRule:     "module tag rule",
		RuntimeTagModel: []runtimeTag{
			{Pattern: "v0.0.x", Meaning: "bootstrap"},
			{Pattern: "v0.y.z", Meaning: "pre-1.0"},
			{Pattern: "v1.0.0", Meaning: "stable"},
		},
		RuntimeTagRule:       "runtime tags are separate",
		BreakingChangeRules:  []string{"remove field"},
		AdditiveChangeRule:   "additive is allowed",
		DownstreamImportRule: "downstream import required",
	}
}
