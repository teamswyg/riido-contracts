# Task Lifecycle (FSM) SSOT

> **이 문서가 `TaskState` 집합 / 합법 전이 / terminal 정의 / FSM invariant 의 SSOT다.**
>
> - 책임: 어떤 상태가 존재하는가, 어떤 transition 이 합법인가, 무엇이 terminal 인가, 어떤 invariant 가 모든 전이에 강제되는가
> - 비책임: `EventType` 의 카탈로그, 각 event 의 payload 스키마, reducer 의 dispatch / 보존 규칙, “어떤 event 가 transition event 인가” 의 정식 분류 — 이 모두는 [`./ir-event-log.md`](./ir-event-log.md) 가 소유한다. 본 문서는 transition trigger 의 **이름** 만 인용하며, 그 이름의 정식 정의는 IR Event Log 의 권한이다.

이 SSOT 는 [context-map.md](./context-map.md) 의 **C1 Task Lifecycle** context 를 채운다. C1 ↔ C2 경계는 §9 가 한 번 더 못박는다.

## 1. 책임 한 줄

> Task 의 상태는 직접 set 되지 않는다. 상태는 항상 **append-only IR 이벤트** 의 reducer 결과로 도출되며, 그 reducer 가 따르는 **합법 전이 표** 와 **invariant** 가 본 문서의 단독 소유다.

## 2. 상태 집합 (15)

현재 transition matrix 의 `FSMSchemaVersion` 은 **1** 이다. transition event 를 `CanonicalEvent` 로 append 할 때는 이 값을 `FSMVersion` 에 기록한다.

상태는 PascalCase 식별자다. 코드/이벤트/문서에서 동일 표기를 쓴다.

| ID | State | 한 줄 의미 | 분류 |
| --- | --- | --- | --- |
| S0 | `Created` | task row 가 생성됐고 아직 큐에 들어가지 않은 상태 (정의 검증 / 초기 메타 채움) | initial |
| S1 | `Queued` | 큐에 들어가 lease 대기 중 | waiting |
| S2 | `Claimed` | 한 runtime 이 lease 를 잡았으나 아직 workdir / process 준비 전 | progressing |
| S3 | `Preparing` | workdir 생성 / native config 주입 / secret scope 확보 중 (provider process 아직 미기동) | progressing |
| S4 | `Running` | provider process 가 기동되어 작업 수행 중 | active |
| S5 | `NeedsInput` | provider 가 사용자/외부에 질문을 던졌고 응답 대기 중. **task 자체는 정상.** | waiting (active subset) |
| S6 | `Blocked` | capability / policy / workspace / security / validation 등 **외부 조건** 으로 진행 불가. agent error 가 아니라 환경 조건. | waiting (recoverable) |
| S7 | `Validating` | provider 가 “완료” 를 신고했지만 daemon 측 게이트가 검증 중. **agent 자기보고는 완료가 아니다.** | active (daemon-side only) |
| S8 | `PatchReady` | daemon validation 이 통과해 산출물 / diff / 결과가 준비됨. 외부 승인이 필요한 경우 다음 단계로 넘어감. | waiting (review-ready) |
| S9 | `HumanReview` | 사람의 승인/반려가 필요. PatchReady 가 자동 승인 정책이면 이 상태를 건너뛸 수도 있다. | waiting (human) |
| S10 | `ReworkQueued` | review 가 반려되거나 자동 재시도 정책으로 새 run 이 필요. **새 `RunID` 가 발급된다.** | waiting (re-entry) |
| S11 | `Completed` | 모든 게이트 통과. **terminal.** | terminal |
| S12 | `Failed` | 회복 불가 실패. agent error / validation 실패 / runtime pin 위반 / timeout 등. **terminal.** | terminal |
| S13 | `Cancelled` | 사용자 / 운영자가 명시적 취소. **terminal.** | terminal |
| S14 | `TimedOut` | 시간 상한 초과. `Running` / `Validating` / `HumanReview` 등에서 진입 가능. **terminal.** | terminal |

분류 의미:

