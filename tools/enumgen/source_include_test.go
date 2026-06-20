package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestCommonLispSourceUsesIncludes(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(body, []byte("(include \"enums/task-state.lisp\")")) {
		t.Fatal("root enum source must load semantic include files")
	}
	if bytes.Contains(body, []byte("(value StateCreated")) {
		t.Fatal("root enum source should not carry enum entries")
	}
}
