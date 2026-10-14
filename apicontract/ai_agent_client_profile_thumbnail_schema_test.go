package apicontract

import "testing"

func verifyProfileThumbnailUploadSchemas(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	requestProps := openAPISchemaProperties(t, openAPI, "CreateAgentProfileThumbnailUploadRequest")
	if _, ok := requestProps["content_type"].(map[string]any); !ok {
		t.Fatalf("CreateAgentProfileThumbnailUploadRequest content_type missing: %#v", requestProps)
	}
	if _, ok := requestProps["content_length_bytes"].(map[string]any); !ok {
		t.Fatalf("CreateAgentProfileThumbnailUploadRequest content_length_bytes missing: %#v", requestProps)
	}
	responseProps := openAPISchemaProperties(t, openAPI, "AgentProfileThumbnailUploadResponse")
	if _, ok := responseProps["upload_url"].(map[string]any); !ok {
		t.Fatalf("AgentProfileThumbnailUploadResponse upload_url missing: %#v", responseProps)
	}
	if _, ok := responseProps["form_fields"].(map[string]any); !ok {
		t.Fatalf("AgentProfileThumbnailUploadResponse form_fields missing: %#v", responseProps)
	}
	responseThumbnail, ok := responseProps["profile_thumbnail_url"].(map[string]any)
	if !ok || responseThumbnail["format"] != "uri" {
		t.Fatalf("AgentProfileThumbnailUploadResponse profile_thumbnail_url schema = %#v", responseProps["profile_thumbnail_url"])
	}
}
