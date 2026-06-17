package progressmessage

const (
	// AssistantPartialCode is a sentinel progress code outside the canonical
	// template catalog range. Receivers must keep the body text verbatim.
	AssistantPartialCode = 9001

	// AssistantPartialKey marks an evolving assistant body update. Consumers
	// should replace the previous partial body for a thread instead of treating
	// every update as a distinct durable progress step.
	AssistantPartialKey = "assistant.partial"
)
