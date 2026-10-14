package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:            "riido-host-integration-evidence.v1",
		ID:                       m.ID,
		Status:                   "verified",
		GeneratedDoc:             m.GeneratedDoc,
		Package:                  m.Package,
		Workflow:                 m.Workflow,
		EvidenceArtifact:         m.EvidenceArtifact,
		DistributionChannelCount: len(model.DistributionChannels),
		StoreManagedChannelCount: len(model.StoreManagedChannels),
		ProviderStatusCount:      len(model.ProviderStatuses),
		NonOwnedSurfaceCount:     len(model.NonOwnedSurfaces),
		DistributionValid:        model.DistributionValid,
		ProviderRoutingValid:     model.ProviderRoutingValid,
		StoreManagedExclusive:    model.StoreManagedExclusive,
		Loop:                     m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
