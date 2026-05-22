# QA-SEL-001 — Selection Replay Hardening

Slice-ID: QA-SEL-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: test-and-validation
Allowed-Paths:
- docs/exec-plans/active/qa-sel-001-selection-replay-hardening.md
- docs/exec-plans/completed/qa-sel-001-selection-replay-hardening.md
- docs/gameplay/spec.md
- docs/verification/tui-e2e-loop.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/content/loader.go
- internal/content/loader_test.go

## 목표

content replay gate가 `e2e_assertions.selection` mismatch를 놓치지 않도록 보강한다.

## 수용 기준

- `replay_status: pass`인 exercise가 `e2e_assertions.selection`을 선언하면 optimal key replay 결과의 selection과 비교한다.
- assertion이 active selection을 요구하는데 replay state에 selection이 없으면 load가 실패한다.
- selection kind, anchor, head, start, end mismatch가 load 실패로 드러난다.
- selection assertion을 생략한 기존 exercise는 현재처럼 통과한다.
- E2E runner와 content replay의 selection 검증 의도를 문서에 반영한다.

## 결과

- `validateReplay`가 `e2e_assertions.selection`을 비교하도록 확장했다.
- active selection 누락과 selection end mismatch를 loader test로 고정했다.
- visual selection 검증이 content load와 TUI E2E 양쪽에서 같은 shape를 보게 문서를 갱신했다.

## 제외한 것

- runtime `Goal`에 selection target 추가
- progress 저장 포맷 변경
- visual engine behavior 변경
- E2E runner assertion semantics 변경

## 검증

- `go test ./internal/content/...`
- `go test ./cmd/e2e-runner/...`
- `go test ./...`
- `git diff --check`
