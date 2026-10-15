package apicontract

import (
	"slices"
	"strings"
	"testing"
)

func TestAIAgentRequestIdentityBehaviorMatchesRequestSchema(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	for _, op := range dsl.Operations {
		if op.OperationID != "assignAIAgentTask" && op.OperationID != "assignAIAgentTaskV2" {
			continue
		}
		scenarios := map[string]bool{}
		for _, scenario := range op.Scenarios {
			if strings.Contains(scenario.Then, "body carries agent_id only") {
				t.Fatalf("%s still forbids optional request identity/instruction", op.OperationID)
			}
			scenarios[scenario.Name] = true
		}
		if !scenarios["same explicit request is idempotent"] || !scenarios["new explicit request preserves busy agent instruction"] {
			t.Fatalf("%s misses retry/new-request distinction", op.OperationID)
		}
	}
	f := loadAIAgentClientContractFixture(t)
	for _, name := range []string{"AIAgentTaskThreadCollectionResponse", "AIAgentTaskThreadHistoryCollectionResponse"} {
		required, _ := f.openAPI.Components.Schemas[name]["required"].([]string)
		if !slices.Contains(required, "assigned_agent_ids") {
			t.Fatalf("%s cannot distinguish absent membership from omitted data", name)
		}
	}
}
