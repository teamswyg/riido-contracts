package apicontract

import (
	"strings"
	"testing"
)

func TestAIAgentClientDesktopPresenceAuthorizationBoundaries(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	var presence, devices IROperation
	for _, op := range f.ir.Operations {
		switch op.OperationID {
		case "getAIAgentDesktopPresenceV2":
			presence = op
		case "listAIAgentDeviceRuntimesV2":
			devices = op
		}
	}
	if presence.Resource != "device" || presence.Action != "read" || presence.Resource != devices.Resource || presence.Action != devices.Action {
		t.Fatal("presence must preserve the existing device read authorization boundary")
	}
	if presence.Auth.Scheme != devices.Auth.Scheme || presence.Auth.Header != "X-Riido-AI-Agent-Token" || presence.Auth.Header != devices.Auth.Header {
		t.Fatal("presence must not introduce an anonymous or different-account identity source")
	}
}

func TestAIAgentClientDesktopPresenceHistoricalAndReadBoundaries(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	var rules string
	for _, policy := range f.ir.Policies {
		if policy.PolicyID == "desktop_presence.v1" {
			rules = strings.Join(policy.Rules, "\n")
		}
	}
	for _, boundary := range []string{
		"another principal's records never count",
		"daemon heartbeats do not advance it",
		"not an active authorization grant",
		"may remain after credential revocation or membership removal",
		"only its observed committed records",
		"integrity failure is an error",
		"does not guarantee cross-replica visibility or history after restart",
		"omitted last_seen_at is not a global proof",
		"first request's network latency may be missed",
		"same page is not guaranteed to be detected",
	} {
		if !strings.Contains(rules, boundary) {
			t.Errorf("presence policy omits a security or freshness boundary: %q", boundary)
		}
	}
}
