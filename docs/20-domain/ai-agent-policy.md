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
- task thread communication for queued, running, and stopped agent work
- menu placement for Riido AI, runtime, and agent management routes
- runtime settings where clients show device/runtime liveness, attached agents,
  and local daemon lifecycle controls
- agent settings where clients add, list owned/public agents by device, and edit
  profile, runtime binding, visibility, and instruction fields

Participant dropdown evidence is `node-id=153-12742`
(`컴포넌트 참여자 드롭다운`). Its annotations state:

- member sorting is 가나다 order and belongs to the existing member/client
  surface
- agent sorting places viewer-owned agents first, then sorts names inside each
  ownership group
- long member/agent names, maximum dropdown height, scrollbar width, and visual
  overflow behavior are client UI composition details

The durable AI Agent contract fact is therefore only the agent side of that
dropdown: visible assignable agents are the viewer's owned agents plus other
users' public agents, and owned agents are ordered first.

Task thread annotation evidence is `node-id=153-15931` (`댓글 소통`).
The file contains Dev Mode annotation categories including `클라이언트 전달`.
In that section, `node-id=153-8545` cites `riido.aiAgent.events.stream`
and `node-id=236-20768` cites `riido.aiAgent.tasks.stop`. Those names are
generated-client consumption hints for the existing client event stream and stop
operation. The durable contract facts remain the API operation ids, paths,
typed stream variants, and task-thread policy defined below.

Menu placement evidence is `node-id=156-19307` (`메뉴바`). It shows
`Menubar/default` and `Menubar/setting` in dark and light variants. The durable
contract fact is only the route affordance: clients expose Riido AI, runtime,
and agent management entry points, while this document owns the shared terms
behind those entries. Visual menu state, ordering among unrelated product
routes, and concrete client route rendering are client-owned.

Runtime settings evidence is `node-id=162-23090` (`런타임 설정페이지`). Its
Dev Mode annotations call out an agent hover popover, a daemon stop modal, and a
restart-in-progress animation. The durable contract facts are:

- the runtime settings route consumes the existing device/runtime read model,
  `GET /v1/client/ai-agent/devices`, plus `device_runtime_snapshot` events
- the agent hover popover uses existing agent profile fields such as `name` and
  `description`; hover timing, layout, and truncation remain client-owned
- daemon stop/restart controls are local desktop/helper lifecycle controls for
  the viewer's device, not SaaS client API operations in this contract

Agent settings evidence is `node-id=164-50215` (`에이전트 설정페이지`) and
`node-id=134-6542` (`에이전트 추가`). The settings annotations call out long
one-line description UI, row/meatball edit entry points, absolute-time tooltip
behavior, runtime dropdown, and model dropdown. The add screen shows profile
photo, name, description, runtime, model, visibility, instruction, and save
controls. The durable contract facts are:

- the add screen needs a client-facing `POST /v1/client/ai-agent/agents`
  operation because the existing update/delete operations do not create an
  owned agent record
- agent rows need a server-authored `updated_at` date-time so clients can render
  list dates and absolute-time tooltips without inventing timestamps
- profile image, name, description, runtime binding, visibility, and instruction
  are the current editable configuration fields
- runtime dropdown candidates come from existing runtime/device read-model data;
  clients own the dropdown rendering
- row click, meatball edit entry, description truncation/wrapping, and timestamp
  formatting are client-owned
- model dropdown labels shown in Figma are planning/UI evidence only until
  `Q-CON-006` settles the runtime model catalog owner

The participant dropdown policy shown in the handoff is:

- member sorting belongs to the existing Riido member/client surface
- agent sorting belongs to this contract: owned agents first, then public
  agents visible through RBAC, with name ordering inside each group

The task thread flow shows agent queue and stop states as task-thread
updates. The control-plane event contract therefore carries task context and a
typed comment-status value instead of asking clients to parse rendered text.
The client command contract also includes explicit task-thread comment submit
and stop actions so web and desktop webview clients do not infer AI Agent work
from generic task comments alone.

## SSOT Dependency Direction

This document owns the shared AI Agent mental model. Downstream repositories may
repeat these terms only as local execution or projection behavior. The cross-SSOT
dependency direction and the top-down / bottom-up harness loop are defined in
[`../30-architecture/ssot-dependency-map.md`](../30-architecture/ssot-dependency-map.md).

For agent settings specifically:

- `profile_thumbnail_url`, `description`, and `instruction` meaning starts here
  and in the `control-plane-ai-agent-client-api.v1` DSL fixture.
- `riido-control-plane` owns HTTP validation, save/update behavior, mock data,
  and generated-client handoff.
- `riido-daemon` owns only runtime consumption of an assigned instruction value;
  it does not own thumbnail presentation, storage, RBAC, or API shape.
- `riido-infra` owns deployment/storage changes only when a future media,
  secret, durability, or topology requirement appears.
- Figma menu placement (`node-id=156-19307`) does not introduce a new contract
  endpoint by itself. Route labels and visual selected state are client-owned;
  this contract owns only the vocabulary and API facts used after those routes
  are opened.
- Figma task-thread annotations (`node-id=153-15931`) can cite generated-client
  call paths such as `riido.aiAgent.events.stream` and
  `riido.aiAgent.tasks.stop`, but this repo owns their canonical operation
  ids/path/typed-event meaning through the DSL/IR/OpenAPI fixture, not the
  client chain syntax or UI micro-interactions.
