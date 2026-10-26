package contextcli

import (
	"encoding/json"
	"testing"
)

func TestEnvelopeWireShape(t *testing.T) {
	env := Envelope{SchemaVersion: EnvelopeSchemaVersion, RequestID: "cli_test", Sources: []Source{}, Warnings: []Warning{}}
	b, err := json.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"schema_version":"1","request_id":"cli_test","data":null,"sources":[],"warnings":[],"page":null,"error":null}`
	if string(b) != want {
		t.Fatalf("envelope = %s", b)
	}
}

func TestStableErrorExitCodes(t *testing.T) {
	cases := map[ErrorCode]ExitCode{
		CodeUsage: 2, CodeSecretInArgv: 2, CodeAuthRequired: 3, CodePermissionDenied: 4, CodeOutOfRunScope: 4,
		CodeNotFound: 5, CodeSnapshotExpired: 5, CodeUnavailable: 6, CodeDaemonUnreachable: 6, CodeRateLimited: 7,
		CodeUnsupported: 8, CodeContextReadUnsupported: 8, CodeUserModeUnavailable: 8, CodeControlPlaneUnsupported: 8,
		CodeIncomplete: 9, CodeConflict: 10, CodeLeaseLost: 10, ErrorCode("future_unknown"): 8,
	}
	for code, want := range cases {
		if got := code.ExitCode(); got != want {
			t.Errorf("%s exit=%d want %d", code, got, want)
		}
	}
}
