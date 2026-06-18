package apicontract

import "testing"

func (s *figmaCoverageTestScope) verifyFinalAssertions(t *testing.T) {
	t.Helper()
	for _, node := range s.manifest.ExpectedTopLevelNodes {
		if !s.seen[node.NodeID] {
			t.Fatalf("expected node %q has no entry", node.NodeID)
		}
	}
	verifyFigmaRuntimeEndpointLabel(t, s.manifest.VerifiedEvidenceNodes, s.entryByNodeID["162:23090"], s.docText)
	verifyFigmaAPIGeneratedAnnotations(t, s.manifest.APIGeneratedAnnotations, s.docText, s.openAPIGeneratedPaths, s.registered, s.entryByNodeID)
	verifyFigmaAPIGeneratedAnnotationInventory(t, s.manifest.APIGeneratedAnnotationInventory, s.docText, s.openAPIGeneratedPaths, s.openAPITransports, s.registered, s.entryByNodeID)
	assertDocumentedFigmaNodeRefsAreRegistered(t, s.registered)
	assertNoStaleOnboardingFixtureWording(t)
	assertNoStaleRuntimeEndpointHostPinned(t)
}
