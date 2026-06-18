package progressmessage

import (
	"fmt"
	"strings"
)

func validateArgs(message MessageDefinition) error {
	args := map[string]bool{}
	for _, arg := range message.Args {
		if err := validateArg(message, arg, args); err != nil {
			return err
		}
	}
	for locale, template := range message.Locales {
		matches := placeholderPattern.FindAllStringSubmatch(template, -1)
		for _, match := range matches {
			if !args[match[1]] {
				return fmt.Errorf("progressmessage: message %d locale %s references unknown arg %q", message.Code, locale, match[1])
			}
		}
	}
	return nil
}

func validateArg(message MessageDefinition, arg MessageArg, args map[string]bool) error {
	if strings.TrimSpace(arg.Name) == "" {
		return fmt.Errorf("progressmessage: message %d has blank arg name", message.Code)
	}
	if arg.Type != "string" && arg.Type != "int" {
		return fmt.Errorf("progressmessage: message %d arg %q has invalid type %q", message.Code, arg.Name, arg.Type)
	}
	if args[arg.Name] {
		return fmt.Errorf("progressmessage: message %d has duplicate arg %q", message.Code, arg.Name)
	}
	args[arg.Name] = true
	return nil
}
