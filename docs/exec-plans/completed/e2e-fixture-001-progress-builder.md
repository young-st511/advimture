# E2E-FIXTURE-001 — Progress Fixture Builder

Slice-ID: E2E-FIXTURE-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: e2e-runner-and-fixtures
Allowed-Paths:
- docs/exec-plans/active/e2e-fixture-001-progress-builder.md
- docs/exec-plans/completed/e2e-fixture-001-progress-builder.md
- docs/verification/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- cmd/e2e-runner/main.go
- cmd/e2e-runner/main_test.go
- test/e2e/*.yaml

## 목표

긴 inline `setup.progress_file` JSON을 줄이고, content playlist 순서가 바뀌어도 E2E fixture가 덜 낡도록 runner에 progress builder를 추가한다.

## 완료 내용

- `setup.complete_before: <exercise-id>`를 추가했다.
- runner가 현재 `content/`의 playable playlist 순서를 읽고 지정 exercise 직전까지 completed progress를 생성하게 했다.
- `setup.progress_file`과 `setup.complete_before` 동시 사용을 오류로 막았다.
- 지정 exercise가 playable sequence에 없으면 오류를 반환하게 했다.
- 긴 progress JSON을 쓰던 후반 playpack/incident E2E를 `complete_before`로 교체했다.

## 제외한 것

- 실제 사용자 HOME 사용
- progress 저장 포맷 변경
- content schema 변경
- playlist unlock 정책 변경
- 외부 fixture 파일 포맷 추가

## 검증 결과

- passed: `go test ./cmd/e2e-runner`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_open_line_full.yaml`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
