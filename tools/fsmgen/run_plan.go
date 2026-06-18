package main

import "path/filepath"

type fsmRunPlan struct {
	Root        string
	Metadata    map[string]fsmMetadata
	PatternDocs map[string]patternDocument
	Files       []generatedArtifact
	Sections    []readmeSection
}

func buildFSMRunPlan() (fsmRunPlan, error) {
	root, err := findRepoRoot()
	if err != nil {
		return fsmRunPlan{}, err
	}
	metadata, err := loadFSMMetadata(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		return fsmRunPlan{}, err
	}
	patternDocs, err := loadPatternDocuments(root, metadata)
	if err != nil {
		return fsmRunPlan{}, err
	}
	if err := verifyConformance(metadata, patternDocs); err != nil {
		return fsmRunPlan{}, err
	}
	files, err := generatedFiles(root, metadata, patternDocs)
	if err != nil {
		return fsmRunPlan{}, err
	}
	sections, err := generatedReadmeSections(metadata)
	if err != nil {
		return fsmRunPlan{}, err
	}
	return fsmRunPlan{
		Root:        root,
		Metadata:    metadata,
		PatternDocs: patternDocs,
		Files:       files,
		Sections:    sections,
	}, nil
}
