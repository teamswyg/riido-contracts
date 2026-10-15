# Task request identity · RIID-5828

2026-10-12

A caller-supplied request_id identifies one explicit request, scoped by workspace, task, agent and authenticated author. Retries retain that identity; a subsequent request uses a new identity even while the same agent is busy. The additive route queues the subsequent request without replacing other agents or dropping its prompt. Legacy callers without request_id retain their existing behavior.

Structured Agent mentions in new task/subtask comments and replies issue requests. Edits, deletions and plain text mentions do not. Assignment membership persists independently of execution completion or failure. Stopping an execution does not remove its membership.

Thread and history collection responses always include assigned_agent_ids (empty when no current membership exists). Assignment BDD scenarios accept optional request_id/instruction and distinguish retries from new explicit requests; neither execution history nor a retry restores removed membership.
