# Riido Contracts Migration Plan

> Riido task: RIID-4637 `[Contracts] 기존 riido_daemon contract 마이그레이션 계획/문서화`

This document defines when facts move from the former private `riido_daemon`
repository into the public `riido-contracts` Go module.

## Goal

`riido-contracts` owns shared contracts only. A contract belongs here when at
least two repositories must agree on the same fact at build time or black-box
test time.

Implementation details stay in the owning runtime repository.

## Promotion Rule

Move a fact into `riido-contracts` only when all conditions are true:

1. `riido-daemon` and `riido-control-plane`, or one of those plus
   `riido-infra`, must consume the same versioned contract.
2. The contract can be represented without importing runtime implementation
   packages.
3. The contract can be versioned, tested, and tagged independently.
4. The owning SSOT doc has been updated before code moves.

If only one repository consumes the fact, keep it local to that repository.

## Candidate Contracts

| Candidate | Source in private `riido_daemon` | Target decision |
| --- | --- | --- |
| IR event type names and envelope shape | `internal/ir`, `docs/20-domain/ir-event-log.md` | Promoted by RIID-4641 into `ir`; EventIngestor remains out of scope. |
| Task lifecycle state and transition fixture | `internal/task`, `docs/20-domain/task-lifecycle.md` | Promoted by RIID-4641 into `task`; `task -> ir` remains one-way. |
| Provider capability fingerprint schema | `internal/provider/capability`, `docs/20-domain/provider-capability.md` | Promote stable schema; keep provider detect logic in daemon. |
| Assignment polling DTOs | `internal/riidoaiserver`, `assignment_contract.riido.json` | Promote request/response contract; keep server store logic in control-plane. |
| RBAC scenario fixtures | `internal/riidoaiserver/*rbac*`, security docs | Promote black-box fixtures, not authorization implementation. |
| Store distribution contract fixtures | `packaging/store`, `tools/storecontract` | Promote only if daemon and infra both validate the same fixture. |

## Repository Boundaries

`riido-contracts` may contain:

- versioned Go constants and DTOs
- JSON schema or generated fixture files
- black-box scenario fixtures shared across repositories
- small validators that do not know about runtime adapters

`riido-contracts` must not contain:

- provider CLI command builders
- daemon process execution code
- control-plane stores or HTTP handlers
- Terraform modules, backend config, or AWS environment data
- private operational evidence

## Versioning

The module is versioned with Git tags. Each promoted contract must state which
version axis it affects:

| Axis | Owner before split | Contract handling |
| --- | --- | --- |
| IR schema | `ir-schema-versioning.md` | Contract tag must match schema doc update. |
| FSM schema | `task-lifecycle.md` | Contract fixture must match transition matrix. |
| Server API | `runtime-versioning.md` + SaaS docs | Control-plane imports tagged contract or generated fixture. |
| Provider capability | `provider-capability.md` | Daemon imports contract only for shared schema. |

## Migration Order

1. Add this migration plan and keep the initial module stdlib-only.
2. Port shared fixtures before shared Go APIs.
3. Tag the first contract release only after a real consumer imports it.
4. Replace duplicated constants in daemon/control-plane with tagged imports.
5. Add cross-repository black-box tests that consume the same fixture.

## Validation Gates

Required for this repository:

```bash
go test ./...
go list -m all
```

When contract fixtures are added, public CI must also verify:

- generated fixture drift
- schema version string alignment
- daemon/control-plane compatibility examples

## Migration Work Map

| Area | Riido task | Target repository |
| --- | --- | --- |
| Daemon runtime | RIID-4636 | `riido-daemon` |
| CLI surface | RIID-4635 | `riido-daemon` |
| Contracts | RIID-4637 | `riido-contracts` |
| Control plane | RIID-4638 | `riido-control-plane` |
| Infrastructure | RIID-4639 | `riido-infra` |
