# Figma AI Agent Coverage

> Riido task: RIID-4809 `[Contracts] Figma AI Agent 화면 커버리지 SSOT manifest`

This document explains how the Figma `v.1.22 AI Agent` file is absorbed into
Riido SSOT documents. The `UI` page owns the detailed generated-client coverage;
the wireframe pages are tracked as supporting evidence or no-diff product
surfaces. Figma is evidence. It is not the durable SSOT for
domain policy, API shape, daemon behavior, or infra topology.

The executable coverage manifest is
[`figma-ai-agent-coverage.riido.json`](figma-ai-agent-coverage.riido.json). It
records the Figma file, page, top-level sections, owner repositories, generated
client paths, verified evidence node IDs, and the top-down / bottom-up rule for
each section. Documentation may cite only node IDs registered in that manifest,
so stale Figma links fail deterministically instead of remaining as hidden
planning drift.

## Source

| Field | Value |
| --- | --- |
| Figma file | `v.1.22 AI Agent` |
| File key | `MUOd9lctoEHASUStN3vUuK` |
| Primary UI page node | `129:5215` |
| Pages inspected | `129:5215` UI, `42:3014` Wireframe - 온보딩, `0:1` Wireframe |
| Page registry authority | Figma Plugin API `figma.root.children`; loaded page child count from `await figma.setCurrentPageAsync(page); page.children.length` |
| Supporting read tools | metadata XML/read output and node-specific metadata lookup |
| Inspection date | `2026-06-02` |

The page registry and the top-level child counts in this document are
authoritative only when they come from the Figma Plugin API
`figma.root.children` and each page's loaded `page.children.length`. For
non-current pages, the inspection script must first call
`await figma.setCurrentPageAsync(page)` and only then read `page.children.length`
and `page.children`. Passive `figma.root.children` page objects can be lazy or
unloaded, so they are only a page registry source. Metadata XML/read tools are
supporting evidence only: they may expand internal node subtrees, return a
scoped XML view, or report lazy/unloaded child counts, so they must not redefine
`expected_pages.child_count` or the top-level page registry.

## Supporting Tool Limitations

`figma-metadata-page-list-underreports-pages.v1` records a concrete 2026-06-02
tooling limitation: calling Figma `get_metadata` without `nodeId` listed only
`129:5215` `UI` for file `MUOd9lctoEHASUStN3vUuK`. The same file's Figma
Plugin API page registry returned `129:5215`, `42:3014`, and `0:1`.

That means the no-`nodeId` metadata page list is supporting evidence only. It
must not remove `expected_pages`, collapse non-UI inventories, or erase legacy
Wireframe coverage. Page registry updates must come from `figma.root.children`,
and child-count updates must come from each loaded page after
`await figma.setCurrentPageAsync(page)`.

## Coverage Rule

Top-down changes start from product/design evidence:

```text
Figma evidence
  -> contracts policy or API DSL
  -> API IR / OpenAPI projection
  -> control-plane executable behavior
  -> generated client handoff
  -> daemon runtime consumption if assignments execute
  -> infra only when topology, secrets, media, durability, or deployment changes
```

Bottom-up changes start from implementation or harness evidence:

```text
repo finding
  -> local repo SSOT records the observed constraint
  -> contracts policy/API changes only if business meaning changes
  -> regenerate projections and update downstream harnesses
  -> update this Figma coverage manifest
```

If a Figma section cannot name an owner and SSOT document, it is not ready for
implementation. It must become either a resolved policy/API fact, a local
presentation fact, a no-diff product surface, or an open question.

## Page Registry

| Figma page | Page | Top-level children | Coverage role |
| --- | --- | --- | --- |
| `129:5215` | UI | 16 | primary detailed API/domain coverage |
| `42:3014` | Wireframe - 온보딩 | 3 | supporting onboarding/platform evidence |
| `0:1` | Wireframe | 28 | legacy planning evidence |

## UI Top-Level Coverage

| Figma node | Section | Status | Owning boundary |
| --- | --- | --- | --- |
| `153:12742` | 컴포넌트 참여자 드롭다운 | covered | contracts policy, API DSL, control-plane, client |
| `153:15931` | 댓글 소통 | covered | task-thread policy, API DSL, control-plane, daemon, client |
| `153:15932` | image 14 | non-decision asset | no independent SSOT decision |
| `153:15935` | 추가 기획 내용 | covered | contracts policy and API DSL before implementation |
| `162:23468` | 논의 필요 | covered | fixture-created agent identity, duplicate names, normal edit/delete lifecycle |
| `156:18767` | image 13 | non-decision asset | no independent SSOT decision |
| `156:19307` | 메뉴바 | covered | client route affordance, control-plane data after route open |
| `156:19308` | Group 6 | non-decision asset | no independent SSOT decision |
| `162:23090` | 런타임 설정페이지 | covered | device/runtime read model, agent-bound daemon detail/control |
| `164:30658` | 데스크탑앱 온보딩-런타임 감지 O | covered | workspace-scoped onboarding fixture/direct create |
| `435:60050` | 데스크탑앱 온보딩-런타임 감지 X | covered | device/runtime liveness, client/product install presentation |
| `164:45736` | image 8 | non-decision asset | no independent SSOT decision |
| `164:45741` | Group 5 | non-decision asset | no independent SSOT decision |
| `236:29749` | 웹 온보딩 | no-diff product surface | auth/team/distribution/product; not AI Agent API |
| `275:22731` | 런타임 설정페이지 엠티 | covered | runtime empty state from device/runtime read model |
| `432:37336` | 에이전트 설정페이지 | covered | agent configuration fields, editability, runtime/model catalog |

