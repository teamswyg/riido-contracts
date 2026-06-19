package main

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"
)

func runRender(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("render", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	manifestPath := fs.String("manifest", defaultManifest, "Figma coverage manifest path")
	if err := fs.Parse(args); err != nil {
		return err
	}
	loadPath := *manifestPath
	if *manifestPath == defaultManifest {
		repoRoot, err := findRepoRoot()
		if err != nil {
			return err
		}
		loadPath = filepath.Join(repoRoot, filepath.FromSlash(*manifestPath))
	}
	m, err := loadManifest(loadPath)
	if err != nil {
		return err
	}
	if err := verifyManifest(m); err != nil {
		return err
	}
	_, err = fmt.Fprint(out, renderManifest(m))
	return err
}
