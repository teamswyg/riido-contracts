package main

import (
	"fmt"
	"os"
	"strings"
)

func attachFSMSections(root string, m manifest) (manifest, error) {
	body, err := os.ReadFile(repoPath(root, generatedDoc))
	if err != nil {
		return manifest{}, fmt.Errorf("read %s: %w", generatedDoc, err)
	}
	for i := range m.FSM.Sections {
		section := &m.FSM.Sections[i]
		sectionBody, err := extractFSMSection(string(body), section.ID)
		if err != nil {
			return manifest{}, err
		}
		section.Body = sectionBody
	}
	return m, nil
}

func extractFSMSection(readme, id string) (string, error) {
	start := fmt.Sprintf("<!-- fsmgen:%s:start -->", id)
	end := fmt.Sprintf("<!-- fsmgen:%s:end -->", id)
	startIndex := strings.Index(readme, start)
	if startIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", generatedDoc, start)
	}
	endIndex := strings.Index(readme[startIndex:], end)
	if endIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", generatedDoc, end)
	}
	endIndex += startIndex + len(end)
	return strings.TrimSpace(readme[startIndex:endIndex]), nil
}
