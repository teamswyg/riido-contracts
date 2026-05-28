# Contracts Integration Matrix

> Riido task: RIID-4713 `[Contracts] Architecture SSOT docs migration`

This repo has no external runtime dependencies. Integration means downstream
compatibility, not live infrastructure.

## Public Gates

| Surface | Verification | External dependency |
| --- | --- | --- |
| full contract module | `go test ./...` | none |
| dependency boundary | `go list -m all` returns only this module | none |
| task/IR coupling | task transition tests reference real IR event constants | none |
| provider capability | fingerprint and protocol-args tests | none |
| host integration | distribution/routing status vocabulary tests | none |
| assignment polling | contract JSON and generated Go constants alignment | none |
| API contract projection | DSL/IR/OpenAPI drift verification plus agent-catalog and AI Agent client projection tests | none |
| SSOT dependency direction | `tools/ssotdeps` verifies the machine-readable dependency manifest, source phrases, and acyclic repo dependency graph | none |
| architecture docs | required docs, package list coverage, stale runtime wording scan | none |

## Downstream Gates

| Consumer | Expected gate |
| --- | --- |
| `riido-daemon` | imports tagged task/IR/provider/hostintegration/assignment packages and runs daemon compatibility tests |
| `riido-control-plane` | imports tagged assignment/provider/hostintegration packages and runs control-plane compatibility tests |
| web frontend | consumes generated OpenAPI projection and runs client/API compatibility tests, including AI Agent enum/sum-type codegen |
| desktop webview | consumes the same generated AI Agent client OpenAPI projection as web |
| `riido-infra` | consumes only tagged public contracts or public facades; no private runtime imports |

Downstream gates must run in their owning repositories. This repo only proves the
contract package is internally consistent and tag-ready.

## Local Commands

```bash
go test ./...
go list -m all
go test ./assignment -run 'AssignmentContract|AssignmentTransition|AssignmentAPI' -count=1
go test ./apicontract -run 'AgentCatalog|AIAgentClient' -count=1
go run ./tools/apicontract verify
go run ./tools/ssotdeps verify
```