- **initial**: task 의 시작점. 정확히 하나(`Created`).
- **progressing**: 시간/작업이 흐르고 있는 비-terminal 상태.
- **active**: provider process 또는 daemon 검증기가 작동 중인 상태. lease pinning 이 강제됨.
- **waiting**: 외부 신호(lease / 사용자 응답 / 외부 조건 해소 / 사람 승인 / 새 run) 가 들어와야 진행됨.
- **terminal**: 더 이상 전이가 일어나지 않는 종착 상태. 정확히 네 개(`Completed` / `Failed` / `Cancelled` / `TimedOut`).

### 2.1 EventScope 와의 매핑

각 state 의 transition event 가 갖는 `EventScope` ([`./ir-schema-versioning.md`](./ir-schema-versioning.md) §1.5):

| 상태 / 전이 시점 | EventScope |
| --- | --- |
| `Created`, `Queued` 진입 | **TaskScope** (아직 어느 runtime 도 잡지 않음 — `RuntimeID`/`CapabilityFingerprint` 부재) |
| `Claimed` 진입부터 모든 이후 transition (`Preparing`/`Running`/`NeedsInput`/`Validating`/`PatchReady`/`HumanReview`/terminal) | **RunScope** (claim 시점에 `RuntimeID` + `CapabilityFingerprint` 결정 — `TaskClaimed` 가 두 식별자를 박는다) |
| `Cancelled` (Queued 에서 취소) | TaskScope |
| `Cancelled` (Running 등에서 취소) | RunScope |
| `Failed` / `Completed` / `TimedOut` | 항상 RunScope (terminal 진입은 항상 run 이 시작된 후이거나 claim 후) |

ingest 는 transition event 의 `Scope` 를 본 표에 따라 박는다. `RuntimeID`/`CapabilityFingerprint` 가 부재인 TaskScope 단계에서 `Running` 으로 직접 점프하는 transition 은 §5 invariant 2 가 차단한다 — `Running` 진입의 전이 트리거 `RunStarted` 는 항상 RunScope 다.

## 3. 합법 전이 표

표 안의 셀: 합법이면 ✓, 불법이면 (빈칸). 행=from, 열=to.

```
                Cr Qu Cm Pr Ru NI Bl Va PR HR RW Co Fa Ca TO
Created (Cr)       ✓                                ✓  ✓
Queued  (Qu)          ✓                          ✓  ✓
Claimed (Cm)             ✓        ✓                 ✓  ✓
Preparing(Pr)               ✓     ✓                 ✓  ✓
Running (Ru)                   ✓  ✓  ✓              ✓  ✓  ✓
NeedsInp(NI)                ✓     ✓                 ✓  ✓  ✓
Blocked (Bl)          ✓        ✓                    ✓  ✓
Validating(Va)                       ✓              ✓  ✓  ✓
PatchRdy(PR)                            ✓        ✓  ✓     ✓
HumanRev(HR)                               ✓     ✓        ✓
ReworkQ (RW)          ✓                                ✓
(terminal: Co/Fa/Ca/TO — 더 이상 전이 없음)
```

읽는 법: 예를 들어 `Running → NeedsInput → Running` 은 합법이지만 `Running → PatchReady` 는 **불법**(반드시 `Validating` 을 거쳐야 함). `Created → Running` 도 불법(반드시 `Queued → Claimed → Preparing` 순).

### 3.1 자주 도는 경로

| 경로 | 의미 |
| --- | --- |
| `Created → Queued → Claimed → Preparing → Running → Validating → PatchReady → HumanReview → Completed` | 표준 happy path (사람 승인 포함) |
| `Created → Queued → Claimed → Preparing → Running → Validating → PatchReady → Completed` | auto-approve 정책이 활성일 때 happy path |
| `Running → NeedsInput → Running` | provider 질문 → 응답 후 복귀 |
| `Running → Blocked → Running` | 외부 조건 해소 후 복귀 |
| `Blocked → Queued` | 외부 조건 해소 + 새 runtime 으로 재배치 필요 시 |
| `HumanReview → ReworkQueued → Queued` | 반려 → 새 RunID 로 재진행 |
| `Validating → Failed` | daemon validation 실패 |
| `Running → Failed (RuntimePinViolated)` | runtime fingerprint 가 도중에 깨짐 |
| `* → Cancelled` | 사용자/운영자 명시적 취소 (대부분 비-terminal 상태에서 가능) |
| `Running / Validating / HumanReview → TimedOut` | 시간 상한 초과 |

