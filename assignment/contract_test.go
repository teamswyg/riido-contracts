package assignment

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sort"
	"testing"
)

func TestAssignmentContractMatchesPackageSurface(t *testing.T) {
	contract := loadContract(t)
	if contract.SchemaVersion != ContractSchemaVersion {
		t.Fatalf("contract schema_version = %q, want %q", contract.SchemaVersion, ContractSchemaVersion)
	}
	if contract.ServiceSchemaVersion != SchemaVersion {
		t.Fatalf("service_schema_version = %q, want %q", contract.ServiceSchemaVersion, SchemaVersion)
	}

	remainingStates := map[AssignmentState]struct{}{}
	for _, state := range AllAssignmentStates() {
		remainingStates[state] = struct{}{}
	}
	contractStates := map[AssignmentState]contractState{}
	for _, state := range contract.AssignmentStates {
		value := AssignmentState(state.Value)
		if !value.Valid() {
			t.Fatalf("contract state %q is missing from package constants", state.Value)
		}
		if _, exists := contractStates[value]; exists {
			t.Fatalf("duplicate contract state %q", state.Value)
		}
		contractStates[value] = state
		delete(remainingStates, value)
		if got := IsTerminal(value); got != state.Terminal {
			t.Fatalf("IsTerminal(%q) = %v, want %v", value, got, state.Terminal)
		}
		if got := IsAgentActive(value); got != state.AgentActive {
			t.Fatalf("IsAgentActive(%q) = %v, want %v", value, got, state.AgentActive)
		}
	}
	if len(remainingStates) != 0 {
		t.Fatalf("package states missing from contract: %v", sortedAssignmentStateSet(remainingStates))
	}

	for _, from := range AllAssignmentStates() {
		allowed := map[AssignmentState]struct{}{}
		for _, transition := range contractStates[from].Transitions {
			allowed[AssignmentState(transition)] = struct{}{}
		}
		for _, to := range AllAssignmentStates() {
			_, inContract := allowed[to]
			want := from == to || inContract
			if got := CanTransition(from, to); got != want {
				t.Fatalf("CanTransition(%q,%q) = %v, want %v", from, to, got, want)
			}
		}
	}

	remainingActions := map[PollAction]struct{}{}
	for _, action := range AllPollActions() {
		remainingActions[action] = struct{}{}
	}
	for _, action := range contract.PollActions {
		value := PollAction(action.Value)
		if !value.Valid() {
			t.Fatalf("contract poll action %q is missing from package constants", action.Value)
		}
		delete(remainingActions, value)
	}
	if len(remainingActions) != 0 {
		t.Fatalf("package poll actions missing from contract: %v", sortedPollActionSet(remainingActions))
	}

	remainingTaskEvents := map[string]struct{}{}
	for _, event := range AllTaskEventTypes() {
		remainingTaskEvents[event] = struct{}{}
	}
	for _, event := range contract.TaskEvents {
		if _, ok := remainingTaskEvents[event.Value]; !ok {
			t.Fatalf("contract task event %q is missing from package constants", event.Value)
		}
		delete(remainingTaskEvents, event.Value)
	}
	if len(remainingTaskEvents) != 0 {
		t.Fatalf("package task events missing from contract: %v", sortedStringSet(remainingTaskEvents))
	}

	if len(contract.AssignmentPayloadFields) != 3 {
		t.Fatalf("assignment payload fields drifted: %#v", contract.AssignmentPayloadFields)
	}
	fields := map[string]contractAssignmentPayloadField{}
	for _, field := range contract.AssignmentPayloadFields {
		fields[field.Name] = field
	}
	instruction := fields["agent_instruction"]
	if instruction.Name != "agent_instruction" ||
		instruction.Source != "agent.instruction" ||
		instruction.MaxLength != 1000 ||
		instruction.Required ||
		instruction.Snapshot != "assignment-created" {
		t.Fatalf("agent instruction assignment payload field drifted: %#v", instruction)
	}
	if instruction.Consumer != "riido-daemon provider-specific runtime instruction placement" {
		t.Fatalf("agent instruction consumer drifted: %q", instruction.Consumer)
	}
	optIn := fields["allow_experimental_runtime"]
	if optIn.Name != "allow_experimental_runtime" ||
		optIn.Source != "agent.runtime.requires_experimental_opt_in" ||
		optIn.MaxLength != 0 ||
		optIn.Required ||
		optIn.Snapshot != "assignment-created" {
		t.Fatalf("experimental runtime opt-in assignment payload field drifted: %#v", optIn)
	}
	if optIn.Consumer != "riido-daemon runtime scheduling experimental opt-in gate" {
		t.Fatalf("experimental runtime opt-in consumer drifted: %q", optIn.Consumer)
	}
	modelID := fields["model_id"]
	if modelID.Name != "model_id" ||
		modelID.Source != "agent.model_id" ||
		modelID.MaxLength != 128 ||
		modelID.Required ||
		modelID.Snapshot != "assignment-created" {
		t.Fatalf("model_id assignment payload field drifted: %#v", modelID)
	}
	if modelID.Consumer != "riido-daemon provider runtime model selection" {
		t.Fatalf("model_id consumer drifted: %q", modelID.Consumer)
	}
}

