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
	manifestPath := fs.String("manifest", defaultManifest, "SSOT dependency manifest path")
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
	if err := verifyManifest(m, root); err != nil {
		return err
	}
	if err := verifyWorkflowBinding(root, m); err != nil {
		return err
	}
	if *checkDoc {
		if err := verifyRenderedDoc(root, m); err != nil {
			return err
		}
	}
	if *evidenceOut != "" {
		if err := writeEvidence(*evidenceOut, evidence{
			SchemaVersion:             evidenceSchemaVersion,
			ID:                        m.ID,
			Status:                    "verified",
			Workflow:                  m.Workflow,
			EvidenceArtifact:          m.EvidenceArtifact,
			FactFilesLoaded:           len(m.FactFiles),
			FactsVerified:             len(m.Facts),
			RepoDependencyFilesLoaded: len(m.RepoDependencyFiles),
			RepoDependenciesChecked:   len(m.RepoDependencies),
			CheckDoc:                  *checkDoc,
			Loop:                      m.Loop,
		}); err != nil {
			return err
		}
	}
	fmt.Fprintf(out, "ssotdeps: verified %d facts and %d repo dependencies\n", len(m.Facts), len(m.RepoDependencies))
	return nil
}
