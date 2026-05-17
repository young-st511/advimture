# ExecPlan: Scenario runtime

Slice-ID: VIM-007
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-007-scenario-runtime.md
- internal/scenario/**

## 목표

검증된 exercise runtime과 scoring 결과를 어드벤처 맥락으로 감싸는 Scenario Runtime을 만든다. 이 모듈은 UI를 렌더링하지 않고, 현재 scenario message, hint 사용, retry, success scoring을 조율한다.

## 범위

- 포함: `internal/scenario` 패키지 생성
- 포함: scenario spec과 run state
- 포함: briefing/success message 선택
- 포함: key 입력을 exercise runtime에 전달
- 포함: hint request count
- 포함: success 시 scoring 평가
- 제외: Bubble Tea view
- 제외: progress 저장
- 제외: YAML loader
- 제외: 복수 exercise mission flow

## 연결된 엔진

- 입력: content compiled exercise
- 내부 사용: runtime session, scoring engine
- 출력: scenario run state
- 다음 연결 후보: TUI Adapter, Progress Adapter

## 구현 계획

1. `Spec`, `Run`, `State`, `StepResult` 타입을 만든다.
2. `NewRun(Spec)`은 runtime session을 시작하고 briefing message를 노출한다.
3. `ApplyKey(key)`는 session에 key를 전달한다.
4. session이 succeeded가 되면 scoring result를 계산하고 success message를 노출한다.
5. `RequestHint()`는 runtime hint를 읽고 hint 사용 수를 기록한다.
6. `Retry()`는 runtime session을 초기화하고 scenario message를 briefing으로 되돌린다.

## 검증 계획

- `go test ./internal/scenario`
- `go test ./...`
- 변경 종료 전 `git diff -- docs/roadmap/ENGINE_TODO.md docs/roadmap/PROGRAM.md docs/exec-plans/active/vim-007-scenario-runtime.md internal/scenario`

## E2E Evidence

이번 slice는 TUI runtime에 연결하지 않으므로 TUI E2E는 필수 evidence가 아니다.

## 승인 체크

- [x] scenario runtime은 Bubble Tea, progress, filesystem에 의존하지 않는다.
- [x] briefing 상태에서 시작한다.
- [x] key 입력으로 exercise가 성공하면 success message와 score가 생긴다.
- [x] hint request가 deterministic하게 기록된다.
- [x] retry가 runtime state와 score를 reset한다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: scenario runtime은 UI 문구를 렌더링하지 않고 현재 message만 노출한다. Bubble Tea adapter가 렌더링 책임을 가져야 하기 때문이다.
- 2026-05-18: hint 사용 수는 scenario runtime에서 기록한다. runtime은 hint 후보만 제공하고, 실제 사용 여부는 플레이 흐름을 조율하는 scenario가 가장 잘 안다.
