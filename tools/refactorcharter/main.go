package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const defaultManifest = "docs/30-architecture/refactoring-charter.riido.json"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "refactorcharter:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("refactorcharter", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	manifestPath := fs.String("manifest", defaultManifest, "refactoring charter manifest path")
	root := fs.String("root", ".", "repo root to scan")
	if err := fs.Parse(args); err != nil {
		return err
	}
	scanRoot := *root
	loadPath := *manifestPath
	if *manifestPath == defaultManifest && *root == "." {
		repoRoot, err := findRepoRoot()
		if err != nil {
			return err
		}
		scanRoot = repoRoot
		loadPath = filepath.Join(repoRoot, filepath.FromSlash(defaultManifest))
	}
	c, err := loadCharter(loadPath)
	if err != nil {
		return err
	}
	if err := verifyCharter(c); err != nil {
		return err
	}
	report, err := scan(scanRoot, c)
	if err != nil {
		return err
	}
	writeReport(out, c, report)
	if enforced(c) && len(report.Findings) > 0 {
		return fmt.Errorf("%d files exceed target_max_lines", len(report.Findings))
	}
	return nil
}