### 3.2 `Cancelled` 의 진입 가능 범위

`Cancelled` 는 다음 비-terminal 상태에서만 진입할 수 있다: `Created`, `Queued`, `Claimed`, `Preparing`, `Running`, `NeedsInput`, `Blocked`, `Validating`, `PatchReady`, `HumanReview`, `ReworkQueued`. terminal 4 개로부터는 진입 불가.

### 3.3 `TimedOut` 의 진입 가능 범위

`TimedOut` 은 시간 상한이 의미 있는 상태에서만 진입 가능: `Running`, `NeedsInput`, `Blocked`, `Validating`, `HumanReview`. `Claimed` 의 시간 상한은 별도 정책으로 `Failed` 로 분류한다. `Preparing` 중 runtime lease 가 만료되면 C5 는 회복 가능한 handoff 로 보고 `Preparing → Blocked → Queued` 경로를 쓴다.

## 4. transition trigger 참조 표

각 transition 은 **transition event** 에 의해서만 일어난다. 아래 이름들은 본 문서가 인용하는 **참조 이름**이고, 정식 카탈로그(payload / 의무 필드 / reducer)는 [`ir-event-log.md`](./ir-event-log.md) 가 소유한다.

| Transition | 참조 event 이름 (인용) | 1차 producer |
| --- | --- | --- |
| `Created → Queued` | `TaskQueued` | 서버/API |
| `Queued → Claimed` | `TaskClaimed` | scheduler (C5) |
| `Claimed → Preparing` | `WorkdirPreparing` | runtime (C4) |
| `Preparing → Running` | `RuntimePinned` + `RunStarted` (둘이 한 쌍) | runtime (C4) |
| `Running → NeedsInput` | `InputRequested` | runtime (C4) |
| `NeedsInput → Running` | `InputProvided` | 서버/API |
| `Running → Blocked` | `BlockerRaised(category=…)` | runtime / scheduler |
| `Blocked → Running` | `BlockerResolved` | scheduler |
| `Blocked → Queued` | `BlockerResolvedRequeue` | scheduler (다른 runtime 필요 시) |
| `Running → Validating` | `RunReportedDone` | runtime (C4) |
| `Validating → PatchReady` | `ValidationPassed` | validation (C8) |
| `Validating → Failed` | `ValidationFailed` | validation (C8) |
| `PatchReady → HumanReview` | `ReviewRequested` | policy/validation |
| `PatchReady → Completed` | `AutoApproved` | policy |
| `HumanReview → Completed` | `HumanApproved` | 서버/API |
| `HumanReview → ReworkQueued` | `HumanRejected(rework=true)` | 서버/API |
| `ReworkQueued → Queued` | `ReworkAccepted` (**새 RunID 발급**) | 서버 |
| `* → Cancelled` | `TaskCancelled(reason=…)` | 서버/API |
| `* (active subset) → TimedOut` | `TaskTimedOut(state=…)` | scheduler |
| `Running / Validating → Failed (pin 위반)` | `RuntimePinViolated` | runtime (C4) |

위 표의 “이름” 은 IR Event Log SSOT 의 카탈로그와 1:1 매칭되어야 한다. 본 문서가 이 표를 갱신할 때는 IR Event Log SSOT 의 동일 PR 갱신이 강제된다(§11 참조).

## 5. Invariants

다음 8 개 invariant 는 모든 transition 에 강제된다.

