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
- runtime settings empty states where clients show missing-current-device
  runtime, no daemon, provider install-card, and Windows app waitlist variants
- agent settings where clients add, list owned/public agents by device, and edit
  profile, runtime binding, visibility, and instruction fields

Onboarding evidence from `node-id=42-3014` includes:

- `node-id=137-6746`: choose the runtime used to create the first agent; the
  inspected UI shows Claude Code and Codex as `감지됨` selectable rows, while
  OpenClaw and Cursor Agent are `감지 안 됨` non-selectable rows
- `node-id=138-7389`: choose an agent template or choose direct configuration;
  the inspected UI shows the starter template rows `리도`, `영실`, `홍도`,
  `지원`, followed by a `직접 설정` row, a pre-selection `다음` button, and a
  right-side preview skeleton
- `node-id=164-26969`: direct configuration is annotated `직접 설정 선택 시
  스크롤`; it dims the starter template rows and expands `직접 설정` into
  `이름` (`리도` placeholder), `설명` (`문제 정의부터 우선순위, 출시 계획까지
  정리합니다.` placeholder), and `지침` (`기능 요청을 문제·목표·성공 기준으로
  재정의하고 PRD, 우선순위, 로드맵, 출시 계획을 구조화합니다. 아이디어는
  가설로 다루며 불확실한 내용은 [확인 필요]로 표시합니다.` placeholder) fields
  with a scroll affordance
- `node-id=164-30192`: workspace selection list shows the selected workspace,
  a scroll affordance, and a `새 워크스페이스` row
- `node-id=164-27719`: template descriptions show up to two lines before
  ellipsis; this is client presentation only
- `node-id=164-30206`: if no selectable AI runtime is installed/detected, the
  flow shows the no-runtime start state with Claude Code, Codex, OpenClaw, and
  Cursor Agent all marked `연결 안 됨` and a `시작하기` CTA

The durable contract fact is that onboarding agent templates are API data, not
frontend hard-coded business copy. Runtime/no-runtime branching is still client
composition over the existing device/runtime read model. Workspace list
selection and the `새 워크스페이스` row are workspace/team/client product
surfaces; they do not add an AI Agent generated operation by themselves.
The all-disconnected provider list and start CTA are also presentation derived
from device/runtime liveness; they do not create provider-install or
provider-start commands in this contract.

Web onboarding evidence is `node-id=236-29749` (`웹 온보딩`). Its frames and
Dev Mode annotations cover macOS app download CTA, sign-up entry points, terms
consent, email sign-up, member invite, Windows waitlist/marketing-consent
variants, chat animation reference, and button progress-bar reference. Those
screens do not add a new AI Agent contract operation by themselves:

- sign-up, login, Google-auth terms consent, email/password validation, and
  member invitation belong to existing auth/team/client product surfaces
- macOS app download and Windows waitlist CTAs are distribution/product facts,
  not runtime/provider install commands
- terms row default-open/closed behavior, row click target, chat animation, and
  progress-bar animation are client presentation facts
- waitlist and marketing-consent mutation ownership is resolved by
  `Q-CON-007` as no-diff for the AI Agent generated API

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

Additional planning evidence is `node-id=153-15935` (`추가 기획 내용`). The
section has no annotation nodes; the visible planning text is the evidence. It
confirms the cross-surface assignment target scope and several non-target
surfaces:

- AI Agent assignment is available only on Riido task and subtask surfaces.
- Existing Riido AI property filling must not recommend or select agents.
- Agent mention is not a supported command surface.
- Device/runtime actions are mediated by the agent access model: public agents
  delegate indirect daemon/runtime execution to workspace members, while private
  agents delegate it only to admins and owners.

The durable contract fact is the target scope, not the page drawing. The
`task_id` path parameter in AI Agent task APIs means a Riido task or subtask id.
Projects, milestones, intakes, generic comments, AI property filling, and
mention surfaces must not reuse task assignment/comment/stop/thread endpoints
without a future owning SSOT and a newly generated operation.

