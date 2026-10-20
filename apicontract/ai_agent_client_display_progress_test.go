package apicontract

import (
	"regexp"
	"testing"
)

func TestAIAgentClientDisplayProgressIsOptionalReadOnlyCopy(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	for _, name := range []string{"AgentThreadProgressLine", "AIAgentTaskThreadHistoryMessage"} {
		t.Run(name, func(t *testing.T) {
			properties := openAPISchemaProperties(t, fixture.openAPI, name)
			field, ok := properties["display_message"].(map[string]any)
			if !ok || field["type"] != "string" {
				t.Fatalf("%s must expose optional display_message string: %#v", name, field)
			}
			if field["maxLength"] != 180 {
				t.Fatalf("%s display_message must limit Unicode length to 180", name)
			}
			pattern, ok := field["pattern"].(string)
			if !ok || pattern == "" {
				t.Fatal("display_message must declare a single-line pattern")
			}
			valid := regexp.MustCompile(pattern)
			if !valid.MatchString("답글 전송 권한을 확인하고 있어요.") || valid.MatchString("one\ntwo") || valid.MatchString("one\rtwo") {
				t.Fatalf("display_message pattern must accept one line and reject CR/LF: %q", pattern)
			}
			schema := fixture.openAPI.Components.Schemas[name]
			for _, property := range openAPISchemaRequired(t, schema, name) {
				if property == "display_message" {
					t.Fatal("display_message must remain optional for older servers")
				}
			}
			for _, internal := range []string{"message_code", "message_args"} {
				if _, exists := properties[internal]; exists {
					t.Fatalf("%s must not expose internal %s", name, internal)
				}
			}
		})
	}
}
