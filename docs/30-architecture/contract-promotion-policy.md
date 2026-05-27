# Contract Promotion Policy

> Riido task: RIID-4713 `[Contracts] Architecture SSOT docs migration`

This policy decides when a fact belongs in `riido-contracts`.

## Promotion Rule

Promote a fact only when all conditions are true:

1. At least two repositories must agree on the same fact at build time or
   black-box test time.
2. The fact can be represented without importing runtime implementation
   packages.
3. The fact can be versioned and tested independently.
4. The owning domain SSOT and `docs/migration/contracts.md` are updated in the
   same PR.
5. At least one downstream repository has a planned or completed import gate.

If only one runtime owns the fact, keep it local to that runtime repository.

## Versioning

The module uses Git tags. Package schema constants are independent from the Go
module tag:

- `task.FSMSchemaVersion` changes when the task transition matrix makes a
  version-affecting change.
- IR event schema versions change per event type.
- `assignment.SchemaVersion` changes when the shared assignment polling API
  changes.
- provider capability and host-integration vocabulary changes remain additive
  unless an existing value changes meaning or is removed.

Changing a package schema constant does not automatically create a tag. A tag is
created only after the PR is merged and downstream consumers are ready to import
it.

## Breaking Change Rules

Breaking changes require an explicit migration slice:

- removing or renaming a state, event type, enum value, DTO field, or JSON field
- changing the meaning of an existing value
- changing transition legality for an existing C1/C10 state pair
- changing required field presence for an existing DTO or IR envelope scope

Additive changes are allowed when older consumers can ignore the new value or
field safely and tests document that compatibility.

## Downstream Import Rule

Promotion is not complete when this repo merges. It is complete when downstream
runtime repos import the tagged module and remove duplicated facts. Migration
docs must name the downstream repo and expected import gate.
