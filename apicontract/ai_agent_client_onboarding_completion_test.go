package apicontract

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

const onboardingCompletionDSLPath = "fixtures/control-plane-ai-agent-client.dsl.riido.json"

// Approved product copy for the onboarding first-task completion comment.
// The policy rules must carry each string verbatim (after the rule label).
var approvedOnboardingCompletionCopy = []string{
	"단계 4개의 기간과 날짜, 산출물을 이 작업의 문서로 정리해 뒀어요. 팀원을 초대하면 단계마다 담당자를 정해 보세요.",
	"Laid out the 4 phases' durations, dates, and deliverables in this task's document. Once you invite teammates, pick an owner for each phase.",
	"질문지를 이 작업의 문서로 남겼어요. \"~라면 쓰시겠어요?\"처럼 답을 유도하는 질문은 뺐어요.",
	"Saved the questions as this task's document. Leading questions like \"Would you use it if…?\" were left out.",
	"출시 전·당일·후에 확인할 항목 9개를 이 작업의 문서로 정리해 뒀어요. 출시일은 오늘부터 7일 뒤로 잡았어요.",
	"Laid out 9 items to check before, on, and after launch in this task's document. The launch date is set 7 days from today.",
}

// Retired copy claimed the template created subtasks. The first-task bundle now
// records its results in the task document only, so these strings must not return.
var retiredOnboardingCompletionCopy = []string{
	"단계 4개를 이 작업의 하위 작업으로 추가하고 마감일을 넣어 뒀어요.",
	"Added the 4 phases as subtasks of this task with due dates.",
	"항목 9개를 이 작업의 하위 작업으로 추가하고 D-3, 당일, D+7 마감일을 넣어 뒀어요.",
	"Added the 9 items as subtasks of this task with D-3, launch-day, and D+7 due dates.",
}

func TestAIAgentClientOnboardingCompletionContract(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	verifyOnboardingCompletionOperation(t, f.openAPI)
	verifyOnboardingCompletionSchemas(t, f.openAPI)
	verifyOnboardingCompletionPolicy(t)
}

func verifyOnboardingCompletionOperation(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	op := openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/tasks/{task_id}/onboarding-completion"]["post"]
	if op.OperationID != "completeAIAgentOnboardingTaskV2" {
		t.Fatalf("onboarding completion operation = %#v", op)
	}
	if op.RiidoClient == nil || op.RiidoClient.GeneratedPath != "v2.aiAgent.tasks.onboardingCompletion.create" {
		t.Fatalf("onboarding completion client metadata = %#v", op.RiidoClient)
	}
	if op.RequestBody == nil {
		t.Fatalf("onboarding completion request body missing")
	}
	if !contains(op.RiidoClient.Invalidates, "v2.aiAgent.tasks.threads") {
		t.Fatalf("onboarding completion must invalidate task threads: %#v", op.RiidoClient.Invalidates)
	}
}

func verifyOnboardingCompletionSchemas(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	name := "CompleteAIAgentOnboardingTaskRequest"
	props := openAPISchemaProperties(t, openAPI, name)
	for _, field := range []string{"agent_id", "template_key", "request_id", "locale"} {
		if _, ok := props[field]; !ok {
			t.Fatalf("%s %s property missing", name, field)
		}
	}
	required := openAPISchemaRequired(t, openAPI.Components.Schemas[name], name)
	if !contains(required, "agent_id") || !contains(required, "template_key") || contains(required, "locale") {
		t.Fatalf("%s required = %#v", name, required)
	}
	template, ok := openAPI.Components.Schemas["AIAgentOnboardingTaskTemplateKey"]["enum"].([]string)
	if !ok || len(template) != 3 || template[0] != "schedule" || template[1] != "interview" || template[2] != "launch" {
		t.Fatalf("AIAgentOnboardingTaskTemplateKey enum = %#v", openAPI.Components.Schemas["AIAgentOnboardingTaskTemplateKey"]["enum"])
	}
	locale, ok := openAPI.Components.Schemas["AIAgentOnboardingLocale"]["enum"].([]string)
	if !ok || len(locale) != 2 || locale[0] != "ko" || locale[1] != "en" {
		t.Fatalf("AIAgentOnboardingLocale enum = %#v", openAPI.Components.Schemas["AIAgentOnboardingLocale"]["enum"])
	}
}

func verifyOnboardingCompletionPolicy(t *testing.T) {
	t.Helper()
	raw, err := os.ReadFile(onboardingCompletionDSLPath)
	if err != nil {
		t.Fatalf("read DSL: %v", err)
	}
	var dsl struct {
		Policies []struct {
			PolicyID string   `json:"policy_id"`
			Rules    []string `json:"rules"`
		} `json:"policies"`
	}
	if err := json.Unmarshal(raw, &dsl); err != nil {
		t.Fatalf("decode DSL: %v", err)
	}
	var rules []string
	for _, policy := range dsl.Policies {
		if policy.PolicyID == "ai_agent_onboarding_completion.v1" {
			rules = policy.Rules
		}
	}
	if len(rules) == 0 {
		t.Fatalf("ai_agent_onboarding_completion.v1 policy missing")
	}
	for _, copyText := range approvedOnboardingCompletionCopy {
		found := false
		for _, rule := range rules {
			if strings.HasSuffix(rule, copyText) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("approved onboarding completion copy missing from policy: %q", copyText)
		}
	}
	verifyRetiredOnboardingCompletionCopy(t, rules)
}

func verifyRetiredOnboardingCompletionCopy(t *testing.T, rules []string) {
	t.Helper()
	for _, retired := range retiredOnboardingCompletionCopy {
		for _, rule := range rules {
			if strings.Contains(rule, retired) {
				t.Fatalf("retired onboarding completion copy still in policy: %q", retired)
			}
		}
	}
}