Task thread annotation evidence is `node-id=153-15931` (`댓글 소통`).
The file contains Dev Mode annotation categories including `클라이언트 전달`.
In that section, `node-id=153-8545` cites `riido.aiAgent.events.stream`
and `node-id=236-20768` cites `riido.aiAgent.tasks.stop`. Those names are
generated-client consumption hints for the existing client event stream and stop
operation. The durable contract facts remain the API operation ids, paths,
typed stream variants, and task-thread policy defined below.
The normal task screen `node-id=236-21379` shows the same boundary in one task
view: a generic task comment input, an AI Agent thread row, a thread reply
input, and a `중지` action. That screen is evidence for composing the existing
thread cold collection, explicit AI Agent comment submit, and explicit stop
operation; it is not a second source for task details, sidebar fields, or client
layout.

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
  `GET /v1/client/ai-agent/devices`, the agent-bound daemon detail endpoint,
  daemon command endpoints, plus device/runtime and daemon status events
- `내 기기` and `다른 기기` groups are both composed from ordinary device and
  runtime records; current-device grouping is a client/session selection, not a
  new API resource class
- the agent hover popover uses existing agent profile fields such as `name` and
  `description`; hover timing, layout, and truncation remain client-owned
- daemon details are SaaS read-model facts exposed through
  `riido.aiAgent.agents.daemon.details`; these include online/offline status,
  uptime, PID, daemon ID, profile, device name, and control state
- daemon start/restart/stop are SaaS command requests exposed through generated
  client endpoints under `riido.aiAgent.agents.daemon.*`;
  the desktop daemon still executes local lifecycle behavior after reading the
  accepted command
- daemon stop acceptance makes that device's runtimes client-visible as
  `offline` through the runtime read model and stream updates

Agent settings evidence is `node-id=164-50215` (`에이전트 설정페이지`),
`node-id=134-6542` (`에이전트 추가`), `node-id=337-24001` (`에이전트`
settings list), and `node-id=432-35713` (`에이전트` list). The settings
annotations call out long one-line description UI, row/meatball edit entry
points, absolute-time tooltip behavior, runtime dropdown, model dropdown,
required add/edit form controls (`node-id=417-21803` / `node-id=432-35544`),
instruction input scroll behavior (`node-id=432-23265` / `node-id=432-35595`),
long device-name dropdown rows (`node-id=I432:22235;6885:15212` /
`node-id=I432:35655;6885:15212`), delete confirmation modals
(`node-id=432-37740` / `node-id=432-38683`), and disabled edit menu
presentation for working agents (`node-id=432-37900` / `node-id=432-38853`).
The list screens add created/update columns, optional-description rows,
online/working/offline labels, edit/delete menu entry points, and the
`node-id=337-24013` rule that the `에이전트 추가` affordance is hidden when no
member-visible runtime is selectable. The add screen shows profile photo, name,
description, runtime, model, visibility, instruction, and save controls. The
durable contract facts are:

- the add screen needs a client-facing `POST /v1/client/ai-agent/agents`
  operation because the existing update/delete operations do not create an
  owned agent record
- agent rows need server-authored `created_at` and `updated_at` date-times so
  clients can render list dates and absolute-time tooltips without inventing
  timestamps
- profile image, name, description, runtime binding, visibility, and instruction
  are the current editable configuration fields
- Figma marks name, runtime, model, and visibility as required form controls;
  API-level required fields are `name`, `runtime_id`, and `visibility`, while
  omitted `model_id` resolves to the selected runtime's default model
- runtime dropdown candidates come from existing runtime/device read-model data;
  clients own the dropdown rendering
- the add affordance visibility is client presentation over the authorized
  device/runtime read model; hiding the button does not create a separate
  eligibility endpoint and does not let clients bypass create validation
