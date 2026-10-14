package apicontract

import "testing"

func (s *figmaCoverageTestScope) loadExpectedPages(t *testing.T) {
	t.Helper()
	s.pages = map[string]figmaCoveragePage{}
	for _, page := range s.manifest.ExpectedPages {
		if page.NodeID == "" || page.Name == "" || page.ChildCount <= 0 {
			t.Fatalf("expected page has invalid field: %+v", page)
		}
		if _, exists := s.pages[page.NodeID]; exists {
			t.Fatalf("duplicate expected page %q", page.NodeID)
		}
		s.pages[page.NodeID] = page
	}
	if _, ok := s.pages[s.manifest.Figma.PageID]; !ok {
		t.Fatalf("primary figma page %q is missing from expected_pages", s.manifest.Figma.PageID)
	}
	if s.pages[s.manifest.Figma.PageID].ChildCount != len(s.manifest.ExpectedTopLevelNodes) {
		t.Fatalf("primary page child_count = %d, expected_top_level_nodes = %d", s.pages[s.manifest.Figma.PageID].ChildCount, len(s.manifest.ExpectedTopLevelNodes))
	}
}

func (s *figmaCoverageTestScope) loadExpectedTopLevelNodes(t *testing.T) {
	t.Helper()
	s.expected = map[string]figmaCoverageNode{}
	for _, node := range s.manifest.ExpectedTopLevelNodes {
		if node.NodeID == "" || node.Name == "" {
			t.Fatalf("expected node has empty field: %+v", node)
		}
		if _, exists := s.expected[node.NodeID]; exists {
			t.Fatalf("duplicate expected node %q", node.NodeID)
		}
		s.expected[node.NodeID] = node
	}
}
