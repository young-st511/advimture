# PLAN-REFRESH-009 — Foundation Forward Plan

Slice-ID: PLAN-REFRESH-009
Created: 2026-05-30
Completed: 2026-05-30
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/active/plan-refresh-009-foundation-forward-plan.md
- docs/exec-plans/completed/plan-refresh-009-foundation-forward-plan.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- docs/README.md

## 목표

문서 cleanup 이후 앞으로의 계획을 계속 참고할 수 있는 rolling plan으로 고정한다. `PROGRAM.md`는 현재 상태, `MIDTERM_TODO.md`는 실행 보드, `FORWARD_PLAN.md`는 2~8주 방향과 운영 원칙을 맡도록 역할을 분리한다.

## 범위

- 포함:
  - Foundation exit 기준과 다음 2~8주 계획 문서화
  - 다음 작업 우선순위와 decision gate 정리
  - 앞으로 문서를 계속 참고/갱신하는 운영 규칙 연결
- 제외:
  - 코드 변경
  - 새 engine/content 구현
  - progress schema 변경
  - release packaging 구현

## 수용 기준

- `docs/roadmap/FORWARD_PLAN.md`가 현재 권장 순서와 장기 후보를 설명한다.
- `PROGRAM.md`와 `MIDTERM_TODO.md`가 새 forward plan을 참조한다.
- `docs/README.md`가 forward plan의 역할과 업데이트 시점을 설명한다.
- 다음 작업은 `FORWARD_PLAN.md`를 먼저 확인하도록 문서화된다.

## 검증 계획

- `git diff --check`
- `rg`로 `FORWARD_PLAN.md` 참조 확인

## Step 1: Rolling Plan

- [x] `FORWARD_PLAN.md` 작성
- [x] Foundation exit criteria와 next sequence 정리

## Step 2: Entry Point Link

- [x] `PROGRAM.md` 참조 연결
- [x] `MIDTERM_TODO.md` 참조 연결
- [x] `docs/README.md` 문서별 안내 갱신

## Step 3: Complete

- [x] `CHANGES.md` 기록
- [x] 검증 후 completed 이동

## 검증 결과

- `rg -n "FORWARD_PLAN|FOUNDATION-EXIT-001|PLAN-REFRESH-009" docs/roadmap docs/README.md`: pass after stale reference sync
- `git diff --check`: pass

## 의사결정 로그

- 2026-05-30: `PROGRAM.md`/`MIDTERM_TODO.md`를 다시 비대하게 만들지 않기 위해 별도 rolling forward plan을 둔다.
- 2026-05-30: 다음 실제 작업명은 `FOUNDATION-EXIT-001`로 두고, `PLAN-REFRESH-009`는 forward plan 문서화 slice로 완료한다.

## 미해결 질문

- 없음.
