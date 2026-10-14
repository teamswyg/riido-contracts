package main

import (
	"strings"
	"testing"
)

func TestVerifyRejectsDuplicateFactID(t *testing.T) {
	m := minimalManifest(t)
	m.Facts = append(m.Facts, m.Facts[0])
	err := verifyManifest(m, testRoot(t))
	if err == nil || !duplicateFactError(err) {
		t.Fatalf("expected duplicate or sorting error, got %v", err)
	}
}

func TestVerifyRejectsMissingSourcePhrase(t *testing.T) {
	m := minimalManifest(t)
	m.Facts[0].SourceRefs[0].RequiredPhrase = "not present"
	err := verifyManifest(m, testRoot(t))
	if err == nil || !strings.Contains(err.Error(), "does not contain phrase") {
		t.Fatalf("expected missing phrase error, got %v", err)
	}
}

func duplicateFactError(err error) bool {
	return strings.Contains(err.Error(), "facts must be sorted") ||
		strings.Contains(err.Error(), "duplicate fact id")
}
