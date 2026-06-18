package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDefaultManifestVerifies(t *testing.T) {
	var out bytes.Buffer
	if err := run([]string{"verify"}, &out); err != nil {
		t.Fatalf("verify default manifest: %v", err)
	}
	if got := out.String(); !strings.Contains(got, "verified 21 facts and 4 repo dependencies") {
		t.Fatalf("unexpected output: %q", got)
	}
}
