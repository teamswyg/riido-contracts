package apicontract

import "testing"

func verifyFigmaAPIGeneratedAnnotationInventory(
	t *testing.T,
	inventory []figmaAPIGeneratedAnnotationGroup,
	docText string,
	openAPIGeneratedPaths map[string]string,
	openAPITransports map[string]figmaOpenAPITransport,
	registered map[string]string,
	entries map[string]figmaCoverageEntry,
) {
	t.Helper()
	if got, want := len(inventory), 20; got != want {
		t.Fatalf("api_generated_annotation_inventory = %d, want %d", got, want)
	}
	scope := figmaAPIGeneratedInventoryScope{
		docText:               docText,
		openAPIGeneratedPaths: openAPIGeneratedPaths,
		openAPITransports:     openAPITransports,
		registered:            registered,
		entries:               entries,
		seenPath:              map[string]bool{},
	}
	for _, group := range inventory {
		scope.verifyGroup(t, group)
	}
	if got, want := scope.totalAnnotations, 90; got != want {
		t.Fatalf("API Generated annotation inventory node annotations = %d, want %d", got, want)
	}
}
