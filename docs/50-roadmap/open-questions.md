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
| Q-CON-006 | Runtime model catalog | Should the AI Agent setting screen's model dropdown be owned as a canonical runtime capability catalog, a control-plane client sub-DSL read model, or a daemon-reported runtime fact? | Resolved in `runtime_model_catalog.v1`: model dropdown values are runtime-scoped catalog records projected under `RuntimeRecord.models`, not top-level API enum values. `model_id` is an opaque runtime-scoped identifier, labels are client display data, omitted create/update values resolve to the selected runtime's default model, and control-plane validates `model_id` against the selected runtime before saving agent configuration. Figma labels from `node-id=164-50215` remain evidence/sample catalog data rather than enum values. |
| Q-CON-007 | Runtime waitlist | Should the Windows app waitlist and marketing-consent mutation shown in Figma `node-id=275-22731` and web onboarding `node-id=236-29749` be owned by the AI Agent client API, a shared user/marketing API, or a separate product system? | Unresolved. Current AI Agent contract exposes only device/runtime liveness through `GET /v1/client/ai-agent/devices`; provider install cards, hover states, Windows waitlist copy, launch-notification CTA, and marketing-consent presentation remain client/product facts until an owning SSOT adds a generated operation. |
