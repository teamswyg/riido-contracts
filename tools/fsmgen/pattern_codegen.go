package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type patternTemplateValue struct {
	Const       string `json:"Const"`
	Value       string `json:"Value"`
	CodeConst   string `json:"CodeConst"`
	StringConst string `json:"StringConst"`
}

type patternTemplateData struct {
	SourcePath     string                 `json:"SourcePath"`
	Package        string                 `json:"Package"`
	Type           string                 `json:"Type"`
	CodeType       string                 `json:"CodeType"`
	StringType     string                 `json:"StringType"`
	Values         []patternTemplateValue `json:"Values"`
	FirstCodeConst string                 `json:"FirstCodeConst"`
	LastCodeConst  string                 `json:"LastCodeConst"`
}

func generatePatternFiles(root string, patternDocs map[string]patternDocument) ([]generatedArtifact, error) {
	sources := make([]string, 0, len(patternDocs))
	for source := range patternDocs {
		sources = append(sources, source)
	}
	sort.Strings(sources)
	files := make([]generatedArtifact, 0, len(sources))
	outputPaths := map[string]string{}
	for _, source := range sources {
		file, err := generatePatternFile(root, source, patternDocs[source])
		if err != nil {
			return nil, err
		}
		if previous, ok := outputPaths[file.Path]; ok {
			return nil, fmt.Errorf("pattern sources %s and %s both generate %s", previous, source, file.Path)
		}
		outputPaths[file.Path] = source
		files = append(files, file)
	}
	return files, nil
}

func generatePatternFile(root, source string, patterns patternDocument) (generatedArtifact, error) {
	data := patternTemplateData{
		SourcePath: source,
		Package:    patterns.SumType.Package,
		Type:       patterns.SumType.Type,
		CodeType:   patterns.SumType.CodeType,
		StringType: patterns.SumType.StringType,
	}
	for _, value := range patterns.SumType.Values {
		suffix := patternConstSuffix(patterns.SumType, value.Const)
		data.Values = append(data.Values, patternTemplateValue{
			Const:       value.Const,
			Value:       value.Value,
			CodeConst:   patterns.SumType.CodeType + suffix,
			StringConst: patterns.SumType.StringType + suffix,
		})
	}
	data.FirstCodeConst = data.Values[0].CodeConst
	data.LastCodeConst = data.Values[len(data.Values)-1].CodeConst
	body, err := renderGoToolTemplate(root, "tools/fsmgen/templates/fsm_pattern.go.gotmpl", data)
	if err != nil {
		return generatedArtifact{}, err
	}
	formatted, err := formatSource(patterns.SumType.OutputPath, body)
	if err != nil {
		return generatedArtifact{}, err
	}
	return generatedArtifact{Path: patterns.SumType.OutputPath, Body: formatted}, nil
}

func renderGoToolTemplate(root, templatePath string, data any) ([]byte, error) {
	tempDir, err := os.MkdirTemp("", "riido-fsmgen-*")
	if err != nil {
		return nil, fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)
	dataPath := filepath.Join(tempDir, "data.json")
	outPath := filepath.Join(tempDir, "out.go")
	dataBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal template data: %w", err)
	}
	if err := os.WriteFile(dataPath, dataBody, 0o644); err != nil {
		return nil, fmt.Errorf("write template data: %w", err)
	}
	cmd := exec.Command("go", "tool", "gotmpl", "-body", filepath.FromSlash(templatePath), "-data", dataPath, "-out", outPath)
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("go tool gotmpl: %w\n%s", err, output)
	}
	body, err := os.ReadFile(outPath)
	if err != nil {
		return nil, fmt.Errorf("read rendered template: %w", err)
	}
	return body, nil
}

func patternConstSuffix(sumType patternSumType, constName string) string {
	if sumType.ConstPrefix != "" && strings.HasPrefix(constName, sumType.ConstPrefix) {
		suffix := strings.TrimPrefix(constName, sumType.ConstPrefix)
		if suffix != "" {
			return suffix
		}
	}
	return constName
}
