package apicontract

import (
	"os"
	"strings"
	"testing"
)

func TestAIAgentClientContractRejectsAmbiguousFutureClientBootstrapWording(t *testing.T) {
	for _, path := range []string{
		"fixtures/control-plane-ai-agent-client.dsl.riido.json",
		"fixtures/control-plane-ai-agent-client.ir.riido.json",
		"fixtures/control-plane-ai-agent-client.openapi.json",
	} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		if strings.Contains(strings.ToLower(string(data)), "future client bootstrap") {
			t.Fatalf("%s contains ambiguous future-client wording; use subsequent aiAgent.bootstrap read wording", path)
		}
	}
}
