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
	manifestPath := fs.String("manifest", defaultManifest, "Figma coverage manifest path")
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
			SchemaVersion:                     evidenceSchemaVersion,
			ID:                                m.ID,
			Status:                            "verified",
			ToolLimitationFilesLoaded:         len(m.ToolLimitationFiles),
			ExpectedTopLevelNodeFilesLoaded:   len(m.ExpectedTopLevelNodeFiles),
			PageInventoryFilesLoaded:          len(m.PageInventoryFiles),
			CoverageEntryFilesLoaded:          len(m.CoverageEntryFiles),
			NonUICoverageEntryFilesLoaded:     len(m.NonUICoverageEntryFiles),
			APIAnnotationInventoryFilesLoaded: len(m.APIAnnotationInventoryFiles),
			APIAnnotationFilesLoaded:          len(m.APIAnnotationFiles),
			VerifiedEvidenceNodeFilesLoaded:   len(m.VerifiedEvidenceNodeFiles),
			EntriesVerified:                   len(m.Entries),
			GeneratedAnnotationsChecked:       len(m.APIGeneratedAnnotationInventory),
			EvidenceNodesVerified:             len(m.VerifiedEvidenceNodes),
			CheckDoc:                          *checkDoc,
		}); err != nil {
			return err
		}
	}
	fmt.Fprintf(out, "figmacoverage: verified %d entries and %d generated annotations\n", len(m.Entries), len(m.APIGeneratedAnnotationInventory))
	return nil
}
