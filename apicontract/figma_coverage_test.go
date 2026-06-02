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
	verifyFigmaCoverageProvenance(t, manifest.StabilizedBy, docPath)
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
	verifyFigmaCoverageInspectionMethod(t, manifest.InspectionMethod, docText)
	verifyFigmaSupportingToolLimitations(t, manifest.SupportingToolLimitations, docText)
	verifyFigmaAPIGeneratedAnnotationContentPolicy(t, manifest.APIGeneratedAnnotationContentPolicy, docText)

	if got, want := len(manifest.ExpectedTopLevelNodes), 16; got != want {
		t.Fatalf("expected_top_level_nodes = %d, want %d", got, want)
	}
	if got, want := len(manifest.ExpectedPages), 3; got != want {
		t.Fatalf("expected_pages = %d, want %d", got, want)
	}
	if got, want := len(manifest.NonUITopLevelNodes), 12; got != want {
		t.Fatalf("non_ui_top_level_nodes = %d, want %d", got, want)
	}
	if got, want := len(manifest.Entries), len(manifest.ExpectedTopLevelNodes); got != want {
		t.Fatalf("entries = %d, want %d", got, want)
	}
	openAPIGeneratedPaths := loadAIAgentClientGeneratedPaths(t)

	pages := map[string]figmaCoveragePage{}
	for _, page := range manifest.ExpectedPages {
		if page.NodeID == "" || page.Name == "" || page.ChildCount <= 0 {
			t.Fatalf("expected page has invalid field: %+v", page)
		}
		if _, exists := pages[page.NodeID]; exists {
			t.Fatalf("duplicate expected page %q", page.NodeID)
		}
		pages[page.NodeID] = page
	}
	if _, ok := pages[manifest.Figma.PageID]; !ok {
		t.Fatalf("primary figma page %q is missing from expected_pages", manifest.Figma.PageID)
	}
	if pages[manifest.Figma.PageID].ChildCount != len(manifest.ExpectedTopLevelNodes) {
		t.Fatalf("primary page child_count = %d, expected_top_level_nodes = %d", pages[manifest.Figma.PageID].ChildCount, len(manifest.ExpectedTopLevelNodes))
	}

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
	for _, page := range manifest.ExpectedPages {
		registerFigmaNode(t, registered, figmaCoverageNode{NodeID: page.NodeID, Name: page.Name}, "expected_pages")
	}
	for _, node := range manifest.ExpectedTopLevelNodes {
		registerFigmaNode(t, registered, node, "expected_top_level_nodes")
	}
	for _, node := range manifest.VerifiedEvidenceNodes {
		registerFigmaNode(t, registered, node, "verified_evidence_nodes")
	}
	nonUIInventory := verifyNonUITopLevelInventory(t, manifest, pages)

	seen := map[string]bool{}
	entryByNodeID := map[string]figmaCoverageEntry{}
	for i, entry := range manifest.Entries {
		expectedNode, ok := expected[entry.NodeID]
		if !ok {
			t.Fatalf("entry %q is not in expected_top_level_nodes", entry.NodeID)
		}
		if seen[entry.NodeID] {
			t.Fatalf("duplicate entry node_id %q", entry.NodeID)
		}
		seen[entry.NodeID] = true
		entryByNodeID[entry.NodeID] = entry
		if entry.Name != expectedNode.Name {
			t.Fatalf("entry %q name = %q, want %q", entry.NodeID, entry.Name, expectedNode.Name)
		}
		assertCoverageDocMentionsEntry(t, docText, entry)
		if manifest.ExpectedTopLevelNodes[i].NodeID != entry.NodeID {
			t.Fatalf("entry order must match expected_top_level_nodes at %d: got %s want %s", i, entry.NodeID, manifest.ExpectedTopLevelNodes[i].NodeID)
		}
		verifyCoverageEntry(t, entry, openAPIGeneratedPaths)
	}
	nonUISeen := map[string]bool{}
	for _, entry := range manifest.NonUITopLevelNodes {
		if _, ok := pages[entry.PageID]; !ok {
			t.Fatalf("non-UI entry %q references unknown page %q", entry.NodeID, entry.PageID)
		}
		if entry.PageID == manifest.Figma.PageID {
			t.Fatalf("non-UI entry %q must not reference primary UI page", entry.NodeID)
		}
		if nonUISeen[entry.NodeID] {
			t.Fatalf("duplicate non-UI entry node_id %q", entry.NodeID)
		}
		nonUISeen[entry.NodeID] = true
		if _, ok := nonUIInventory[entry.PageID][entry.NodeID]; !ok {
			t.Fatalf("non-UI coverage entry %q is missing from loaded top-level inventory for page %q", entry.NodeID, entry.PageID)
		}
		registerFigmaNode(t, registered, figmaCoverageNode{NodeID: entry.NodeID, Name: entry.Name}, "non_ui_top_level_nodes")
		assertCoverageDocMentionsEntry(t, docText, entry)
		verifyCoverageEntry(t, entry, openAPIGeneratedPaths)
	}
	verifySemanticLegacyWireframeCoverage(t, manifest.NonUITopLevelNodes, nonUIInventory, entryByNodeID, docText)
	for pageID, nodes := range nonUIInventory {
		for nodeID, node := range nodes {
			if nonUISeen[nodeID] {
				continue
			}
			registerFigmaNodeIfAbsent(t, registered, node, "non_ui_top_level_inventory page "+pageID)
		}
	}

	for _, node := range manifest.ExpectedTopLevelNodes {
		if !seen[node.NodeID] {
			t.Fatalf("expected node %q has no entry", node.NodeID)
		}
	}

	verifyFigmaRuntimeEndpointLabel(t, manifest.VerifiedEvidenceNodes, entryByNodeID["162:23090"], docText)
	verifyFigmaAPIGeneratedAnnotations(t, manifest.APIGeneratedAnnotations, docText, openAPIGeneratedPaths, registered, entryByNodeID)
	verifyFigmaAPIGeneratedAnnotationInventory(t, manifest.APIGeneratedAnnotationInventory, docText, openAPIGeneratedPaths, registered, entryByNodeID)
	assertDocumentedFigmaNodeRefsAreRegistered(t, registered)
	assertNoStaleOnboardingFixtureWording(t)
	assertNoStaleRuntimeEndpointHostPinned(t)
}

