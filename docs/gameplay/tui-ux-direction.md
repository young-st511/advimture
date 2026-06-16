# TUI UX Direction

> UI/UX 본격 개발을 위한 통합 결론이다. 세계관은 화면의 분위기를 만들지만, Vim 학습 목표보다 앞서지 않는다.

## 현재 판단

현재 TUI는 Vim 훈련 도구로는 기능한다. mode, cursor, input limit, required key coaching, success grade, key count, progress 저장이 모두 노출되어 있어 Agent와 플레이어가 같은 상태를 볼 수 있다.

Foundation UI pass 이후 화면 정보 계층은 Mission HUD, Runbook Console, floating modal, status line으로 정리됐다. 다만 사용자 스크린샷 이후 실패/성공 modal이 실제 overlay보다 console 뒤 삽입 블록처럼 보이고 action line이 CTA처럼 보이지 않는 P1 UX blocker가 확인됐다.

- 현재 조작 목표는 Mission HUD에서 먼저 보이며, review/daily 정보는 보조 line으로 내려갔다.
- tutorial과 incident가 같은 하단 안내 언어를 사용하던 문제는 `TRAINING BRIEF`와 `OPERATOR JUDGMENT` FocusPanel 분리로 개선했다.
- `Mode`, `Status`, `Cursor`는 status line에 남아 있지만, 화면 주 정보는 복구 작전과 Vim 목표 중심으로 이동했다.
- buffer/cursor/selection은 색상 없이도 검증 가능하지만, `[]`, `{}` 기호 삽입이 긴 줄과 visual selection에서 과밀해진다.
- E2E screen artifact는 `screen_timeline.txt`와 `screen_final.txt`를 통해 긴 route도 검토 가능해졌다.

## UI 원칙

- Vim 학습 목표가 1순위다.
- 초반 tutorial은 정답 key를 명시해도 된다. 첫 학습에서는 `이번에 익힐 키`, `왜 쓰는지`, `남은 입력`을 우선한다.
- incident는 정답 key보다 선택 판단을 우선한다. 기본 화면에서는 `상황`, `판단 목표`, `복구 단계`를 먼저 보여주고 key sequence는 hint/failure에서 점진 공개한다.
- review/daily 언어는 세계관에 맞지만 보조 정보다. 현재 exercise 목표보다 높은 계층에 두지 않는다.
- Runbook Dispatch 톤은 “상황 1문장 + Vim 조작 목표 1문장”을 넘기지 않는다.
- UI 검증은 화면 문자열뿐 아니라 app_state의 의미 필드를 우선한다.

## Release-Quality 기준

출시 가능한 품질의 TUI는 화려함보다 현재 행동을 숨기지 않는 안정성을 먼저 통과해야 한다.

- 첫 화면에서 현재 목표와 편집 대상 buffer가 바로 보인다.
- tutorial은 새 command를 숨기지 않고, incident는 정답 sequence를 hint/failure 전까지 과하게 노출하지 않는다.
- 실패 화면은 실패 이유와 `다시 시도` primary action footer를 유지한다.
- 성공 화면은 복구 기록, context별 review motivation, 다음 행동 primary action footer를 유지한다.
- review/daily 정보는 현재 mission보다 높은 시각 계층으로 올라오지 않는다.
- 80x24/80x30 계열 viewport에서 action/hint line이 clipping으로 사라지지 않는다.
- style/color pass를 하더라도 color 없는 evidence와 app_state 계약은 유지한다.

현재 판정은 `docs/roadmap/PLAYABLE_QUALITY_BASELINE_2026-06-02.md`를 따른다.

## 현재 정보 구조

```text
Header
  ADVIMTURE / current playlist or incident / exercise index / status

Mission HUD
  current title
  situation sentence + Vim operation or judgment goal
  recovery status as a compact secondary line
  running/mode cue line

Runbook Console
  runbook buffer with cursor and selection
  failed/succeeded floating modal with primary action footer

Status Line
  mode / status / cursor / inputs left / command or search line
```

넓은 터미널에서는 `Console + Side Rail`을 고려할 수 있다. 좁은 터미널에서는 위 구조를 세로로 쌓는다.

## 상세 기획 루프 2차 결론

