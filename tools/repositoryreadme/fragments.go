package main

func loadFragments(base string, m *manifest) error {
	for _, fragmentPath := range m.Fragments {
		fragment, err := loadFragment(repoPath(base, fragmentPath))
		if err != nil {
			return err
		}
		mergeFragment(m, fragment)
	}
	return nil
}

func mergeFragment(m *manifest, fragment manifestFragment) {
	m.Summary = append(m.Summary, fragment.Summary...)
	m.Owns = append(m.Owns, fragment.Owns...)
	m.DoesNotOwn = append(m.DoesNotOwn, fragment.DoesNotOwn...)
	m.Rationale = append(m.Rationale, fragment.Rationale...)
	m.DocLinks = append(m.DocLinks, fragment.DocLinks...)
	m.Packages = append(m.Packages, fragment.Packages...)
	m.FSM.Intro = append(m.FSM.Intro, fragment.FSM.Intro...)
	m.FSM.Sections = append(m.FSM.Sections, fragment.FSM.Sections...)
	m.Decisions = append(m.Decisions, fragment.Decisions...)
	m.Verification = append(m.Verification, fragment.Verification...)
	m.Rules = append(m.Rules, fragment.Rules...)
	m.RequiredMarkers = append(m.RequiredMarkers, fragment.RequiredMarkers...)
	m.ForbiddenLiterals = append(m.ForbiddenLiterals, fragment.ForbiddenLiterals...)
	if completeLoop(fragment.Loop) {
		m.Loop = fragment.Loop
	}
}
