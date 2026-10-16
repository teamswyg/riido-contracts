package apicontract

import "testing"

const desktopPresencePath = "/v2/client/workspaces/{workspace_id}/ai-agent/desktop-presence"

func TestAIAgentClientDesktopPresenceContract(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	verifyDesktopPresenceOperation(t, f.openAPI)
	verifyDesktopPresenceSchema(t, f.openAPI)
}

func verifyDesktopPresenceOperation(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	op := openAPI.Paths[desktopPresencePath]["get"]
	if op.OperationID != "getAIAgentDesktopPresenceV2" || op.RiidoKind != "query" {
		t.Fatalf("desktop presence operation = %#v", op)
	}
	if op.RiidoClient == nil || op.RiidoClient.GeneratedPath != "v2.aiAgent.desktopPresence" {
		t.Fatalf("desktop presence client metadata = %#v", op.RiidoClient)
	}
	if len(op.RiidoScopes) != 1 || op.RiidoScopes[0] != "ai-agent:device:read" {
		t.Fatalf("desktop presence scopes = %#v", op.RiidoScopes)
	}
	if op.RiidoRBAC != "desktop_presence.v1" {
		t.Fatalf("desktop presence rbac policy = %q", op.RiidoRBAC)
	}
}

func verifyDesktopPresenceSchema(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	name := "AIAgentDesktopPresenceResponse"
	props := openAPISchemaProperties(t, openAPI, name)
	if len(props) != 2 {
		t.Fatalf("%s must expose only last_seen_at and checked_at: %#v", name, props)
	}
	for _, field := range []string{"last_seen_at", "checked_at"} {
		if _, ok := props[field]; !ok {
			t.Fatalf("%s %s property missing", name, field)
		}
	}
	required := openAPISchemaRequired(t, openAPI.Components.Schemas[name], name)
	if len(required) != 1 || required[0] != "checked_at" {
		t.Fatalf("%s required = %#v", name, required)
	}
}
