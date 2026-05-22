# UI-EVIDENCE-001 — Screen Evidence Timeline

Slice-ID: UI-EVIDENCE-001
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: e2e-runner-and-docs
Allowed-Paths:
- docs/exec-plans/active/ui-evidence-001-screen-timeline.md
- docs/exec-plans/completed/ui-evidence-001-screen-timeline.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- docs/verification/tui-e2e-loop.md
- docs/verification/tui-ui-qa-contract.md
- cmd/e2e-runner/main.go
- cmd/e2e-runner/main_test.go
- test/e2e/playable_review_queue.yaml

## 목표

TUI UI 변경을 반복할 때 Agent가 cleaned 누적 화면 흐름을 명시적으로 확인할 수 있도록 E2E evidence를 보강한다. 첫 범위는 flaky한 terminal frame parser를 새로 만들지 않고, 기존 cleaned terminal text를 `screen_timeline.txt`라는 명확한 evidence로 추가한다.

## 완료 내용

- E2E scenario에 `evidence.save_screen_timeline` 옵션을 추가했다.
- 옵션 선언 시 artifacts 아래 `screen_timeline.txt`를 저장한다.
- `summary.json`에 `screen_timeline_evidence`를 기록한다.
- 대표 UI QA scenario인 `playable_review_queue.yaml`가 timeline evidence를 남기도록 갱신했다.
- 기존 `screen.txt`, `raw.log`, `app_state.json`, `progress.json` 동작은 유지했다.

## 검증 결과

- `go test ./cmd/e2e-runner/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- `git diff --check`
- artifacts 확인: `artifacts/e2e/playable_review_queue/screen_timeline.txt`, `summary.json`의 `screen_timeline_evidence: true`

## 후속 작업

다음 UI 루프에서는 실제 viewport parser가 필요해질 때 `screen_final.txt` 또는 `frames/*.txt`를 별도 ExecPlan으로 열어 구현한다.
