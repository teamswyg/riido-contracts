package discovery

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"testing"
	"time"

	providercatalog "github.com/teamswyg/riido-contracts/provider/catalog"
)

func TestPolicyAcceptsBoundedLocalResolutionHints(t *testing.T) {
	now := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	policy := validPolicy(now)
	if err := policy.ValidateAt(now); err != nil {
		t.Fatal(err)
	}
}

func TestPolicyRejectsAbsoluteOrEscapingPaths(t *testing.T) {
	now := time.Now().UTC()
	for _, unsafe := range []string{"/Applications/ChatGPT.app", `C:\\Program Files\\Cursor`, "../bin/codex"} {
		policy := validPolicy(now)
		policy.Rules[0].Candidates[0].RelativePath = unsafe
		if err := policy.Validate(); err == nil {
			t.Fatalf("expected %q to fail", unsafe)
		}
	}
}

func TestPolicyRejectsUnknownRootAndExcessLifetime(t *testing.T) {
	now := time.Now().UTC()
	policy := validPolicy(now)
	policy.Rules[0].Candidates[0].Root = "arbitrary-root"
	if err := policy.Validate(); err == nil {
		t.Fatal("expected unknown root to fail")
	}
	policy = validPolicy(now)
	policy.ExpiresAt = now.Add(MaxPolicyLifetime + time.Second)
	if err := policy.Validate(); err == nil {
		t.Fatal("expected excess lifetime to fail")
	}
}

func TestPolicyRejectsRootFromAnotherOS(t *testing.T) {
	policy := validPolicy(time.Now().UTC())
	policy.Rules[0].Candidates[0].Root = RootProgramFiles
	if err := policy.Validate(); err == nil {
		t.Fatal("expected Windows root on darwin rule to fail")
	}
}

func TestPolicyRejectsExecutableNameFromAnotherProvider(t *testing.T) {
	policy := validPolicy(time.Now().UTC())
	policy.Rules[0].Candidates[0].RelativePath = "ChatGPT.app/Contents/Resources/bash"
	if err := policy.Validate(); err == nil {
		t.Fatal("expected arbitrary executable name to fail")
	}
}

func TestSignedEnvelopeIsBounded(t *testing.T) {
	envelope := SignedEnvelope{
		SchemaVersion: EnvelopeSchemaVersion,
		KeyID:         "provider-policy-2026-09",
		Algorithm:     SignatureAlgorithm,
		Payload:       base64.StdEncoding.EncodeToString([]byte(`{"schema_version":"riido-provider-discovery-policy.v1"}`)),
		Signature:     base64.StdEncoding.EncodeToString(make([]byte, 64)),
	}
	if err := envelope.Validate(); err != nil {
		t.Fatal(err)
	}
	envelope.Algorithm = "arbitrary"
	if err := envelope.Validate(); err == nil {
		t.Fatal("expected unknown signature algorithm to fail")
	}
}

func TestSignedEnvelopeRoundTripAndTamperRejection(t *testing.T) {
	now := time.Now().UTC()
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	envelope, err := Sign(validPolicy(now), "provider-policy-test", privateKey)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := envelope.Verify(map[string]ed25519.PublicKey{"provider-policy-test": publicKey}, now); err != nil {
		t.Fatal(err)
	}
	envelope.Payload = base64.StdEncoding.EncodeToString([]byte(`{"schema_version":"riido-provider-discovery-policy.v1"}`))
	if _, err := envelope.Verify(map[string]ed25519.PublicKey{"provider-policy-test": publicKey}, now); err == nil {
		t.Fatal("expected tampered payload to fail")
	}
}

func validPolicy(now time.Time) Policy {
	return Policy{
		SchemaVersion: PolicySchemaVersion,
		Revision:      "2026-09-01.1",
		IssuedAt:      now,
		ExpiresAt:     now.Add(24 * time.Hour),
		Rules: []Rule{{
			Provider: providercatalog.KindCodex,
			OS:       "darwin",
			Arch:     "arm64",
			Candidates: []Candidate{{
				Root:         RootSystemApplications,
				RelativePath: "ChatGPT.app/Contents/Resources/codex",
			}},
		}},
	}
}
