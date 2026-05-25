# PLAYTEST-EVIDENCE-001 — Foundation Route Evidence Review

Slice-ID: PLAYTEST-EVIDENCE-001
Created: 2026-05-25
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/verification/tui-e2e-loop.md
- docs/verification/tui-ui-qa-contract.md
- docs/exec-plans/active/playtest-evidence-001-foundation-routes.md
- docs/exec-plans/completed/ux-language-001-route-action-copy.md
- test/e2e/
- artifacts/e2e/

## 목표

FTUE, incident, command-choice 대표 루트의 E2E evidence를 다시 보고, 다음 HUD/UX polish가 필요한 지점을 evidence 기반 backlog로 정리한다.

## 범위

- 포함:
  - `playable_ftue_first_five_route`
  - `playable_incident_001_full` 또는 최신 incident 대표
  - `playable_command_choice_scope`
  - screen timeline/final screen/app_state evidence 확인
  - 다음 slice 후보 정리
- 제외:
  - renderer 구현 변경
  - content target/optimal key 변경
  - progress 저장 포맷 변경

## 수용 기준

- 대표 route evidence가 최신 action copy를 반영한다.
- 어색함은 주관적 감상만이 아니라 evidence 파일과 연결해 기록한다.
- 다음 slice 후보가 `HUD-DENSITY-001`, `HELP-AFFORDANCE-001`, 또는 별도 bugfix로 분류된다.
- 필요한 E2E만 다시 실행하고 `git diff --check`를 통과한다.

## Step 1: Evidence Generation

- [ ] 대표 E2E rerun
- [ ] final/timeline/app_state evidence 위치 확인

## Step 2: Review

- [ ] FTUE route UI 흐름 점검
- [ ] incident route UI 흐름 점검
- [ ] command-choice route UI 흐름 점검

## Step 3: Backlog

- [ ] 다음 slice 후보 정리
- [ ] `git diff --check`
