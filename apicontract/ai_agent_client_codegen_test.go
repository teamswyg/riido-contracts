package apicontract

import "testing"

func TestAIAgentClientDSLKeepsEnumsAndSumTypesCodegenSafe(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	fixture.verifyIdentity(t)
	fixture.verifyCoreSchemas(t)
	fixture.verifyBootstrapSchema(t)
	fixture.verifyClientModules(t)
	fixture.verifyOnboarding(t)
	fixture.verifyProfileThumbnailUpload(t)
	fixture.verifyAgentConfigurationRequests(t)
	fixture.verifyV2WorkspaceContracts(t)
	fixture.verifyCommentKindSchema(t)
	fixture.verifyTaskThreadContracts(t)
	fixture.verifyAgentRecordSchemas(t)
	fixture.verifyRuntimeSchemas(t)
}
