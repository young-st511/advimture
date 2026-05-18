# ExecPlan: Command intent scoring

Slice-ID: SCORING-002
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/spec.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/scoring-002-command-intent.md
- docs/exec-plans/completed/scoring-002-command-intent.md
- internal/scoring/scoring.go
- internal/scoring/scoring_test.go
- internal/scenario/run.go
- internal/scenario/run_test.go
- test/e2e/playable_constraint_required_key.yaml

## 목표

목표 상태 도달 여부와 별개로 “이번 문항에서 의도한 command를 사용했는가”가 점수와 코칭에 명시적으로 반영되도록 scoring 계약을 정리한다.

## 수용 기준

- scoring input/result는 runtime failure reason을 보존한다.
- `required_keys_missing` 실패는 `Passed=false`, `IntentSatisfied=false`, `Grade=F`로 평가된다.
- 일반 성공은 `IntentSatisfied=true`로 평가된다.
- scenario는 runtime failure reason을 scoring input에 전달한다.
- 기존 constraint E2E는 required key 누락 시 실패/Grade F/coaching을 검증한다.

## 범위

- 포함: scoring result metadata
- 포함: scenario scoring input 연결
- 포함: unit test
- 제외: 추가 감점 공식 변경
- 제외: 성공했지만 exact key가 아닌 경우의 세부 코칭 UI

## 검증 계획

- `go test ./internal/scoring/...`
- `go test ./internal/scenario/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_required_key.yaml`
- `git diff --check`

## 검증 결과

- `go test ./internal/scoring/...`: pass
- `go test ./internal/scenario/...`: pass
- `go test ./...`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_required_key.yaml`: pass
- `git diff --check`: pass

## 작업 항목

- [x] scoring input/result에 intent metadata를 추가한다.
- [x] scenario가 runtime failure reason을 전달한다.
- [x] required key 누락 scoring test를 추가한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
