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

### `control-plane-ai-agent-client-api.v1`

- DSL: `apicontract/fixtures/control-plane-ai-agent-client.dsl.riido.json`
- IR: `apicontract/fixtures/control-plane-ai-agent-client.ir.riido.json`
- OpenAPI: `apicontract/fixtures/control-plane-ai-agent-client.openapi.json`

It covers the v1.22 AI Agent client surface used by Riido web and the desktop
webview:

- `GET /v1/client/ai-agent/bootstrap`
- `GET /v1/client/ai-agent/devices`
- `GET /v1/client/ai-agent/tasks/{task_id}/assignable-agents`
- `GET /v1/client/ai-agent/tasks/{task_id}/threads`
- `POST /v1/client/ai-agent/tasks/{task_id}/comments`
- `POST /v1/client/ai-agent/tasks/{task_id}/stop`
- `POST /v1/client/ai-agent/agents`
- `GET /v1/client/ai-agent/agents/{agent_id}/editability`
- `PATCH /v1/client/ai-agent/agents/{agent_id}`
- `DELETE /v1/client/ai-agent/agents/{agent_id}`
- `GET /v1/client/ai-agent/events`

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

Onboarding template selection from Figma `node-id=138-7389` is projected from
`ClientBootstrapResponse.agent_templates`. The DSL does not add a
template-selection command and does not encode `직접 설정` as an
`AgentOnboardingTemplate`. Clients render row selection, pre-selection disabled
next state, and preview skeleton/popover state from the ordered template data
and local selection state.

Onboarding direct setting from Figma `node-id=164-26969` is projected through
the existing `POST /v1/client/ai-agent/agents` command. The expanded `이름`,
`설명`, and `지침` fields map to
`CreateAgentConfigurationRequest.name`, `description`, and `instruction`.
`runtime_id`, `visibility`, optional profile image, and optional model selection
remain the normal create request fields. The DSL does not add a direct-setting
command or a fifth template record.

Task-thread history is projected as a cold collection at
`GET /v1/client/ai-agent/tasks/{task_id}/threads`. The response contains
historical thread records, including visible threads created while the viewer
was on another screen, and may include one `active_stream` HATEOAS link. The
link is omitted when the screen can render from cold history only, which keeps
SSE connection decisions deterministic for generated clients.

Runtime settings empty states from Figma `node-id=275-22731` are projected from
the same device/runtime read model as `GET /v1/client/ai-agent/devices`.
Provider install-card hover, external provider installation links, Windows app
waitlist copy, and marketing-consent presentation are not generated operations
in this fixture. They require a separate owning SSOT before the DSL adds a
waitlist or marketing mutation.

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

## Boundary

This contract owns:

- DSL schema version `riido-api-dsl.v1`
- IR schema version `riido-api-ir.v1`
- OpenAPI projection generation from the IR
- shared operation ids, paths, methods, schema names, auth scope patterns, RBAC
  policy ids, and BDD scenario ids
- top-level API enum and sum-type definitions that must survive DSL -> IR ->
  OpenAPI projection for client codegen
- deterministic fixture drift verification

This contract does not own:

- control-plane HTTP handlers
- request authorization implementation
- RBAC evaluator implementation
- persistence, stores, metrics, or SSE fan-out
- frontend code generation output
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
