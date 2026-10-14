package main

import "fmt"

func verifyManifest(m manifest) error {
	if m.SchemaVersion != "riido-ai-agent-assigned-profile-map-evidence-manifest.v1" {
		return fmt.Errorf("unexpected manifest schema %s", m.SchemaVersion)
	}
	if m.ID == "" || m.GeneratedDoc == "" || m.PolicyID == "" {
		return fmt.Errorf("manifest identity is incomplete")
	}
	if m.DSLFixture == "" || m.IRFixture == "" || m.OpenAPIFixture == "" {
		return fmt.Errorf("manifest fixture paths are incomplete")
	}
	if m.RequiredOperation.OperationID == "" || len(m.RequiredPolicyRules) == 0 {
		return fmt.Errorf("manifest must define operation and policy expectations")
	}
	return verifyLoop(m.Loop)
}
