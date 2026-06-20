# IR Event Log SSOT

> **이 문서가 `EventType` 카탈로그 / transition 분류 / `CanonicalEvent.Payload` 의 필드 셋 / reducer dispatch 와 보존 / unknown raw 필드 정책의 SSOT다.**
>
> - 책임: 어떤 이벤트가 존재하는가, 어떤 이벤트가 FSM transition 인가, 각 이벤트의 payload 가 무엇을 담는가, reducer 가 어떻게 dispatch / 보존하는가, raw 표현을 어떻게 흡수하는가
> - 비책임: 9 종 실행 지문 + 2 종 runtime identity + `FSMVersion` 의 의무 규칙은 [`./ir-schema-versioning.md`](./ir-schema-versioning.md) 가 소유. `TaskState` 집합과 합법 transition matrix 는 [`./task-lifecycle.md`](./task-lifecycle.md) 가 소유. 본 문서는 그 둘과 1:1 정합으로만 존재한다.

이 SSOT 는 [context-map.md](./context-map.md) 의 **C2 IR Event Log** context 를 채운다.

## 1. 위치 / 책임 한 줄

> **이벤트 로그가 truth-of-record 다.** `TaskState` / `TaskIR` / 모든 통계 view 는 본 로그의 reducer 결과로만 도출된다. 본 문서는 그 로그에 들어갈 수 있는 **유일한 EventType 집합** 과 그 dispatch 규칙의 단일 진실이다.

## 2. EventType taxonomy (분류)

각 EventType 은 정확히 하나의 카테고리에 속한다. 카테고리는 (a) producer context (b) FSM 영향 유무 두 축으로 정렬된다.

| Cat | 이름 | 1차 producer | FSM 영향 |
| --- | --- | --- | --- |
| **A** | Task lifecycle transitions | server/API + C5 scheduler + C4 runtime + C8 validation | **transition** (모두 `FSMVersion` 의무) |
| **B** | Runtime registry / capability lifecycle | C3 capability + C5 scheduler | 부분 transition (예: `TaskBlocked` 로 이어지는 `BlockerRaised`) |
| **C** | Provider raw → canonical (ACL 산물) | C4 adapter | 비-transition |
| **D** | Validation | C8 | 일부 transition |
| **E** | Workspace / config injection | C6 | 비-transition |
| **F** | Security / policy | C7 | 비-transition (단 정책 위반은 BlockerRaised 로 transition 가능) |
| **G** | Upgrade / runtime change | cross-cutting | 부분 transition |
| **H** | Administrative / audit | server / operator | 비-transition |

“transition event” 인지 여부는 본 문서의 카탈로그 표(§3) 의 **Transition** 열이 SSOT 다. task-lifecycle.md 의 generated Transition Surface 에 있는 reference event 이름이 본 표에 있어야 하며, 불일치는 PR 단계에서 거절된다.

## 3. EventType 카탈로그

표 안 표기:

- **Transition**: ✓ 면 FSM transition event(→ `FSMVersion` 의무 / reducer 가 state 전이 계산). ─ 면 비-transition.
- **Payload (필수 키)**: `CanonicalEvent.Payload` 안의 의무 키. envelope 의 공통 필드와 scope 별 의무 필드는 `ir-schema-versioning.md` §1.5/§2 가 owner — 여기서는 Payload 만.
- **Notes**: 발행 시점 또는 reducer 가 주의해야 할 점.

### 3.0 EventType → EventScope 할당

`CanonicalEvent.Scope` 는 발행 인스턴스의 실제 scope 다 (`ir-schema-versioning.md` §1.5). 각 EventType 마다 **legal scope 집합** 이 정해진다 — 대부분 단일 scope, 일부는 발생 맥락에 따라 두 scope 모두 합법.

**카테고리 기본 scope:**