- model dropdown candidates come from `RuntimeRecord.models` as runtime-scoped
  catalog records; `model_id` is opaque per runtime and model labels are
  display data, not generated enum values
- row click, meatball edit entry, delete-confirmation modals, disabled edit
  tooltip/cursor behavior, input scroll limits, long device-name presentation,
  description truncation/wrapping, status-label copy/color, and timestamp
  formatting are client-owned

The participant dropdown policy shown in the handoff is:

- member sorting belongs to the existing Riido member/client surface
- agent sorting belongs to this contract: owned agents first, then public
  agents visible through RBAC, with display-name ordering inside each group
- agent names are mutable and non-unique, so equal display names use
  `agent_id` as the stable tie-breaker. Clients render the display name, but
  must not treat it as identity.

The same handoff marks participant-dropdown presentation constraints at
`node-id=153-12742`: long member/agent names must remain visually contained,
the dropdown's visible height caps at 520px, and a visible scrollbar may expand
the rendered width from 240px to 254px. These are client presentation
constraints. The API contract only guarantees stable identifiers and the agent
ordering above; it does not encode pixel sizes.

The task thread flow shows agent queue and stop states as task-thread
updates. The control-plane event contract therefore carries task context and a
typed comment-status value instead of asking clients to parse rendered text.
The client command contract also includes explicit task-thread comment submit
and stop actions so web and desktop webview clients do not infer AI Agent work
from generic task comments alone.
Selecting an agent from the task participant dropdown is a separate generated
assignment command, not a generic task participant mutation. The command creates
the initial AI Agent task-thread row with typed
`comment_kind=assignment_started` unless the selected agent is already busy, in
which case the existing queued tuple applies. Removing an active or queued agent
from task participants uses the generated unassign command, which stops the
agent work with `comment_kind=stopped_by_user_request`; hiding stopped content is
client-owned presentation around the returned typed status.
In the normal active screen (`node-id=236-21379`), generic task comments and
AI-Agent-directed thread replies are visually adjacent. The contract boundary is
that AI-Agent-directed messages use the explicit comments operation with
`agent_id` and optional `source_comment_id`; generic task comments remain the
existing task product surface until routed through that operation.
In the busy-agent screen (`node-id=153-8761`), assigning/commenting to an agent
that is already working creates a queued task-thread row. The contract fact is
the typed status tuple `comment_kind=queued_by_busy_agent`,
`assignment_state=queued`, and `work_status=queued`; the Korean copy shown in
Figma is client presentation around that tuple. The visible `중지` affordance
continues to use the explicit task stop operation and does not create a second
cancel endpoint.
In the stopped-by-deleted-agent screen (`node-id=227-19354`, `작업 중지`),
Figma shows a Riido-authored task-thread row with the copy "에이전트가 삭제되어
진행 중이던 작업이 중지됐어요." The contract fact is not that copy. The canonical
fact is that deleting an agent with queued or running assignments force-stops
those assignments and projects a typed task-thread status with
`comment_kind=stopped_by_agent_deleted`, `assignment_state=stopped`, and a
terminal agent work status. This reuses `DELETE
/v1/client/ai-agent/agents/{agent_id}` as the command. It does not introduce a
new client stop endpoint; the exact Korean message, `리도` actor label, "방금 전"
timestamp, hidden stop affordance state, avatar, and row layout remain
client/task presentation over the typed status.

## SSOT Dependency Direction

This document owns the shared AI Agent mental model. Downstream repositories may
repeat these terms only as local execution or projection behavior. The cross-SSOT
dependency direction and the top-down / bottom-up harness loop are defined in
[`../30-architecture/ssot-dependency-map.md`](../30-architecture/ssot-dependency-map.md).

For agent settings specifically:

- `profile_thumbnail_url`, `description`, and `instruction` meaning starts here
  and in the `control-plane-ai-agent-client-api.v1` DSL fixture.
