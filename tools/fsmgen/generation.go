package main

func generatedFiles(root string, metadata map[string]fsmMetadata, patternDocs map[string]patternDocument) ([]generatedArtifact, error) {
	files, err := generatedFSMFiles(metadata)
	if err != nil {
		return nil, err
	}
	patternFiles, err := generatePatternFiles(root, patternDocs)
	if err != nil {
		return nil, err
	}
	files = append(patternFiles, files...)
	return files, nil
}
