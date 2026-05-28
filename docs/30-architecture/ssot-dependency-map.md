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
| `instruction` is optional client-authored agent guidance text capped at 1000 characters | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) plus the AI Agent API DSL fixture | Control-plane validates/stores/projects it. Daemon may consume the assigned value for prompt/native-config materialization. Infra acts only if a future storage/secret/media requirement appears. |
| Agent editability requires zero assigned tasks | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) and the API DSL BDD scenarios | Control-plane implements the executable HTTP/store behavior and emits client events. |
| Admin/owner/public-private visibility vocabulary | [`../20-domain/ai-agent-policy.md`](../20-domain/ai-agent-policy.md) and API fixture policy ids | Control-plane owns the executable RBAC evaluator and request authorization boundary. |
| DSL -> IR -> OpenAPI projection rules | [`../20-domain/api-contract-projection.md`](../20-domain/api-contract-projection.md) | Control-plane mirrors generated fixtures and owns local generator drift checks. |
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
  thumbnails and normal instruction text. They must not redefine API shape or
  daemon execution.
