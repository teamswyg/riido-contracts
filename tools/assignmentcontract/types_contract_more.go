package main

type approvalContract struct {
	Owner                 string        `json:"owner"`
	TimeoutTerminalStatus string        `json:"timeout_terminal_status"`
	Statuses              []statusValue `json:"statuses"`
	Decisions             []namedValue  `json:"decisions"`
}

type statusValue struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Terminal bool   `json:"terminal"`
}

type payloadField struct {
	Name               string `json:"name"`
	Source             string `json:"source"`
	MaxLength          int    `json:"max_length"`
	UserInputMaxLength int    `json:"user_input_max_length,omitempty"`
	Required           bool   `json:"required"`
	Snapshot           string `json:"snapshot"`
	Consumer           string `json:"consumer"`
}
