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

## Stabilization Provenance

The executable manifest's `stabilized_by` list is the contracts-owned source
for downstream projection provenance. `riido-control-plane` and `riido-daemon`
may mirror this list to explain which coverage history their narrower local
projection consumed, but they must not invent a different upstream history.

Current stabilization history:

- `teamswyg/riido-contracts#38`
- `teamswyg/riido-contracts#39`
- `teamswyg/riido-contracts#45`
- `teamswyg/riido-contracts#46`
- `teamswyg/riido-contracts#51`
- `teamswyg/riido-contracts#52`
- `teamswyg/riido-contracts#54`
- `teamswyg/riido-contracts#55`
- `teamswyg/riido-contracts#56`
- `teamswyg/riido-contracts#57`
- `teamswyg/riido-contracts#58`
- `teamswyg/riido-contracts#60`
- `teamswyg/riido-contracts#62`
- `teamswyg/riido-contracts#63`
- `teamswyg/riido-contracts#64`
- `teamswyg/riido-contracts#65`
- `teamswyg/riido-contracts#66`

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

`figma-headless-file-key-placeholder.v1` records another Figma Plugin API
runtime limitation: a 2026-06-02 live `use_figma` root/category inspection for
file `MUOd9lctoEHASUStN3vUuK` returned `figma.fileKey=headless`, even though the
same invocation successfully read the AI Agent pages and annotation categories.
The authoritative file identity remains the Figma URL/tool input file key and
the manifest's `figma.file_key`, not the headless runtime placeholder. That
placeholder must not overwrite `figma.file_key`, `expected_pages`, or downstream
projection source identity.

`figma-onboarding-page-load-timeout.v1` records a separate 2026-06-02 tooling
limitation: current live reads of `node-id=42:3014` (`Wireframe - 온보딩`) can
time out after 120s when using Figma `get_metadata(nodeId=42:3014)` or
`use_figma` scripts that attempt `await figma.setCurrentPageAsync(page)`. Direct
Figma Plugin API `getNodeByIdAsync` lookups for registered nodes `236:33845` and
`236:33847` still returned the six onboarding `riido.*` `API Generated`
annotations. A timeout is supporting evidence only. It must not rewrite `expected_pages`,
remove page `42:3014`, remove onboarding inventory, or mark onboarding generated
paths unresolved; in other words, the timeout must not
become onboarding generated paths unresolved. The authoritative value for the
onboarding page remains the previously captured loaded Figma Plugin API
inventory, registered node-id evidence in this manifest, and direct registered
node lookup fallback for onboarding API Generated annotations.

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
| `42:3014` | Wireframe - 온보딩 | 83 | supporting onboarding/platform evidence |
| `0:1` | Wireframe | 28 | legacy planning evidence |

## API Generated Annotation Content

Figma Dev Mode annotations that start with `riido.` belong to category `700:0`
`API Generated`. Each annotation must keep this handoff shape:

1. `riido.*` generated-client facade path
2. `종류: Query | Mutation | SSE Stream`
3. `배경: ...` Korean frontend consumption background

The executable manifest keeps the durable `operation_kind` and `background`
values in `api_generated_annotation_inventory`; the live Figma annotation text is
handoff evidence, not a second API SSOT.
`operation_kind` is not free text: it must match the generated OpenAPI
transport. `text/event-stream` responses are `SSE Stream`, non-stream `GET`
operations are `Query`, and non-`GET` operations are `Mutation`.

The previous custom category `39:0` `클라이언트 전달` is retired. The 2026-06-02
inspection found zero annotations using it across the inspected pages. It still
exists as an unused Figma category definition because the current Figma MCP
category objects expose only data fields, not callable `remove` or `setLabel`
methods, so this automation could not delete the definition.

The 2026-06-02 lightweight Figma Plugin API annotation traversal found:

| Figma page | `riido.*` annotations | `API Generated` annotations | Missing kind | Missing background |
| --- | ---: | ---: | ---: | ---: |
| `129:5215` UI | 53 | 53 | 0 | 0 |
| `42:3014` Wireframe - 온보딩 | 6 | 6 | 0 | 0 |
| `0:1` Wireframe | 0 | 0 | 0 | 0 |

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
The same loaded inspection found page `42:3014` `Wireframe - 온보딩` has 83
top-level children. Several top-level layers have blank or repeated Figma names,
so the executable inventory keeps stable descriptive labels while preserving
the node IDs.

