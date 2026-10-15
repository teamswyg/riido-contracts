package apicontract

import (
	"fmt"
	"strings"
)

type QueryParameter struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required,omitempty"`
	Minimum  *int   `json:"minimum,omitempty"`
	Default  *int   `json:"default,omitempty"`
}

func validateQueryParameters(params []QueryParameter) error {
	seen := map[string]bool{}
	for _, p := range params {
		if p.Name == "" || strings.TrimSpace(p.Name) != p.Name || seen[p.Name] {
			return fmt.Errorf("apicontract: invalid/duplicate query parameter %q", p.Name)
		}
		seen[p.Name] = true
		if p.Type != "string" && p.Type != "integer" && p.Type != "boolean" {
			return fmt.Errorf("apicontract: unsupported query parameter type %q", p.Type)
		}
		if (p.Minimum != nil || p.Default != nil) && p.Type != "integer" {
			return fmt.Errorf("apicontract: integer query constraints require integer type")
		}
		if p.Minimum != nil && p.Default != nil && *p.Default < *p.Minimum {
			return fmt.Errorf("apicontract: query default is below minimum")
		}
	}
	return nil
}

func operationParameters(path string, query []QueryParameter) []OpenAPIParameter {
	out := pathParameters(path)
	for _, p := range query {
		schema := map[string]any{"type": p.Type}
		if p.Minimum != nil {
			schema["minimum"] = *p.Minimum
		}
		if p.Default != nil {
			schema["default"] = *p.Default
		}
		out = append(out, OpenAPIParameter{Name: p.Name, In: "query", Required: p.Required, Schema: schema})
	}
	return out
}
