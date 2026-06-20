# IR Schema Versioning SSOT

> **이 문서가 `CanonicalEvent` / `TaskIR` / FSM event 의 schema version 정책과 reducer 보존 규칙의 SSOT다.**
>
> - 책임: IR 의 어떤 필드가 버전을 갖는가, append-only 이벤트 로그와 view migration 의 분리, 새 schemaVersion 도입 절차, 모든 IR 행이 의무적으로 갖는 “버전 지문”
> - 비책임: 어떤 `EventType` 들이 존재하는지의 카탈로그는 [`ir-event-log.md`](./ir-event-log.md) 가 소유. FSM 의 상태 집합은 [`task-lifecycle.md`](./task-lifecycle.md) 가 소유. capability 와의 정합 게이트는 [`../30-architecture/compatibility-gate.md`](../30-architecture/compatibility-gate.md) 가 소유

이 SSOT 는 [ADR 0001](../40-decisions/0001-versioning-is-capability-based.md) 의 6번 규칙(“모든 IR 행은 어떤 버전 묶음 위에서 만들어졌는지 자체기술적이어야 한다”)의 구현이다.

## 0. 본 문서가 박는 핵심 invariant (요약)

1. **`CanonicalEvent` 는 append-only 다.** 한 번 발행된 이벤트의 payload 를 mutation 하지 않는다. 정정이 필요하면 새 이벤트(`Type=Correction`)를 append 한다.
2. **`TaskIR` view 는 재생성 가능하다.** reducer/migration 으로 옛 이벤트들로부터 새 shape 의 view 를 다시 만들 수 있다. 디스크에 저장된 `TaskIR` 은 캐시이지 truth-of-record 가 아니다.
3. **의무 필드는 scope 별로 다르다 (`EventScope` 4 종 — §1.5)**. 모든 이벤트가 같은 필드 세트를 갖지 않는다. 예: `TaskQueued` 는 아직 claim 전이라 `RuntimeID`/`CapabilityFingerprint`/`NativeConfigVersion` 이 없고, `RuntimeDetected` 는 `TaskID` 가 없다. 공통 envelope 필드 + scope 별 의무 필드 표는 §1.5 가 소유한다.
4. **`FSMVersion` 은 transition event 에만 의무**. 비-transition 이벤트에는 `FSMVersion=0` 이 합법. (§2.3)
5. **runtime identity 페어 invariant**: `RunScope` 이벤트는 모두 `(RuntimeID, CapabilityFingerprint)` 를 가지며, 같은 `(TaskID, RunID)` 의 모든 `RunScope` 이벤트는 동일한 페어를 가진다(runtime pinning — §2.2). 다른 scope 이벤트는 본 페어를 요구하지 않는다.
6. **fake placeholder 금지**. 아직 결정되지 않은 `RuntimeID` / `CapabilityFingerprint` / `NativeConfigVersion` 을 `"unknown"` / `"none"` / `"pending"` / `""` 같은 sentinel 문자열로 채우지 않는다. 필드의 **존재 여부** 는 `EventScope` 가 결정한다. 부재해야 할 필드는 부재로 둔다.

이 여섯이 흔들리면 capability-based versioning 의 “행위 자체기술성” 약속이 깨진다. ADR 0001 의 7 약속 전체는 [`runtime-versioning.md`](./runtime-versioning.md) §0 가 소유한다.

## 1. 두 가지 버전 축의 분리

IR 에서는 다음 두 축을 **혼동하지 않는다**.

| 축 | 누가 갖는가 | 변화 방식 |
| --- | --- | --- |
| **EventSchemaVersion** | `CanonicalEvent` 의 한 줄 한 줄 | append-only — 한 번 발행된 `(eventType, eventSchemaVersion)` 의 의미는 영구 동결 |
| **TaskIRSchemaVersion** | reducer 가 만든 `TaskIR` snapshot/view | migratable — 새 reducer 가 옛 events 를 새 IR shape 로 재계산 가능 |

핵심 원칙:

> **이벤트는 append-only, view 만 migrate.**

