package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func fixtureRepo(t *testing.T, manifestBody string) string {
	t.Helper()
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "go.mod"), "module fixture\n")
	mustWrite(t, filepath.Join(root, "docs", "50-roadmap", "open-questions.riido.json"), manifestBody)
	mustWrite(t, filepath.Join(root, ".github", "workflows", "open-questions.yml"), fixtureWorkflow())
	return root
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := writeFile(path, []byte(body)); err != nil {
		t.Fatal(err)
	}
}

func fixtureWorkflow() string {
	return "steps:\n- run: go run ./tools/openquestions -check-doc -evidence-out out/open-questions-evidence.json\n- uses: actions/upload-artifact@v4\n  with:\n    name: open-questions-evidence\n    path: out/open-questions-evidence.json\n    if-no-files-found: error\n"
}

func fixtureManifest() string {
	return `{"schema_version":"riido-contracts-open-questions.v1","id":"open","title":"Open","riido_task":"RIID","generated_doc":"docs/50-roadmap/open-questions.md","workflow":".github/workflows/open-questions.yml","evidence_artifact":"open-questions-evidence","questions":[{"id":"Q1","area":"a","status":"open","question":"q","current_stance":"s","next_artifact":"n"},{"id":"Q2","area":"a","status":"resolved","question":"q","current_stance":"s","next_artifact":""}],"loop":{"observation":"o","hypothesis":"h","execute":"x","evaluate":"e","retrospective":"r"}}`
}

func duplicateFixtureManifest() string {
	return strings.Replace(fixtureManifest(), `"id":"Q2"`, `"id":"Q1"`, 1)
}
