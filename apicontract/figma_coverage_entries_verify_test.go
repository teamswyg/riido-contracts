package apicontract

import "testing"

func (s *figmaCoverageTestScope) verifyPrimaryEntries(t *testing.T) {
	t.Helper()
	for i, entry := range s.manifest.Entries {
		expectedNode, ok := s.expected[entry.NodeID]
		if !ok {
			t.Fatalf("entry %q is not in expected_top_level_nodes", entry.NodeID)
		}
		if s.seen[entry.NodeID] {
			t.Fatalf("duplicate entry node_id %q", entry.NodeID)
		}
		s.seen[entry.NodeID] = true
		s.entryByNodeID[entry.NodeID] = entry
		if entry.Name != expectedNode.Name {
			t.Fatalf("entry %q name = %q, want %q", entry.NodeID, entry.Name, expectedNode.Name)
		}
		assertCoverageDocMentionsEntry(t, s.docText, entry)
		if s.manifest.ExpectedTopLevelNodes[i].NodeID != entry.NodeID {
			t.Fatalf("entry order must match expected_top_level_nodes at %d: got %s want %s", i, entry.NodeID, s.manifest.ExpectedTopLevelNodes[i].NodeID)
		}
		verifyCoverageEntry(t, entry, s.openAPIGeneratedPaths)
	}
}
