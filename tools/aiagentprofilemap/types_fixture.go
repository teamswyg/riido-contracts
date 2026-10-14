package main

type contractFixture struct {
	Operations []operation  `json:"operations"`
	Schemas    []schema     `json:"schemas"`
	Policies   []policySpec `json:"policies"`
}

type operation struct {
	OperationID string       `json:"operation_id"`
	Kind        string       `json:"kind"`
	Method      string       `json:"method"`
	Path        string       `json:"path"`
	Client      clientSpec   `json:"client"`
	RBACPolicy  string       `json:"rbac_policy"`
	Response    *responseRef `json:"response"`
	Scenarios   []scenario   `json:"scenarios"`
}

type clientSpec struct {
	GeneratedPath string `json:"generated_path"`
	CacheTag      string `json:"cache_tag"`
}

type responseRef struct {
	Status int    `json:"status"`
	Ref    string `json:"ref"`
}

type scenario struct {
	Name string `json:"name"`
}
