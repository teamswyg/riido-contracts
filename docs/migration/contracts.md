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
| Provider capability fingerprint schema | `internal/provider/capability`, `docs/20-domain/provider-capability.md` | Promoted by RIID-4642 into `provider/capability`; keep provider detect logic in daemon. |
| Distribution channel + provider routing status vocabulary | `internal/hostintegration`, `docs/20-domain/distribution-host-integration.md` | Promoted by RIID-4670 into `hostintegration`; keep app data roots, IPC, grants, provider discovery, and review/demo mode in runtime repos. |
| Assignment polling DTOs | `internal/riidoaiserver`, `assignment_contract.riido.json` | Promoted by RIID-4687 into `assignment`; keep server store logic, health/metrics adapters, HTTP/SSE, authZ, and persistence in control-plane. |
| API DSL / IR / OpenAPI projection | control-plane HTTP docs and web UI contract needs | Added by RIID-4718 as candidate shared projection fixtures; OpenAPI is generated from IR and is not SSOT. |
| AI Agent policy and client API projection | v1.22 AI Agent Figma flow and control-plane client surface | Added by RIID-4720 as shared policy vocabulary and API projection fixtures; keep handlers, daemon probing, and generated client code in owner repos. |
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
| Distribution / provider routing vocabulary | `distribution-host-integration.md` | Daemon and control-plane import shared enum values before provider status migration. |
| Assignment polling API | `assignment-polling.md` | Daemon and control-plane import tagged assignment DTOs before daemon SaaS adapter migration. |

## Migration Order

1. Add this migration plan and keep the initial module stdlib-only.
2. Port shared fixtures before shared Go APIs.
3. Tag the first contract release only after a real consumer imports it.
4. Replace duplicated constants in daemon/control-plane with tagged imports.
5. Add cross-repository black-box tests that consume the same fixture.

## Current Migration Slices

### RIID-4713 — architecture SSOT docs migration

This slice restores the public architecture SSOT set for the split-repo
contracts boundary.

This slice does:

- add `docs/20-domain/context-map.md` for public context ownership
- add `docs/30-architecture/module-decomposition.md` for package/import rules
- add `docs/30-architecture/contract-promotion-policy.md` for promotion,
  tagging, and breaking-change policy
- add `docs/30-architecture/integration-matrix.md` for public CI and
  downstream compatibility gates
- add `docs/50-roadmap/open-questions.md` for unresolved shared-contract
  decisions
- add a focused public architecture-docs GitHub Actions workflow
- link README and domain docs to the new SSOT set

This slice does not add a new public contract API, create a Go module tag, move
runtime implementation, move Terraform/AWS/deployment evidence, or commit
private fixtures.

### RIID-4718 — API DSL IR OpenAPI projection

This slice adds the first shared API contract projection fixture for web
frontend/control-plane compatibility.

This slice does:

- add `apicontract` with the Domain DSL -> API IR -> OpenAPI projection model
- add `control-plane-agent-catalog-api.v1` as the first DSL fixture
- generate `riido-api-ir.v1` and OpenAPI 3.1 JSON from the DSL
- keep OpenAPI as a generated projection rather than the SSOT
- add `tools/apicontract verify` for deterministic fixture drift checks
- document the API projection boundary and add focused public CI

This slice does not move control-plane HTTP handlers, authorization/RBAC
implementation, frontend implementation, generated frontend client code,
production bearer tokens, IdP config, Terraform, AWS data, or deployment
evidence.

### RIID-4720 — AI Agent policy vocabulary and client API projection

This slice adds the v1.22 AI Agent policy contract for web and desktop webview
clients.

This slice does:

- document the shared vocabulary for runtime, agent, control plane, device,
  daemon, and client
- document agent deletion, runtime deletion, daemon detection, agent editing,
  and runtime-output parsing policy
- extend API DSL/IR with top-level enum and sum-type definitions
- add `control-plane-ai-agent-client-api.v1` fixtures for bootstrap,
  device/runtime listing, task assignable-agent listing, editability checks,
  task-thread comment submit, task-thread stop, mutation, deletion, and SSE
  client events
- keep OpenAPI as generated projection for Orval or compatible client codegen

This slice does not implement control-plane handlers, daemon runtime detection,
desktop/web UI code, generated client code, or provider runtime parsing.

## Validation Gates

Required for this repository:

```bash
go test ./...
go list -m all
go run ./tools/apicontract verify
```

Architecture-doc migration PRs must also pass:

```bash
test -f docs/20-domain/context-map.md
test -f docs/30-architecture/module-decomposition.md
test -f docs/30-architecture/contract-promotion-policy.md
test -f docs/30-architecture/integration-matrix.md
test -f docs/50-roadmap/open-questions.md
go test ./...
```

When contract fixtures are added, public CI must also verify:

- generated fixture drift
- schema version string alignment
- daemon/control-plane compatibility examples
- API DSL / IR / OpenAPI projection drift

## Migration Work Map

| Area | Riido task | Target repository |
| --- | --- | --- |
| Daemon runtime | RIID-4636 | `riido-daemon` |
| CLI surface | RIID-4635 | `riido-daemon` |
| Contracts | RIID-4637 | `riido-contracts` |
| Control plane | RIID-4638 | `riido-control-plane` |
| Infrastructure | RIID-4639 | `riido-infra` |
