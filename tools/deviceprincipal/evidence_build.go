package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:         "riido-device-principal-evidence.v1",
		ID:                    m.ID,
		Status:                "verified",
		GeneratedDoc:          m.GeneratedDoc,
		Package:               m.Package,
		Workflow:              m.Workflow,
		EvidenceArtifact:      m.EvidenceArtifact,
		PrincipalCount:        len(model.PrincipalKinds),
		DaemonHeaderCount:     len(model.DaemonHeaders),
		ClientHeaderCount:     len(model.ClientHeaders),
		SnapshotInterval:      model.SnapshotInterval,
		RuntimeStaleAfter:     model.RuntimeStaleAfter,
		OwnershipEdgeCount:    len(model.OwnershipEdges),
		BindingSourceCount:    len(model.BindingSources),
		ExcludedFallbackCount: len(model.ExcludedFallbacks),
		SecretSinkCount:       len(model.SecretSinks),
		DependencyPhraseCount: model.DependencyPhraseCount,
		BindingFieldCount:     len(model.BindingFields),
		PolicyRuleCount:       len(model.PolicyRules),
		Loop:                  m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