이벤트 로그는 절대 다시 쓰지 않는다. 우리가 IR shape 을 진화시키고 싶다면, 새 `TaskIRSchemaVersion` 을 가진 reducer 를 추가하고 옛 이벤트들을 그대로 다시 흘려보낸다. 옛 reducer 도 영구 보존된다(같은 logs 의 다른 view 로 함께 살 수 있게).

## 1.5 EventScope — scope 별 의무 필드

모든 `CanonicalEvent` 가 동일한 필드 세트를 갖지 않는다. 예: `TaskQueued` 는 아직 어느 runtime 이 잡을지 모르므로 `RuntimeID` 가 없고, `RuntimeDetected` 는 어느 task 와도 무관하다. 그래서 본 문서는 **4 종 EventScope** 를 두고, scope 별 의무 필드를 분리한다.

### 1.5.1 4 EventScope

| Scope | 의미 | 대표 이벤트 |
| --- | --- | --- |
| `SystemScope` | daemon / 운영자 / 정책 번들 등 전체 시스템 레벨. 특정 runtime 이나 task 에 묶이지 않음. | `DaemonStarted` (장차), `PolicyBundleLoaded`, `PolicyBundleSwitched`, system-level `OperatorNote`/`Snapshot` |
| `RuntimeScope` | runtime 슬롯 또는 capability 스냅샷에 대한 이벤트. 특정 task 에 묶이지 않음. | `RuntimeRegistered`, `RuntimeRejected`, `RuntimeFingerprintChanged`, `CapabilityReevaluated`, `RuntimeHandshakeOK`(시스템 차원의 healthcheck 인 경우), `UpgradeDetected` |
| `TaskScope` | task 라이프사이클의 매우 이른 단계 — 아직 어느 runtime/run 에도 묶이지 않은 이벤트. | `TaskCreated`, `TaskQueued`, (그리고 Queued 에서 발생한 `TaskCancelled` 같은 경우) |
| `RunScope` | 특정 task 의 특정 run 안에서 발생한 이벤트. claim 이후 모든 lifecycle 이 여기 속한다. | `TaskClaimed`, `WorkdirPreparing`, `RuntimePinned`, `RunStarted`, 모든 Cat C provider draft, `ValidationStarted`/`Passed`/`Failed`, `ReviewRequested`, `HumanApproved`/`Rejected`, `RuntimePinViolated`, `TaskFailed`, `TaskCompleted` 등 |

scope 의 관계는 부분 순서가 아니라 **분류** 다. 한 이벤트는 정확히 하나의 scope 를 갖는다. 단, 일부 EventType 은 발생 시점에 따라 두 scope 중 하나가 될 수 있다(아래 §1.5.4).

### 1.5.2 공통 envelope (모든 scope 공통 의무)

| 필드 | 의미 |
| --- | --- |
| `EventID` | ULID / UUID7. ingest 가 발급 |
| `OccurredAt` | RFC3339Nano timestamp |
| `EventSchemaVersion` | (Type, SchemaVersion) 의 의미는 영구 동결 |
| `EventScope` | 본 §1.5 의 4 종 중 하나 |
| `Type` | EventType 카탈로그 멤버 ([`./ir-event-log.md`](./ir-event-log.md)) |
| `ActorKind` | 서버 결정 — `agent`/`daemon`/`human`/`system` |
| `ActorID` | 서버 결정 |
| `RiidoDaemonVersion` | 이벤트를 만든 데몬 semver |
| `PolicyBundleVersion` | 이벤트 시점에 활성이던 정책 번들 버전 |
| `Payload` | 이 EventType + EventSchemaVersion 의 허용 키 집합 |
| `Unknown` | adapter ACL 잔여 |

### 1.5.3 scope 별 추가 의무 필드

