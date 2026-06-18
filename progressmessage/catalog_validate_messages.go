package progressmessage

import (
	"errors"
	"fmt"
	"strings"
)

func validateMessages(messages []MessageDefinition) error {
	if len(messages) == 0 {
		return errors.New("progressmessage: messages are required")
	}
	if len(messages) > MaxMessages {
		return fmt.Errorf("progressmessage: messages = %d, max %d", len(messages), MaxMessages)
	}
	codes := map[int]bool{}
	keys := map[string]bool{}
	for _, message := range messages {
		if err := validateMessageIdentity(message, codes, keys); err != nil {
			return err
		}
		if err := validateMessageLocales(message); err != nil {
			return err
		}
		if err := validateArgs(message); err != nil {
			return err
		}
	}
	return nil
}

func validateMessageIdentity(message MessageDefinition, codes map[int]bool, keys map[string]bool) error {
	if message.Code <= 0 {
		return errors.New("progressmessage: code must be positive")
	}
	if codes[message.Code] {
		return fmt.Errorf("progressmessage: duplicate code %d", message.Code)
	}
	codes[message.Code] = true
	if strings.TrimSpace(message.Key) == "" {
		return fmt.Errorf("progressmessage: message %d key is required", message.Code)
	}
	if keys[message.Key] {
		return fmt.Errorf("progressmessage: duplicate key %q", message.Key)
	}
	keys[message.Key] = true
	return validateUsage(message)
}
