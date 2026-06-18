package main

import "sort"

func factsSorted(facts []fact) bool {
	ids := make([]string, 0, len(facts))
	for _, fact := range facts {
		ids = append(ids, fact.ID)
	}
	return stringsSorted(ids)
}

func repoDependenciesSorted(deps []repoDependency) bool {
	ids := make([]string, 0, len(deps))
	for _, dep := range deps {
		ids = append(ids, dep.ID)
	}
	return stringsSorted(ids)
}

func stringsSorted(values []string) bool {
	return sort.SliceIsSorted(values, func(i, j int) bool {
		return values[i] < values[j]
	})
}
