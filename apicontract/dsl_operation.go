package apicontract

type DSLOperation struct {
	OperationID string      `json:"operation_id"`
	Kind        string      `json:"kind"`
	Summary     string      `json:"summary"`
	Description string      `json:"description,omitempty"`
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Resource    string      `json:"resource"`
	Action      string      `json:"action"`
	Client      *ClientMeta `json:"client,omitempty"`
	Auth        Auth        `json:"auth"`
	RBACPolicy  string      `json:"rbac_policy,omitempty"`
	Request     *MessageRef `json:"request,omitempty"`
	Response    ResponseRef `json:"response"`
	Scenarios   []Scenario  `json:"scenarios,omitempty"`
}

type Auth struct {
	Scheme string   `json:"scheme"`
	Header string   `json:"header,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type MessageRef struct {
	Ref      string `json:"ref"`
	Required bool   `json:"required,omitempty"`
}

type ResponseRef struct {
	Status      int    `json:"status"`
	Ref         string `json:"ref"`
	ContentType string `json:"content_type,omitempty"`
}

type Scenario struct {
	Name  string `json:"name"`
	Given string `json:"given"`
	When  string `json:"when"`
	Then  string `json:"then"`
}
