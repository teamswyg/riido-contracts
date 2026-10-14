package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRunWritesEvidence(t *testing.T) {
	root := fixtureRepo(t, fixtureManifest())
	out := filepath.Join(root, "out", "questions.json")
	if err := run([]string{"-root", root, "-write-doc", "-check-doc", "-evidence-out", out}); err != nil {
		t.Fatalf("run: %v", err)
	}
	body, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read evidence: %v", err)
	}
	var got evidence
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode evidence: %v", err)
	}
	if got.Status != "advisory_open_questions" || got.OpenCount != 1 || got.ResolvedCount != 1 {
		t.Fatalf("unexpected evidence: %+v", got)
	}
}

func TestRejectsDuplicateQuestionID(t *testing.T) {
	root := fixtureRepo(t, duplicateFixtureManifest())
	if err := run([]string{"-root", root, "-write-doc"}); err == nil {
		t.Fatal("expected duplicate id failure")
	}
}
