package apicontract

type IROperation struct {
	OperationID string      `json:"operation_id"`
	Kind        string      `json:"kind"`
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Resource    string      `json:"resource"`
	Action      string      `json:"action"`
	Client      *ClientMeta `json:"client,omitempty"`
	Summary     string      `json:"summary"`
	Auth        Auth        `json:"auth"`
	RBACPolicy  string      `json:"rbac_policy,omitempty"`
	Request     *MessageRef `json:"request,omitempty"`
	Response    ResponseRef `json:"response"`
	ScenarioIDs []string    `json:"scenario_ids,omitempty"`
}

type IRScenario struct {
	ScenarioID  string `json:"scenario_id"`
	OperationID string `json:"operation_id"`
	Name        string `json:"name"`
	Given       string `json:"given"`
	When        string `json:"when"`
	Then        string `json:"then"`
}
