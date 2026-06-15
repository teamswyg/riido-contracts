package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func generatedReadmeSections(metadata map[string]fsmMetadata) ([]readmeSection, error) {
	taskMeta, err := requireFSMMetadata(metadata, "task", "TaskTransitionCode")
	if err != nil {
		return nil, err
	}
	taskStart, err := taskStateCodesFromConsts(taskMeta.StartPoints)
	if err != nil {
		return nil, err
	}
	taskEnd, err := taskStateCodesFromConsts(taskMeta.EndPoints)
	if err != nil {
		return nil, err
	}
	assignmentMeta, err := requireFSMMetadata(metadata, "assignment", "AssignmentTransitionCode")
	if err != nil {
		return nil, err
	}
	assignmentStart, err := assignmentStateCodesFromConsts(assignmentMeta.StartPoints)
	if err != nil {
		return nil, err
	}
	assignmentEnd, err := assignmentStateCodesFromConsts(assignmentMeta.EndPoints)
	if err != nil {
		return nil, err
	}
	return []readmeSection{
		{
			ID:      taskMeta.ReadmeSection,
			Content: mermaidFence(taskMermaid(taskStart, taskEnd)),
		},
		{
			ID:      assignmentMeta.ReadmeSection,
			Content: mermaidFence(assignmentMermaid(assignmentStart, assignmentEnd)),
		},
	}, nil
}

func mermaidFence(body string) string {
	return "```mermaid\n" + body + "```\n"
}

func writeReadmeSections(root string, sections []readmeSection) error {
	path := filepath.Join(root, readmePath)
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	updated := string(body)
	for _, section := range sections {
		var replaceErr error
		updated, replaceErr = replaceSection(updated, section)
		if replaceErr != nil {
			return replaceErr
		}
	}
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", readmePath, err)
	}
	return nil
}

func verifyReadmeSections(root string, sections []readmeSection) error {
	path := filepath.Join(root, readmePath)
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	current := string(body)
	for _, section := range sections {
		got, err := extractSection(current, section.ID)
		if err != nil {
			return err
		}
		if got != section.Content {
			return fmt.Errorf("%s section %q drifted; run go run ./tools/fsmgen generate", readmePath, section.ID)
		}
	}
	return nil
}

func replaceSection(body string, section readmeSection) (string, error) {
	start := sectionStart(section.ID)
	end := sectionEnd(section.ID)
	startIndex := strings.Index(body, start)
	if startIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, start)
	}
	contentStart := startIndex + len(start)
	endIndex := strings.Index(body[contentStart:], end)
	if endIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, end)
	}
	contentEnd := contentStart + endIndex
	replacement := "\n" + section.Content
	return body[:contentStart] + replacement + body[contentEnd:], nil
}

func extractSection(body, id string) (string, error) {
	start := sectionStart(id)
	end := sectionEnd(id)
	startIndex := strings.Index(body, start)
	if startIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, start)
	}
	contentStart := startIndex + len(start)
	endIndex := strings.Index(body[contentStart:], end)
	if endIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, end)
	}
	contentEnd := contentStart + endIndex
	content := body[contentStart:contentEnd]
	return strings.TrimPrefix(content, "\n"), nil
}

func sectionStart(id string) string {
	return "<!-- fsmgen:" + id + ":start -->"
}

func sectionEnd(id string) string {
	return "<!-- fsmgen:" + id + ":end -->"
}
