# PLAN-REFRESH-006 — Foundation Milestone Health Check

Slice-ID: PLAN-REFRESH-006
Created: 2026-05-25
Status: completed
Completed: 2026-05-25
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/active/ftue-001-first-five-minute-route.md
- docs/exec-plans/completed/plan-refresh-006-foundation-milestone.md

## 목표

UI-HUD-003 이후의 중간 점검 결과를 반영해 `Vim Learning Foundation`의 다음 2~3주 실행 순서를 고정한다. 새 Vim 엔진 기능을 추가하기보다 첫 플레이 루프, HUD playtest, command-choice 적용, no-schema daily route, progress v2 결정 게이트 순서로 제품 milestone을 묶는다.

## 입력

- 사용자 결정: 권장 방향으로 진행한다.
- SubAgent 종합:
  - 새 Vim 기능 추가보다 milestone packaging이 우선이다.
  - 첫 5분 플레이 루프가 아직 미확인이다.
  - UI-HUD-003은 좋은 기본 구조지만 playtest/polish가 필요하다.
  - command-choice drill은 docs-only에서 playable로 승격할 때다.
  - progress schema v2는 no-schema daily loop를 먼저 확인한 뒤 결정한다.

## 결정

다음 중기 플랜은 아래 순서로 고정한다.

1. `FTUE-001`: First 5-Minute Route
2. `UI-PLAYTEST-001`: HUD Usability Pass
3. `CHOICE-PLAY-001`: Command Choice Drill Playable
4. `DAILY-ROUTE-001`: No-Schema Daily Loop
5. `PROGRESS-V2-DECISION-001`: Progress v2 Decision Gate

## 완료 기준

- `PROGRAM.md`의 active slice가 낡은 completed slice에서 `FTUE-001`로 갱신된다.
- `MIDTERM_TODO.md`에 다음 2~3주 중기 플랜과 Foundation exit criteria가 기록된다.
- `FTUE-001` ExecPlan이 active로 생성되어 다음 작업을 바로 시작할 수 있다.
- 저장 포맷 변경은 이번 중기 플랜의 마지막 결정 게이트 전까지 보류한다.

## 검증 결과

- `git diff --check`

## 후속 작업

`FTUE-001`을 실행해 현재 첫 실행 루트를 evidence 기반으로 점검하고, 5분 안에 `이동 → 생존 → 빠른 이동 → 작은 수정`의 효용을 느낄 수 있는 canonical route를 고정한다.
