# PLAN-REFRESH-003 — Applied Learning and World Frame

Slice-ID: PLAN-REFRESH-003
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/completed/plan-refresh-003-applied-world-frame.md
- docs/gameplay/world-frame.md
- docs/gameplay/scenario-tone.md
- docs/roadmap/PRODUCT.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- docs/roadmap/decisions/0005-world-frame-runbook-dispatch.md
- docs/README.md

## 목표

중간 점검과 세계관 아이디에이션 결과를 반영해 다음 중기 플랜을 고정한다.

## 결정

- 세계관 기본 프레임은 `원격 시설 복구국 / Runbook Dispatch`로 한다.
- 개별 incident에는 `침묵한 릴레이 기지` 감각을 얇게 섞는다.
- 다음 구현 루프는 새 Vim command 추가가 아니라 incident UX 보강으로 시작한다.
- visual mode는 selection contract와 E2E assertion 확장 이후 구현한다.

## 수용 기준

- completed: world frame decision이 문서화된다.
- completed: scenario tone 문서가 새 프레임을 참조한다.
- completed: 다음 중기 플랜이 `MIDTERM_TODO.md`에 기록된다.
- completed: `PROGRAM.md` 활성 slice가 다음 실행 대상으로 갱신된다.

## 검증 결과

- passed: `git diff --check`

