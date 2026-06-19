package main

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"
)

func runVerify(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("verify", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	manifestPath := fs.String("manifest", defaultManifest, "contract promotion policy manifest path")
	checkDoc := fs.Bool("check-doc", false, "verify generated markdown is up to date")
	evidenceOut := fs.String("evidence-out", "", "optional evidence JSON output path")
	if err := fs.Parse(args); err != nil {
		return err
	}
	root := "."
	loadPath := *manifestPath
	if *manifestPath == defaultManifest {
		repoRoot, err := findRepoRoot()
		if err != nil {
			return err
		}
		root = repoRoot
		loadPath = filepath.Join(repoRoot, filepath.FromSlash(*manifestPath))
	}
	m, err := loadManifest(loadPath)
	if err != nil {
		return err
	}
	if err := verifyManifest(m); err != nil {
		return err
	}
	if *checkDoc {
		if err := verifyRenderedDoc(root, m); err != nil {
			return err
		}
	}
	if *evidenceOut != "" {
		if err := writeEvidence(*evidenceOut, evidence{
			SchemaVersion:              evidenceSchemaVersion,
			ID:                         m.ID,
			Status:                     "verified",
			PromotionConditionsChecked: len(m.PromotionConditions),
			BreakingRulesChecked:       len(m.BreakingChangeRules),
			RuntimeTagRulesChecked:     len(m.RuntimeTagModel),
			CheckDoc:                   *checkDoc,
		}); err != nil {
			return err
		}
	}
	fmt.Fprintf(out, "promotionpolicy: verified %d promotion conditions and %d breaking rules\n", len(m.PromotionConditions), len(m.BreakingChangeRules))
	return nil
}
