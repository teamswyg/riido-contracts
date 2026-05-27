# Domain Contract Docs

This directory carries the public SSOT documents for the contracts currently
implemented in this module.

RIID-4641 migrates:

- C1 Task Lifecycle contract into `task`
- C2 IR Event Log contract into `ir`
- C3 Provider Capability contract into `provider/capability`

The EventIngestor implementation, policy redaction catalog, provider adapters,
server stores, and Terraform deployment code remain outside this repository.
They may consume these contracts, but they do not become contract module
implementation details.
