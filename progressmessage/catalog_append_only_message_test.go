package progressmessage

import "testing"

func assertMessageCodeAppendOnly(
	t *testing.T,
	old MessageDefinition,
	now MessageDefinition,
) {
	t.Helper()
	if now.Code == 0 {
		t.Fatalf("progress message code %d (%s) was removed. Mark it reserved instead.", old.Code, old.Key)
	}
	if now.Key != old.Key {
		t.Fatalf("progress message code %d changed key from %q to %q.", old.Code, old.Key, now.Key)
	}
	if argNames(now.Args) != argNames(old.Args) {
		t.Fatalf("progress message code %d changed args from %q to %q.", old.Code, argNames(old.Args), argNames(now.Args))
	}
}
