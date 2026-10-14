package main

func minimalManifest() manifest {
	return manifest{
		SchemaVersion: schemaVersion,
		ID:            "architecture",
		RiidoTask:     "RIID-TEST",
		Summary:       "contracts architecture",
		GeneratedDocs: generatedDocs{
			ModuleDecomposition: "docs/30-architecture/module-decomposition.md",
			IntegrationMatrix:   "docs/30-architecture/integration-matrix.md",
		},
		RequiredDocs:  []string{"docs/20-domain/context-map.md"},
		ContextMapDoc: "docs/20-domain/context-map.md",
		PackageRules: packageRules{
			Dependency:  "stdlib only",
			Forbidden:   "no runtime imports",
			TagBoundary: "tag boundary",
		},
		Packages: []packageEntry{{Path: "task", Display: "`task`", Role: "states", MustNotOwn: "runtime"}},
		ContractShape: contractShape{
			Allowed:   []string{"DTO structs"},
			Forbidden: []string{"network listeners"},
		},
		PublicGates:     []publicGate{{Surface: "module", Verification: "go test ./...", ExternalDependency: "none"}},
		DownstreamGates: []downstreamGate{{Consumer: "daemon", ExpectedGate: "imports tagged module"}},
		LocalCommands:   []string{"go test ./..."},
		StaleScanPaths:  []string{"docs/20-domain/context-map.md"},
		StaleRuntimeWords: []string{
			"AWS_SECRET",
		},
		Loop: evidenceLoop{
			Observation:   "observation",
			Hypothesis:    "hypothesis",
			Execute:       "execute",
			Evaluate:      "evaluate",
			Retrospective: "retrospective",
		},
	}
}
