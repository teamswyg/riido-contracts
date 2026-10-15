package main

import "testing"

func TestPersistentParticipantRemovalObservesLatestAssignment(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	m, err := readLooseJSONFile[manifest](resolve(root, defaultManifestPath))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{m.DSLFixture, m.IRFixture} {
		fixture, err := readLooseJSONFile[struct {
			Schemas    []schema `json:"schemas"`
			Components []struct {
				Schema schema `json:"schema"`
			} `json:"components"`
		}](resolve(root, path))
		if err != nil {
			t.Fatal(err)
		}
		for _, component := range fixture.Components {
			fixture.Schemas = append(fixture.Schemas, component.Schema)
		}
		record := findSchema(fixture.Schemas, "TaskAgentParticipantRecord")
		if !schemaHasField(record, "latest_assignment_id") {
			t.Fatalf("%s omits the latest assignment guard; a historical waiting execution must not select the removal target", path)
		}
		for _, required := range record.Required {
			if required == "latest_assignment_id" {
				t.Fatalf("%s requires an assignment for a participant whose execution history may be absent", path)
			}
		}
	}
}
