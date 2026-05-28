# AI Agent Policy

> Riido task: RIID-4720 `[Contracts] AI Agent policy vocabulary and client API projection`

This document is the SSOT for the AI Agent policy vocabulary shared by
`riido-contracts`, `riido-control-plane`, `riido-daemon`, and clients.

## Product Evidence

The client surface follows the Figma v1.22 AI Agent flow:

- file: `v.1.22 AI Agent`
- planning node: `42:3014`
- UI dev handoff node: `129:5215`
- planning URL:
  `https://www.figma.com/design/MUOd9lctoEHASUStN3vUuK/v.1.22-AI-Agent?node-id=42-3014&p=f&m=dev`
- UI dev handoff URL:
  `https://www.figma.com/design/MUOd9lctoEHASUStN3vUuK/v.1.22-AI-Agent?node-id=129-5215&p=f&m=dev`

The planning flow includes onboarding, workspace selection, agent selection,
manual agent configuration, desktop-app prompting, and agent setting screens.
Those screens are consumed by both Riido web and the desktop app webview. The
UI dev handoff page adds Ready-for-dev surfaces for:

- task participant dropdowns where agents appear beside members
- task comment communication for queued, running, and stopped agent work
- menu placement for Riido AI, runtime, and agent management routes

The participant dropdown policy shown in the handoff is:

- member sorting belongs to the existing Riido member/client surface
- agent sorting belongs to this contract: owned agents first, then public
  agents visible through RBAC, with name ordering inside each group

The comment communication flow shows agent queue and stop states as task-thread
updates. The control-plane event contract therefore carries task context and a
typed comment-status value instead of asking clients to parse rendered text.
The client command contract also includes explicit task-thread comment submit
and stop actions so web and desktop webview clients do not infer AI Agent work
from generic task comments alone.

## Ubiquitous Language

| Term | Meaning |
| --- | --- |
| Runtime | A provider runtime installed on a device, such as Codex, Claude Code, Cursor, or OpenClaw. The device owner owns the runtime. |
| Agent | A task-assignable abstraction of a configured runtime created through Riido by a workspace admin or by the creator-owner. The creator owns the agent. |
| Control Plane | The Riido SaaS surface that applies the AI Agent feature to daemon and client workflows. |
| Device | The machine running `riido-daemon`, indirectly first signed in to Riido by the owning account. A device owns runtimes. |
| Daemon | The `riido-daemon` artifact controlled by the desktop app. It detects and controls runtimes and reports state to the control plane. |
| Client | Riido web or Riido desktop app webview inside a Riido workspace. |

## Policy Assertions

### Agent Deletion

When an agent is deleted from Riido web or desktop:

- queued tasks assigned to that agent are forcibly unassigned
- running tasks assigned to that agent are forcibly stopped
- deletion completes only after the control plane has applied those assignment
  effects

### Runtime Deletion

When a runtime is deleted on the device by the device user, local system, or
any PC-controlling actor:

- if at least one assigned agent references the runtime, the runtime becomes
  `offline`
- if no assigned agent references the runtime, the runtime is treated as
  absent, as if it never existed
- reinstalling or redetecting the same runtime on the same device makes the
  runtime follow only the online/offline policy; there are no extra exceptions

### Daemon Detection

The daemon detects runtimes immediately after start and then every 30 seconds.
Detection results are queued and reduced by the actor model. Locking is not
needed per runtime because same-runtime edits are time-sliced by the actor.

Detection errors or missed detections are client-visible as `offline`.

### Agent Editing

An agent is editable only when it has no assigned tasks. Clients must be able to
ask whether an agent can be edited before enabling edit controls, and the
control plane must also emit editability changes through the client event
stream.

Agent names are mutable and not unique.

### Runtime Output Parsing

Assigning an agent may instruct its runtime to emit parseable progress markers.
The parser can derive client-visible agent work status from runtime output. A
terminal marker such as `task-end` is allowed when it is owned by the runtime
prompt/adapter contract rather than inferred by the client.

## API Codegen Rule

Control-plane API enum values and sum-type variants are contract values, not
free text. They must be defined in the Domain DSL, preserved in the API IR, and
then projected into OpenAPI for client code generation.

OpenAPI is a generated projection. Clients may consume the OpenAPI projection
with Orval or a compatible generator, but generated client code is not owned by
this repository.

## Contract Fixtures

The client-facing contract fixture is
`control-plane-ai-agent-client-api.v1`:

- DSL: `../../apicontract/fixtures/control-plane-ai-agent-client.dsl.riido.json`
- IR: `../../apicontract/fixtures/control-plane-ai-agent-client.ir.riido.json`
- OpenAPI:
  `../../apicontract/fixtures/control-plane-ai-agent-client.openapi.json`

It covers:

- `GET /v1/client/ai-agent/bootstrap`
- `GET /v1/client/ai-agent/devices`
- `GET /v1/client/ai-agent/tasks/{task_id}/assignable-agents`
- `GET /v1/client/ai-agent/agents/{agent_id}/editability`
- `PATCH /v1/client/ai-agent/agents/{agent_id}`
- `DELETE /v1/client/ai-agent/agents/{agent_id}`
- `GET /v1/client/ai-agent/events`

The event stream uses a discriminated sum type, `ClientStreamEvent`, so client
codegen can produce safe branches for runtime snapshots, agent editability, and
agent work-status updates.

## Boundary

This document and `apicontract` own shared policy vocabulary, API enum values,
sum-type shape, operation ids, paths, and BDD scenario ids.

They do not own control-plane handlers, daemon runtime probing, client UI code,
Orval output, store implementations, task cancellation workers, or provider CLI
execution.