| Category | 기본 scope | 비고 |
| --- | --- | --- |
| Cat A — Task lifecycle transitions | 대부분 **RunScope**. 단 `TaskCreated`/`TaskQueued` 는 **TaskScope** (claim 전). `TaskCancelled` 는 발생 시점에 따라 TaskScope 또는 RunScope. | claim 시점부터 RunScope 로 “승격” |
| Cat B — Runtime registry / capability | **RuntimeScope**. capability snapshot 결정 전인 `RuntimeRejected` 도 RuntimeScope(`RuntimeID` 후보를 알지만 `CapabilityFingerprint` 부재 가능). | task 무관 |
| Cat C — Provider raw → canonical | 항상 **RunScope** (provider process 는 항상 한 run 안에서 동작) | |
| Cat D — Validation | 항상 **RunScope** | |
| Cat E — Workspace / config injection | 항상 **RunScope** (workdir 은 항상 한 run 의 것) | |
| Cat F — Security / policy | `PolicyBundleLoaded`/`PolicyBundleSwitched` 는 **SystemScope**. `PolicyViolationDetected`/`SecretsScopeIssued`/`SecretsScopeRevoked` 는 발생 맥락에 따라 TaskScope 또는 RunScope. | |
| Cat G — Upgrade / runtime change | `UpgradeDetected`/`UpgradePolicyReevaluated` 는 **RuntimeScope**. `DrainStarted`/`DrainTimedOut` 은 **SystemScope** (풀 전체 영향) 또는 **RuntimeScope** (단일 runtime drain). `TaskHandedOff` 는 **RunScope**. | |
| Cat H — Administrative / audit | polymorphic. `Correction` 은 정정 대상 이벤트의 scope 와 일치. `OperatorNote` 는 SystemScope/TaskScope/RunScope 모두 가능. `Snapshot` 도 동일. | 인스턴스마다 `Scope` 명시 |

**멀티 scope EventType 표** (인스턴스 발행 시 ingest 가 한 scope 를 골라 박음):

| EventType | 합법 scope |
| --- | --- |
| `TaskCancelled` | TaskScope(Queued 에서 취소), RunScope(Running 등에서 취소) |
| `TaskCreated`, `TaskQueued` | TaskScope (only) |
| `BlockerRaised`, `BlockerResolved`, `BlockerResolvedRequeue` | TaskScope (claim 전 — 예: capability 매트릭스에 맞는 runtime 부재), RunScope (run 중) |
| `PolicyViolationDetected` | TaskScope, RunScope |
| `SecretsScopeIssued`, `SecretsScopeRevoked` | TaskScope, RunScope |
| `DrainStarted`, `DrainTimedOut` | SystemScope, RuntimeScope |
| `Correction`, `OperatorNote`, `Snapshot` | SystemScope, RuntimeScope, TaskScope, RunScope |

EventType 카탈로그가 새 EventType 을 추가하면 그 행에 (a) 기본 scope (b) 합법 scope 집합 (c) 멀티 scope 인 경우 ingest 결정 규칙을 함께 명시한다.

### 3.0.1 RunScope 의 NCV 분류 (3 묶음 + 동적 phase check)

RunScope EventType 은 `NativeConfigVersion` 의무 여부에 따라 **세** 묶음으로 갈린다 — pre-execute 와 execution-bound 사이에 “상황에 따라 다른” 묶음이 있다. 일부 EventType (예: `TaskFailed`, `BlockerRaised`, `RuntimePinViolated`) 은 `Preparing` 중에도 `Running` 중에도 발생 가능해서 EventType 단독으로 NCV 의무를 단정할 수 없다.

[`./ir-schema-versioning.md`](./ir-schema-versioning.md) §1.5.3.1 가 동적 규칙의 owner. 본 표는 EventType → 분류 매핑.

| 묶음 | NCV 의무 | 멤버 EventType |
| --- | --- | --- |
| **PreExecuteOnly** | 부재 허용 (envelope 단독으로 결정 가능) | `TaskClaimed`, `WorkdirPreparing`, `WorkdirCreated`, `RuntimePinned`, `RuntimeHandshakeOK` |
| **ExecutionBoundOnly** | **필수** (envelope 단독으로 결정 가능) | `NativeConfigInjected`, `ConfigTemplateReinjected`, `RunStarted`, `RunReportedDone`, 모든 Cat C (`TextDelta`, `ReasoningDelta`, `ToolCallStarted`, `ToolCallFinished`, `FileChanged`, `CommandStarted`, `CommandFinished`, `SessionPinned`, `ApprovalRequested`, `ApprovalResolved`, `StatusUpdate`, `UsageDelta`, `LogLine`, `ProviderUnknownEvent`), 모든 Cat D (`ValidationStarted`, `ValidationRuleExecuted`, `ValidationPassed`, `ValidationFailed`), `ReviewRequested`, `AutoApproved`, `HumanApproved`, `HumanRejected`, `WorkdirArchived`, `InputRequested`, `InputProvided` |
| **PhaseDependent** | **envelope 단독으로 결정 불가** — run phase 가 필요 | `BlockerRaised`, `BlockerResolved`, `BlockerResolvedRequeue`, `TaskFailed`, `TaskCancelled`, `TaskTimedOut`, `RuntimePinViolated`, `ReworkAccepted` |

근거:

