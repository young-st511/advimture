# NEXT Playtest Review — 2026-06-07

## 판정

P0/P1 blocker는 보이지 않는다. `SEARCH-THEN-SCOPE-APPLIED-001 -> BRACKET-PAIR-HARDEN-001` 이후 playable loop는 content/engine/E2E 기준으로 안정적이다.

다음 최선은 새 Vim command나 release candidate prep이 아니라 **`REVIEW-LOOP-MOTIVATION-001`**이다.

목표는 success debrief, 잔류 리스크, 오늘의 복구 루트, 다음 출격 언어를 다듬어 “다음에 무엇을 반복해야 하는지”가 더 자연스럽게 읽히게 만드는 것이다.

## Evidence Inventory

| 축 | Evidence | 확인 결과 |
|----|----------|-----------|
| First tour | `artifacts/e2e/playable_ftue_first_five_route/screen_final.txt` | 첫 이동/생존/빠른 이동/작은 수정 흐름은 통과한다. 성공 modal은 grade, 목표 입력, 다음 단계, 잔류 리스크를 안정적으로 보여준다. |
| Bracket pair tutorial | `artifacts/e2e/playable_bracket_pair_full/screen_final.txt` | `di(`, `ci(`, `yi(`, `di{`, `ci{`, `yi{` 6문항이 한 playlist로 닫힌다. 최종 success copy는 delimiter 보존과 register 재사용 이유를 설명한다. |
| First dispatch arc | `artifacts/e2e/playable_incident_001_full/screen_final.txt`, `playable_incident_002_full`, `playable_incident_003_full` | 001 원인 신호 추적, 002 구조 재동기화, 003 오염 구간 격리 흐름이 final evidence에서 읽힌다. |
| Judgment drill | `artifacts/e2e/playable_command_choice_scope/screen_final.txt` | command-choice 6 beat는 scope/range/inline/reuse/repeat/line reuse 판단을 설명한다. 다음 runbook action도 안정적이다. |
| Search then scope | `artifacts/e2e/playable_incident_008_full/screen_final.txt` | `/breach`로 marker를 찾고 linewise scope를 판단하는 1-beat incident가 명확하다. 마지막 incident로 `출격 완료`를 표시한다. |
| Review loop | `artifacts/e2e/playable_review_queue/app_state.json` | review queue, primary exercise, daily route, focus panel action id는 안정적이다. 다만 success modal의 `잔류 리스크`와 `다음 출격` 언어가 tutorial 중에도 반복되어 약간 기계적으로 읽힌다. |

## Option Review

| Option | 장점 | 리스크 | 판정 |
|--------|------|--------|------|
| Command-choice breadth | Advimture 정체성인 “상황에 맞는 Vim 도구 선택”을 더 강화한다. | 현재 `incident-005` 6 beat와 `incident-008`이 이미 충분히 판단 훈련을 보여준다. 바로 늘리면 route가 무거워질 수 있다. | 다음 후보로 보류 |
| Tutorial pacing / mid-tour compression | tutorial이 길어지는 문제를 직접 다룬다. | playlist 순서와 학습 곡선을 건드릴 가능성이 있어 review보다 크다. | 후속 playtest 후보 |
| Review-loop motivation polish | schema/engine/content 추가 없이 반복 플레이 동기를 강화한다. 현재 evidence에서 가장 자주 노출되는 어색함이다. | copy/UI wording contract를 건드리므로 focused E2E가 필요하다. | 추천 |
| Deeper pair hardening | `a(`, nested pair, escaped delimiter 등 실용성을 더 늘린다. | engine scope가 커지고 지금 evidence가 요구하지 않는다. | later |
| Release candidate prep | 출시 준비 문서와 tag 후보를 묶을 수 있다. | 사용자가 바로 출시할 생각은 없다고 했다. | 열지 않음 |

## Recommended Slice

### REVIEW-LOOP-MOTIVATION-001 — Review Loop Motivation Polish

목표:

- success modal에서 primary action과 review motivation을 더 명확히 분리한다.
- tutorial 중에는 `다음 출격`이 실제 action처럼 읽히지 않도록 review/daily route copy를 조정한다.
- incident 완료와 마지막 dispatch에서는 `잔류 리스크`/`다음 출격` 언어를 유지하되, 반복해야 하는 이유가 더 플레이어 친화적으로 보이게 한다.

포함:

- `FocusPanel` success lines의 review/daily wording spot polish
- tutorial success와 incident success의 review label 차이 검토
- `playable_review_queue`, `playable_debrief_success`, `playable_ftue_first_five_route`, viewport modal E2E 확인
- `docs/gameplay/spec.md`, `docs/gameplay/tui-screen-contract.md`, `docs/roadmap/UX_BACKLOG_001.md` 동기화

제외:

- progress schema 변경
- content schema 변경
- 새 Vim command/engine capability
- 새 dependency
- release candidate/tag 작업

완성 기준:

- success modal에서 primary action과 review queue 안내가 서로 다른 의미로 읽힌다.
- tutorial 중 review 안내는 “지금 누를 다음 버튼”이 아니라 “나중에 재점검할 대상”으로 읽힌다.
- incident success는 Runbook Dispatch 톤을 유지한다.
- app_state `ui.focus_panel.actions[].id` 계약은 유지한다.
- `go test ./internal/playable ./internal/playableview ./internal/e2estate`와 `go test ./...`가 통과한다.
- focused E2E: `playable_review_queue`, `playable_debrief_success`, `playable_ftue_first_five_route`, `playable_viewport_success_modal_80x24`.

## Not Now

- `COMMAND-CHOICE-BREADTH-002`: 다음 content 후보로 좋지만, 이번 evidence에서는 review motivation보다 덜 급하다.
- `PAIR-HARDENING-LATER`: nested/escaped/around/count는 engine risk가 크므로 실제 콘텐츠 병목이 생긴 뒤 연다.
- `RELEASE-CANDIDATE-001`: 출시 가능한 수준으로 개발하는 목표와 실제 공개 후보를 묶는 작업은 분리한다.
