package apicontract

import (
	"bytes"
	"encoding/json"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestAgentCatalogDSLGeneratesIRAndOpenAPI(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	if ir.SchemaVersion != IRSchemaVersion || ir.SourceSchemaVersion != DSLSchemaVersion {
		t.Fatalf("IR versions = %q / %q", ir.SchemaVersion, ir.SourceSchemaVersion)
	}
	if got, want := len(ir.Operations), 5; got != want {
		t.Fatalf("IR operations = %d, want %d", got, want)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	if openAPI.OpenAPI != OpenAPIVersion {
		t.Fatalf("OpenAPI version = %q", openAPI.OpenAPI)
	}
	list := openAPI.Paths["/v1/agent-catalog"]["get"]
	if list.OperationID != "listAgents" {
		t.Fatalf("list operation = %+v", list)
	}
	if len(list.RiidoScopes) != 1 || list.RiidoScopes[0] != "agent-catalog:read" {
		t.Fatalf("list scopes = %v", list.RiidoScopes)
	}
	update := openAPI.Paths["/v1/agent-catalog/{agent_id}"]["patch"]
	if len(update.Parameters) != 1 || update.Parameters[0].Name != "agent_id" {
		t.Fatalf("update path parameters = %+v", update.Parameters)
	}
	if update.RiidoRBAC != "agent_catalog_visibility.v1" {
		t.Fatalf("update rbac = %q", update.RiidoRBAC)
	}
}

func TestAgentCatalogGeneratedFixturesDoNotDrift(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	assertFixture(t, "fixtures/control-plane-agent-catalog.ir.riido.json", ir)
	assertFixture(t, "fixtures/control-plane-agent-catalog.openapi.json", openAPI)
}

func TestGenerateIRRejectsInvalidDSL(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	dsl.SchemaVersion = "riido-api-dsl.v0"
	if _, err := GenerateIR(dsl); err == nil {
		t.Fatal("expected unsupported schema version error")
	}
}

func TestAIAgentClientDSLKeepsEnumsAndSumTypesCodegenSafe(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	if got, want := len(ir.Enums), 12; got != want {
		t.Fatalf("IR enums = %d, want %d", got, want)
	}
	if got, want := len(ir.SumTypes), 1; got != want {
		t.Fatalf("IR sum types = %d, want %d", got, want)
	}
	if ir.ContractID != "control-plane-ai-agent-client-api.v2" {
		t.Fatalf("IR contract_id = %q", ir.ContractID)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	if openAPI.Info.Title != "control-plane-ai-agent-client-api.v2" {
		t.Fatalf("OpenAPI title = %q", openAPI.Info.Title)
	}
	runtimeAvailability := openAPI.Components.Schemas["RuntimeAvailability"]
	values, ok := runtimeAvailability["enum"].([]string)
	if !ok || len(values) != 2 || values[0] != "online" || values[1] != "offline" {
		t.Fatalf("RuntimeAvailability enum = %#v", runtimeAvailability["enum"])
	}
	streamEvent := openAPI.Components.Schemas["ClientStreamEvent"]
	if _, ok := streamEvent["oneOf"].([]map[string]any); !ok {
		t.Fatalf("ClientStreamEvent oneOf missing: %#v", streamEvent)
	}
	streamOperation := openAPI.Paths["/v1/client/ai-agent/events"]["get"]
	if _, ok := streamOperation.Responses["200"].Content["text/event-stream"]; !ok {
		t.Fatalf("stream response content = %#v", streamOperation.Responses["200"].Content)
	}
	if len(openAPI.RiidoClientModules) != 2 ||
		openAPI.RiidoClientModules[0].Module != "aiAgent" ||
		openAPI.RiidoClientModules[1].Module != "v2" {
		t.Fatalf("client modules = %#v", openAPI.RiidoClientModules)
	}
	if _, ok := openAPI.Components.SecuritySchemes["riidoAIAgentToken"]; !ok {
		t.Fatalf("riidoAIAgentToken security scheme missing: %#v", openAPI.Components.SecuritySchemes)
	}
	if _, ok := openAPI.Components.Schemas["AgentOnboardingTemplate"]; ok {
		t.Fatalf("AgentOnboardingTemplate must not be exposed")
	}
	fixtureList := openAPI.Paths["/v1/client/ai-agent/onboarding/fixtures"]["get"]
	if fixtureList.OperationID != "listAIAgentOnboardingFixtures" ||
		fixtureList.RiidoClient == nil ||
		fixtureList.RiidoClient.CacheTag != "aiAgent.onboarding.fixtures" ||
		fixtureList.RiidoClient.GeneratedPath != "aiAgent.onboarding.fixtures" ||
		fixtureList.RiidoRBAC != "agent_onboarding_fixtures.v1" {
		t.Fatalf("fixture list operation = %#v", fixtureList)
	}
	fixtureCreate := openAPI.Paths["/v1/client/ai-agent/onboarding/fixtures/{fixture_id}/agents"]["post"]
	if fixtureCreate.OperationID != "createAIAgentFromOnboardingFixture" ||
		fixtureCreate.RequestBody == nil ||
		fixtureCreate.RiidoClient == nil ||
		fixtureCreate.RiidoClient.GeneratedPath != "aiAgent.onboarding.fixtures.createAgent" ||
		!contains(fixtureCreate.RiidoClient.Invalidates, "aiAgent.bootstrap") {
		t.Fatalf("fixture create operation = %#v", fixtureCreate)
	}
	if len(fixtureCreate.Parameters) != 1 || fixtureCreate.Parameters[0].Name != "fixture_id" {
		t.Fatalf("fixture create parameters = %#v", fixtureCreate.Parameters)
	}
	agentCreateV2 := openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/agents"]["post"]
	if agentCreateV2.OperationID != "createAIAgentV2" ||
		agentCreateV2.RequestBody == nil ||
		agentCreateV2.RiidoClient == nil ||
		agentCreateV2.RiidoClient.GeneratedPath != "v2.aiAgent.agents.create" ||
		!contains(agentCreateV2.RiidoClient.Invalidates, "v2.aiAgent.bootstrap") {
		t.Fatalf("v2 agent create operation = %#v", agentCreateV2)
	}
	if len(agentCreateV2.Parameters) != 1 || agentCreateV2.Parameters[0].Name != "workspace_id" {
		t.Fatalf("v2 agent create parameters = %#v", agentCreateV2.Parameters)
	}
	profileThumbnailUpload := openAPI.Paths["/v1/client/ai-agent/profile-thumbnails/uploads"]["post"]
	if profileThumbnailUpload.OperationID != "createAIAgentProfileThumbnailUpload" ||
		profileThumbnailUpload.RequestBody == nil ||
		profileThumbnailUpload.RiidoClient == nil ||
		profileThumbnailUpload.RiidoClient.GeneratedPath != "aiAgent.profileThumbnails.uploads.create" ||
		profileThumbnailUpload.RiidoRBAC != "agent_profile_thumbnail_upload.v1" {
		t.Fatalf("profile thumbnail upload operation = %#v", profileThumbnailUpload)
	}
	profileThumbnailUploadV2 := openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/profile-thumbnails/uploads"]["post"]
	if profileThumbnailUploadV2.OperationID != "createAIAgentProfileThumbnailUploadV2" ||
		profileThumbnailUploadV2.RequestBody == nil ||
		profileThumbnailUploadV2.RiidoClient == nil ||
		profileThumbnailUploadV2.RiidoClient.GeneratedPath != "v2.aiAgent.profileThumbnails.uploads.create" ||
		profileThumbnailUploadV2.RiidoRBAC != "agent_profile_thumbnail_upload.v1" {
		t.Fatalf("v2 profile thumbnail upload operation = %#v", profileThumbnailUploadV2)
	}
	if len(profileThumbnailUploadV2.Parameters) != 1 || profileThumbnailUploadV2.Parameters[0].Name != "workspace_id" {
		t.Fatalf("v2 profile thumbnail upload parameters = %#v", profileThumbnailUploadV2.Parameters)
	}
	createRequest := openAPI.Components.Schemas["CreateAgentConfigurationRequest"]
	createRequestDescription, ok := createRequest["description"].(string)
	if !ok ||
		!strings.Contains(createRequestDescription, "프로필 사진") ||
		!strings.Contains(createRequestDescription, "런타임") ||
		!strings.Contains(createRequestDescription, "모델") ||
		!strings.Contains(createRequestDescription, "지침") {
		t.Fatalf("CreateAgentConfigurationRequest description must explain Figma agent setting fields: %q", createRequestDescription)
	}
	createRequestRequired, ok := createRequest["required"].([]string)
	if !ok ||
		!contains(createRequestRequired, "name") ||
		!contains(createRequestRequired, "visibility") ||
		!contains(createRequestRequired, "runtime_id") ||
		contains(createRequestRequired, "model_id") {
		t.Fatalf("CreateAgentConfigurationRequest required = %#v", createRequest["required"])
	}
	createRequestProps, ok := createRequest["properties"].(map[string]any)
	if !ok {
		t.Fatalf("CreateAgentConfigurationRequest properties missing: %#v", createRequest)
	}
	for _, propertyName := range []string{"name", "profile_thumbnail_url", "description", "runtime_id", "model_id", "visibility", "instruction"} {
		if _, ok := createRequestProps[propertyName].(map[string]any); !ok {
			t.Fatalf("CreateAgentConfigurationRequest missing %s: %#v", propertyName, createRequestProps)
		}
	}
	createThumbnail, ok := createRequestProps["profile_thumbnail_url"].(map[string]any)
	if !ok || createThumbnail["format"] != "uri" {
		t.Fatalf("CreateAgentConfigurationRequest profile_thumbnail_url schema = %#v", createRequestProps["profile_thumbnail_url"])
	}
	createDescription, ok := createRequestProps["description"].(map[string]any)
	if !ok || createDescription["maxLength"] != 160 {
		t.Fatalf("CreateAgentConfigurationRequest description schema = %#v", createRequestProps["description"])
	}
	createInstruction, ok := createRequestProps["instruction"].(map[string]any)
	if !ok || createInstruction["maxLength"] != 1000 {
		t.Fatalf("CreateAgentConfigurationRequest instruction schema = %#v", createRequestProps["instruction"])
	}
	profileThumbnailUploadRequest := openAPI.Components.Schemas["CreateAgentProfileThumbnailUploadRequest"]
	profileThumbnailUploadRequestProps, ok := profileThumbnailUploadRequest["properties"].(map[string]any)
	if !ok {
		t.Fatalf("CreateAgentProfileThumbnailUploadRequest properties missing: %#v", profileThumbnailUploadRequest)
	}
	if _, ok := profileThumbnailUploadRequestProps["content_type"].(map[string]any); !ok {
		t.Fatalf("CreateAgentProfileThumbnailUploadRequest content_type missing: %#v", profileThumbnailUploadRequestProps)
	}
	if _, ok := profileThumbnailUploadRequestProps["content_length_bytes"].(map[string]any); !ok {
		t.Fatalf("CreateAgentProfileThumbnailUploadRequest content_length_bytes missing: %#v", profileThumbnailUploadRequestProps)
	}
	profileThumbnailUploadResponse := openAPI.Components.Schemas["AgentProfileThumbnailUploadResponse"]
	profileThumbnailUploadResponseProps, ok := profileThumbnailUploadResponse["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AgentProfileThumbnailUploadResponse properties missing: %#v", profileThumbnailUploadResponse)
	}
	if _, ok := profileThumbnailUploadResponseProps["upload_url"].(map[string]any); !ok {
		t.Fatalf("AgentProfileThumbnailUploadResponse upload_url missing: %#v", profileThumbnailUploadResponseProps)
	}
	if _, ok := profileThumbnailUploadResponseProps["form_fields"].(map[string]any); !ok {
		t.Fatalf("AgentProfileThumbnailUploadResponse form_fields missing: %#v", profileThumbnailUploadResponseProps)
	}
	responseThumbnail, ok := profileThumbnailUploadResponseProps["profile_thumbnail_url"].(map[string]any)
	if !ok || responseThumbnail["format"] != "uri" {
		t.Fatalf("AgentProfileThumbnailUploadResponse profile_thumbnail_url schema = %#v", profileThumbnailUploadResponseProps["profile_thumbnail_url"])
	}
	updateRequest := openAPI.Components.Schemas["UpdateAgentConfigurationRequest"]
	updateRequestDescription, ok := updateRequest["description"].(string)
	if !ok ||
		!strings.Contains(updateRequestDescription, "프로필 사진") ||
		!strings.Contains(updateRequestDescription, "런타임") ||
		!strings.Contains(updateRequestDescription, "모델") ||
		!strings.Contains(updateRequestDescription, "지침") {
		t.Fatalf("UpdateAgentConfigurationRequest description must explain Figma agent setting fields: %q", updateRequestDescription)
	}
	updateRequestProps, ok := updateRequest["properties"].(map[string]any)
	if !ok {
		t.Fatalf("UpdateAgentConfigurationRequest properties missing: %#v", updateRequest)
	}
	for _, propertyName := range []string{"name", "profile_thumbnail_url", "description", "runtime_id", "model_id", "visibility", "instruction"} {
		if _, ok := updateRequestProps[propertyName].(map[string]any); !ok {
			t.Fatalf("UpdateAgentConfigurationRequest missing %s: %#v", propertyName, updateRequestProps)
		}
	}
	agentV2 := openAPI.Components.Schemas["AgentClientRecordV2"]
	agentV2Props, ok := agentV2["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AgentClientRecordV2 properties missing: %#v", agentV2)
	}
	if _, ok := agentV2Props["workspace_id"].(map[string]any); !ok {
		t.Fatalf("AgentClientRecordV2 workspace_id missing: %#v", agentV2Props)
	}
	v2Required, ok := agentV2["required"].([]string)
	if !ok || !contains(v2Required, "workspace_id") {
		t.Fatalf("AgentClientRecordV2 required = %#v", agentV2["required"])
	}
	v2Bootstrap := openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/bootstrap"]["get"]
	if v2Bootstrap.RiidoClient == nil ||
		v2Bootstrap.RiidoClient.GeneratedPath != "v2.aiAgent.bootstrap" ||
		v2Bootstrap.RiidoClient.CacheTag != "v2.aiAgent.bootstrap" {
		t.Fatalf("v2 bootstrap client metadata = %#v", v2Bootstrap.RiidoClient)
	}
	if !strings.Contains(v2Bootstrap.Summary, "v2.aiAgent.bootstrap.agents[]") {
		t.Fatalf("v2 bootstrap summary must name agents[] source for generated comments: %q", v2Bootstrap.Summary)
	}
	assignedProfiles := openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/tasks/assigned-agent-profiles"]["get"]
	if assignedProfiles.OperationID != "listWorkspaceAssignedAgentProfilesV2" ||
		assignedProfiles.RiidoClient == nil ||
		assignedProfiles.RiidoClient.GeneratedPath != "v2.aiAgent.tasks.assignedAgentProfiles" ||
		assignedProfiles.RiidoClient.CacheTag != "v2.aiAgent.tasks.assignedAgentProfiles" ||
		assignedProfiles.RiidoRBAC != "workspace_assigned_agent_profile_map.v1" {
		t.Fatalf("assigned profiles operation = %#v", assignedProfiles)
	}
	if len(assignedProfiles.Parameters) != 1 || assignedProfiles.Parameters[0].Name != "workspace_id" {
		t.Fatalf("assigned profiles parameters = %#v", assignedProfiles.Parameters)
	}
	threadMessageCreate := openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/threads/{thread_id}/messages"]["post"]
	if threadMessageCreate.OperationID != "createAIAgentTaskThreadMessage" ||
		threadMessageCreate.RiidoClient == nil ||
		threadMessageCreate.RiidoClient.GeneratedPath != "aiAgent.tasks.threadMessages.create" {
		t.Fatalf("thread message create operation = %#v", threadMessageCreate)
	}
	fixtureSchema := openAPI.Components.Schemas["AgentOnboardingFixture"]
	fixtureProps, ok := fixtureSchema["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AgentOnboardingFixture properties missing: %#v", fixtureSchema)
	}
	fixtureID, ok := fixtureProps["fixture_id"].(map[string]any)
	if !ok || fixtureID["description"] == "" {
		t.Fatalf("fixture_id description missing: %#v", fixtureProps["fixture_id"])
	}
	if _, ok := fixtureProps["tmp_color"].(map[string]any); !ok {
		t.Fatalf("fixture tmp_color schema missing: %#v", fixtureProps)
	}
	bootstrapSchema := openAPI.Components.Schemas["ClientBootstrapResponse"]
	bootstrapProps, ok := bootstrapSchema["properties"].(map[string]any)
	if !ok {
		t.Fatalf("ClientBootstrapResponse properties missing: %#v", bootstrapSchema)
	}
	if _, ok := bootstrapProps["agent_templates"]; ok {
		t.Fatalf("ClientBootstrapResponse must not expose agent_templates: %#v", bootstrapProps)
	}
	bootstrapAgents, ok := bootstrapProps["agents"].(map[string]any)
	if !ok {
		t.Fatalf("ClientBootstrapResponse agents property missing: %#v", bootstrapProps["agents"])
	}
	bootstrapAgentsDescription, ok := bootstrapAgents["description"].(string)
	if !ok ||
		!strings.Contains(bootstrapAgentsDescription, "agent_id") ||
		!strings.Contains(bootstrapAgentsDescription, "tasks.assignableAgents") {
		t.Fatalf("ClientBootstrapResponse agents description must explain settings/list agent_id source and task dropdown boundary: %q", bootstrapAgentsDescription)
	}
	daemonDetail := openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/daemon"]["get"]
	if daemonDetail.RiidoClient == nil || daemonDetail.RiidoClient.CacheTag != "aiAgent.agents.daemon" {
		t.Fatalf("daemon detail client metadata = %#v", daemonDetail.RiidoClient)
	}
	if len(daemonDetail.Parameters) != 1 || daemonDetail.Parameters[0].Name != "agent_id" {
		t.Fatalf("daemon detail parameters = %#v", daemonDetail.Parameters)
	}
	daemonStop := openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/daemon/stop"]["post"]
	if daemonStop.RiidoClient == nil || !contains(daemonStop.RiidoClient.Invalidates, "aiAgent.devices.runtimes") {
		t.Fatalf("daemon stop client metadata = %#v", daemonStop.RiidoClient)
	}
	editability := openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/editability"]["get"]
	if editability.RiidoRBAC != "agent_mutation_safety.v1" {
		t.Fatalf("editability rbac = %q", editability.RiidoRBAC)
	}
	assignable := openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/assignable-agents"]["get"]
	if len(assignable.Parameters) != 1 || assignable.Parameters[0].Name != "task_id" {
		t.Fatalf("assignable-agent parameters = %#v", assignable.Parameters)
	}
	assignableList := openAPI.Components.Schemas["AgentClientListResponse"]
	assignableListProps, ok := assignableList["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AgentClientListResponse properties missing: %#v", assignableList)
	}
	assignableAgents, ok := assignableListProps["agents"].(map[string]any)
	if !ok {
		t.Fatalf("AgentClientListResponse agents property missing: %#v", assignableListProps["agents"])
	}
	assignableAgentsDescription, ok := assignableAgents["description"].(string)
	if !ok ||
		!strings.Contains(assignableAgentsDescription, "agent_id") ||
		!strings.Contains(assignableAgentsDescription, "tasks.assign") {
		t.Fatalf("AgentClientListResponse agents description must explain task dropdown agent_id source: %q", assignableAgentsDescription)
	}
	threads := openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/threads"]["get"]
	if threads.RiidoRBAC != "task_thread_cold_collection.v1" {
		t.Fatalf("task threads rbac = %q", threads.RiidoRBAC)
	}
	if len(threads.Parameters) != 1 || threads.Parameters[0].Name != "task_id" {
		t.Fatalf("task threads parameters = %#v", threads.Parameters)
	}
	commentKind := openAPI.Components.Schemas["AgentTaskCommentKind"]
	commentValues, ok := commentKind["enum"].([]string)
	if !ok || len(commentValues) == 0 || commentValues[0] != "queued_by_busy_agent" {
		t.Fatalf("AgentTaskCommentKind enum = %#v", commentKind["enum"])
	}
	agentRecord := openAPI.Components.Schemas["AgentClientRecord"]
	recordProps, ok := agentRecord["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AgentClientRecord properties missing: %#v", agentRecord)
	}
	thumbnail, ok := recordProps["profile_thumbnail_url"].(map[string]any)
	if !ok || thumbnail["format"] != "uri" {
		t.Fatalf("profile_thumbnail_url schema = %#v", recordProps["profile_thumbnail_url"])
	}
	description, ok := recordProps["description"].(map[string]any)
	if !ok || description["maxLength"] != 160 {
		t.Fatalf("description schema = %#v", recordProps["description"])
	}
	instruction, ok := recordProps["instruction"].(map[string]any)
	if !ok || instruction["maxLength"] != 1000 {
		t.Fatalf("instruction schema = %#v", recordProps["instruction"])
	}
	if _, ok := recordProps["model_id"].(map[string]any); !ok {
		t.Fatalf("model_id schema missing: %#v", recordProps)
	}
	if _, ok := recordProps["tmp_color"].(map[string]any); !ok {
		t.Fatalf("tmp_color schema missing: %#v", recordProps)
	}
	assignedProfileMap := openAPI.Components.Schemas["AssignedAgentProfileMapResponse"]
	assignedProfileMapProps, ok := assignedProfileMap["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AssignedAgentProfileMapResponse properties missing: %#v", assignedProfileMap)
	}
	assignedProfileValues, ok := assignedProfileMapProps["assigned_agent_profiles"].(map[string]any)
	if !ok {
		t.Fatalf("assigned_agent_profiles property missing: %#v", assignedProfileMapProps)
	}
	additional, ok := assignedProfileValues["additionalProperties"].(map[string]any)
	if !ok || additional["$ref"] != "#/components/schemas/AssignedAgentProfile" {
		t.Fatalf("assigned_agent_profiles additionalProperties = %#v", assignedProfileValues["additionalProperties"])
	}
	agentID, ok := recordProps["agent_id"].(map[string]any)
	if !ok {
		t.Fatalf("agent_id schema missing: %#v", recordProps)
	}
	agentIDDescription, ok := agentID["description"].(string)
	if !ok ||
		!strings.Contains(agentIDDescription, "bootstrap.agents[]") ||
		!strings.Contains(agentIDDescription, "tasks.assignableAgents.agents[]") {
		t.Fatalf("agent_id description must explain bootstrap vs assignableAgents source: %q", agentIDDescription)
	}
	createdAt, ok := recordProps["created_at"].(map[string]any)
	if !ok || createdAt["format"] != "date-time" {
		t.Fatalf("created_at schema = %#v", recordProps["created_at"])
	}
	updatedAt, ok := recordProps["updated_at"].(map[string]any)
	if !ok || updatedAt["format"] != "date-time" {
		t.Fatalf("updated_at schema = %#v", recordProps["updated_at"])
	}
	runtimeRecord := openAPI.Components.Schemas["RuntimeRecord"]
	runtimeProps, ok := runtimeRecord["properties"].(map[string]any)
	if !ok {
		t.Fatalf("RuntimeRecord properties missing: %#v", runtimeRecord)
	}
	runtimeRequired, ok := runtimeRecord["required"].([]string)
	if !ok || !contains(runtimeRequired, "requires_experimental_opt_in") {
		t.Fatalf("RuntimeRecord required = %#v", runtimeRecord["required"])
	}
	optIn, ok := runtimeProps["requires_experimental_opt_in"].(map[string]any)
	if !ok || optIn["type"] != "boolean" {
		t.Fatalf("RuntimeRecord requires_experimental_opt_in schema = %#v", runtimeProps["requires_experimental_opt_in"])
	}
	models, ok := runtimeProps["models"].(map[string]any)
	if !ok || models["type"] != "array" {
		t.Fatalf("RuntimeRecord models schema = %#v", runtimeProps["models"])
	}
	providerVersion, ok := runtimeProps["provider_version"].(map[string]any)
	if !ok || providerVersion["type"] != "string" {
		t.Fatalf("RuntimeRecord provider_version schema = %#v", runtimeProps["provider_version"])
	}
	daemonRecord := openAPI.Components.Schemas["DeviceDaemonRecord"]
	daemonRequired, ok := daemonRecord["required"].([]string)
	if !ok || !contains(daemonRequired, "supported_actions") {
		t.Fatalf("DeviceDaemonRecord required = %#v", daemonRecord["required"])
	}
	daemonStatusEvent := openAPI.Components.Schemas["DeviceDaemonStatusEvent"]
	daemonStatusProps, ok := daemonStatusEvent["properties"].(map[string]any)
	if !ok {
		t.Fatalf("DeviceDaemonStatusEvent properties missing: %#v", daemonStatusEvent)
	}
	if _, ok := daemonStatusProps["daemon"].(map[string]any); !ok {
		t.Fatalf("DeviceDaemonStatusEvent daemon schema missing: %#v", daemonStatusProps)
	}
	runtimeModel := openAPI.Components.Schemas["RuntimeModelRecord"]
	modelRequired, ok := runtimeModel["required"].([]string)
	if !ok || !contains(modelRequired, "model_id") || !contains(modelRequired, "is_default") {
		t.Fatalf("RuntimeModelRecord required = %#v", runtimeModel["required"])
	}
	threadCollection := openAPI.Components.Schemas["AIAgentTaskThreadCollectionResponse"]
	threadCollectionProps, ok := threadCollection["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AIAgentTaskThreadCollectionResponse properties missing: %#v", threadCollection)
	}
	if _, ok := threadCollectionProps["active_stream"].(map[string]any); !ok {
		t.Fatalf("active_stream schema missing: %#v", threadCollectionProps["active_stream"])
	}
	progressEvent := openAPI.Components.Schemas["AgentThreadProgressEvent"]
	progressRequired, ok := progressEvent["required"].([]string)
	if !ok || !contains(progressRequired, "thread_id") {
		t.Fatalf("AgentThreadProgressEvent required = %#v", progressEvent["required"])
	}
}

