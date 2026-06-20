# Provider Capability SSOT

> **이 문서가 `ProviderCapability` 모델과 protocol-별 어댑터 분류의 SSOT다.**
>
> - 책임: provider CLI 가 “무엇을 할 수 있는지” 의 도메인 표현, capability detection 알고리즘, protocol → 어댑터 매핑, `CompatibilityStatus` 등급의 정의
> - 비책임: Claude Code / Codex CLI 자체의 옵션 카탈로그는 [`../10-research/claude-code.md`](../10-research/claude-code.md), [`../10-research/codex.md`](../10-research/codex.md) 가 소유. capability 게이트가 “언제” 동작하는지는 [`../30-architecture/compatibility-gate.md`](../30-architecture/compatibility-gate.md) 가 소유.

이 SSOT 는 [ADR 0001 — 버저닝은 capability 기반](../40-decisions/0001-versioning-is-capability-based.md) 의 직접 구현이다.

## 0. 본 문서가 박는 핵심 invariant (요약)

다음 다섯 가지는 capability 도메인에서 단단히 박혀야 한다. 다른 절들은 이 다섯의 정교한 표현일 뿐이다.

1. **`ProtocolKind` 가 adapter 선택의 1차 key 다.** `ProviderKind` 가 “claude” 라는 사실만으로 어댑터를 고르지 않는다. wrapper / fork / experimental surface 가 같은 ProviderKind 아래 공존하기 때문이다.
2. **provider binary 의 `DetectedVersion` 은 raw signal 일 뿐 실행 조건이 아니다.** 분기는 항상 surface flag 와 `ProtocolKind`/`ProtocolVersion` 으로 한다. 버전 숫자 비교(`>=`/`<`)로 분기하는 코드는 PR 단계에서 거절된다.
3. **`CompatibilityStatus` 는 네 입력의 결합 *요약* 일 뿐 단일 진실이 아니다.** 입력 (a) capability detection (b) protocol maturity (c) policy bundle 호환 (d) 어댑터 회귀 테스트 — 네 입력은 보수적 AND 로 단일 enum 으로 요약되지만, scheduler 가 정확한 게이트 결정을 내리려면 **보조 필드**(`ProtocolMaturity`, `RequiresExperimentalOptIn`, `MissingCapabilities`, `BlockedReasons`, `DegradedReasons`)를 함께 읽어야 한다. enum 한 값으로 정보가 손실되면 안 된다. 등급 우선순위는 **blocked > experimental > degraded > supported**.
4. **`ExposesUnsafePermissionBypass=true` 는 사용 허가가 아니라 risk signal 이다.** Claude 의 `--permission-mode bypassPermissions`, Codex 의 `--yolo` / `--dangerously-bypass-approvals-and-sandbox` 처럼 approval/tool-use 우회 표면이 노출되어 있음을 의미할 뿐, **사용 가부의 최종 결정** 은 downstream C7 Security / Policy 와 C4 Provider Runtime harness gate 가 내린다. 어댑터가 이 flag 만 보고 우회 모드를 켜는 코드는 PR 단계에서 거절된다. Codex `danger-full-access` 처럼 sandbox enum 에 포함될 수 있는 값은 provider 가 노출한 sandbox surface facts 로 기록할 수 있지만, 그것만으로 Riido 가 해당 실행 envelope 를 채택했다는 뜻은 아니다.
5. **event stream 형식 이름은 provider-neutral 로 둔다.** Claude 의 `stream-json` 과 Codex 의 `exec --json` 은 같은 NDJSON 형식이지만 명칭이 다르다. 도메인 capability 는 `SupportsStructuredEventStream` + `EventStreamFormat` enum 으로 표현하고, provider-specific 호출 플래그명은 `ClaudeSurface` / `CodexSurface` (아래 §8)에 따로 둔다.

위 다섯 invariant 는 [ADR 0001](../40-decisions/0001-versioning-is-capability-based.md) 결정의 capability 도메인 표현이다. ADR 의 7 약속 전체는 [`runtime-versioning.md`](./runtime-versioning.md) §0 가 소유한다.

## 1. 왜 capability 가 1차 시민인가

`if claudeVersion >= "X"` 분기는 다음 세 경우 모두 깨진다.

1. minor 업그레이드가 인자/이벤트 스키마/sandbox 기본값을 조용히 바꿈.
2. 같은 이름의 CLI 가 wrapper / fork / 사내 gateway 형태로 동시 존재.
3. 새 기능이 fingerprint version 보다 빨리 들어옴(예: `--include-partial-messages` 같은 신규 플래그).

Riido 는 fingerprint version 을 **저장은 하지만 분기 조건으로 쓰지 않는다**. 분기 조건은 항상 `ProviderCapability` 다.

## 2. ProviderCapability 모델

도메인 타입(Go 식별자 기준, 실제 패키지는 `provider/capability` 가 소유).

