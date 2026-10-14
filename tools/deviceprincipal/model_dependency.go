package main

import (
	"fmt"
	"path/filepath"

	"github.com/teamswyg/riido-contracts/deviceprincipal"
)

const dependencyMapPath = "docs/30-architecture/ssot-dependency-map.riido.json"

func verifyDependencyFact(root, factID string) (int, error) {
	facts, err := loadDependencyFacts(root)
	if err != nil {
		return 0, err
	}
	fact, ok := findFact(facts, factID)
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

func loadDependencyFacts(root string) ([]dependencyFact, error) {
	path := resolve(root, dependencyMapPath)
	doc, err := readLooseJSONFile[dependencyMap](path)
	if err != nil {
		return nil, err
	}
	facts := append([]dependencyFact(nil), doc.Facts...)
	for _, file := range doc.FactFiles {
		include, err := readLooseJSONFile[dependencyFactDocument](filepath.Join(filepath.Dir(path), file))
		if err != nil {
			return nil, err
		}
		facts = append(facts, include.Fact)
	}
	return facts, nil
}

func findFact(facts []dependencyFact, factID string) (dependencyFact, bool) {
	for _, fact := range facts {
		if fact.ID == factID {
			return fact, true
		}
	}
	return dependencyFact{}, false
}
