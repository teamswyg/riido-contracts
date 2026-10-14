package main

import "path/filepath"

type enumRunPlan struct {
	Root  string
	Files map[string][]byte
}

func buildEnumRunPlan() (enumRunPlan, error) {
	root, err := findRepoRoot()
	if err != nil {
		return enumRunPlan{}, err
	}
	doc, err := loadDocument(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		return enumRunPlan{}, err
	}
	files, err := generatedFiles(doc)
	if err != nil {
		return enumRunPlan{}, err
	}
	return enumRunPlan{Root: root, Files: files}, nil
}
