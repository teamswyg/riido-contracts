# SSOT Dependency Map

> Riido task: RIID-4753 `[Contracts/Server/Daemon/Infra] Agent setting SSOT dependency direction`

This document is the meta-SSOT for how AI Agent configuration facts move
between Riido SSOT documents and repositories. It does not replace the domain
SSOTs. It defines the dependency direction between them so duplicated wording is
either a linked projection or a defect.

## Rule

One durable fact still has one owner. A downstream SSOT may repeat that fact
only when it names the upstream owner and limits itself to local execution,
projection, or harness behavior.

If a downstream test, implementation, cost constraint, or deployment constraint
proves that the upstream model is wrong or incomplete, the next change flows
bottom-up: document the local finding, escalate the owning upstream SSOT, then
regenerate or update downstream projections.

## Agent Configuration Ownership

| Fact | Owning SSOT | Downstream projection |
| --- | --- | --- |
| Agent means a task-assignable abstraction of a configured runtime | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) | Control-plane handlers, daemon runtime prompts, clients, and infra docs link to this language. |
| `profile_thumbnail_url` is an optional HTTPS image URL string | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture | Control-plane validates/stores/projects it. Clients render it. Daemon ignores it. Infra acts only if a future media/storage SSOT replaces URL-only storage. |
| `description` is optional client-authored one-line agent summary text capped at 160 characters | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture | Control-plane validates/stores/projects it. Clients render or truncate it. Daemon ignores it. Infra acts only if a future search, media, durability, or moderation SSOT changes storage requirements. |
| `instruction` is optional client-authored agent guidance text capped at 1000 characters | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture | Control-plane validates/stores/projects it. Daemon may consume the assigned value for prompt/native-config materialization. Infra acts only if a future storage/secret/media requirement appears. |
| Agent editability requires zero assigned tasks | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) and the API DSL BDD scenarios | Control-plane implements the executable HTTP/store behavior and emits client events. |
| Admin/owner/public-private visibility vocabulary | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) and API fixture policy ids | Control-plane owns the executable RBAC evaluator and request authorization boundary. |
| DSL -> IR -> OpenAPI projection rules | [`../20-domain/api-contract-projection.md`](../20-domain/api-contract-projection.md) | Control-plane mirrors generated fixtures and owns local generator drift checks. |
| Participant dropdown AI Agent visibility and ordering | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma annotation evidence at `node-id=153-12742` | Control-plane implements `GET /v1/client/ai-agent/tasks/{task_id}/assignable-agents`. Client owns member sorting, long-name rendering, max height, scrollbar width, checkbox layout, and mixed member/agent visual composition. Daemon and infra do not change for dropdown presentation alone. |
| AI Agent assignment target scope | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma planning evidence at `node-id=153-15935` | Control-plane implements task/subtask-scoped generated operations under `/v1/client/ai-agent/tasks/{task_id}/...` and must not add project, milestone, intake, AI property filler, or mention operations without a new owning SSOT. Client owns whether non-target surfaces hide, disable, or omit agent UI. Daemon consumes only SaaS assignments after target validation. Infra is no-diff unless a future target surface adds durability, stream, queue, secret, or deployment topology. |
| AI Agent, runtime, and agent-management menu route affordance | Figma `node-id=156-19307` as cited by [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) | Client owns concrete menu rendering, ordering, selected state, and route wiring. Control-plane serves data after the route is opened. Daemon and infra do not change for menu placement alone. |
| Task-thread cold collection, comment submit, progress stream, stop operation, queued state, and stopped-by-deleted-agent semantics | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma annotation evidence at `node-id=153-15931`, normal task-screen evidence at `node-id=236-21379`, busy-agent evidence at `node-id=153-8761`, and stopped-by-deleted-agent evidence at `node-id=227-19354` | Control-plane implements `GET /v1/client/ai-agent/tasks/{task_id}/threads`, `POST /v1/client/ai-agent/tasks/{task_id}/comments`, `POST /v1/client/ai-agent/tasks/{task_id}/stop`, `DELETE /v1/client/ai-agent/agents/{agent_id}`, HTTP/SSE, and generated-client call chains. Daemon produces parsed progress batches with thread identity and consumes cancellation/stop commands from SaaS, not UI clicks or rendered copy. Client owns scroll, hover, modal, animation, reply-input rendering, send-button state, task sidebar fields, Riido actor label, Korean copy, hidden action state, timestamp wording, avatar rendering, and rendered thread composition. Infra acts only if a future topology/cost/evidence SSOT changes deployment requirements. |
| Runtime settings device/runtime read model and local daemon lifecycle boundary | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma annotation evidence at `node-id=162-23090` | Control-plane implements `GET /v1/client/ai-agent/devices` and `device_runtime_snapshot` projection. Daemon owns local status/start/stop lifecycle behavior for the current desktop device. Client owns the agent hover popover, daemon stop modal, restart animation, and any local helper composition. Infra acts only if a future topology/storage/evidence SSOT changes deployment requirements. |
| Runtime settings empty states, provider install-card hover, and Windows waitlist variants | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma evidence at `node-id=275-22731` and web-onboarding waitlist evidence at `node-id=236-29749` | Control-plane still exposes only device/runtime read-model data through `GET /v1/client/ai-agent/devices`. Client/product owns install-card links, hover states, Windows app waitlist copy, and marketing-consent presentation. A generated waitlist/marketing mutation needs a separate owning SSOT before contracts/control-plane add an operation. Daemon and infra do not change for this presentation-only slice. |
| Runtime-scoped model catalog and agent `model_id` selection | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, resolving `Q-CON-006` from Figma `node-id=164-50215` and add-screen `node-id=134-6542` | Control-plane projects `RuntimeRecord.models`, validates create/update `model_id` against the selected runtime, defaults omitted values, and projects saved `model_label`. Client owns dropdown rendering. Daemon consumes only the already-authorized selected model for runtime execution. Infra acts only if a future deployment-backed model catalog/store/cache/index/evidence SSOT changes topology. |
| Agent setting add/list/edit fields and update timestamp | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma annotation evidence at `node-id=164-50215` and add-screen evidence at `node-id=134-6542` | Control-plane implements `POST /v1/client/ai-agent/agents`, `updated_at` projection, and mutation-time refresh. Client owns save-button enablement, row/meatball edit entry, long-description presentation, timestamp formatting/tooltip, and dropdown rendering. Daemon consumes only assigned runtime/model values after upstream policy supplies them. Infra acts only if a future model catalog, storage, search, media, or evidence SSOT changes topology. |
| Onboarding runtime selection, template catalog, and no-runtime branch | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture, with Figma onboarding evidence at `node-id=42-3014`, `node-id=137-6746`, `node-id=138-7389`, `node-id=164-30192`, and `node-id=164-30206` | Control-plane projects `agent_templates` through `GET /v1/client/ai-agent/bootstrap`, device/runtime liveness through `GET /v1/client/ai-agent/devices`, and validates the selected `runtime_id` through create/update mutations. Client owns runtime radio rendering, `감지됨` / `감지 안 됨` labels, row dimming, template row selection, direct-setting expansion, scroll, two-line ellipsis, workspace selection/list scrolling/create-new entry points, all-disconnected provider-list rendering, the `시작하기` CTA, and the decision to skip template selection when device/runtime data has no selectable runtime. Daemon supplies runtime liveness/detection facts; infra acts only if template media/storage/search/topology changes appear. |
| Web onboarding auth/team/distribution presentation | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md), with Figma evidence at `node-id=236-29749` | Sign-up/login, Google-auth terms consent, email/password validation, terms row click behavior, and member invite/link-copy belong to auth/team/client product surfaces. macOS app download and Windows waitlist CTAs are distribution/product facts. Chat/progress animations are presentation facts. None of these become AI Agent API, daemon, or Terraform facts without a separate owning SSOT. |
| Runtime prompt/native-config consumption | `riido-daemon` C4/C6 SSOTs | Contracts own only the instruction value semantics, not process execution or provider files. |
| Production deployment/storage topology | `riido-infra` Terraform/architecture SSOTs | Contracts do not create buckets, secrets, CDNs, or persistent stores. |

