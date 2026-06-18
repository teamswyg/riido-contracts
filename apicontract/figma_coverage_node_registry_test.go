package apicontract

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func registerFigmaNode(t *testing.T, registered map[string]string, node figmaCoverageNode, source string) {
	t.Helper()
	if strings.TrimSpace(node.NodeID) == "" || strings.TrimSpace(node.Name) == "" {
		t.Fatalf("%s has empty node field: %+v", source, node)
	}
	if existing, exists := registered[node.NodeID]; exists {
		t.Fatalf("duplicate Figma node %q in %s; already registered by %s", node.NodeID, source, existing)
	}
	registered[node.NodeID] = source
}

func registerFigmaNodeIfAbsent(t *testing.T, registered map[string]string, node figmaCoverageNode, source string) {
	t.Helper()
	if strings.TrimSpace(node.NodeID) == "" || strings.TrimSpace(node.Name) == "" {
		t.Fatalf("%s has empty node field: %+v", source, node)
	}
	if _, exists := registered[node.NodeID]; exists {
		return
	}
	registered[node.NodeID] = source
}

func registerFigmaNodeIDIfAbsent(t *testing.T, registered map[string]string, nodeID, source string) {
	t.Helper()
	if strings.TrimSpace(nodeID) == "" {
		t.Fatalf("%s has empty node id", source)
	}
	if _, exists := registered[nodeID]; exists {
		return
	}
	registered[nodeID] = source
}

func assertDocumentedFigmaNodeRefsAreRegistered(t *testing.T, registered map[string]string) {
	t.Helper()
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
			assertFigmaNodeRefsInFileAreRegistered(t, root, registered)
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
				assertFigmaNodeRefsInFileAreRegistered(t, path, registered)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s: %v", root, err)
		}
	}
}
