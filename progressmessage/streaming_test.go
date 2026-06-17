package progressmessage

import "testing"

func TestAssistantPartialStreamingContractIsOutsideCatalog(t *testing.T) {
	if AssistantPartialKey != "assistant.partial" {
		t.Fatalf("AssistantPartialKey = %q", AssistantPartialKey)
	}
	if AssistantPartialCode <= MaxMessages {
		t.Fatalf("AssistantPartialCode = %d, want outside catalog max %d", AssistantPartialCode, MaxMessages)
	}
	if rendered, ok := Render(AssistantPartialCode, nil, DefaultLocale); ok || rendered != "" {
		t.Fatalf("assistant partial must not render through catalog: ok=%v rendered=%q", ok, rendered)
	}
}
