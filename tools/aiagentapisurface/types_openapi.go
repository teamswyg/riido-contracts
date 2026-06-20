package main

type openAPIDoc struct {
	Paths map[string]map[string]openAPIOperation `json:"paths"`
}

type openAPIOperation struct {
	OperationID string `json:"operationId"`
}
