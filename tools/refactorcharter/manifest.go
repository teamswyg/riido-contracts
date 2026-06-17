package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const schemaVersion = "riido-refactoring-charter.v1"

func loadCharter(path string) (charter, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return charter{}, fmt.Errorf("read charter: %w", err)
	}
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.DisallowUnknownFields()
	var c charter
	if err := dec.Decode(&c); err != nil {
		return charter{}, fmt.Errorf("decode charter: %w", err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); err != io.EOF {
		return charter{}, errors.New("decode charter: trailing data")
	}
	return c, nil
}

func verifyCharter(c charter) error {
	if c.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", c.SchemaVersion, schemaVersion)
	}
	if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.RiidoTask) == "" {
		return errors.New("id and riido_task are required")
	}
	if c.Mode != "advisory" && c.Mode != "enforced" {
		return fmt.Errorf("mode = %q, want advisory or enforced", c.Mode)
	}
	if c.LineBudget.TargetMaxLines <= 0 {
		return errors.New("line_budget.target_max_lines must be positive")
	}
	if len(c.SemanticUnits) == 0 || len(c.RequiredArtifacts) == 0 {
		return errors.New("semantic_units and required_artifacts are required")
	}
	if len(c.Scan.Roots) == 0 || len(c.Scan.IncludeExtensions) == 0 {
		return errors.New("scan.roots and scan.include_extensions are required")
	}
	return nil
}
