package assignment

import (
	"strings"
	"testing"
)

func TestReferenceFolders(t *testing.T) {
	r := &ReferenceFolders{DeviceID: "device", RuntimeID: "runtime", Paths: []string{"/자료/with spaces", "/자료/with spaces", `C:\Users\Test\자료`}}
	got, err := r.Normalize()
	if err != nil || len(got.Paths) != 2 {
		t.Fatalf("normalize: %v %v", got, err)
	}
	r.Paths[0] = "/changed"
	if got.Paths[0] != "/자료/with spaces" {
		t.Fatal("input alias")
	}
	clone := got.Clone()
	clone.Paths[0] = "/changed"
	if got.Paths[0] != "/자료/with spaces" {
		t.Fatal("clone alias")
	}
	if !got.Matches("device", "runtime") || got.Matches("other", "runtime") {
		t.Fatal("binding")
	}
	if !strings.Contains(got.Instruction(), "/자료/with spaces") {
		t.Fatal("instruction")
	}
	for _, path := range []string{"", "relative/path", "~/folder", "/bad\x00path"} {
		_, err := (&ReferenceFolders{DeviceID: "device", RuntimeID: "runtime", Paths: []string{path}}).Normalize()
		if err == nil {
			t.Fatalf("accepted %q", path)
		}
	}
}
