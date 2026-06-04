package progressmessage

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	DSLSchemaVersion = "riido-progress-message-dsl.v1"
	IRSchemaVersion  = "riido-progress-message-ir.v1"
	ContractID       = "ai-agent-progress-message-catalog.v1"
	MaxMessages      = 15
	DefaultLocale    = "ko"

	UsageRequired = "required"
	UsageActive   = "active"
	UsageReserved = "reserved"
)

var placeholderPattern = regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)\}\}`)

//go:embed catalog.ir.riido.json
var embeddedIR []byte

type DSLDocument struct {
	SchemaVersion string              `json:"schema_version"`
	ContractID    string              `json:"contract_id"`
	Description   string              `json:"description,omitempty"`
	AppendOnly    bool                `json:"append_only"`
	MaxMessages   int                 `json:"max_messages"`
	Messages      []MessageDefinition `json:"messages"`
}

type MessageDefinition struct {
	Code        int               `json:"code"`
	Key         string            `json:"key"`
	Usage       string            `json:"usage"`
	Category    string            `json:"category"`
	Description string            `json:"description,omitempty"`
	Args        []MessageArg      `json:"args,omitempty"`
	Locales     map[string]string `json:"locales"`
}

type MessageArg struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required,omitempty"`
}

type IRDocument struct {
	SchemaVersion       string              `json:"schema_version"`
	ContractID          string              `json:"contract_id"`
	SourceSchemaVersion string              `json:"source_schema_version"`
	AppendOnly          bool                `json:"append_only"`
	MaxMessages         int                 `json:"max_messages"`
	Messages            []MessageDefinition `json:"messages"`
}

func Catalog() (IRDocument, error) {
	var ir IRDocument
	dec := json.NewDecoder(bytes.NewReader(embeddedIR))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&ir); err != nil {
		return IRDocument{}, fmt.Errorf("progressmessage: decode embedded IR: %w", err)
	}
	if err := ValidateIR(ir); err != nil {
		return IRDocument{}, err
	}
	return ir, nil
}

func GenerateIR(dsl DSLDocument) (IRDocument, error) {
	if err := ValidateDSL(dsl); err != nil {
		return IRDocument{}, err
	}
	messages := append([]MessageDefinition(nil), dsl.Messages...)
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Code < messages[j].Code
	})
	return IRDocument{
		SchemaVersion:       IRSchemaVersion,
		ContractID:          dsl.ContractID,
		SourceSchemaVersion: dsl.SchemaVersion,
		AppendOnly:          dsl.AppendOnly,
		MaxMessages:         dsl.MaxMessages,
		Messages:            messages,
	}, nil
}

func ValidateDSL(dsl DSLDocument) error {
	if dsl.SchemaVersion != DSLSchemaVersion {
		return fmt.Errorf("progressmessage: unsupported DSL schema_version %q", dsl.SchemaVersion)
	}
	if dsl.ContractID != ContractID {
		return fmt.Errorf("progressmessage: contract_id = %q, want %q", dsl.ContractID, ContractID)
	}
	if !dsl.AppendOnly {
		return errors.New("progressmessage: append_only must be true")
	}
	if dsl.MaxMessages != MaxMessages {
		return fmt.Errorf("progressmessage: max_messages = %d, want %d", dsl.MaxMessages, MaxMessages)
	}
	return validateMessages(dsl.Messages)
}

func ValidateIR(ir IRDocument) error {
	if ir.SchemaVersion != IRSchemaVersion {
		return fmt.Errorf("progressmessage: unsupported IR schema_version %q", ir.SchemaVersion)
	}
	if ir.SourceSchemaVersion != DSLSchemaVersion {
		return fmt.Errorf("progressmessage: source_schema_version = %q, want %q", ir.SourceSchemaVersion, DSLSchemaVersion)
	}
	if ir.ContractID != ContractID {
		return fmt.Errorf("progressmessage: contract_id = %q, want %q", ir.ContractID, ContractID)
	}
	if !ir.AppendOnly {
		return errors.New("progressmessage: append_only must be true")
	}
	if ir.MaxMessages != MaxMessages {
		return fmt.Errorf("progressmessage: max_messages = %d, want %d", ir.MaxMessages, MaxMessages)
	}
	if !messagesSorted(ir.Messages) {
		return errors.New("progressmessage: IR messages must be sorted by code")
	}
	return validateMessages(ir.Messages)
}

func Render(code int, args map[string]string, locale string) (string, bool) {
	ir, err := Catalog()
	if err != nil {
		return "", false
	}
	for _, message := range ir.Messages {
		if message.Code != code {
			continue
		}
		template := message.Locales[strings.TrimSpace(locale)]
		if template == "" {
			template = message.Locales[DefaultLocale]
		}
		if template == "" {
			return "", false
		}
		return renderTemplate(template, args), true
	}
	return "", false
}

func MarshalCanonical(value any) ([]byte, error) {
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(out, '\n'), nil
}

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
		switch message.Usage {
		case UsageRequired, UsageActive, UsageReserved:
		default:
			return fmt.Errorf("progressmessage: message %d has invalid usage %q", message.Code, message.Usage)
		}
		if strings.TrimSpace(message.Category) == "" {
			return fmt.Errorf("progressmessage: message %d category is required", message.Code)
		}
		if message.Locales[DefaultLocale] == "" {
			return fmt.Errorf("progressmessage: message %d must define %s locale", message.Code, DefaultLocale)
		}
		if message.Locales["en"] == "" {
			return fmt.Errorf("progressmessage: message %d must define en locale", message.Code)
		}
		if err := validateArgs(message); err != nil {
			return err
		}
	}
	return nil
}

func validateArgs(message MessageDefinition) error {
	args := map[string]bool{}
	for _, arg := range message.Args {
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

func messagesSorted(messages []MessageDefinition) bool {
	for i := 1; i < len(messages); i++ {
		if messages[i-1].Code > messages[i].Code {
			return false
		}
	}
	return true
}

func renderTemplate(template string, args map[string]string) string {
	return placeholderPattern.ReplaceAllStringFunc(template, func(match string) string {
		parts := placeholderPattern.FindStringSubmatch(match)
		if len(parts) != 2 {
			return match
		}
		value := strings.TrimSpace(args[parts[1]])
		if value == "" {
			return "not provided"
		}
		return value
	})
}
