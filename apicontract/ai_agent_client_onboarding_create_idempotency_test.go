package apicontract

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestOnboardingCreateHasSeparateOptionalRequestIDSchema(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	base := openAPISchemaProperties(t, f.openAPI, "CreateAgentConfigurationRequest")
	onboarding := openAPISchemaProperties(t, f.openAPI, "CreateOnboardingAgentRequest")
	if len(base) != 7 || len(onboarding) != 8 {
		t.Fatalf("direct/fixture create property counts = %d/%d", len(base), len(onboarding))
	}
	for name, property := range base {
		if !reflect.DeepEqual(property, onboarding[name]) {
			t.Fatalf("fixture create changed ordinary configuration property %s", name)
		}
	}
	required := openAPISchemaRequired(t, f.openAPI.Components.Schemas["CreateOnboardingAgentRequest"], "CreateOnboardingAgentRequest")
	if contains(required, "request_id") || !reflect.DeepEqual(required,
		openAPISchemaRequired(t, f.openAPI.Components.Schemas["CreateAgentConfigurationRequest"], "CreateAgentConfigurationRequest")) {
		t.Fatalf("fixture create must preserve required configuration and keep request_id optional: %v", required)
	}
	for path, operations := range f.openAPI.Paths {
		op, ok := operations["post"]
		if !ok || op.OperationID != "createAIAgent" && op.OperationID != "createAIAgentV2" &&
			op.OperationID != "createAIAgentFromOnboardingFixture" && op.OperationID != "createAIAgentFromOnboardingFixtureV2" {
			continue
		}
		want := "#/components/schemas/CreateAgentConfigurationRequest"
		if strings.Contains(path, "/onboarding/fixtures/") {
			want = "#/components/schemas/CreateOnboardingAgentRequest"
		}
		if op.RequestBody == nil || op.RequestBody.Content["application/json"].Schema["$ref"] != want {
			t.Fatalf("%s request body must use %s", path, want)
		}
	}
}

func TestOnboardingCreateRequestIDIsBoundedASCII(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	p := openAPISchemaProperties(t, f.openAPI, "CreateOnboardingAgentRequest")["request_id"].(map[string]any)
	if p["minLength"] != 1 || p["maxLength"] != 128 {
		t.Fatalf("request_id length bounds = %#v", p)
	}
	pattern, ok := p["pattern"].(string)
	if !ok {
		t.Fatal("request_id pattern missing")
	}
	r := regexp.MustCompile(pattern)
	for _, v := range []string{"a", "onboard.A_1:retry-2", strings.Repeat("a", 128)} {
		if !r.MatchString(v) {
			t.Fatalf("request_id should accept %q", v)
		}
	}
	for _, v := range []string{"", "has space", "한글", "slash/value", "line\n", strings.Repeat("a", 129)} {
		if r.MatchString(v) {
			t.Fatalf("request_id should reject %q", v)
		}
	}
}

func TestOnboardingCreatePolicyDefinesDurableReplayBoundary(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	var rules string
	for _, policy := range dsl.Policies {
		if policy.PolicyID == "agent_onboarding_fixtures.v1" {
			rules = strings.Join(policy.Rules, "\n")
		}
	}
	for _, marker := range []string{"CreateOnboardingAgentRequest", "authorized principal", "workspace", "operation", "request_id", "201", "409", "503", "24 hours", "deleted", "hashed", "completion", "without request_id"} {
		if !strings.Contains(rules, marker) {
			t.Fatalf("fixture create replay policy missing %q", marker)
		}
	}
}