func verifyNonUITopLevelInventory(t *testing.T, manifest figmaCoverageManifest, pages map[string]figmaCoveragePage) map[string]map[string]figmaCoverageNode {
	t.Helper()
	if len(manifest.NonUITopLevelInventory) == 0 {
		t.Fatalf("non_ui_top_level_inventory must record loaded non-UI page children")
	}
	inventory := map[string]map[string]figmaCoverageNode{}
	for _, pageInventory := range manifest.NonUITopLevelInventory {
		page, ok := pages[pageInventory.PageID]
		if !ok {
			t.Fatalf("non-UI inventory references unknown page %q", pageInventory.PageID)
		}
		if pageInventory.PageID == manifest.Figma.PageID {
			t.Fatalf("non-UI inventory must not reference primary UI page %q", pageInventory.PageID)
		}
		if _, exists := inventory[pageInventory.PageID]; exists {
			t.Fatalf("duplicate non-UI inventory page %q", pageInventory.PageID)
		}
		if got, want := len(pageInventory.Nodes), page.ChildCount; got != want {
			t.Fatalf("non-UI inventory page %q nodes = %d, want loaded child_count %d", pageInventory.PageID, got, want)
		}
		nodes := map[string]figmaCoverageNode{}
		for _, node := range pageInventory.Nodes {
			if strings.TrimSpace(node.NodeID) == "" || strings.TrimSpace(node.Name) == "" {
				t.Fatalf("non-UI inventory page %q has invalid node: %+v", pageInventory.PageID, node)
			}
			if _, exists := nodes[node.NodeID]; exists {
				t.Fatalf("duplicate non-UI inventory node %q on page %q", node.NodeID, pageInventory.PageID)
			}
			nodes[node.NodeID] = node
		}
		inventory[pageInventory.PageID] = nodes
	}
	for _, page := range pages {
		if page.NodeID == manifest.Figma.PageID {
			continue
		}
		if _, ok := inventory[page.NodeID]; !ok {
			t.Fatalf("non-UI page %q is missing loaded top-level inventory", page.NodeID)
		}
	}
	return inventory
}

func verifyFigmaCoverageInspectionMethod(t *testing.T, method figmaCoverageInspectionMethod, docText string) {
	t.Helper()
	if method.ID != "figma-plugin-api-page-registry.v1" {
		t.Fatalf("inspection_method.id = %q", method.ID)
	}
	if strings.TrimSpace(method.Authority) != "Figma Plugin API via use_figma" {
		t.Fatalf("inspection_method.authority = %q", method.Authority)
	}
	if method.PageRegistryExpression != "figma.root.children" {
		t.Fatalf("inspection_method.page_registry_expression = %q", method.PageRegistryExpression)
	}
	if method.TopLevelChildCountExpression != "await figma.setCurrentPageAsync(page); page.children.length" {
		t.Fatalf("inspection_method.top_level_child_count_expression = %q", method.TopLevelChildCountExpression)
	}
	if len(method.SupportingTools) == 0 {
		t.Fatalf("inspection_method.supporting_tools must name non-authoritative read tools")
	}
	rule := strings.ToLower(method.Rule)
	for _, needle := range []string{"metadata", "supporting evidence", "must not redefine page-level child counts", "lazy/unloaded"} {
		if !strings.Contains(rule, needle) {
			t.Fatalf("inspection_method.rule must contain %q: %q", needle, method.Rule)
		}
	}
	for _, needle := range []string{"figma.root.children", "await figma.setCurrentPageAsync(page)", "page.children.length", "Metadata XML/read", "supporting evidence only", "lazy/unloaded"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe inspection method with %q", needle)
		}
	}
}

