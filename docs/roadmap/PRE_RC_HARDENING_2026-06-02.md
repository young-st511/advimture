# PRE-RC Hardening Review — 2026-06-02

## 목적

`RELEASE-CANDIDATE-001`로 첫 공개 후보를 묶기 전에, 기존 playable loop의 약한 지점을 evidence 기준으로 한 번 더 확인한다.

## Evidence

실행한 대표 route:

- `playable_ftue_first_five_route`
  - `artifacts/e2e/playable_ftue_first_five_route/screen_final.txt`
  - `artifacts/e2e/playable_ftue_first_five_route/screen_timeline.txt`
  - `artifacts/e2e/playable_ftue_first_five_route/app_state.json`
- `playable_incident_hint_affordance`
  - `artifacts/e2e/playable_incident_hint_affordance/screen_final.txt`
  - `artifacts/e2e/playable_incident_hint_affordance/screen_timeline.txt`
  - `artifacts/e2e/playable_incident_hint_affordance/app_state.json`
- `playable_incident_001_full`
  - `artifacts/e2e/playable_incident_001_full/screen_final.txt`
  - `artifacts/e2e/playable_incident_001_full/screen_timeline.txt`
  - `artifacts/e2e/playable_incident_001_full/app_state.json`
- `playable_review_queue`
  - `artifacts/e2e/playable_review_queue/screen_final.txt`
  - `artifacts/e2e/playable_review_queue/screen_timeline.txt`
  - `artifacts/e2e/playable_review_queue/app_state.json`

## Findings

### P0/P1 blocker

없음.

- 첫 5분 route는 tutorial 0~3 첫 문항까지 막힘 없이 진행된다.
- 대표 incident 001 route는 검색, 단어 변경, open line, yank/put, range substitute를 하나의 runbook으로 끝까지 진행한다.
- review queue 성공 debrief는 `잔류 리스크`, `다음 출격`, `Next` action을 유지한다.

### P2 hardening

1. Incident hint cue가 한 줄에 몰려 120폭에서도 긴 hint가 잘릴 수 있었다.
   - 조치: running mission cue를 terminal width 기준으로 여러 줄에 감싸도록 변경했다.
   - 검증: `playable_incident_hint_affordance`를 80x30 viewport smoke로 낮추고, `원인 신호를`, `잡습니다.`, `참고 명령: /`, `ui.focus_panel` evidence를 검증한다.

2. 긴 incident full route 일부가 `app_state.json` evidence를 저장하지 않았다.
   - 조치: incident 001/002/003/004/006/007 full route에 `save_app_state: true`를 추가했다.
   - 검증: `playable_incident_001_full` focused run에서 `app_state.json` 생성과 success focus panel을 확인했다. 전체 route는 `make release-check`에서 다시 검증한다.

## 판정

현재 hardening 이후에도 새 engine/content/schema/progress 저장 포맷 변경은 없다. 전체 release gate를 통과했고, `PRE-RC-HARDENING-001`은 completed로 이동했다. 다음 권장 slice는 `RELEASE-CANDIDATE-001`이다.