1. **Task state 는 직접 set 되지 않는다.** transition event 를 IR 로그에 append 한 뒤 reducer 가 새 state 를 도출하는 길만 합법이다. SQL `UPDATE tasks SET state=...` 같은 직접 변이는 본 SSOT 위반.
2. **`Running` 진입 전에 `(RuntimeID, CapabilityFingerprint)` pinning 이 완료되어야 한다.** `Preparing → Running` 의 트리거는 `RuntimePinned` 와 `RunStarted` 가 한 쌍이며, `RuntimePinned` 가 먼저 append 되어 있지 않으면 reducer 가 거절한다.
3. **`Running` / `Validating` 중 `(RuntimeID, CapabilityFingerprint)` 가 바뀌면 `RuntimePinViolated` → `Failed`.** silent recovery 금지. (`runtime-upgrade-flow.md` §3 와 일관.)
4. **`Completed` 진입은 두 입력 중 하나가 필요하다**: (a) `ValidationPassed` + `AutoApproved` 또는 (b) `ValidationPassed` + `HumanApproved`. **agent 자기보고만으로 `Completed` 진입 불가.**
5. **terminal 4 개(`Completed` / `Failed` / `Cancelled` / `TimedOut`)는 출발지 없음.** 어떤 transition event 도 terminal 에서 발생할 수 없다(reducer 가 거절).
6. **`ReworkQueued → Queued` 는 새 `RunID` 를 발급한다.** 같은 task 에서 옛 `RunID` 의 이벤트들은 보존되지만 새 run 은 새 `RunID` 로 시작한다. 새 RunID 가 없는 `ReworkQueued → Queued` 전이는 invariant 위반.
7. **`NeedsInput` 는 “질문” 이지 “실패” 가 아니다.** 어떤 reducer 코드도 `NeedsInput` 진입을 `Failed` 의 precursor 로 간주해서는 안 된다. `NeedsInput` 의 timeout 은 `TimedOut` 으로만 분기한다.
8. **`Blocked` 는 “외부 조건” 이지 “agent error” 가 아니다.** category 가 `CAPABILITY_*` / `POLICY_*` / `WORKSPACE_*` / `SECURITY_*` / `VALIDATION_INPUT_MISSING` 같은 외부 환경 사유여야 한다. agent 가 코드를 깨뜨린 경우는 `Validating → Failed` 또는 `Running → Failed` 로 흘러야 한다.
9. **`Running` 진입의 사전조건은 `WorkspacePrepared` + `NativeConfigVersion` 확정**. `Claim` 의 사전조건이 아니다. C5 scheduler 는 task 를 claim 할 때 workspace feasibility 만 볼 수 있고, `WorkspacePrepared` state 자체를 claim 전에 요구하지 않는다. `Preparing → Running` 전이 트리거 (`RunStarted`) 는 FSM Orchestrator / Pre-Execute Gate (G5 + G-S2) 가 다음 모두를 확인한 후에만 발행된다: workspace prepared, native config injected, runtime pinned, security pre-execute gate 통과. (자세한 게이트 정의는 [`../30-architecture/compatibility-gate.md`](../30-architecture/compatibility-gate.md) G5, [`./security.md`](./security.md) §4 G-S2.)

## 6. terminal 의 의미

