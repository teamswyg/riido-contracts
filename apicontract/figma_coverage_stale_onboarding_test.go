package apicontract

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertNoStaleOnboardingFixtureWording(t *testing.T) {
	t.Helper()
	forbidden := []string{
		"starter-agent",
		"starter agent",
		"starter agents",
		"starter row",
		"starter rows",
		"starter fixture",
		"starter fixtures",
		"starter-fixture",
	}
	for _, root := range []string{
		filepath.FromSlash("../docs"),
		filepath.FromSlash("fixtures"),
		filepath.FromSlash("../README.md"),
	} {
		info, err := os.Stat(root)
		if err != nil {
			t.Fatalf("stat %s: %v", root, err)
		}
		if !info.IsDir() {
			assertNoStaleOnboardingFixtureWordingInFile(t, root, forbidden)
			continue
		}
		err = filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() {
				return nil
			}
			switch filepath.Ext(path) {
			case ".md", ".json":
				assertNoStaleOnboardingFixtureWordingInFile(t, path, forbidden)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s for stale onboarding fixture wording: %v", root, err)
		}
	}
}

func assertNoStaleOnboardingFixtureWordingInFile(t *testing.T, path string, forbidden []string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	haystack := strings.ToLower(string(data))
	for _, phrase := range forbidden {
		if strings.Contains(haystack, phrase) {
			t.Fatalf("%s contains stale onboarding fixture wording %q; use onboarding fixture wording instead", path, phrase)
		}
	}
}
