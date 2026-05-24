# TUI UX Direction

> UI/UX 본격 개발을 위한 통합 결론이다. 세계관은 화면의 분위기를 만들지만, Vim 학습 목표보다 앞서지 않는다.

## 현재 판단

현재 TUI는 Vim 훈련 도구로는 기능한다. mode, cursor, input limit, required key coaching, success grade, key count, progress 저장이 모두 노출되어 있어 Agent와 플레이어가 같은 상태를 볼 수 있다.

하지만 화면 정보 계층은 아직 제품 UI라기보다 테스트 가능한 개발 UI에 가깝다.

- 현재 조작 목표보다 `재점검 대상`, `오늘의 복구 루트`가 먼저 보인다.
- tutorial과 incident가 같은 하단 안내 언어를 사용하던 문제는 `TRAINING BRIEF`와 `OPERATOR JUDGMENT` FocusPanel 분리로 1차 개선했다.
- `A D V I M T U R E - Playable Slice`, `Mode`, `Status`, `Cursor`, `ACTION`은 세계관보다 디버그 표면에 가깝다.
- buffer/cursor/selection은 색상 없이도 검증 가능하지만, `[]`, `{}` 기호 삽입이 긴 줄과 visual selection에서 과밀해진다.
- E2E screen artifact는 누적 ANSI log 성격이 강해 현재 frame 검증을 보조하기 어렵다.

## UI 원칙

- Vim 학습 목표가 1순위다.
- 초반 tutorial은 정답 key를 명시해도 된다. 첫 학습에서는 `이번에 익힐 키`, `왜 쓰는지`, `남은 입력`을 우선한다.
- incident는 정답 key보다 선택 판단을 우선한다. 기본 화면에서는 `상황`, `판단 목표`, `복구 단계`를 먼저 보여주고 key sequence는 hint/failure에서 점진 공개한다.
- review/daily 언어는 세계관에 맞지만 보조 정보다. 현재 exercise 목표보다 높은 계층에 두지 않는다.
- Runbook Dispatch 톤은 “상황 1문장 + Vim 조작 목표 1문장”을 넘기지 않는다.
- UI 검증은 화면 문자열뿐 아니라 app_state의 의미 필드를 우선한다.

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
  failed/succeeded floating modal

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
- Status Line: mode, cursor, inputs left를 한 줄에 보여준다.
- Mission HUD cue: `TRAINING BRIEF`로 `훈련 키`, `힌트`, `재시도`, `다음`을 현재 목표 바로 아래에 압축한다.

Tutorial에서는 `재점검 대상`, `오늘의 복구 루트`가 현재 목표보다 앞에 오지 않는다. terminal size가 있는 화면에서는 `복구 현황:` 한 줄로 `MISSION` HUD 안에 접는다.

### Incident Screen

Incident는 “배운 도구를 조합하는 Runbook Dispatch”다. 여기서는 정답 key sequence를 첫 화면부터 전부 말하지 않는다.

- Header: `ADVIMTURE / Runbook Dispatch / relay-004 / 1 of 5`
- Briefing: 상황 1문장 + 판단 목표 1문장
- Console: 복구 대상 buffer
- Status Line: mode, cursor, command/search line
- Mission HUD cue: `OPERATOR JUDGMENT`로 `판단`, `힌트`, `재시도`, `다음 조치`를 현재 목표와 함께 표시한다.

Incident focus panel의 첫 줄은 `훈련 키`보다 `판단:`을 우선한다. 정답 key는 hint 요청 또는 실패 후 회복 메시지에서 점진 공개한다.

### Success/Debrief Screen

성공 화면은 점수판이 아니라 복구 기록이다.

- `복구 기록`: 이번 grade/key count
- `최단 복구 기록`: progress 기반 best record
- `잔류 리스크`: 다음 review candidate
- `오늘의 복구 루트`: 남은 review count
- `Next` 또는 `Next tutorial`: 다음 입력
- terminal size가 있는 화면에서는 `RUNBOOK CONSOLE` 안의 `RUNBOOK SEALED` floating modal로 표시한다.

### Visual/Selection Screen

Visual mode는 텍스트 자체가 더 길어지지 않는 표현이 이상적이다. 단, 색상 없는 환경과 E2E 안정성을 위해 첫 단계에서는 기존 `[]`, `{}` fallback을 유지할 수 있다. 추후 renderer 분리 뒤 lipgloss style 기반 표현으로 이동한다.

## 개발 순서

1. **UX-QA-001**: review/daily route를 typed app_state assertion과 evidence로 검증한다.
2. **UI-CONTRACT-001**: header/briefing/console/status/action panel의 영역 contract를 고정한다.
3. **UI-RENDER-001**: 렌더링을 `internal/playableview` 같은 순수 renderer로 분리한다.
4. **UI-HIERARCHY-001**: 현재 exercise 목표를 상위에 두고 review/daily를 보조 정보로 낮춘다.
5. **UI-MODE-001**: tutorial과 incident 안내 언어를 분리한다.
6. **UI-EVIDENCE-001**: 현재 frame 또는 frame timeline evidence를 추가한다.
7. **UI-MODAL-001**: 안내/성공/실패/debrief를 console 위 FocusPanel로 승격하고 `app_state.ui.focus_panel`로 검증한다.
8. **UI-LAYOUT-001**: `tea.WindowSizeMsg`를 반영해 FocusPanel을 terminal width에 맞춰 중앙 정렬한다.
9. **UI-MODAL-002**: FocusPanel을 고정 modal layer로 렌더링하고 실패/성공/거절 피드백을 panel 안에 모은다.
10. **UI-HUD-003**: 화면을 Mission HUD + Console Core로 재정렬하고 실패/성공/debrief를 console 안 floating modal로 표시한다.

## 후속 UI 방향

- split cockpit wide layout은 Mission HUD + Console Core가 안정된 뒤 별도 slice에서 다룬다.
- 완전한 입력 차단 modal state는 별도 slice에서 다룬다.
- pre-start modal은 modal open 상태의 입력을 Vim runtime key trace와 분리해야 하므로 후속 slice에서 다룬다.
- 세계관 polish는 FocusPanel title과 짧은 상태 용어를 중심으로 확장하고, Vim 학습 목표보다 길어지지 않게 한다.

## 첫 구현 범위

첫 개발은 화면 리디자인이 아니라 QA 기반부터 시작한다. `UX-QA-001`은 기존 화면을 크게 바꾸지 않고, review/daily route를 문자열 `contains` 대신 typed `app_state.review` assertion으로 검증하며, E2E evidence에 app_state/progress snapshot을 남긴다.
