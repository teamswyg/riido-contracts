package apicontract

import (
	"strings"
	"testing"
)

func (s *figmaCoverageTestScope) verifyManifestEnvelope(t *testing.T) {
	t.Helper()
	if s.manifest.SchemaVersion != "riido-figma-ai-agent-coverage.v1" {
		t.Fatalf("schema_version = %q", s.manifest.SchemaVersion)
	}
	if s.manifest.ID != "figma-v1-22-ai-agent-ui-coverage" {
		t.Fatalf("id = %q", s.manifest.ID)
	}
	if s.manifest.RiidoTask != "RIID-4809" {
		t.Fatalf("riido_task = %q", s.manifest.RiidoTask)
	}
	verifyFigmaCoverageProvenance(t, s.manifest.StabilizedBy, s.docPath)
	if s.manifest.HumanDoc != "docs/30-architecture/figma-ai-agent-coverage.md" {
		t.Fatalf("human_doc = %q", s.manifest.HumanDoc)
	}
	if s.manifest.Figma.FileKey != "MUOd9lctoEHASUStN3vUuK" || s.manifest.Figma.PageID != "129:5215" {
		t.Fatalf("figma source = %+v", s.manifest.Figma)
	}
}

func (s *figmaCoverageTestScope) verifyManifestPolicy(t *testing.T) {
	t.Helper()
	policy := s.manifest.CoveragePolicy
	if strings.TrimSpace(policy.TopDown) == "" || strings.TrimSpace(policy.BottomUp) == "" {
		t.Fatalf("coverage policy must name top-down and bottom-up loops: %+v", policy)
	}
	if !strings.Contains(s.docText, "Figma is evidence") {
		t.Fatalf("coverage doc must say Figma is evidence, not durable SSOT")
	}
	verifyFigmaCoverageInspectionMethod(t, s.manifest.InspectionMethod, s.docText)
	verifyFigmaSupportingToolLimitations(t, s.manifest.SupportingToolLimitations, s.docText)
	verifyFigmaAPIGeneratedAnnotationContentPolicy(t, s.manifest.APIGeneratedAnnotationContentPolicy, s.docText)
}

func (s *figmaCoverageTestScope) verifyManifestCounts(t *testing.T) {
	t.Helper()
	if got, want := len(s.manifest.ExpectedTopLevelNodes), 16; got != want {
		t.Fatalf("expected_top_level_nodes = %d, want %d", got, want)
	}
	if got, want := len(s.manifest.ExpectedPages), 3; got != want {
		t.Fatalf("expected_pages = %d, want %d", got, want)
	}
	if got, want := len(s.manifest.NonUITopLevelNodes), 12; got != want {
		t.Fatalf("non_ui_top_level_nodes = %d, want %d", got, want)
	}
	if got, want := len(s.manifest.Entries), len(s.manifest.ExpectedTopLevelNodes); got != want {
		t.Fatalf("entries = %d, want %d", got, want)
	}
}
