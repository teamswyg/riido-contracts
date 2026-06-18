package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyProfileThumbnailUpload(t *testing.T) {
	t.Helper()
	profileThumbnailUpload := f.openAPI.Paths["/v1/client/ai-agent/profile-thumbnails/uploads"]["post"]
	if profileThumbnailUpload.OperationID != "createAIAgentProfileThumbnailUpload" ||
		profileThumbnailUpload.RequestBody == nil ||
		profileThumbnailUpload.RiidoClient == nil ||
		profileThumbnailUpload.RiidoClient.GeneratedPath != "aiAgent.profileThumbnails.uploads.create" ||
		profileThumbnailUpload.RiidoRBAC != "agent_profile_thumbnail_upload.v1" {
		t.Fatalf("profile thumbnail upload operation = %#v", profileThumbnailUpload)
	}
	f.verifyProfileThumbnailUploadV2(t)
	verifyProfileThumbnailUploadSchemas(t, f.openAPI)
}

func (f aiAgentClientContractFixture) verifyProfileThumbnailUploadV2(t *testing.T) {
	t.Helper()
	upload := f.openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/profile-thumbnails/uploads"]["post"]
	if upload.OperationID != "createAIAgentProfileThumbnailUploadV2" ||
		upload.RequestBody == nil ||
		upload.RiidoClient == nil ||
		upload.RiidoClient.GeneratedPath != "v2.aiAgent.profileThumbnails.uploads.create" ||
		upload.RiidoRBAC != "agent_profile_thumbnail_upload.v1" {
		t.Fatalf("v2 profile thumbnail upload operation = %#v", upload)
	}
	if len(upload.Parameters) != 1 || upload.Parameters[0].Name != "workspace_id" {
		t.Fatalf("v2 profile thumbnail upload parameters = %#v", upload.Parameters)
	}
}
