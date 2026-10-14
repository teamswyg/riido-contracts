package apicontract

import (
	"fmt"
	"strings"
)

func validateAuth(operationID string, auth Auth) error {
	switch auth.Scheme {
	case "bearer":
		if strings.TrimSpace(auth.Header) != "" {
			return fmt.Errorf("apicontract: operation %q bearer auth must not set header", operationID)
		}
	case "apiKey":
		if strings.TrimSpace(auth.Header) == "" {
			return fmt.Errorf("apicontract: operation %q apiKey auth must set header", operationID)
		}
	default:
		return fmt.Errorf("apicontract: operation %q has unsupported auth scheme %q", operationID, auth.Scheme)
	}
	return nil
}
