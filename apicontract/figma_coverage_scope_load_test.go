package apicontract

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func loadFigmaCoverageTestScope(t *testing.T) *figmaCoverageTestScope {
	t.Helper()
	manifestPath := filepath.FromSlash("../docs/30-architecture/figma-ai-agent-coverage.riido.json")
	docPath := filepath.FromSlash("../docs/30-architecture/figma-ai-agent-coverage.md")
	manifest := loadFigmaCoverageManifest(t, manifestPath)
	doc, err := os.ReadFile(docPath)
	if err != nil {
		t.Fatalf("read coverage doc: %v", err)
	}
	return &figmaCoverageTestScope{
		manifest:              manifest,
		docPath:               docPath,
		docText:               string(doc),
		registered:            map[string]string{},
		entryByNodeID:         map[string]figmaCoverageEntry{},
		openAPIGeneratedPaths: loadAIAgentClientGeneratedPaths(t),
		openAPITransports:     loadAIAgentClientGeneratedPathTransports(t),
		seen:                  map[string]bool{},
		nonUISeen:             map[string]bool{},
	}
}

func loadFigmaCoverageManifest(t *testing.T, path string) figmaCoverageManifest {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read coverage manifest: %v", err)
	}
	var manifest figmaCoverageManifest
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&manifest); err != nil {
		t.Fatalf("decode coverage manifest: %v", err)
	}
	var trailing struct{}
	if err := dec.Decode(&trailing); !errors.Is(err, io.EOF) {
		t.Fatalf("decode coverage manifest: trailing JSON document: %v", err)
	}
	return manifest
}
