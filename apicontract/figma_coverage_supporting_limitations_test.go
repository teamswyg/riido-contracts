package apicontract

import "testing"

func verifyFigmaSupportingToolLimitations(t *testing.T, limitations []figmaSupportingToolLimitation, docText string) {
	t.Helper()
	if len(limitations) == 0 {
		t.Fatalf("supporting_tool_limitations must record non-authoritative tooling failure modes")
	}
	index := figmaSupportingToolLimitationIndex(limitations)
	verifyMetadataPageListLimitation(t, index.metadataPageList, docText)
	verifyHeadlessFileKeyLimitation(t, index.headlessFileKey, docText)
	verifyOnboardingPageLoadTimeoutLimitation(t, index.onboardingPageLoadTimeout, docText)
}
