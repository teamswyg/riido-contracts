package main

import (
	"encoding/json"
	"os"
)

func newEvidence(model model) evidence {
	m := model.Manifest
	return evidence{
		SchemaVersion:          "riido-contracts-ir-event-log-evidence.v1",
		ID:                     m.ID,
		Status:                 "verified",
		GeneratedDoc:           m.GeneratedDoc,
		Package:                m.Package,
		Workflow:               m.Workflow,
		EvidenceArtifact:       m.EvidenceArtifact,
		EventCount:             model.EventCount,
		TransitionCount:        model.TransitionCount,
		NonTransitionCount:     model.NonTransitionCount,
		TaskFSMTransitionCount: model.TaskFSMTransitionCount,
		TaskFSMTriggerCount:    model.TaskFSMTriggerCount,
		CanonicalEventFields:   model.CanonicalEventFields,
		ReduceResultFields:     model.ReduceResultFields,
		NativeConfigCounts:     model.NativeConfigCounts,
		Loop:                   m.Loop,
	}
}

func writeEvidence(path string, model model) error {
	body, err := json.MarshalIndent(newEvidence(model), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o644)
}
