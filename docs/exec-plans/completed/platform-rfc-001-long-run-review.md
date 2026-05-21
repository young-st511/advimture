# PLATFORM-RFC-001 — Long-Run Review Platform RFC

Slice-ID: PLATFORM-RFC-001
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/platform-rfc-001-long-run-review.md
- docs/roadmap/PLATFORM_RFC_001.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- docs/gameplay/spec.md

## 목표

Utility command playpack 이후 Advimture를 장기 반복 학습 플랫폼으로 키우기 위한 mastery, spaced review, daily run, progress schema 후보를 문서화한다.

## 범위

- 포함:
  - 현재 progress 저장 모델에서 가능한 read-only review 기능 정리
  - 저장 포맷 변경이 필요한 기능과 승인 게이트 정리
  - progress schema v2 후보 구조 제안
  - 다음 구현 루프 후보 분리
- 제외:
  - `internal/progress/` 코드 변경
  - 저장 JSON schema 변경
  - daily run UI 구현
  - review engine 구현

## 수용 기준

- completed: 저장 포맷 변경 없이 가능한 기능과 불가능한 기능을 분리한다.
- completed: mastery, recovery, spaced review, daily run의 첫 기준을 정의한다.
- completed: schema v2 후보는 제안으로만 두고 구현 승인으로 해석되지 않게 한다.
- completed: 다음 구현 루프를 저장 변경 없는 루프와 schema 변경 루프로 나눈다.

## 검증 결과

- passed: 문서 범위 diff review
- passed: `git diff --check`

코드 변경이 없으므로 Go test는 최종 묶음 검증에서 기존 회귀 확인 용도로만 실행한다.
