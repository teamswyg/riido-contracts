package main

func minimalManifest() manifest {
	return manifest{
		SchemaVersion: schemaVersion,
		ID:            "migration",
		RiidoTask:     "RIID-TEST",
		GeneratedDoc:  "docs/migration/contracts.md",
		Goal:          "goal",
		PromotionRule: promotionRule{Conditions: []string{"condition"}, Fallback: "fallback"},
		CandidateContracts: []candidateContract{{
			Candidate: "IR", Source: "private", Decision: "promote",
		}},
		RepositoryBoundaries: repositoryBoundaries{
			MayContain: []string{"DTOs"}, MustNotContain: []string{"runtime code"},
		},
		Versioning:     versioning{Intro: "versioning", Axes: []versionAxis{{Axis: "IR", OwnerBeforeSplit: "ir", ContractHandling: "tag"}}},
		MigrationOrder: []string{"add docs"},
		MigrationSlices: []migrationSlice{{
			Title: "RIID-TEST", Intro: []string{"intro"}, Does: []string{"work"}, DoesNot: "This slice does not deploy.",
		}},
		ValidationGates: validationGates{RequiredCommands: []string{"go test ./..."}, FixtureChecks: []string{"drift"}},
		MigrationWorkMap: []workMapEntry{{
			Area: "Contracts", RiidoTask: "RIID-TEST", TargetRepository: "riido-contracts",
		}},
		Workflow:         ".github/workflows/architecture-docs.yml",
		EvidenceArtifact: "migration-ledger-evidence",
		Loop: evidenceLoop{
			Observation: "observation", Hypothesis: "hypothesis", Execute: "execute",
			Evaluate: "evaluate", Retrospective: "retrospective",
		},
	}
}
