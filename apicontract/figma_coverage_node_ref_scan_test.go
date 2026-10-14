package apicontract

import (
	"net/url"
	"os"
	"regexp"
	"testing"
)

var (
	figmaNodeIDRefPattern      = regexp.MustCompile(`node-id=([A-Za-z0-9%][A-Za-z0-9:;%_-]*)`)
	numericFigmaURLNodePattern = regexp.MustCompile(`^([0-9]+)-([0-9]+)$`)
)

func assertFigmaNodeRefsInFileAreRegistered(t *testing.T, path string, registered map[string]string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	for _, match := range figmaNodeIDRefPattern.FindAllStringSubmatch(string(data), -1) {
		nodeID := normalizeFigmaNodeID(match[1])
		if _, ok := registered[nodeID]; !ok {
			t.Fatalf("%s cites unregistered Figma node-id=%s; add it to figma-ai-agent-coverage.riido.json or remove the stale citation", path, match[1])
		}
	}
}

func normalizeFigmaNodeID(raw string) string {
	unescaped, err := url.QueryUnescape(raw)
	if err != nil {
		unescaped = raw
	}
	if match := numericFigmaURLNodePattern.FindStringSubmatch(unescaped); match != nil {
		return match[1] + ":" + match[2]
	}
	return unescaped
}
