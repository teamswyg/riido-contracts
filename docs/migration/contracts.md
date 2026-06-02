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
  task-thread cold collection, task-thread comment submit, task-thread stop,
  mutation, deletion, and SSE client events
- keep OpenAPI as generated projection for Orval or compatible client codegen

This slice does not implement control-plane handlers, daemon runtime detection,
desktop/web UI code, generated client code, or provider runtime parsing.

### RIID-4827 — bootstrap scenario wording SSOT clarification

This slice clarifies a BDD scenario sentence in the AI Agent client contract
without changing API shape.

This slice does:

- replace ambiguous future-client/bootstrap wording with subsequent
  `aiAgent.bootstrap` read wording
- regenerate the derived IR fixture from the DSL
- add a regression gate so the ambiguous wording cannot return to the
  canonical client contract fixtures

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4857 — Figma onboarding loaded inventory and draft-create SSOT

This slice closes the `42:3014` `Wireframe - 온보딩` loaded-page inventory gap
and absorbs the Figma `node-id=432:46849` onboarding order-change memo.

The loaded Figma Plugin API inspection says `42:3014` has 83 top-level
children. Earlier contracts coverage recorded only three top-level nodes, which
made the manifest contradict its own loaded-page child-count rule. The same
Figma pass shows an order-change memo: `에이전트 생성 → 런타임 선택 → 워크스페이스
선택`. Contracts interpret that as client-local onboarding draft/configuration
selection before final submit, not as a durable workspace-less agent create
command.

This slice does:

- update `figma-ai-agent-coverage.riido.json` so the `42:3014` loaded inventory
  has 83 top-level nodes and `node-id=432:46849` is a coverage-bearing planning
  decision
- update `ai-agent-policy.md`, `api-contract-projection.md`, and the API DSL so
  local onboarding drafts remain client-owned and final v2 fixture/direct create
  still requires `workspace_id` in the URL and `runtime_id` in the request
- regenerate the derived AI Agent client IR/OpenAPI fixtures from the DSL
- add `onboarding-draft-ordering` to the SSOT dependency manifest so downstream
  repos can mirror the boundary without redefining it
- require the Figma coverage gate to expect the updated provenance and the new
  non-UI coverage-bearing node count

This slice does not add persisted draft state, a workspace-less create route,
new authorization scopes, new RBAC policy ids, generated client delivery
automation, frontend code, control-plane handlers, daemon runtime behavior,
Terraform, AWS data, or deployment evidence.

### RIID-4828 — task-thread v2 submitComment coverage gate

This slice closes a coverage drift between the contracts-owned Figma manifest
and the downstream control-plane generated-client projection.

This slice does:

- add `v2.aiAgent.tasks.submitComment` to the task-thread Figma coverage
  generated-path hints
- keep `aiAgent.tasks.threadMessages.create` as the canonical next-instruction
  operation while documenting the compatibility `submitComment` route in both
  v1 and v2 generated surfaces
- require the human coverage doc to mention each generated path recorded by the
  executable coverage manifest

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4829 — Figma non-UI page coverage SSOT

This slice extends the Figma AI Agent coverage manifest from the primary UI
page to the whole inspected Figma file.

This slice does:

- add a deterministic page registry for `UI`, `Wireframe - 온보딩`, and
  `Wireframe`
- add non-UI top-level evidence entries for desktop onboarding, macOS,
  Windows, and legacy planning wireframes
- classify platform wireframes as no-diff product surfaces unless a future SSOT
  adds a durable AI Agent API or daemon fact
- require coverage tests to register and document non-UI top-level evidence
  instead of leaving it as hidden Figma drift

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4830 — Figma coverage inspection method SSOT

This slice makes the Figma coverage inspection authority explicit.

This slice does:

- record that the page registry and top-level child counts come from the Figma
  Plugin API `figma.root.children` and `page.children.length`
- classify metadata XML/read output and node-specific metadata lookup as
  supporting evidence only
- require the human coverage doc and executable manifest to agree on that
  inspection method
- prevent future coverage updates from treating expanded metadata XML subtrees
  as page-level child-count drift

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4832 — Figma lazy page load coverage drift gate

This slice closes a gap in the Figma coverage inspection method.

The prior RIID-4830 wording treated `page.children.length` as authoritative
without saying that non-current pages must be loaded first. A fresh Figma Plugin
API inspection showed that passive page-registry reads can under-report lazy
page children: page `0:1` (`Wireframe`) looked like it had one child until the
script loaded the page with `await figma.setCurrentPageAsync(page)`, after
which the loaded top-level count was 28.

This slice does:

- make loaded-page inspection part of the coverage SSOT:
  `await figma.setCurrentPageAsync(page); page.children.length`
- update the `Wireframe` page child count from 1 to the loaded count 28
- add `non_ui_top_level_inventory` so every loaded non-UI top-level node is
  registered without pretending every legacy layer owns a policy/API decision
- keep `non_ui_top_level_nodes` as the narrower list of non-UI nodes that carry
  a real coverage decision
