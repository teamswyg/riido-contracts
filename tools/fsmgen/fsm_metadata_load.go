package main

import (
	"errors"
	"fmt"
	"os"
)

func loadFSMMetadata(path string) (map[string]fsmMetadata, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return nil, err
	}
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return nil, errors.New("root form must be (enum-gen ...)")
	}
	return fsmMetadataFromRoot(root)
}

func fsmMetadataFromRoot(root node) (map[string]fsmMetadata, error) {
	metadata := map[string]fsmMetadata{}
	for _, form := range root.list[1:] {
		if !isTransitionForm(form) {
			continue
		}
		if err := addFSMMetadataForm(metadata, form); err != nil {
			return nil, err
		}
	}
	return metadata, nil
}

func isTransitionForm(form node) bool {
	return !form.isAtom() && len(form.list) > 0 && atom(form.list[0]) == "transitions"
}
