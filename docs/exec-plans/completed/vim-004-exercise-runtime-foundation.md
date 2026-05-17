# ExecPlan: Exercise runtime foundation

Slice-ID: VIM-004
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-004-exercise-runtime-foundation.md
- internal/runtime/**

## 목표

`internal/vimengine` 위에 첫 미들웨어 엔진인 Exercise Runtime을 만든다. 이 runtime은 TUI, progress, scenario를 모르고, 하나의 exercise session을 시작하고 key 입력을 적용하며 목표 상태 도달 여부를 판정한다.

## 범위

- 포함: `internal/runtime` 패키지 생성
- 포함: exercise spec, goal, session state
- 포함: key trace 기록
- 포함: cursor/mode/lines 목표 판정
- 포함: retry와 hint
- 제외: Bubble Tea 연결
- 제외: YAML content loader
- 제외: scoring/grade 산정
- 제외: progress 저장
- 제외: 새 Vim command 구현

## 연결된 엔진

- 입력: `internal/vimengine`
- 출력: runtime session state
- 다음 연결 후보: Content Schema, Grader, TUI Adapter

## 구현 계획

1. `Exercise`, `Goal`, `Session`, `Status` 타입을 만든다.
2. `NewSession(Exercise)`가 초기 Vim state를 copy해서 보관한다.
3. `ApplyKey(key)`는 `vimengine.Apply`를 호출하고 key trace를 남긴다.
4. 입력 후 `Goal`이 충족되면 session status를 `succeeded`로 바꾼다.
5. `Retry()`는 초기 state와 key trace를 리셋하고 attempt를 증가시킨다.
6. `CurrentHint()`는 key trace 길이에 따라 hint를 노출한다.

## 검증 계획

- `go test ./internal/runtime`
- `go test ./...`
- 변경 종료 전 `git diff -- docs/roadmap/ENGINE_TODO.md docs/roadmap/PROGRAM.md docs/exec-plans/active/vim-004-exercise-runtime-foundation.md internal/runtime`

## E2E Evidence

이번 slice는 TUI runtime에 연결하지 않으므로 TUI E2E는 필수 evidence가 아니다. TUI Adapter slice에서 E2E schema가 부족하면 구현을 중단하고 assertion을 먼저 보강한다.

## 승인 체크

- [x] runtime은 Bubble Tea, 파일 시스템, progress 저장에 의존하지 않는다.
- [x] 정답 key trace가 목표 cursor에 도달하면 `succeeded`가 된다.
- [x] 미완료 입력은 `running` 상태를 유지한다.
- [x] unsupported key도 trace와 vim event에 남는다.
- [x] retry가 초기 state, key trace, status를 재설정한다.
- [x] hint가 deterministic하게 노출된다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: `internal/runtime`은 `vimengine`만 import한다. TUI, progress, content loader와 분리해서 exercise session 계약을 먼저 고정하기 위함이다.
- 2026-05-18: scoring/grade는 이번 slice에 넣지 않았다. 성공 판정과 점수 산정을 분리해야 이후 힌트, 최적 키, 실수 평가를 독립적으로 바꿀 수 있다.
