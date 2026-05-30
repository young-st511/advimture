# UX Backlog 001 — Foundation Polish Aftercare

Created: 2026-05-26
Status: current
Last reviewed: 2026-05-30

## 판단

Content breadth로 넘어가기 전에 막는 P0 UX 리스크는 없다.

현재 Foundation은 current mission, console, failed/succeeded floating modal, hint/retry/quit affordance, review/daily motivation을 E2E로 검증한다. 다음 단계의 UX 작업은 콘텐츠 확장을 막는 선행 조건이 아니라, 콘텐츠가 늘면서 반복 플레이 품질을 높이는 후속 backlog로 관리한다.

## P0

없음.

P0로 승격하는 조건:

- visible affordance가 실제 입력과 다시 어긋난다.
- failed/succeeded modal의 다음 행동 line이 잘려 플레이 진행이 막힌다.
- 새 콘텐츠가 `app_state` 없이 화면 문자열만으로만 검증되어 E2E가 불안정해진다.

## P1

### UI-RAIL-001 — Wide Layout Side Rail

- 왜 중요한가: 콘텐츠가 길어질수록 review/daily, command memory, route progress를 한 줄 HUD에 계속 압축하기 어렵다.
- 언제 열 것인가: incident가 6개 이상으로 늘거나 daily/review 표시가 현재 mission briefing을 다시 밀어낼 때.
- 어떻게 검증할 것인가: desktop width에서 console line index가 고정되고, mobile/narrow width에서는 기존 vertical HUD로 fallback하는 renderer tests와 E2E screenshot/text evidence를 추가한다.

### UI-COMMAND-MEMORY-001 — Learned Command Memory

- 왜 중요한가: 콘텐츠가 늘면 플레이어가 방금 배운 command를 기억하기 어렵다. 다만 기본 화면에 command list를 크게 노출하면 판단 훈련을 약화할 수 있다.
- 언제 열 것인가: applied incident에서 같은 command를 반복 실패하거나, command-choice가 정답 암기/찍기처럼 보이는 evidence가 쌓일 때.
- 어떻게 검증할 것인가: tutorial에서는 current command를 짧게 노출하고, incident에서는 hint/failure에서만 command memory를 열어주는 E2E를 둔다.

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

## Content Breadth Readiness

다음 content breadth slice는 진행 가능하다.

- HUD/help/choice/success polish가 full E2E를 통과했다.
- progress 저장 포맷 변경이 필요하지 않다.
- P0 UX blocker가 없다.
- 새 콘텐츠는 기존 engine command만 사용하고, 필요한 경우 command-choice/applied incident layer로 추가한다.

추천 다음 콘텐츠:

1. `CONTENT-BREADTH-002`: mission/review loop가 닫혔으므로 기존 engine 기반 applied content를 늘린다.
2. `UI-COMMAND-MEMORY-001`: applied incident에서 반복 실패 evidence가 쌓이면 hint/failure 중심 command memory를 연다.
3. `UI-RAIL-001`: mission/review/daily 정보가 HUD를 다시 밀어내면 wide layout side rail을 연다.

## Completed / Retired

- `UI-EVIDENCE-003 — Final Frame Evidence`: `UI-EVIDENCE-002`에서 runner의 `screen_final.txt` 저장을 구현했고, `E2E-EVIDENCE-008`에서 long incident route가 final/timeline evidence를 남기도록 고정했다.
