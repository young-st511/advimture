# UX-QA-001 — Review Route Evidence Contract

Slice-ID: UX-QA-001
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/ux-qa-001-review-route-evidence.md
- docs/exec-plans/completed/ux-qa-001-review-route-evidence.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/tui-ui-qa-contract.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- cmd/e2e-runner/main.go
- cmd/e2e-runner/main_test.go
- test/e2e/playable_review_queue.yaml

## 목표

UI 리디자인 전에 review/daily route를 문자열 검증이 아니라 typed app_state assertion과 재확인 가능한 evidence로 검증한다.

## 범위

- 포함:
  - `assert.app_state.review` typed assertion
  - E2E evidence의 `app_state.json`, `progress.json` snapshot 저장
  - `playable_review_queue.yaml`의 review contains 제거
  - UI/QA 방향 문서화
- 제외:
  - 실제 UI 레이아웃 변경
  - progress 저장 포맷 변경
  - content schema 변경
  - 새 의존성 추가
  - screen final/frame parser 구현

## 수용 기준

- completed: `assert.app_state.review.queue_count`를 검증할 수 있다.
- completed: `assert.app_state.review.primary_exercise_id`를 검증할 수 있다.
- completed: `assert.app_state.review.primary_reason`을 검증할 수 있다.
- completed: `assert.app_state.review.daily_route`를 검증할 수 있다.
- completed: `playable_review_queue.yaml`은 review 검증에 JSON 문자열 `contains`를 사용하지 않는다.
- completed: evidence 옵션으로 `app_state.json`과 `progress.json` snapshot을 저장할 수 있다.
- completed: temp HOME 삭제 후에도 evidence directory에서 app state와 progress를 읽을 수 있다.

## 완료 내용

- E2E runner에 `assert.app_state.review` typed assertion을 추가했다.
- `evidence.save_app_state`, `evidence.save_progress`를 추가했다.
- `runResult`가 temp HOME 정리 전에 app state/progress bytes를 수집한다.
- `playable_review_queue.yaml`은 review 상태를 typed assertion으로 검증한다.
- UI/UX 통합 결론은 `docs/gameplay/tui-ux-direction.md`에, QA 계약은 `docs/verification/tui-ui-qa-contract.md`에 기록했다.

## 검증 결과

- passed: `go test ./cmd/e2e-runner/...`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
