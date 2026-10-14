package apicontract

import "testing"

func (s *figmaCoverageTestScope) verifyNonUIEntries(t *testing.T) {
	t.Helper()
	for _, entry := range s.manifest.NonUITopLevelNodes {
		if _, ok := s.pages[entry.PageID]; !ok {
			t.Fatalf("non-UI entry %q references unknown page %q", entry.NodeID, entry.PageID)
		}
		if entry.PageID == s.manifest.Figma.PageID {
			t.Fatalf("non-UI entry %q must not reference primary UI page", entry.NodeID)
		}
		if s.nonUISeen[entry.NodeID] {
			t.Fatalf("duplicate non-UI entry node_id %q", entry.NodeID)
		}
		s.nonUISeen[entry.NodeID] = true
		if _, ok := s.nonUIInventory[entry.PageID][entry.NodeID]; !ok {
			t.Fatalf("non-UI coverage entry %q is missing from loaded top-level inventory for page %q", entry.NodeID, entry.PageID)
		}
		registerFigmaNode(t, s.registered, figmaCoverageNode{NodeID: entry.NodeID, Name: entry.Name}, "non_ui_top_level_nodes")
		assertCoverageDocMentionsEntry(t, s.docText, entry)
		verifyCoverageEntry(t, entry, s.openAPIGeneratedPaths)
	}
	verifySemanticLegacyWireframeCoverage(t, s.manifest.NonUITopLevelNodes, s.nonUIInventory, s.entryByNodeID, s.docText)
	s.registerNonUIInventoryRemainder(t)
}