이번 UI는 “게임처럼 꾸미는 것”보다 “조작 판단이 즉시 보이는 복구 콘솔”을 목표로 한다. 화면의 첫 시선은 항상 현재 exercise의 제목과 Vim 목표로 가야 한다.

### Tutorial Screen

초반 tutorial은 “첫 투어”다. 플레이어가 아직 Vim 단어를 모를 수 있으므로 훈련 키를 직접 알려준다.

- Header: `ADVIMTURE / Tutorial / 1 of 4 / running`
- Briefing: 현재 상황과 익힐 조작을 짧게 표시한다.
- Console: 수정할 텍스트를 가장 넓게 보여준다.
- Status Line: terminal size가 있는 HUD에서는 `NORMAL · running · cursor 0:2`처럼 개발식 label 없이 한 줄로 보여준다.
- Mission HUD cue: `TRAINING BRIEF`로 `훈련 키`, `힌트`, `재시도`, `다음`을 현재 목표 바로 아래에 압축한다.
- Mission HUD cue는 `기억할 명령: ...`을 짧게 보여줘, 콘텐츠가 늘어나도 방금 배우는 Vim command를 놓치지 않게 한다.

Tutorial에서는 `재점검 대상`, `오늘의 복구 루트`가 현재 목표보다 앞에 오지 않는다. terminal size가 있는 running 화면에서는 `복구 메모: 재점검 N건 · 다음: <title>` 한 줄로 축약하고, action cue 뒤에 둔다.

Incident running 화면에서는 `복구 현황: 재점검 N건 · 잔류: <title>`로 축약해 세계관 메타 정보는 유지하되 현재 목표와 console 접근을 늦추지 않는다.

### Incident Screen

Incident는 “배운 도구를 조합하는 Runbook Dispatch”다. 여기서는 정답 key sequence를 첫 화면부터 전부 말하지 않는다.

- Header: `ADVIMTURE / Runbook Dispatch / relay-004 / 1 of 5`
- Briefing: 상황 1문장 + 판단 목표 1문장
- Console: 복구 대상 buffer
- Status Line: mode, cursor, command/search line
- Mission HUD cue: `OPERATOR JUDGMENT`로 `판단`, `힌트`, `재시도`, `다음 조치`를 현재 목표와 함께 표시한다.

Incident focus panel의 첫 줄은 `훈련 키`보다 `판단:`을 우선한다. 정답 key와 command memory는 기본 running 화면에서 숨기고, hint 요청 또는 실패 후 `참고 명령`/회복 메시지로 점진 공개한다.

### Success/Debrief Screen

성공 화면은 점수판이 아니라 복구 기록이다.

- `복구 기록`: 이번 grade/key count
- `최단 복구 기록`: progress 기반 best record
- Tutorial success: `재점검 메모`와 `나중에 다시 풀기`로 다음 review candidate를 보조 동기로 보여준다.
- Incident success: `잔류 리스크`와 `다음 출격 후보`로 다음 review candidate를 Runbook Dispatch 톤에 맞게 보여준다.
- `오늘의 복구 루트`: primary review 대상, 이유, 남은 review count
- `다음 행동`: modal footer의 primary action prefix. 뒤에 `다음 단계`, `다음 튜토리얼`, `다음 runbook`, `다음 출격`, `출격 완료`, `다시 시도` 같은 action label을 붙인다.
- `보조 행동`: modal footer의 secondary action prefix. 보통 `종료: q`를 붙인다.
- E2E 의미 검증은 화면 label 문자열이 아니라 `app_state.ui.focus_panel.actions[].id`를 우선한다.
- terminal size가 있는 화면에서는 `RUNBOOK CONSOLE` 안의 `RUNBOOK SEALED` floating modal로 표시한다.

### Visual/Selection Screen

Visual mode는 텍스트 자체가 더 길어지지 않는 표현이 이상적이다. 단, 색상 없는 환경과 E2E 안정성을 위해 첫 단계에서는 기존 `[]`, `{}` fallback을 유지할 수 있다. 추후 renderer 분리 뒤 lipgloss style 기반 표현으로 이동한다.

## 완료된 Foundation UI 순서