- **PreExecuteOnly**: `TaskClaimed` ~ `RunStarted` 사이 또는 그 단계에서만 발생. workspace 가 아직 native config 를 materialize 하지 않은 시점이므로 NCV 부재 합법.
- **ExecutionBoundOnly**: `NativeConfigInjected` 이후의 “실행 사실” 이벤트. provider draft / validation / review 는 정의상 NCV 가 결정된 후 발생.
- **PhaseDependent**: `Preparing` 중 fail 일 수도, `Running` 중 fail 일 수도 있는 이벤트. 예시:
  - `TaskFailed`: `Preparing` 중 workspace 준비 실패면 NCV 없음. `Validating` 중 fail 이면 NCV 있음.
  - `RuntimePinViolated`: `Running` 직전 G5 핸드셰이크에서 pin 위반을 감지하면 NCV 가 아직 없을 수 있음. `Running` 중 위반이면 NCV 있음.
  - `BlockerRaised`: `Preparing` 중 capability missing 으로 raise 면 NCV 없음. `Running` 중 policy violation 으로 raise 면 NCV 있음.

### 3.0.2 동적 검사의 분담

- **`ValidateEnvelope(event)` 정적 검사**:
  - `PreExecuteOnly` → NCV 부재 허용
  - `ExecutionBoundOnly` → NCV 필수
  - `PhaseDependent` → envelope 단독으로 NCV 의무를 단정하지 않음 (NCV 있어도/없어도 envelope-valid)
- **`ValidateEnvelopeWithRunContext(event, RunContext)` 동적 검사**:
  - `PhaseDependent` 이고 `RunContext.NativeConfigEstablished == true` → NCV 필수
  - `PhaseDependent` 이고 `RunContext.NativeConfigEstablished == false` → NCV 부재 허용
  - 이 검사는 EventIngestor / FSM Orchestrator / reducer (run 별 view 가 있는 곳)가 수행
- “Native config established” 의 정의: 같은 `RunID` 의 IR 로그에 `NativeConfigInjected` 또는 그 이후 execution-bound 이벤트가 이미 한 번이라도 append 되었는가. orchestrator 가 IR 로그에서 결정.

### 3.1 Cat A — Task lifecycle transitions

| EventType | Transition | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- | --- |
| `TaskCreated` | ✓ | server | `goal`, `constraints[]` | 초기 진입. 이 전에는 task row 가 없음 |
| `TaskQueued` | ✓ | server | `priority` (optional) | `Created → Queued` |
| `TaskClaimed` | ✓ | scheduler (C5) | `runtimeID`, `capabilityFingerprint` | `Queued → Claimed`. (`runtime fingerprint` 는 9 종 지문 + identity 와 별도 payload 키로 한 번 더 명시 — claim 의 의도/계약을 명확히 기록) |
| `WorkdirPreparing` | ✓ | runtime (C4) | `workdirPath` | `Claimed → Preparing` |
| `RuntimePinned` | ✓ | runtime (C4) | `runtimeID`, `capabilityFingerprint` | `Preparing → Running` 의 두 트리거 중 첫째. 이 event 가 append 되지 않았는데 `RunStarted` 가 오면 reducer 가 거절(invariant 2) |
| `RunStarted` | ✓ | runtime (C4) | `runID`, `pid` (optional) | `Preparing → Running` 의 두 트리거 중 둘째. `RuntimePinned` 와 같은 `(RuntimeID, CapabilityFingerprint)` 를 가져야 한다 |
| `InputRequested` | ✓ | runtime (C4) | `prompt`, `inputType` ("approval"/"clarify"/"choice") | `Running → NeedsInput` |
| `InputProvided` | ✓ | server | `response` | `NeedsInput → Running` |
| `BlockerRaised` | ✓ | runtime / scheduler | `category` (enum), `reason`, `subject` | `Running → Blocked` 또는 `Preparing → Blocked`. category 값은 task lifecycle invariant anchor 의 외부 조건 범위와 맞아야 함 |
| `BlockerResolved` | ✓ | scheduler | `category` (해소된 카테고리) | `Blocked → Running` |
| `BlockerResolvedRequeue` | ✓ | scheduler | `category` | `Blocked → Queued` (다른 runtime 필요) |
| `RunReportedDone` | ✓ | runtime (C4) | `summary` (optional) | `Running → Validating`. agent 자기보고 — 이것 자체로 완료 아님 (invariant 4) |
| `ValidationPassed` | ✓ | validation (C8) | `validationID`, `passedRules[]` | `Validating → PatchReady` |
| `ValidationFailed` | ✓ | validation (C8) | `validationID`, `failedRules[]`, `terminal=true` | `Validating → Failed` |
| `ReviewRequested` | ✓ | policy/validation | `reviewerHint` (optional) | `PatchReady → HumanReview` |
| `AutoApproved` | ✓ | policy | `policyID` | `PatchReady → Completed` |
| `HumanApproved` | ✓ | server | `approverID` | `HumanReview → Completed` |
| `HumanRejected` | ✓ | server | `rework: bool`, `comment` | `HumanReview → ReworkQueued` (`rework=true`) 또는 `HumanReview → Cancelled` (`rework=false`) |
| `ReworkAccepted` | ✓ | server | `prevRunID`, `newRunID` | `ReworkQueued → Queued`. **새 RunID 발급은 본 event 의 invariant** |
| `TaskCancelled` | ✓ | server | `reason`, `byActor` | 비-terminal 어디서나 진입 가능 |
| `TaskTimedOut` | ✓ | scheduler | `fromState`, `limit`, `elapsed` | active subset 에서 진입 가능 |
| `RuntimePinViolated` | ✓ | runtime (C4) | `expected.runtimeID`, `expected.capabilityFingerprint`, `observed.*` | `Running → Failed` / `Validating → Failed`. silent recovery 금지 (invariant 3) |
| `TaskFailed` | ✓ | runtime / validation | `category`, `reason`, `terminal=true` | `Failed` 로의 명시적 진입 (`RuntimePinViolated`/`ValidationFailed` 외의 사유) |