```go
type ProviderCapability struct {
    // identity
    RuntimeID             string  // 등록된 runtime 슬롯의 stable ID — IR 이벤트의 RuntimeID 와 동일
    ProviderKind          string  // "claude" | "codex" | "claude-wrapper" | ...
    ProtocolKind          string  // §4 참조
    AdapterID             string  // 이 capability 를 만든 어댑터의 등록명
    AdapterVersion        string  // 어댑터 코드 자체의 semver
    ProtocolVersion       string  // 어댑터가 인지한 protocol surface 식별자

    // discovery
    ExecutablePath        string
    Argv0                 string
    DetectedVersion       string  // fingerprint only — 분기 조건 아님
    DetectedFingerprint   string  // checksum(binary) + CLI banner hash — binary 자체의 지문
    CapabilityFingerprint string  // 이 capability snapshot 의 결정적 해시 (§2.1)
    DiscoveredAt          time.Time

    // surface flags — provider-neutral (의미 단위, "버전 ≥ X" 분기 금지)
    // -- event stream surface
    SupportsStructuredEventStream bool              // NDJSON / JSON-RPC notifications 중 하나라도 지원
    EventStreamFormat             EventStreamFormat // "ndjson" | "json-rpc-notifications" | "text-only" | "unknown"
    SupportsPartialDeltas         bool              // 토큰/문장 단위 부분 델타 이벤트
    // -- 세션/재개 표면
    SupportsResume                bool
    SupportsSessionID             bool   // 외부에서 session/thread id 를 부여/고정 가능
    SupportsSessionPin            bool   // pin 의 강도: 단일 인스턴스가 점유함을 데몬이 강제할 수 있는가
    SupportsSystemPrompt          bool
    SupportsMaxTurns              bool
    // -- tool / file 이벤트
    SupportsToolEvents            bool
    SupportsFileEvents            bool
    SupportsUsageMetrics          bool
    // -- 권한/보안 표면 (provider-neutral)
    SupportsPermissionControl     bool   // 권한/승인 primitive 존재 (Claude --permission-mode, Codex approval policy 등)
    ExposesUnsafePermissionBypass bool   // ※ risk signal — 사용 허가 아님 (§0 invariant 4)
    SupportsApprovalProtocol      bool   // 양방향 approval 프로토콜 (codex app-server 등)
    SupportsSandbox               bool
    SupportsManagedSettings       bool   // 외부 강제 설정 채널
    // -- 확장 표면
    SupportsHookEvents            bool   // PreToolUse / PostToolUse / PermissionRequest / SessionStart 등
    SupportsMCP                   bool
    SupportsWorktree              bool
    SupportsJSONSchemaTools       bool

    // safety surface defaults (observed/profile-derived facts, not Riido launch decisions)
    DefaultSandboxMode       string  // "read-only" | "workspace-write" | "danger-full-access" | "unknown"
    DefaultApprovalPolicy    string  // "on-request" | "on-failure" | "never" | "untrusted" | "unknown"
    HasNetworkOffDefault     bool

    // compatibility envelope (rich gate signals — §5)
    CompatibilityStatus       CompatibilityStatus    // summary (blocked > experimental > degraded > supported)
    ProtocolMaturity          ProtocolMaturity       // "stable" | "experimental" | "deprecated" | "unknown"
    RequiresExperimentalOptIn bool                   // task.allowExperimentalRuntime=true 가 필요한가
    MissingCapabilities       []CapabilityName       // probe 또는 매트릭스가 “부재” 로 결론낸 capability
    BlockedReasons            []CompatibilityReason  // 비어있지 않으면 status=blocked
    DegradedReasons           []CompatibilityReason  // 비어있지 않으면 status≥degraded
    MinSupportedVersion       string
    MaxTestedVersion          string

    // provider-specific surface (해당 provider 한 곳만 채워짐; 다른 곳은 nil)
    ClaudeSurface *ClaudeSurface
    CodexSurface  *CodexSurface

    // raw bag — 알려지지 않은 surface 보존
    Unknown                  map[string]any
}

type EventStreamFormat string

const (
    EventStreamFormatUnknown              EventStreamFormat = "unknown"
    EventStreamFormatTextOnly             EventStreamFormat = "text-only"
    EventStreamFormatNDJSON               EventStreamFormat = "ndjson"                  // claude --output-format stream-json, codex exec --json
    EventStreamFormatJSONRPCNotifications EventStreamFormat = "json-rpc-notifications"  // codex app-server
)

type ProtocolMaturity string

const (
    ProtocolMaturityUnknown      ProtocolMaturity = "unknown"
    ProtocolMaturityStable       ProtocolMaturity = "stable"
    ProtocolMaturityExperimental ProtocolMaturity = "experimental"
    ProtocolMaturityDeprecated   ProtocolMaturity = "deprecated"
)

type CapabilityName string  // e.g. "structured-event-stream", "session-resume", "approval-protocol"

type CompatibilityReason struct {
    Code    string  // "PROBE_FAILED" | "POLICY_INCOMPATIBLE" | "ADAPTER_REGRESSION_FAILED" | "MIN_VERSION" | ...
    Subject string  // 문제가 된 capability/policy 식별자
    Detail  string  // 사람이 읽는 설명
}
```

규칙(이 문서가 소유):

