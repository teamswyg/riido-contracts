package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyOnboarding(t *testing.T) {
	t.Helper()
	fixtureList := f.openAPI.Paths["/v1/client/ai-agent/onboarding/fixtures"]["get"]
	if fixtureList.OperationID != "listAIAgentOnboardingFixtures" ||
		fixtureList.RiidoClient == nil ||
		fixtureList.RiidoClient.CacheTag != "aiAgent.onboarding.fixtures" ||
		fixtureList.RiidoClient.GeneratedPath != "aiAgent.onboarding.fixtures" ||
		fixtureList.RiidoRBAC != "agent_onboarding_fixtures.v1" {
		t.Fatalf("fixture list operation = %#v", fixtureList)
	}
	f.verifyOnboardingCreate(t)
	f.verifyOnboardingFixtureSchema(t)
}

func (f aiAgentClientContractFixture) verifyOnboardingCreate(t *testing.T) {
	t.Helper()
	fixtureCreate := f.openAPI.Paths["/v1/client/ai-agent/onboarding/fixtures/{fixture_id}/agents"]["post"]
	if fixtureCreate.OperationID != "createAIAgentFromOnboardingFixture" ||
		fixtureCreate.RequestBody == nil ||
		fixtureCreate.RiidoClient == nil ||
		fixtureCreate.RiidoClient.GeneratedPath != "aiAgent.onboarding.fixtures.createAgent" ||
		!contains(fixtureCreate.RiidoClient.Invalidates, "aiAgent.bootstrap") {
		t.Fatalf("fixture create operation = %#v", fixtureCreate)
	}
	if len(fixtureCreate.Parameters) != 1 || fixtureCreate.Parameters[0].Name != "fixture_id" {
		t.Fatalf("fixture create parameters = %#v", fixtureCreate.Parameters)
	}
}

func (f aiAgentClientContractFixture) verifyOnboardingFixtureSchema(t *testing.T) {
	t.Helper()
	fixtureProps := openAPISchemaProperties(t, f.openAPI, "AgentOnboardingFixture")
	fixtureID, ok := fixtureProps["fixture_id"].(map[string]any)
	if !ok || fixtureID["description"] == "" {
		t.Fatalf("fixture_id description missing: %#v", fixtureProps["fixture_id"])
	}
	if _, ok := fixtureProps["tmp_color"].(map[string]any); !ok {
		t.Fatalf("fixture tmp_color schema missing: %#v", fixtureProps)
	}
}
