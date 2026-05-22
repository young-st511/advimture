# PLATFORM-REVIEW-002 — Review Daily Route Motivation

Slice-ID: PLATFORM-REVIEW-002
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/platform-review-002-daily-route.md
- docs/exec-plans/completed/platform-review-002-daily-route.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/e2estate/
- internal/playable/
- test/e2e/playable_review_queue.yaml

## 목표

저장 포맷을 변경하지 않고 기존 review queue를 "오늘의 복구 루트"로도 표현해 반복 플레이 동기를 강화한다.

## 수용 기준

- completed: `internal/progress/` 저장 JSON 구조를 변경하지 않는다.
- completed: daily route는 앱 실행 시 content library와 progress v1에서 계산한 review queue만 읽는다.
- completed: 메인 화면은 review queue와 함께 오늘의 복구 루트 요약을 표시한다.
- completed: 성공 debrief는 다음 복구 루트가 남아 있음을 action panel에서 보여준다.
- completed: E2E state summary는 review queue count와 primary candidate를 노출한다.
- completed: focused E2E는 화면 문구, key trace, app_state raw assertion으로 검증한다.

## 제외 항목

- progress schema v2
- daily streak/calendar 저장
- spaced review due date 저장
- review 후보 파일 저장
- 새 메뉴 또는 playlist 생성

## 완료 내용

- TUI 상단과 성공 action panel에 `오늘의 복구 루트: N건 대기`를 표시한다.
- `e2estate.State`에 저장 변경 없는 테스트용 `review` summary를 추가했다.
- `playable_review_queue` E2E가 화면, key trace, progress, app_state raw review summary를 함께 검증한다.

## 검증 결과

- passed: `go test ./internal/e2estate/...`
- passed: `go test ./internal/playable/... ./internal/review/...`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`

## 저장 포맷 확인

`internal/progress/`와 progress JSON schema는 변경하지 않았다.
