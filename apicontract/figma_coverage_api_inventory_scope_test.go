package apicontract

type figmaAPIGeneratedInventoryScope struct {
	docText               string
	openAPIGeneratedPaths map[string]string
	openAPITransports     map[string]figmaOpenAPITransport
	registered            map[string]string
	entries               map[string]figmaCoverageEntry
	seenPath              map[string]bool
	totalAnnotations      int
}