- require coverage tests to prove inventory length matches each non-UI page's
  loaded `child_count`

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4834 — Figma API Generated annotation normalization gate

This slice closes a coverage gap around Figma Dev Mode category `700:0`
`API Generated`.

This slice does:

- add `api_generated_annotations` to the Figma AI Agent coverage manifest,
  replacing the older handoff-oriented field name with the current
  `API Generated` category vocabulary
- preserve Figma facade examples such as `riido.aiAgent.events.stream` and
  `riido.aiAgent.tasks.stop`
- normalize those examples to canonical generated paths
  `aiAgent.events.stream` and `aiAgent.tasks.stop`
- require the coverage gate to prove that the canonical paths exist in OpenAPI
  and in the task-thread coverage entry
- document that `상세내용은 작업중입니다` on `node-id=153:8545` is stale
  Figma handoff copy, not a missing endpoint

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4838 — Legacy Wireframe semantic node coverage gate

This slice closes a gap left by the loaded-page Figma coverage inventory.

The prior manifest registered every loaded top-level node on legacy page `0:1`
(`Wireframe`), but semantically meaningful frames such as `런타임`, `에이전트`,
`에이전트 수정`, `에이전트 추가`, `데몬 상세`, and `런타임 상세` could still remain
inventory-only. That made the page count deterministic, but it did not prove
that old design frames had been absorbed into the current UI coverage entries.

This slice does:

- promote the legacy runtime, agent, agent add/edit, daemon detail, and runtime
  detail frames from inventory-only records to explicit `non_ui_top_level_nodes`
  coverage decisions
- record `absorbed_by_top_level_node_id` so each legacy frame points to the
  current UI section that owns the durable meaning
- require the coverage test to fail when one of those semantic legacy nodes is
  missing, lacks generated-path inheritance, or points to a non-covered current
  UI section
- keep Figma as evidence only: the current UI section, policy doc, and API DSL
  still own the actual domain/API meaning

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4840 — Figma metadata page-list limitation guard

This slice records a Figma read-tool failure mode that can otherwise collapse
the coverage SSOT.

On 2026-06-02, `get_metadata` without `nodeId` returned only page `129:5215`
`UI` for file `MUOd9lctoEHASUStN3vUuK`. The Figma Plugin API page registry for
the same file returned `129:5215`, `42:3014`, and `0:1`, and the loaded
`0:1` page still has 28 top-level children. Therefore the no-`nodeId` metadata
page list is supporting evidence only.

This slice does:

- add `supporting_tool_limitations` to the Figma coverage manifest
- document that no-`nodeId` `get_metadata` must not remove `expected_pages`,
  non-UI inventories, or legacy Wireframe coverage
- require the coverage test to preserve the three authoritative page IDs from
  the Figma Plugin API

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma onboarding timeout provenance catch-up

This slice repairs the Figma coverage stabilization history after the
onboarding page load timeout limitation was added.

`teamswyg/riido-contracts#60` changed executable Figma coverage meaning by
adding `figma-onboarding-page-load-timeout.v1` to
`supporting_tool_limitations`. Before this slice, the canonical
`stabilized_by` list still stopped at `teamswyg/riido-contracts#58`, so
downstream mirrors could not prove that they consumed the latest timeout
limitation state.

This slice does:

- append `teamswyg/riido-contracts#60` to the Figma coverage `stabilized_by`
  list
- require the coverage test to expect that provenance entry in order
- keep the update documentation-only so no API, daemon, infra, generated client,
  or deployment behavior changes

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma API Generated annotation content policy

This slice records the live Figma annotation content rule requested for frontend
handoff.

The Figma file now uses category `700:0` / `API Generated` for generated-client
handoff annotations. A lightweight Figma Plugin API traversal on 2026-06-02 found
59 `riido.*` annotations across the inspected file: 53 on `129:5215` `UI`, 6 on
`42:3014` `Wireframe - 온보딩`, and 0 on `0:1` `Wireframe`. Every live `riido.*`
annotation is in `API Generated`, has an operation kind (`Query`, `Mutation`, or
`SSE Stream`), and has Korean `배경:` handoff text.

This slice does:

- append `teamswyg/riido-contracts#62` to the Figma coverage `stabilized_by`
  list
- add `api_generated_annotation_content_policy` to the Figma coverage manifest
- record the required handoff label shape: path, `종류`, then `배경`
- record page-level live inspection counts for UI, onboarding, and legacy
  wireframe pages
- require the coverage test to verify category, label-format policy, page
  totals, zero missing operation kinds, and zero missing backgrounds

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma retired client-delivery annotation category

This slice records the remaining state of the old Figma annotation category
after `API Generated` became the generated-client handoff category.

The current Figma file still defines category `39:0` / `클라이언트 전달`, but a
2026-06-02 whole-file annotation usage scan found zero annotations using it.
The current Figma MCP exposes category entries as data objects without callable
`remove` or `setLabel` methods, so this automation could not delete the unused
category definition from the file.

This slice does:

- add the old category to `api_generated_annotation_content_policy` as a retired
  category