1. **UX-QA-001**: review/daily route를 typed app_state assertion과 evidence로 검증했다.
2. **UI-CONTRACT-001**: header/briefing/console/status/action panel의 영역 contract를 고정했다.
3. **UI-RENDER-001**: 렌더링을 `internal/playableview` 순수 renderer로 분리했다.
4. **UI-HIERARCHY-001**: 현재 exercise 목표를 상위에 두고 review/daily를 보조 정보로 낮췄다.
5. **UI-MODE-001**: tutorial과 incident 안내 언어를 분리했다.
6. **UI-EVIDENCE-001**: screen timeline evidence를 추가했다.
7. **UI-MODAL-001**: 안내/성공/실패/debrief를 FocusPanel 모델로 승격하고 `app_state.ui.focus_panel`로 검증했다.
8. **UI-LAYOUT-001**: `tea.WindowSizeMsg`를 반영해 FocusPanel을 terminal width에 맞춰 중앙 정렬했다.
9. **UI-MODAL-002**: 실패/성공/거절 피드백을 panel 안에 모았다.
10. **UI-HUD-003**: 화면을 Mission HUD + Console Core로 재정렬하고 실패/성공/debrief를 console 안 floating modal로 표시했다.
11. **UI-EVIDENCE-002 / E2E-EVIDENCE-008**: final screen evidence와 long route timeline/final evidence를 고정했다.
12. **UI-POLISH-002**: tutorial command memory와 incident hint/failure command memory를 `FocusPanel`에 얇게 추가했다.
13. **UI-CONSOLE-POLISH-001**: success/failure action line을 한국어 label과 internal `action.id`로 분리하고, focused E2E가 action 의미를 검증하게 했다.
14. **UI-MODAL-ACTION-HIERARCHY-001 / Todo 3**: failed/succeeded modal을 terminal viewport overlay로 배치하고, action footer를 `다음 행동`/`보조 행동`으로 body와 분리했다.
15. **UI-MODAL-ACTION-HIERARCHY-001 / Todo 4**: running `hint`/`quit`을 `FocusPanel.actions`로 구조화하고 HUD에서는 `보조 행동` utility cue로 분리했다.
16. **UI-MODAL-ACTION-HIERARCHY-001 / Todo 5**: hint revealed line을 `힌트 내용  ... · 등급에 영향`으로 바꾸고 failure modal hint도 secondary action으로 분리했다.
17. **UI-MODAL-ACTION-HIERARCHY-001 / Todo 6**: incident briefing에서 exact command sequence 과노출을 줄이고 상황/판단 목표를 먼저 보이게 다듬었다.
18. **UI-MODAL-ACTION-HIERARCHY-001 / Todo 7**: 대표 success/failure/hint/incident/review evidence를 재생성하고, modal/action hierarchy blocker를 completed로 닫았다.

## 후속 UI 방향

- Foundation polish 이후 현재 진행 불가 P0/P1 UX 리스크는 없다.
- split cockpit wide layout은 Mission HUD + Console Core가 안정된 뒤 별도 slice에서 다룬다.
- 완전한 입력 차단 modal state는 별도 slice에서 다룬다.
- pre-start modal은 modal open 상태의 입력을 Vim runtime key trace와 분리해야 하므로 후속 slice에서 다룬다.
- 세계관 polish는 FocusPanel title과 짧은 상태 용어를 중심으로 확장하고, Vim 학습 목표보다 길어지지 않게 한다.
- 과거 action line은 `UI-CONSOLE-POLISH-001`에서 한국어 screen label과 machine-readable action id로 분리했다. 이후 UI copy 변경은 `actions.id`를 유지한 채 label만 조정한다.
- Header의 track label은 첫 투어와 incident 전환을 분리하는 제품 신호다. `Tutorial`과 `Runbook Dispatch`는 첫 viewport에서 유지한다.
- 새 UI 개선은 먼저 “현재 Vim 조작 목표를 더 잘 보이게 하는가?”를 통과해야 하며, review/daily/세계관 정보가 mission briefing보다 앞서면 다시 backlog로 되돌린다.

## 현재 완료 범위

Foundation UI는 QA 기반, renderer 분리, Mission HUD, floating modal, final/timeline evidence까지 완료됐다. 다음 UI 작업은 `docs/roadmap/UX_BACKLOG_001.md`의 P1/P2 backlog를 기준으로 열고, `PLAN-REFRESH-009`에서 content/platform 우선순위와 함께 다시 판단한다.
