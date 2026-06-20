package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func scanDocs(root string, m manifest) (scanReport, error) {
	var docs []docRecord
	for _, scanRoot := range m.ScanRoots {
		base := resolve(root, scanRoot)
		err := filepath.WalkDir(base, func(path string, entry os.DirEntry, err error) error {
			if err != nil || entry.IsDir() || filepath.Ext(path) != ".md" {
				return err
			}
			record, err := classifyDoc(root, path, m.GeneratedMarkers)
			if err != nil {
				return err
			}
			docs = append(docs, record)
			return nil
		})
		if err != nil {
			return scanReport{}, err
		}
	}
	for _, scanFile := range m.ScanFiles {
		record, err := classifyDoc(root, resolve(root, scanFile), m.GeneratedMarkers)
		if err != nil {
			return scanReport{}, err
		}
		docs = append(docs, record)
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].Path < docs[j].Path })
	return summarize(root, docs)
}

func classifyDoc(root, path string, markers []string) (docRecord, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return docRecord{}, err
	}
	text := string(body)
	record := docRecord{
		Path:                rel(root, path),
		Lines:               strings.Count(text, "\n"),
		HasGeneratedMarker:  containsAny(text, markers),
		HasExecutableMarker: strings.Contains(text, "Executable SSOT"),
		HasAdjacentManifest: adjacentManifestExists(path),
	}
	record.Classification = classify(record)
	return record, nil
}
