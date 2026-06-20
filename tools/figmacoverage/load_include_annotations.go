package main

import "fmt"

func loadAnnotationInventoryDocument(path string) (annotationInventory, error) {
	var doc annotationInventoryDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return annotationInventory{}, err
	}
	if doc.SchemaVersion != annotationInventoryVersion {
		return annotationInventory{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.Inventory, nil
}

func loadAnnotationDocument(path string) (annotation, error) {
	var doc annotationDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return annotation{}, err
	}
	if doc.SchemaVersion != annotationSchemaVersion {
		return annotation{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.Annotation, nil
}
