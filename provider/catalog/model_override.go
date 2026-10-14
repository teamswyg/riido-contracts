package catalog

import "strings"

func ModelOverride(provider, modelID string) string {
	modelID = strings.TrimSpace(modelID)
	if modelID == "" || modelID == DefaultModelID(provider) {
		return ""
	}
	return modelID
}
