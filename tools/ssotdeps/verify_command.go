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
	fmt.Fprintf(out, "ssotdeps: verified %d facts and %d repo dependencies\n", len(m.Facts), len(m.RepoDependencies))
	return nil
}
