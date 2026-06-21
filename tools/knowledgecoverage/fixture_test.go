package main

import (
	"path/filepath"
	"testing"
)

func fixtureRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "go.mod"), "module fixture\n")
	mustWrite(t, filepath.Join(root, "docs", "manual.md"), "# Manual\n")
	mustWrite(t, filepath.Join(root, "docs", "generated.md"), "# Generated\n\n<!-- Code generated; DO NOT EDIT. -->\n")
	mustWrite(t, filepath.Join(root, "docs", "executable-knowledge.riido.json"), fixtureManifest())
	mustWrite(t, filepath.Join(root, "docs", "loopless.riido.json"), `{"schema_version":"fixture.v1"}`)
	mustWrite(t, filepath.Join(root, ".github", "workflows", "executable-knowledge-coverage.yml"), fixtureWorkflow())
	return root
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := writeFile(path, []byte(body)); err != nil {
		t.Fatal(err)
	}
}

func fixtureManifest() string {
	return `{"schema_version":"riido-contracts-executable-knowledge-coverage.v1","id":"fixture","title":"Fixture","generated_doc":"docs/executable-knowledge.md","workflow":".github/workflows/executable-knowledge-coverage.yml","evidence_artifact":"executable-knowledge-coverage","scan_roots":["docs"],"generated_markers":["Code generated"],"loop":{"observation":"o","hypothesis":"h","execute":"x","evaluate":"e","retrospective":"r"}}`
}

func fixtureWorkflow() string {
	return `steps:
- run: go run ./tools/knowledgecoverage -check-doc -evidence-out out/executable-knowledge-coverage.json
- uses: actions/upload-artifact@v4
  with:
    name: executable-knowledge-coverage
    path: out/executable-knowledge-coverage.json
    if-no-files-found: error
`
}
