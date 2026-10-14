package main

import "errors"

func verifyManifest(root string, m manifest) error {
	if m.SchemaVersion != schemaVersion || !filled(m.ID, m.Title, m.RiidoTask) {
		return errors.New("manifest identity is incomplete")
	}
	if !filled(m.GeneratedDoc, m.Workflow, m.EvidenceArtifact) || len(m.Questions) == 0 {
		return errors.New("generated doc, workflow, artifact, and questions are required")
	}
	if !filled(m.Loop.Observation, m.Loop.Hypothesis, m.Loop.Execute, m.Loop.Evaluate, m.Loop.Retrospective) {
		return errors.New("complete evidence loop is required")
	}
	if err := verifyQuestions(m.Questions); err != nil {
		return err
	}
	return verifyWorkflow(root, m)
}
