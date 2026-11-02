# Shared reaction emoji identity

Imported unchanged from riido-monolith-service PR885 (`0bc7bd293ce3bb531fa42402adc10e50214769b5`), `internal/platform/reactionemoji`. The fixed v1 emoji catalog is the existing Riido storage normalization contract. Control-plane thread reactions use the same normalization so optional presentation selectors, mixed case and glyph aliases do not create distinct actors on one chip. This package owns reusable emoji identity; service-specific reaction aggregation and storage remain in their services.
