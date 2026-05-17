# ExecPlan: Content schema foundation

Slice-ID: VIM-005
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-005-content-schema-foundation.md
- internal/content/**

## 목표

문서/데이터로 표현될 exercise spec을 `internal/runtime`이 실행할 수 있는 runtime exercise로 변환하는 Content Schema 모듈을 만든다. 이번 slice는 YAML loader가 아니라, content와 runtime 사이의 타입 경계를 고정하는 작업이다.

## 범위

- 포함: `internal/content` 패키지 생성
- 포함: exercise spec, state spec, goal spec, hint spec
- 포함: spec validation
- 포함: spec을 runtime exercise로 compile
- 포함: expected/allowed key metadata 보존
- 제외: YAML/JSON 파일 로더
- 제외: scenario skinning
- 제외: scoring/grade
- 제외: Bubble Tea 연결

## 연결된 엔진

- 입력: content exercise spec
- 출력: `runtime.Exercise` + content metadata
- 다음 연결 후보: Grader / Scoring Engine

## 구현 계획

1. `ExerciseSpec`, `StateSpec`, `GoalSpec`, `HintSpec` 타입을 만든다.
2. `CompileExercise(ExerciseSpec)`가 validation 후 `runtime.Exercise`를 반환한다.
3. mode 문자열은 `vimengine.Mode`로 명시적으로 변환한다.
4. cursor는 buffer line 범위 안에 있어야 한다.
5. expected/allowed keys는 runtime에 섞지 않고 compiled metadata로 보존한다.

## 검증 계획

- `go test ./internal/content`
- `go test ./...`
- 변경 종료 전 `git diff -- docs/roadmap/ENGINE_TODO.md docs/roadmap/PROGRAM.md docs/exec-plans/active/vim-005-content-schema-foundation.md internal/content`

## E2E Evidence

이번 slice는 TUI runtime에 연결하지 않으므로 TUI E2E는 필수 evidence가 아니다.

## 승인 체크

- [x] 유효한 exercise spec이 runtime exercise로 compile된다.
- [x] compile된 exercise를 runtime session으로 실행할 수 있다.
- [x] missing id, invalid mode, out-of-range cursor를 거부한다.
- [x] expected/allowed key metadata가 copy로 보존된다.
- [x] content 모듈은 TUI, progress, filesystem에 의존하지 않는다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: YAML/JSON loader는 이번 slice에서 제외했다. 먼저 content spec과 runtime exercise 사이의 순수 타입 경계를 고정하기 위함이다.
- 2026-05-18: `ExpectedKeys`와 `AllowedKeys`는 runtime에 넣지 않고 compiled metadata로 보존했다. runtime은 진행만, scoring/validation 정책은 다음 엔진에서 다루기 위함이다.
