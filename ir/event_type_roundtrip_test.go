package ir

import "testing"

func TestEventTypeGeneratedCodeRoundTrip(t *testing.T) {
	if got := EventTaskQueued.Code(); got != EventTypeCodeTaskQueued {
		t.Fatalf("EventTaskQueued.Code() = %v, want %v", got, EventTypeCodeTaskQueued)
	}
	if got := EventTypeCodeTaskQueued.StringValue(); got != EventTypeStringTaskQueued {
		t.Fatalf("EventTypeCodeTaskQueued.StringValue() = %q, want %q", got, EventTypeStringTaskQueued)
	}
	if got := EventTypeCodeTaskQueued.EventType(); got != EventTaskQueued {
		t.Fatalf("EventTypeCodeTaskQueued.EventType() = %q, want %q", got, EventTaskQueued)
	}
	if ParseEventTypeCode("not-an-event") != EventTypeCodeUnknown {
		t.Fatal("unknown event type must parse to EventTypeCodeUnknown")
	}
	if !EventTypeCodeTaskQueued.IsTransition() {
		t.Fatal("generated EventTypeCode transition classification drifted")
	}
}
