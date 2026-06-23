package apicontract

import "testing"

func TestAIAgentTaskThreadHistoryV3Contract(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	path := "/v3/client/workspaces/{workspace_id}/ai-agent/tasks/{task_id}/threads"
	threads := fixture.openAPI.Paths[path]["get"]
	if threads.OperationID != "listAIAgentTaskThreadHistoryV3" ||
		threads.RiidoClient == nil ||
		threads.RiidoClient.GeneratedPath != "v3.aiAgent.tasks.threads" {
		t.Fatalf("v3 thread history operation = %#v", threads)
	}
	if threads.RiidoRBAC != "task_thread_history_collection.v1" {
		t.Fatalf("v3 thread history rbac = %q", threads.RiidoRBAC)
	}
}

func TestAIAgentTaskThreadHistoryV3Schemas(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	history := openAPISchemaProperties(t, fixture.openAPI, "AIAgentTaskThreadHistoryCollectionResponse")
	verifyHistoryThreads(t, history)
	verifyHistorySnapshots(t, history)
	message := openAPISchemaProperties(t, fixture.openAPI, "AIAgentTaskThreadHistoryMessage")
	if message["role"].(map[string]any)["$ref"] != "#/components/schemas/AIAgentTaskThreadMessageRole" {
		t.Fatalf("history message role schema = %#v", message["role"])
	}
}

func verifyHistoryThreads(t *testing.T, history map[string]any) {
	t.Helper()
	threads, ok := history["threads"].(map[string]any)
	if !ok {
		t.Fatalf("history threads schema = %#v", history["threads"])
	}
	items, ok := threads["items"].(map[string]any)
	if !ok || items["$ref"] != "#/components/schemas/AIAgentTaskThreadHistoryRecord" {
		t.Fatalf("history threads items = %#v", threads["items"])
	}
}

func verifyHistorySnapshots(t *testing.T, history map[string]any) {
	t.Helper()
	snapshots, ok := history["agent_snapshots"].(map[string]any)
	if !ok {
		t.Fatalf("history snapshots schema = %#v", history["agent_snapshots"])
	}
	additional, ok := snapshots["additionalProperties"].(map[string]any)
	if !ok || additional["$ref"] != "#/components/schemas/AIAgentTaskThreadAgentSnapshot" {
		t.Fatalf("history snapshot map schema = %#v", snapshots["additionalProperties"])
	}
}
