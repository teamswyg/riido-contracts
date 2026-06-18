package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyFact(root, humanDoc string, f fact) error {
	if err := requireID("id", f.ID); err != nil {
		return err
	}
	if strings.TrimSpace(f.Fact) == "" {
		return errors.New("fact is required")
	}
	if strings.TrimSpace(f.HumanDocPhrase) == "" {
		return errors.New("human_doc_phrase is required")
	}
	if !strings.Contains(humanDoc, f.HumanDocPhrase) {
		return fmt.Errorf("human_doc does not contain phrase %q", f.HumanDocPhrase)
	}
	if err := verifyOwner(root, f.Owner); err != nil {
		return err
	}
	if len(f.SourceRefs) == 0 {
		return errors.New("source_refs are required")
	}
	if err := verifySourceRefs(root, f.SourceRefs); err != nil {
		return err
	}
	if len(f.Downstreams) == 0 {
		return errors.New("downstreams are required")
	}
	return verifyDownstreams(f.Downstreams)
}
