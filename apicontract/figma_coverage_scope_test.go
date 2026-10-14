package apicontract

type figmaCoverageTestScope struct {
	manifest              figmaCoverageManifest
	docPath               string
	docText               string
	pages                 map[string]figmaCoveragePage
	expected              map[string]figmaCoverageNode
	registered            map[string]string
	nonUIInventory        map[string]map[string]figmaCoverageNode
	entryByNodeID         map[string]figmaCoverageEntry
	openAPIGeneratedPaths map[string]string
	openAPITransports     map[string]figmaOpenAPITransport
	seen                  map[string]bool
	nonUISeen             map[string]bool
}
