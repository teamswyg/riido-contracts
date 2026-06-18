package main

import (
	"errors"
	"fmt"
)

func documentFromNode(root node) (document, error) {
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return document{}, errors.New("root form must be (enum-gen ...)")
	}
	var doc document
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 {
			return document{}, errors.New("enum-gen children must be lists")
		}
		if err := appendDocumentForm(&doc, form); err != nil {
			return document{}, err
		}
	}
	return doc, validateDocument(doc)
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
