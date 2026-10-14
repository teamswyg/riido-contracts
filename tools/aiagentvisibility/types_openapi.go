package main

type openAPIDoc struct {
	Paths map[string]map[string]openAPIOperation `json:"paths"`
}

type openAPIOperation struct {
	OperationID string                     `json:"operationId"`
	Responses   map[string]openAPIResponse `json:"responses"`
}

type openAPIResponse struct {
	Content map[string]openAPIMedia `json:"content"`
}

type openAPIMedia struct {
	Schema openAPISchema `json:"schema"`
}

type openAPISchema struct {
	Ref string `json:"$ref"`
}
