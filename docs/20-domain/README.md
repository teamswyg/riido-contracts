# Domain Contract Docs

This directory carries the public SSOT documents for the contracts currently
implemented in this module.

The split-repo context map is [`context-map.md`](context-map.md). Promotion,
package decomposition, and downstream integration gates are owned by
[`../30-architecture/contract-promotion-policy.md`](../30-architecture/contract-promotion-policy.md),
[`../30-architecture/module-decomposition.md`](../30-architecture/module-decomposition.md),
and
[`../30-architecture/integration-matrix.md`](../30-architecture/integration-matrix.md).

RIID-4641 migrates:

- C1 Task Lifecycle contract into `task`
- C2 IR Event Log contract into `ir`
- C3 Provider Capability contract into `provider/capability`

RIID-4670 migrates:

- C11/C10 distribution channel and provider routing status vocabulary into
  `hostintegration`

RIID-4687 migrates:

- C10 assignment polling DTOs, service schema identifiers, assignment state
  transition predicates, poll actions, task event type values, and agent runtime
  binding DTOs into `assignment`

RIID-4718 adds:

- C10 API contract projection fixtures into `apicontract`, where Domain DSL and
  API IR are SSOT and OpenAPI is generated

RIID-4720 adds:

- C10 AI Agent policy vocabulary in [`ai-agent-policy.md`](ai-agent-policy.md)
- generated AI Agent onboarding fixture evidence in
  [`ai-agent-onboarding.md`](ai-agent-onboarding.md)
- generated AI Agent configuration and mutation-safety evidence in
  [`ai-agent-configuration.md`](ai-agent-configuration.md)
- `control-plane-ai-agent-client-api.v2` API projection fixtures, including
  top-level enum and sum-type definitions for client codegen safety

RIID-4868 adds:

- device principal enrollment and daemon credential semantics in
  [`device-principal.md`](device-principal.md)

Current AI Agent progress-message work adds:

- fixed, translated, append-only runtime progress copy in
  [`progress-message-catalog.md`](progress-message-catalog.md), with the
  executable catalog under `progressmessage/catalog.dsl.riido.json`

The EventIngestor implementation, policy redaction catalog, provider adapters,
server stores, and Terraform deployment code remain outside this repository.
They may consume these contracts, but they do not become contract module
implementation details.

Unresolved promotion/versioning questions are tracked in
[`../50-roadmap/open-questions.md`](../50-roadmap/open-questions.md).
