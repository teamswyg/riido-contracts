package apicontract

import (
	"bytes"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestFigmaAIAgentCoverageManifest(t *testing.T) {
	manifestPath := filepath.FromSlash("../docs/30-architecture/figma-ai-agent-coverage.riido.json")
	docPath := filepath.FromSlash("../docs/30-architecture/figma-ai-agent-coverage.md")

	var manifest figmaCoverageManifest
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("read coverage manifest: %v", err)
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&manifest); err != nil {
		t.Fatalf("decode coverage manifest: %v", err)
	}

	if manifest.SchemaVersion != "riido-figma-ai-agent-coverage.v1" {
		t.Fatalf("schema_version = %q", manifest.SchemaVersion)
	}
	if manifest.ID != "figma-v1-22-ai-agent-ui-coverage" {
		t.Fatalf("id = %q", manifest.ID)
	}
	if manifest.RiidoTask != "RIID-4809" {
		t.Fatalf("riido_task = %q", manifest.RiidoTask)
	}
	if manifest.HumanDoc != "docs/30-architecture/figma-ai-agent-coverage.md" {
		t.Fatalf("human_doc = %q", manifest.HumanDoc)
	}
	if manifest.Figma.FileKey != "MUOd9lctoEHASUStN3vUuK" || manifest.Figma.PageID != "129:5215" {
		t.Fatalf("figma source = %+v", manifest.Figma)
	}
	if strings.TrimSpace(manifest.CoveragePolicy.TopDown) == "" || strings.TrimSpace(manifest.CoveragePolicy.BottomUp) == "" {
		t.Fatalf("coverage policy must name top-down and bottom-up loops: %+v", manifest.CoveragePolicy)
	}

	doc, err := os.ReadFile(docPath)
	if err != nil {
		t.Fatalf("read coverage doc: %v", err)
	}
	docText := string(doc)
	if !strings.Contains(docText, "Figma is evidence") {
		t.Fatalf("coverage doc must say Figma is evidence, not durable SSOT")
	}

	if got, want := len(manifest.ExpectedTopLevelNodes), 16; got != want {
		t.Fatalf("expected_top_level_nodes = %d, want %d", got, want)
	}
	if got, want := len(manifest.Entries), len(manifest.ExpectedTopLevelNodes); got != want {
		t.Fatalf("entries = %d, want %d", got, want)
	}
	openAPIGeneratedPaths := loadAIAgentClientGeneratedPaths(t)

	expected := map[string]figmaCoverageNode{}
	for _, node := range manifest.ExpectedTopLevelNodes {
		if node.NodeID == "" || node.Name == "" {
			t.Fatalf("expected node has empty field: %+v", node)
		}
		if _, exists := expected[node.NodeID]; exists {
			t.Fatalf("duplicate expected node %q", node.NodeID)
		}
		expected[node.NodeID] = node
	}
	registered := map[string]string{}
	registerFigmaNode(t, registered, figmaCoverageNode{NodeID: manifest.Figma.PageID, Name: manifest.Figma.PageName}, "figma.page_id")
	for _, node := range manifest.ExpectedTopLevelNodes {
		registerFigmaNode(t, registered, node, "expected_top_level_nodes")
	}
	for _, node := range manifest.VerifiedEvidenceNodes {
		registerFigmaNode(t, registered, node, "verified_evidence_nodes")
	}

	seen := map[string]bool{}
	for i, entry := range manifest.Entries {
		expectedNode, ok := expected[entry.NodeID]
		if !ok {
			t.Fatalf("entry %q is not in expected_top_level_nodes", entry.NodeID)
		}
		if seen[entry.NodeID] {
			t.Fatalf("duplicate entry node_id %q", entry.NodeID)
		}
		seen[entry.NodeID] = true
		if entry.Name != expectedNode.Name {
			t.Fatalf("entry %q name = %q, want %q", entry.NodeID, entry.Name, expectedNode.Name)
		}
		if !strings.Contains(docText, entry.NodeID) || !strings.Contains(docText, entry.Name) {
			t.Fatalf("coverage doc must mention node %s %s", entry.NodeID, entry.Name)
		}
		for _, generatedPath := range entry.GeneratedPaths {
			if !docMentionsGeneratedPath(docText, generatedPath) {
				t.Fatalf("coverage doc must mention generated path %q for node %s", generatedPath, entry.NodeID)
			}
		}
		if manifest.ExpectedTopLevelNodes[i].NodeID != entry.NodeID {
			t.Fatalf("entry order must match expected_top_level_nodes at %d: got %s want %s", i, entry.NodeID, manifest.ExpectedTopLevelNodes[i].NodeID)
		}
		verifyCoverageEntry(t, entry, openAPIGeneratedPaths)
	}

	for _, node := range manifest.ExpectedTopLevelNodes {
		if !seen[node.NodeID] {
			t.Fatalf("expected node %q has no entry", node.NodeID)
		}
	}

	assertDocumentedFigmaNodeRefsAreRegistered(t, registered)
	assertNoStaleOnboardingFixtureWording(t)
}

