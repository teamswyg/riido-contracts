package ir

import "testing"

func TestActorKindValues(t *testing.T) {
	kinds := []ActorKind{ActorAgent, ActorDaemon, ActorHuman, ActorSystem}
	if len(kinds) != 4 {
		t.Fatalf("expected 4 actor kinds, got %d", len(kinds))
	}
}
