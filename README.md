# riido-contracts

Shared Riido contracts, schemas, and cross-repository Go module.

This repository is the public contract boundary between `riido-daemon`,
`riido-control-plane`, and private deployment infrastructure. It should contain
stable protocol/contract artifacts only when more than one repository needs the
same fact.

## Module

```text
github.com/teamswyg/riido-contracts
```

## Repository Boundary

This repository may contain:

- shared event and API contract versions
- JSON schema or generated fixtures
- cross-repository black-box contract test fixtures
- small Go packages that are intentionally shared

This repository must not contain:

- daemon implementation details
- control-plane implementation details
- Terraform state, AWS account details, or deployment secrets
- provider CLI binaries

## Current Packages

- `assignment`: C10 SaaS assignment polling DTOs, schema identifiers, state
  transition predicates, poll action values, task event type values, and agent
  runtime binding DTOs shared by daemon and control-plane repositories.
- `ir`: C2 IR event log contract, event catalog, envelope validation, and pure
  reducer contract.
- `hostintegration`: C11/C10 distribution channel and provider routing status
  vocabulary shared by daemon metadata and control-plane provider status
  contracts.
- `provider/capability`: C3 provider capability model, protocol identifiers,
  compatibility status, protocol-critical args, and capability fingerprinting.
- `task`: C1 task lifecycle states and transition matrix. This package depends
  on `ir`; `ir` must not depend on `task`.

## Verification

```bash
go test ./...
go list -m all
```

## License

Apache-2.0.
