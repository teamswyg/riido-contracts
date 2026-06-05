# Device Principal

> Riido task: RIID-4868 `[Contracts] device principal 소유권 SSOT 정의`

This document is the SSOT for how Riido binds a desktop-launched daemon to the
same user identity that the web or desktop webview already authenticated.

## Ownership

This contract owns:

- the distinction between `UserPrincipal` and `DevicePrincipal`
- device enrollment semantics
- `device_id` and `device_secret` credential meaning
- device/runtime ownership direction
- dynamic `AgentRuntimeBinding` derivation direction
- daemon credential transport vocabulary
- secret storage and non-exposure rules that downstream repos must preserve

This contract does not own:

- existing Riido workspace list/create APIs
- existing Riido JWT verification implementation
- control-plane persistence adapters
- Electron secure-storage implementation details
- daemon process execution, provider CLI execution, or local runtime detection
- frontend screen composition

## Principal Model

`UserPrincipal` is the authenticated Riido user identity resolved from the
existing Riido login token. Browser and desktop-webview AI Agent client
requests may carry that token through `X-Riido-AI-Agent-Token`, and the
control plane may ask the existing Riido API server to validate it.

`DevicePrincipal` is a daemon-facing identity resolved from a registered
`device_id` plus a `device_secret`. A daemon must not be authenticated by a
browser user JWT. The control plane may remember which user enrolled the
device, but the daemon proves only possession of the device credential.

## Ownership Graph

The durable model is:

```text
UserPrincipal
  -> Device
      -> Runtime

Workspace
  -> Agent
      -> selected Runtime binding
```

A device is account-owned through `owner_principal_id`. A runtime is owned
through its device owner. An agent is workspace-owned and stores the creator as
`owner_principal_id` inside that workspace. `workspace_id` scopes agent
creation, listing, assignment, and RBAC; it is not the canonical owner of a
device.

When device enrollment is requested from a desktop app, the request may carry a
`workspace_id` only so the existing user-token authorizer can evaluate the
user's current workspace membership and return a stable user principal. That
workspace context does not make the enrolled device workspace-owned.

`team_id`, Open API workspace keys, and existing workspace Open API credentials
are not part of this principal model. They must not be used to identify a
daemon, match a browser session to a daemon, select an agent runtime binding, or
authorize assignment polling. Those values may exist only in legacy
task-context reader integrations outside this DevicePrincipal contract.
This is a hard exclusion, not a fallback path: a valid DevicePrincipal test or
staging smoke check must not require `team_id`, `teamId`, an OpenAPI task
context URL, an Open API key, or `X-Workspace-Api-Key` to connect browser-visible
AI Agent state with daemon-visible runtime state.

## Enrollment

The desktop app owns the enrollment trigger. After the user is logged in inside
the desktop webview, Electron main process may call the control plane to enroll
or reuse the local machine.

The enrollment command is authenticated as `UserPrincipal`. A successful
response returns:

- `device_id`: durable device identifier safe to store and display
- `device_secret`: one-time secret used by the daemon to authenticate as
  `DevicePrincipal`

The control plane stores only a hash of `device_secret`. The raw secret is
returned once and must not be re-readable through any client or daemon API. If
the desktop app loses the secret, it must enroll again or use a future rotation
flow.

## Daemon Credential Transport

Daemon-originated control-plane calls use explicit device credential headers:

- `X-Riido-Device-ID`
- `X-Riido-Device-Secret`

These headers are intentionally separate from `X-Riido-AI-Agent-Token` and
from `Authorization: Bearer ...`. Frontend/client calls use the AI Agent token
header. Daemon calls use device credential headers. A generated frontend client
must not accidentally receive or forward `device_secret`.

Compatibility static bearer tokens may remain for local tests, review accounts,
or migration slices, but production daemon identity is `DevicePrincipal`.

## Dynamic Agent Runtime Binding

Desktop-launched daemon assignment must not depend on a static
`RIIDO_AI_SERVER_AGENT_BINDINGS_JSON` deployment secret. The binding a daemon can
see is derived by the control plane from three durable facts:

- the `DevicePrincipal` authenticated by `device_id` and `device_secret`
- the latest daemon runtime snapshot for that device
- workspace-scoped agents whose saved `runtime_id` points at one of those
  runtime records

The derived DTO is `assignment.AgentRuntimeBinding`. `riido-contracts` owns the
DTO shape through the assignment package, but it does not own the query store or
HTTP implementation. `riido-control-plane` owns the projection endpoint that
returns the current bindings to a daemon, and `riido-daemon` owns polling that
projection before it polls an agent-specific assignment queue.

The dependency direction is:

```text
Desktop UserPrincipal login
  -> device enrollment
  -> DevicePrincipal credential
  -> daemon runtime snapshot
  -> control-plane AgentRuntimeBinding projection
  -> daemon assignment poll
```

If a runtime is removed from the device, the next runtime snapshot must make the
runtime unavailable for newly derived bindings. Existing agent records are not
deleted by runtime disappearance; they become actionable again only when the
same device/runtime identity is detected and reported again.

## Device And Runtime Liveness

The daemon refreshes device/runtime liveness by sending a daemon runtime
snapshot every 5 seconds while the desktop-launched daemon is running. The
snapshot is daemon-level and should aggregate the current device runtime rows
where possible; it must not create one SaaS request per runtime unless a
provider-specific transport forces that shape.

SSOT phrase for dependency checks: daemon runtime snapshot every 5 seconds.

The control plane treats a device/runtime read model as stale when the latest
runtime snapshot for that device has not been refreshed for 20 seconds. Stale
devices and runtimes are still historical records, but client-facing runtime
reads must project them as unavailable:

- device daemon availability becomes `offline`
- runtime availability becomes `offline`
- runtime detection state becomes `missing`
- stale runtimes must not be used for newly derived `AgentRuntimeBinding`
  projection

This policy does not add a generated frontend route or change the response
shape. Frontend code keeps reading the existing device/runtime fields and
renders the current availability values. A later fresh runtime snapshot may
make the same device/runtime identity available again.

## Secret Non-Exposure

The raw `device_secret` must not be written to:

- browser local storage, cookies, or webview JavaScript state
- renderer IPC payloads
- command-line arguments
- daemon status output
- logs, task evidence, SSE events, or generated client responses

Desktop may pass the secret to daemon through an environment variable or a
protected local config file owned by Electron main process. The downstream
implementation chooses the OS-specific storage mechanism, but it must preserve
the non-exposure rule above.

## Revocation And Rotation

This slice defines issuance and verification. Revocation and rotation are
future commands unless a downstream implementation needs them to ship safely.
If they are added, they must preserve:

- `device_id` remains stable while a secret rotates
- a revoked secret cannot authenticate future daemon requests
- revocation does not delete historical device/runtime audit facts
- workspace-scoped agent RBAC still controls who can act through an agent

## Downstream Use

`riido-control-plane` implements enrollment, hash storage, request
verification, and black-box tests.

`riido-daemon` consumes `device_id` and `device_secret` as SaaS credentials and
uses them for poll, heartbeat, event, and progress requests.

`riido-desktop` owns registering the local device after login, storing the raw
secret safely, and launching the daemon with the credential. `riido-client`
does not need to change for this contract; it keeps consuming protected SaaS
read models.
