package main

import (
	"flag"
	"io"
	"path/filepath"
)

func quietFlagSet(name string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs
}

func manifestFlag(fs *flag.FlagSet) *string {
	return fs.String("manifest", defaultManifest, "context map manifest path")
}

func loadDefaultedManifest(path string) (string, manifest, error) {
	root := "."
	loadPath := path
	if path == defaultManifest {
		repoRoot, err := findRepoRoot()
		if err != nil {
			return "", manifest{}, err
		}
		root = repoRoot
		loadPath = filepath.Join(repoRoot, filepath.FromSlash(path))
	}
	m, err := loadManifest(loadPath)
	if err != nil {
		return "", manifest{}, err
	}
	return root, m, verifyManifest(m)
}