func verifyFigmaSupportingToolLimitations(t *testing.T, limitations []figmaSupportingToolLimitation, docText string) {
	t.Helper()
	if len(limitations) == 0 {
		t.Fatalf("supporting_tool_limitations must record non-authoritative tooling failure modes")
	}
	var metadataPageList figmaSupportingToolLimitation
	var headlessFileKey figmaSupportingToolLimitation
	var onboardingPageLoadTimeout figmaSupportingToolLimitation
	for _, limitation := range limitations {
		if limitation.ID == "figma-metadata-page-list-underreports-pages.v1" {
			metadataPageList = limitation
		}
		if limitation.ID == "figma-headless-file-key-placeholder.v1" {
			headlessFileKey = limitation
		}
		if limitation.ID == "figma-onboarding-page-load-timeout.v1" {
			onboardingPageLoadTimeout = limitation
		}
	}
	if metadataPageList.ID == "" {
		t.Fatalf("supporting_tool_limitations must include figma-metadata-page-list-underreports-pages.v1")
	}
	if !strings.Contains(metadataPageList.Tool, "get_metadata") || !strings.Contains(metadataPageList.Tool, "without nodeId") {
		t.Fatalf("metadata limitation tool must name no-nodeId get_metadata: %+v", metadataPageList)
	}
	for _, needle := range []string{"only page 129:5215 UI", "MUOd9lctoEHASUStN3vUuK"} {
		if !strings.Contains(metadataPageList.ObservedResult, needle) {
			t.Fatalf("metadata limitation observed_result must contain %q: %q", needle, metadataPageList.ObservedResult)
		}
	}
	requiredPages := map[string]bool{"129:5215": false, "42:3014": false, "0:1": false}
	for _, pageID := range metadataPageList.AuthoritativeResult {
		if _, ok := requiredPages[pageID]; ok {
			requiredPages[pageID] = true
		}
	}
	for pageID, seen := range requiredPages {
		if !seen {
			t.Fatalf("metadata limitation authoritative_result is missing page %s: %+v", pageID, metadataPageList.AuthoritativeResult)
		}
	}
	rule := strings.ToLower(metadataPageList.Rule)
	for _, needle := range []string{"supporting evidence only", "must not remove expected_pages", "non-ui inventories"} {
		if !strings.Contains(rule, needle) {
			t.Fatalf("metadata limitation rule must contain %q: %q", needle, metadataPageList.Rule)
		}
	}
	for _, needle := range []string{"figma-metadata-page-list-underreports-pages.v1", "get_metadata", "without `nodeId`", "`129:5215`", "`42:3014`", "`0:1`", "must not remove `expected_pages`"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe metadata page-list limitation with %q", needle)
		}
	}
	if headlessFileKey.ID == "" {
		t.Fatalf("supporting_tool_limitations must include figma-headless-file-key-placeholder.v1")
	}
	for _, needle := range []string{"use_figma", "figma.fileKey"} {
		if !strings.Contains(headlessFileKey.Tool, needle) {
			t.Fatalf("headless file-key limitation tool must contain %q: %+v", needle, headlessFileKey)
		}
	}
	for _, needle := range []string{"MUOd9lctoEHASUStN3vUuK", "figma.fileKey=headless", "pages and annotation categories"} {
		if !strings.Contains(headlessFileKey.ObservedResult, needle) {
			t.Fatalf("headless file-key observed_result must contain %q: %q", needle, headlessFileKey.ObservedResult)
		}
	}
	for _, needle := range []string{"MUOd9lctoEHASUStN3vUuK", "v.1.22 AI Agent"} {
		if !stringSliceContains(headlessFileKey.AuthoritativeResult, needle) {
			t.Fatalf("headless file-key authoritative_result must contain %q: %+v", needle, headlessFileKey.AuthoritativeResult)
		}
	}
	headlessRule := strings.ToLower(headlessFileKey.Rule)
	for _, needle := range []string{"supporting evidence only", "must not overwrite figma.file_key", "expected_pages", "downstream projection source identity"} {
		if !strings.Contains(headlessRule, needle) {
			t.Fatalf("headless file-key rule must contain %q: %q", needle, headlessFileKey.Rule)
		}
	}
	for _, needle := range []string{"figma-headless-file-key-placeholder.v1", "`figma.fileKey=headless`", "`MUOd9lctoEHASUStN3vUuK`", "authoritative file identity", "must not overwrite `figma.file_key`"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe headless file-key limitation with %q", needle)
		}
	}
	if onboardingPageLoadTimeout.ID == "" {
		t.Fatalf("supporting_tool_limitations must include figma-onboarding-page-load-timeout.v1")
	}
	for _, needle := range []string{"get_metadata", "42:3014", "Plugin API page load"} {
		if !strings.Contains(onboardingPageLoadTimeout.Tool, needle) {
			t.Fatalf("onboarding load timeout tool must contain %q: %+v", needle, onboardingPageLoadTimeout)
		}
	}
	for _, needle := range []string{"time out after 120s", "Wireframe - 온보딩", "setCurrentPageAsync", "236:33845", "236:33847", "six onboarding riido.* API Generated annotations"} {
		if !strings.Contains(onboardingPageLoadTimeout.ObservedResult, needle) {
			t.Fatalf("onboarding load timeout observed_result must contain %q: %q", needle, onboardingPageLoadTimeout.ObservedResult)
		}
	}
	for _, needle := range []string{"42:3014", "child_count=83", "non_ui_top_level_inventory", "236:33845", "236:33847", "onboarding_api_generated_annotations=6"} {
		if !stringSliceContains(onboardingPageLoadTimeout.AuthoritativeResult, needle) {
			t.Fatalf("onboarding load timeout authoritative_result must contain %q: %+v", needle, onboardingPageLoadTimeout.AuthoritativeResult)
		}
	}
	onboardingRule := strings.ToLower(onboardingPageLoadTimeout.Rule)
	for _, needle := range []string{"supporting evidence only", "must not rewrite expected_pages", "remove page 42:3014", "onboarding generated paths", "direct registered-node lookup"} {
		if !strings.Contains(onboardingRule, needle) {
			t.Fatalf("onboarding load timeout rule must contain %q: %q", needle, onboardingPageLoadTimeout.Rule)
		}
	}
	for _, needle := range []string{"figma-onboarding-page-load-timeout.v1", "get_metadata(nodeId=42:3014)", "after 120s", "`Wireframe - 온보딩`", "`236:33845`", "`236:33847`", "six onboarding `riido.*` `API Generated`", "must not rewrite `expected_pages`", "onboarding generated paths unresolved"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe onboarding page load timeout with %q", needle)
		}
	}
}

func verifyFigmaCoverageProvenance(t *testing.T, stabilizedBy []string, docPath string) {
	t.Helper()
	want := []string{
		"teamswyg/riido-contracts#38",
		"teamswyg/riido-contracts#39",
		"teamswyg/riido-contracts#45",
		"teamswyg/riido-contracts#46",
		"teamswyg/riido-contracts#51",
		"teamswyg/riido-contracts#52",
		"teamswyg/riido-contracts#54",
		"teamswyg/riido-contracts#55",
		"teamswyg/riido-contracts#56",
		"teamswyg/riido-contracts#57",
		"teamswyg/riido-contracts#58",
		"teamswyg/riido-contracts#60",
		"teamswyg/riido-contracts#62",
		"teamswyg/riido-contracts#63",
		"teamswyg/riido-contracts#64",
	}
	if len(stabilizedBy) != len(want) {
		t.Fatalf("stabilized_by = %d entries, want %d: %+v", len(stabilizedBy), len(want), stabilizedBy)
	}
	for i := range want {
		if stabilizedBy[i] != want[i] {
			t.Fatalf("stabilized_by[%d] = %q, want %q; full list = %+v", i, stabilizedBy[i], want[i], stabilizedBy)
		}
	}
	doc, err := os.ReadFile(docPath)
	if err != nil {
		t.Fatalf("read coverage doc for provenance: %v", err)
	}
	docText := string(doc)
	for _, pr := range want {
		if !strings.Contains(docText, pr) {
			t.Fatalf("coverage doc must mention stabilization provenance %q", pr)
		}
	}
	if !strings.Contains(docText, "`stabilized_by`") ||
		!strings.Contains(docText, "downstream projection") {
		t.Fatalf("coverage doc must explain stabilized_by as the downstream projection mirror source")
	}
}

