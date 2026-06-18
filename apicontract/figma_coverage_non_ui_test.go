package apicontract

import (
	"strings"
	"testing"
)

func verifyNonUITopLevelInventory(t *testing.T, manifest figmaCoverageManifest, pages map[string]figmaCoveragePage) map[string]map[string]figmaCoverageNode {
	t.Helper()
	if len(manifest.NonUITopLevelInventory) == 0 {
		t.Fatalf("non_ui_top_level_inventory must record loaded non-UI page children")
	}
	inventory := map[string]map[string]figmaCoverageNode{}
	for _, pageInventory := range manifest.NonUITopLevelInventory {
		page, ok := pages[pageInventory.PageID]
		if !ok {
			t.Fatalf("non-UI inventory references unknown page %q", pageInventory.PageID)
		}
		if pageInventory.PageID == manifest.Figma.PageID {
			t.Fatalf("non-UI inventory must not reference primary UI page %q", pageInventory.PageID)
		}
		if _, exists := inventory[pageInventory.PageID]; exists {
			t.Fatalf("duplicate non-UI inventory page %q", pageInventory.PageID)
		}
		if got, want := len(pageInventory.Nodes), page.ChildCount; got != want {
			if !figmaNonUIInventoryDriftDocumented(manifest.SupportingToolLimitations, pageInventory.PageID, got, want) {
				t.Fatalf("non-UI inventory page %q nodes = %d, want loaded child_count %d", pageInventory.PageID, got, want)
			}
		}
		nodes := map[string]figmaCoverageNode{}
		for _, node := range pageInventory.Nodes {
			if strings.TrimSpace(node.NodeID) == "" || strings.TrimSpace(node.Name) == "" {
				t.Fatalf("non-UI inventory page %q has invalid node: %+v", pageInventory.PageID, node)
			}
			if _, exists := nodes[node.NodeID]; exists {
				t.Fatalf("duplicate non-UI inventory node %q on page %q", node.NodeID, pageInventory.PageID)
			}
			nodes[node.NodeID] = node
		}
		inventory[pageInventory.PageID] = nodes
	}
	for _, page := range pages {
		if page.NodeID == manifest.Figma.PageID {
			continue
		}
		if _, ok := inventory[page.NodeID]; !ok {
			t.Fatalf("non-UI page %q is missing loaded top-level inventory", page.NodeID)
		}
	}
	return inventory
}
