package apicontract

import "testing"

func TestFigmaAIAgentCoverageManifest(t *testing.T) {
	scope := loadFigmaCoverageTestScope(t)
	scope.verifyManifestEnvelope(t)
	scope.verifyManifestPolicy(t)
	scope.verifyManifestCounts(t)
	scope.loadExpectedPages(t)
	scope.loadExpectedTopLevelNodes(t)
	scope.registerManifestEvidence(t)
	scope.verifyPrimaryEntries(t)
	scope.verifyNonUIEntries(t)
	scope.verifyFinalAssertions(t)
}
