package main

import (
	"time"

	"github.com/teamswyg/riido-contracts/ir"
)

func baseEvent(scope ir.EventScope) ir.CanonicalEvent {
	return ir.CanonicalEvent{
		EventID:             "ev_schema_probe",
		OccurredAt:          time.Date(2026, 6, 20, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:  1,
		Scope:               scope,
		Type:                ir.EventPolicyBundleLoaded,
		ActorKind:           ir.ActorSystem,
		RiidoDaemonVersion:  "0.1.0",
		PolicyBundleVersion: "pb-1",
	}
}

func validSystemEvent() ir.CanonicalEvent {
	return baseEvent(ir.EventScopeSystem)
}
