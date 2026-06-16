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
- command/search prompt는 해당 mode에서 입력 중일 때만 화면에 표시한다. succeeded/failed floating modal에서는 이전 command/search 입력을 `Command: ...`로 다시 노출하지 않고, 재현 evidence는 `app_state.command`로 유지한다.

### 5. FocusPanel / Floating Feedback Modal

역할: 다음 행동을 명확히 안내한다.

위치:

- terminal size가 있는 화면은 `MISSION` HUD → `RUNBOOK CONSOLE` → status line 순서를 따른다.
- running/mode-specific 안내는 `MISSION` HUD 안의 짧은 cue로 접는다. 긴 hint나 command memory는 terminal width 기준으로 여러 줄에 감싸며, 중요한 action/hint 문구를 truncation으로 잃지 않는다.
- running HUD에서는 cue line을 review/daily summary보다 먼저 배치한다.
- tutorial running HUD의 review/daily는 `복구 메모: 재점검 N건 · 다음: <title>`처럼 축약한다.
- incident running HUD의 review/daily는 `복구 현황: 재점검 N건 · 잔류: <title>`처럼 축약한다.
- HUD briefing은 terminal width를 기준으로 최대 2줄까지 wrap하고, 초과분은 `...`로 축약할 수 있다.
- failed/succeeded/debrief 안내는 `RUNBOOK CONSOLE` 위에서 floating modal로 표시하되, buffer 뒤에 단순 append된 일반 본문 블록처럼 보이면 안 된다.
- floating modal은 terminal width/height가 알려진 경우 viewport 기준으로 horizontal/vertical placement를 계산한다.
- 좁은 화면에서는 modal width가 terminal width를 넘지 않도록 줄어든다.
- failed/succeeded modal은 console label과 buffer 위치를 밀지 않고, status/grade line보다 높은 decision layer로 표시한다.

상태별 규칙:

- running tutorial: `MISSION` HUD cue에 command memory, 필요한 경우 훈련 키를 표시하고, `hint`/`quit`은 `보조 행동  힌트: ? · 종료: q` utility action으로 분리한다. command memory와 coach key가 같은 정보를 말하면 중복 노출하지 않는다.
- running incident: `MISSION` HUD cue에 판단 cue를 표시하고, `hint`/`quit`은 `보조 행동  힌트: ? · 종료: q` utility action으로 분리한다. command memory는 hint/failure 후에만 `참고 명령`으로 점진 공개
- command/search mode: 입력 중인 prompt와 실행/취소 방법
- insert/search/command/visual mode cue는 한국어 action label로 표현하고, 실제 입력 처리와 맞지 않는 일반 hint/quit 안내를 섞지 않는다.
- failed: floating modal에 실패 이유, 남은 입력, attempts, recovery hint, primary retry action footer
- succeeded: floating modal에 복구 기록, best record, runbook completion, context별 review motivation, primary next/complete action footer
- failed/succeeded feedback은 briefing이 아니라 panel 본문에 둔다.
- failed/succeeded floating modal 주변에는 detailed review/daily line을 다시 올리지 않는다. review/daily 의미는 modal 내부의 review motivation과 `app_state.review`로 유지한다.
- tutorial success의 review motivation은 `재점검 메모`/`나중에 다시 풀기`로 표시해 실제 primary action처럼 읽히지 않게 한다.
- incident success의 review motivation은 `잔류 리스크`/`다음 출격 후보`로 표시해 Runbook Dispatch 후보성을 유지한다.

구조:

- `kind`: `training`, `incident`, `failure`, `success`, `mode`
- `title`: `TRAINING BRIEF`, `OPERATOR JUDGMENT`, `RECOVERY REQUIRED`, `STEP SEALED`, `명령 모드` 등
- `lines`: 사용자에게 보일 안내 문구
- `actions`: retry/next/hint/quit 같은 조작 의미. 화면에는 `label`을 action footer로 표시하고, E2E는 `id`로 검증한다.
- failed modal은 `RECOVERY CHECK`, success modal은 `RUNBOOK SEALED` heading으로 감싸되, app_state의 원래 focus panel kind/title/lines는 유지한다.
- floating modal이 추가하는 보조 label은 `실수`, `힌트`, `배운 점`, `기록`처럼 한국어로 표시한다.
- action footer는 modal body의 기록/힌트/review motivation과 분리한다. 화면에서는 primary action을 `다음 행동`, secondary action을 `보조 행동` prefix로 표시한다. `retry`/`next*`/`dispatch_complete`/`playlist_complete` 같은 primary action은 `quit`이나 `hint` 같은 secondary action보다 먼저 읽혀야 한다.

Action label 계약:

| id | label | 의미 |
|----|-------|------|
| `retry` | `다시 시도: r 또는 enter` | 실패한 exercise 재시도 |
| `next` | `다음 단계: enter` | 같은 playlist의 다음 exercise |
| `next_tutorial` | `다음 튜토리얼: enter` | 다음 tutorial playlist |
| `next_runbook` | `다음 runbook: enter` | 다음 incident/runbook |
| `next_dispatch` | `다음 출격: enter` | review queue primary exercise로 재출격 |
| `dispatch_complete` | `출격 완료` | incident path 완료 |
| `playlist_complete` | `플레이리스트 완료` | tutorial path 완료 |
| `hint` | `힌트: ?` | 현재 exercise hint 열기 |
| `quit` | `종료: q` | 현재 화면 종료 |

금지:

- 현재 목표보다 먼저 focus panel이 시선을 빼앗는 것
- `복구 현황`이 별도 큰 pre-console section으로 console 접근을 늦추는 것
- incident 첫 화면에서 모든 정답 key를 과하게 노출하는 것

## Tutorial과 Incident의 차이

| 항목 | Tutorial | Incident |
|------|----------|----------|
| 목적 | 새 Vim command 학습 | 배운 command 선택/조합 |
| 키 노출 | 직접 노출 허용, `기억할 명령` 표시 | 기본 화면에서는 숨기고 hint/failure에서 `참고 명령`으로 점진 공개 |
| 문구 | `훈련 키`, `연습`, `재점검 메모`, `나중에 다시 풀기`, `다음 튜토리얼` | `판단`, `복구 단계`, `잔류 리스크`, `다음 출격 후보`, `다음 runbook`, `출격 완료` |
| review/daily | 보조 정보 | 세계관 메타 정보로 활용 |

## 검증 기준

- E2E는 화면 문구와 함께 app_state를 본다.
- focus panel은 `app_state.ui.focus_panel`으로 kind/title/lines/actions를 검증할 수 있어야 한다.
- review/daily는 `app_state.review` typed assertion으로 검증한다.
- visual selection은 app_state selection object로 검증한다.
- 화면 레이아웃 변경 후에도 key trace와 progress assertion이 유지되어야 한다.
