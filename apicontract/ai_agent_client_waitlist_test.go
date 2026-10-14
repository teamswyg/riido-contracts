package apicontract

import (
	"strings"
	"testing"
)

func TestAIAgentClientOpenAPIDoesNotExposeRuntimeWaitlistMutation(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	for path, methods := range fixture.openAPI.Paths {
		for method, operation := range methods {
			haystack := strings.ToLower(path + " " + method + " " + operation.OperationID + " " + operation.Summary)
			for _, forbidden := range []string{"waitlist", "marketing", "consent"} {
				if strings.Contains(haystack, forbidden) {
					t.Fatalf("AI Agent OpenAPI exposed %q in %s %s (%s)", forbidden, method, path, operation.OperationID)
				}
			}
		}
	}
}
