package capability

import "testing"

func TestFingerprintChangesOnPolicyBundle(t *testing.T) {
	base := CapabilityFingerprintInput{
		ProviderKind:        "claude",
		PolicyBundleVersion: "v1",
	}
	fp1, _ := ComputeCapabilityFingerprint(base)
	base.PolicyBundleVersion = "v2"
	fp2, _ := ComputeCapabilityFingerprint(base)
	if fp1 == fp2 {
		t.Fatal("fingerprint must change when PolicyBundleVersion changes")
	}
}
