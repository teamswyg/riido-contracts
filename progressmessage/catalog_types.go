package progressmessage

type DSLDocument struct {
	SchemaVersion string              `json:"schema_version"`
	ContractID    string              `json:"contract_id"`
	Description   string              `json:"description,omitempty"`
	AppendOnly    bool                `json:"append_only"`
	MaxMessages   int                 `json:"max_messages"`
	MessageFiles  []string            `json:"message_files,omitempty"`
	Messages      []MessageDefinition `json:"messages,omitempty"`
}

type MessageDocument struct {
	SchemaVersion string            `json:"schema_version"`
	Message       MessageDefinition `json:"message"`
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
	MessageFiles        []string            `json:"message_files,omitempty"`
	Messages            []MessageDefinition `json:"messages,omitempty"`
}
