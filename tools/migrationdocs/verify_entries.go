package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyEntries(m manifest) error {
	if len(m.RepositoryBoundaries.MayContain) == 0 || len(m.RepositoryBoundaries.MustNotContain) == 0 {
		return errors.New("repository boundary lists are required")
	}
	if len(m.Versioning.Axes) == 0 || len(m.ValidationGates.RequiredCommands) == 0 {
		return errors.New("version axes and required validation commands are required")
	}
	for _, candidate := range m.CandidateContracts {
		if blank(candidate.Candidate) || blank(candidate.Source) || blank(candidate.Decision) {
			return errors.New("candidate contract entries must be complete")
		}
	}
	for _, slice := range m.MigrationSlices {
		if blank(slice.Title) || (len(slice.Intro) == 0 && len(slice.Does) == 0 && blank(slice.DoesNot)) {
			return errors.New("migration slices must have a title and content")
		}
		if migrationSliceHasHeading(slice) {
			return fmt.Errorf("migration slice %q embeds a top-level heading", slice.Title)
		}
	}
	for _, entry := range m.MigrationWorkMap {
		if blank(entry.Area) || blank(entry.RiidoTask) || blank(entry.TargetRepository) {
			return errors.New("migration work map entries must be complete")
		}
	}
	return nil
}

func migrationSliceHasHeading(slice migrationSlice) bool {
	for _, line := range slice.Intro {
		if strings.HasPrefix(strings.TrimSpace(line), "## ") {
			return true
		}
	}
	for _, line := range slice.Does {
		if strings.HasPrefix(strings.TrimSpace(line), "## ") {
			return true
		}
	}
	return strings.Contains(slice.DoesNot, " ## ")
}
