package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyManifest(m manifest) error {
	if m.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", m.SchemaVersion, schemaVersion)
	}
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.GeneratedDoc) || blank(m.Summary) {
		return errors.New("id, riido_task, generated_doc, and summary are required")
	}
	if len(m.PromotionConditions) != 5 {
		return fmt.Errorf("promotion_conditions=%d, want 5", len(m.PromotionConditions))
	}
	if !containsPhrase(m.PromotionConditions, "docs/migration/contracts.md") {
		return errors.New("promotion conditions must require migration docs update")
	}
	if len(m.SchemaVersionAxes) == 0 || len(m.RuntimeTagModel) != 3 {
		return errors.New("schema_version_axes and three runtime_tag_model entries are required")
	}
	if len(m.BreakingChangeRules) == 0 || blank(m.DownstreamImportRule) {
		return errors.New("breaking_change_rules and downstream_import_rule are required")
	}
	return nil
}

func blank(s string) bool {
	return strings.TrimSpace(s) == ""
}
