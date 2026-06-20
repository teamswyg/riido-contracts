package main

import (
	"fmt"

	"github.com/teamswyg/riido-contracts/deviceprincipal"
)

const dependencyMapPath = "docs/30-architecture/ssot-dependency-map.riido.json"

func verifyDependencyFact(root, factID string) (int, error) {
	doc, err := readLooseJSONFile[dependencyMap](resolve(root, dependencyMapPath))
	if err != nil {
		return 0, err
	}
	fact, ok := findFact(doc.Facts, factID)
	if !ok {
		return 0, fmt.Errorf("%s missing fact %s", dependencyMapPath, factID)
	}
	for _, phrase := range deviceprincipal.DependencyPhrases() {
		if !factHasPhrase(fact, phrase) {
			return 0, fmt.Errorf("%s missing dependency phrase %q", factID, phrase)
		}
	}
	return len(fact.SourceRef), nil
}

func findFact(facts []dependencyFact, factID string) (dependencyFact, bool) {
	for _, fact := range facts {
		if fact.ID == factID {
			return fact, true
		}
	}
	return dependencyFact{}, false
}
