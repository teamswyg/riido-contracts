# Task request identity · RIID-5828

2026-10-12

A caller-supplied request_id identifies one explicit request, scoped by workspace, task, agent and authenticated author. Retries retain that identity; a subsequent request uses a new identity even while the same agent is busy. The additive route queues the subsequent request without replacing other agents or dropping its prompt. Legacy callers without request_id retain their existing behavior.

Structured Agent mentions in new task/subtask comments and replies issue requests. Edits, deletions and plain text mentions do not. Assignment membership persists independently of execution completion or failure. Stopping an execution does not remove its membership.

Thread and history collection responses always include assigned_agent_ids (empty when no current membership exists). Assignment BDD scenarios accept optional request_id/instruction and distinguish retries from new explicit requests; neither execution history nor a retry restores removed membership.

조회 페이지도 작업 계약에 포함한다. 선택적 query_parameters의 이름·타입·기본값·최솟값을 DSL→IR→OpenAPI로 보존하며, 누락된 쿼리 인자로 실행 이력 페이지가 유실되지 않게 한다.
