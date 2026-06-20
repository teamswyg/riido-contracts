package main

type apiFixture struct {
	ContractID string        `json:"contract_id"`
	Operations []operation   `json:"operations"`
	SumTypes   []sumTypeSpec `json:"sum_types"`
}

type operation struct {
	OperationID string `json:"operation_id"`
	Method      string `json:"method"`
	Path        string `json:"path"`
}

type sumTypeSpec struct {
	Name     string        `json:"name"`
	Variants []variantSpec `json:"variants"`
}

type variantSpec struct {
	Kind   string `json:"kind"`
	Schema string `json:"schema"`
}
