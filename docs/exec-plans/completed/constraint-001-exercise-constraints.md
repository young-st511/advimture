# ExecPlan: Exercise constraint runtime

Slice-ID: CONSTRAINT-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- content/exercises/first-five-minutes.yaml
- docs/gameplay/spec.md
- docs/gameplay/content-requirements.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/constraint-001-exercise-constraints.md
- docs/exec-plans/completed/constraint-001-exercise-constraints.md
- Makefile
- internal/content/exercise.go
- internal/content/exercise_test.go
- internal/content/loader.go
- internal/content/loader_test.go
- internal/runtime/session.go
- internal/runtime/session_test.go
- internal/scenario/run.go
- internal/scenario/run_test.go
- internal/tuiadapter/adapter.go
- internal/playable/model.go
- internal/playable/model_test.go
- test/e2e/**

## 목표

Exercise가 의도한 command를 학습시키도록 최대 입력 수, 필수 입력, 금지 입력 제약을 콘텐츠 schema와 runtime에 연결한다.

## 영향 도메인

- Content: YAML exercise가 `constraints`를 선언할 수 있다.
- Runtime: 금지 입력, 최대 입력 초과, 필수 입력 누락은 실패 상태로 전이한다.
- Scenario: 실패 시 mentor failure copy를 보여주고 재시도를 허용한다.
- Playable: 실패 상태에서 `r`과 `enter` 모두 retry로 동작한다.
- Verification: 우회 입력, 최대 입력 초과, retry 복구를 TUI E2E로 검증한다.

## 수용 기준

- exercise YAML은 `constraints.max_inputs`, `constraints.required_keys`, `constraints.forbidden_keys`, `constraints.attempt_limit`를 로드한다.
- `max_inputs`는 0 이상이어야 하며, 0은 제한 없음이다.
- `attempt_limit`는 0 이상이어야 하며, 0은 무제한이다.
- forbidden key 입력은 Vim state를 진행시키지 않고 즉시 failed 상태가 된다.
- max input을 초과하는 입력은 Vim state를 진행시키지 않고 즉시 failed 상태가 된다.
- 목표 상태에 도달했더라도 `required_keys`가 trace에 없으면 succeeded가 아니라 failed가 된다.
- failed 상태에서는 progress를 저장하지 않는다.
- failed 상태에서 `r` 또는 `enter`를 누르면 같은 exercise를 재시도한다.
- 실패 화면은 실패 메시지, 남은 입력 수, retry 안내를 보여준다.
- 기존 progress 저장 포맷은 변경하지 않는다.

## 범위

- 포함: exercise constraint schema와 loader validation
- 포함: runtime failed status와 failure reason
- 포함: scenario failure copy 연결
- 포함: playable retry handling
- 포함: 첫 movement exercise의 constraint fixture와 E2E
- 제외: attempt_limit enforcement
- 제외: 정교한 strategy/semantic shortcut detector
- 제외: scoring rubric의 command-intent 감점 상세화
- 제외: 별도 실패 화면 레이아웃 대개편

## 검증 계획

- `go test ./internal/content/...`
- `go test ./internal/runtime/...`
- `go test ./internal/scenario/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 검증 결과

- `go test ./internal/content/...`: pass
- `go test ./internal/runtime/...`: pass
- `go test ./internal/scenario/...`: pass
- `go test ./internal/playable/...`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass
- `git diff --check`: pass

## 작업 항목

- [x] constraint schema와 compile mapping을 추가한다.
- [x] loader validation과 content fixture를 추가한다.
- [x] runtime failed status와 constraint enforcement를 추가한다.
- [x] scenario/playable 실패 메시지와 retry 흐름을 연결한다.
- [x] E2E fixtures를 추가한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
