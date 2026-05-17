# ExecPlan: Vim engine contract hardening

Slice-ID: VIM-003
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/vim-003-vimengine-contract-hardening.md
- docs/exec-plans/completed/vim-002-vimengine-foundation.md
- internal/vimengine/**

## 목표

`internal/vimengine`을 다음 단계의 exercise runtime이 안정적으로 사용할 수 있게 만든다. 이번 slice는 새 command를 늘리는 것이 아니라, 엔진 계약과 상태 경계를 단단하게 만드는 작업이다.

## 범위

- 포함: 초기 `State` 주입 API
- 포함: key 상수 정의
- 포함: key trace를 순차 적용하는 helper
- 포함: 외부에서 들어온 비정상 `State`의 정규화 테스트
- 포함: event 순서와 상태 copy 경계 테스트
- 제외: `w`, `b`, `e` 등 새 Vim command
- 제외: exercise runtime 구현
- 제외: TUI 연결
- 제외: Neovim oracle runner 구현

## 구현 계획

1. `NewWithState(State)`를 추가해 exercise가 원하는 초기 buffer/cursor/mode를 주입할 수 있게 한다.
2. `KeyH`, `KeyJ`, `KeyK`, `KeyL`, `KeyEsc` 상수를 추가해 문자열 오타를 줄인다.
3. `ApplyKeys(State, []string)`와 `Engine.ApplyKeys([]string)`를 추가해 key trace 재생을 지원한다.
4. `normalizeState`의 cursor/desired column 경계를 테스트로 고정한다.
5. `Result.Events`가 key 입력 순서를 보존하는지 테스트한다.

## 검증 계획

- `go test ./internal/vimengine`
- `go test ./...`
- 변경 종료 전 `git diff -- docs/roadmap/PROGRAM.md docs/exec-plans/active/vim-003-vimengine-contract-hardening.md docs/exec-plans/completed/vim-002-vimengine-foundation.md internal/vimengine`

## E2E Evidence

이번 slice는 TUI runtime에 연결하지 않으므로 TUI E2E는 필수 evidence가 아니다. runtime 또는 UI adapter에 연결되는 순간 E2E schema가 부족하면 구현을 중단하고 assertion을 먼저 보강한다.

## 승인 체크

- [x] 외부 `State`를 주입해도 engine 내부 상태가 copy로 보호된다.
- [x] key trace 적용 결과가 단일 key 반복 적용과 일치한다.
- [x] event 순서가 입력 순서와 일치한다.
- [x] 비정상 cursor 값이 안전하게 정규화된다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: 새 command 구현 대신 engine 계약을 강화했다. 다음 runtime layer가 상태 주입과 key trace replay를 안정적으로 사용할 수 있게 하기 위함이다.
- 2026-05-18: `DesiredCol`은 짧은 줄 이동 후 긴 줄로 복귀하는 Vim식 세로 이동을 위해 현재 `Col`보다 클 수 있다. 단, 외부에서 들어온 `DesiredCol < Col` 상태는 정규화한다.