## Top-Down Loop

Top-down work starts when product, policy, or design changes the mental model.

```text
product/design evidence
  -> contracts canonical policy and DSL
  -> API IR and OpenAPI projection
  -> control-plane executable HTTP/SSE behavior
  -> generated client handoff
  -> daemon runtime consumption when assignments execute
  -> infra deployment work only when topology, secrets, media, or durability change
```

The harness at each level catches drift:

- contracts: DSL/IR/OpenAPI fixture verification
- control-plane: HTTP black-box tests, generator drift tests, smoke tests
- daemon: provider/runtime/workdir tests and real-runtime integration checks
- infra: Terraform authoring, plan/evidence, contract, and Mermaid projection
  checks

## Bottom-Up Loop

Bottom-up work starts when a lower layer discovers a contradiction.

```text
implementation / harness / operations finding
  -> local repo SSOT records the observed constraint
  -> if business meaning changes, escalate to contracts
  -> regenerate projections and update downstream harnesses
  -> keep deprecated/superseded API surfaces until clients can migrate
```

Examples:

- A frontend usability issue may enter through control-plane API design, but a
  vocabulary or policy change must move up to `ai-agent-policy.md`.
- A daemon prompt-placement limitation may enter through daemon C4/C6 docs, but
  changing the meaning or limit of `instruction` must move up to contracts.
