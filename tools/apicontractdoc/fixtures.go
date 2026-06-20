package main

import (
	"os"

	"github.com/teamswyg/riido-contracts/apicontract"
)

func loadFixtureSummaries(root string, refs []fixtureRef) ([]fixtureSummary, map[string]bool, error) {
	summaries := make([]fixtureSummary, 0, len(refs))
	paths := map[string]bool{}
	for _, ref := range refs {
		ir, err := loadJSON[apicontract.IRDocument](resolve(root, ref.IR), ref.IR)
		if err != nil {
			return nil, nil, err
		}
		for _, path := range []string{ref.DSL, ref.IR, ref.OpenAPI} {
			if _, err := os.Stat(resolve(root, path)); err != nil {
				return nil, nil, err
			}
		}
		summaries = append(summaries, summarizeFixture(ref, ir))
		collectGeneratedPaths(paths, ir)
	}
	return summaries, paths, nil
}
