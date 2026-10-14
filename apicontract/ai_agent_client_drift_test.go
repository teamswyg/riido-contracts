package apicontract

import "testing"

func TestAIAgentClientGeneratedFixturesDoNotDrift(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	assertFixture(t, "fixtures/control-plane-ai-agent-client.ir.riido.json", fixture.ir)
	assertFixture(t, "fixtures/control-plane-ai-agent-client.openapi.json", fixture.openAPI)
}
