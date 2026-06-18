package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyDaemonContracts(t *testing.T) {
	t.Helper()
	daemonDetail := f.openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/daemon"]["get"]
	if daemonDetail.RiidoClient == nil || daemonDetail.RiidoClient.CacheTag != "aiAgent.agents.daemon" {
		t.Fatalf("daemon detail client metadata = %#v", daemonDetail.RiidoClient)
	}
	if len(daemonDetail.Parameters) != 1 || daemonDetail.Parameters[0].Name != "agent_id" {
		t.Fatalf("daemon detail parameters = %#v", daemonDetail.Parameters)
	}
	daemonStop := f.openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/daemon/stop"]["post"]
	if daemonStop.RiidoClient == nil || !contains(daemonStop.RiidoClient.Invalidates, "aiAgent.devices.runtimes") {
		t.Fatalf("daemon stop client metadata = %#v", daemonStop.RiidoClient)
	}
	editability := f.openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/editability"]["get"]
	if editability.RiidoRBAC != "agent_mutation_safety.v1" {
		t.Fatalf("editability rbac = %q", editability.RiidoRBAC)
	}
}