func verifySemanticLegacyWireframeCoverage(t *testing.T, entries []figmaCoverageEntry, inventory map[string]map[string]figmaCoverageNode, primaryEntries map[string]figmaCoverageEntry, docText string) {
	t.Helper()
	required := map[string]struct {
		name       string
		absorbedBy string
	}{
		"13:3789": {"런타임", "162:23090"},
		"86:9988": {"런타임", "162:23090"},
		"17:3551": {"에이전트", "432:37336"},
		"17:4231": {"에이전트 수정", "432:37336"},
		"84:9846": {"에이전트 추가", "432:37336"},
		"17:2871": {"데몬 상세", "162:23090"},
		"17:3111": {"런타임 상세", "162:23090"},
	}
	byNode := map[string]figmaCoverageEntry{}
	for _, entry := range entries {
		byNode[entry.NodeID] = entry
	}
	for nodeID, want := range required {
		entry, ok := byNode[nodeID]
		if !ok {
			t.Fatalf("semantic legacy Wireframe node %s %s must be promoted from inventory to non_ui_top_level_nodes", nodeID, want.name)
		}
		if entry.PageID != "0:1" {
			t.Fatalf("semantic legacy node %s page_id = %q, want 0:1", nodeID, entry.PageID)
		}
		if entry.Name != want.name {
			t.Fatalf("semantic legacy node %s name = %q, want %q", nodeID, entry.Name, want.name)
		}
		if entry.CoverageStatus != "covered" {
			t.Fatalf("semantic legacy node %s coverage_status = %q, want covered", nodeID, entry.CoverageStatus)
		}
		if entry.EvidenceKind != "figma_legacy_wireframe_section" {
			t.Fatalf("semantic legacy node %s evidence_kind = %q, want figma_legacy_wireframe_section", nodeID, entry.EvidenceKind)
		}
		if _, ok := inventory["0:1"][nodeID]; !ok {
			t.Fatalf("semantic legacy node %s is not present in loaded Wireframe inventory", nodeID)
		}
		if entry.AbsorbedByTopLevelNodeID != want.absorbedBy {
			t.Fatalf("semantic legacy node %s absorbed_by_top_level_node_id = %q, want %q", nodeID, entry.AbsorbedByTopLevelNodeID, want.absorbedBy)
		}
		absorbed, ok := primaryEntries[want.absorbedBy]
		if !ok {
			t.Fatalf("semantic legacy node %s absorbs into missing primary UI entry %s", nodeID, want.absorbedBy)
		}
		if absorbed.CoverageStatus != "covered" {
			t.Fatalf("semantic legacy node %s absorbs into non-covered primary entry %s", nodeID, want.absorbedBy)
		}
		if len(entry.GeneratedPaths) == 0 {
			t.Fatalf("semantic legacy node %s must name the generated paths inherited from its current UI entry", nodeID)
		}
		for _, generatedPath := range entry.GeneratedPaths {
			if !stringSliceContains(absorbed.GeneratedPaths, generatedPath) {
				t.Fatalf("semantic legacy node %s generated path %q is not covered by absorbed primary UI entry %s", nodeID, generatedPath, want.absorbedBy)
			}
		}
		facts := strings.Join(entry.CoveredFacts, "\n")
		if !strings.Contains(facts, "absorbed by the current UI") {
			t.Fatalf("semantic legacy node %s must explain current UI absorption: %q", nodeID, facts)
		}
		if !strings.Contains(docText, nodeID) || !strings.Contains(docText, want.absorbedBy) {
			t.Fatalf("coverage doc must mention semantic legacy node %s and absorbed primary entry %s", nodeID, want.absorbedBy)
		}
	}
}