func TestAIAgentClientGeneratedFixturesDoNotDrift(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	assertFixture(t, "fixtures/control-plane-ai-agent-client.ir.riido.json", ir)
	assertFixture(t, "fixtures/control-plane-ai-agent-client.openapi.json", openAPI)
}

func TestAIAgentClientContractRejectsAmbiguousFutureClientBootstrapWording(t *testing.T) {
	for _, path := range []string{
		"fixtures/control-plane-ai-agent-client.dsl.riido.json",
		"fixtures/control-plane-ai-agent-client.ir.riido.json",
		"fixtures/control-plane-ai-agent-client.openapi.json",
	} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		if strings.Contains(strings.ToLower(string(data)), "future client bootstrap") {
			t.Fatalf("%s contains ambiguous future-client wording; use subsequent aiAgent.bootstrap read wording", path)
		}
	}
}

func TestAIAgentClientOpenAPIDoesNotExposeRuntimeWaitlistMutation(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	for path, methods := range openAPI.Paths {
		for method, operation := range methods {
			haystack := strings.ToLower(path + " " + method + " " + operation.OperationID + " " + operation.Summary)
			for _, forbidden := range []string{"waitlist", "marketing", "consent"} {
				if strings.Contains(haystack, forbidden) {
					t.Fatalf("AI Agent OpenAPI exposed %q in %s %s (%s)", forbidden, method, path, operation.OperationID)
				}
			}
		}
	}
}

func loadTestDSL(t *testing.T, path string) DSLDocument {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read DSL fixture: %v", err)
	}
	var dsl DSLDocument
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		t.Fatalf("decode DSL fixture: %v", err)
	}
	return dsl
}

func assertFixture(t *testing.T, path string, value any) {
	t.Helper()
	want, err := MarshalCanonical(value)
	if err != nil {
		t.Fatalf("marshal fixture: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("%s drifted; run go run ./tools/apicontract generate", path)
	}
}

func contains(values []string, needle string) bool {
	return slices.Contains(values, needle)
}
