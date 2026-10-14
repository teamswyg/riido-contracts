package assignment

import "testing"

func assertAssignmentPayloadFields(t *testing.T, contract executableContract) {
	t.Helper()
	if len(contract.AssignmentPayloadFields) != 3 {
		t.Fatalf("assignment payload fields drifted: %#v", contract.AssignmentPayloadFields)
	}
	fields := map[string]contractAssignmentPayloadField{}
	for _, field := range contract.AssignmentPayloadFields {
		fields[field.Name] = field
	}
	assertAgentInstructionField(t, fields["agent_instruction"])
	assertExperimentalRuntimeField(t, fields["allow_experimental_runtime"])
	assertModelIDField(t, fields["model_id"])
}

func assertAgentInstructionField(t *testing.T, field contractAssignmentPayloadField) {
	t.Helper()
	if field.Name != "agent_instruction" ||
		field.Source != "agent.instruction" ||
		field.MaxLength != 1000 ||
		field.Required ||
		field.Snapshot != "assignment-created" {
		t.Fatalf("agent instruction assignment payload field drifted: %#v", field)
	}
	if field.Consumer != "riido-daemon provider-specific runtime instruction placement" {
		t.Fatalf("agent instruction consumer drifted: %q", field.Consumer)
	}
}