## Non-UI Page Coverage

`non_ui_top_level_inventory` in the executable manifest records every loaded
top-level node on the non-UI pages. `non_ui_top_level_nodes` remains narrower:
it lists only the non-UI nodes that carry a coverage decision. This separation
prevents lazy page loading from hiding Figma drift while avoiding fake policy
ownership for old screenshots, duplicated legacy frames, or asset-only layers.
The 2026-06-02 loaded inspection found page `0:1` `Wireframe` has 28 top-level
children after `await figma.setCurrentPageAsync(page)`, even though a passive
page-registry read can report it as a single child.

| Figma page | Figma node | Section | Status | Coverage rule |
| --- | --- | --- | --- | --- |
| `42:3014` | `164:30657` | 데스크탑앱 온보딩 | covered | absorbed by UI onboarding fixture/direct-create API coverage; desktop owns launch presentation |
| `42:3014` | `188:27707` | macOS | no-diff product surface | macOS install/launch presentation does not create an AI Agent endpoint |
| `42:3014` | `188:27708` | windows | no-diff product surface | Windows install/waitlist/launch-notification presentation stays outside AI Agent generated API |
| `0:1` | `13:3789` | 런타임 | covered | legacy runtime list absorbed by current UI runtime settings `162:23090` and `devices.runtimes` generated paths |
| `0:1` | `86:9988` | 런타임 | covered | expanded legacy runtime frame absorbed by current UI runtime settings `162:23090` and agent-bound daemon detail |
| `0:1` | `17:3551` | 에이전트 | covered | legacy agent list absorbed by current UI agent settings `432:37336` and normal agent lifecycle paths |
| `0:1` | `17:4231` | 에이전트 수정 | covered | legacy agent edit absorbed by current UI agent settings `432:37336`, editability, and `updateConfiguration` |
| `0:1` | `84:9846` | 에이전트 추가 | covered | legacy agent add absorbed by current UI agent settings `432:37336` and ordinary direct create paths |
| `0:1` | `17:2871` | 데몬 상세 | covered | legacy daemon detail absorbed by current UI runtime settings `162:23090` and agent-bound daemon command paths |
| `0:1` | `17:3111` | 런타임 상세 | covered | legacy runtime detail absorbed by current UI runtime settings `162:23090`; no standalone runtime-detail operation |
| `0:1` | `153:15934` | 추가 기획 내용 | covered | legacy planning evidence resolved by UI page `153:15935` and task/subtask assignment scope |

Legacy node absorption is explicit. When a non-UI Wireframe node carries a
current product/API meaning, the manifest records
`absorbed_by_top_level_node_id` so the old frame cannot become a second SSOT.
If a future Figma change revives one of those legacy frames with new business
meaning, the change must update the current UI coverage entry or create a new
owning SSOT before implementation begins.

## Generated Client Path Hints

The manifest intentionally stores generated paths that frontend developers
would search for in generated TypeScript comments. These are projections of the
API DSL, not hand-authored route names.

| UI area | Main generated paths |
| --- | --- |
| Participant dropdown | `aiAgent.tasks.assignableAgents`, `aiAgent.tasks.assign`, `aiAgent.tasks.unassign`, `v2.aiAgent.tasks.assignableAgents`, `v2.aiAgent.tasks.assign`, `v2.aiAgent.tasks.unassign` |
| Task thread | `aiAgent.tasks.threads`, `aiAgent.tasks.threadMessages.create`, `aiAgent.tasks.submitComment`, `aiAgent.tasks.stop`, `aiAgent.events.stream`, `v2.aiAgent.tasks.threads`, `v2.aiAgent.tasks.threadMessages.create`, `v2.aiAgent.tasks.submitComment`, `v2.aiAgent.tasks.stop`, `v2.aiAgent.events.stream` |
| Runtime settings | `aiAgent.devices.runtimes`, `aiAgent.agents.daemon.details`, `aiAgent.agents.daemon.*`, `v2.aiAgent.devices.runtimes`, `v2.aiAgent.agents.daemon.details`, `v2.aiAgent.agents.daemon.*` |
| Onboarding | `aiAgent.onboarding.fixtures`, `aiAgent.onboarding.fixtures.createAgent`, `aiAgent.agents.create`, `v2.aiAgent.onboarding.fixtures`, `v2.aiAgent.onboarding.fixtures.createAgent`, `v2.aiAgent.agents.create` |
| Agent settings | `aiAgent.bootstrap`, `aiAgent.agents.create`, `aiAgent.agents.updateConfiguration`, `aiAgent.agents.delete`, `aiAgent.agents.editability`, `v2.aiAgent.bootstrap`, `v2.aiAgent.agents.*` |

