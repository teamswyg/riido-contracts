package capability

import "testing"

func TestFingerprintDeterministic(t *testing.T) {
	in := fingerprintFixture()
	fp1, err := ComputeCapabilityFingerprint(in)
	if err != nil {
		t.Fatalf("compute 1: %v", err)
	}
	fp2, err := ComputeCapabilityFingerprint(in)
	if err != nil {
		t.Fatalf("compute 2: %v", err)
	}
	if fp1 != fp2 {
		t.Fatalf("fingerprint not deterministic: %s vs %s", fp1, fp2)
	}
	if len(fp1) != 64 {
		t.Fatalf("expected 64-char SHA-256 hex, got %d", len(fp1))
	}
}

func TestFingerprintMapOrderIndependent(t *testing.T) {
	a := CapabilityFingerprintInput{
		ImportantSurfaceFlags: map[string]any{"zzz": true, "aaa": false, "mmm": 42},
	}
	b := CapabilityFingerprintInput{
		ImportantSurfaceFlags: map[string]any{"aaa": false, "mmm": 42, "zzz": true},
	}
	fa, err := ComputeCapabilityFingerprint(a)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := ComputeCapabilityFingerprint(b)
	if err != nil {
		t.Fatal(err)
	}
	if fa != fb {
		t.Fatalf("fingerprint depends on map iteration order: %s vs %s", fa, fb)
	}
}
