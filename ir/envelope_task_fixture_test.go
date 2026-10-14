package ir

import "time"

func validTaskScopeEvent(eventType EventType) CanonicalEvent {
	return CanonicalEvent{
		EventID:             "ev_task",
		OccurredAt:          time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:  1,
		Scope:               EventScopeTask,
		Type:                eventType,
		ActorKind:           ActorDaemon,
		ActorID:             "daemon-1",
		RiidoDaemonVersion:  "0.1.0",
		PolicyBundleVersion: "pb-1",
		TaskID:              "task_1",
		FSMVersion:          1,
	}
}
