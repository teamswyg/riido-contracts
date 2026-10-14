package main

const manifestLoopSampleLimit = 3

func scanManifestLoops(root string, sources []manifestLoopSource) (manifestLoopReport, error) {
	report := manifestLoopReport{}
	missingByGroup := map[string]int{}
	missingSamples := map[string][]string{}
	paths, err := manifestPaths(root)
	if err != nil {
		return report, err
	}
	for _, path := range paths {
		group := manifestGroup(root, path)
		switch manifestLoopStatusWithSources(root, path, sources) {
		case "direct":
			report.Complete++
			report.Direct++
			continue
		case "delegated":
			report.Complete++
			report.Delegated++
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

func manifestDocHasLoop(doc map[string]any) bool {
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
