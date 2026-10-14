package progressmessage

import "testing"

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
