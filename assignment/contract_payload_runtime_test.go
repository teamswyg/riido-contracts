package assignment

import "testing"

func assertExperimentalRuntimeField(t *testing.T, field contractAssignmentPayloadField) {
	t.Helper()
	if field.Name != "allow_experimental_runtime" ||
		field.Source != "agent.runtime.requires_experimental_opt_in" ||
		field.MaxLength != 0 ||
		field.Required ||
		field.Snapshot != "assignment-created" {
		t.Fatalf("experimental runtime opt-in assignment payload field drifted: %#v", field)
	}
	if field.Consumer != "riido-daemon runtime scheduling experimental opt-in gate" {
		t.Fatalf("experimental runtime opt-in consumer drifted: %q", field.Consumer)
	}
}

func assertModelIDField(t *testing.T, field contractAssignmentPayloadField) {
	t.Helper()
	if field.Name != "model_id" ||
		field.Source != "agent.model_id" ||
		field.MaxLength != 128 ||
		field.Required ||
		field.Snapshot != "assignment-created" {
		t.Fatalf("model_id assignment payload field drifted: %#v", field)
	}
	if field.Consumer != "riido-daemon provider runtime model selection" {
		t.Fatalf("model_id consumer drifted: %q", field.Consumer)
	}
}
