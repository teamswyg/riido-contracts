package main

import "errors"

func verifyEntries(m manifest) error {
	for _, link := range m.ArchitectureLinks {
		if blank(link.Label) || blank(link.Path) {
			return errors.New("architecture links must be complete")
		}
	}
	for _, change := range m.Changes {
		if blank(change.Task) || blank(change.Verb) || len(change.Items) == 0 {
			return errors.New("change entries must be complete")
		}
		for _, item := range change.Items {
			if blank(item) {
				return errors.New("change item must not be blank")
			}
		}
	}
	return nil
}
