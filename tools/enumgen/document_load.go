package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func loadDocument(path string) (document, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return document{}, fmt.Errorf("resolve %s: %w", path, err)
	}
	doc, err := loadDocumentAt(abs, filepath.Dir(abs), map[string]bool{})
	if err != nil {
		return document{}, err
	}
	return doc, validateDocument(doc)
}

func loadDocumentAt(path, base string, seen map[string]bool) (document, error) {
	if seen[path] {
		return document{}, fmt.Errorf("include cycle at %s", path)
	}
	seen[path] = true
	defer delete(seen, path)
	body, err := os.ReadFile(path)
	if err != nil {
		return document{}, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return document{}, err
	}
	return documentFromLoadedNode(root, base, seen)
}

func documentFromLoadedNode(root node, base string, seen map[string]bool) (document, error) {
	if !root.isAtom() && len(root.list) > 0 && atom(root.list[0]) == "enum-gen" {
		return documentFromRoot(root, base, seen)
	}
	var doc document
	if err := appendDocumentForm(&doc, root); err != nil {
		return document{}, err
	}
	return doc, nil
}
