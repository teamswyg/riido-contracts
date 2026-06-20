package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func verifyGeneratedReaders(root string, m manifest) error {
	docDir := filepath.Dir(resolve(root, m.GeneratedDoc))
	doc := renderDoc(m)
	for _, reader := range m.RequiredGeneratedReaders {
		if strings.Contains(reader, "/") || strings.Contains(reader, "..") {
			return fmt.Errorf("generated reader %q must be a local doc link", reader)
		}
		if _, err := os.Stat(filepath.Join(docDir, reader)); err != nil {
			return fmt.Errorf("generated reader %q is missing: %w", reader, err)
		}
		if !strings.Contains(doc, "]("+reader+")") {
			return fmt.Errorf("generated reader %q is not linked from generated doc", reader)
		}
	}
	if strings.Count(doc, "node-id=") == 0 || countAPIPaths(doc) == 0 {
		return fmt.Errorf("policy doc must include Figma node and API path references")
	}
	return nil
}
