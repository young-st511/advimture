# E2E-007 — Selection App State Assertion

Slice-ID: E2E-007
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: e2e-state-and-runner
Allowed-Paths:
- docs/exec-plans/active/e2e-007-selection-app-state-assertion.md
- docs/exec-plans/completed/e2e-007-selection-app-state-assertion.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/verification/selection-app-state-contract.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/e2estate/state.go
- internal/e2estate/state_test.go
- internal/content/loader.go
- internal/content/loader_test.go
- cmd/e2e-runner/main.go
- cmd/e2e-runner/main_test.go

## 목표

VISUAL-GAP-002에서 정의한 `selection` object를 E2E state summary, content `e2e_assertions`, runner `assert.app_state`에 추가한다.

## 완료 내용

- `internal/e2estate.State`에 optional `selection`을 추가했다.
- content `e2e_assertions.selection`을 YAML loader에서 보존한다.
- runner `assert.app_state.selection`이 `active`, `kind`, `anchor`, `head`, `start`, `end`를 검증한다.
- selection assertion이 있으면 다른 app_state assertion 없이도 app state summary를 읽는다.
- 기존 E2E fixture는 selection 없이 계속 통과한다.

## 제외한 것

- visual engine 구현
- TUI selection rendering 구현
- visual mode content/playpack 추가

## 검증 결과

- passed: `go test ./internal/e2estate/...`
- passed: `go test ./internal/content/...`
- passed: `go test ./cmd/e2e-runner`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
