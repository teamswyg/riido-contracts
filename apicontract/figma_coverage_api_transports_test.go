package apicontract

import (
	"net/http"
	"strings"
	"testing"
)

func loadAIAgentClientGeneratedPathTransports(t *testing.T) map[string]figmaOpenAPITransport {
	t.Helper()
	openAPI := loadAIAgentClientOpenAPI(t)
	out := map[string]figmaOpenAPITransport{}
	for path, methods := range openAPI.Paths {
		for method, operation := range methods {
			if operation.RiidoClient == nil || strings.TrimSpace(operation.RiidoClient.GeneratedPath) == "" {
				continue
			}
			out[operation.RiidoClient.GeneratedPath] = figmaOpenAPITransport{
				Method:       strings.ToUpper(method),
				Path:         path,
				ContentTypes: figmaOpenAPIContentTypes(operation),
			}
		}
	}
	return out
}

func figmaOpenAPIContentTypes(operation OpenAPIOperation) map[string]bool {
	contentTypes := map[string]bool{}
	for _, response := range operation.Responses {
		for contentType := range response.Content {
			contentTypes[contentType] = true
		}
	}
	return contentTypes
}

func operationKindForOpenAPITransport(transport figmaOpenAPITransport) string {
	if transport.ContentTypes["text/event-stream"] {
		return "SSE Stream"
	}
	if transport.Method == http.MethodGet {
		return "Query"
	}
	return "Mutation"
}
