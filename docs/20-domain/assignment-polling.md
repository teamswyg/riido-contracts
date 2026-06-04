# Assignment Polling Contract

> Riido task: RIID-4687 `[Contracts] assignment polling DTO contract migration`

This document is the public SSOT for the C10 assignment polling contract that
must be shared by `riido-daemon` and `riido-control-plane`.

## Responsibility

`riido-contracts/assignment` owns only the shared contract surface:

- `riido-ai-server.v1` service schema version
- `riido-ai-server-contract.v1` executable contract fixture
- assignment state values and terminal/agent-active classification
- legal assignment state transitions
- daemon poll action values
- task event type values
- assignment, poll, heartbeat, agent-event, task-event, and agent runtime
  binding DTOs

The package does not own control-plane store actors, HTTP handlers, SSE fan-out,
metrics routes, health routes, request authorization, provider status stores,
daemon provider process execution, daemon local API behavior, Terraform,
DynamoDB/EventBridge adapters, AWS account settings, secret values, or
deployment evidence.

## Executable Contract

The executable fixture is
[`assignment/assignment_contract.riido.json`](../../assignment/assignment_contract.riido.json).

That fixture owns the canonical lists for:

- assignment states
- terminal and agent-active state flags
- legal transitions from each state
- poll actions
- task event type values
- assignment payload fields that the control plane snapshots for daemon runtime
  composition

Markdown must link to the fixture instead of redefining the matrix. The contract
test decodes the fixture as one strict JSON document: unknown fields and
trailing JSON documents are failures. A new executable contract fact must update
the test schema and this human explanation in the same work unit.

## DTO Surface

The public shared DTOs are:

- `AssignRequest`
- `Assignment`
- `PollRequest`
- `PollResponse`
- `AgentHeartbeatRequest`
- `AgentHeartbeatResponse`
- `AgentEventRequest`
- `AgentEventResponse`
- `TaskEvent`
- `AgentRuntimeBinding`

`Assignment.agent_instruction` is the assignment-created snapshot of the
configured agent's `instruction` value. It is optional and keeps the same
1000-character limit owned by [`ai-agent-policy.md`](ai-agent-policy.md). The
control plane copies it into `AssignRequest` / `Assignment` at assignment
creation time; later agent edits do not rewrite an already-created assignment.
The daemon consumes this snapshot only when composing provider-specific runtime
instructions.

`Assignment.allow_experimental_runtime` is the assignment-created snapshot that
lets the daemon execute a runtime capability marked
`requires_experimental_opt_in=true`. The control plane derives the value from
the selected agent/runtime binding at assignment creation time; the daemon must
not infer it from provider name or local environment. This field is false by
default and omitted from JSON unless the assignment explicitly opts in.

`Assignment.model_id` is the assignment-created snapshot of the selected
agent's validated runtime-scoped model. The control plane resolves omitted
agent create/update values to the selected runtime default before assignment,
then copies the saved `agent.model_id` into `AssignRequest` / `Assignment`.
Daemon consumes this value only as provider runtime model selection input; it
must not re-resolve model defaults from provider name, environment, team id, or
Open API key configuration.

`PollResponse.action=start` means the control plane leased a queued assignment
to the polling daemon/runtime. `PollResponse.action=active` means the same
daemon/runtime identity is still considered to hold an active assignment. A
daemon that already has the task in local in-flight state may ignore `active`;
a daemon that lost local in-flight state after restart may rebuild the same
`TaskRequest` from the returned assignment snapshot. The control plane must
therefore keep `active` payloads self-contained in the same way as `start`.

While an assignment is `leased`, `ready`, or `running`, the daemon sends an
assignment heartbeat every 5 seconds. The control plane treats the active
assignment lease as stale when it has not been refreshed for 20 seconds. A stale
active assignment is terminally failed before the same agent can claim later
queued work, and later daemon heartbeats/events for that expired lease must not
revive it. The heartbeat response refreshes only assignments that still hold the
active lease; a daemon must treat a missing refreshed assignment as a server-side
stale/cancel signal and stop the local provider run.

`AgentRuntimeBinding` is the shared DTO that lets a daemon know which
workspace-created agent may poll through one of its runtime slots. The DTO shape
is shared here, but the binding list is not a static deployment secret in the
current DevicePrincipal path. Control-plane derives it from persisted agent
configuration plus the latest daemon runtime snapshot for the authenticated
device.

`Health` and `MetricsSnapshot` remain in `riido-control-plane` for now because
they are control-plane adapter/read-model contracts rather than daemon polling
contracts.

## Handoff

After this contract is tagged, `riido-control-plane` should replace duplicated
assignment constants and DTO declarations with aliases or imports from
`github.com/teamswyg/riido-contracts/assignment`.

After the control-plane consumer is on the tagged contract, `riido-daemon` can
migrate its control-plane SaaS adapter without importing private
`internal/riidoaiserver` packages.
