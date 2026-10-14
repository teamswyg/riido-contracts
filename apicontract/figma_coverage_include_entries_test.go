package apicontract

import (
	"path/filepath"
	"testing"
)

func loadFigmaCoverageEntryIncludes(t *testing.T, base string, m *figmaCoverageManifest) {
	t.Helper()
	m.Entries = append(
		m.Entries,
		loadFigmaCoverageEntries(t, base, m.CoverageEntryFiles)...,
	)
	m.NonUITopLevelNodes = append(
		m.NonUITopLevelNodes,
		loadFigmaCoverageEntries(t, base, m.NonUICoverageEntryFiles)...,
	)
}

func loadFigmaCoverageEntries(t *testing.T, base string, files []string) []figmaCoverageEntry {
	t.Helper()
	entries := make([]figmaCoverageEntry, 0, len(files))
	for _, file := range files {
		path := filepath.Join(base, file)
		var doc figmaCoverageEntryDocument
		loadFigmaCoverageStrictJSON(t, path, &doc)
		if doc.SchemaVersion != figmaCoverageEntrySchemaVersion {
			t.Fatalf("%s schema_version = %q", path, doc.SchemaVersion)
		}
		entries = append(entries, doc.Entry)
	}
	return entries
}
