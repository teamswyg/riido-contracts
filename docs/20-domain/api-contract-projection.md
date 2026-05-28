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
- `POST /v1/client/ai-agent/tasks/{task_id}/comments`
- `POST /v1/client/ai-agent/tasks/{task_id}/stop`
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
ordered line batches, not by provider raw output text.

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