### 3.2 Cat B — Runtime registry / capability lifecycle

| EventType | Transition | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- | --- |
| `RuntimeRegistered` | ─ | C5 | `runtimeID`, `capabilityFingerprint`, `compatibilityStatus`, `protocolMaturity` | G2 통과 시 |
| `RuntimeRejected` | ─ | C5 | `runtimeID` (or candidate), `reason`, `compatibilityStatus=blocked` | G2 실패 시 |
| `RuntimeFingerprintChanged` | ─ | C5 | `runtimeID`, `from.detectedFingerprint`, `to.detectedFingerprint` | G3 fingerprint 변화 감지 시. 후속으로 capability 재탐지 |
| `CapabilityReevaluated` | ─ | C3/C5 | `runtimeID`, `from.capabilityFingerprint`, `to.capabilityFingerprint`, `diff[]` | 재탐지 완료 시 |
| `LeaseInvalidated` | ─ | C5 | `runtimeID`, `reason`, `affectedTasks[]` | G3 실패 시 |
| `RuntimeHandshakeOK` | ─ | C4/C5 | `runtimeID`, `handshakeKind` ("appserver-initialize"/"stream-json-firstline") | G5 통과 시 |

### 3.3 Cat C — Provider raw → canonical (ACL 산물)

비-transition 이벤트들이다. FSM 에는 영향이 없지만 IR replay / audit / UI 에 핵심.

| EventType | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- |
| `TextDelta` | C4 adapter | `text` | streaming text chunk |
| `ReasoningDelta` | C4 adapter | `text`, `private: bool` (default true) | thinking / hidden reasoning |
| `ToolCallStarted` | C4 adapter | `toolName`, `args` | provider 의 tool 호출 시작 |
| `ToolCallFinished` | C4 adapter | `toolName`, `result` 또는 `error` | tool 결과 |
| `FileChanged` | C4 adapter | `path`, `kind` ("create"/"edit"/"delete"), `diff` (optional) | agent 가 파일을 바꿈 — daemon 측 git diff 와 교차 검증 |
| `CommandStarted` | C4 adapter | `cmd`, `cwd`, `sandboxMode` | provider 가 shell command 실행 |
| `CommandFinished` | C4 adapter | `cmd`, `exitCode`, `durationMs` | 결과 |
| `SessionPinned` | C4 adapter | `providerSessionID`, `providerThreadID` (optional) | session 식별자 영속화 — Multica 의 “early pin” 학습. resume 의 1차 키 |
| `ApprovalRequested` | C4 adapter | `approvalID`, `kind`, `payload` | provider 가 approval 프로토콜 사용 시. Claude `control_request.request_id` 는 provider-neutral run event 의 `ToolRef.ProviderRequestID` 로 보존된 뒤 ingest 시 `approvalID` 로 승격 가능 |
| `ApprovalResolved` | C4 adapter | `approvalID`, `decision` ("approve"/"deny"/"defer") | server 응답 후. Claude 는 `ProviderInputBuilder` 가 `CommandApproveTool` / `CommandRejectTool` 을 `control_response` 로 직렬화 |
| `StatusUpdate` | C4 adapter | `text` | provider 의 상태 보고(예: “파일 읽는 중”) |
| `UsageDelta` | C4 adapter | `usage.promptTokens`, `usage.completionTokens`, ... | provider usage metric delta |
| `LogLine` | C4 adapter | `level`, `text` | provider stderr 의 로그 라인 |
| `ProviderUnknownEvent` | C4 adapter | `rawType`, `rawPayload` | 어댑터가 알려지지 않은 raw event type 을 만났을 때. **FSM 전이 일으키지 않음** (provider-capability §6 / ir-schema-versioning §2.1 의 unknown 보존 원칙) |

