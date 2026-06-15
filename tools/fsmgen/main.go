package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	modulePath  = "github.com/teamswyg/riido-contracts"
	sourcePath  = "enumgen/enums.lisp"
	generatedBy = "fsm gen"
	readmePath  = "README.md"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "fsmgen:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	command := "verify"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}
	fs := flag.NewFlagSet("fsmgen", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		return err
	}
	root, err := findRepoRoot()
	if err != nil {
		return err
	}
	metadata, err := loadFSMMetadata(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		return err
	}
	patternDocs, err := loadPatternDocuments(root, metadata)
	if err != nil {
		return err
	}
	if err := verifyConformance(metadata, patternDocs); err != nil {
		return err
	}
	files, err := generatedFiles(root, metadata, patternDocs)
	if err != nil {
		return err
	}
	sections, err := generatedReadmeSections(metadata)
	if err != nil {
		return err
	}
	switch command {
	case "conformance":
		fmt.Fprintf(out, "fsmgen: conformance verified %d FSMs against %d profile(s)\n", len(metadata), profileCount(patternDocs))
		return nil
	case "generate":
		for _, file := range files {
			path := filepath.Join(root, filepath.FromSlash(file.Path))
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				return fmt.Errorf("mkdir %s: %w", filepath.Dir(file.Path), err)
			}
			if err := os.WriteFile(path, file.Body, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", file.Path, err)
			}
		}
		if err := writeReadmeSections(root, sections); err != nil {
			return err
		}
		fmt.Fprintf(out, "fsmgen: generated %d files and %d README sections\n", len(files), len(sections))
		return nil
	case "verify":
		for _, file := range files {
			path := filepath.Join(root, filepath.FromSlash(file.Path))
			got, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read %s: %w", file.Path, err)
			}
			if !bytes.Equal(got, file.Body) {
				return fmt.Errorf("%s drifted; run go run ./tools/fsmgen generate", file.Path)
			}
		}
		if err := verifyReadmeSections(root, sections); err != nil {
			return err
		}
		fmt.Fprintf(out, "fsmgen: verified %d files and %d README sections\n", len(files), len(sections))
		return nil
	default:
		return errors.New("usage: go run ./tools/fsmgen [verify|generate|conformance]")
	}
}

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		body, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err == nil && strings.Contains(string(body), "module "+modulePath) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("cannot find riido-contracts repo root")
		}
		dir = parent
	}
}
