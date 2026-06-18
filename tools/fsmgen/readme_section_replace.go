package main

func replaceReadmeSections(body string, sections []readmeSection) (string, error) {
	updated := body
	for _, section := range sections {
		next, err := replaceSection(updated, section)
		if err != nil {
			return "", err
		}
		updated = next
	}
	return updated, nil
}

func verifyReadmeSectionContent(body string, sections []readmeSection) error {
	for _, section := range sections {
		got, err := extractSection(body, section.ID)
		if err != nil {
			return err
		}
		if got != section.Content {
			return readmeSectionDriftError(section.ID)
		}
	}
	return nil
}
