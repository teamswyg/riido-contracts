package main

type contractFixture struct {
	Operations []operation `json:"operations"`
	Schemas    []schema    `json:"schemas"`
}

type operation struct {
	OperationID string       `json:"operation_id"`
	Kind        string       `json:"kind"`
	Method      string       `json:"method"`
	Path        string       `json:"path"`
	Request     *requestRef  `json:"request"`
	Response    *responseRef `json:"response"`
	Scenarios   []scenario   `json:"scenarios"`
}

type requestRef struct {
	Ref      string `json:"ref"`
	Required bool   `json:"required"`
}

type responseRef struct {
	Status int    `json:"status"`
	Ref    string `json:"ref"`
}

type scenario struct {
	Name string `json:"name"`
}
