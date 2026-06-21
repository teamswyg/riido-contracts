package main

import (
	"encoding/json"
	"os"
)

const manifestLoopSampleLimit = 3

func scanManifestLoops(root string) (manifestLoopReport, error) {
	report := manifestLoopReport{}
	missingByGroup := map[string]int{}
	missingSamples := map[string][]string{}
	paths, err := manifestPaths(root)
	if err != nil {
		return report, err
	}
	for _, path := range paths {
		group := manifestGroup(root, path)
		if manifestHasLoop(path) {
			report.Complete++
			continue
		}
		report.Missing++
		missingByGroup[group]++
		if len(missingSamples[group]) < manifestLoopSampleLimit {
			missingSamples[group] = append(missingSamples[group], rel(root, path))
		}
	}
	report.MissingGroups = manifestGroups(missingByGroup)
	report.MissingSamples = orderedManifestSamples(report.MissingGroups, missingSamples)
	return report, nil
}

func manifestHasLoop(path string) bool {
	var doc map[string]any
	if err := readJSON(path, &doc); err != nil {
		return false
	}
	loop, ok := doc["loop"].(map[string]any)
	if !ok {
		return false
	}
	for _, key := range []string{"observation", "hypothesis", "execute", "evaluate", "retrospective"} {
		value, ok := loop[key].(string)
		if !ok || value == "" {
			return false
		}
	}
	return true
}

func readJSON(path string, v any) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
