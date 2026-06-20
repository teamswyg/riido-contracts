package main

type evidence struct {
	SchemaVersion       string       `json:"schema_version"`
	ID                  string       `json:"id"`
	Status              string       `json:"status"`
	Manifest            string       `json:"manifest"`
	GeneratedDoc        string       `json:"generated_doc"`
	FragmentCount       int          `json:"fragment_count"`
	DocLinkCount        int          `json:"doc_link_count"`
	PackageCount        int          `json:"package_count"`
	FSMSectionCount     int          `json:"fsm_section_count"`
	VerificationCount   int          `json:"verification_count"`
	RequiredMarkerCount int          `json:"required_marker_count"`
	Loop                evidenceLoop `json:"loop"`
}

func buildEvidence(m manifest) evidence {
	return evidence{
		SchemaVersion: evidenceSchema, ID: m.ID, Status: "verified",
		Manifest: defaultManifest, GeneratedDoc: generatedDoc, FragmentCount: len(m.Fragments),
		DocLinkCount: len(m.DocLinks), PackageCount: len(m.Packages),
		FSMSectionCount: len(m.FSM.Sections), VerificationCount: len(m.Verification),
		RequiredMarkerCount: len(m.RequiredMarkers), Loop: m.Loop,
	}
}