func verifyCoverageEntry(t *testing.T, entry figmaCoverageEntry, openAPIGeneratedPaths map[string]string) {
	t.Helper()
	if strings.TrimSpace(entry.CoverageStatus) == "" {
		t.Fatalf("entry %q coverage_status is required", entry.NodeID)
	}
	switch entry.CoverageStatus {
	case "covered", "no_diff_product_surface", "planning_evidence":
		if len(entry.SSOTDocs) == 0 {
			t.Fatalf("entry %q must link ssot_docs", entry.NodeID)
		}
		if len(entry.OwnerRepos) == 0 {
			t.Fatalf("entry %q must name owner_repos", entry.NodeID)
		}
		if strings.TrimSpace(entry.DirectionLoop.TopDown) == "" || strings.TrimSpace(entry.DirectionLoop.BottomUp) == "" {
			t.Fatalf("entry %q must define both direction loops", entry.NodeID)
		}
		for _, doc := range entry.SSOTDocs {
			assertCoverageLocalRefExists(t, doc)
		}
		for _, generatedPath := range entry.GeneratedPaths {
			if _, ok := openAPIGeneratedPaths[generatedPath]; !ok {
				t.Fatalf("entry %q references unknown generated path %q", entry.NodeID, generatedPath)
			}
		}
	case "non_decision_asset":
		if strings.TrimSpace(entry.Reason) == "" {
			t.Fatalf("non-decision entry %q must explain reason", entry.NodeID)
		}
		if len(entry.SSOTDocs) != 0 || len(entry.OwnerRepos) != 0 {
			t.Fatalf("non-decision entry %q must not invent owners or SSOT docs", entry.NodeID)
		}
	default:
		t.Fatalf("entry %q has unknown coverage_status %q", entry.NodeID, entry.CoverageStatus)
	}
}

func loadAIAgentClientGeneratedPaths(t *testing.T) map[string]string {
	t.Helper()
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	out := map[string]string{}
	for path, methods := range openAPI.Paths {
		for method, operation := range methods {
			if operation.RiidoClient == nil || strings.TrimSpace(operation.RiidoClient.GeneratedPath) == "" {
				continue
			}
			out[operation.RiidoClient.GeneratedPath] = strings.ToUpper(method) + " " + path
		}
	}
	return out
}

func docMentionsGeneratedPath(docText, generatedPath string) bool {
	if strings.Contains(docText, generatedPath) {
		return true
	}
	lastDot := strings.LastIndex(generatedPath, ".")
	if lastDot < 0 {
		return false
	}
	return strings.Contains(docText, generatedPath[:lastDot]+".*")
}

func assertCoverageLocalRefExists(t *testing.T, ref string) {
	t.Helper()
	path := ref
	if before, _, ok := strings.Cut(ref, "#"); ok {
		path = before
	}
	if strings.TrimSpace(path) == "" {
		t.Fatalf("empty local ref in %q", ref)
	}
	if _, err := os.Stat(filepath.FromSlash("../" + path)); err != nil {
		t.Fatalf("local ref %q does not exist: %v", ref, err)
	}
}

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

func assertNoStaleOnboardingFixtureWording(t *testing.T) {
	t.Helper()
	forbidden := []string{
		"starter-agent",
		"starter agent",
		"starter agents",
		"starter row",
		"starter rows",
		"starter fixture",
		"starter fixtures",
		"starter-fixture",
	}
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
			assertNoStaleOnboardingFixtureWordingInFile(t, root, forbidden)
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
				assertNoStaleOnboardingFixtureWordingInFile(t, path, forbidden)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s for stale onboarding fixture wording: %v", root, err)
		}
	}
}

func assertNoStaleOnboardingFixtureWordingInFile(t *testing.T, path string, forbidden []string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	haystack := strings.ToLower(string(data))
	for _, phrase := range forbidden {
		if strings.Contains(haystack, phrase) {
			t.Fatalf("%s contains stale onboarding fixture wording %q; use onboarding fixture wording instead", path, phrase)
		}
	}
}

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

var (
	figmaNodeIDRefPattern      = regexp.MustCompile(`node-id=([A-Za-z0-9%][A-Za-z0-9:;%_-]*)`)
	numericFigmaURLNodePattern = regexp.MustCompile(`^([0-9]+)-([0-9]+)$`)
)

type figmaCoverageManifest struct {
	SchemaVersion         string               `json:"schema_version"`
	ID                    string               `json:"id"`
	RiidoTask             string               `json:"riido_task"`
	HumanDoc              string               `json:"human_doc"`
	RelatedManifests      []string             `json:"related_manifests"`
	Figma                 figmaCoverageSource  `json:"figma"`
	CoveragePolicy        figmaCoveragePolicy  `json:"coverage_policy"`
	ExpectedTopLevelNodes []figmaCoverageNode  `json:"expected_top_level_nodes"`
	VerifiedEvidenceNodes []figmaCoverageNode  `json:"verified_evidence_nodes"`
	Entries               []figmaCoverageEntry `json:"entries"`
}

type figmaCoverageSource struct {
	FileKey          string `json:"file_key"`
	FileName         string `json:"file_name"`
	PageID           string `json:"page_id"`
	PageName         string `json:"page_name"`
	InspectedAt      string `json:"inspected_at"`
	InspectionSource string `json:"inspection_source"`
}

type figmaCoveragePolicy struct {
	Summary  string `json:"summary"`
	TopDown  string `json:"top_down"`
	BottomUp string `json:"bottom_up"`
}

type figmaCoverageNode struct {
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}

type figmaCoverageEntry struct {
	NodeID         string                 `json:"node_id"`
	Name           string                 `json:"name"`
	CoverageStatus string                 `json:"coverage_status"`
	EvidenceKind   string                 `json:"evidence_kind"`
	SSOTDocs       []string               `json:"ssot_docs,omitempty"`
	OwnerRepos     []string               `json:"owner_repos,omitempty"`
	GeneratedPaths []string               `json:"generated_paths,omitempty"`
	CoveredFacts   []string               `json:"covered_facts,omitempty"`
	DirectionLoop  figmaCoverageDirection `json:"direction_loop,omitempty"`
	Reason         string                 `json:"reason,omitempty"`
}

type figmaCoverageDirection struct {
	TopDown  string `json:"top_down,omitempty"`
	BottomUp string `json:"bottom_up,omitempty"`
}
