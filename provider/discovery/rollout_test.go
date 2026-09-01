package discovery

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestRolloutPreservesDaemonV1AndBoundsTelemetry(t *testing.T) {
	raw, err := os.ReadFile("rollout.riido.json")
	if err != nil {
		t.Fatal(err)
	}
	var manifest any
	if err := json.Unmarshal(raw, &manifest); err != nil {
		t.Fatal(err)
	}
	for _, required := range []string{
		"/v1/daemon/agent-bindings", "immutable",
		"/v2/daemon/agent-bindings", "retries v1", "optional signed",
		"riido-client-v2", "standalone provider-discovery endpoint",
		"forbidden_dimensions", "device_id", "redacted diagnostic_message",
		"unchanged five-second",
	} {
		if !strings.Contains(string(raw), required) {
			t.Fatalf("rollout manifest must contain %q", required)
		}
	}
}