| Terminal | 의미 | 사후 행동 |
| --- | --- | --- |
| `Completed` | 모든 게이트 통과 + 산출물 보존 | 산출물은 archived; lease 반환 |
| `Failed` | 회복 불가 실패 (agent / validation / runtime 위반 / preparing 실패 등) | 운영자 분석 대상; replay 가능 |
| `Cancelled` | 명시적 취소 | 산출물 일부만 보존 (정책 따라 결정 — #15 security/policy) |
| `TimedOut` | 시간 상한 | 정책에 따라 자동 ReworkQueued 진입은 **불가** (terminal 이므로). 새 task 를 만들어야 한다. |

terminal 도달 후의 “재시도” 는 새 `TaskID` 또는 `ReworkQueued` 비-terminal 경로로만 가능하다.

## 7. ReworkQueued 의 RunID 규칙

`ReworkQueued → Queued` 가 일어날 때 다음을 수행한다.

1. 새 `RunID` 발급 (정책: ulid).
2. 옛 `RunID` 의 IR 이벤트들은 **보존** (append-only invariant).
3. 새 run 의 첫 이벤트는 `ReworkAccepted(prevRunID=...)` 이며 이후 모든 이벤트의 `RunID` 가 새 값.
4. `TaskIR.Runs[]` 에는 옛 run 과 새 run 모두 보존(`ProviderRunIR` 한 행씩). 옛 run 의 `Outcome` 은 “rework” 로 표기.

## 8. NeedsInput vs Blocked vs Failed 분리

세 상태가 자주 혼동되므로 본 문서가 한 번 더 못박는다.

| 상황 | 정답 상태 |
| --- | --- |
| provider 가 사용자/외부에 질문을 던졌고 응답이 오면 진행 가능 | `NeedsInput` |
| capability 매트릭스가 task requirement 를 만족 못 함 | `Blocked(category=CAPABILITY_MISSING)` |
| 정책 번들이 새 메이저로 바뀌어 재검사 필요 | `Blocked(category=POLICY_REEVAL)` |
| workdir 생성 실패 / git checkout 실패 | `Blocked(category=WORKSPACE_PREP_FAILED)` 또는 (회복 불가면) `Failed` |
| sandbox 위반 / protected path 손상 | `Blocked(category=SECURITY_VIOLATION)` 또는 (의도적/심각) `Failed` |
| validation rule 의 input 누락 (예: 필요한 fixture 가 없음) | `Blocked(category=VALIDATION_INPUT_MISSING)` |
| validation rule 실행 결과 실패 | `Validating → Failed` (terminal) |
| agent 가 코드를 컴파일 깨뜨림 | `Validating → Failed` |
| runtime pinning 위반 | `Running/Validating → Failed (RuntimePinViolated)` |

핵심: **회복 가능 = `Blocked` / `NeedsInput`. 회복 불가 = `Failed`.**

## 9. C1 ↔ C2 경계 (재확인)

| 책임 | 소유 |
| --- | --- |
| `TaskState` enum (15 멤버) | **C1 (본 문서)** |
| 합법 transition matrix | **C1 (본 문서)** |
| terminal 정의 | **C1 (본 문서)** |
| transition invariant (§5) | **C1 (본 문서)** |
| `EventType` 카탈로그 (모든 이벤트의 이름·payload·종류) | **C2** ([`./ir-event-log.md`](./ir-event-log.md)) |
| “어떤 EventType 이 transition event 인가” 의 정식 분류 | **C2** |
| `CanonicalEvent.Payload` 스키마 | **C2** |
| reducer dispatch / 보존 규칙 | **C2** |
| unknown raw 필드 보존 | **C2** |

본 문서 §4 의 “참조 event 이름” 은 C2 카탈로그와 정확히 매칭되어야 한다. 두 SSOT 가 어긋나면 PR 단계에서 거절한다.

## 10. 미결정 / 오픈 이슈

`open-questions.md` 위임.

- `Q-FSM-001`: 해결됨. `Claimed` 상태에서 runtime lease 가 만료되면 `Failed`(시간 상한이 아니라 “준비 단계 실패”)로 분류한다. `Preparing` / `Running` 중 runtime lease 만료는 C5 handoff 정책에 따라 `Blocked → Queued` 로 재배치할 수 있다.
- `Q-FSM-002`: auto-approve 정책이 활성일 때도 `HumanReview` 를 “건너뛰지 않고 자동 통과” 로 기록할지(audit trail 강화).
- `Q-FSM-003`: `Blocked → Queued` 시 lease 가 다른 runtime 으로 넘어갈 수 있는데, 같은 RunID 를 유지할지 새로 발급할지. 현재안: 같은 RunID(왜냐하면 run 이 “시작” 되지는 않은 상태이므로).
- `Q-FSM-004`: `NeedsInput` 에 머무는 시간 상한의 기본값.

## 11. version-affecting changes

본 문서는 **C1** 의 SSOT 이므로 다음 변경은 강한 라벨을 동반한다.

- 새 state 추가는 `change:additive` 단, transition 표와 §4 reference event 표가 함께 갱신되어야 함.
- 기존 state 의 의미 변경 또는 제거는 `change:breaking-policy` + `FSMSchemaVersion` 증가 ([`ir-schema-versioning.md`](./ir-schema-versioning.md) §1 참조).
- transition trigger 이름 변경은 **C2 SSOT (`ir-event-log.md`) 의 동시 PR 갱신** 이 강제된다. 한 쪽만 바뀌면 reducer 가 사라진 EventType 으로 dispatch 하다 실패한다.