| Figma page | Figma node | Section | Status | Coverage rule |
| --- | --- | --- | --- | --- |
| `42:3014` | `164:30657` | 데스크탑앱 온보딩 | covered | absorbed by UI onboarding fixture/direct-create API coverage; desktop owns launch presentation |
| `42:3014` | `188:27707` | macOS | no-diff product surface | macOS install/launch presentation does not create an AI Agent endpoint |
| `42:3014` | `188:27708` | windows | no-diff product surface | Windows install/waitlist/launch-notification presentation stays outside AI Agent generated API |
| `42:3014` | `432:46849` | Ex AI - 온보딩 순서 변경 메모 | covered | revised onboarding order is client-local draft/config selection first, but durable v2 create still requires selected workspace and runtime at final submit |
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

## API Generated Annotations

Figma Dev Mode category `700:0` / `API Generated` is handoff evidence for
frontend generated-client consumption. It does not own API names. A Figma label
that starts with `riido.` is treated as a client facade example, and the durable
contracts generated path is the same value without the leading `riido.` prefix.
Current labels keep the facade path on the first line so Figma search works,
then add Korean handoff context:

```text
riido.aiAgent.tasks.stop
종류: Mutation
배경: 작업 중인 Agent에게 중지 요청을 보냅니다. daemon은 이 요청을 읽어 provider 실행을 강제 중지합니다.
```

| Figma node | Figma label | Canonical generated path | Resolution |
| --- | --- | --- | --- |
| `153:8545` | `riido.aiAgent.events.stream` / `종류: SSE Stream` | `aiAgent.events.stream` | The stream hint is a client facade example over the task-thread active stream handoff, covered by `153:15931`, API DSL/OpenAPI, and generated-client comments. |
| `236:20768` | `riido.aiAgent.tasks.stop` / `종류: Mutation` | `aiAgent.tasks.stop` | The stop hint is a client facade example over the task-thread stop operation, covered by `153:15931`, API DSL/OpenAPI, and generated-client comments. |

The executable manifest keeps two layers for this category. The canonical field
names are `api_generated_annotations` and
`api_generated_annotation_inventory`, matching the Figma category rather than
the older handoff wording. The first field keeps the two original
representative task-thread hints, while the inventory field registers the
screen-level inventory below. Those annotations are still evidence, not API
authority. Their first-line paths must already exist in the API DSL / OpenAPI
projection and generated TypeScript comments, and every group must carry
`Query`, `Mutation`, or `SSE Stream` plus Korean background text.

