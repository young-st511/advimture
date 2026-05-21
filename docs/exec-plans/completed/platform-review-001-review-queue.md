# PLATFORM-REVIEW-001 — Read-Only Review Queue

Slice-ID: PLATFORM-REVIEW-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/platform-review-001-review-queue.md
- docs/exec-plans/completed/platform-review-001-review-queue.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- internal/review/
- internal/playable/
- test/e2e/
- Makefile

## 목표

저장 포맷을 변경하지 않고 기존 progress v1 `Missions`와 content library만 읽어 복습 후보를 계산하고, TUI에 세계관적으로 의미 있는 review queue를 표시한다.

세계관 명칭:

- 메인 첫 화면: `재진단 큐`
- playlist/exercise 성공 후: `잔류 리스크`

## 범위

- 포함:
  - `internal/review` 순수 계산 패키지
  - 낮은 best grade, optimal보다 높은 best keystrokes, incomplete exercise 추천
  - playable 첫 화면 상단의 read-only review summary
  - 성공 action panel의 다음 review recommendation
  - model-level test와 focused E2E
- 제외:
  - `internal/progress/` 저장 JSON 구조 변경
  - spaced review date 저장
  - daily streak/calendar
  - review candidate를 별도 파일에 저장

## 수용 기준

- completed: review 계산은 content library와 progress snapshot을 입력으로 받고 progress를 mutate하지 않는다.
- completed: incomplete exercise는 가장 높은 priority로 추천된다.
- completed: completed exercise 중 best grade가 S가 아니면 review 후보가 된다.
- completed: completed exercise 중 best keystrokes가 exercise optimal key count보다 크면 review 후보가 된다.
- completed: TUI 첫 화면은 review 후보가 있을 때 `재진단 큐`를 표시한다.
- completed: exercise 성공 action panel은 review 후보가 있을 때 `잔류 리스크`를 표시한다.
- completed: E2E는 temp HOME progress fixture만 사용한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/review/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/playable/... ./internal/review/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`

## 저장 포맷 확인

`internal/progress/`와 progress JSON schema는 변경하지 않았다.