### 3.4 Cat D — Validation

| EventType | Transition | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- | --- |
| `ValidationStarted` | ─ | C8 | `validationID`, `rules[]`, `policyBundleVersion` | `Validating` 진입 직후 |
| `ValidationRuleExecuted` | ─ | C8 | `validationID`, `ruleID`, `result` ("pass"/"fail"/"skip"/"error") | rule 단위 |
| `ValidationPassed` | ✓ (Cat A 와 동일 행) | C8 | (위 §3.1 참조) | 전체 통과 |
| `ValidationFailed` | ✓ (Cat A 와 동일 행) | C8 | (위 §3.1 참조) | 전체 실패 |

`ValidationStarted` 와 `ValidationRuleExecuted` 는 비-transition (관측 이벤트). `ValidationPassed`/`ValidationFailed` 만 transition.

### 3.5 Cat E — Workspace / config injection

| EventType | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- |
| `WorkdirCreated` | C6 | `workdirPath`, `taskID` | per-task workdir 생성 |
| `NativeConfigInjected` | C6 | `files[]` (예: `["CLAUDE.md", "AGENTS.md"]`), `nativeConfigVersion` | 주입된 파일 목록 |
| `WorkdirArchived` | C6 | `workdirPath`, `archiveURI` (optional) | terminal 후 |
| `ConfigTemplateReinjected` | C6 | `workdirPath`, `from.nativeConfigVersion`, `to.nativeConfigVersion` | T-CONFIG 트리거 후 (`runtime-upgrade-flow.md` §6) |

### 3.6 Cat F — Security / policy

| EventType | Transition | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- | --- |
| `PolicyBundleLoaded` | ─ | C7 | `policyBundleVersion`, `source` | daemon 부트 시 |
| `PolicyBundleSwitched` | ─ | C7 | `from.version`, `to.version` | 활성 번들 교체 |
| `PolicyViolationDetected` | ─ | C7 | `category` ("PROTECTED_PATH"/"SANDBOX"/"PERMISSION_BYPASS"/"SECRET_LEAK"), `subject`, `severity` | 정책 위반 감지. 후속으로 `BlockerRaised(category=POLICY_*)` 가 발행될 수 있음 |
| `SecretsScopeIssued` | ─ | C7 | `scopeID`, `ttlSeconds`, `purpose` | scoped token 발급 (값 자체는 IR 에 적지 않음) |
| `SecretsScopeRevoked` | ─ | C7 | `scopeID`, `reason` | 회수 |

### 3.7 Cat G — Upgrade / runtime change

