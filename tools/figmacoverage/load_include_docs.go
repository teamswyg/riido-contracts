package main

import (
	"fmt"
	"path/filepath"
)

func loadCoverageEntryDocument(path string) (coverageEntry, error) {
	var doc coverageEntryDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return coverageEntry{}, err
	}
	if doc.SchemaVersion != entrySchemaVersion {
		return coverageEntry{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.Entry, nil
}

func loadPageInventoryDocument(path string) (pageInventory, error) {
	var doc pageInventoryDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return pageInventory{}, err
	}
	if doc.SchemaVersion != pageInventorySchemaVersion {
		return pageInventory{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	for _, file := range doc.Inventory.NodeFiles {
		node, err := loadNodeDocument(filepath.Join(filepath.Dir(path), file))
		if err != nil {
			return pageInventory{}, err
		}
		doc.Inventory.Nodes = append(doc.Inventory.Nodes, node)
	}
	return doc.Inventory, nil
}

func loadNodeDocument(path string) (node, error) {
	var doc nodeDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return node{}, err
	}
	if doc.SchemaVersion != nodeSchemaVersion {
		return node{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.Node, nil
}

func loadToolLimitationDocument(path string) (toolLimitation, error) {
	var doc toolLimitationDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return toolLimitation{}, err
	}
	if doc.SchemaVersion != toolLimitationSchemaVersion {
		return toolLimitation{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.Limitation, nil
}
