# Distribution Host Integration Contracts

> Riido task: RIID-4670 `[Contracts] distribution/provider routing status contract migration`

This document is the public SSOT for the C11/C10 host-integration contract
vocabulary shared by daemon-side distribution metadata and control-plane
provider status routing.

## Ownership

`riido-contracts/hostintegration` owns only vocabulary that more than one
repository must compile against:

- distribution channel values
- store-managed distribution classification
- server-facing provider routing status values

This package does not own app data root selection, local IPC, workspace grants,
consent ledgers, review/demo mode, privacy metadata allowlist artifacts,
provider executable discovery, provider login probes, provider process
execution, control-plane HTTP handlers, Terraform, AWS accounts, or deployment
evidence.

## Distribution Channels

The public distribution channels are:

- `developer-id`
- `mac-app-store`
- `msix-sideload`
- `msix-store`
- `dev-local`

`DistributionChannel.Valid` is the executable gate for this vocabulary.
`DistributionChannel.StoreManaged` returns true only for `mac-app-store` and
`msix-store`, because those channels are subject to store-review constraints.

## Provider Routing Status

The public provider routing statuses are:

- `available`
- `login-required`
- `unsupported`
- `store-blocked`

`ProviderRoutingStatus.Valid` is the executable gate for this vocabulary.

These statuses are server-facing. They intentionally do not contain provider
executable paths, workspace absolute paths, provider tokens, API keys, or raw
environment values.

## Migration State

RIID-4670 moves the shared vocabulary from the former private
`riido_daemon/internal/hostintegration` package into the public
`riido-contracts/hostintegration` package.

The daemon may later replace its private enum definitions with this package.
The control plane may consume this package before moving provider status DTOs.
Both follow-up migrations must keep private imports out of public repositories.
