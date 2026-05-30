# Midterm Todo

> 현재 중기 보드만 둔다. 과거 중기 플랜의 긴 히스토리는 `docs/roadmap/archive/history/MIDTERM_TODO_2026-05-29.md`와 `docs/exec-plans/completed/`를 본다.

## 현재 상태

Status: planning needed

현재 active slice는 없다. 가장 최근 중기 플랜인 Long Incident Evidence Hardening은 완료됐고, 다음 단계는 `PLAN-REFRESH-009`로 Foundation exit review와 다음 중기 플랜을 고정하는 것이다.

## 다음 중기 플랜 후보

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-009 | proposed | Foundation exit review와 다음 중기 플랜 확정 |
| 2 | QUOTE-PAIR-HARDEN-001 | candidate | `ci'`, `ci(`, `ci{` 등 quote/pair text object hardening |
| 3 | PLATFORM-REVIEW-003 | candidate | 저장 포맷 변경 없는 mission/review/game loop 강화 |
| 4 | CONTENT-BREADTH-002 | candidate | 기존 engine만 사용하는 applied incident/command-choice 확장 |

권장은 `PLAN-REFRESH-009`를 먼저 열고, 그 결과로 engine hardening, platform loop, content breadth 중 하나를 고르는 것이다.

## PLAN-REFRESH-009 출구 조건

| 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|-----------|-------------|------|---------------|
| Long Incident Evidence Hardening 완료 | foundation exit review, 출시 후보 범위, 다음 중기 플랜 | docs review, 최신 E2E evidence 확인 | stale roadmap을 근거로 새 engine/content를 열지 않음 |

## 다음 결정 기준

- **새 engine을 열 때**: `docs/roadmap/ENGINE_TODO.md`의 hardening candidates와 `docs/gameplay/vim-curriculum-map.md`의 coverage gaps를 먼저 본다.
- **게임성을 보강할 때**: `docs/roadmap/UX_BACKLOG_001.md`, `docs/roadmap/PLATFORM_RFC_001.md`, 최신 playtest review를 먼저 본다.
- **content를 늘릴 때**: 기존 implemented engine만 사용할 수 있는지 확인하고, long route에는 final/timeline E2E evidence를 남긴다.

## 최근 완료된 중기 플랜

| ID | 완료일 | 요약 |
|----|--------|------|
| E2E-EVIDENCE-008 | 2026-05-29 | 긴 incident와 command-choice applied route의 final/timeline evidence bundle 고정 |
| Applied Mastery Runs | 2026-05-28 | `incident-006`, quote value reuse-choice, `incident-007` 완료 |
| Inline Target Motions | 2026-05-26 | `f/t`, `df/dt/cf/ct`, tutorial, command-choice 적용 beat 완료 |
| Foundation Playtest / UX Polish | 2026-05-26 | HUD/help/choice/success modal polish와 command-choice content breadth 완료 |

## 운영 규칙

- 이 파일은 현재/다음 중기 계획만 유지한다.
- 완료된 긴 계획표는 `docs/roadmap/archive/history/`로 이동한다.
- "다음 중기 플랜" 섹션이 여러 개 쌓이면 stale 위험으로 보고 정리한다.