| UI area | Representative Figma nodes | Facade path | Kind | Background shown in Figma |
| --- | --- | --- | --- | --- |
| Participant dropdown / task details | `227:16920`, `227:17993`, `236:33845`, `236:33847` | `riido.aiAgent.tasks.assignableAgents` | Query | 참여자 드롭다운에서 현재 task/subtask에 배정할 수 있는 Agent 목록을 조회합니다. |
| Participant dropdown / task details | `227:16920`, `227:17993`, `236:33845`, `236:33847` | `riido.aiAgent.tasks.assign` | Mutation | 작업에 Agent를 참여자로 배정하고 daemon이 런타임으로 작업을 시작할 수 있는 서버 상태를 만듭니다. |
| Participant dropdown / task details | `227:16920`, `227:17993`, `236:33845`, `236:33847` | `riido.aiAgent.tasks.unassign` | Mutation | 참여자에서 Agent를 제거합니다. 진행 중이면 중지 요청/큐 해제 흐름으로 이어집니다. |
| Task thread | `153:8592`, `236:21379` | `riido.aiAgent.tasks.threads` | Query | 작업의 완료/진행 중 Agent thread cold collection을 조회합니다. active_stream이 있으면 SSE로 이어집니다. |
| Task thread | `153:8588` | `riido.aiAgent.tasks.threadMessages.create` | Mutation | 특정 thread_id에 다음 지시/답글을 추가하고 Agent 응답을 이어서 요청합니다. |
| Task thread | `153:8588` | `riido.aiAgent.tasks.submitComment` | Mutation | 호환용 댓글 제출 경로입니다. thread_id 없이 입력하면 서버가 적절한 thread 응답 흐름을 처리합니다. |
| Task thread | `153:8545`, `236:21379` | `riido.aiAgent.events.stream` | SSE Stream | threads 조회 결과에 active_stream이 있을 때만 연결해 진행 상태와 thread 갱신 이벤트를 받습니다. |
| Task thread | `236:20768` | `riido.aiAgent.tasks.stop` | Mutation | 작업 중인 Agent에게 중지 요청을 보냅니다. daemon은 이 요청을 읽어 provider 실행을 강제 중지합니다. |
| Runtime and agent settings | `160:10339`, `160:10418`, `559:34626`, `559:34704`, `164:34470`, `164:34476`, `164:34483`, `164:34496`, `435:72699`, `432:22231`, `432:35651`, `432:22617`, `432:35707` | `riido.aiAgent.devices.runtimes` | Query | 계정 소유 device에서 감지된 runtime 목록과 온라인/오프라인 상태를 조회합니다. 화면은 SaaS 값을 신뢰합니다. |
| Runtime settings | `160:10654`, `559:34714` | `riido.aiAgent.agents.daemon.details` | Query | Agent에 연결된 daemon/runtime 상세와 제어 가능 상태를 SaaS 기준으로 조회합니다. |
| Runtime settings | `160:14712`, `164:23904` | `riido.aiAgent.agents.daemon.stop` | Mutation | SaaS에 daemon 중지 요청을 남깁니다. daemon은 요청을 읽은 뒤 스스로 종료합니다. |
| Runtime settings | `160:16169`, `164:23977` | `riido.aiAgent.agents.daemon.restart` | Mutation | SaaS에 daemon 재시작 요청을 남깁니다. daemon은 polling으로 요청을 읽어 실행합니다. |
| Onboarding | `164:30672`, `164:30681`, `164:30690`, `164:30699` | `riido.aiAgent.onboarding.fixtures` | Query | 리도/영실/홍도/지원처럼 제품이 제공하는 초기값 목록을 조회합니다. template entity가 아니라 fixture입니다. |
| Onboarding | `164:33556` | `riido.aiAgent.onboarding.fixtures.createAgent` | Mutation | 선택한 fixture 값을 기반으로 일반 Agent를 생성합니다. fixture 자체를 생성하는 기능은 아닙니다. |
| Agent settings / direct setting | `337:24013`, `432:37349`, `134:6584`, `432:35493`, `164:30708`, `164:33556` | `riido.aiAgent.agents.create` | Mutation | 직접 설정 화면에서 워크스페이스 안에 새 Agent를 생성합니다. 신규 v2 create는 workspace_id를 포함합니다. |
| Agent settings | `337:24013`, `432:37349` | `riido.aiAgent.bootstrap` | Query | AI Agent 설정/온보딩 초기 화면에 필요한 agent 요약, 권한, 기본 상태를 조회합니다. |
| Agent settings | `417:21803`, `432:35544` | `riido.aiAgent.agents.updateConfiguration` | Mutation | 할당 작업이 없는 Agent의 이름, 썸네일, 설명, 지침, 런타임, 모델, 공개 범위를 저장합니다. |
| Agent settings | `417:21803`, `432:35544`, `432:38529`, `432:38855` | `riido.aiAgent.agents.editability` | Query | Agent를 수정할 수 있는지 먼저 조회합니다. 할당된 작업이 있으면 저장/수정 UI는 막혀야 합니다. |
| Agent settings | `432:37746`, `432:38689` | `riido.aiAgent.agents.delete` | Mutation | Agent 삭제를 요청합니다. 진행/예약 중 작업은 서버 정책에 따라 중지 또는 큐 해제됩니다. |

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
RIID-4857 adds loaded inventory closure for the `42:3014` onboarding page and
registers the `432:46849` order-change planning memo as a coverage-bearing
decision: the UI may let the user choose/configure an agent draft before runtime
and workspace selection, but contracts still expose only final create commands
that require the selected workspace/runtime context.

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
