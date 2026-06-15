package progressmessage

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestProgressMessageDSLGeneratesIR(t *testing.T) {
	dsl := loadTestDSL(t)
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	assertFixture(t, "catalog.ir.riido.json", ir)
}

func TestProgressMessageCatalogRendersKoreanMessage(t *testing.T) {
	got, ok := Render(1101, map[string]string{
		"label":       "팀 프로젝트",
		"description": "팀의 프로젝트 목록, 진행 상태, 우선순위와 담당자 정보를 조회해 요약을 준비 중. . .",
	}, "ko")
	if !ok {
		t.Fatal("Render returned false")
	}
	want := "팀 프로젝트 수집 중 - 팀의 프로젝트 목록, 진행 상태, 우선순위와 담당자 정보를 조회해 요약을 준비 중. . ."
	if got != want {
		t.Fatalf("Render = %q, want %q", got, want)
	}
}

func TestProgressMessageCatalogIsAppendOnlyAgainstHEAD(t *testing.T) {
	out, err := baselineCatalogFromGit()
	if err != nil {
		t.Skip("progress message catalog is new in this checkout")
	}
	var previous IRDocument
	if err := json.Unmarshal(out, &previous); err != nil {
		t.Fatalf("decode previous catalog: %v", err)
	}
	current, err := Catalog()
	if err != nil {
		t.Fatalf("Catalog: %v", err)
	}
	currentByCode := map[int]MessageDefinition{}
	for _, message := range current.Messages {
		currentByCode[message.Code] = message
	}
	for _, old := range previous.Messages {
		now, ok := currentByCode[old.Code]
		if !ok {
			t.Fatalf("progress message code %d (%s) was removed. Progress message codes are append-only; removing a code can break persisted backend events, daemon payloads, and replayed SSE records. Mark it reserved instead.", old.Code, old.Key)
		}
		if now.Key != old.Key {
			t.Fatalf("progress message code %d changed key from %q to %q. Codes are append-only and their identity is immutable.", old.Code, old.Key, now.Key)
		}
		if argNames(now.Args) != argNames(old.Args) {
			t.Fatalf("progress message code %d changed args from %q to %q. Add a new code instead of mutating an existing backend/daemon payload shape.", old.Code, argNames(old.Args), argNames(now.Args))
		}
	}
}

func baselineCatalogFromGit() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if base, err := exec.CommandContext(ctx, "git", "merge-base", "HEAD", "origin/main").Output(); err == nil {
		ref := strings.TrimSpace(string(base))
		if ref != "" {
			if out, err := exec.CommandContext(ctx, "git", "show", ref+":progressmessage/catalog.ir.riido.json").Output(); err == nil {
				return out, nil
			}
		}
	}
	return exec.CommandContext(ctx, "git", "show", "HEAD^:progressmessage/catalog.ir.riido.json").Output()
}

func loadTestDSL(t *testing.T) DSLDocument {
	t.Helper()
	body, err := os.ReadFile("catalog.dsl.riido.json")
	if err != nil {
		t.Fatalf("read DSL: %v", err)
	}
	var dsl DSLDocument
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		t.Fatalf("decode DSL: %v", err)
	}
	return dsl
}

func assertFixture(t *testing.T, path string, value any) {
	t.Helper()
	want, err := MarshalCanonical(value)
	if err != nil {
		t.Fatalf("MarshalCanonical: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("%s drifted; run go run ./tools/progressmessage generate", path)
	}
}

func argNames(args []MessageArg) string {
	names := make([]string, 0, len(args))
	for _, arg := range args {
		names = append(names, arg.Name)
	}
	return strings.Join(names, ",")
}
