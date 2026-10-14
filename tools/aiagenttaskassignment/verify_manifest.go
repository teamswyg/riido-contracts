package main

import "fmt"

func verifyManifest(m manifest) error {
	if m.SchemaVersion != "riido-ai-agent-task-assignment-evidence-manifest.v1" {
		return fmt.Errorf("unexpected manifest schema %s", m.SchemaVersion)
	}
	if m.ID == "" || m.Title == "" || m.GeneratedDoc == "" {
		return fmt.Errorf("manifest identity is incomplete")
	}
	if m.DSLFixture == "" || m.IRFixture == "" || m.OpenAPIFixture == "" {
		return fmt.Errorf("manifest fixture paths are incomplete")
	}
	if len(m.RequiredOperations) == 0 || len(m.RequiredSchemaFields) == 0 {
		return fmt.Errorf("manifest must define operation and schema expectations")
	}
	return verifyLoop(m.Loop)
}
