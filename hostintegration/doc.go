// Package hostintegration owns shared C11/C10 host-integration contracts.
//
// The current package surface is intentionally small: distribution channel
// vocabulary and server-facing provider routing status vocabulary. Runtime
// host adapters, app data paths, local IPC, consent ledgers, workspace grants,
// provider discovery, and provider process execution stay in the runtime
// repositories.
package hostintegration
