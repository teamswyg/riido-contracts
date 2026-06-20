package main

import (
	"fmt"
	"io"
)

func runVerify(args []string, out io.Writer) error {
	fs := quietFlagSet("verify")
	manifestPath := manifestFlag(fs)
	checkDoc := fs.Bool("check-doc", false, "verify generated markdown is up to date")
	evidenceOut := fs.String("evidence-out", "", "optional evidence JSON output path")
	if err := fs.Parse(args); err != nil {
		return err
	}
	root, m, err := loadDefaultedManifest(*manifestPath)
	if err != nil {
		return err
	}
	if err := verifyRepository(root, m, *checkDoc); err != nil {
		return err
	}
	if *evidenceOut != "" {
		if err := writeEvidence(*evidenceOut, newEvidence(m, *checkDoc)); err != nil {
			return err
		}
	}
	fmt.Fprintf(out, "migrationdocs: verified slices=%d candidates=%d\n", len(m.MigrationSlices), len(m.CandidateContracts))
	return nil
}
