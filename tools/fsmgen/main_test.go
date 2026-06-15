package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestVerifyGeneratedFSMFiles(t *testing.T) {
	if err := run([]string{"verify"}, io.Discard); err != nil {
		t.Fatal(err)
	}
}

func TestConformanceCommand(t *testing.T) {
	var out bytes.Buffer
	if err := run([]string{"conformance"}, &out); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "conformance verified 2 FSMs against 1 profile(s)") {
		t.Fatalf("conformance output = %q", out.String())
	}
}
