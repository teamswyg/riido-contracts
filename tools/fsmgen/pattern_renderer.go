package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func generatePatternSection(
	root, outputPath string,
	section patternSection,
	data patternTemplateData,
) (generatedArtifact, error) {
	path := patternSectionOutputPath(outputPath, section.Name)
	body, err := renderGoToolTemplate(root, section.Template, data)
	if err != nil {
		return generatedArtifact{}, err
	}
	formatted, err := formatSource(path, body)
	if err != nil {
		return generatedArtifact{}, err
	}
	return generatedArtifact{Path: path, Body: formatted}, nil
}

func renderGoToolTemplate(root, templatePath string, data any) ([]byte, error) {
	tempDir, err := os.MkdirTemp("", "riido-fsmgen-*")
	if err != nil {
		return nil, fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)
	return renderGoToolTemplateInTemp(root, templatePath, tempDir, data)
}

func renderGoToolTemplateInTemp(root, templatePath, tempDir string, data any) ([]byte, error) {
	dataPath := filepath.Join(tempDir, "data.json")
	outPath := filepath.Join(tempDir, "out.go")
	if err := writeTemplateData(dataPath, data); err != nil {
		return nil, err
	}
	cmd := exec.Command("go", "tool", "gotmpl", "-body", filepath.FromSlash(templatePath), "-data", dataPath, "-out", outPath)
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("go tool gotmpl: %w\n%s", err, output)
	}
	return os.ReadFile(outPath)
}

func writeTemplateData(path string, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal template data: %w", err)
	}
	if err := os.WriteFile(path, body, 0o644); err != nil {
		return fmt.Errorf("write template data: %w", err)
	}
	return nil
}
