package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func loadPatternDocument(path string) (patternDocument, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return patternDocument{}, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return patternDocument{}, err
	}
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "fsm-pattern-gen" {
		return patternDocument{}, errors.New("pattern root form must be (fsm-pattern-gen ...)")
	}
	return patternDocumentFromRoot(root)
}

func patternDocumentFromRoot(root node) (patternDocument, error) {
	doc := patternDocument{Profiles: map[string]conformanceProfile{}}
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 {
			return patternDocument{}, errors.New("fsm-pattern-gen children must be lists")
		}
		if err := applyPatternDocumentForm(&doc, form); err != nil {
			return patternDocument{}, err
		}
	}
	return validatePatternDocument(doc)
}

func loadPatternDocuments(root string, metadata map[string]fsmMetadata) (map[string]patternDocument, error) {
	sources, err := patternSources(metadata)
	if err != nil {
		return nil, err
	}
	docs := make(map[string]patternDocument, len(sources))
	for source := range sources {
		doc, err := loadPatternDocument(filepath.Join(root, filepath.FromSlash(source)))
		if err != nil {
			return nil, err
		}
		docs[source] = doc
	}
	return docs, nil
}