- `model_id` meaning starts here and in the same DSL fixture. It is validated
  against the selected runtime's `RuntimeModelRecord` catalog and defaults to
  the selected runtime's default model when omitted.
- onboarding template catalog meaning starts here and is projected through the
  same client bootstrap fixture so clients can render template names,
  descriptions, role labels, thumbnails, and copyable instructions without
  owning the template source text.
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
  client chain syntax or UI micro-interactions. Long-body scroll/focus,
  viewer-away notification rendering, hover buttons, stop modals, and progress
  animations are client presentation facts around returned thread records.
- Figma normal task-thread screen (`node-id=236-21379`) can cite that a task
  screen needs both user-to-agent comment submission and active-thread stop
  affordances. This repo already owns those as
  `POST /v1/client/ai-agent/tasks/{task_id}/comments` and
  `POST /v1/client/ai-agent/tasks/{task_id}/stop`; the right-side task details
  panel, reply input rendering, send button state, and agent row layout remain
  client/task surface facts.
- Figma queued task-thread screen (`node-id=153-8761`) can cite the busy-agent
  auto-comment copy and `중지` affordance. This repo owns only the typed queued
  semantics: `queued_by_busy_agent`, `assignment_state=queued`, and
  `work_status=queued`, plus reuse of the existing stop action. The exact copy,
  timestamp wording, row layout, and agent avatar rendering are client/task
  presentation facts.
- Figma stopped-by-deleted-agent task-thread screen (`node-id=227-19354`) can
  cite the stopped row shown after an assigned agent is deleted. This repo owns
  only the typed forced-stop semantics: `stopped_by_agent_deleted`,
  `assignment_state=stopped`, and reuse of the existing agent delete operation.
  The exact Korean copy, Riido actor label, timestamp wording, row layout,
  hidden action state, and agent/avatar rendering are client/task presentation
  facts.
- Figma participant dropdown annotations (`node-id=153-12742`) can cite sort and
  overflow behavior, but this repo owns only AI Agent visibility and owned-first
  agent ordering. Member sorting and visual dropdown constraints are client
  surface facts.
- Figma additional planning section (`node-id=153-15935`) can cite agent-bound
  device/runtime actions, no agent recommendation in existing AI property
  filling, no agent mention command surface, and task/subtask-only assignment.
  This repo owns the target-scope contract fact and the access fact that public
  agents delegate indirect daemon/runtime execution to workspace members while
  private agents delegate it only to admins and owners. Project, milestone,
  intake, property-filling, and mention surfaces need a separate owning SSOT
  before they become generated AI Agent operations.
- Figma runtime settings annotations (`node-id=162-23090`) can cite runtime
  liveness, agent hover, daemon stop modal, and restart animation behavior. This
  repo owns the device/runtime read-model policy, agent-bound daemon detail
  projection, agent-bound daemon start/restart/stop command contract, and the
  fact that a daemon stop makes affected runtimes offline through the existing
  liveness policy. Client hover/modal/animation behavior is downstream
  presentation.
- Figma runtime settings empty-state annotations (`node-id=275-22731`) can cite
  provider install cards, hover states, Windows app waitlist copy, and marketing
  consent presentation. This repo owns only the device/runtime liveness data
  used to decide whether the current device, daemon, and runtime rows are
  present. Provider installation URLs, waitlist subscription, and marketing
  consent mutation need a separate owning SSOT before they become generated API
  operations.
- Figma onboarding annotations (`node-id=42-3014`) can cite scroll, workspace
  selector list behavior (`node-id=164-30192`), two-line ellipsis,
  no-installed-AI start behavior (`node-id=164-30206`), and direct-setting
  expansion. This repo owns only the onboarding template catalog data shape and
  the runtime liveness facts used by clients to choose which step to show.
