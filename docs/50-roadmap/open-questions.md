# Contracts Open Questions

> Riido task: RIID-4713 `[Contracts] Architecture SSOT docs migration`

This file owns unresolved shared-contract decisions.

| ID | Area | Question | Current working stance |
| --- | --- | --- | --- |
| Q-CON-001 | Contract fixtures | Should cross-repo black-box fixtures live as JSON under this repo or generated Go testdata packages? | Keep Go constants/DTOs first; add JSON fixtures only when two repos consume the same file. |
| Q-CON-002 | RBAC fixtures | Should agent catalog RBAC scenarios move here? | Promote only scenario fixtures, not authorization implementation, after daemon or another client consumes them. |
| Q-CON-003 | Store distribution | Should store distribution contract fixtures move from daemon to contracts? | Promote only if daemon and infra both validate the same public fixture. |
| Q-CON-004 | Tag cadence | Should each promoted package get an immediate tag or batch tags by migration wave? | Tag after merge when a downstream import gate is ready. |
| Q-CON-005 | Version labels | Should package schema versions include the module tag? | No. Package schema constants and Go module tags remain separate axes. |
