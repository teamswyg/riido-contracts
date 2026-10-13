# Assignment workspace identity

RIID-5827 adds optional `workspace_id` to AssignRequest and Assignment.
The control-plane sets this from the authenticated workspace and validated task
context. Daemon repository bindings use that exact workspace/task pair.
Omission remains compatible with legacy senders; the field does not authorize
filesystem access or replace principal/task permission checks.
This change preserves the current main run, session, approval and worktree fields.
