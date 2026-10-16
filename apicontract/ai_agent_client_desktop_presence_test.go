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
	if len(op.Security) != 1 {
		t.Fatalf("desktop presence must require the existing token security scheme")
	}
	if _, ok := op.Security[0]["riidoAIAgentToken"]; !ok {
		t.Fatalf("desktop presence token security scheme missing")
	}
	if op.RequestBody != nil || len(op.Parameters) != 1 || op.Parameters[0].Name != "workspace_id" || op.Parameters[0].In != "path" || !op.Parameters[0].Required {
		t.Fatalf("desktop presence must remain a workspace-authorized body-free read")
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
		property, ok := props[field].(map[string]any)
		if !ok {
			t.Fatalf("%s %s property missing", name, field)
		}
		if property["type"] != "string" || property["format"] != "date-time" {
			t.Fatalf("%s %s must remain a server date-time, not an identifier or flag", name, field)
		}
	}
	required := openAPISchemaRequired(t, openAPI.Components.Schemas[name], name)
	if len(required) != 1 || required[0] != "checked_at" {
		t.Fatalf("%s required = %#v", name, required)
	}
}
