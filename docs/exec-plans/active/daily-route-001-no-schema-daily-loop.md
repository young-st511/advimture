# DAILY-ROUTE-001 — No-Schema Daily Loop

Slice-ID: DAILY-ROUTE-001
Created: 2026-05-25
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/exec-plans/active/daily-route-001-no-schema-daily-loop.md
- docs/exec-plans/completed/choice-play-001-command-choice-playable.md
- internal/review/
- internal/playable/
- internal/e2estate/
- test/e2e/

## 목표

progress v1 저장 포맷을 바꾸지 않고, 이미 있는 review queue와 content library만으로 “오늘 다시 들어올 이유”를 더 선명하게 만든다.

## 범위

- 포함:
  - daily route 문구를 단순 개수에서 primary review 대상과 command 영역이 드러나는 짧은 요약으로 개선
  - app_state/e2e summary에서 daily route 문구 검증
  - focused E2E와 unit test
- 제외:
  - `internal/progress/` 저장 JSON 구조 변경
  - mastery score, spaced review due date, streak 저장
  - 새 content schema

## 수용 기준

- 완료하지 않은 mission이 있으면 daily route가 primary mission과 reason을 포함한다.
- 모든 mission이 완료된 경우에도 기존처럼 빈/완료 상태를 안정적으로 표현한다.
- UI 문구는 `복구국` 세계관에 맞되 command 학습 목표를 가리지 않는다.
- E2E state summary 또는 focused E2E가 daily route 문구를 검증한다.
- `go test ./...`, focused E2E, `git diff --check`를 통과한다.

## Step 1: Current Contract Review

- [ ] review queue/daily route 현재 구현 확인
- [ ] app_state summary assertion 확인

## Step 2: No-Schema Implementation

- [ ] daily route copy/summary 개선
- [ ] unit test 갱신
- [ ] focused E2E 갱신 또는 추가

## Step 3: Verification

- [ ] focused E2E
- [ ] `go test ./...`
- [ ] `git diff --check`
