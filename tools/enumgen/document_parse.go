package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func documentFromRoot(root node, base string, seen map[string]bool) (document, error) {
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return document{}, errors.New("root form must be (enum-gen ...)")
	}
	var doc document
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 {
			return document{}, errors.New("enum-gen children must be lists")
		}
		if isIncludeForm(form) {
			included, err := loadDocumentInclude(base, form, seen)
			if err != nil {
				return document{}, err
			}
			doc.Enums = append(doc.Enums, included.Enums...)
			doc.Transitions = append(doc.Transitions, included.Transitions...)
			continue
		}
		if err := appendDocumentForm(&doc, form); err != nil {
			return document{}, err
		}
	}
	return doc, nil
}

func appendDocumentForm(doc *document, form node) error {
	switch atom(form.list[0]) {
	case "enum":
		spec, err := parseEnum(form)
		if err != nil {
			return err
		}
		doc.Enums = append(doc.Enums, spec)
	case "transitions":
		spec, err := parseTransitions(form)
		if err != nil {
			return err
		}
		doc.Transitions = append(doc.Transitions, spec)
	default:
		return fmt.Errorf("unknown form %q", atom(form.list[0]))
	}
	return nil
}

func loadDocumentInclude(base string, form node, seen map[string]bool) (document, error) {
	path, err := includePath(base, form)
	if err != nil {
		return document{}, err
	}
	return loadDocumentAt(path, filepath.Dir(path), seen)
}
