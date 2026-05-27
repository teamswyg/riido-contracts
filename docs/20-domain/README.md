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

The EventIngestor implementation, policy redaction catalog, provider adapters,
server stores, and Terraform deployment code remain outside this repository.
They may consume these contracts, but they do not become contract module
implementation details.

Unresolved promotion/versioning questions are tracked in
[`../50-roadmap/open-questions.md`](../50-roadmap/open-questions.md).
