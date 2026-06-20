package main

import "fmt"

func build(root string, opt options) (manifest, []fixtureSummary, string, error) {
	m, err := loadJSON[manifest](resolve(root, opt.manifest), "manifest")
	if err != nil {
		return manifest{}, nil, "", err
	}
	if err := verifyManifest(m); err != nil {
		return manifest{}, nil, "", err
	}
	if err := verifyWorkflow(root, m); err != nil {
		return manifest{}, nil, "", err
	}
	summaries, generatedPaths, err := loadFixtureSummaries(root, m.Fixtures)
	if err != nil {
		return manifest{}, nil, "", err
	}
	for _, path := range m.RequiredGeneratedPaths {
		if !generatedPaths[path] {
			return manifest{}, nil, "", fmt.Errorf("required generated path %q is missing", path)
		}
	}
	return m, summaries, renderDoc(m, summaries), nil
}