## Runtime Endpoint Labels

The runtime settings section contains an endpoint-looking text node at
`node-id=129:17930`. That node is registered as Figma evidence only. It does not
define the canonical AI Agent base URL, does not create a generated client path,
and must not be copied into public docs as a live host. In short, it is not a
canonical base URL, generated path, or live host export. Control-plane
deployment and smoke workflows use configured environment variable names for
live base URL values, while generated clients receive only operation paths and
caller-provided configuration.

## Client Delivery Annotations

Figma Dev Mode category `39:0` / `클라이언트 전달` is handoff evidence for
frontend generated-client consumption. It does not own API names. A Figma label
that starts with `riido.` is treated as a client facade example, and the durable
contracts generated path is the same value without the leading `riido.` prefix.

| Figma node | Figma label | Canonical generated path | Resolution |
| --- | --- | --- | --- |
| `153:8545` | `riido.aiAgent.events.stream` plus `상세내용은 작업중입니다` | `aiAgent.events.stream` | The Korean `상세내용은 작업중입니다` copy is stale handoff text, not a missing API. The stream is covered by the task-thread entry `153:15931`, API DSL/OpenAPI, and generated-client comments. |
| `236:20768` | `riido.aiAgent.tasks.stop` | `aiAgent.tasks.stop` | The stop hint is a client facade example over the task-thread stop operation, covered by `153:15931`, API DSL/OpenAPI, and generated-client comments. |

## Ownership Notes

- `docs/20-domain/ai-agent-policy.md` owns the durable business language:
  agent, runtime, device, daemon, visibility, editability, task-thread, fixture,
  workspace-scoped v2 API, and no-runtime semantics.
- Figma `node-id=162-23475` closes the former `논의 필요` note: agents created
  from onboarding fixtures are ordinary agents, may share the same display name,
  are identified by `agent_id`, and follow normal update/delete/editability/RBAC
  rules.
- `docs/20-domain/api-contract-projection.md` owns DSL -> IR -> OpenAPI
  projection mechanics and generated-client path searchability.
- `apicontract/fixtures/control-plane-ai-agent-client.dsl.riido.json` owns the
  executable API contract that is generated into IR/OpenAPI.
- `docs/30-architecture/ssot-dependency-map.md` owns cross-repo dependency
  direction and duplicate-audit rules.
- `riido-control-plane` owns HTTP/SSE execution and generated-client delivery.
- `riido-daemon` owns local runtime detection, provider execution, polling,
  provider-specific instruction placement, and process lifecycle behavior.
- `riido-client` and `riido-desktop` own rendering, route wiring, copy, scroll,
  animation, and desktop daemon launch trigger presentation.
- `riido-infra` acts only when deployment topology, secrets, storage, media,
  durability, cost, or evidence requirements change.

## Current Annotation Pass

The 2026-06-02 Figma annotation pass keeps the same 16 top-level sections on
primary page `129:5215`, records loaded non-UI page inventories, and keeps
coverage-bearing non-UI sections from pages `42:3014` and `0:1` distinct from
inventory-only legacy layers. Legacy Wireframe nodes for runtime, agent,
agent add/edit, daemon detail, and runtime detail are now explicitly absorbed by
the current runtime settings or agent settings UI entries instead of remaining
silent inventory-only frames. The task-thread section now also cites
motion references
(`node-id=153:12743`, `node-id=236:21467`), stop modals
(`node-id=236:20762`, `node-id=236:21048`), viewer-away rows
(`node-id=153:8749`, `node-id=236:21199`), and long-body focus rows
(`node-id=153:8592`, `node-id=236:20848`). These nodes are registered as
evidence so downstream docs can cite them deterministically, but they remain
client presentation facts unless a later SSOT changes the typed thread API.
The runtime settings endpoint-looking label (`node-id=129:17930`) is likewise
registered as evidence so downstream docs can reject it as a canonical base URL
or generated path.

## Verification

```bash
go test ./apicontract -run TestFigmaAIAgentCoverageManifest
go test ./...
```

The test verifies that each known primary UI top-level Figma node has an entry,
that the file's known pages and non-UI top-level evidence nodes are registered,
that legacy semantic Wireframe nodes point to the current UI entry that absorbed
them, that real decision sections link to SSOT documents and owner repos, that
documentation does not cite unregistered `node-id=...` values, and that
non-decision assets stay explicitly classified instead of silently becoming
work.
