package main

import (
	"errors"
	"fmt"
	"strings"
)

// runtimeProfile is a document manifest, not a public assignment wire DTO.
// Each value names a safety boundary consumed by the daemon and server.
type runtimeProfile struct {
	InitialWorkspacePreparationBudget string `json:"initial_workspace_preparation_budget"`
	ProviderRecoveryElapsedBudget     string `json:"provider_recovery_elapsed_budget"`
	ReadyACKRequirement               string `json:"ready_ack_requirement"`
	ReadyReportRetry                  string `json:"ready_report_retry"`
	TerminalEventDedupe               string `json:"terminal_event_dedupe"`
	TerminalEventKeyConflict          string `json:"terminal_event_key_conflict"`
}

func validateRuntimeProfile(p runtimeProfile) error {
	if p.InitialWorkspacePreparationBudget != "host-workspace-policy" || p.ProviderRecoveryElapsedBudget != "exclude-initial-workspace-preparation" {
		return errors.New("normal initial workspace preparation must retain host policy outside the provider recovery budget")
	}
	if p.ReadyACKRequirement != "accepted-original-execution-before-provider-dispatch" || p.ReadyReportRetry != "bounded-with-same-execution-and-stable-event-key" {
		return errors.New("dispatch requires an accepted original execution ACK and bounded identity-preserving report retries")
	}
	if p.TerminalEventDedupe != "durable-original-event-ack-across-replicas" || p.TerminalEventKeyConflict != "reject-without-state-change" {
		return errors.New("terminal event keys require durable original ACK replay and non-mutating conflict rejection")
	}
	return nil
}

func renderRuntimeProfile(out *strings.Builder, p runtimeProfile) {
	out.WriteString("## 실행 전 프로파일의 검증 가능한 안전 규칙\n\n| 경계 | 규칙 |\n| --- | --- |\n")
	for _, row := range [][2]string{
		{"정상 최초 workspace 준비", p.InitialWorkspacePreparationBudget},
		{"Provider 복구 누적 기한", p.ProviderRecoveryElapsedBudget},
		{"Ready 수락과 dispatch", p.ReadyACKRequirement},
		{"Ready 보고 재시도", p.ReadyReportRetry},
		{"terminal 중복 수신", p.TerminalEventDedupe},
		{"terminal key 충돌", p.TerminalEventKeyConflict},
	} {
		fmt.Fprintf(out, "| %s | `%s` |\n", row[0], row[1])
	}
	out.WriteString("\n이 manifest 검증은 소비자 구현·배포·실모델 검증을 대신하지 않는다.\n\n")
}