func assertCoverageDocMentionsEntry(t *testing.T, docText string, entry figmaCoverageEntry) {
	t.Helper()
	if !strings.Contains(docText, entry.NodeID) || !strings.Contains(docText, entry.Name) {
		t.Fatalf("coverage doc must mention node %s %s", entry.NodeID, entry.Name)
	}
	for _, generatedPath := range entry.GeneratedPaths {
		if !docMentionsGeneratedPath(docText, generatedPath) {
			t.Fatalf("coverage doc must mention generated path %q for node %s", generatedPath, entry.NodeID)
		}
	}
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

func verifyFigmaAPIGeneratedAnnotations(t *testing.T, annotations []figmaAPIGeneratedAnnotation, docText string, openAPIGeneratedPaths map[string]string, registered map[string]string, entries map[string]figmaCoverageEntry) {
	t.Helper()
	if got, want := len(annotations), 2; got != want {
		t.Fatalf("api_generated_annotations = %d, want %d", got, want)
	}
	seen := map[string]bool{}
	for _, annotation := range annotations {
		if seen[annotation.NodeID] {
			t.Fatalf("duplicate API Generated annotation node %q", annotation.NodeID)
		}
		seen[annotation.NodeID] = true
		if _, ok := registered[annotation.NodeID]; !ok {
			t.Fatalf("API Generated annotation %q is not a registered Figma evidence node", annotation.NodeID)
		}
		if annotation.TopLevelNodeID != "153:15931" || annotation.CoverageEntryNodeID != "153:15931" {
			t.Fatalf("API Generated annotation %q must resolve through task-thread top-level entry 153:15931: %+v", annotation.NodeID, annotation)
		}
		if annotation.CategoryID != "700:0" || annotation.CategoryLabel != "API Generated" {
			t.Fatalf("API Generated annotation %q category drifted: %+v", annotation.NodeID, annotation)
		}
		if !strings.HasPrefix(annotation.FigmaGeneratedPath, "riido.") {
			t.Fatalf("API Generated annotation %q must preserve the Figma facade path: %q", annotation.NodeID, annotation.FigmaGeneratedPath)
		}
		canonical := strings.TrimPrefix(annotation.FigmaGeneratedPath, "riido.")
		if annotation.CanonicalGeneratedPath != canonical {
			t.Fatalf("API Generated annotation %q canonical path = %q, want %q", annotation.NodeID, annotation.CanonicalGeneratedPath, canonical)
		}
		if _, ok := openAPIGeneratedPaths[annotation.CanonicalGeneratedPath]; !ok {
			t.Fatalf("API Generated annotation %q references unknown OpenAPI generated path %q", annotation.NodeID, annotation.CanonicalGeneratedPath)
		}
		entry, ok := entries[annotation.CoverageEntryNodeID]
		if !ok {
			t.Fatalf("API Generated annotation %q references missing coverage entry %q", annotation.NodeID, annotation.CoverageEntryNodeID)
		}
		if !stringSliceContains(entry.GeneratedPaths, annotation.CanonicalGeneratedPath) {
			t.Fatalf("API Generated annotation %q canonical path %q is not in coverage entry %q generated paths", annotation.NodeID, annotation.CanonicalGeneratedPath, entry.NodeID)
		}
		for _, needle := range []string{annotation.NodeID, annotation.FigmaGeneratedPath, annotation.CanonicalGeneratedPath, annotation.CategoryLabel} {
			if !strings.Contains(docText, needle) {
				t.Fatalf("coverage doc must mention API Generated annotation %q", needle)
			}
		}
		if strings.Contains(annotation.FigmaLabel, "작업중") {
			if annotation.ResolutionStatus != "resolved_stale_handoff_copy" {
				t.Fatalf("API Generated annotation %q stale Figma copy must be explicitly resolved: %+v", annotation.NodeID, annotation)
			}
			if !strings.Contains(annotation.Resolution, "stale") || !strings.Contains(docText, "상세내용은 작업중입니다") {
				t.Fatalf("API Generated annotation %q stale copy resolution is not documented", annotation.NodeID)
			}
		}
	}
}

func verifyFigmaAPIGeneratedAnnotationContentPolicy(t *testing.T, policy figmaAPIGeneratedAnnotationContentRule, docText string) {
	t.Helper()
	if policy.CategoryID != "700:0" || policy.CategoryLabel != "API Generated" {
		t.Fatalf("API Generated annotation content category drifted: %+v", policy)
	}
	if len(policy.LabelFormat) != 3 {
		t.Fatalf("API Generated annotation label_format = %d entries, want 3", len(policy.LabelFormat))
	}
	for _, needle := range []string{"riido.*", "종류", "Query", "Mutation", "SSE Stream", "배경", "Korean"} {
		if !strings.Contains(strings.Join(policy.LabelFormat, "\n")+"\n"+policy.Rule, needle) {
			t.Fatalf("API Generated annotation content policy must mention %q: %+v", needle, policy)
		}
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must mention API Generated annotation content policy %q", needle)
		}
	}
	if !strings.Contains(policy.Rule, "must not become a second API SSOT") {
		t.Fatalf("API Generated annotation content policy must prevent second SSOT drift: %q", policy.Rule)
	}
	verifyFigmaAPIGeneratedRetiredCategories(t, policy.RetiredCategories, docText)
	scan := policy.LiveInspection
	if scan.ObservedAt != "2026-06-02" || !strings.Contains(scan.Tool, "use_figma") {
		t.Fatalf("API Generated annotation live inspection provenance drifted: %+v", scan)
	}
	expected := map[string]figmaAPIGeneratedAnnotationLivePageCounter{
		"129:5215": {
			PageID:               "129:5215",
			PageName:             "UI",
			RiidoAnnotationCount: 53,
			APIGeneratedCount:    53,
		},
		"42:3014": {
			PageID:               "42:3014",
			PageName:             "Wireframe - 온보딩",
			RiidoAnnotationCount: 6,
			APIGeneratedCount:    6,
		},
		"0:1": {
			PageID:               "0:1",
			PageName:             "Wireframe",
			RiidoAnnotationCount: 0,
			APIGeneratedCount:    0,
		},
	}
	if len(scan.PageCounts) != len(expected) {
		t.Fatalf("API Generated annotation page_counts = %d, want %d", len(scan.PageCounts), len(expected))
	}
	var totalRiido, totalAPIGenerated int
	for _, page := range scan.PageCounts {
		want, ok := expected[page.PageID]
		if !ok {
			t.Fatalf("unexpected API Generated annotation live page count: %+v", page)
		}
		if page.PageName != want.PageName || page.RiidoAnnotationCount != want.RiidoAnnotationCount || page.APIGeneratedCount != want.APIGeneratedCount {
			t.Fatalf("API Generated annotation live page count for %s = %+v, want %+v", page.PageID, page, want)
		}
		if page.MissingOperationKind != 0 || page.MissingBackground != 0 {
			t.Fatalf("API Generated annotation live page count has missing content: %+v", page)
		}
		totalRiido += page.RiidoAnnotationCount
		totalAPIGenerated += page.APIGeneratedCount
		for _, needle := range []string{page.PageID, page.PageName} {
			if !strings.Contains(docText, needle) {
				t.Fatalf("coverage doc must mention API Generated annotation live page count %q", needle)
			}
		}
	}
	if scan.TotalRiidoAnnotations != totalRiido || scan.TotalAPIGeneratedAnnotations != totalAPIGenerated {
		t.Fatalf("API Generated annotation live totals = riido:%d/api:%d, want riido:%d/api:%d", scan.TotalRiidoAnnotations, scan.TotalAPIGeneratedAnnotations, totalRiido, totalAPIGenerated)
	}
	if totalRiido != 59 || totalAPIGenerated != 59 {
		t.Fatalf("API Generated annotation live totals = riido:%d/api:%d, want 59/59", totalRiido, totalAPIGenerated)
	}
}