- Figma web onboarding annotations (`node-id=236-29749`) can cite sign-up,
  terms, member-invite, app-download, waitlist, and animation behavior. This repo
  owns none of those as AI Agent contract facts until a separate owning SSOT
  promotes a durable auth/team/distribution/waitlist operation. Current AI Agent
  contract facts stay `agent_templates`, agent create/update/read models, and
  device/runtime liveness.

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
device/runtime read model and the agent-bound daemon read/command model.

The control-plane client API exposes device/runtime liveness and attached-agent
metadata for current-device, other-device, and agent-accessible groups through
`GET /v1/client/ai-agent/devices`. The runtime settings screen reads daemon
detail through `GET /v1/client/ai-agent/agents/{agent_id}/daemon`; this
response owns the SaaS-visible daemon facts needed by the settings page:
online/offline status, uptime, PID, daemon ID, profile, device name, supported
actions, and current control state.

Start, restart, and stop are distinct SaaS command requests:

- `POST /v1/client/ai-agent/agents/{agent_id}/daemon/start`
- `POST /v1/client/ai-agent/agents/{agent_id}/daemon/restart`
- `POST /v1/client/ai-agent/agents/{agent_id}/daemon/stop`

Those commands are not direct local socket calls from the frontend. The client
asks the SaaS control plane, the daemon reads the accepted command, and the
daemon performs local lifecycle control. The generated frontend path is
`riido.aiAgent.agents.daemon.details/start/restart/stop`.

Daemon detail and daemon commands follow the agent access boundary. Publishing
an agent as public delegates indirect daemon/runtime execution and daemon
command requests to workspace members for that agent. Keeping an agent private
limits the same path to workspace admins and the agent owner. Knowing a
`device_id` alone is never enough to inspect or control a daemon.

When a stop command is accepted, affected runtimes on that device become
client-visible as `offline` through the same runtime liveness read model. The
stream also includes typed daemon status changes so clients can render
`재시작 중`, `중지`, or `시작` affordances without polling the local daemon.

### Agent Onboarding

The AI Agent onboarding flow is a client composition over bootstrap, device
runtime data, and agent creation.

The runtime selection step from `node-id=137-6746` is composed from
`DeviceRecord.runtimes`. A runtime is selectable for onboarding only when the
client can submit its `runtime_id` to agent creation and the read model marks
that runtime `availability=online` and `detection_state=detected`. Korean
labels such as `감지됨` / `감지 안 됨`, radio state, row dimming, and `다음` /
`다음에 하기` button presentation are client-owned rendering. The control
plane still validates the selected `runtime_id` at agent create/update time.

The control-plane bootstrap response carries an ordered onboarding template
catalog. A template is a copyable default for an agent configuration and
contains `template_id`, `name`, optional `role_label`, optional
`profile_thumbnail_url`, `description`, and `instruction`. Template text and
instruction defaults are contract data so frontend clients do not hard-code the
behavioral meaning of starter agents.

The template-selection step from `node-id=138-7389` is projected from
`ClientBootstrapResponse.agent_templates` in response order. Current Figma
evidence shows four starter rows, `리도`, `영실`, `홍도`, and `지원`, but the
rows are not frontend-owned copy. The `직접 설정` row is not an
`AgentOnboardingTemplate`; it is a client presentation entry that lets the user
continue to explicit agent configuration. The pre-selection disabled `다음`
button and the right-side preview skeleton/popover are also client presentation
over the selected or unselected template state.

Selecting a template does not create a separate domain entity. The client still
creates an agent through `POST /v1/client/ai-agent/agents` with selected runtime,
visibility, and copied profile fields. Direct configuration uses the same create
operation without choosing a template.

The direct-configuration expansion from `node-id=164-26969` also does not create
a separate command or template row. The expanded `이름`, `설명`, and `지침`
fields project to `CreateAgentConfigurationRequest.name`, `description`, and
`instruction`. The previously selected runtime continues to supply `runtime_id`,
and visibility uses the create request policy/default chosen by the client. The
dimmed starter-template rows, scroll bar, placeholder copy, and expanded-row
layout are client presentation facts.

