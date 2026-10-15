package main

import "testing"

func TestRecoveryRuntimeProfileRejectsUnsafeBoundaries(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*runtimeProfile)
	}{
		{"missing_profile", func(p *runtimeProfile) { *p = runtimeProfile{} }},
		{"clone_cancelled_by_retry_deadline", func(p *runtimeProfile) { p.InitialWorkspacePreparationBudget = "provider-recovery-deadline" }},
		{"clone_consumes_retry_budget", func(p *runtimeProfile) { p.ProviderRecoveryElapsedBudget = "include-workspace-preparation" }},
		{"dispatch_before_ack", func(p *runtimeProfile) { p.ReadyACKRequirement = "best-effort-report" }},
		{"unbounded_ready_retry", func(p *runtimeProfile) { p.ReadyReportRetry = "retry-forever" }},
		{"replica_local_terminal_dedupe", func(p *runtimeProfile) { p.TerminalEventDedupe = "process-memory-only" }},
		{"overwrite_conflicting_key", func(p *runtimeProfile) { p.TerminalEventKeyConflict = "overwrite-receipt" }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := loadDesign(t)
			tc.mutate(&d.RuntimeProfile)
			if err := validate(d); err == nil {
				t.Fatal("unsafe runtime profile accepted")
			}
		})
	}
}
