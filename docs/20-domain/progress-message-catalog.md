# AI Agent Progress Message Catalog

> Owner: `progressmessage/catalog.dsl.riido.json`

AI Agent runtime progress copy is a shared contract between daemon, control
plane, and client-facing SSE projection. The catalog keeps provider/runtime
progress cheap and deterministic: daemon sends an integer message code plus
small arguments, control-plane renders the fixed localized copy, and frontend
clients keep receiving the existing `message` string.

AI Agent runtime progress messages are fixed, translated, integer-coded, append-only, and rendered before public SSE delivery.

## Rules

- Message codes are integer, stable, and append-only.
- A code must never be removed or reassigned. If product copy is no longer used,
  keep the code and mark usage as `reserved`.
- The active catalog is intentionally small. `max_messages=15` prevents the SSE
  surface from becoming an unbounded prompt-generated copy channel.
- `usage=required` messages are system-critical and must remain supported by
  daemon/control-plane fallbacks.
- `usage=active` messages may be emitted by current daemon runtime parsing.
- `usage=reserved` messages are not part of active daemon emission yet, but are
  held for compatibility or future product copy.
- Public frontend SSE shape does not change. The generated/internal code may
  carry `message_code` and `message_args`, but the public UI still consumes the
  rendered `message`.

## Projection

```text
progressmessage/catalog.dsl.riido.json
  -> progressmessage/catalog.ir.riido.json
  -> daemon structured telemetry parser
  -> control-plane progress renderer
  -> existing SSE message string
```

Provider prompts should prefer:

```text
<riido_log>{"code":1101,"args":{"label":"팀 프로젝트","description":"팀의 프로젝트 목록, 진행 상태, 우선순위와 담당자 정보를 조회해 요약을 준비 중. . ." }}<end>
```

Legacy raw Korean progress is tolerated only as a daemon-side compatibility
fallback. New provider/runtime integrations should emit the structured JSON
payload so control-plane can render fixed copy without inferring meaning from
free text.
