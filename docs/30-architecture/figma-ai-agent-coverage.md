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
| Inspection date | `2026-06-02` |

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
| `0:1` | Wireframe | 1 | legacy planning evidence |

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

| Figma page | Figma node | Section | Status | Coverage rule |
| --- | --- | --- | --- | --- |
| `42:3014` | `164:30657` | 데스크탑앱 온보딩 | covered | absorbed by UI onboarding fixture/direct-create API coverage; desktop owns launch presentation |
| `42:3014` | `188:27707` | macOS | no-diff product surface | macOS install/launch presentation does not create an AI Agent endpoint |
| `42:3014` | `188:27708` | windows | no-diff product surface | Windows install/waitlist/launch-notification presentation stays outside AI Agent generated API |
| `0:1` | `153:15934` | 추가 기획 내용 | covered | legacy planning evidence resolved by UI page `153:15935` and task/subtask assignment scope |

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
primary page `129:5215` and additionally registers all non-UI top-level
sections from pages `42:3014` and `0:1`. The task-thread section now also cites
motion references
(`node-id=153:12743`, `node-id=236:21467`), stop modals
(`node-id=236:20762`, `node-id=236:21048`), viewer-away rows
(`node-id=153:8749`, `node-id=236:21199`), and long-body focus rows
(`node-id=153:8592`, `node-id=236:20848`). These nodes are registered as
evidence so downstream docs can cite them deterministically, but they remain
client presentation facts unless a later SSOT changes the typed thread API.

## Verification

```bash
go test ./apicontract -run TestFigmaAIAgentCoverageManifest
go test ./...
```

The test verifies that each known primary UI top-level Figma node has an entry,
that the file's known pages and non-UI top-level evidence nodes are registered,
that real decision sections link to SSOT documents and owner repos, that
documentation does not cite unregistered `node-id=...` values, and that
non-decision assets stay explicitly classified instead of silently becoming
work.