(자세한 규칙은 [`../30-architecture/runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) 가 소유. 본 표는 카탈로그만.)

| EventType | Transition | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- | --- |
| `UpgradeDetected` | ─ | cross | `trigger` ("T-PROV"/"T-ADAP"/"T-DAEMON"/"T-POLICY"/"T-CONFIG"), `details` | G3 |
| `UpgradePolicyReevaluated` | ─ | cross | `result` ("passed"/"partial"/"blocked"), `notes` | risk profile 재검사 후 |
| `DrainStarted` | ─ | C5 | `targetRuntimeID`, `deadline` | T-DAEMON 시작 |
| `DrainTimedOut` | ─ | C5 | `targetRuntimeID` | drainDeadline 초과 |
| `TaskHandedOff` | ─ | C5 | `taskID`, `fromRuntime`, `toRuntime`, `lastEventID` | handoff 완료 |

### 3.8 Cat H — Administrative / audit

| EventType | Producer | Payload (필수) | Notes |
| --- | --- | --- | --- |
| `Correction` | server / operator | `correctsEventID`, `note`, `byActor` | 잘못 append 된 event 의 정정. **옛 event 를 mutate 하지 않고 새 event 로 표기** (append-only invariant) |
| `OperatorNote` | operator | `note` | 운영자 주석 |
| `Snapshot` | reducer | `taskIRSchemaVersion`, `lastEventID`, `snapshotURI` | replay 가속용 view snapshot 발행 (옵션 §7) |

## 4. transition vs 비-transition 의 invariant

본 문서가 강제하는 규칙:

1. **transition event 만 FSM state 를 바꾼다.** 비-transition event 는 IR 에 append 되어도 `TaskState` 가 그대로다.
2. **transition event 는 모두 `FSMVersion` 의무 필드를 가진다** (`ir-schema-versioning.md` §0 invariant 4).
3. **task-lifecycle.md 의 generated Transition Surface reference event 이름은 본 §3.1 표와 정확히 매칭** 되어야 한다. 둘 중 하나만 갱신되는 PR 은 거절.
4. **알려지지 않은 raw event type** 은 `ProviderUnknownEvent` 로만 흡수된다. 정의되지 않은 새 `EventType` 식별자를 임의로 발행하는 코드는 PR 단계에서 거절.

## 5. reducer dispatch + 보존

### 5.0 Reducer purity invariant (단단히 박는다)

> **Reducer 는 순수 함수다. 절대 `CanonicalEvent` 를 append 하지 않는다.**

event-sourcing 구조에서 reducer 가 side-effect 를 가지면 replay 마다 로그가 바뀐다 — idempotency 와 auditability 가 동시에 깨진다. 따라서 본 SSOT 는 다음 권한 분리를 강제한다.

| 역할 | 권한 | 입력 | 출력 |
| --- | --- | --- | --- |
| **Reducer** | event 를 **읽기만** 한다 | 이미 로그에 존재하는 `CanonicalEvent[]` | `TaskIR` view + `ReducerDiagnostic[]` + (실패 시) `ReducerError`. 이게 전부 — append 권한 없음, EventIngestor 호출 권한 없음. |
| **EventIngestor** (single Append API) | `CanonicalEvent` 를 append 하는 **유일한 코드 경로** | authorized caller 의 호출 + adapter 의 `ProviderEventDraft` | append-only record 에 필요한 identity / ordering / runtime identity / attribution / schema / timestamp 정책을 최종 확정한 뒤 적재. 확정 항목 목록: `EventID`, sequence/ordering metadata, `RuntimeID`, `CapabilityFingerprint`, `ActorKind`, `ActorID`, `EventSchemaVersion`, `OccurredAt`/`IngestedAt` 정책. 적재 직전 [`security-redaction.md`](./security-redaction.md) 의 C7 redaction catalog 로 `Payload` / `Unknown` 을 2차 스캔하고, redaction 이 발생하면 같은 sink batch 에 `PolicyViolationDetected(category="SECRET_LEAK_ATTEMPTED")` audit event 를 함께 append 한다. |
| **Authorized callers**: RunController (C4 orchestration), C6 workspace lifecycle handler, FSM Orchestrator, server transition layer, validation runner result handler, runtime scheduler result handler | EventIngestor 를 **호출 가능**. 직접 writer 를 갖지 않는다. | 외부 신호(API 호출, runtime 보고, validation 결과, workspace 준비/주입/archive, 운영자 명령) + `Provider.Drafts()` drain (RunController) | EventIngestor API 호출 |
| **Adapter (C4)** | event 를 append 하지 **않고** EventIngestor 를 **호출하지도** 않는다 — 관측만 | provider raw stdout/RPC | `ProviderEventDraft` (정규화 초안). 정식 정의는 [`./provider-runtime.md`](./provider-runtime.md) (C4) 가 소유 |

> **단단히 박는 한 줄**: `CanonicalEvent` append 는 **단일 EventIngestor API** 를 통해서만 수행된다. RunController(C4 orchestration), C6 workspace lifecycle handler, FSM Orchestrator, server transition layer, validation runner result handler, runtime scheduler result handler 가 그 API 의 **authorized caller** 이며 직접 writer 를 갖지 않는다. Adapter 구현체는 EventIngestor 를 모른다 — `Provider.Drafts()` 채널을 drain 하고 `EventIngestor.AppendDraft(...)` 를 호출하는 것은 **RunController** 다. Reducer 는 `ReduceResult` / `ReducerError` / `ReducerDiagnostic` 만 반환한다.

이 권한 분리가 깨지면 다음 invariant 들이 동시에 흔들린다.

- actor attribution 서버 결정 (§9): adapter 가 직접 append 하면 attribution 이 클라이언트/CLI 측 결정으로 누출.
- runtime identity 일관성: `RuntimeID` / `CapabilityFingerprint` 는 lease 와 짝지어져 있어 단일 API 에서만 일관되게 찍을 수 있다.
- ordering / sequence / fencing: 단일 API 가 `EventID` 와 sequence 를 부여해야 같은 task 의 이벤트가 단조 정렬된다.
- `EventSchemaVersion` 결정: caller 마다 다른 schema 버전을 직접 박으면 reducer 보존 규칙(`ir-schema-versioning.md` §4)이 깨진다.

코드 레벨 강제:

- `riido-contracts/ir` 의 Reducer contract 는 writer/appender / EventIngestor dependency 를 갖지 않는다.
- Adapter 패키지는 writer / EventIngestor dependency 를 직접 갖지 않는다. adapter 는 `ProviderEventDraft` 채널로만 ingest 와 통신한다.
- EventIngestor 구현은 현재 RIID-4641 범위 밖이다. 이관 전까지 daemon/control-plane 쪽 구현이 `riido-contracts/ir` 타입을 소비하고, 이관 시에도 writer port 만 알고 filesystem 을 import 하지 않아야 한다.
- EventIngestor 의 2차 secret redaction 은 C7 policy catalog 를 호출한다. C2 는 패턴 의미를 재해석하지 않고, redacted event 와 audit event 를 같은 append 호출에 담는다.

### 5.1 dispatch 규칙

reducer 는 `(Type, EventSchemaVersion) → state_delta + view_delta` 로 dispatch 한다.

1. 새 `(Type, N+1)` 추가는 항상 추가형(`change:additive`). 옛 `(Type, N)` 의 코드는 **영구 보존**.
2. 옛 페어의 동작 변경은 새 `EventSchemaVersion` 으로만 가능. 옛 코드는 삭제 금지(`ir-schema-versioning.md` §4 와 일관).
3. **미지의 `(Type, ?)` 를 만나면**: reducer 는 부분 reduce 를 하지 않고 `ReducerError(code="IR_REDUCER_INCOMPAT", eventID=...)` 를 반환한다. reducer 는 append 하지 않는다.
   - **live ingest 상황**에서는 호출자(`EventIngestor` 또는 FSM Orchestrator)가 그 ReducerError 를 받고 서버 actor 로 `BlockerRaised(category=IR_REDUCER_INCOMPAT)` 를 append 한다. attribution 은 §9 actor attribution invariant 를 따른다.
   - **replay 상황**(예: `riido-cli ir rebuild`)에서는 로그를 mutation 하지 않는다. view 생성을 중단하거나(전체 실패) 부분 view + 진단 목록을 운영자에게 보고한다.
4. 같은 `RunID` 의 transition 시퀀스는 **단조** 여야 한다. 이미 본 시퀀스 번호보다 작은 transition 이 들어오면 reducer 가 `ReducerError(code="OUT_OF_ORDER_TRANSITION", ...)` 를 반환한다(같은 EventID 재처리는 무영향 — idempotency).

### 5.2 ReducerError / ReducerDiagnostic

- `ReducerError`: 치명적 — 이 task 의 view 를 더 이상 진행 불가. 호출자가 받아서 처리.
- `ReducerDiagnostic`: 비치명적 — reduce 는 계속 진행하지만 운영자가 봐야 할 관측. 예: 알려진 옛 schemaVersion 의 deprecated 필드 사용.

둘 다 reducer 의 **반환값** 이다. reducer 가 자체적으로 event 를 발행하거나 로그를 mutate 하지 않는다.

## 6. unknown raw 필드 / ACL 보존

어댑터(C4)는 provider raw 표현을 `CanonicalEvent` 로 변환할 때 다음을 지킨다.

1. raw payload 의 **알려지지 않은 키** 는 `CanonicalEvent.Unknown` 에 보존 (provider-capability.md §6 + ir-schema-versioning.md §2.1 의 unknown 정신).
2. raw event 의 **알려지지 않은 type** 은 `ProviderUnknownEvent(rawType, rawPayload)` 로 변환. FSM 전이 없음.
3. raw 표현을 “해석” 해서 의미를 추가한 경우 `CanonicalEvent.Payload.derived = true` 표기.

본 정책은 IR replay 안정성의 핵심이다. 어댑터가 “알려지지 않은 데이터” 를 드롭하면 그 시점의 사실이 영구 소실된다.

## 7. replay / snapshot 정책

1. **replay 는 reducer 가 IR 이벤트들을 처음부터 다시 흘려보내는 것**으로 정의된다. snapshot 이 있어도 truth-of-record 는 이벤트 로그.
2. snapshot 은 **선택적 가속기**이다. 한 task 가 N 개 이벤트를 넘으면 reducer 가 `Snapshot` 이벤트를 append 하고 view 를 저장할 수 있다. 다음 replay 는 그 snapshot 부터 재개.
3. snapshot 의 `taskIRSchemaVersion` 이 현재 reducer 의 schema 와 다르면 snapshot 을 무시하고 처음부터 replay (silent migration 금지).
4. snapshot 보관 기간 / 임계값은 운영 정책 (`open-questions.md`).

## 8. 인접 SSOT 와의 계약

| 인접 SSOT | 본 문서가 요구/제공 |
| --- | --- |
| [`./task-lifecycle.md`](./task-lifecycle.md) (C1) | §3.1 의 transition event 이름이 generated Transition Surface 와 1:1 매칭. transition 여부의 정식 분류는 본 문서. |
| [`./ir-schema-versioning.md`](./ir-schema-versioning.md) | `CanonicalEvent` 필수 필드(9+2+FSMVersion) 규칙은 그 문서가 소유. 본 문서는 카탈로그와 dispatch 만. |
| [`./provider-capability.md`](./provider-capability.md) (C3) | Cat B/G 의 capability 관련 이벤트는 capability fingerprint 변경 흐름에 대응 |
| [`./provider-runtime.md`](./provider-runtime.md) (C4) | Cat C 가 어댑터 ACL 의 출력. raw → canonical 매핑 표는 C4 가 소유. 본 문서는 canonical 의 EventType 명만 소유 |
| [`./runtime-scheduling.md`](./runtime-scheduling.md) (C5) | Cat A 중 `TaskClaimed`/`BlockerResolved`/`TaskTimedOut`, Cat B 전체, Cat G 의 `Drain*`/`TaskHandedOff` 는 C5 producer |
| [`./workspace.md`](./workspace.md) (C6) | Cat E 의 producer |
| [`./security.md`](./security.md) (C7) | Cat F 의 producer |
| [`./validation.md`](./validation.md) (C8) | Cat D + Cat A 의 `ValidationPassed`/`ValidationFailed` 의 producer |
| [`../30-architecture/compatibility-gate.md`](../30-architecture/compatibility-gate.md) | Cat B/G 의 발행 “시점” 을 소유 |
| [`../30-architecture/runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) | Cat G 의 “발행 조건” 을 소유 |

## 9. actor attribution invariant

(Multica 의 actor attribution 버그 학습 — 자료조사 §5.9.3)

1. **`ActorKind` / `ActorID` 는 서버측 transition 처리 시에만 결정된다.** 클라이언트/CLI/agent process 가 직접 결정하지 않는다.
2. agent 가 발행한 event(예: `TextDelta`, `ToolCallStarted`)는 `ActorKind=agent` + `ActorID=<runID>` 로 일괄 기록되며, 그 이상의 attribution 은 agent 입력으로 못 바꾼다.
3. human 의 행위(`HumanApproved`, `HumanRejected`, `OperatorNote`, `Correction`)는 인증된 사용자 컨텍스트에서만 `ActorKind=human` + `ActorID=<userID>`. 인증되지 않은 채널의 동일 event 는 거절.

## 10. 미결정 / 오픈 이슈

`open-questions.md` 위임.

- `Q-IR-001`: 이벤트 로그 물리 저장(append-only 파일 vs DB table). 현재안: DB table + WAL.
- `Q-IR-006`: snapshot 의 보관 기간 / 정밀도 / 주기 / 저장 위치.
- `Q-IR-003`: `ProviderUnknownEvent` 가 일정 임계 이상 누적될 때 자동 `Blocked(category=IR_UNKNOWN_OVERFLOW)` 전이를 둘지.
- `Q-IR-004`: `Correction` event 가 reducer view 에 어떻게 반영되는가(옛 이벤트의 “효과 취소” 방식: undo-then-redo vs ledger-style).

## 11. version-affecting changes

- 새 EventType 추가: `change:additive` (transition 표 / reference 표 동시 갱신 시).
- 기존 EventType payload 필수 키 추가: `change:additive` (옛 reducer 와 호환 유지).
- 기존 EventType payload 의 의미 변경: `change:breaking-ir` + `EventSchemaVersion` 증가 + 옛 reducer 보존.
- transition ↔ 비-transition 분류 변경: `change:breaking-ir` + `FSMSchemaVersion` 증가 + task-lifecycle generated Transition Surface 와 동시 갱신.
- Category 자체 변경(추가/통합/제거): `change:breaking-policy` + module-decomposition.md 갱신.
