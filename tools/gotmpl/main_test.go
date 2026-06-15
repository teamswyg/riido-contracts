package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRenderTemplateToStdout(t *testing.T) {
	tempDir := t.TempDir()
	templatePath := filepath.Join(tempDir, "body.gotmpl")
	dataPath := filepath.Join(tempDir, "data.json")
	if err := os.WriteFile(templatePath, []byte("hello {{ .Name }}"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dataPath, []byte(`{"Name":"fsm"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := run([]string{"-body", templatePath, "-data", dataPath}, &out); err != nil {
		t.Fatal(err)
	}
	if got := out.String(); got != "hello fsm" {
		t.Fatalf("output = %q, want %q", got, "hello fsm")
	}
}

func TestRenderTemplateToFile(t *testing.T) {
	tempDir := t.TempDir()
	templatePath := filepath.Join(tempDir, "body.gotmpl")
	dataPath := filepath.Join(tempDir, "data.json")
	outputPath := filepath.Join(tempDir, "out", "body.go")
	if err := os.WriteFile(templatePath, []byte("{{ .Package }}"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dataPath, []byte(`{"Package":"fsmmeta"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"-body", templatePath, "-data", dataPath, "-out", outputPath}, &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != "fsmmeta" {
		t.Fatalf("output file = %q, want %q", got, "fsmmeta")
	}
}
