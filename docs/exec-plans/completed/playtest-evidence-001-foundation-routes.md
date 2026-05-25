# PLAYTEST-EVIDENCE-001 — Foundation Route Evidence Review

Slice-ID: PLAYTEST-EVIDENCE-001
Created: 2026-05-25
Status: completed
Completed: 2026-05-26
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

- [x] 대표 E2E rerun
- [x] final/timeline/app_state evidence 위치 확인

## Step 2: Review

- [x] FTUE route UI 흐름 점검
- [x] incident route UI 흐름 점검
- [x] command-choice route UI 흐름 점검

## Step 3: Backlog

- [x] 다음 slice 후보 정리
- [x] `git diff --check`

## RedTeam 실행 결과

SubAgent RedTeam은 코드/문서 수정 없이 E2E runner로 대표 루트를 직접 실행하고 evidence를 검토했다.

실행:
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`

결과:
- 세 scenario 모두 `scenario passed`
- evidence:
  - `artifacts/e2e/playable_ftue_first_five_route/screen.txt`
  - `artifacts/e2e/playable_incident_001_full/screen.txt`
  - `artifacts/e2e/playable_command_choice_scope/screen.txt`
  - `artifacts/e2e/playable_command_choice_scope/key_trace.txt`

검증:
- `git diff --check`

## Findings

| 우선순위 | 문제 | Evidence | 판단 |
|----------|------|----------|------|
| P1 | Mission HUD의 review/daily 문구가 현재 목표보다 먼저 눈에 들어온다. | FTUE/incident screen의 `복구 현황: 재점검 대상... 오늘의 복구 루트...` | `HUD-DENSITY-001`에서 즉시 해결 |
| P1 | briefing/review 문구가 화면 폭에서 잘리거나 어색하게 끊긴다. | command-choice screen의 `오염된 줄 묶음을 골`, FTUE screen의 `외 2건` | `HUD-DENSITY-001`에서 줄바꿈/축약 정책으로 해결 |
| P1 | command-choice drill이 아직 도구 선택보다 `V j d` 암기처럼 보일 수 있다. | command-choice key trace `V`, `j`, `d`; cue `판단: 목표 상태...` | HUD 이후 `CHOICE-JUDGMENT-001` 후보로 분리 |
| P2 | success modal 내부 정보가 다소 중복된다. | `RUNBOOK SEALED`, `STEP SEALED`, `Runbook: 1/1`, `Dispatch complete` | HUD 이후 `SUCCESS-MODAL-001` 후보로 분리 |
| P2 | `?`, retry, quit 안내는 노출되지만 실제 입력 UX 검증은 부족하다. | 대표 루트는 `?`를 누르지 않음 | `HELP-AFFORDANCE-001`에서 검증 |

## 해결 순서

1. `HUD-DENSITY-001`
   - current mission title/briefing/action cue를 최우선으로 둔다.
   - review/daily는 tutorial 초반에서 더 짧게 접고, incident에서는 세계관 메타 정보로 유지하되 길이를 제한한다.
   - terminal width 기반 wrapping/truncation 정책을 renderer test로 고정한다.
2. `HELP-AFFORDANCE-001`
   - `?`, retry, quit 안내가 실제 입력 처리와 일치하는지 focused route를 추가한다.
3. `CHOICE-JUDGMENT-001`
   - command-choice에서 왜 linewise visual을 고르는지 비교 판단 cue를 강화한다.
4. `SUCCESS-MODAL-001`
   - 성공 modal의 중복 heading/record density를 줄인다.
