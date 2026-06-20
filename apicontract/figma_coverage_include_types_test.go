package apicontract

type figmaCoverageEntryDocument struct {
	SchemaVersion string             `json:"schema_version"`
	Entry         figmaCoverageEntry `json:"entry"`
}

type figmaCoveragePageInventoryDocument struct {
	SchemaVersion string                      `json:"schema_version"`
	Inventory     figmaNonUITopLevelInventory `json:"inventory"`
}

type figmaCoverageToolLimitationDocument struct {
	SchemaVersion string                        `json:"schema_version"`
	Limitation    figmaSupportingToolLimitation `json:"limitation"`
}

type figmaCoverageAnnotationInventoryDocument struct {
	SchemaVersion string                           `json:"schema_version"`
	Inventory     figmaAPIGeneratedAnnotationGroup `json:"inventory"`
}

type figmaCoverageAnnotationDocument struct {
	SchemaVersion string                      `json:"schema_version"`
	Annotation    figmaAPIGeneratedAnnotation `json:"annotation"`
}

type figmaCoverageNodeDocument struct {
	SchemaVersion string            `json:"schema_version"`
	Node          figmaCoverageNode `json:"node"`
}
