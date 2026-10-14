package apicontract

import (
	"strings"
	"testing"
)

func verifyFigmaAPIGeneratedLiveInspection(t *testing.T, scan figmaAPIGeneratedAnnotationLiveScan, docText string) {
	t.Helper()
	if scan.ObservedAt != "2026-06-03" || !strings.Contains(scan.Tool, "use_figma") || !strings.Contains(scan.Tool, "categoryId") {
		t.Fatalf("API Generated annotation live inspection provenance drifted: %+v", scan)
	}
	expected := expectedFigmaAPIGeneratedLivePageCounts()
	if len(scan.PageCounts) != len(expected) {
		t.Fatalf("API Generated annotation page_counts = %d, want %d", len(scan.PageCounts), len(expected))
	}
	var totalRiido, totalAPIGenerated int
	for _, page := range scan.PageCounts {
		verifyFigmaAPIGeneratedLivePageCount(t, page, expected, docText)
		totalRiido += page.RiidoAnnotationCount
		totalAPIGenerated += page.APIGeneratedCount
	}
	if scan.TotalRiidoAnnotations != totalRiido || scan.TotalAPIGeneratedAnnotations != totalAPIGenerated {
		t.Fatalf("API Generated annotation live totals = riido:%d/api:%d, want riido:%d/api:%d", scan.TotalRiidoAnnotations, scan.TotalAPIGeneratedAnnotations, totalRiido, totalAPIGenerated)
	}
	if totalRiido != 90 || totalAPIGenerated != 90 {
		t.Fatalf("API Generated annotation live totals = riido:%d/api:%d, want 90/90", totalRiido, totalAPIGenerated)
	}
}

func verifyFigmaAPIGeneratedLivePageCount(t *testing.T, page figmaAPIGeneratedAnnotationLivePageCounter, expected map[string]figmaAPIGeneratedAnnotationLivePageCounter, docText string) {
	t.Helper()
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
	for _, needle := range []string{page.PageID, page.PageName} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must mention API Generated annotation live page count %q", needle)
		}
	}
}
