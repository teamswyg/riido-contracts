package main

func manifestInventorySamples(root string, groups []manifestGroupCount, limit int) ([]manifestGroupSample, error) {
	byGroup := map[string][]string{}
	paths, err := manifestPaths(root)
	for _, path := range paths {
		group := manifestGroup(root, path)
		if len(byGroup[group]) < limit {
			byGroup[group] = append(byGroup[group], rel(root, path))
		}
	}
	return orderedManifestSamples(groups, byGroup), err
}

func orderedManifestSamples(groups []manifestGroupCount, byGroup map[string][]string) []manifestGroupSample {
	samples := make([]manifestGroupSample, 0, len(groups))
	for _, group := range groups {
		samples = append(samples, manifestGroupSample{Group: group.Group, Paths: byGroup[group.Group]})
	}
	return samples
}
