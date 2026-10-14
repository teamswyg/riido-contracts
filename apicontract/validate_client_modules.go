package apicontract

import (
	"errors"
	"fmt"
	"strings"
)

func validateClientModules(modules []ClientModule) (map[string]struct{}, error) {
	out := map[string]struct{}{}
	for _, module := range modules {
		name := strings.TrimSpace(module.Module)
		if name == "" {
			return nil, errors.New("apicontract: client module name is required")
		}
		if _, exists := out[name]; exists {
			return nil, fmt.Errorf("apicontract: duplicate client module %q", name)
		}
		out[name] = struct{}{}
		for _, namespace := range module.Namespaces {
			if len(namespace.Path) == 0 {
				return nil, fmt.Errorf("apicontract: client module %q has empty namespace path", name)
			}
			for _, segment := range namespace.Path {
				if strings.TrimSpace(segment) == "" {
					return nil, fmt.Errorf("apicontract: client module %q has blank namespace path segment", name)
				}
			}
		}
	}
	return out, nil
}
