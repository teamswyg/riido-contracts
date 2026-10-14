package capability

import "testing"

func TestFingerprintChangesOnSurfaceFlag(t *testing.T) {
	base := CapabilityFingerprintInput{
		ProviderKind: "claude",
		ImportantSurfaceFlags: map[string]any{
			"ExposesUnsafePermissionBypass": false,
		},
	}
	fp1, _ := ComputeCapabilityFingerprint(base)
	base.ImportantSurfaceFlags = map[string]any{
		"ExposesUnsafePermissionBypass": true,
	}
	fp2, _ := ComputeCapabilityFingerprint(base)
	if fp1 == fp2 {
		t.Fatal("fingerprint must change when an important surface flag changes")
	}
}
