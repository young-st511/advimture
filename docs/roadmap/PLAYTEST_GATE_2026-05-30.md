# Playtest Gate — 2026-05-30

Status: completed
ExecPlan: `docs/exec-plans/completed/playtest-gate-001-fresh-release-gate.md`

## 판정

Release candidate blocker는 확인되지 않았다.

- P0: 없음
- P1: 없음
- P2: first-run polish 후보 있음
- P3: post-MVP backlog 후보 있음

따라서 다음 권장 slice는 `RC-BLOCKER-FIX-001`이 아니라 `FIRST-RUN-POLISH-001`이다.

## 확인한 Evidence

Release gate:

- `make release-check`: pass
  - `make test`: pass
  - `make build`: pass
  - `make e2e-playable`: pass

대표 evidence:

- `artifacts/e2e/playable_ftue_first_five_route/screen_timeline.txt`
- `artifacts/e2e/playable_review_queue/screen_timeline.txt`
- `artifacts/e2e/playable_hjkl_success/screen_final.txt`
- `artifacts/e2e/playable_text_object_quote_pair_full/screen_final.txt`
- `artifacts/e2e/playable_command_choice_scope/screen_final.txt`
- `artifacts/e2e/playable_incident_001_full/screen_final.txt`
- `artifacts/e2e/playable_incident_002_full/screen_final.txt`
- `artifacts/e2e/playable_incident_003_full/screen_final.txt`
- `artifacts/e2e/playable_incident_007_full/screen_final.txt`

## P0/P1 Blocker Review

| 축 | 판정 | 근거 |
|----|------|------|
| README 실행성 | pass | `make release-check`가 README의 공개 전 검증 흐름과 일치한다. |
| 첫 5분 | pass | FTUE route가 tutorial 0~3 초반까지 진행되고 success/debrief/next tutorial이 보인다. |
| tutorial 확장 | pass | open-line, repeat, search, quote, visual, char-find full route가 통과한다. |
| incident 감각 | pass | incident 001/002/003/005/007 final screen이 runbook 복구 완료 구조로 읽힌다. |
| 진행/복습 | pass | debrief가 `이번 복구`, `최단 복구`, `잔류 리스크`, `다음 출격`을 보여준다. app_state review도 유지된다. |
| 저장 안전성 | pass | E2E는 temp HOME을 사용하고 실제 progress를 건드리지 않는다. |

## P2 First-Run Polish 후보

### 1. Running cue line 밀도

첫 tutorial running 화면의 cue line이 한 줄에 많은 정보를 싣는다.

예:

```text
TRAINING BRIEF · Inputs left: 2/2 · 기억할 명령: l · Coach: 훈련 키 l · ?: hint  q: quit
```

진행은 가능하지만 첫 플레이어가 `기억할 명령`, `Coach`, `?: hint`, `q: quit`을 동시에 받아야 해서 밀도가 높다.

후보:

- tutorial running cue를 2개 의미 단위로 나누되 높이 흔들림은 만들지 않는다.
- `기억할 명령`과 `Coach` 중복을 줄인다.
- incident 기본 judgment cue는 정답 노출 없이 유지한다.

### 2. Review/daily line 길이

success 화면 상단의 `재점검 대상 + 오늘의 복구 루트` line은 세계관적으로 좋지만 긴 제목에서 시선을 빼앗을 수 있다.

후보:

- success modal 안의 `잔류 리스크`/`다음 출격`이 이미 핵심을 담으므로, 상단 line은 running 화면 위주로 축약한다.
- app_state review assertion은 유지한다.

### 3. Cursor rendering과 quote 문자 혼동

quote 주변에서 cursor marker가 `["]`처럼 보이며, 처음 보는 플레이어에게 실제 buffer 문자처럼 읽힐 수 있다.

후보:

- cursor rendering 자체를 바꾸기보다, quote/text-object tutorial의 문구에서 "대괄호는 커서 표시"를 초반 한 번만 자연스럽게 알려준다.
- 더 큰 renderer 변경은 post-MVP로 둔다.

### 4. Tutorial full route evidence 편차

긴 incident와 command-choice는 `screen_final.txt`/`screen_timeline.txt` evidence가 잘 남지만, 일부 mid tutorial full route는 `screen.txt` 중심이다.

후보:

- FIRST-RUN-POLISH 또는 QA slice에서 open-line/repeat/search/visual/char-find 대표 fixture에 final/timeline evidence를 추가한다.

### 5. Viewport smoke 부재

현재 release gate는 주로 기존 fixture terminal size를 따른다. 80x24 같은 작은 권장 하한 근처에서 modal/action line이 잘리는지 별도 smoke가 있으면 더 안전하다.

후보:

- `playable_viewport_success_modal_80x24`
- `playable_viewport_failure_modal_80x24`
- assertion은 화면 문구보다 `ui.focus_panel` app_state와 action line 존재를 함께 본다.

## P3 Post-MVP 후보

- `i(`, `i{` bracket pair hardening
- line reuse choice
- search-then-act incident
- progress schema v2 / spaced review
- terminal cell-grid parser

## 다음 권장 Slice

`FIRST-RUN-POLISH-001`

목표:

- 새 engine/content/schema 없이 첫 실행 tutorial cue, review/daily line, viewport smoke evidence만 좁게 다듬는다.

제외:

- progress 저장 포맷 변경
- 새 Vim command
- 새 incident run
- renderer 대개편