1. `ProviderCapability` 는 **불변값** 이다. detection 마다 새 인스턴스가 생성되며, 같은 runtime row 의 `CapabilityFingerprint` 가 바뀌면 lease 가 무효화된다.
2. 비교/분기는 **불리언 surface flag** 와 `ProtocolKind`/`ProtocolVersion` 으로만 한다. `DetectedVersion` 으로 분기하는 코드는 PR 단계에서 거절된다.
3. `Unknown` 필드는 “알려지지 않은 surface 를 드롭하지 않는다” 원칙의 보존소다(이벤트 ACL 의 `payload.unknown[]` 과 같은 정신).
4. runtime 식별자는 `RuntimeID` 하나로 통일. `ProviderID` 같은 모호한 이름은 금지어다. provider 의 종류는 `ProviderKind`, protocol 의 종류는 `ProtocolKind` — 셋의 의미를 한 식별자로 묶지 않는다.
5. `DefaultSandboxMode` / `DefaultApprovalPolicy` 는 detection 또는 static profile 에서 관찰한 provider safety surface fact 다. Scheduler / adapter 는 이 값을 “Riido 가 실제로 그 mode 로 실행한다”는 결정으로 해석하면 안 된다. Downstream C4/C7 이 provider 별 trusted-runtime envelope 를 명시 채택하고 harness evidence 를 추가해야 실제 실행 정책이 된다.

### 2.1 CapabilityFingerprint 계산

> `CapabilityFingerprint` 는 **runtime capability snapshot** 의 fingerprint 다 — “이 runtime 이 무엇을 할 수 있는가” 를 잡는다. **execution context** (어느 task 의 어느 run 이 어느 native config 위에서 도는가) 는 별도 축이며, 본 fingerprint 에 포함하지 않는다. 그 이유:
>
> - `TaskClaimed` 시점에는 아직 workspace 가 준비되지 않아 `NativeConfigVersion` 이 존재하지 않는다. 그런데 `TaskClaimed` 는 RunScope 이고 `CapabilityFingerprint` 를 요구한다. fingerprint 가 NCV 를 입력으로 가지면 claim 자체가 불가능해진다(`#27i` 라운드의 발견).
> - runtime lease / heartbeat / re-detection trigger 는 “runtime 이 같은가” 만 판단하면 된다. native config 는 task 별로 다르고 lease 단위로 변하므로 runtime fingerprint 와 묶일 일이 없다.
>
> 따라서 **runtime pinning** 의 정확한 페어는 `(RuntimeID, CapabilityFingerprint)` 이지만, **실행 맥락 pinning** 은 추가로 `(PolicyBundleVersion, NativeConfigVersion)` 을 동반한다 (`task-lifecycle.md` Invariant Anchors / `runtime-upgrade-flow.md` §3).

`CapabilityFingerprint` 는 다음 10 개 입력 + 1 surface flag 집합을 정렬된 JSON 으로 직렬화한 뒤 SHA-256 으로 해시한 hex 문자열이다. 입력 순서/직렬화 형식은 본 문서가 소유한다(어댑터 별 자유 구현 불가 — 같은 capability 가 다른 데몬에서도 같은 fingerprint 를 만들어야 lease handoff 가 동작한다).

| # | 입력 |
| --- | --- |
| 1 | `ProviderKind` |
| 2 | `ProtocolKind` |
| 3 | `ProviderVersion` (raw `DetectedFingerprint` 가 아니라 자기 신고 버전 문자열) |
| 4 | `DetectedFingerprint` |
| 5 | `AdapterID` |
| 6 | `AdapterVersion` |
| 7 | `ProtocolVersion` |
| 8 | `DefaultSandboxMode` |
| 9 | `DefaultApprovalPolicy` |
| 10 | `PolicyBundleVersion` (runtime eligibility 평가에 쓰인 정책 번들 버전; task 별 NCV 와 다름) |
| 11 | **important surface flags** — 본 SSOT 가 정의하는 “capability 결정에 영향을 주는 flag 집합”(아래 표) |

