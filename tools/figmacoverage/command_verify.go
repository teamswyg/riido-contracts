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
	fmt.Fprintf(out, "figmacoverage: verified %d entries and %d generated annotations\n", len(m.Entries), len(m.APIGeneratedAnnotationInventory))
	return nil
}
