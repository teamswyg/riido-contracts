# Assignment workspace identity

RIID-5827 adds optional `workspace_id` to AssignRequest and Assignment. The server
sets it from the authenticated principal when accepting client work and preserves
it through poll and persisted snapshots. Daemons use it for exact host-owned
repository bindings; they never treat a prompt as an authorized filesystem path.
Older payloads omit it; existing component/task fallback remains for compatibility.
This field is additive and does not create a workspace grant or alter authorization.