| Scope | 추가 의무 | 부재해야 함 |
| --- | --- | --- |
| `SystemScope` | (공통 envelope 외 없음) | `TaskID`, `RunID`, `RuntimeID`, `CapabilityFingerprint`, `ProviderKind`, `ProtocolKind`, `ProviderVersion`, `AdapterID`, `AdapterVersion`, `ProtocolVersion`, `NativeConfigVersion`, `FSMVersion` |
| `RuntimeScope` | `RuntimeID`, `CapabilityFingerprint` (capability snapshot 이 결정된 시점부터). `ProviderKind`, `ProtocolKind`, `ProviderVersion` 같은 capability 관련 필드는 그 시점에 결정된 것만 채움(없으면 부재). | `TaskID`, `RunID`, `FSMVersion`, `NativeConfigVersion` |
| `TaskScope` | `TaskID` | `RunID`, `RuntimeID`, `CapabilityFingerprint`, `ProviderKind`, `ProtocolKind`, `ProviderVersion`, `AdapterID`, `AdapterVersion`, `ProtocolVersion`, `NativeConfigVersion`. **FSMVersion**: TaskScope transition event 는 `FSMVersion` 필수(예: `TaskQueued` 는 `Created → Queued` transition). |
| `RunScope` | 3 tier 로 분리 (§1.5.3.1). **Run identity**: `TaskID`, `RunID`, `RuntimeID`, `CapabilityFingerprint`. **Runtime capability**: `ProviderKind`, `ProtocolKind`, `ProviderVersion`, `AdapterID`, `AdapterVersion`, `ProtocolVersion`. **Execution context (조건부)**: `NativeConfigVersion` — execution-bound RunScope event 에만 필수, pre-execute RunScope event 는 부재 허용. transition event 는 `FSMVersion` 추가 필수. | (없음 — RunScope 는 최대 envelope) |

### 1.5.3.1 RunScope 의 3 tier — pre-execute vs execution-bound

`TaskClaimed` 시점에는 `(RuntimeID, CapabilityFingerprint)` 가 결정되지만 workspace 가 아직 준비되지 않아 `NativeConfigVersion` 이 부재한다. 그래서 RunScope 의무 필드를 3 tier 로 분리한다.

- **Run identity (RunScope 전체 의무)** — `TaskID`, `RunID`, `RuntimeID`, `CapabilityFingerprint`. 어떤 task 의 어떤 run 이 어떤 runtime + capability snapshot 위에 있는가.
- **Runtime capability (RunScope 전체 의무)** — `ProviderKind`, `ProtocolKind`, `ProviderVersion`, `AdapterID`, `AdapterVersion`, `ProtocolVersion`. capability 결정에 쓰인 표면.
- **Execution context (조건부 의무)** — `NativeConfigVersion`. EventType 을 **세 묶음** 으로 분류한다(envelope 단독 결정 가능한 두 묶음 + run phase 가 필요한 한 묶음):

  | 분류 | envelope 단독 NCV 의무 | 대표 EventType |
  | --- | --- | --- |
  | **PreExecuteOnly** | 부재 허용 | `TaskClaimed`, `WorkdirPreparing`, `WorkdirCreated`, `RuntimePinned`, `RuntimeHandshakeOK` |
  | **ExecutionBoundOnly** | **필수** | `NativeConfigInjected`, `ConfigTemplateReinjected`, `RunStarted`, `RunReportedDone`, 모든 Cat C, 모든 Cat D, `ReviewRequested`/`AutoApproved`/`HumanApproved`/`HumanRejected`, `WorkdirArchived`, `InputRequested`/`InputProvided` |
  | **PhaseDependent** | 단독 결정 불가 — run phase 필요 | `BlockerRaised`, `BlockerResolved`, `BlockerResolvedRequeue`, `TaskFailed`, `TaskCancelled`, `TaskTimedOut`, `RuntimePinViolated`, `ReworkAccepted` |

  분류 전체 표의 정식 owner 는 [`./ir-event-log.md`](./ir-event-log.md) 와 `ir.NativeConfigRequirementOf`.

#### 1.5.3.2 PhaseDependent 의 동적 규칙

`PhaseDependent` EventType 은 같은 RunID 가 어느 phase 에 있느냐에 따라 NCV 의무가 달라진다. envelope alone 으로 단정하지 않고, 호출자가 run context 를 같이 넘긴다.

- `ValidateEnvelope(event)` (정적 검사): `PhaseDependent` 이면 NCV 부재/존재 모두 envelope-valid.
- `ValidateEnvelopeWithRunContext(event, runContext)` (동적 검사):
  - `runContext.NativeConfigEstablished == true` (즉, 같은 RunID 에 `NativeConfigInjected` 또는 그 이후 ExecutionBoundOnly 이벤트가 이미 한 번이라도 append 됨) → NCV **필수**
  - `runContext.NativeConfigEstablished == false` → NCV 부재 허용