> **빠진 입력**: `NativeConfigVersion`. NCV 는 task 별 execution context 다. runtime capability 와 무관하며 본 fingerprint 의 입력이 아니다. 본 fingerprint 가 바뀌면 lease 가 무효화되지만, NCV 가 바뀐다고 lease 가 무효화되지는 않는다(다만 진행 중 task 의 NCV 가 도중에 바뀌면 [`../30-architecture/runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) §2 T-CONFIG 매트릭스가 발동).

“important surface flags” 의 멤버:

- `SupportsStructuredEventStream`, `EventStreamFormat`, `SupportsPartialDeltas`
- `SupportsResume`, `SupportsSessionID`, `SupportsSessionPin`, `SupportsSystemPrompt`, `SupportsMaxTurns`
- `SupportsToolEvents`, `SupportsFileEvents`, `SupportsUsageMetrics`
- `SupportsPermissionControl`, `ExposesUnsafePermissionBypass`, `SupportsApprovalProtocol`, `SupportsSandbox`, `SupportsManagedSettings`
- `SupportsHookEvents`, `SupportsMCP`, `SupportsWorktree`, `SupportsJSONSchemaTools`
- `ProtocolMaturity`, `CompatibilityStatus`, `RequiresExperimentalOptIn`

provider-specific surface(`ClaudeSurface`/`CodexSurface`) 의 필드는 본 해시 입력에 **포함하지 않는다** — 해시 안정성을 위해 vendor-only 필드는 별도 보조 fingerprint(`ProviderSurfaceFingerprint`) 로 따로 둘 수 있으나, 1차 게이트 결정에는 `CapabilityFingerprint` 만 쓴다.

### 2.2 세 fingerprint 의 안정성 비교

| 필드 | 안정성 | 무엇이 바뀌면 새 값 | 1차 용도 |
| --- | --- | --- | --- |
| `RuntimeID` | 가장 안정 | runtime 슬롯 재등록 | DB lease 의 owning runtime 식별 |
| `CapabilityFingerprint` | 중간 | capability 입력 10 + surface flag 집합 중 어느 하나라도 바뀜 (§2.1 — NCV 는 입력 아님) | lease 무효화 트리거, IR 이벤트 식별 |
| `DetectedFingerprint` | 가장 변동 | binary 파일 자체 변경 | capability 재탐지 트리거 |

흐름: `DetectedFingerprint` 변경 → capability 재탐지 → 새 `CapabilityFingerprint` 산출 → `RuntimeID` 의 lease 가 옛 fingerprint 와 다르면 무효화 → upgrade flow 매트릭스(§2 of `runtime-upgrade-flow.md`) 진입.

## 3. detection 알고리즘

정적 매트릭스 + 동적 probe 를 결합한다.

### 3.1 정적 매트릭스 (`provider/capability` 또는 daemon adapter ACL 이 소유)

```
(ProviderKind, DetectedVersionRange) → DefaultCapabilityProfile
```

- 어댑터 저장소 안에 코드로 보관한다.
- 각 cell 은 “이 버전 범위에서 알려진 surface 기본값”을 표현한다.
- 외부 사실 SSOT(`../10-research/claude-code.md`, `../10-research/codex.md`)가 갱신될 때 함께 PR 로 들어와야 한다.
- 정적 매트릭스가 우리가 신뢰하는 “lower bound” — probe 실패 시에도 이 값까지는 보장.

### 3.2 동적 probe

매트릭스 위에 다음 probe 를 얹어 surface flag 를 “verified” 로 끌어올린다.

| Probe | 목적 | 검증 대상 flag |
| --- | --- | --- |
| `--help` 파싱 | 옵션 존재 여부 | `SupportsStructuredEventStream`, `SupportsResume`, `SupportsHookEvents`, `SupportsMCP`, `SupportsWorktree`, `SupportsPermissionControl`, `ExposesUnsafePermissionBypass`, `SupportsManagedSettings` |
| `app-server initialize` 핸드셰이크 | JSON-RPC capabilities | `SupportsSessionPin`, `SupportsApprovalProtocol`, `EventStreamFormat=json-rpc-notifications`, `ProtocolVersion` |
| dry-run 한 NDJSON line 파싱 | 이벤트 키 셋 추론 | `EventStreamFormat=ndjson`, `SupportsPartialDeltas`, `SupportsToolEvents`, `SupportsFileEvents` |
| sandbox 기본값 헬프/매뉴얼 매칭 | `DefaultSandboxMode`, `DefaultApprovalPolicy`, `HasNetworkOffDefault` | 안전 정책 게이트용 |
| provider-specific surface fill-in | Claude / Codex 전용 호출 플래그 인지 | `ClaudeSurface.*` 또는 `CodexSurface.*` (§8) |

probe 결과는 `Unknown` 에 그대로 보존되고, surface flag 는 매트릭스 ↔ probe 가 **둘 다 OK 일 때만** `true` 로 승격된다(보수적 OR 가 아니라 보수적 AND).

### 3.3 재탐지(re-detection) 트리거

다음 중 하나라도 일어나면 capability 를 다시 계산한다.

- `claude --version` / `codex --version` 의 fingerprint 가 바뀜
- 바이너리 파일 mtime 또는 hash 가 바뀜
- 어댑터 자체가 새 minor 버전으로 올라옴(`AdapterVersion` 변경)
- 데몬 시작 / 주기적 헬스체크
- 정책 번들(`policy-bundle`) 이 갱신되어 sandbox/permission 게이트가 바뀜

재탐지가 일어나면 [compatibility gate](../30-architecture/compatibility-gate.md) 가 진행 중 task 와의 정합을 다시 평가하고, [runtime upgrade flow](../30-architecture/runtime-upgrade-flow.md) 가 task state 별 허용/거절을 결정한다.

## 4. protocol 별 어댑터 분류

어댑터는 **버전별로 분기하지 않는다**. protocol-level 다형성으로 다음 4 종을 둔다.

| Adapter | ProtocolKind | 표면 | 기본 `CompatibilityStatus` | 단계 |
| --- | --- | --- | --- | --- |
| `ClaudeStreamJSONAdapter` | `claude-stream-json` | `claude -p --output-format stream-json` (+ `--input-format stream-json` / `--include-partial-messages` when supported) | `supported` | **MVP stable** |
| `CodexExecJSONLAdapter` | `codex-exec-jsonl` | `codex exec --json` NDJSON 이벤트 스트림. 공식 CLI reference 가 **Stable** 로 분류 | `supported` | **MVP stable** |
| `CodexAppServerAdapter` | `codex-app-server` | `codex app-server --listen stdio://` JSON-RPC + thread/turn/item notification. 공식 CLI reference 가 **Experimental** 로 분류(“development/debugging 용도이며 변경될 수 있다”) | `experimental` | **Spike / Long-term** |
| `OpenClawAgentJSONAdapter` | `openclaw-agent-json` | `openclaw agent --local --json --session-id ...` JSON/NDJSON result stream | `experimental` | **MVP volatile** |
| `CursorAgentStreamJSONAdapter` | `cursor-agent-stream-json` | `cursor-agent -p ... --output-format stream-json --workspace ...` | `experimental` | **MVP volatile** |
| `ClaudeCompatibleWrapperAdapter` | `claude-compatible-wrapper` | stock 과 동일 인자 표면을 흉내내는 사내 gateway/wrapper CLI | wrapper 매니페스트 자기신고에 의존; 기본은 `experimental` | Future |
| (예시) `OpenCodeAdapter` | `opencode-*` | TBD | TBD | Future |

규칙:

1. 어댑터 식별자는 `ProtocolKind` 다. `ProviderKind=claude` 하나만 가지고 어댑터를 고르면 wrapper 변종을 흡수할 수 없다.
2. 같은 ProtocolKind 안에서 minor surface 차이는 `Unknown` 보존 + capability flag 로 풀고, 새 ProtocolKind 가 필요한 변화(예: stream-json 메이저 스키마 변경) 가 나오면 **새 어댑터를 추가**하고 옛 어댑터는 유지한다.
3. wrapper 어댑터는 정책 SSOT([`security.md`](./security.md)) 가 정의한 “신뢰 등급(trust tier)” 을 capability 에 기록해야 한다 — wrapper 가 sandbox 를 우회할 수 있는 외부 도구일 수 있기 때문.
4. **production path 는 stable surface 만**. `CodexAppServerAdapter` 가 더 풍부한 thread lifecycle 을 제공하지만, app-server 자체가 공식적으로 experimental 인 동안에는 MVP 의 production 경로로 쓰지 않는다. `task.allowExperimentalRuntime=true` 가 명시된 작업에서만 활성화된다.

### 4.1 Codex 두 어댑터의 운영 분기

| 항목 | `CodexExecJSONLAdapter` | `CodexAppServerAdapter` |
| --- | --- | --- |
| 공식 분류 | Stable | Experimental |
| 사용 시점 | MVP / production task | spike, 신기능 검증, 장기 thread lifecycle 필요 시 |
| 이벤트 명명 | dot 형식 (`turn.started`, `item.completed`, ...) | slash 형식 (`turn/started`, `item/completed`, ...) |
| 정규화 책임 | 어댑터 ACL 가 둘을 같은 `CanonicalEvent.Type` 으로 흡수 | 동일 |
| FSM 매핑 | `Running` 의 turn 단위 | thread/turn/item 3단 매핑 |
| 자기 신고 핸드셰이크 | 없음(stdout 첫 line 검증) | `initialize` JSON-RPC capabilities |
| 변경 위험 | 낮음 | 높음 (app-server 표면이 minor 단위로 바뀔 수 있음) |

두 어댑터가 같은 ProviderKind(`codex`) 를 공유하더라도, 같은 task 가 두 ProtocolKind 사이를 움직이지 않는다(runtime pinning — [`runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) §3).

### 4.2 agentbridge DetectResult reconciliation

Daemon adapter 의 `DetectResult` 는 C4 adapter ACL 의 raw probe 결과다. C3 `ProviderCapability` 와 `CapabilityFingerprint` 로 승격하는 구현은 daemon runtime actor 가 담당한다.

규칙:

1. runtimeactor 는 `DetectResult` 를 그대로 scheduler 조건으로 쓰지 않고, `ProviderCapability` 로 reconcile 한 뒤 daemon status 에 `protocol_kind`, `adapter_id`, `adapter_version`, `protocol_version`, `compatibility_status`, `capability_fingerprint`, `detected_fingerprint`, provider-neutral support flags 를 노출한다. 이 필드들은 RunScope `CanonicalEvent` 의 runtime capability tier 를 채우는 입력이며, EventIngestor 가 생겼을 때 같은 snapshot 을 사용한다.
2. `RIIDO_POLICY_BUNDLE_VERSION` 은 capability fingerprint 의 `PolicyBundleVersion` 입력이다. C7 policy bundle loader 가 생기기 전까지 기본값은 `policy-bundle.local.v0` 이며, 운영자가 env 로 override 할 수 있다.
3. `DetectedFingerprint` 는 `DetectResult.Executable` 이 읽을 수 있는 절대 regular file path 일 때 그 파일 내용의 SHA-256 hex 로 채운다. `Executable` 이 PATH 이름뿐이거나, 존재하지 않거나, regular file 이 아니거나, 읽을 수 없으면 빈 값으로 남긴다. 이 경우에도 capability fingerprint 는 provider/version/surface/policy 입력으로 deterministic 하게 계산된다.
4. `DetectedVersion` 은 여전히 raw signal 이며 분기 조건이 아니다. reconciliation 은 version 비교가 아니라 protocol profile + surface flags + policy version 으로만 status 를 산출한다.

### 4.3 Protocol-critical args blocklist

Protocol-critical args 는 adapter 가 provider process 와 맺은 protocol 계약 자체를 구성하는 CLI flag 다. caller 의 free-form `CustomArgs` 가 이 flag 를 다시 넣어 protocol 을 바꾸면 C4 adapter ACL 이 provider stream/RPC 를 더 이상 신뢰할 수 없으므로, adapter 는 process spawn 전에 반드시 drop 하고 `StartCommand.DroppedArgs` 로 보고해야 한다. 실행 카탈로그는 `provider/capability.ProtocolCriticalArgs` 가 소유한다.

| ProtocolKind | protocol-critical args | 이유 |
| --- | --- | --- |
| `claude-stream-json` | `-p`, `--print`, `--output-format`, `--input-format`, `--permission-mode`, `--mcp-config`, `--strict-mcp-config`, `--verbose` | prompt/stdin framing, stream-json envelope, permission mode, MCP config strictness, verbose event shape 를 adapter 가 고정한다. |
| `codex-exec-jsonl` | `--json` | NDJSON event stream 을 adapter ACL 입력으로 고정한다. 현재 repo 의 executable Codex builder 는 `codex-app-server` 이므로 이 row 는 future exec builder 의 C3 계약이다. |
| `codex-app-server` | `--listen` | JSON-RPC transport 를 `stdio://` 로 고정한다. caller 가 TCP/Unix/WebSocket transport 로 바꾸면 daemon child-process boundary 가 깨진다. |
| `openclaw-agent-json` | `--local`, `--json`, `--session-id`, `--message`, `--model`, `--system-prompt` | local execution, JSON stream/result shape, session id, prompt/message framing, agent/profile selection, system prompt fallback 을 adapter 가 고정한다. |
| `cursor-agent-stream-json` | `-p`, `--output-format`, `--yolo` | prompt flag, stream-json event shape, unsafe auto-approval surface 를 adapter/security policy 경로 밖에서 바꿀 수 없게 한다. |

규칙:

1. `BlockedArgs()` 는 위 catalog 를 참조해야 한다. provider package 가 별도 list 를 재정의하면 SSOT 위반이다.
2. `agentbridge.FilterBlockedArgs` 는 bare form(`--flag value`) 과 equals form(`--flag=value`) 을 모두 인식해야 한다.
3. protocol-critical blocklist 는 unsafe bypass allowlist 와 다르다. Codex `--yolo` / `--dangerously-bypass-approvals-and-sandbox` 같은 approval-bypass surface 의 최종 사용 가부는 downstream C7 Security / Policy gate 가 소유한다. Codex `danger-full-access` 는 sandbox mode vocabulary 에 등장할 수 있지만 이 문서가 실행 채택을 결정하지 않는다. Downstream C4 Provider Runtime 이 provider 별 trusted-runtime envelope 로 명시 채택하고 harness evidence 를 붙인 경우에만 실제 실행 선택이 된다. 이 절은 “free-form custom args 로 protocol/policy 경계를 우회할 수 없다”는 adapter ACL 계약만 소유한다.
4. 새 ProtocolKind 또는 새 protocol-critical flag 를 추가하면 본 표, `provider/capability.ProtocolCriticalArgs`, provider adapter black-box gate 를 같은 PR 에서 갱신해야 한다.

## 5. CompatibilityStatus 와 보조 필드

`CompatibilityStatus` 는 capability 한 인스턴스의 **요약 등급** 이다. scheduler 가 정확한 결정을 내리려면 보조 필드를 함께 읽어야 한다 — 요약만으로는 “왜 그렇게 됐는지” 와 “무엇이 누락됐는지” 의 정보가 소실되기 때문이다.

### 5.1 등급의 정의 (요약값)

| 등급 | 의미 | 게이트 동작 |
| --- | --- | --- |
| `supported` | 네 입력 모두 OK | 모든 capability 가 task 에 사용 가능 |
| `degraded` | capability detection 부분 실패 또는 어댑터 회귀 일부 실패 | 누락 surface 를 요구하는 task 는 `Blocked` 로 빠지고 그 외에는 진행 |
| `experimental` | protocol maturity 가 experimental 이거나, 매트릭스에 없는 새 버전 / wrapper 첫 등장 | task 가 `task.allowExperimentalRuntime=true` 일 때만 실행 |
| `blocked` | 정적 매트릭스에서 known-incompatible, policy bundle 비호환, 또는 probe 가 위험 신호 반환 | claim 자체가 거절됨 |

### 5.2 보조 필드 (information loss 방지)

요약이 하나라도 정보 손실을 일으키는 경우(예: `experimental` 이면서 동시에 `degraded`)를 피하기 위해 scheduler 는 다음 필드를 함께 읽는다.

| 필드 | 의미 |
| --- | --- |
| `ProtocolMaturity` | protocol 자체의 공식 분류 (stable / experimental / deprecated / unknown). 요약 등급과 독립적으로 유지. |
| `RequiresExperimentalOptIn` | 실행하려면 task 측 `allowExperimentalRuntime=true` 가 필요한가. `experimental` 일 때 true 가 되지만, 다른 등급에서도 wrapper 변종 등 이유로 true 가 될 수 있다. |
| `MissingCapabilities` | probe/매트릭스가 “부재” 로 결론낸 capability 이름 목록. `degraded` 일 때 비어있지 않다. |
| `BlockedReasons` | 비어있지 않으면 요약 = `blocked`. 각 reason 은 `(Code, Subject, Detail)`. |
| `DegradedReasons` | 비어있지 않으면 요약 ≥ `degraded`. |

### 5.3 등급 우선순위 (요약 압축 규칙)

여러 신호가 동시에 발생할 때 요약은 다음 우선순위로 압축된다 — 잃은 정보는 §5.2 의 보조 필드에 영구 보존된다.

```
blocked > experimental > degraded > supported
```

### 5.4 계산식 (4 입력 결합)

```
detection   ∈ {ok, partial, fail}      // §3 capability detection 결과
maturity    ∈ {stable, experimental, deprecated}   // ProtocolMaturity
policyCompat∈ {ok, incompatible}       // 현재 활성 policy bundle 과의 호환성
adapterTest ∈ {pass, partial, fail}    // 어댑터 회귀 테스트 fixture 결과
```

조합 규칙 (요약 등급 산출):

1. `policyCompat=incompatible` → 무조건 `blocked`. `BlockedReasons += {Code: "POLICY_INCOMPATIBLE"}`.
2. `adapterTest=fail` → 무조건 `blocked`. `BlockedReasons += {Code: "ADAPTER_REGRESSION_FAILED"}`.
3. `detection=fail` → 무조건 `blocked`. `BlockedReasons += {Code: "DETECTION_FAILED"}`.
4. `maturity ∈ {experimental, deprecated}` → 최소 `experimental` 등급. `RequiresExperimentalOptIn=true`.
5. `detection=partial` → 최소 `degraded`. `DegradedReasons += {Code: "PROBE_PARTIAL"}` + `MissingCapabilities` 채움.
6. `adapterTest=partial` → 최소 `degraded`. `DegradedReasons += {Code: "ADAPTER_REGRESSION_PARTIAL"}`.
7. 위에 걸리지 않으면 `supported`.

요약 등급 변환의 **시점** 은 [compatibility gate](../30-architecture/compatibility-gate.md) (G2 register / G3 heartbeat) 가 결정한다. 이 문서는 등급의 **정의·우선순위·계산식·보조 필드** 만 소유한다.

## 6. Min/MaxTestedVersion 정책

- `MinSupportedVersion` 보다 낮은 fingerprint version 은 항상 `blocked`. 매트릭스에 명시되지 않은 더 낮은 버전도 마찬가지로 `blocked`.
- `MaxTestedVersion` 보다 높은 fingerprint version 은 자동으로 `experimental`. 운영자가 명시적으로 한 PR 로 매트릭스의 `MaxTestedVersion` 을 올리고, 동시에 fixture 회귀 테스트가 통과되어야 `supported` 로 승격한다.
- `MinSupportedVersion`/`MaxTestedVersion` 은 fingerprint version 그대로의 텍스트가 아니라 **semver 비교 함수**로 평가한다. fingerprint 가 semver 가 아닌 wrapper(예: 사내 빌드 해시) 의 경우 `MaxTestedVersion` 검사를 skip 하고 자동으로 `experimental` 로 둔다.

## 7. Wrapper 변종 처리

내부 gateway/wrapper CLI 는 다음 조건 하에 같은 capability 모델로 흡수된다.

1. wrapper 가 `riido-wrapper-manifest.json` 을 `--print-manifest` 같은 표준 옵션으로 노출하고, 거기에 `protocolKind`, `protocolVersion`, `surfaceFlags`, `trustTier`, `originalBinary` 를 자기 신고한다.
2. wrapper 가 매니페스트를 노출하지 않으면 자동으로 `experimental` + `degraded` 로 분류하고, capability 는 정적 매트릭스의 **최소 보장값** 까지만 신뢰한다.
3. wrapper 가 stock 보다 “더 적게 허용”(예: file write 차단) 한 경우 그 사실이 capability 로 반영되어야 한다. wrapper 의 자기 신고를 capability AND 정적 매트릭스 의 **교집합** 으로 계산한다.

이 절은 Multica 가 외부 issue 로 노출했던 “Claude-compatible wrapper runtime 등록” 요구사항을 (참고: [`../10-research/multica.md`](../10-research/multica.md)) 우리 SSOT 안에서 1급으로 흡수한다.

## 8. Provider-specific surface (ClaudeSurface / CodexSurface)

provider-neutral capability 만으로는 어댑터 코드가 “정확히 어떤 CLI 플래그/매개변수를 쓰는가” 를 표현할 수 없다. 그 vendor-only 디테일은 본 구조체들이 가진다. capability 한 인스턴스에서 본 인스턴스의 `ProviderKind` 에 해당하는 surface 하나만 채워지고, 나머지는 `nil` 이다.

```go
type ClaudeSurface struct {
    // CLI 옵션 노출 여부 (Claude 전용 플래그명)
    OutputStreamJSON     bool  // claude --output-format stream-json
    InputStreamJSON      bool  // claude --input-format stream-json
    PartialMessages      bool  // claude --include-partial-messages
    HookEvents           bool  // PreToolUse / PostToolUse / PermissionRequest / SessionStart 등
    PermissionMode       bool  // claude --permission-mode (default | acceptEdits | plan | bypassPermissions)
    ManagedSettings      bool  // managed/local/project/user precedence 적용 가능
    AddDir               bool  // claude --add-dir
    MCPConfig            bool  // claude --mcp-config
    StrictMCPConfig      bool  // claude --strict-mcp-config

    // permission-mode 의 enum 값 중 우리가 인지한 것
    PermissionModeValues []string  // 예: ["default", "acceptEdits", "plan", "bypassPermissions"]
}

type CodexSurface struct {
    // 두 호출 모드의 가용성
    ExecMode             bool  // codex exec (Stable)
    AppServerMode        bool  // codex app-server (Experimental)

    // app-server 시 인지 가능한 JSON-RPC method 집합
    AppServerMethods     []string  // 예: ["initialize", "thread/start", "thread/resume", "turn/start", "turn/interrupt"]

    // sandbox / approval 의 enum 값
    SandboxModes         []string  // 예: ["read-only", "workspace-write", "danger-full-access"]
    ApprovalModes        []string  // 예: ["untrusted", "on-request", "never"]

    // AGENTS.md 발견 알고리즘
    AgentsMdGlobalChain  bool      // global override → global base → root → ... → cwd
    AgentsMdSizeCapBytes int       // 예: 32768 (32 KiB)

    // 우회 surface (risk signal — 사용 허가 아님)
    Yolo                       bool  // codex --yolo
    DangerouslyBypassFlag      bool  // codex --dangerously-bypass-approvals-and-sandbox
}
```

규칙:

1. `ClaudeSurface` 와 `CodexSurface` 의 필드는 **CLI/RPC 표면에서 직접 검출된 사실** 만 적는다. policy 도메인 결정은 [`security.md`](./security.md) 가 소유한다.
2. provider 가 새 옵션을 추가하면(예: Claude 새 `--something`) 본 구조체에 필드를 추가하는 것은 **additive** 다. 옵션 제거 또는 의미 변경은 `change:breaking-ir`.
3. `ClaudeSurface.PermissionMode=true` + `ClaudeSurface.PermissionModeValues` 가 `"bypassPermissions"` 를 포함하면, 상위 `ExposesUnsafePermissionBypass=true` 로 승격된다. `CodexSurface.Yolo` 또는 `CodexSurface.DangerouslyBypassFlag` 도 동일.
4. wrapper 어댑터의 경우, wrapper 매니페스트가 어떤 surface 를 “감춘다” 고 선언하면 그 필드는 `false` 가 된다(§7 wrapper 변종 교집합 규칙).

## 9. 인접 SSOT 와의 계약

| 인접 SSOT | 이 문서가 제공/요구하는 것 |
| --- | --- |
| [`ir-event-log.md`](./ir-event-log.md) | 모든 `CanonicalEvent` 가 9 종 실행 지문(`ProviderKind`/`ProtocolKind` 와 7 개 version 필드)을 기록한다. FSM transition event 에는 `FSMVersion` 추가. 자세한 강제 규칙은 [`ir-schema-versioning.md`](./ir-schema-versioning.md) §2 |
| [`task-lifecycle.md`](./task-lifecycle.md) | `Claimed`/`Preparing` 진입 시 capability ⊇ task.requirement 검사 통과가 invariant |
| [`runtime-scheduling.md`](./runtime-scheduling.md) | `RuntimeLease` 는 capability fingerprint 를 포함한다. fingerprint 가 바뀌면 lease 무효 |
| [`security.md`](./security.md) | wrapper 의 `trustTier`, sandbox 기본값, `ExposesUnsafePermissionBypass=true` 의 최종 사용 가부 결정. 본 문서는 “노출 여부” 만 신고, 사용 결정은 security 가 소유 |
| [`runtime-versioning.md`](./runtime-versioning.md) | 본 capability 가 어떤 11 축 위에서 의미를 갖는지의 hub. D·E·F 축이 본 문서로 위임됨 |
| [`compatibility-gate.md`](../30-architecture/compatibility-gate.md) | `CompatibilityStatus` 와 §5.2 보조 필드를 함께 사용하는 게이트의 단일 진입점 |
| [`runtime-upgrade-flow.md`](../30-architecture/runtime-upgrade-flow.md) | 재탐지가 일어났을 때 task state 별 처리 |

## 10. 미결정/오픈 이슈

`open-questions.md` 의 다음 항목으로 위임한다.

- `Q-CAP-001`: probe 가 비용을 발생시킬 때(예: app-server cold start) probe 캐시 TTL.
- `Q-CAP-002`: wrapper 매니페스트 표준의 외부화 — 공개 spec 으로 둘지, 사내 전용으로 둘지.
- `Q-CAP-003`: experimental capability 를 활성화한 task 가 만든 IR 이벤트의 별도 표시 방식.
- `Q-CAP-004`: `ClaudeSurface.PermissionModeValues` / `CodexSurface.SandboxModes` 처럼 enum 멤버 자체가 minor 단위로 늘어나는 케이스의 PR 라벨 분류.

## 11. version-affecting changes

이 문서가 정의하는 provider-neutral surface flag 의 추가는 **additive** 다. 기존 flag 의 의미 변경, 제거, provider-specific surface 의 필드 의미 변경, 또는 protocol 분류 통합은 **breaking-ir** 로 간주되어 PR 라벨 `change:breaking-ir` 가 강제된다. 자세한 PR 라벨 규칙은 [`runtime-versioning.md`](./runtime-versioning.md) §6 가 소유한다.
