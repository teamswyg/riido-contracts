package apicontract

import (
	"path/filepath"
	"testing"
)

func loadFigmaCoverageIncludes(t *testing.T, base string, m *figmaCoverageManifest) {
	t.Helper()
	loadFigmaCoveragePolicyIncludes(t, base, m)
	loadFigmaCoverageNodeIncludes(t, base, m)
	loadFigmaCoveragePageInventories(t, base, m)
	loadFigmaCoverageEntryIncludes(t, base, m)
	loadFigmaCoverageAnnotationIncludes(t, base, m)
}

func loadFigmaCoveragePolicyIncludes(t *testing.T, base string, m *figmaCoverageManifest) {
	t.Helper()
	for _, file := range m.ToolLimitationFiles {
		path := filepath.Join(base, file)
		var doc figmaCoverageToolLimitationDocument
		loadFigmaCoverageStrictJSON(t, path, &doc)
		if doc.SchemaVersion != figmaCoverageToolLimitationSchemaVersion {
			t.Fatalf("%s schema_version = %q", path, doc.SchemaVersion)
		}
		m.SupportingToolLimitations = append(m.SupportingToolLimitations, doc.Limitation)
	}
}
