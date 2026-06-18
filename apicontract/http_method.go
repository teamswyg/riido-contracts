package apicontract

import (
	"strings"
)

func methodAllowed(method string) bool {
	switch strings.ToUpper(method) {
	case "GET", "POST", "PATCH", "DELETE":
		return true
	default:
		return false
	}
}
