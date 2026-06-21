package main

import (
	"path/filepath"
	"testing"
)

func TestManifestLoopStatusAcceptsLoopSource(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "source.riido.json"), fixtureManifest())
	mustWrite(t, filepath.Join(root, "target.riido.json"), `{"schema_version":"fixture.v1","loop_source":"source.riido.json"}`)
	if got := manifestLoopStatus(root, filepath.Join(root, "target.riido.json")); got != "delegated" {
		t.Fatalf("status = %q", got)
	}
}
