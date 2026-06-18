package apicontract

import (
	"strings"
	"testing"
)

func (f aiAgentClientContractFixture) verifyAgentConfigurationRequests(t *testing.T) {
	t.Helper()
	verifyAgentConfigurationCreateRequest(t, f.openAPI)
	verifyAgentConfigurationUpdateRequest(t, f.openAPI)
}

func verifyAgentConfigurationCreateRequest(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	createRequest := openAPI.Components.Schemas["CreateAgentConfigurationRequest"]
	verifyAgentConfigurationDescription(t, "CreateAgentConfigurationRequest", createRequest)
	required := openAPISchemaRequired(t, createRequest, "CreateAgentConfigurationRequest")
	if !contains(required, "name") || !contains(required, "visibility") || !contains(required, "runtime_id") || contains(required, "model_id") {
		t.Fatalf("CreateAgentConfigurationRequest required = %#v", createRequest["required"])
	}
	verifyAgentConfigurationProperties(t, "CreateAgentConfigurationRequest", openAPISchemaPropertiesFromSchema(t, createRequest, "CreateAgentConfigurationRequest"))
}

func verifyAgentConfigurationUpdateRequest(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	updateRequest := openAPI.Components.Schemas["UpdateAgentConfigurationRequest"]
	verifyAgentConfigurationDescription(t, "UpdateAgentConfigurationRequest", updateRequest)
	verifyAgentConfigurationProperties(t, "UpdateAgentConfigurationRequest", openAPISchemaPropertiesFromSchema(t, updateRequest, "UpdateAgentConfigurationRequest"))
}

func verifyAgentConfigurationDescription(t *testing.T, schemaName string, schema map[string]any) {
	t.Helper()
	description, ok := schema["description"].(string)
	if !ok || !strings.Contains(description, "프로필 사진") || !strings.Contains(description, "런타임") || !strings.Contains(description, "모델") || !strings.Contains(description, "지침") {
		t.Fatalf("%s description must explain Figma agent setting fields: %q", schemaName, description)
	}
}
