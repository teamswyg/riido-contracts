package apicontract

import (
	"strings"
	"testing"
)

func (s *figmaAPIGeneratedInventoryScope) verifyGroupPath(t *testing.T, group figmaAPIGeneratedAnnotationGroup) {
	t.Helper()
	if s.seenPath[group.CanonicalGeneratedPath] {
		t.Fatalf("duplicate API Generated annotation inventory generated path %q", group.CanonicalGeneratedPath)
	}
	s.seenPath[group.CanonicalGeneratedPath] = true
	if _, ok := s.openAPIGeneratedPaths[group.CanonicalGeneratedPath]; !ok {
		t.Fatalf("API Generated annotation group references unknown OpenAPI generated path %q", group.CanonicalGeneratedPath)
	}
	transport, ok := s.openAPITransports[group.CanonicalGeneratedPath]
	if !ok {
		t.Fatalf("API Generated annotation group %q has no OpenAPI transport evidence", group.CanonicalGeneratedPath)
	}
	v2Path := "v2." + group.CanonicalGeneratedPath
	v2Transport, ok := s.openAPITransports[v2Path]
	if !ok {
		t.Fatalf("API Generated annotation group %q must keep OpenAPI v2 counterpart %q", group.CanonicalGeneratedPath, v2Path)
	}
	if strings.TrimSpace(group.Background) == "" {
		t.Fatalf("API Generated annotation group %q must explain background", group.CanonicalGeneratedPath)
	}
	verifyInventoryOperationKind(t, group, transport, v2Transport, v2Path)
}

func verifyInventoryOperationKind(t *testing.T, group figmaAPIGeneratedAnnotationGroup, transport, v2Transport figmaOpenAPITransport, v2Path string) {
	t.Helper()
	if !map[string]bool{"Query": true, "Mutation": true, "SSE Stream": true}[group.OperationKind] {
		t.Fatalf("API Generated annotation group %q operation_kind = %q", group.CanonicalGeneratedPath, group.OperationKind)
	}
	if wantKind := operationKindForOpenAPITransport(transport); group.OperationKind != wantKind {
		t.Fatalf("API Generated annotation group %q operation_kind = %q, want %q from OpenAPI transport %+v", group.CanonicalGeneratedPath, group.OperationKind, wantKind, transport)
	}
	if v2Kind := operationKindForOpenAPITransport(v2Transport); group.OperationKind != v2Kind {
		t.Fatalf("API Generated annotation group %q operation_kind = %q, want %q from v2 OpenAPI transport %q %+v", group.CanonicalGeneratedPath, group.OperationKind, v2Kind, v2Path, v2Transport)
	}
}