func verifyFigmaAPIGeneratedRetiredCategories(t *testing.T, categories []figmaAPIGeneratedAnnotationRetiredCategory, docText string) {
	t.Helper()
	if len(categories) != 1 {
		t.Fatalf("API Generated retired categories = %d, want 1", len(categories))
	}
	retired := categories[0]
	if retired.CategoryID != "39:0" || retired.CategoryLabel != "클라이언트 전달" {
		t.Fatalf("unexpected retired API Generated category: %+v", retired)
	}
	if retired.RetirementStatus != "unused_not_deleted" || retired.LiveUsageCount != 0 {
		t.Fatalf("retired API Generated category must stay unused_not_deleted with zero live usage: %+v", retired)
	}
	if retired.ObservedAt != "2026-06-02" || !strings.Contains(retired.ToolLimitation, "remove/setLabel") {
		t.Fatalf("retired API Generated category must record automation limitation: %+v", retired)
	}
	for _, needle := range []string{retired.CategoryID, retired.CategoryLabel, "retired", "zero annotations"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must mention retired API Generated category %q", needle)
		}
	}
}

func verifyFigmaAPIGeneratedAnnotationInventory(t *testing.T, inventory []figmaAPIGeneratedAnnotationGroup, docText string, openAPIGeneratedPaths map[string]string, registered map[string]string, entries map[string]figmaCoverageEntry) {
	t.Helper()
	if got, want := len(inventory), 19; got != want {
		t.Fatalf("api_generated_annotation_inventory = %d, want %d", got, want)
	}
	allowedKinds := map[string]bool{"Query": true, "Mutation": true, "SSE Stream": true}
	seenPath := map[string]bool{}
	totalAnnotations := 0
	for _, group := range inventory {
		if strings.TrimSpace(group.UIArea) == "" {
			t.Fatalf("API Generated annotation group has empty ui_area: %+v", group)
		}
		if group.CategoryID != "700:0" || group.CategoryLabel != "API Generated" {
			t.Fatalf("API Generated annotation group %q category drifted: %+v", group.FigmaGeneratedPath, group)
		}
		if !strings.HasPrefix(group.FigmaGeneratedPath, "riido.") {
			t.Fatalf("API Generated annotation group must preserve Figma facade path: %q", group.FigmaGeneratedPath)
		}
		canonical := strings.TrimPrefix(group.FigmaGeneratedPath, "riido.")
		if group.CanonicalGeneratedPath != canonical {
			t.Fatalf("API Generated annotation group %q canonical path = %q, want %q", group.FigmaGeneratedPath, group.CanonicalGeneratedPath, canonical)
		}
		if seenPath[group.CanonicalGeneratedPath] {
			t.Fatalf("duplicate API Generated annotation inventory generated path %q", group.CanonicalGeneratedPath)
		}
		seenPath[group.CanonicalGeneratedPath] = true
		if _, ok := openAPIGeneratedPaths[group.CanonicalGeneratedPath]; !ok {
			t.Fatalf("API Generated annotation group references unknown OpenAPI generated path %q", group.CanonicalGeneratedPath)
		}
		if !allowedKinds[group.OperationKind] {
			t.Fatalf("API Generated annotation group %q operation_kind = %q", group.CanonicalGeneratedPath, group.OperationKind)
		}
		if strings.TrimSpace(group.Background) == "" {
			t.Fatalf("API Generated annotation group %q must explain background", group.CanonicalGeneratedPath)
		}
		if len(group.Sources) == 0 {
			t.Fatalf("API Generated annotation group %q must name sources", group.CanonicalGeneratedPath)
		}
		annotationCount := 0
		for _, source := range group.Sources {
			if strings.TrimSpace(source.PageID) == "" || strings.TrimSpace(source.TopLevelNodeID) == "" || strings.TrimSpace(source.CoverageEntryNodeID) == "" {
				t.Fatalf("API Generated annotation group %q has invalid source: %+v", group.CanonicalGeneratedPath, source)
			}
			entry, ok := entries[source.CoverageEntryNodeID]
			if !ok {
				t.Fatalf("API Generated annotation group %q references missing coverage entry %q", group.CanonicalGeneratedPath, source.CoverageEntryNodeID)
			}
			if !stringSliceContains(entry.GeneratedPaths, group.CanonicalGeneratedPath) {
				t.Fatalf("API Generated annotation group %q canonical path is not covered by source entry %q", group.CanonicalGeneratedPath, source.CoverageEntryNodeID)
			}
			registerFigmaNodeIDIfAbsent(t, registered, source.TopLevelNodeID, "api_generated_annotation_inventory top-level "+group.CanonicalGeneratedPath)
			if len(source.NodeIDs) == 0 {
				t.Fatalf("API Generated annotation group %q source %q must list node_ids", group.CanonicalGeneratedPath, source.TopLevelNodeID)
			}
			sourceSeen := map[string]bool{}
			for _, nodeID := range source.NodeIDs {
				if strings.TrimSpace(nodeID) == "" {
					t.Fatalf("API Generated annotation group %q source has empty node id", group.CanonicalGeneratedPath)
				}
				if sourceSeen[nodeID] {
					t.Fatalf("API Generated annotation group %q source %q duplicates node %q", group.CanonicalGeneratedPath, source.TopLevelNodeID, nodeID)
				}
				sourceSeen[nodeID] = true
				registerFigmaNodeIDIfAbsent(t, registered, nodeID, "api_generated_annotation_inventory "+group.CanonicalGeneratedPath)
				annotationCount++
			}
		}
		if group.AnnotationCount != annotationCount {
			t.Fatalf("API Generated annotation group %q annotation_count = %d, want node count %d", group.CanonicalGeneratedPath, group.AnnotationCount, annotationCount)
		}
		totalAnnotations += annotationCount
		for _, needle := range []string{group.UIArea, group.FigmaGeneratedPath, group.CanonicalGeneratedPath, group.OperationKind, group.Background} {
			if !strings.Contains(docText, needle) {
				t.Fatalf("coverage doc must mention API Generated annotation inventory %q", needle)
			}
		}
	}
	if got, want := totalAnnotations, 59; got != want {
		t.Fatalf("API Generated annotation inventory node annotations = %d, want %d", got, want)
	}
}

