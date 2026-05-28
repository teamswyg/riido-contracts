# Contracts Module Decomposition

> Riido task: RIID-4713 `[Contracts] Architecture SSOT docs migration`

`github.com/teamswyg/riido-contracts` is a public Go module for shared Riido
contracts. It is intentionally small and standard-library only.

## Packages

| Package/path | Role | Must not own |
| --- | --- | --- |
| root `contracts` | module identity constants | runtime behavior, tag orchestration |
| `task` | C1 task lifecycle states and transition matrix | scheduler, local task DB, validation execution |
| `ir` | C2 canonical event envelope, event catalog, reducer contract | event storage adapters, ingestion pipelines |
| `provider/capability` | C3 provider capability and fingerprint vocabulary | provider detection, process launch, real CLI probing |
| `hostintegration` | C11/C10 distribution and provider routing status vocabulary | app data roots, IPC, consent storage, store helper implementation |
| `assignment` | C10 assignment polling DTOs and state vocabulary | control-plane store actor, HTTP/SSE handlers, authorization, persistence |
| `apicontract` | C10 API DSL, API IR, generated OpenAPI projection fixtures, API enum/sum-type vocabulary | control-plane handlers, frontend implementation, authorization/RBAC implementation, generated client code |
| `tools/apicontract` | deterministic fixture drift verifier/generator | runtime code generation output, network calls, third-party parser dependencies |

## Dependency Rules

Allowed:

- Go standard library
- package-local tests
- `task -> ir`

Forbidden:

- third-party dependencies
- imports from `riido-daemon`, `riido-control-plane`, or `riido-infra`
- Terraform, AWS, Docker, provider CLI, or credential packages
- runtime adapter packages

`go list -m all` must return only this module.

## Contract Shape

A contract package may contain:

- schema-version constants
- enum/string vocabulary
- DTO structs with JSON tags
- pure validation helpers
- pure transition/classification helpers
- generated projection fixtures
- codegen-safe enum and sum-type contract shapes
- black-box fixture tests

A contract package must not contain:

- network listeners
- process execution
- file-system persistence adapters
- secret lookup
- cloud API calls
- runtime-specific policy decisions

## Tag Boundary

The Go module tag is the distribution boundary. Downstream repositories consume
contracts by tag, then replace duplicated constants or fixtures in their own
repos. A contract is not considered promoted until at least one downstream repo
imports the tagged module and passes its own compatibility gate.
