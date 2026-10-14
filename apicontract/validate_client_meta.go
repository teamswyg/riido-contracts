package apicontract

import (
	"fmt"
	"strings"
)

func validateClientMeta(operationID, method string, meta ClientMeta, modules map[string]struct{}) error {
	if _, ok := modules[meta.Module]; !ok {
		return fmt.Errorf("apicontract: operation %q references unknown client module %q", operationID, meta.Module)
	}
	if len(meta.FacadePath) == 0 {
		return fmt.Errorf("apicontract: operation %q missing client facade_path", operationID)
	}
	for _, segment := range meta.FacadePath {
		if strings.TrimSpace(segment) == "" {
			return fmt.Errorf("apicontract: operation %q has blank client facade_path segment", operationID)
		}
	}
	if strings.EqualFold(method, "GET") && strings.TrimSpace(meta.CacheTag) == "" {
		return fmt.Errorf("apicontract: operation %q missing client cache_tag", operationID)
	}
	if generatedPath := strings.TrimSpace(meta.GeneratedPath); generatedPath != "" {
		want := generatedClientPath(meta)
		if generatedPath != want {
			return fmt.Errorf("apicontract: operation %q has client generated_path %q, want %q", operationID, generatedPath, want)
		}
	}
	return nil
}
