package main

import (
	"fmt"
	"strings"
)

func replaceSection(body string, section readmeSection) (string, error) {
	start := sectionStart(section.ID)
	end := sectionEnd(section.ID)
	bounds, err := readmeSectionBounds(body, start, end)
	if err != nil {
		return "", err
	}
	replacement := "\n" + section.Content
	return body[:bounds.ContentStart] + replacement + body[bounds.ContentEnd:], nil
}

func extractSection(body, id string) (string, error) {
	start := sectionStart(id)
	end := sectionEnd(id)
	bounds, err := readmeSectionBounds(body, start, end)
	if err != nil {
		return "", err
	}
	content := body[bounds.ContentStart:bounds.ContentEnd]
	return strings.TrimPrefix(content, "\n"), nil
}

func sectionStart(id string) string {
	return "<!-- fsmgen:" + id + ":start -->"
}

func sectionEnd(id string) string {
	return "<!-- fsmgen:" + id + ":end -->"
}

func readmeSectionDriftError(id string) error {
	return fmt.Errorf("%s section %q drifted; run go run ./tools/fsmgen generate", readmePath, id)
}
