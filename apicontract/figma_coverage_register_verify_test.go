package apicontract

import "testing"

func (s *figmaCoverageTestScope) registerManifestEvidence(t *testing.T) {
	t.Helper()
	for _, page := range s.manifest.ExpectedPages {
		registerFigmaNode(t, s.registered, figmaCoverageNode{NodeID: page.NodeID, Name: page.Name}, "expected_pages")
	}
	for _, node := range s.manifest.ExpectedTopLevelNodes {
		registerFigmaNode(t, s.registered, node, "expected_top_level_nodes")
	}
	for _, node := range s.manifest.VerifiedEvidenceNodes {
		registerFigmaNode(t, s.registered, node, "verified_evidence_nodes")
	}
	s.nonUIInventory = verifyNonUITopLevelInventory(t, s.manifest, s.pages)
}

func (s *figmaCoverageTestScope) registerNonUIInventoryRemainder(t *testing.T) {
	t.Helper()
	for pageID, nodes := range s.nonUIInventory {
		for nodeID, node := range nodes {
			if s.nonUISeen[nodeID] {
				continue
			}
			registerFigmaNodeIfAbsent(t, s.registered, node, "non_ui_top_level_inventory page "+pageID)
		}
	}
}
