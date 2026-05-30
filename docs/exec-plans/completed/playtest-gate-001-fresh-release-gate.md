# ExecPlan: PLAYTEST-GATE-001 Fresh Playtest Release Gate

Status: completed
Date: 2026-05-30
Review: `docs/roadmap/PLAYTEST_GATE_2026-05-30.md`

## Goal

README 기준으로 처음 실행하는 플레이어 관점에서 Advimture release candidate를 점검하고, 출시 전 blocker와 후속 polish wishlist를 분리한다.

## Context

- `RELEASE-READINESS-001`로 README, `make build`, `make test`, `make release-check`가 정리됐다.
- Foundation engine, tutorial/incident content, E2E loop는 통과 중이다.
- 다음 위험은 기능 부족보다 fresh run에서 보이는 막힘, UI clipping, 안내 오해, incident fatigue다.

## Scope

포함:

- README 실행 흐름과 release gate 확인
- 첫 5분 canonical route evidence 재검토
- tutorial 확장 구간 대표 route evidence 재검토
- incident 3개 이상 final/timeline evidence 재검토
- debrief/review loop evidence 재검토
- terminal width/height 차이에 따른 modal/HUD clipping smoke 필요성 판단
- blocker/P2 polish/post-release 후보 분류
- 필요 시 후속 bugfix ExecPlan 후보 작성

제외:

- 새 Vim engine 기능
- 새 content/playpack 추가
- 저장 포맷 변경
- content schema 변경
- UI 대개편
- 새 의존성 추가

## Review Rubric

| 등급 | 의미 | 처리 |
|------|------|------|
| P0 | 실행/진행/저장/검증이 막힌다 | 즉시 blocker fix slice |
| P1 | 첫 플레이어가 필수 조작을 오해하거나 action을 못 본다 | release 전 fix slice |
| P2 | 게임감/문구/밀도 개선이 필요하지만 진행은 가능하다 | first-run polish 후보 |
| P3 | 출시 후 개선해도 되는 확장/취향/세계관 보강 | post-MVP backlog |

## Acceptance Criteria

- [x] `make release-check`가 통과한다.
- [x] README 실행/검증 안내가 실제 명령과 맞는지 확인한다.
- [x] 첫 5분 route evidence를 검토하고 P0/P1/P2/P3로 분류한다.
- [x] tutorial 확장 route evidence를 최소 3개 이상 검토한다.
- [x] incident full route evidence를 최소 3개 이상 검토한다.
- [x] review/debrief route evidence를 검토한다.
- [x] terminal UI clipping 위험을 확인하고 필요 시 viewport smoke 후보를 남긴다.
- [x] `docs/roadmap/PLAYTEST_GATE_2026-05-30.md`에 결과를 기록한다.
- [x] P0/P1이 있으면 다음 active 후보를 blocker fix로 둔다. 없으면 first-run polish 또는 release candidate prep으로 둔다.

## Verification

- `env PATH=/opt/homebrew/bin:$PATH make release-check`: pass
- Evidence spot review:
  - `artifacts/e2e/playable_ftue_first_five_route/screen_timeline.txt`
  - `artifacts/e2e/playable_review_queue/screen_timeline.txt`
  - `artifacts/e2e/playable_hjkl_success/screen_final.txt`
  - `artifacts/e2e/playable_text_object_quote_pair_full/screen_final.txt`
  - `artifacts/e2e/playable_command_choice_scope/screen_final.txt`
  - `artifacts/e2e/playable_incident_001_full/screen_final.txt`
  - `artifacts/e2e/playable_incident_002_full/screen_final.txt`
  - `artifacts/e2e/playable_incident_003_full/screen_final.txt`
  - `artifacts/e2e/playable_incident_007_full/screen_final.txt`

참고: Codex sandbox에서는 `go build` 중 사용자 Go module stat cache 쓰기 경고가 한 줄 출력됐지만 exit code는 0이었다.

## Result

- P0/P1 blocker는 확인되지 않았다.
- 다음 권장 slice는 `FIRST-RUN-POLISH-001`이다.
- P2 후보는 running cue density, review/daily line length, quote 주변 cursor marker 혼동, mid tutorial evidence 편차, viewport smoke 부재다.
