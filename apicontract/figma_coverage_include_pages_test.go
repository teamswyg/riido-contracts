package apicontract

import (
	"path/filepath"
	"testing"
)

func loadFigmaCoveragePageInventories(t *testing.T, base string, m *figmaCoverageManifest) {
	t.Helper()
	for _, file := range m.PageInventoryFiles {
		path := filepath.Join(base, file)
		var doc figmaCoveragePageInventoryDocument
		loadFigmaCoverageStrictJSON(t, path, &doc)
		if doc.SchemaVersion != figmaCoveragePageInventorySchemaVersion {
			t.Fatalf("%s schema_version = %q", path, doc.SchemaVersion)
		}
		loadFigmaCoverageInventoryNodes(t, path, &doc.Inventory)
		m.NonUITopLevelInventory = append(m.NonUITopLevelInventory, doc.Inventory)
	}
}

func loadFigmaCoverageInventoryNodes(
	t *testing.T,
	path string,
	inventory *figmaNonUITopLevelInventory,
) {
	t.Helper()
	base := filepath.Dir(path)
	for _, file := range inventory.NodeFiles {
		var doc figmaCoverageNodeDocument
		nodePath := filepath.Join(base, file)
		loadFigmaCoverageStrictJSON(t, nodePath, &doc)
		if doc.SchemaVersion != figmaCoverageNodeSchemaVersion {
			t.Fatalf("%s schema_version = %q", nodePath, doc.SchemaVersion)
		}
		inventory.Nodes = append(inventory.Nodes, doc.Node)
	}
}
