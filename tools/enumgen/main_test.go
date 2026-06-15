package main

import (
	"bytes"
	"io"
	"path/filepath"
	"testing"
)

func TestGeneratedFilesFromCommonLispSource(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatal(err)
	}
	doc, err := loadDocument(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		t.Fatal(err)
	}
	if got := len(doc.Enums); got != 4 {
		t.Fatalf("enum count = %d, want 4", got)
	}
	if got := len(doc.Transitions); got != 2 {
		t.Fatalf("transition count = %d, want 2", got)
	}
	files, err := generatedFiles(doc)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{
		"task/task_state_enum_gen.go",
		"task/task_transition_code_enum_gen.go",
		"ir/event_type_enum_gen.go",
		"assignment/assignment_state_enum_gen.go",
		"assignment/assignment_transition_code_enum_gen.go",
		"assignment/poll_action_enum_gen.go",
	} {
		if len(files[path]) == 0 {
			t.Fatalf("missing generated file %s", path)
		}
	}
	if !bytes.Contains(files["task/task_state_enum_gen.go"], []byte("type TaskStateCode uint16")) {
		t.Fatal("task state generated file missing iota-backed code type")
	}
	if !bytes.Contains(files["task/task_state_enum_gen.go"], []byte("type TaskStateString string")) {
		t.Fatal("task state generated file missing named string type")
	}
}

func TestVerifyGeneratedFiles(t *testing.T) {
	if err := run([]string{"verify"}, io.Discard); err != nil {
		t.Fatal(err)
	}
}
