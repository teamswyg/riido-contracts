package apicontract

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/teamswyg/riido-contracts/assignment"
)

func TestExecutionContextVocabularySurvivesAPIProjection(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	want := map[string][]string{
		"workspace_kind":      assignment.ExecutionWorkspaceKinds(),
		"verification_status": assignment.ExecutionVerificationStatuses(),
		"diagnostic_code":     assignment.ExecutionDiagnosticCodes(),
		"next_action":         assignment.ExecutionNextActions(),
	}
	// Compare the public projection as JSON so this test also guards omitted
	// property constraints without depending on internal generator structures.
	body, err := json.Marshal(fixture.openAPI)
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Components struct {
			Schemas map[string]struct {
				Properties map[string]struct {
					Enum []string `json:"enum"`
				} `json:"properties"`
			} `json:"schemas"`
		} `json:"components"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		t.Fatal(err)
	}
	properties := doc.Components.Schemas["AIAgentExecutionContext"].Properties
	if len(properties) != len(want) {
		t.Fatal("execution context exposed unexpected fields or lost its schema")
	}
	for field, values := range want {
		if !reflect.DeepEqual(properties[field].Enum, values) {
			t.Fatalf("%s vocabulary drift: %v", field, properties[field].Enum)
		}
	}
}