- A client rendering issue may enter through control-plane/client handoff, but
  changing the meaning or limit of `description` must move up to contracts.
- A storage or media moderation requirement may enter through infra, but
  replacing URL-only thumbnail storage must first create or update a media
  contract instead of silently changing the API fixture.

## Duplicate Audit

The current duplicated wording is intentional only in these forms:

- `ai-agent-policy.md` owns canonical agent setting meaning. Other repos may
  restate the fields as local behavior only.
- `api-contract-projection.md` owns DSL/IR/OpenAPI mechanics. Control-plane may
  restate the generated fixture flow because it runs the local mock/generator
  harness.
- Control-plane RBAC and editability docs may repeat visible behavior because
  they own the executable evaluator. They must not redefine the shared
  vocabulary without a contracts change.
- Daemon docs may describe how instruction text enters runtime prompts or
  native config. They must not redefine storage, length, RBAC, or thumbnail
  policy.
- Infra docs may explain that no Terraform diff is required for URL-only
  thumbnails, one-line descriptions, and normal instruction text. They must not
  redefine API shape or daemon execution.
- Menu-placement docs may restate that Figma shows AI/runtime/agent-management
  route affordances. They must not turn client menu rendering into a new API,
  daemon runtime, or Terraform requirement without a separate owning SSOT.
- Task-thread annotation docs may restate that Figma cites
  `riido.aiAgent.events.stream` and `riido.aiAgent.tasks.stop`. They may also
  restate that task screens read the cold thread collection before following an
  advertised active stream, and that assignment-created-while-viewer-away
  records remain visible through that cold collection. They may restate that
  Figma `node-id=153-8761` renders a busy-agent queued row, but only the typed
  `queued_by_busy_agent`/`queued` status tuple is canonical here. They may also
  restate that Figma `node-id=227-19354` renders a stopped row after agent
  deletion, but only the typed `stopped_by_agent_deleted`/`stopped` status tuple
  and reuse of the existing delete command are canonical here. They must not
  turn generated client chain names, Korean display copy, Riido actor labels,
  timestamp wording, scroll/focus behavior, hover states, modals, hidden action
  state, row layout, avatar rendering, or animation references into canonical
  API or daemon facts.
- Participant-dropdown docs may restate that Figma shows member sorting,
  owned-first agent sorting, long-name states, and scroll/height constraints.
  They must not turn client presentation facts into API, daemon, or Terraform
  requirements. Only agent visibility and owned-first agent ordering are
  canonical AI Agent contract facts.
- Assignment-target-scope docs may restate that Figma `node-id=153-15935` says
  only tasks and subtasks can receive Agent assignment, existing AI property
  filling does not recommend agents, and agent mentions are unsupported. They
  must not create placeholder generated helpers for projects, milestones,
  intakes, property filling, or mentions. Only task/subtask-scoped generated
  operations are canonical today.
- Runtime-settings docs may restate that Figma shows runtime liveness, attached
  agents, an agent hover popover, a daemon stop modal, restart animation, empty
  runtime install-card states, and Windows waitlist variants. They must not turn
  hover/modal/animation presentation, external provider install links,
  marketing-consent UI, or desktop-local command composition into API, daemon,
  or Terraform requirements without a separate owning SSOT. Only the
  device/runtime read model and local daemon lifecycle boundary are canonical
  cross-repo facts.
- Agent-setting docs may restate that Figma shows add/list/edit screens,
  profile photo, row/meatball edit entry, long-description UI, absolute-time
  tooltips, runtime dropdowns, save-button enablement, and model dropdowns. They
  must not hard-code model candidates as contract enum values. Model candidates
  are canonical only as runtime-scoped `RuntimeModelRecord` catalog data under
  the device/runtime read model. Profile/configuration fields, create/update
  operations, editability, runtime binding, visibility, `model_id`, and
  `updated_at` are canonical contract facts today.
- Onboarding docs may restate that Figma shows runtime choice, template choice,
  direct setting, scroll, description ellipsis, and no-installed-AI branches.
  They must not turn client presentation into new API commands. The template
  catalog is canonical API data; no-runtime branching is derived from existing
  runtime/device liveness.
- Web-onboarding docs may restate that Figma shows sign-up, terms consent,
  member invite, app download, Windows waitlist, and animation references. They
  must not turn auth/team/product/distribution presentation into AI Agent API,
  daemon, or Terraform facts without a new owning SSOT.
