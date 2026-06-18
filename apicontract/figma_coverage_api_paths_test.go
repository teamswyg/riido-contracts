package apicontract

import (
	"strings"
	"testing"
)

func canonicalPathFromFigmaFacade(path string) string {
	out := strings.TrimPrefix(path, "riido.")
	out = strings.TrimPrefix(out, "v2.")
	return out
}

func loadAIAgentClientGeneratedPaths(t *testing.T) map[string]string {
	t.Helper()
	openAPI := loadAIAgentClientOpenAPI(t)
	out := map[string]string{}
	for path, methods := range openAPI.Paths {
		for method, operation := range methods {
			if operation.RiidoClient == nil || strings.TrimSpace(operation.RiidoClient.GeneratedPath) == "" {
				continue
			}
			out[operation.RiidoClient.GeneratedPath] = strings.ToUpper(method) + " " + path
		}
	}
	return out
}

func docMentionsGeneratedPath(docText, generatedPath string) bool {
	if strings.Contains(docText, generatedPath) {
		return true
	}
	lastDot := strings.LastIndex(generatedPath, ".")
	if lastDot < 0 {
		return false
	}
	return strings.Contains(docText, generatedPath[:lastDot]+".*")
}
