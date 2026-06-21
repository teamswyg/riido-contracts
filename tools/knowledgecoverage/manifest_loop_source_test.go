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

func TestManifestLoopStatusAcceptsExternalLoopSource(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "source.riido.json"), fixtureManifest())
	mustWrite(t, filepath.Join(root, "target.riido.json"), `{"schema_version":"fixture.v1"}`)
	sources := []manifestLoopSource{{ID: "fixture", LoopSource: "source.riido.json", Paths: []string{"target.riido.json"}}}
	if got := manifestLoopStatusWithSources(root, filepath.Join(root, "target.riido.json"), sources); got != "delegated" {
		t.Fatalf("status = %q", got)
	}
}

func TestManifestLoopStatusAcceptsExternalPrefix(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "source.riido.json"), fixtureManifest())
	mustWrite(t, filepath.Join(root, "owned", "target.riido.json"), `{"schema_version":"fixture.v1"}`)
	sources := []manifestLoopSource{{ID: "fixture", LoopSource: "source.riido.json", PathPrefixes: []string{"owned/"}}}
	if got := manifestLoopStatusWithSources(root, filepath.Join(root, "owned", "target.riido.json"), sources); got != "delegated" {
		t.Fatalf("status = %q", got)
	}
}
