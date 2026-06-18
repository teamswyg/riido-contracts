package progressmessage

import (
	"fmt"
	"strings"
)

func validateUsage(message MessageDefinition) error {
	switch message.Usage {
	case UsageRequired, UsageActive, UsageReserved:
		return nil
	default:
		return fmt.Errorf("progressmessage: message %d has invalid usage %q", message.Code, message.Usage)
	}
}

func validateMessageLocales(message MessageDefinition) error {
	if strings.TrimSpace(message.Category) == "" {
		return fmt.Errorf("progressmessage: message %d category is required", message.Code)
	}
	if message.Locales[DefaultLocale] == "" {
		return fmt.Errorf("progressmessage: message %d must define %s locale", message.Code, DefaultLocale)
	}
	if message.Locales["en"] == "" {
		return fmt.Errorf("progressmessage: message %d must define en locale", message.Code)
	}
	return nil
}
