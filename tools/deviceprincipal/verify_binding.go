package main

import "fmt"

var requiredBindingFields = []string{
	"agent_id",
	"daemon_id",
	"device_id",
	"runtime_id",
	"runtime_provider",
}

func verifyBindingFields(model model) error {
	for _, want := range requiredBindingFields {
		if !contains(model.BindingFields, want) {
			return fmt.Errorf("AgentRuntimeBinding missing json field %q", want)
		}
	}
	return nil
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
