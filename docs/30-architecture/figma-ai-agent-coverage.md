# Figma AI Agent Coverage

> Riido task: RIID-4809 `[Contracts] Figma AI Agent 화면 커버리지 SSOT manifest`

This document explains how the Figma `v.1.22 AI Agent` UI page is absorbed
into Riido SSOT documents. Figma is evidence. It is not the durable SSOT for
domain policy, API shape, daemon behavior, or infra topology.

The executable coverage manifest is
[`figma-ai-agent-coverage.riido.json`](figma-ai-agent-coverage.riido.json). It
records the Figma file, page, top-level sections, owner repositories, generated
client paths, and the top-down / bottom-up rule for each section.

## Source

| Field | Value |
| --- | --- |
| Figma file | `v.1.22 AI Agent` |
| File key | `MUOd9lctoEHASUStN3vUuK` |
| UI page node | `129:5215` |
| Inspection date | `2026-06-01` |

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

## Top-Level Coverage

| Figma node | Section | Status | Owning boundary |
| --- | --- | --- | --- |
| `153:12742` | 컴포넌트 참여자 드롭다운 | covered | contracts policy, API DSL, control-plane, client |
| `153:15931` | 댓글 소통 | covered | task-thread policy, API DSL, control-plane, daemon, client |
| `153:15932` | image 14 | non-decision asset | no independent SSOT decision |
| `153:15935` | 추가 기획 내용 | covered | contracts policy and API DSL before implementation |
| `162:23468` | 논의 필요 | planning evidence | open question or explicit owner required before code |
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

## Generated Client Path Hints

The manifest intentionally stores generated paths that frontend developers
would search for in generated TypeScript comments. These are projections of the
API DSL, not hand-authored route names.

| UI area | Main generated paths |
| --- | --- |
| Participant dropdown | `aiAgent.tasks.assignableAgents`, `aiAgent.tasks.assign`, `aiAgent.tasks.unassign`, `v2.aiAgent.tasks.assignableAgents`, `v2.aiAgent.tasks.assign`, `v2.aiAgent.tasks.unassign` |
| Task thread | `aiAgent.tasks.threads`, `aiAgent.tasks.threadMessages.create`, `aiAgent.tasks.stop`, `aiAgent.events.stream`, `v2.aiAgent.tasks.threads`, `v2.aiAgent.tasks.threadMessages.create`, `v2.aiAgent.tasks.stop`, `v2.aiAgent.events.stream` |
| Runtime settings | `aiAgent.devices.runtimes`, `aiAgent.agents.daemon.details`, `aiAgent.agents.daemon.*`, `v2.aiAgent.devices.runtimes`, `v2.aiAgent.agents.daemon.details`, `v2.aiAgent.agents.daemon.*` |
| Onboarding | `aiAgent.onboarding.fixtures`, `aiAgent.onboarding.fixtures.createAgent`, `aiAgent.agents.create`, `v2.aiAgent.onboarding.fixtures`, `v2.aiAgent.onboarding.fixtures.createAgent`, `v2.aiAgent.agents.create` |
| Agent settings | `aiAgent.bootstrap`, `aiAgent.agents.create`, `aiAgent.agents.updateConfiguration`, `aiAgent.agents.delete`, `aiAgent.agents.editability`, `v2.aiAgent.bootstrap`, `v2.aiAgent.agents.*` |

## Ownership Notes

- `docs/20-domain/ai-agent-policy.md` owns the durable business language:
  agent, runtime, device, daemon, visibility, editability, task-thread, fixture,
  workspace-scoped v2 API, and no-runtime semantics.
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

## Verification

```bash
go test ./apicontract -run TestFigmaAIAgentCoverageManifest
go test ./...
```

The test verifies that each known top-level Figma node has an entry, that real
decision sections link to SSOT documents and owner repos, and that
non-decision assets stay explicitly classified instead of silently becoming
work.
