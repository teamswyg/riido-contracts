package main

import "path/filepath"

func loadAnnotationIncludes(base string, m *manifest) error {
	for _, file := range m.APIAnnotationInventoryFiles {
		item, err := loadAnnotationInventoryDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.APIGeneratedAnnotationInventory = append(m.APIGeneratedAnnotationInventory, item)
	}
	for _, file := range m.APIAnnotationFiles {
		annotation, err := loadAnnotationDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.APIGeneratedAnnotations = append(m.APIGeneratedAnnotations, annotation)
	}
	return nil
}
