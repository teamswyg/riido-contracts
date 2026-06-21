package main

const (
	manualSampleLimit   = 10
	manifestSampleLimit = 3
)

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
			if doc.Classification == "generated_reader" {
				report.GeneratedAdjacentCount++
			}
			if doc.Classification == "executable_reader" {
				report.ExecutableAdjacentCount++
			}
		}
	}
	count, groups, err := countManifests(root)
	report.ManifestInventory = count
	report.ManifestInventoryByGroup = groups
	if err != nil {
		return report, err
	}
	report.ManifestInventorySamples, err = manifestInventorySamples(root, groups, manifestSampleLimit)
	if err != nil {
		return report, err
	}
	report.ManifestLoops, err = scanManifestLoops(root)
	return report, err
}
