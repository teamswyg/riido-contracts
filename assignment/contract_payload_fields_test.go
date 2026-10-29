package assignment

import "testing"

func assertAssignmentPayloadFields(t *testing.T, contract executableContract) {
	t.Helper()
	if len(contract.AssignmentPayloadFields) != 5 {
		t.Fatalf("assignment payload fields drifted: %#v", contract.AssignmentPayloadFields)
	}
	fields := map[string]contractAssignmentPayloadField{}
	for _, field := range contract.AssignmentPayloadFields {
		fields[field.Name] = field
	}
	reference := fields["reference_folders"]
	if reference.Source != "agent.reference_folders" || reference.Required || reference.Snapshot != "assignment-created" {
		t.Fatalf("reference folders assignment payload drifted: %#v", reference)
	}
	assertAgentInstructionField(t, fields["agent_instruction"])
	assertExperimentalRuntimeField(t, fields["allow_experimental_runtime"])
	assertModelIDField(t, fields["model_id"])
	field := fields["recovery_policy"]
	if field.Name != "recovery_policy" || field.Required || field.MaxLength != 2048 || field.Snapshot != "assignment-created" || field.Source != "authorized assignment request; same owner/device verified fallback inventory" || field.Consumer != "riido-daemon finite pre-dispatch retry/circuit and explicitly authorized provider/model fallback" {
		t.Fatalf("recovery policy assignment payload field drifted: %#v", field)
	}
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