func verifyFigmaRuntimeEndpointLabel(t *testing.T, evidence []figmaCoverageNode, runtimeEntry figmaCoverageEntry, docText string) {
	t.Helper()
	var found bool
	for _, node := range evidence {
		if node.NodeID == "129:17930" {
			found = true
			if !strings.Contains(strings.ToLower(node.Name), "endpoint") {
				t.Fatalf("runtime endpoint-looking evidence node must explain its role: %+v", node)
			}
		}
	}
	if !found {
		t.Fatal("runtime settings endpoint-looking label node-id=129:17930 must be registered as verified evidence")
	}
	if runtimeEntry.NodeID != "162:23090" {
		t.Fatalf("runtime settings coverage entry missing: %+v", runtimeEntry)
	}
	facts := strings.Join(runtimeEntry.CoveredFacts, "\n")
	normalizedDocText := strings.Join(strings.Fields(docText), " ")
	for _, needle := range []string{
		"node-id=129:17930",
		"not a canonical base URL",
		"generated path",
		"live host export",
	} {
		if !strings.Contains(facts, needle) {
			t.Fatalf("runtime settings coverage must classify endpoint-looking Figma label with %q: %q", needle, facts)
		}
		if !strings.Contains(normalizedDocText, needle) {
			t.Fatalf("coverage doc must mention runtime endpoint-looking label rule with %q", needle)
		}
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

func stringSliceContains(items []string, want string) bool {
	for _, item := range items {
		if item == want {
			return true
		}
	}
	return false
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

func registerFigmaNodeIfAbsent(t *testing.T, registered map[string]string, node figmaCoverageNode, source string) {
	t.Helper()
	if strings.TrimSpace(node.NodeID) == "" || strings.TrimSpace(node.Name) == "" {
		t.Fatalf("%s has empty node field: %+v", source, node)
	}
	if _, exists := registered[node.NodeID]; exists {
		return
	}
	registered[node.NodeID] = source
}

func registerFigmaNodeIDIfAbsent(t *testing.T, registered map[string]string, nodeID, source string) {
	t.Helper()
	if strings.TrimSpace(nodeID) == "" {
		t.Fatalf("%s has empty node id", source)
	}
	if _, exists := registered[nodeID]; exists {
		return
	}
	registered[nodeID] = source
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

func assertNoStaleRuntimeEndpointHostPinned(t *testing.T) {
	t.Helper()
	forbidden := "desktop-api." + "riido.ai"
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
			assertFileDoesNotContain(t, root, forbidden)
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
				assertFileDoesNotContain(t, path, forbidden)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s for stale runtime endpoint host: %v", root, err)
		}
	}
}

func assertFileDoesNotContain(t *testing.T, path, forbidden string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if strings.Contains(string(data), forbidden) {
		t.Fatalf("%s pins stale endpoint-looking Figma host; cite node-id=129:17930 and explain it is not canonical instead", path)
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
	SchemaVersion                       string                                 `json:"schema_version"`
	ID                                  string                                 `json:"id"`
	RiidoTask                           string                                 `json:"riido_task"`
	StabilizedBy                        []string                               `json:"stabilized_by"`
	HumanDoc                            string                                 `json:"human_doc"`
	RelatedManifests                    []string                               `json:"related_manifests"`
	Figma                               figmaCoverageSource                    `json:"figma"`
	InspectionMethod                    figmaCoverageInspectionMethod          `json:"inspection_method"`
	SupportingToolLimitations           []figmaSupportingToolLimitation        `json:"supporting_tool_limitations"`
	CoveragePolicy                      figmaCoveragePolicy                    `json:"coverage_policy"`
	APIGeneratedAnnotationContentPolicy figmaAPIGeneratedAnnotationContentRule `json:"api_generated_annotation_content_policy"`
	ExpectedPages                       []figmaCoveragePage                    `json:"expected_pages"`
	ExpectedTopLevelNodes               []figmaCoverageNode                    `json:"expected_top_level_nodes"`
	NonUITopLevelInventory              []figmaNonUITopLevelInventory          `json:"non_ui_top_level_inventory"`
	VerifiedEvidenceNodes               []figmaCoverageNode                    `json:"verified_evidence_nodes"`
	NonUITopLevelNodes                  []figmaCoverageEntry                   `json:"non_ui_top_level_nodes"`
	APIGeneratedAnnotations             []figmaAPIGeneratedAnnotation          `json:"api_generated_annotations"`
	APIGeneratedAnnotationInventory     []figmaAPIGeneratedAnnotationGroup     `json:"api_generated_annotation_inventory"`
	Entries                             []figmaCoverageEntry                   `json:"entries"`
}

type figmaCoverageSource struct {
	FileKey          string `json:"file_key"`
	FileName         string `json:"file_name"`
	PageID           string `json:"page_id"`
	PageName         string `json:"page_name"`
	InspectedAt      string `json:"inspected_at"`
	InspectionSource string `json:"inspection_source"`
}

type figmaCoverageInspectionMethod struct {
	ID                           string   `json:"id"`
	Authority                    string   `json:"authority"`
	PageRegistryExpression       string   `json:"page_registry_expression"`
	TopLevelChildCountExpression string   `json:"top_level_child_count_expression"`
	SupportingTools              []string `json:"supporting_tools"`
	Rule                         string   `json:"rule"`
}

type figmaSupportingToolLimitation struct {
	ID                  string   `json:"id"`
	Tool                string   `json:"tool"`
	ObservedAt          string   `json:"observed_at"`
	ObservedResult      string   `json:"observed_result"`
	AuthoritativeSource string   `json:"authoritative_source"`
	AuthoritativeResult []string `json:"authoritative_result"`
	Rule                string   `json:"rule"`
}

type figmaCoveragePolicy struct {
	Summary  string `json:"summary"`
	TopDown  string `json:"top_down"`
	BottomUp string `json:"bottom_up"`
}

type figmaAPIGeneratedAnnotationContentRule struct {
	CategoryID        string                                       `json:"category_id"`
	CategoryLabel     string                                       `json:"category_label"`
	LabelFormat       []string                                     `json:"label_format"`
	Rule              string                                       `json:"rule"`
	RetiredCategories []figmaAPIGeneratedAnnotationRetiredCategory `json:"retired_categories"`
	LiveInspection    figmaAPIGeneratedAnnotationLiveScan          `json:"live_inspection"`
}

type figmaAPIGeneratedAnnotationRetiredCategory struct {
	CategoryID       string `json:"category_id"`
	CategoryLabel    string `json:"category_label"`
	RetirementStatus string `json:"retirement_status"`
	LiveUsageCount   int    `json:"live_usage_count"`
	ObservedAt       string `json:"observed_at"`
	ToolLimitation   string `json:"tool_limitation"`
}

type figmaAPIGeneratedAnnotationLiveScan struct {
	ObservedAt                   string                                       `json:"observed_at"`
	Tool                         string                                       `json:"tool"`
	PageCounts                   []figmaAPIGeneratedAnnotationLivePageCounter `json:"page_counts"`
	TotalRiidoAnnotations        int                                          `json:"total_riido_annotations"`
	TotalAPIGeneratedAnnotations int                                          `json:"total_api_generated_annotations"`
}

type figmaAPIGeneratedAnnotationLivePageCounter struct {
	PageID               string `json:"page_id"`
	PageName             string `json:"page_name"`
	RiidoAnnotationCount int    `json:"riido_annotation_count"`
	APIGeneratedCount    int    `json:"api_generated_count"`
	MissingOperationKind int    `json:"missing_operation_kind"`
	MissingBackground    int    `json:"missing_background"`
}

type figmaCoverageNode struct {
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}

type figmaCoveragePage struct {
	NodeID     string `json:"node_id"`
	Name       string `json:"name"`
	ChildCount int    `json:"child_count"`
}

type figmaNonUITopLevelInventory struct {
	PageID string              `json:"page_id"`
	Nodes  []figmaCoverageNode `json:"nodes"`
}

type figmaCoverageEntry struct {
	NodeID                   string                 `json:"node_id"`
	PageID                   string                 `json:"page_id,omitempty"`
	Name                     string                 `json:"name"`
	CoverageStatus           string                 `json:"coverage_status"`
	EvidenceKind             string                 `json:"evidence_kind"`
	AbsorbedByTopLevelNodeID string                 `json:"absorbed_by_top_level_node_id,omitempty"`
	SSOTDocs                 []string               `json:"ssot_docs,omitempty"`
	OwnerRepos               []string               `json:"owner_repos,omitempty"`
	GeneratedPaths           []string               `json:"generated_paths,omitempty"`
	CoveredFacts             []string               `json:"covered_facts,omitempty"`
	DirectionLoop            figmaCoverageDirection `json:"direction_loop,omitempty"`
	Reason                   string                 `json:"reason,omitempty"`
}

type figmaAPIGeneratedAnnotation struct {
	NodeID                 string `json:"node_id"`
	TopLevelNodeID         string `json:"top_level_node_id"`
	CoverageEntryNodeID    string `json:"coverage_entry_node_id"`
	CategoryID             string `json:"category_id"`
	CategoryLabel          string `json:"category_label"`
	FigmaLabel             string `json:"figma_label"`
	FigmaGeneratedPath     string `json:"figma_generated_path"`
	CanonicalGeneratedPath string `json:"canonical_generated_path"`
	ResolutionStatus       string `json:"resolution_status"`
	Resolution             string `json:"resolution"`
}

type figmaAPIGeneratedAnnotationGroup struct {
	UIArea                 string                              `json:"ui_area"`
	CategoryID             string                              `json:"category_id"`
	CategoryLabel          string                              `json:"category_label"`
	FigmaGeneratedPath     string                              `json:"figma_generated_path"`
	CanonicalGeneratedPath string                              `json:"canonical_generated_path"`
	OperationKind          string                              `json:"operation_kind"`
	Background             string                              `json:"background"`
	AnnotationCount        int                                 `json:"annotation_count"`
	Sources                []figmaAPIGeneratedAnnotationSource `json:"sources"`
}

type figmaAPIGeneratedAnnotationSource struct {
	PageID              string   `json:"page_id"`
	TopLevelNodeID      string   `json:"top_level_node_id"`
	CoverageEntryNodeID string   `json:"coverage_entry_node_id"`
	NodeIDs             []string `json:"node_ids"`
}

type figmaCoverageDirection struct {
	TopDown  string `json:"top_down,omitempty"`
	BottomUp string `json:"bottom_up,omitempty"`
}
