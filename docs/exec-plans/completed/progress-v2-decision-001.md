# PROGRESS-V2-DECISION-001 — Progress Schema v2 Decision

Slice-ID: PROGRESS-V2-DECISION-001
Created: 2026-05-25
Status: completed
Completed: 2026-05-25
Scope-Mode: docs-only
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PLATFORM_RFC_001.md
- docs/gameplay/spec.md
- docs/exec-plans/active/progress-v2-decision-001.md
- docs/exec-plans/completed/daily-route-001-no-schema-daily-loop.md

## 목표

FTUE, command-choice, no-schema daily loop 결과를 바탕으로 progress schema v2를 지금 승인할지, 더 많은 evidence 전까지 보류할지 결정한다.

## 범위

- 포함:
  - progress v1로 가능한 반복 학습 범위 정리
  - v2가 필요한 조건과 최소 승인안 정리
  - 다음 중기 플랜에 반영할 decision 작성
- 제외:
  - `internal/progress/` 코드/JSON 변경
  - migration 구현
  - mastery/spaced review runtime 구현

## 수용 기준

- v1로 충분한 범위와 부족한 범위가 분리되어 있다.
- v2 승인/보류 결정이 `docs/roadmap/PLATFORM_RFC_001.md` 또는 본 ExecPlan에 명시된다.
- 결정은 Foundation Product Loop의 다음 중기 플랜 입력으로 쓸 수 있다.
- `git diff --check`를 통과한다.

## Step 1: Evidence Review

- [x] FTUE/command-choice/daily 결과 확인
- [x] progress v1 한계 정리

## Step 2: Decision

- [x] v2 승인/보류 판단
- [x] 다음 플랜 입력 작성

## Step 3: Verification

- [ ] `git diff --check`

## 결정

Progress schema v2는 지금 구현하지 않고 보류한다.

현재 evidence:
- FTUE route는 exercise completion, best grade, best keystrokes만으로 충분히 검증된다.
- command-choice playable도 새 저장 필드 없이 progress v1 mission completion으로 연결된다.
- daily route는 content library와 review queue 계산만으로 primary 대상과 이유를 보여줄 수 있다.

v1로 충분한 범위:
- 첫 실행 route 이어가기
- exercise별 완료 여부와 best record 표시
- incomplete, low grade, key count 기반 review 후보 계산
- 저장하지 않는 daily route 추천

v2가 필요해지는 조건:
- 실패했지만 성공하지 않은 attempt를 다음 세션 복구 후보로 유지해야 할 때
- command cluster별 mastery level을 저장해야 할 때
- last reviewed/due date 기반 spaced review가 필요할 때
- daily streak/history가 제품 핵심 동기가 될 때

다음 플랜 입력:
- 당장은 progress v2보다 Foundation playtest/UX polish와 content breadth를 우선한다.
- v2를 다시 열 때는 `schema_version`, `review`, `mastery`, `daily_runs` 중 `review` 최소 필드부터 별도 승인한다.
- `internal/progress/` 변경은 계속 사용자 승인 필수 boundary로 유지한다.

## 검증 결과

- `git diff --check`