func TestAssignmentTransitionBDDScenarios(t *testing.T) {
	cases := []struct {
		name       string
		from, to   AssignmentState
		want       bool
		acceptance string
	}{
		{
			name:       "daemon can lease queued work",
			from:       AssignmentQueued,
			to:         AssignmentLeased,
			want:       true,
			acceptance: "poll start can claim queued assignment",
		},
		{
			name:       "daemon reports active assignment as running",
			from:       AssignmentReady,
			to:         AssignmentRunning,
			want:       true,
			acceptance: "daemon event can move ready assignment to running",
		},
		{
			name:       "running completes after agent event",
			from:       AssignmentRunning,
			to:         AssignmentCompleted,
			want:       true,
			acceptance: "terminal success is reached only from running",
		},
		{
			name:       "terminal state cannot restart",
			from:       AssignmentCompleted,
			to:         AssignmentRunning,
			want:       false,
			acceptance: "completed assignment is terminal",
		},
		{
			name:       "queued cannot skip to completed",
			from:       AssignmentQueued,
			to:         AssignmentCompleted,
			want:       false,
			acceptance: "assignment must pass through daemon-active states before success",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanTransition(tc.from, tc.to); got != tc.want {
				t.Fatalf("%s: CanTransition(%q,%q) = %v, want %v", tc.acceptance, tc.from, tc.to, got, tc.want)
			}
		})
	}
}

func loadContract(t *testing.T) executableContract {
	t.Helper()
	data, err := os.ReadFile("assignment_contract.riido.json")
	if err != nil {
		t.Fatalf("read assignment contract: %v", err)
	}
	var contract executableContract
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&contract); err != nil {
		t.Fatalf("unmarshal assignment contract: %v", err)
	}
	var trailing struct{}
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		t.Fatal("assignment contract must contain exactly one JSON document")
	}
	return contract
}

type executableContract struct {
	SchemaVersion           string                           `json:"schema_version"`
	ServiceSchemaVersion    string                           `json:"service_schema_version"`
	AssignmentStates        []contractState                  `json:"assignment_states"`
	PollActions             []contractValue                  `json:"poll_actions"`
	TaskEvents              []contractValue                  `json:"task_events"`
	AssignmentPayloadFields []contractAssignmentPayloadField `json:"assignment_payload_fields"`
}

type contractState struct {
	Name        string   `json:"name"`
	Value       string   `json:"value"`
	AgentActive bool     `json:"agent_active"`
	Terminal    bool     `json:"terminal"`
	Transitions []string `json:"transitions"`
}

type contractValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type contractAssignmentPayloadField struct {
	Name      string `json:"name"`
	Source    string `json:"source"`
	MaxLength int    `json:"max_length"`
	Required  bool   `json:"required"`
	Snapshot  string `json:"snapshot"`
	Consumer  string `json:"consumer"`
}

func sortedAssignmentStateSet(values map[AssignmentState]struct{}) []AssignmentState {
	keys := make([]AssignmentState, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

func sortedPollActionSet(values map[PollAction]struct{}) []PollAction {
	keys := make([]PollAction, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

func sortedStringSet(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
