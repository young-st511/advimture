# TUI Screen Contract

> `UI-CONTRACT-001`의 결과물이다. 화면 영역과 정보 우선순위를 고정해 이후 renderer 분리와 UI 개편의 기준으로 삼는다.

## 목표

Advimture의 TUI는 Vim 학습 게임이면서 원격 시설 복구국의 콘솔이어야 한다. 화면은 현재 조작 목표를 가장 먼저 보여주고, progress/review 정보는 플레이 동기를 주는 보조 정보로 둔다.

## 영역

### 1. Header

역할: 플레이어가 현재 어디에 있는지 빠르게 인식한다.

필수 정보:

- 제품명: `ADVIMTURE`
- 현재 track: `Tutorial` 또는 `Runbook Dispatch`
- 현재 playlist/incident title
- exercise index: `1/4`
- runtime status: `running`, `failed`, `succeeded`

금지:

- 긴 review 문구
- hint 문구
- debug-only label

### 2. Briefing

역할: 지금 해야 할 일을 알려준다.

규칙:

- 상황 1문장 + Vim 조작/판단 목표 1문장
- tutorial은 새 key 의미를 직접 설명할 수 있다.
- incident는 정답 key sequence보다 판단 목표를 우선한다.

### 3. Console

역할: 플레이어가 실제로 편집/탐색하는 텍스트 표면이다.

필수 정보:

- buffer lines
- cursor
- visual selection
- command/search mode일 때 입력 중인 command line과 연결되는 상태

초기 구현:

- 현재 `[]` cursor, `{}` selection fallback을 유지할 수 있다.
- renderer 분리 후 style 기반 cursor/selection으로 개선한다.

### 4. Status Line

역할: Vim 상태를 짧게 요약한다.

필수 정보:

- mode
- status
- cursor
- inputs left, 해당할 때만
- command/search prompt, 해당할 때만

규칙:

- status line은 한 줄을 기본으로 한다.
- terminal size가 있는 HUD 화면은 `NORMAL · running · cursor 0:2` 형태를 사용한다. terminal size가 없는 fallback은 legacy text를 유지할 수 있다.

### 5. FocusPanel / Floating Feedback Modal

역할: 다음 행동을 명확히 안내한다.

위치:

- terminal size가 있는 화면은 `MISSION` HUD → `RUNBOOK CONSOLE` → status line 순서를 따른다.
- running/mode-specific 안내는 `MISSION` HUD 안의 짧은 cue line으로 접는다.
- failed/succeeded/debrief 안내는 `RUNBOOK CONSOLE` 안에서 floating modal로 표시한다.
- floating modal은 terminal width가 알려진 경우 horizontal center 정렬을 사용한다.
- 좁은 화면에서는 modal width가 terminal width를 넘지 않도록 줄어든다.
- failed/succeeded modal은 console label과 buffer 위치를 밀지 않는다.

상태별 규칙:

- running tutorial: `MISSION` HUD cue line에 훈련 키, hint, quit
- running incident: `MISSION` HUD cue line에 판단 cue, hint, quit
- command/search mode: 입력 중인 prompt와 실행/취소 방법
- failed: floating modal에 실패 이유, 남은 입력, attempts, retry
- succeeded: floating modal에 복구 기록, best record, runbook completion, residual risk, next action
- failed/succeeded feedback은 briefing이 아니라 panel 본문에 둔다.

구조:

- `kind`: `training`, `incident`, `failure`, `success`, `mode`
- `title`: `TRAINING BRIEF`, `OPERATOR JUDGMENT`, `RECOVERY REQUIRED`, `STEP SEALED`, `COMMAND CHANNEL` 등
- `lines`: 사용자에게 보일 안내 문구
- failed modal은 `RECOVERY CHECK`, success modal은 `RUNBOOK SEALED` heading으로 감싸되, app_state의 원래 focus panel kind/title/lines는 유지한다.

금지:

- 현재 목표보다 먼저 focus panel이 시선을 빼앗는 것
- `복구 현황`이 별도 큰 pre-console section으로 console 접근을 늦추는 것
- incident 첫 화면에서 모든 정답 key를 과하게 노출하는 것

## Tutorial과 Incident의 차이

| 항목 | Tutorial | Incident |
|------|----------|----------|
| 목적 | 새 Vim command 학습 | 배운 command 선택/조합 |
| 키 노출 | 직접 노출 허용 | hint/failure에서 점진 공개 |
| 문구 | `훈련 키`, `연습`, `Next tutorial` | `판단`, `복구 단계`, `잔류 리스크`, `Next runbook`, `Dispatch complete` |
| review/daily | 보조 정보 | 세계관 메타 정보로 활용 |

## 검증 기준

- E2E는 화면 문구와 함께 app_state를 본다.
- focus panel은 `app_state.ui.focus_panel`으로 kind/title/lines를 검증할 수 있어야 한다.
- review/daily는 `app_state.review` typed assertion으로 검증한다.
- visual selection은 app_state selection object로 검증한다.
- 화면 레이아웃 변경 후에도 key trace와 progress assertion이 유지되어야 한다.