이 동적 규칙은 EventIngestor / FSM Orchestrator / reducer 같이 run 별 view 가 있는 계층이 강제한다. envelope validator 는 EventType 만으로 단정할 수 있는 부분까지만 책임진다.

> **요약**: `runtime pinning` = `(RuntimeID, CapabilityFingerprint)`. **execution pinning** = 그 위에 `PolicyBundleVersion` + `NativeConfigVersion` 추가. NCV 는 capability fingerprint 의 입력이 아니다 (provider-capability generated Invariant Anchors 의 명시적 결정 — #27i 라운드). NCV 의 의무 여부는 EventType 의 정적 분류(PreExecuteOnly / ExecutionBoundOnly / PhaseDependent) + run context (PhaseDependent 일 때만) 의 조합으로 결정된다 (#27j 라운드).

### 1.5.4 멀티 scope EventType

같은 `Type` 이라도 발생 맥락에 따라 다른 scope 가 될 수 있다. 다음 EventType 은 두 scope 모두 합법:

- `TaskCancelled` — Queued 에서 발생하면 `TaskScope` (RuntimeID 없음), Running 에서 발생하면 `RunScope`.
- `BlockerRaised` / `BlockerResolved` / `BlockerResolvedRequeue` — task-pre-claim 사유면 `TaskScope`, run 중이면 `RunScope`.
- `Correction` / `OperatorNote` — system-wide 면 `SystemScope`, task 묶이면 `TaskScope` 또는 `RunScope`.

이 경우 **인스턴스의 실제 scope** 는 발행 시점에 ingest 가 결정해 `EventScope` 필드에 박는다. 멀티 scope 인 EventType 의 합법 scope 집합은 [`./ir-event-log.md`](./ir-event-log.md) 와 `ir` envelope validator 가 owner.

### 1.5.5 fake placeholder 절대 금지

- `"unknown"` / `"none"` / `"pending"` / `""` (빈 문자열) 등 sentinel 값으로 `RuntimeID` / `CapabilityFingerprint` / `NativeConfigVersion` / `ProviderVersion` / `AdapterID` 등 식별자 필드를 채우지 않는다.
- “이 단계에서는 그 값이 아직 없다” 면 **EventScope 를 한 단계 낮춰** 해당 필드를 정당하게 부재시킨다.
- ingest 와 reducer 모두 sentinel 문자열을 발견하면 거절(`ReducerError(code="FAKE_PLACEHOLDER")` 또는 ingest 측 envelope 검증 실패).

## 2. CanonicalEvent 의 의무 필드 (scope-aware “실행 지문 / runtime fingerprint”)

모든 `CanonicalEvent` 는 공통 envelope (§1.5.2) 를 가지며, 자기 `EventScope` 에 해당하는 추가 필드(§1.5.3) 를 갖는다. 필드 누락이나 sentinel 채움이 발견되면 ingest envelope 검증 또는 reducer 가 이벤트를 거절한다.

“실행 지문” 의 9 종 (`ProviderKind`, `ProtocolKind`, `ProviderVersion`, `AdapterID`, `AdapterVersion`, `ProtocolVersion`, `RiidoDaemonVersion`, `PolicyBundleVersion`, `NativeConfigVersion`) 은 7 version + 2 kind 의 분류로 구성된다. kind 는 semver 가 아니므로 version 과 같은 슬롯으로 묶지 않는다. 그러나 본 §1.5 처럼 이 9 종 전부가 모든 이벤트의 의무는 아니다 — 의무 범위는 scope 가 결정한다.

```go
type CanonicalEvent struct {
    // 공통 envelope — 모든 scope 의무 (§1.5.2)
    EventID              string       // ULID / UUID7. ingest 가 발급
    OccurredAt           time.Time    // RFC3339Nano
    EventSchemaVersion   int          // (Type, SchemaVersion) 의 의미는 영구 동결
    Scope                EventScope   // System / Runtime / Task / Run (§1.5.1)
    Type                 string       // EventType 카탈로그 멤버
    ActorKind            string       // 서버 결정 — agent / daemon / human / system
    ActorID              string       // 서버 결정
    RiidoDaemonVersion   string       // 데몬 semver
    PolicyBundleVersion  string       // 발행 시점 활성 정책 번들 버전 (SystemScope 도 포함)
    Payload              map[string]any
    Unknown              map[string]any

    // scope 별 추가 필드 — §1.5.3 표대로 의무 / 부재 규칙 다름
    TaskID                string     // TaskScope/RunScope 필수
    RunID                 string     // RunScope 필수
    RuntimeID             string     // RuntimeScope(capability snapshot 결정 후) / RunScope 필수
    CapabilityFingerprint string     // RuntimeScope(capability snapshot 결정 후) / RunScope 필수
    ProviderKind          string     // RunScope 필수, RuntimeScope 선택
    ProtocolKind          string     // RunScope 필수, RuntimeScope 선택
    ProviderVersion       string     // RunScope 필수
    AdapterID             string     // RunScope 필수
    AdapterVersion        string     // RunScope 필수
    ProtocolVersion       string     // RunScope 필수
    NativeConfigVersion   string     // execution-bound RunScope 필수 (§1.5.3.1). pre-execute RunScope 부재 허용. 다른 scope 부재.

    // FSM transition event 전용 (TaskScope/RunScope 의 transition 일 때 필수)
    FSMVersion            int        // 0 은 비-transition 이벤트만 허용
}

type EventScope string

const (
    EventScopeSystem  EventScope = "system"
    EventScopeRuntime EventScope = "runtime"
    EventScopeTask    EventScope = "task"
    EventScopeRun     EventScope = "run"
)
```

부재 규칙 위반(`Unknown` 같은 sentinel 채움, 또는 부재해야 할 scope 에서 필드가 차 있음)은 ingest envelope 검증에서 거절된다. fake placeholder 금지 — §1.5.5.

### 2.1 runtime identity vs runtime fingerprint

`RuntimeID` 와 `CapabilityFingerprint` 는 9 종 실행 지문과 별개의 2 종 **runtime identity** 다. 9 종 지문이 “어떤 소프트웨어 스택 위에서 관측됐는가” 를 답한다면, 이 둘은 “정확히 어느 runtime 인스턴스 / 어느 capability 스냅샷 위에서 관측됐는가” 를 답한다.

| 필드 | 안정성 | 무엇이 바뀌면 새 값이 되는가 |
| --- | --- | --- |
| `RuntimeID` | 가장 안정 — 같은 등록 슬롯이면 재탐지에도 유지 | runtime 슬롯이 새로 등록될 때만 |
| `CapabilityFingerprint` | 중간 — capability 스냅샷이 바뀔 때마다 | provider binary, 어댑터, surface flag, 기본 sandbox/approval, policy bundle 중 하나라도 바뀔 때. `NativeConfigVersion` 은 task/run execution context 라서 입력이 아니다. |
| `ProviderVersion` (raw `DetectedFingerprint`) | 가장 변동성 큼 | provider binary 자체가 바뀔 때 |

`CapabilityFingerprint` 의 정확한 입력은 [`provider-capability.md`](./provider-capability.md) generated reader 와 `provider/capability.CapabilityFingerprintInput` 이 소유한다. 본 문서는 “이벤트에 필수” 라는 강제 규칙만 가진다.

### 2.2 runtime pinning 과의 정합 (RunScope 한정)

같은 `(TaskID, RunID)` 의 모든 **RunScope** 이벤트는 **동일한 `(RuntimeID, CapabilityFingerprint)` 위에서 발생** 해야 한다. 같은 run 안에서 둘 중 하나라도 바뀌면 그 run 은 `RuntimePinViolated` 이벤트와 함께 `Failed` 로 강제 전이된다(자세한 규칙은 [`../30-architecture/runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) §3 “runtime pinning 메커니즘”).

TaskScope / RuntimeScope / SystemScope 이벤트는 본 invariant 의 적용 대상이 아니다(애초에 RunID/RuntimeID 페어가 부재하거나 task 와 무관).

### 2.3 FSMVersion 의 강제 범위 (scope 기반)

- **강제**: TaskScope 또는 RunScope 이벤트 중 `Type` 이 FSM 전이를 일으키는 이벤트(예: `TaskQueued`, `TaskClaimed`, `RuntimePinned`, `BlockerRaised`, `ValidationPassed`, `TaskCompleted`, `ReworkQueued` 등) 면 `FSMVersion` 이 0 이 아닌 값이어야 한다.
- **선택**: 비-transition 이벤트(예: `TextDelta`, `ReasoningDelta`, `ToolCallStarted`, `FileChanged`, `CommandFinished`) 면 `FSMVersion` 을 0(미기재) 으로 둘 수 있다.
- **부재**: SystemScope / RuntimeScope 이벤트는 FSM 영향이 없으므로 `FSMVersion` 을 0(부재)로 둔다.
- 어느 EventType 이 transition 인가의 정식 분류 + 어느 scope 가 합법인가는 [`ir-event-log.md`](./ir-event-log.md) 와 `ir` package predicates 가 소유한다. 본 문서는 강제 규칙만 둔다.

규칙:

1. `EventSchemaVersion` 의 초기값은 `1`. 의미가 바뀌면 `(Type, EventSchemaVersion)` 페어를 새로 만들고 reducer 에 케이스를 추가한다. 옛 케이스는 삭제 금지.
2. `Payload` 는 자유 키 맵이지만, 각 `(Type, EventSchemaVersion)` 의 허용 키 집합은 코드 안에 schema 로 박힌다(예: 향후 `ir/eventschemas/`).
3. `Unknown` 은 어댑터 ACL 단계에서 “수용은 했지만 도메인 의미를 못 정한” raw 필드를 그대로 옮겨 담는다. reducer 는 `Unknown` 을 무시하지만 **drop 하지 않는다**(나중에 재해석할 수 있게).
4. `ActorKind` 는 서버 transition 단계에서 결정한다. 클라이언트/CLI 가 직접 결정하지 않는다(이 규칙은 Multica 의 actor-attribution 버그를 회피하기 위한 invariant).

## 3. TaskIR 의 의무 필드

`TaskIR` 은 reducer 가 produce 하는 view 다.

```go
type TaskIR struct {
    TaskID               string
    TaskIRSchemaVersion  int       // 이 view 를 만든 reducer 의 schema version
    ReducedByDaemon      string    // 마지막 reducer 가 돌았던 데몬 semver
    LastEventID          string    // 어느 EventID 까지 반영한 view 인가
    LastEventSeq         int64     // 단조 증가 sequence (per task)

    // FSM
    FSMSchemaVersion     int       // task-lifecycle 의 transition 표 버전
    State                TaskState // 현재 상태
    StateEnteredAt       time.Time

    // provenance per provider run
    Runs                 []ProviderRunIR

    // 도메인 view (요약/관심사별로 인덱싱된 derived data)
    ToolCalls            []ToolCallIR
    FileEffects          []FileEffectIR
    Decisions            []DecisionIR
    Blockers             []BlockerIR
    Validations          []ValidationIR
    Artifacts            []ArtifactIR
    RiskFlags            []RiskFlagIR
}

type ProviderRunIR struct {
    RunID                string
    Capability           ProviderCapability  // 그 run 시작 시점의 capability snapshot
    StartedAt            time.Time
    EndedAt              *time.Time
    Outcome              string              // "completed" | "failed" | "cancelled" | "timed-out"
}
```

규칙:

1. `TaskIR` 은 **derived state** 다. 디스크에 캐시되더라도 truth-of-record 가 아니다. truth-of-record 는 이벤트 로그.
2. `FSMSchemaVersion` 과 `TaskIRSchemaVersion` 은 함께 진화할 수도, 독립일 수도 있다(예: FSM 변화 없이 view 필드만 추가).
3. reducer 가 미지의 `(Type, EventSchemaVersion)` 을 만나면 view 를 “부분 적용” 하지 않는다 — 그 task 의 reducer 가 통째로 fail 하고 `Blocked(reason=IR_REDUCER_INCOMPATIBLE)` 이벤트가 새로 append 된다. silent partial reduce 금지.
4. `Runs[*].Capability` 는 그 run 동안 **불변** 이다. capability 가 바뀌면 새 run 이 시작되거나 task 가 `Blocked` 로 빠진다(자세한 정책은 [`runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) 가 소유).

## 4. reducer 보존 규칙

reducer 는 `(EventType, EventSchemaVersion) → state_delta` 로 dispatch 한다.

- 새 케이스 추가는 항상 허용(additive).
- 기존 케이스의 동작 변경은 새 `EventSchemaVersion` 을 만들 때만 허용. 같은 페어의 옛 코드는 영구 보존.
- reducer 코드 파일 자체는 한 번 머지되면 그 분기에 대해서는 비변경(`go:build` 태그로 디렉토리 분리). 새 dispatch 는 새 파일에 추가.
- 새 `TaskIRSchemaVersion` 출시 시 옛 view 를 자동으로 “rebuild” 하는 마이그레이션 도구가 함께 머지되어야 한다(`riido-cli ir rebuild --to v2`). 옛 view 는 별도 column/테이블에 보관한다.

## 5. 새 EventSchemaVersion 도입 절차

PR 라벨 `change:breaking-ir` 가 강제하는 체크리스트:

1. 새 `(Type, N+1)` 케이스를 reducer 에 추가.
2. 옛 `(Type, N)` 케이스 코드 변경 없음을 lint 로 보장.
3. fixture 회귀: 옛 task 의 이벤트 로그 캡처를 그대로 흘려도 옛 view 의 동등성이 유지되는지 테스트.
4. 정적 capability 매트릭스 갱신(필요 시).
5. compatibility gate 에 “이 task 의 IR 은 reducer ≥ v? 가 필요하다” 라는 요구사항이 추가.
6. ADR 또는 기존 ADR 갱신(결정의 SSOT 가 흔들리는 경우).

## 6. 어댑터 ACL 과의 계약

어댑터는 provider raw 표현을 `CanonicalEvent` 로 변환할 때 다음을 지킨다(이 규칙은 본 문서가 소유).

1. raw payload 의 **알려지지 않은 키** 는 드롭하지 않고 `Unknown` 으로 보존.
2. raw event 의 **알려지지 않은 type** 은 `Type=ProviderUnknownEvent`, `Payload={rawType: ..., rawPayload: ...}` 로 변환한다. 이 이벤트는 FSM 전이를 일으키지 않는다(로그/감사용).
3. 어댑터가 raw 표현을 “해석” 해서 도메인 이벤트로 바꿀 때, 정보 손실이 생기면 그 이벤트와 같은 `CanonicalEvent` 에 `Payload.derived=true` 를 표기한다.

이 규칙은 [`provider-capability.md`](./provider-capability.md) 의 “unknown 보존” 정책과 일관된다.

## 7. 인접 SSOT 와의 계약

| 인접 SSOT | 이 문서가 요구/제공하는 것 |
| --- | --- |
| [`ir-event-log.md`](./ir-event-log.md) | `EventType` 카탈로그가 본 문서의 의무 필드 셋을 채워야 함 |
| [`task-lifecycle.md`](./task-lifecycle.md) | FSM 전이는 `(Type, EventSchemaVersion)` 페어에 의해서만 일어남. unknown event 는 전이 없음 |
| [`provider-capability.md`](./provider-capability.md) | `Runs[*].Capability` 의 정의 |
| [`runtime-versioning.md`](./runtime-versioning.md) | A/B/C/D/E/F/G 축 중 C 축이 본 문서에 위임됨 |
| [`../30-architecture/compatibility-gate.md`](../30-architecture/compatibility-gate.md) | reducer 호환성 게이트가 본 문서의 §4 보존 규칙을 인용 |

## 8. 미결정/오픈 이슈

`open-questions.md` 위임.

- `Q-IR-001`: 이벤트 로그 물리 저장(append-only 파일 vs DB table) 결정.
- `Q-IR-002`: `Payload.derived=true` 의 더 강한 표기(별도 컬럼 vs 페이로드 키) 선택.
- `Q-IR-006`: replay 성능을 위한 snapshot 주기/저장 위치.

## 9. version-affecting changes

이 문서가 정의하는 “의무 필드 셋” 의 추가는 **breaking-ir**(필드 추가만 해도 옛 데몬이 그 필드를 채우지 않은 채 이벤트를 발행하면 정합성이 깨짐). 따라서 의무 필드 변경 PR 은 항상 `change:breaking-ir` + 마이그레이션 도구 + fixture 회귀가 함께 머지된다.
