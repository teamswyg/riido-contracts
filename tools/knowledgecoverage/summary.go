package main

import (
	"os"
	"path/filepath"
	"strings"
)

const manualSampleLimit = 10

func summarize(root string, docs []docRecord) (scanReport, error) {
	report := scanReport{Docs: docs, ScannedCount: len(docs)}
	for _, doc := range docs {
		switch doc.Classification {
		case "generated_reader":
			report.GeneratedCount++
		case "executable_reader":
			report.ExecutableCount++
		case "manual_reader":
			report.ManualCount++
			if len(report.ManualSamples) < manualSampleLimit {
				report.ManualSamples = append(report.ManualSamples, doc)
			}
		}
		if doc.HasAdjacentManifest {
			report.AdjacentCount++
		}
	}
	count, err := countManifests(root)
	report.ManifestInventory = count
	return report, err
}

func countManifests(root string) (int, error) {
	count := 0
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}
		if entry.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".riido.json") {
			count++
		}
		return nil
	})
	return count, err
}