If no selectable runtime is online/detected for the viewer, the client skips the
template-selection and direct-setting steps and shows the no-installed-AI start
state from the planning screen. That branch does not introduce a new control
plane command. It is derived from the existing runtime/device read model.
In `node-id=164-30206`, the client can still show provider rows for Claude Code,
Codex, OpenClaw, and Cursor Agent as `연결 안 됨` and let the user continue with
`시작하기`; those rows and CTA are not executable provider install/start
operations.

Runtime settings empty states from `node-id=275-22731` use the same read model.
When the current device has no daemon, no runtime, or no selectable current
runtime, the API does not add a provider-install command. Clients can render
provider install cards and hover states from product copy and external provider
links, but provider CLIs remain external user-installed tools. The Windows app
waitlist and marketing-consent button states shown in the Figma section are not
part of this AI Agent client API. `Q-CON-007` fixes this as a no-diff decision:
they require a separate product/marketing SSOT and a separate generated
operation before any API helper can exist.

Web onboarding from `node-id=236-29749` reinforces that boundary. The macOS app
download CTA is a product/distribution route to a desktop artifact, not a
provider CLI install command. The Windows launch-notification and
marketing-consent variants follow the same `Q-CON-007` no-diff decision as
`node-id=275-22731`. Google sign-up terms consent, email sign-up terms rows,
member invite input/link-copy, and animation references are client/auth/team
presentation surfaces and do not change the AI Agent DSL/IR/OpenAPI projection.

Workspace selection, workspace creation entry points, template row selection,
`직접 설정` row rendering, disabled-next state before selection, scroll
behavior, description ellipsis, and preview-popover layout remain client-owned
presentation behavior.

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
copied into `Assignment.agent_instruction` when the control plane creates an
agent assignment. Empty text is allowed. The current client API limit is 1000
characters. Longer values are rejected by the control plane before the agent
configuration is saved. Assignment-time snapshotting is intentional: editing an
agent after a task has been assigned does not rewrite already-queued or running
runtime instructions.

This contract treats `instruction` as provider-neutral agent guidance. It does
not decide whether Claude, Codex, OpenClaw, Cursor, or a future runtime obeys
the text more effectively through a prompt prefix, system prompt, native config,
or another provider surface. That placement/effectiveness strategy is owned by
the daemon provider-runtime SSOT. Cross-provider evidence must consume the
assignment-created `Assignment.agent_instruction` snapshot and link back here
only for the value semantics and limit.

Profile field creation and updates follow the same RBAC and mutation safety
rules as name, visibility, and runtime binding updates. Creation stamps
`owner_principal_id` from the authorized principal and binds only a selected
runtime that is present in the authorized selectable device/runtime read model.
For a non-admin viewer this normally means a viewer-owned runtime or a runtime
made available through a public agent; an admin can use runtime rows made
visible by workspace RBAC. Local daemon detail/control follows the agent access
boundary, not a standalone device-owner-only rule. After creation, admin may mutate all agents,
owner may mutate owned agents, and no agent can be edited while it has assigned tasks.

### Agent List Timestamps

The agent client record carries required `created_at` and `updated_at`
date-times owned by the control plane. `created_at` is immutable after agent
creation. `updated_at` changes when editable agent configuration is saved. Both
fields are used by clients for agent-list dates and absolute-time tooltips.
Clients own relative/absolute formatting and tooltip presentation; they must
not synthesize or rewrite the stored timestamps.

These timestamps are distinct from runtime heartbeat, daemon liveness, and
provider session progress time. Runtime freshness remains owned by the
device/runtime read model.

### Runtime Model Dropdown

The Figma agent setting screen shows a model dropdown with provider-specific
labels. The owning contract decision is `runtime_model_catalog.v1`: model
candidates are projected as `RuntimeModelRecord` values under each
`RuntimeRecord`, not as top-level API enum values.

