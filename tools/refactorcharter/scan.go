package main

import (
	"bufio"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func scan(root string, c charter) (scanReport, error) {
	var report scanReport
	for _, scanRoot := range c.Scan.Roots {
		base := filepath.Join(root, filepath.FromSlash(scanRoot))
		err := filepath.WalkDir(base, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			rel = filepath.ToSlash(rel)
			if !includedFile(rel, c.Scan.IncludeExtensions) || generatedFile(root, rel, c.Scan) {
				return nil
			}
			lines, err := countLines(path)
			if err != nil {
				return err
			}
			report.FilesScanned++
			if lines > c.LineBudget.TargetMaxLines {
				report.Findings = append(report.Findings, finding{Path: rel, Lines: lines})
			}
			return nil
		})
		if err != nil {
			return scanReport{}, err
		}
	}
	slices.SortFunc(report.Findings, func(a, b finding) int {
		if b.Lines != a.Lines {
			return b.Lines - a.Lines
		}
		return strings.Compare(a.Path, b.Path)
	})
	return report, nil
}

func includedFile(path string, extensions []string) bool {
	ext := filepath.Ext(path)
	return slices.Contains(extensions, ext)
}

func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
