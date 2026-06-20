package main

import "strings"

type model struct {
	Manifest             manifest
	Document             string
	TermCount            int
	PolicyAssertionCount int
	FigmaNodeRefCount    int
	APIPathRefCount      int
	GeneratedReaderCount int
}

func build(root string, opt options) (model, error) {
	m, err := loadManifest(resolve(root, opt.manifest))
	if err != nil {
		return model{}, err
	}
	if err := verifyManifest(m); err != nil {
		return model{}, err
	}
	if err := verifyGeneratedReaders(root, m); err != nil {
		return model{}, err
	}
	if err := verifyWorkflow(root, m); err != nil {
		return model{}, err
	}
	doc := renderDoc(m)
	return model{
		Manifest:             m,
		Document:             doc,
		TermCount:            len(m.RequiredTerms),
		PolicyAssertionCount: len(m.RequiredPolicyAssertions),
		FigmaNodeRefCount:    strings.Count(doc, "node-id="),
		APIPathRefCount:      countAPIPaths(doc),
		GeneratedReaderCount: len(m.RequiredGeneratedReaders),
	}, nil
}
