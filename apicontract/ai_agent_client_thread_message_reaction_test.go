package apicontract

import (
	"regexp"
	"strings"
	"testing"
)

const (
	threadMessageReactionsPath      = "/v2/client/workspaces/{workspace_id}/ai-agent/tasks/{task_id}/threads/{thread_id}/messages/{message_id}/reactions"
	threadMessageReactionDeletePath = threadMessageReactionsPath + "/{emoji}"
	threadMessageReactionPolicy     = "task_thread_message_reaction.v1"
)

func TestAIAgentClientThreadMessageReactionOperations(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	cases := []struct {
		path, method, operationID, generatedPath string
		hasBody                                  bool
		pathParams                               []string
	}{
		{
			path:          threadMessageReactionsPath,
			method:        "post",
			operationID:   "addAIAgentTaskThreadMessageReactionV2",
			generatedPath: "v2.aiAgent.tasks.threadMessages.reactions.add",
			hasBody:       true,
			pathParams:    []string{"workspace_id", "task_id", "thread_id", "message_id"},
		},
		{
			path:          threadMessageReactionDeletePath,
			method:        "delete",
			operationID:   "deleteAIAgentTaskThreadMessageReactionV2",
			generatedPath: "v2.aiAgent.tasks.threadMessages.reactions.delete",
			pathParams:    []string{"workspace_id", "task_id", "thread_id", "message_id", "emoji"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.operationID, func(t *testing.T) {
			op := fixture.openAPI.Paths[tc.path][tc.method]
			if op.OperationID != tc.operationID || op.RiidoKind != "command" {
				t.Fatalf("operation = %#v", op)
			}
			if op.RiidoClient == nil || op.RiidoClient.GeneratedPath != tc.generatedPath {
				t.Fatalf("client metadata = %#v", op.RiidoClient)
			}
			// 사람 반응은 thread message 작성과 같은 권한으로만 남길 수 있다.
			if strings.Join(op.RiidoScopes, ",") != "ai-agent:write,task:{task_id}:comment" || op.RiidoRBAC != threadMessageReactionPolicy {
				t.Fatalf("auth = %#v / %q", op.RiidoScopes, op.RiidoRBAC)
			}
			if (op.RequestBody != nil) != tc.hasBody {
				t.Fatalf("request body presence = %v, want %v", op.RequestBody != nil, tc.hasBody)
			}
			verifyPathParameters(t, op.Parameters, tc.pathParams)
			if _, ok := op.Responses["200"]; !ok {
				t.Fatalf("responses = %#v", op.Responses)
			}
		})
	}
}

func verifyPathParameters(t *testing.T, params []OpenAPIParameter, want []string) {
	t.Helper()
	got := map[string]bool{}
	for _, param := range params {
		if param.In == "path" && param.Required {
			got[param.Name] = true
		}
	}
	for _, name := range want {
		if !got[name] {
			t.Fatalf("missing required path parameter %q in %#v", name, params)
		}
	}
}

func TestAIAgentClientThreadMessageReactionSchemas(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	schemas := fixture.openAPI.Components.Schemas

	kinds, ok := schemas["AIAgentMessageReactionActorKind"]["enum"].([]string)
	if !ok || strings.Join(kinds, ",") != "member,agent" {
		t.Fatalf("actor kind enum = %#v", schemas["AIAgentMessageReactionActorKind"]["enum"])
	}

	reaction := openAPISchemaProperties(t, fixture.openAPI, "AIAgentMessageReaction")
	for _, field := range []string{"emoji", "count", "actors"} {
		if _, ok := reaction[field]; !ok {
			t.Fatalf("AIAgentMessageReaction %s missing", field)
		}
	}
	emoji, _ := reaction["emoji"].(map[string]any)
	pattern, _ := emoji["pattern"].(string)
	valid := regexp.MustCompile(pattern)
	// monolith 댓글 reaction과 같은 정규화 값만 받는다.
	for _, value := range []string{"1f440", "2705", "274c", "1f468-200d-1f4bb"} {
		if !valid.MatchString(value) {
			t.Fatalf("emoji pattern %q must accept %q", pattern, value)
		}
	}
	for _, value := range []string{"👀", "1F440", ":eyes:", "", "1f440 "} {
		if valid.MatchString(value) {
			t.Fatalf("emoji pattern %q must reject %q", pattern, value)
		}
	}

	history := openAPISchemaProperties(t, fixture.openAPI, "AIAgentTaskThreadHistoryMessage")
	if _, ok := history["reactions"]; !ok {
		t.Fatal("v3 history message must expose reactions")
	}
	for _, required := range openAPISchemaRequired(t, schemas["AIAgentTaskThreadHistoryMessage"], "AIAgentTaskThreadHistoryMessage") {
		if required == "reactions" {
			t.Fatal("reactions must stay optional for older servers and progress messages")
		}
	}

	request := openAPISchemaProperties(t, fixture.openAPI, "AddAIAgentTaskThreadMessageReactionRequest")
	if len(request) != 1 {
		t.Fatalf("add reaction request must only carry emoji, never an actor: %#v", request)
	}
}

func TestAIAgentClientThreadMessageReactionPolicy(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	var rules []string
	for _, policy := range fixture.ir.Policies {
		if policy.PolicyID == threadMessageReactionPolicy {
			rules = policy.Rules
		}
	}
	if len(rules) == 0 {
		t.Fatalf("%s policy missing", threadMessageReactionPolicy)
	}
	joined := strings.Join(rules, "\n")
	// RIID-5904에서 정한 상태별 반응과 알림 규칙이 빠지지 않게 고정한다.
	for _, required := range []string{
		"progress messages are never reactable",
		"clients cannot create, delete, or impersonate agent reactions",
		"completed replaces 1f440 with 2705",
		"failed replaces 1f440 with 274c",
		"cancelling and cancelled remove 1f440",
		"needs_input and waiting_for_user keep 1f440",
		"never fails, blocks, or retries the agent run",
		"a member reaction notifies the message author as an unread Riido notification",
		"an agent reaction never creates a Riido notification",
		"agent task desktop push of the same run lifecycle",
		"no notification is created when the author is an agent",
		"Slack and Discord reaction sync does not apply to agent reactions",
		"has a matching entry in agent_snapshots",
	} {
		if !strings.Contains(joined, required) {
			t.Fatalf("%s must keep rule containing %q", threadMessageReactionPolicy, required)
		}
	}
}