- Figma participant dropdown annotations (`node-id=153-12742`) can cite sort and
  overflow behavior, but this repo owns only AI Agent visibility and owned-first
  agent ordering. Member sorting and visual dropdown constraints are client
  surface facts.
- Figma runtime settings annotations (`node-id=162-23090`) can cite runtime
  liveness, agent hover, daemon stop modal, and restart animation behavior. This
  repo owns only the device/runtime read-model policy and the fact that a local
  daemon stop eventually makes affected runtimes offline through the existing
  liveness policy. Client hover/modal/animation behavior and local helper
  command composition are downstream facts.

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

### Runtime Settings

The runtime settings surface is a client composition over the control-plane
device/runtime read model and, for the viewer's current desktop device, the
local daemon/helper control surface.

The control-plane client API exposes device/runtime liveness and attached agent
metadata; it does not expose a SaaS command that stops or restarts a user's
daemon. When a desktop client stops its local daemon, the resulting control-plane
effect is represented through the same runtime liveness policy: affected
runtimes become `offline` when heartbeat/detection state is missing or expired.

Restart is not a distinct contract operation. A client or helper may compose it
from local daemon lifecycle controls, while this contract only requires that the
subsequent device/runtime read model converges to the current online/offline
state.

### Agent Editing

An agent is editable only when it has no assigned tasks. Clients must be able to
ask whether an agent can be edited before enabling edit controls, and the
control plane must also emit editability changes through the client event
stream.

Agent names are mutable and not unique.

### Agent Profile Configuration

Agent profile presentation belongs to the agent configuration, not to a task
thread or runtime. The client-facing agent record therefore carries optional
`profile_thumbnail_url`, `description`, and `instruction` fields wherever
agents are returned.

The thumbnail value is an HTTPS image URL string. Binary upload, image
resizing, CDN storage, and moderation are outside this contract until a separate
media/storage contract owns them.

The description value is client-authored text used as a short, one-line agent
summary in agent list and edit surfaces. Empty text is allowed. The current
client API limit is 160 characters. Longer values are rejected by the control
plane before the agent configuration is saved. UI truncation/wrapping policy is
owned by the client and must not change the stored value.

The instruction value is client-authored text that is saved with the agent and
used by the control plane/daemon when composing runtime prompts. Empty text is
allowed. The current client API limit is 1000 characters. Longer values are
rejected by the control plane before the agent configuration is saved.

Profile field creation and updates follow the same RBAC and mutation safety
rules as name, visibility, and runtime binding updates. Creation stamps
`owner_principal_id` from the authorized principal and binds only a selected
viewer-owned runtime. After creation, admin may mutate all agents, owner may
mutate owned agents, and no agent can be edited while it has assigned tasks.

### Agent Update Timestamp

The agent client record carries a required `updated_at` date-time owned by the
control plane. It changes when editable agent configuration is saved and is used
by clients for agent-list update dates and absolute-time tooltips. Clients own
relative/absolute formatting and tooltip presentation; they must not synthesize
or rewrite the stored timestamp.

The timestamp is distinct from runtime heartbeat, daemon liveness, and provider
session progress time. Runtime freshness remains owned by the device/runtime
read model.

### Runtime Model Dropdown

The Figma agent setting screen shows a model dropdown with provider-specific
labels. Those labels are not contract enum values today. This contract does not
add `model_id` to agent configuration until the runtime model catalog ownership
question is resolved in
[`../50-roadmap/open-questions.md#q-con-006`](../50-roadmap/open-questions.md#q-con-006).

Daemon adapters may accept a model value only as part of an already-authorized
runtime execution request. They do not own the catalog of values exposed to
clients.

### Runtime Output Parsing

Assigning an agent may instruct its runtime to emit parseable progress markers.
The parser can derive client-visible agent work status from runtime output. A
terminal marker such as `task-end` is allowed when it is owned by the runtime
prompt/adapter contract rather than inferred by the client.

### Task Thread Progress Batching

Runtime progress visible in the client task thread is derived only from explicit
runtime progress markers, such as `<riido_log>...<end>`. The daemon must not
relay provider raw token streams directly to SaaS or clients.

The daemon batches parsed progress lines before SaaS ingest. The default cadence
is 10 seconds while a task is running, and pending progress is flushed before a
terminal assignment event. The control plane fans accepted batches out through
the client event stream as the typed `agent_thread_progress` variant of
`ClientStreamEvent`.

Clients render task-thread progress from that typed payload and never parse
provider output text.

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
- `POST /v1/client/ai-agent/tasks/{task_id}/comments`
- `POST /v1/client/ai-agent/tasks/{task_id}/stop`
- `POST /v1/client/ai-agent/agents`
- `GET /v1/client/ai-agent/agents/{agent_id}/editability`
- `PATCH /v1/client/ai-agent/agents/{agent_id}`
- `DELETE /v1/client/ai-agent/agents/{agent_id}`
- `GET /v1/client/ai-agent/events`

The event stream uses a discriminated sum type, `ClientStreamEvent`, so client
codegen can produce safe branches for runtime snapshots, agent editability, and
agent work-status updates. Runtime progress intended for the task thread is the
`agent_thread_progress` variant and carries ordered progress lines.

## Boundary

This document and `apicontract` own shared policy vocabulary, API enum values,
sum-type shape, operation ids, paths, and BDD scenario ids.

They do not own control-plane handlers, daemon runtime probing, client UI code,
Orval output, store implementations, task cancellation workers, or provider CLI
execution.
