# UX Backlog 001 — Foundation Polish Aftercare

Created: 2026-05-26
Status: current
Last reviewed: 2026-06-20

## 판단

`UI-MODAL-ACTION-HIERARCHY-001`은 완료됐고, `v0.2.0`도 배포됐다. 배포 후 동일 기준으로 UX 재검토한 결과 확인된 modal placement, running hint/quit utility wrapping, 한글 final-frame evidence, action language/evidence metadata follow-up은 `UI-MODAL-ACTION-HIERARCHY-002`로 닫았다.

현재 Foundation은 current mission, console, failed/succeeded floating modal, hint/retry/quit affordance, review/daily motivation, command memory cue를 E2E로 검증한다. 사용자 스크린샷과 SubAgent 감사로 확인된 modal/action hierarchy P1 blocker와 post-release follow-up은 현재 닫힌 상태다.

## P0

없음.

P0로 승격하는 조건:

- visible affordance가 실제 입력과 다시 어긋난다.
- failed/succeeded modal의 다음 행동 line이 잘려 플레이 진행이 막힌다.
- 새 콘텐츠가 `app_state` 없이 화면 문자열만으로만 검증되어 E2E가 불안정해진다.

## P1

### UI-MODAL-ACTION-HIERARCHY-002 — Modal/Hint Post-release Hardening

- 상태: completed
- 왜 중요한가: `001` 이후에도 modal이 buffer/status 뒤에 붙은 블록처럼 읽히는 final viewport가 있고, hint utility action과 한글 evidence reconstruction도 사람이 보는 품질 기준에서 blind spot이 남았다.
- 완료 내용: 80x24 success/failure/hint evidence가 modal placement, action footer, hint utility line, 한글 readability를 검증한다.
- 검증: `go test ./...`, `make e2e-playable`, `git diff --check`.

### UI-RAIL-001 — Wide Layout Side Rail

- 왜 중요한가: 콘텐츠가 길어질수록 review/daily, command memory, route progress를 한 줄 HUD에 계속 압축하기 어렵다.
- 언제 열 것인가: incident가 6개 이상으로 늘거나 daily/review 표시가 현재 mission briefing을 다시 밀어낼 때.
- 어떻게 검증할 것인가: desktop width에서 console line index가 고정되고, mobile/narrow width에서는 기존 vertical HUD로 fallback하는 renderer tests와 E2E screenshot/text evidence를 추가한다.

### UI-COMMAND-MEMORY-001 — Learned Command Memory

- 왜 중요한가: 콘텐츠가 늘면 플레이어가 방금 배운 command를 기억하기 어렵다. 다만 기본 화면에 command list를 크게 노출하면 판단 훈련을 약화할 수 있다.
- 상태: completed as `UI-POLISH-002` minimal slice.
- 결과: tutorial에서는 `기억할 명령`, incident에서는 hint/failure 후 `참고 명령`으로 command memory를 점진 공개한다.
- 후속: command memory가 길어지거나 command-choice가 여전히 찍기처럼 보이면 side rail 또는 command history로 확장한다.

### UI-FRAME-TIMELINE-001 — Step Frame Evidence

- 왜 중요한가: `screen_final.txt`와 `screen_timeline.txt`는 현재 충분하지만, 복잡한 modal transition을 다룰 때 send/wait 단위 frame이 필요할 수 있다.
- 언제 열 것인가: final/timeline evidence로는 특정 입력 직후의 transient UI 회귀를 설명하지 못할 때.
- 어떻게 검증할 것인가: `frames/*.txt`를 선택적으로 저장하고 summary에 frame count/path를 남긴다.

## P2

### UI-STYLE-001 — Color/Emphasis Pass

- 왜 중요한가: 현재 ASCII fallback은 검증에 좋지만, 실제 플레이 감각은 아직 개발 UI에 가깝다.
- 언제 열 것인가: content breadth 2회 이상 완료 후 주요 flow가 안정되면.
- 어떻게 검증할 것인가: color 없는 환경에서도 의미가 보존되는 snapshot test를 유지하고, style-only 변경은 app_state 계약을 변경하지 않는다.

### UI-PRESTART-001 — Briefing Modal Before Runtime Input

- 왜 중요한가: 장기적으로 tutorial episode 시작 전 “첫 투어/회상” 느낌을 줄 수 있다.
- 언제 열 것인가: episode 단위 목표 설명이 mission HUD 두 줄로 부족해질 때.
- 어떻게 검증할 것인가: modal open 상태의 `enter`는 Vim key trace에 들어가지 않아야 하며, E2E key_trace와 app_state로 분리 검증한다.

