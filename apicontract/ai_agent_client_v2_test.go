package apicontract

import (
	"strings"
	"testing"
)

func (f aiAgentClientContractFixture) verifyV2WorkspaceContracts(t *testing.T) {
	t.Helper()
	f.verifyV2AgentCreate(t)
	verifyAgentClientRecordV2(t, f.openAPI)
	v2Bootstrap := f.openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/bootstrap"]["get"]
	if v2Bootstrap.RiidoClient == nil || v2Bootstrap.RiidoClient.GeneratedPath != "v2.aiAgent.bootstrap" || v2Bootstrap.RiidoClient.CacheTag != "v2.aiAgent.bootstrap" {
		t.Fatalf("v2 bootstrap client metadata = %#v", v2Bootstrap.RiidoClient)
	}
	if !strings.Contains(v2Bootstrap.Summary, "v2.aiAgent.bootstrap.agents[]") {
		t.Fatalf("v2 bootstrap summary must name agents[] source for generated comments: %q", v2Bootstrap.Summary)
	}
	f.verifyAssignedProfiles(t)
}

func (f aiAgentClientContractFixture) verifyV2AgentCreate(t *testing.T) {
	t.Helper()
	agentCreateV2 := f.openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/agents"]["post"]
	if agentCreateV2.OperationID != "createAIAgentV2" || agentCreateV2.RequestBody == nil || agentCreateV2.RiidoClient == nil || agentCreateV2.RiidoClient.GeneratedPath != "v2.aiAgent.agents.create" || !contains(agentCreateV2.RiidoClient.Invalidates, "v2.aiAgent.bootstrap") {
		t.Fatalf("v2 agent create operation = %#v", agentCreateV2)
	}
	if len(agentCreateV2.Parameters) != 1 || agentCreateV2.Parameters[0].Name != "workspace_id" {
		t.Fatalf("v2 agent create parameters = %#v", agentCreateV2.Parameters)
	}
}
