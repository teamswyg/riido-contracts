# API Contract Projection

> Riido task: RIID-4718 `[Contracts] API DSL IR OpenAPI projection`

This document is the SSOT for shared API contract projection fixtures in
`riido-contracts`.

## Ownership

API contracts use this flow:

```text
Domain DSL -> canonical API IR -> OpenAPI projection
```

The Domain DSL describes the human-owned resource, policy, schema, and operation
facts. The API IR is the canonical machine-readable contract used for drift
checks. OpenAPI is generated from the IR for web clients, mock servers, docs,
and black-box HTTP checks.

OpenAPI is not the SSOT. If an OpenAPI artifact disagrees with the API IR, the
OpenAPI artifact is regenerated or rejected.

This projection is downstream of domain policy. When API schema text repeats an
AI Agent setting fact, such as `profile_thumbnail_url` or `instruction`, it is a
projection of [`ai-agent-policy.md`](ai-agent-policy.md). The dependency rule is
recorded in
[`../30-architecture/ssot-dependency-map.md`](../30-architecture/ssot-dependency-map.md).

## Current Fixture

Current fixtures:

### `control-plane-agent-catalog-api.v1`

- DSL: `apicontract/fixtures/control-plane-agent-catalog.dsl.riido.json`
- IR: `apicontract/fixtures/control-plane-agent-catalog.ir.riido.json`
- OpenAPI: `apicontract/fixtures/control-plane-agent-catalog.openapi.json`

It covers the agent catalog routes currently exposed by the control plane:

- `GET /v1/agent-catalog`
- `POST /v1/agent-catalog`
- `GET /v1/agent-catalog/{agent_id}`
- `PATCH /v1/agent-catalog/{agent_id}`
- `DELETE /v1/agent-catalog/{agent_id}`

The fixture includes the shared RBAC policy identifier
`agent_catalog_visibility.v1`: admin can read/mutate public and private agents,
owner is admin-equivalent only for owned agents, and non-admin users read owned
agents plus other users' public agents.

### `control-plane-ai-agent-client-api.v2`

- DSL: `apicontract/fixtures/control-plane-ai-agent-client.dsl.riido.json`
- IR: `apicontract/fixtures/control-plane-ai-agent-client.ir.riido.json`
- OpenAPI: `apicontract/fixtures/control-plane-ai-agent-client.openapi.json`

It covers the v1.22 AI Agent client surface used by Riido web and the desktop
webview:

- `GET /v1/client/ai-agent/bootstrap`
- `GET /v1/client/ai-agent/onboarding/fixtures`
- `POST /v1/client/ai-agent/onboarding/fixtures/{fixture_id}/agents`
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

The v2 client surface is additive, not a migration. It duplicates the AI Agent
routes under `/v2/client/workspaces/{workspace_id}/ai-agent/...` and uses
generated client paths under `riido.v2.aiAgent.*`. The selected workspace comes
from existing Riido workspace UI/API surfaces; AI Agent does not add workspace
list/create operations. v1 routes stay present so existing UI tests and mock
callers do not break, but new workspace-scoped client work must target v2.

For v2 agent creation, `workspace_id` is required as a path parameter rather
than a request-body field. The server stamps workspace ownership from the URL
path, stamps `owner_principal_id` from authorization, and rejects attempts to
derive workspace ownership from client-authored body data.

The fixture includes explicit DSL/IR enums for runtime kind, runtime
availability, runtime detection state, agent editability, agent work status,
assignment state, task comment status, client kind, and agent visibility. It
also includes the `ClientStreamEvent` sum type so client codegen receives a
discriminated event union rather than ad hoc strings. Runtime progress intended
for task threads is carried by the `agent_thread_progress` event variant with
`thread_id` and ordered line batches, not by provider raw output text.

Onboarding runtime selection from Figma `node-id=137-6746` is projected from
the same `DeviceRecord.runtimes` values: `runtime_id`,
`RuntimeRecord.availability`, and `RuntimeRecord.detection_state`. The DSL does
not add a runtime-selection command. The client renders detected/non-detected
rows from the read model, and agent create/update mutations validate the
selected `runtime_id`.

Onboarding fixture selection from Figma `node-id=138-7389` is projected from
`AgentOnboardingFixtureListResponse.fixtures`. The DSL does not expose a
template entity or template CRUD, and it does not encode `직접 설정` as an
`AgentOnboardingFixture`. Clients render row selection, pre-selection disabled
next state, and preview skeleton/popover state from the ordered fixture data and
local selection state.

Onboarding direct setting from Figma `node-id=164-26969` is projected through
the existing `POST /v1/client/ai-agent/agents` command. The expanded `이름`,
`설명`, and `지침` fields map to
`CreateAgentConfigurationRequest.name`, `description`, and `instruction`.
`runtime_id`, `visibility`, optional profile image, and optional model selection
remain the normal create request fields. The DSL does not add a direct-setting
command or a fifth fixture record. Selecting `리도`, `영실`, `홍도`, or `지원`
uses `POST /v1/client/ai-agent/onboarding/fixtures/{fixture_id}/agents` with a
complete `CreateAgentConfigurationRequest` body, so the created result is a
normal agent rather than a fixture/template entity.