### UI-INPUT-LOG-001 — Recent Input Trace

- 왜 중요한가: 실패 원인을 스스로 복기하는 데 도움이 된다.
- 언제 열 것인가: efficiency/review run에서 key count 개선을 본격적으로 다룰 때.
- 어떻게 검증할 것인가: progress 저장 변경 없이 runtime key_trace tail만 표시하고, E2E key_trace와 화면 표시가 일치하는지 확인한다.

### UI-ACTION-LANGUAGE-001 — Action Line Language Contract

- 상태: completed by `UI-CONSOLE-POLISH-001` on 2026-06-06.
- 왜 중요했는가: success/failure modal의 action line은 한때 화면 표시와 E2E priority marker를 겸했다.
- 완료 내용: 화면 표시 label은 `다시 시도`, `다음 단계`, `다음 런북`, `출격 완료`처럼 제품 톤에 맞췄고, E2E는 `app_state.ui.focus_panel.actions[].id`로 action 의미를 검증한다.
- 검증: `playable_viewport_success_modal_80x24`, `playable_viewport_failure_modal_80x24`, `playable_incident_hint_affordance`, `playable_command_mismatch_feedback`.

### UI-REVIEW-MOTIVATION-001 — Success Review Motivation Language

- 상태: completed by `REVIEW-LOOP-MOTIVATION-001` on 2026-06-07.
- 왜 중요했는가: tutorial success에서도 `잔류 리스크`/`다음 출격`이 반복되어 review 후보 안내가 실제 primary action처럼 읽혔다.
- 완료 내용: tutorial success는 `재점검 메모`/`나중에 다시 풀기`, incident success는 `잔류 리스크`/`다음 출격 후보`를 표시한다. action id/label 계약과 progress 저장 포맷은 변경하지 않았다.
- 검증: `playable_review_queue`, `playable_debrief_success`, `playable_ftue_first_five_route`, `playable_viewport_success_modal_80x24`.

## Content Breadth Readiness

다음 content breadth slice는 진행 가능하다.

- HUD/help/choice/success polish가 full E2E를 통과했다.
- progress 저장 포맷 변경이 필요하지 않다.
- 진행 불가 P0/P1 UX blocker는 현재 없다.
- 새 콘텐츠는 기존 engine command만 사용하고, 필요한 경우 command-choice/applied incident layer로 추가한다.

추천 다음 콘텐츠:

1. `CONTENT-BREADTH-002`: mission/review loop가 닫혔으므로 기존 engine 기반 applied content를 늘린다.
2. `UI-RAIL-001`: mission/review/daily 정보가 HUD를 다시 밀어내면 wide layout side rail을 연다.
3. `UI-STYLE-001`: 첫 release polish에서 color/emphasis pass가 필요해지면 연다.

## Completed / Retired

- `UI-MODAL-ACTION-HIERARCHY-001 — Modal Decision Surface`: failed/succeeded modal을 viewport overlay decision surface로 재배치하고, primary/secondary action footer, running hint/quit action, hint cost copy, incident progressive briefing, final-frame E2E evidence를 보강했다.
- `UI-MODAL-ACTION-HIERARCHY-002 — Modal/Hint Post-release Hardening`: modal을 console surface 위 decision layer로 고정하고, running hint utility action 분리, 한글 final-frame evidence, `다음 런북` action label, summary evidence metadata를 보강했다.
- `UI-EVIDENCE-003 — Final Frame Evidence`: `UI-EVIDENCE-002`에서 runner의 `screen_final.txt` 저장을 구현했고, `E2E-EVIDENCE-008`에서 long incident route가 final/timeline evidence를 남기도록 고정했다.
- `UI-COMMAND-MEMORY-001 — Learned Command Memory`: `UI-POLISH-002`에서 tutorial/hint/failure command memory cue를 구현했다.
- `UI-ACTION-LANGUAGE-001 — Action Line Language Contract`: `UI-CONSOLE-POLISH-001`에서 screen label과 action id를 분리했다.
- `UI-REVIEW-MOTIVATION-001 — Success Review Motivation Language`: `REVIEW-LOOP-MOTIVATION-001`에서 tutorial/incident success review motivation 언어를 분리했다.