`model_id` is an opaque runtime-scoped identifier. A `model_id` that is valid
for one runtime is not valid for another runtime unless that runtime's catalog
also declares it. The client renders `RuntimeModelRecord.label` as display data
and must not branch on labels or hard-code Figma sample labels as enum members.

When an agent is created or updated:

- if `model_id` is omitted, the control plane stores the selected runtime's
  default model
- if `model_id` is present, the control plane validates it against the selected
  runtime catalog before saving
- if `runtime_id` changes and `model_id` is omitted, the new runtime's default
  model is selected
- if `runtime_id` changes and `model_id` is present, the model must belong to
  the new runtime

Daemon adapters may accept the selected model value only as part of an
already-authorized runtime execution request. They do not own the client-facing
model catalog or dropdown labels.

Figma can still present the model dropdown as a required control because every
selectable runtime must expose exactly one default model. That default is a
deterministic selected value, not a new API required-field rule. This avoids a
contract break for generated clients that omit `model_id` while still letting
clients render a non-empty required model control.

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
`ClientStreamEvent`. Each progress event carries `thread_id`; clients must
apply streamed progress to that thread id instead of inferring the target from a
task id alone.

Clients render task-thread progress from that typed payload and never parse
provider output text.

### Task Thread Cold Collection

Task screens read AI Agent thread history through
`GET /v1/client/ai-agent/tasks/{task_id}/threads` before opening an SSE
connection. The response returns all visible historical AI Agent task threads
for the task. It includes `active_stream` only when there is currently one
active AI Agent assignment for the task.

If an agent assignment/comment was created while the viewer was on another
screen, the control plane still persists the task thread and returns it in this
cold collection when the viewer later opens the task. The client may scroll or
focus the visible thread to match Figma, but the canonical API fact is only that
the persisted `thread_id` record is visible and that `active_stream` is
advertised only while that same thread is still active.

Completed, stopped, failed, or otherwise cold task-thread collections omit
`active_stream`; clients must not connect the SSE stream just because a task has
historical AI Agent comments. When a task was assigned, completed, and assigned
again later, the collection can contain multiple thread records, while
`active_stream`, if present, targets only the currently active `thread_id`.

When an agent delete command force-stops queued or running assignments, the
later cold collection still returns the stopped thread records. The
client-visible reason is the typed `stopped_by_agent_deleted` comment kind, not
a parsed Korean sentence. SSE may carry the same typed status while the task
screen is open; if the viewer opens the task later, the cold collection alone is
sufficient and `active_stream` stays absent unless a new active assignment
exists.

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
- `POST /v1/client/ai-agent/tasks/{task_id}/assignment`
- `DELETE /v1/client/ai-agent/tasks/{task_id}/assignment`
- `GET /v1/client/ai-agent/tasks/{task_id}/threads`
- `POST /v1/client/ai-agent/tasks/{task_id}/comments`
- `POST /v1/client/ai-agent/tasks/{task_id}/stop`
- `POST /v1/client/ai-agent/agents`
- `GET /v1/client/ai-agent/agents/{agent_id}/editability`
- `PATCH /v1/client/ai-agent/agents/{agent_id}`
- `DELETE /v1/client/ai-agent/agents/{agent_id}`
- `GET /v1/client/ai-agent/events`

`GET /v1/client/ai-agent/bootstrap` also carries the onboarding template
catalog for the Figma onboarding template-selection screen.

The event stream uses a discriminated sum type, `ClientStreamEvent`, so client
codegen can produce safe branches for runtime snapshots, agent editability, and
agent work-status updates. Runtime progress intended for the task thread is the
`agent_thread_progress` variant and carries `thread_id` plus ordered progress
lines.

## Boundary

This document and `apicontract` own shared policy vocabulary, API enum values,
sum-type shape, operation ids, paths, and BDD scenario ids.

They do not own control-plane handlers, daemon runtime probing, client UI code,
Orval output, store implementations, task cancellation workers, or provider CLI
execution.
