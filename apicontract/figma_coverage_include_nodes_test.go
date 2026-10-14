package apicontract

import (
	"path/filepath"
	"testing"
)

func loadFigmaCoverageNodeIncludes(t *testing.T, base string, m *figmaCoverageManifest) {
	t.Helper()
	m.ExpectedTopLevelNodes = append(
		m.ExpectedTopLevelNodes,
		loadFigmaCoverageNodes(t, base, m.ExpectedTopLevelNodeFiles)...,
	)
	m.VerifiedEvidenceNodes = append(
		m.VerifiedEvidenceNodes,
		loadFigmaCoverageNodes(t, base, m.VerifiedEvidenceNodeFiles)...,
	)
}

func loadFigmaCoverageNodes(t *testing.T, base string, files []string) []figmaCoverageNode {
	t.Helper()
	nodes := make([]figmaCoverageNode, 0, len(files))
	for _, file := range files {
		path := filepath.Join(base, file)
		var doc figmaCoverageNodeDocument
		loadFigmaCoverageStrictJSON(t, path, &doc)
		if doc.SchemaVersion != figmaCoverageNodeSchemaVersion {
			t.Fatalf("%s schema_version = %q", path, doc.SchemaVersion)
		}
		nodes = append(nodes, doc.Node)
	}
	return nodes
}
