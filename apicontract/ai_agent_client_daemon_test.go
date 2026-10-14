package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyDaemonContracts(t *testing.T) {
	t.Helper()
	daemonList := f.openAPI.Paths["/v1/client/ai-agent/devices/{device_id}/daemons"]["get"]
	if daemonList.RiidoClient == nil || daemonList.RiidoClient.CacheTag != "aiAgent.devices.daemons" {
		t.Fatalf("daemon list client metadata = %#v", daemonList.RiidoClient)
	}
	if len(daemonList.Parameters) != 1 || daemonList.Parameters[0].Name != "device_id" {
		t.Fatalf("daemon list parameters = %#v", daemonList.Parameters)
	}
	daemonListV2 := f.openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/devices/{device_id}/daemons"]["get"]
	if daemonListV2.RiidoClient == nil || daemonListV2.RiidoClient.CacheTag != "v2.aiAgent.devices.daemons" {
		t.Fatalf("v2 daemon list client metadata = %#v", daemonListV2.RiidoClient)
	}
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
