package deviceprincipal

import "testing"

func TestDevicePrincipalCadence(t *testing.T) {
	if RuntimeSnapshotIntervalSeconds != 5 {
		t.Fatalf("snapshot interval = %d", RuntimeSnapshotIntervalSeconds)
	}
	if RuntimeStaleAfterSeconds != 20 {
		t.Fatalf("stale after = %d", RuntimeStaleAfterSeconds)
	}
}

func TestCredentialHeaderBoundary(t *testing.T) {
	for _, daemonHeader := range DaemonCredentialHeaders() {
		for _, clientHeader := range ClientCredentialHeaders() {
			if daemonHeader == clientHeader {
				t.Fatalf("daemon header overlaps client header: %s", daemonHeader)
			}
		}
	}
}

func TestExcludedFallbackVocabulary(t *testing.T) {
	want := map[string]bool{"team_id": true, "teamId": true, "X-Workspace-Api-Key": true}
	for _, fallback := range ExcludedFallbacks() {
		delete(want, fallback)
	}
	for missing := range want {
		t.Fatalf("missing excluded fallback %q", missing)
	}
}
