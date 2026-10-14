package apicontract

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertNoStaleRuntimeEndpointHostPinned(t *testing.T) {
	t.Helper()
	forbidden := "desktop-api." + "riido.ai"
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
			assertFileDoesNotContain(t, root, forbidden)
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
				assertFileDoesNotContain(t, path, forbidden)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s for stale runtime endpoint host: %v", root, err)
		}
	}
}

func assertFileDoesNotContain(t *testing.T, path, forbidden string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if strings.Contains(string(data), forbidden) {
		t.Fatalf("%s pins stale endpoint-looking Figma host; cite node-id=129:17930 and explain it is not canonical instead", path)
	}
}