- record `live_usage_count=0`
- document that `700:0` / `API Generated` is the active generated-client
  handoff category
- require the coverage test to fail if the retired category regains usage or is
  confused with the active category

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma onboarding direct-node fallback evidence

This slice refines the onboarding page-load timeout limitation with the fallback
that succeeded during live Figma inspection.

The full `42:3014` `Wireframe - 온보딩` page traversal can still time out after
120s. However, direct Figma Plugin API `getNodeByIdAsync` reads for registered
nodes `236:33845` and `236:33847` returned the six onboarding `riido.*`
`API Generated` annotations with operation kind and Korean background text.

This slice does:

- keep full loaded-page inventory as the authority for `child_count=83`
- record `236:33845`, `236:33847`, and
  `onboarding_api_generated_annotations=6` as fallback authoritative results
  inside `figma-onboarding-page-load-timeout.v1`
- require the coverage test to preserve that direct-node fallback evidence

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma headless file-key placeholder limitation

This slice records a small but important Figma Plugin API runtime limitation
found during the continuing whole-file audit.

A live `use_figma` root/category inspection on 2026-06-02 successfully read the
AI Agent file's pages and annotation categories, but `figma.fileKey` returned the
placeholder value `headless` instead of `MUOd9lctoEHASUStN3vUuK`. Contracts keep
the user-provided Figma URL/tool input key and the manifest's `figma.file_key` as
the authoritative file identity.

This slice does:

- add `figma-headless-file-key-placeholder.v1` to
  `supporting_tool_limitations`
- document that `figma.fileKey=headless` is supporting evidence only
- require the coverage test to reject using that placeholder as manifest or
  downstream projection source identity

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma API Generated operation-kind transport guard

This slice tightens the Figma API Generated annotation guard.

Figma annotations expose `종류: Query | Mutation | SSE Stream` for frontend
handoff readability, but that value must not become a second source of truth.
Contracts derive the canonical transport from the generated OpenAPI operation:
`text/event-stream` responses are `SSE Stream`, non-stream `GET` operations are
`Query`, and non-`GET` operations are `Mutation`.

This slice does:

- document the transport-derived meaning of `operation_kind`
- require `api_generated_annotation_inventory.operation_kind` to match the
  generated OpenAPI method and response content type
- keep Figma text as handoff evidence, not as an API operation-kind authority

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma onboarding page load timeout limitation

This slice records a live Figma tooling limitation found while continuing the
whole-file AI Agent coverage audit.

Current live reads of `node-id=42:3014` (`Wireframe - 온보딩`) can time out
after 120s when using Figma `get_metadata(nodeId=42:3014)` or `use_figma`
scripts that attempt `await figma.setCurrentPageAsync(page)`. This is different
from the no-`nodeId` metadata page-list under-report: the page is known and
registered, but the current tool call cannot reliably reload it.

This slice does:

- add `figma-onboarding-page-load-timeout.v1` to
  `supporting_tool_limitations`
- document that the timeout is supporting evidence only
- require the coverage gate to preserve page `42:3014`, its captured
  `child_count=83`, its onboarding inventory, and its generated path coverage

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### Figma API Generated provenance catch-up

This slice repairs the Figma coverage stabilization history after the API
Generated annotation passes.

`teamswyg/riido-contracts#56`, `#57`, and `#58` changed executable Figma
coverage meaning: #56 registered the screen-level API Generated annotation
inventory, #57 moved the Figma category to `700:0` / `API Generated`, and #58
renamed the manifest fields to `api_generated_annotations` and
`api_generated_annotation_inventory`. Before this slice, the canonical
`stabilized_by` list still stopped at `teamswyg/riido-contracts#55`, so
downstream mirrors could not prove that they consumed the latest annotation
coverage state.

This slice does:

- append `teamswyg/riido-contracts#56`, `#57`, and `#58` to the Figma coverage
  `stabilized_by` list
- require the coverage test to expect those provenance entries in order
- keep the update documentation-only so no API, daemon, infra, generated client,
  or deployment behavior changes

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

### RIID-4849 — Figma coverage upstream provenance SSOT guard

This slice moves Figma coverage stabilization history back to the canonical
contracts manifest.

Control-plane and daemon downstream projections already mirror the contracts
coverage history (`teamswyg/riido-contracts#38`, `#39`, `#45`, `#46`, `#51`,
and `#52`) to explain which upstream Figma coverage state they consumed. Before
this slice, the canonical contracts manifest did not own that list, so
downstreams had to preserve it from local knowledge instead of mirroring a
contracts field.

This slice does:

- add top-level `stabilized_by` to
  `docs/30-architecture/figma-ai-agent-coverage.riido.json`
- document that `stabilized_by` is the downstream projection mirror source
- require `TestFigmaAIAgentCoverageManifest` to fail if the list is missing,
  out of order, or absent from the human coverage doc

This slice does not change routes, schemas, authorization, RBAC, generated
client delivery, frontend code, control-plane handlers, Terraform, AWS data, or
deployment evidence.

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