Task-thread history is projected as a cold collection at
`GET /v1/client/ai-agent/tasks/{task_id}/threads`. The response contains
historical thread records, including visible threads created while the viewer
was on another screen, and may include one `active_stream` HATEOAS link. The
link is omitted when the screen can render from cold history only, which keeps
SSE connection decisions deterministic for generated clients.

Runtime settings empty states from Figma `node-id=275-22731` are projected from
the same device/runtime read model as `GET /v1/client/ai-agent/devices`.
Provider install-card hover, external provider installation links, Windows app
waitlist copy, and marketing-consent presentation are client/product facts, not
generated AI Agent operations in this fixture. `Q-CON-007` resolves that as a
no-diff boundary for `control-plane-ai-agent-client-api.v2`: the DSL/IR/OpenAPI
projection must not add waitlist or marketing-consent operations unless a future
product/marketing SSOT defines a separate generated surface.

Agent list timestamps from Figma `node-id=432-35713` are projected through
`AgentClientRecord.created_at` and `AgentClientRecord.updated_at`.
`created_at` is stamped when the agent is created and remains immutable;
`updated_at` is refreshed when editable agent configuration is saved. Clients
own shortened date formatting, absolute-time tooltip presentation, row-click /
meatball edit entry, delete-menu placement, no-description row layout, and
status-label copy/color.

Agent records also carry optional profile presentation fields:
`profile_thumbnail_url` is an HTTPS image URL string, `description` is a
client-authored one-line summary capped at 160 characters, and `instruction` is
client-authored text capped at 1000 characters by the contract projection.

For `GET /v1/client/ai-agent/tasks/{task_id}/assignable-agents`, the
projection preserves the participant-dropdown ordering contract: viewer-owned
agents first, display-name order inside each ownership group, and `agent_id` as
the stable tie-breaker when display names are equal. Pixel-level dropdown
constraints from Figma are intentionally not API fields.

Task participant assignment uses the same task-scoped API family. `POST
/v1/client/ai-agent/tasks/{task_id}/assignment` assigns one visible agent and
returns the task's AI Agent thread action response. The first agent-authored row
uses typed `comment_kind=assignment_started` unless the selected agent is busy,
in which case the existing queued tuple is returned. `DELETE
/v1/client/ai-agent/tasks/{task_id}/assignment` is the participant-removal
command and uses the existing stopped-by-user typed status; whether stopped rows
are hidden in the task comment UI remains client presentation.

## Generated Client Delivery Boundary

Generated client delivery PRs are review handoffs. This contract owns the API
projection facts that must survive into generated code: operation ids, paths,
schemas, enum values, sum-type variants, lifecycle metadata, replacement
guidance, and deprecation/removal semantics.

Generated client path searchability is also a contract projection fact. The DSL
input owns `client.module` and `client.facade_path`; the IR/OpenAPI projection
derives `client.generated_path` as `module + "." + facade_path`. Codegen must
emit that path into generated comments, and should also expose the module-local
search path such as `tasks.threadMessages.create`, so frontend developers can
find the intended generated facade route without knowing operation ids or HTTP
paths first.

`riido-control-plane` owns the downstream delivery workflow that turns the
OpenAPI projection into generated React Query files and opens a PR against
`riido-client`. That workflow may create or update a generated branch, but it
must not auto-merge the client PR. `riido-client` owns the final review,
application integration, and merge decision for generated files.

Control-plane delivery details such as tag trigger policy, branch naming,
generated path allowlist, generator pinning, and release manifest shape are
downstream execution rules. They may repeat this review-handoff rule only by
linking back to this SSOT.

## Boundary

This contract owns:

- DSL schema version `riido-api-dsl.v1`
- IR schema version `riido-api-ir.v1`
- OpenAPI projection generation from the IR
- shared operation ids, paths, methods, schema names, auth scope patterns, RBAC
  policy ids, and BDD scenario ids
- top-level API enum and sum-type definitions that must survive DSL -> IR ->
  OpenAPI projection for client codegen
- generated-client lifecycle and review-handoff semantics
- derived generated client path metadata used for searchable generated comments
- deterministic fixture drift verification

This contract does not own:

- control-plane HTTP handlers
- request authorization implementation
- RBAC evaluator implementation
- persistence, stores, metrics, or SSE fan-out
- frontend code generation output
- generated client branch creation, PR creation mechanics, or client PR merge
  decisions
- production bearer tokens, IdP configuration, Terraform, or deployment
  evidence

## Validation

The deterministic gate is:

```bash
go run ./tools/apicontract verify
go test ./apicontract -run 'AgentCatalog|AIAgentClient' -count=1
```

`verify` regenerates IR and OpenAPI in memory from the DSL and rejects any drift
from the checked-in generated artifacts.
