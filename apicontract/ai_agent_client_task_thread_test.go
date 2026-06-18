package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyTaskThreadContracts(t *testing.T) {
	t.Helper()
	threadMessageCreate := f.openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/threads/{thread_id}/messages"]["post"]
	if threadMessageCreate.OperationID != "createAIAgentTaskThreadMessage" ||
		threadMessageCreate.RiidoClient == nil ||
		threadMessageCreate.RiidoClient.GeneratedPath != "aiAgent.tasks.threadMessages.create" {
		t.Fatalf("thread message create operation = %#v", threadMessageCreate)
	}
	f.verifyAssignableAgents(t)
	f.verifyThreadsCollection(t)
	verifyTaskThreadSchemas(t, f.openAPI)
}

func (f aiAgentClientContractFixture) verifyThreadsCollection(t *testing.T) {
	t.Helper()
	threads := f.openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/threads"]["get"]
	if threads.RiidoRBAC != "task_thread_cold_collection.v1" {
		t.Fatalf("task threads rbac = %q", threads.RiidoRBAC)
	}
	if len(threads.Parameters) != 1 || threads.Parameters[0].Name != "task_id" {
		t.Fatalf("task threads parameters = %#v", threads.Parameters)
	}
}
