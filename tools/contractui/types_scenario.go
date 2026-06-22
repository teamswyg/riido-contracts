package main

type scenario struct {
	Name  string `json:"name"`
	Given string `json:"given"`
	When  string `json:"when"`
	Then  string `json:"then"`
}
