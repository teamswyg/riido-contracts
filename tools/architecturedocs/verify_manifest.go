package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyManifest(m manifest) error {
	if m.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", m.SchemaVersion, schemaVersion)
	}
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.Summary) {
		return errors.New("id, riido_task, and summary are required")
	}
	if blank(m.GeneratedDocs.ModuleDecomposition) || blank(m.GeneratedDocs.IntegrationMatrix) {
		return errors.New("generated docs are required")
	}
	if len(m.RequiredDocs) == 0 || len(m.Packages) == 0 || len(m.PublicGates) == 0 {
		return errors.New("required_docs, packages, and public_gates are required")
	}
	if len(m.ContractShape.Allowed) == 0 || len(m.ContractShape.Forbidden) == 0 {
		return errors.New("contract_shape allowed and forbidden values are required")
	}
	if len(m.LocalCommands) == 0 || len(m.StaleRuntimeWords) == 0 {
		return errors.New("local_commands and stale_runtime_words are required")
	}
	return verifyLoop(m.Loop)
}

func verifyLoop(loop evidenceLoop) error {
	if blank(loop.Observation) || blank(loop.Hypothesis) || blank(loop.Execute) {
		return errors.New("loop observation, hypothesis, and execute are required")
	}
	if blank(loop.Evaluate) || blank(loop.Retrospective) {
		return errors.New("loop evaluate and retrospective are required")
	}
	return nil
}

func blank(value string) bool {
	return strings.TrimSpace(value) == ""
}
