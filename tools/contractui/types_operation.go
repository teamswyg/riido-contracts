package main

type operation struct {
	OperationID string      `json:"operation_id"`
	Kind        string      `json:"kind"`
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Resource    string      `json:"resource"`
	Action      string      `json:"action"`
	Summary     string      `json:"summary"`
	Client      *clientMeta `json:"client,omitempty"`
	Auth        auth        `json:"auth"`
	RBACPolicy  string      `json:"rbac_policy,omitempty"`
	Request     *messageRef `json:"request,omitempty"`
	Response    responseRef `json:"response"`
	ScenarioIDs []string    `json:"scenario_ids,omitempty"`
	Scenarios   []scenario  `json:"scenarios,omitempty"`
	PathParams  []string    `json:"path_params,omitempty"`
}

type clientMeta struct {
	Module        string   `json:"module"`
	FacadePath    []string `json:"facade_path,omitempty"`
	GeneratedPath string   `json:"generated_path,omitempty"`
	CacheTag      string   `json:"cache_tag,omitempty"`
	Invalidates   []string `json:"invalidates,omitempty"`
}

type auth struct {
	Scheme string   `json:"scheme"`
	Header string   `json:"header,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type messageRef struct {
	Ref      string `json:"ref"`
	Required bool   `json:"required,omitempty"`
}

type responseRef struct {
	Status      int    `json:"status"`
	Ref         string `json:"ref"`
	ContentType string `json:"content_type,omitempty"`
}
