# PROGRESS-V2-DECISION-001 — Progress Schema v2 Decision

Slice-ID: PROGRESS-V2-DECISION-001
Created: 2026-05-25
Status: active
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

- [ ] FTUE/command-choice/daily 결과 확인
- [ ] progress v1 한계 정리

## Step 2: Decision

- [ ] v2 승인/보류 판단
- [ ] 다음 플랜 입력 작성

## Step 3: Verification

- [ ] `git diff --check`
