package main

import (
	"fmt"
	"sort"
	"strings"
)

func generatePatternFiles(root string, patternDocs map[string]patternDocument) ([]generatedArtifact, error) {
	sources := make([]string, 0, len(patternDocs))
	for source := range patternDocs {
		sources = append(sources, source)
	}
	sort.Strings(sources)
	files := make([]generatedArtifact, 0, len(sources))
	outputPaths := map[string]string{}
	for _, source := range sources {
		sourceFiles, err := generatePatternFilesForSource(root, source, patternDocs[source])
		if err != nil {
			return nil, err
		}
		for _, file := range sourceFiles {
			if previous, ok := outputPaths[file.Path]; ok {
				return nil, fmt.Errorf("pattern sources %s and %s both generate %s", previous, source, file.Path)
			}
			outputPaths[file.Path] = source
			files = append(files, file)
		}
	}
	return files, nil
}

func generatePatternFilesForSource(root, source string, patterns patternDocument) ([]generatedArtifact, error) {
	data := patternTemplateDataFrom(source, patterns)
	files := make([]generatedArtifact, 0, len(patternSections()))
	for _, section := range patternSections() {
		file, err := generatePatternSection(root, patterns.SumType.OutputPath, section, data)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func patternConstSuffix(sumType patternSumType, constName string) string {
	if sumType.ConstPrefix != "" && strings.HasPrefix(constName, sumType.ConstPrefix) {
		suffix := strings.TrimPrefix(constName, sumType.ConstPrefix)
		if suffix != "" {
			return suffix
		}
	}
	return constName
}
