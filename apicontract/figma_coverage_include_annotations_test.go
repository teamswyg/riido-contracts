package apicontract

import (
	"path/filepath"
	"testing"
)

func loadFigmaCoverageAnnotationIncludes(t *testing.T, base string, m *figmaCoverageManifest) {
	t.Helper()
	for _, file := range m.APIGeneratedAnnotationInventoryFiles {
		path := filepath.Join(base, file)
		var doc figmaCoverageAnnotationInventoryDocument
		loadFigmaCoverageStrictJSON(t, path, &doc)
		if doc.SchemaVersion != figmaCoverageAnnotationInventorySchemaVersion {
			t.Fatalf("%s schema_version = %q", path, doc.SchemaVersion)
		}
		m.APIGeneratedAnnotationInventory = append(m.APIGeneratedAnnotationInventory, doc.Inventory)
	}
	for _, file := range m.APIGeneratedAnnotationFiles {
		path := filepath.Join(base, file)
		var doc figmaCoverageAnnotationDocument
		loadFigmaCoverageStrictJSON(t, path, &doc)
		if doc.SchemaVersion != figmaCoverageAnnotationSchemaVersion {
			t.Fatalf("%s schema_version = %q", path, doc.SchemaVersion)
		}
		m.APIGeneratedAnnotations = append(m.APIGeneratedAnnotations, doc.Annotation)
	}
}
