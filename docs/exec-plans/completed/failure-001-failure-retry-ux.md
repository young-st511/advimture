# ExecPlan: Failure and retry UX

Slice-ID: FAILURE-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/spec.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/failure-001-failure-retry-ux.md
- docs/exec-plans/completed/failure-001-failure-retry-ux.md
- internal/runtime/session.go
- internal/runtime/session_test.go
- internal/playable/model.go
- internal/playable/model_test.go
- test/e2e/playable_constraint_forbidden_retry.yaml
- test/e2e/playable_constraint_max_inputs.yaml

## 목표

실패가 학습 루프를 끊지 않도록 실패 화면에 원인, 입력 여유, 시도 횟수, 재시도 명령을 안정적으로 표시한다.

## 수용 기준

- failed state는 현재 attempt count와 attempt limit metadata를 노출한다.
- `attempt_limit: 0`은 UI에서 `unlimited`로 표시한다.
- failed 화면은 `Retry: r or enter`를 보여준다.
- `r`와 `enter`는 같은 exercise를 retry하고 attempt count를 증가시킨다.
- retry는 progress를 저장하지 않는다.
- 기존 progress 저장 포맷은 변경하지 않는다.

## 범위

- 포함: runtime state metadata
- 포함: playable failed UI copy
- 포함: unit/E2E assertion 보강
- 제외: attempt_limit 강제 실패
- 제외: 별도 modal/화면 전환

## 검증 계획

- `go test ./internal/runtime/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_forbidden_retry.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_max_inputs.yaml`
- `git diff --check`

## 검증 결과

- `go test ./internal/runtime/...`: pass
- `go test ./internal/playable/...`: pass
- `go test ./...`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_forbidden_retry.yaml`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_max_inputs.yaml`: pass
- `git diff --check`: pass

## 작업 항목

- [x] runtime state에 attempt limit metadata를 노출한다.
- [x] failed UI에 attempt count를 표시한다.
- [x] unit/E2E assertion을 보강한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
