package apicontract

import (
	"slices"
)

func stringSliceContains(items []string, want string) bool {
	return slices.Contains(items, want)
}
