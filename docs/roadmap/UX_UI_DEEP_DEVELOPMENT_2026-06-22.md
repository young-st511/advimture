# UX/UI Deep Development 2026-06-22

## 결론

Advimture의 UI는 "터미널 안의 게임"보다 "게임 세계 안의 터미널 장비"가 되어야 한다. Vim command는 키 암기 과제가 아니라 원격 복구국 오퍼레이터가 보내는 판단 명령이고, 화면은 그 판단이 어떤 신호, 위험, 복구 진척을 만들었는지 즉시 보여줘야 한다.

다음 UX/UI 개선의 핵심은 세 가지다.

1. Action을 본문에서 분리해 항상 누를 수 있는 조작으로 보이게 한다.
2. Running HUD를 현재 목표, 판단, 입력 반응 중심으로 줄인다.
3. Incident를 단일 문제 목록이 아니라 dispatch route와 recovery report로 장면화한다.

## 제품 경험 목표

플레이어가 한 화면에서 느껴야 하는 감각:

- "지금 무엇을 해야 하는지 바로 안다."
- "내 Vim 입력이 세계 안의 장비 명령처럼 반응한다."
- "실패해도 왜 이 선택이 틀렸는지 배운다."
- "다음 문제는 그냥 다음 exercise가 아니라 다음 복구 단계다."
- "힌트는 정답 공개가 아니라 추리 보조다."

비목표:

- 상용 게임 UI를 시각적으로 복제하지 않는다.
- 텍스트 양을 늘려 세계관을 설명하지 않는다.
- Animation이 mission/console을 밀어내지 않는다.
- Progress schema, content schema, Vim engine capability를 UX polish 이유만으로 흔들지 않는다.

## Core UX Model

Advimture의 플레이 루프를 다음 5단계로 재정의한다.

```text
1. Dispatch     지금 복구할 node와 판단 목표를 받는다.
2. Inspect      console buffer를 보고 어떤 Vim 도구를 쓸지 결정한다.
3. Execute      Vim 입력을 보내고 signal/echo가 반응한다.
4. Report       성공/실패 결과를 recovery report로 받는다.
5. Route        다음 node, review 후보, 남은 risk를 확인한다.
```

현재 구현은 2, 3, 4가 기능적으로 존재한다. 부족한 것은 1과 5의 장면감, 그리고 3의 반응성과 4의 학습적 설명력이다.

## Information Architecture

### Base 80-column Layout

80x24를 기준으로 기본 화면은 다음 정보만 강하게 보여준다.

```text
ADVIMTURE | Tutorial or Runbook Dispatch | Episode | 2/5 | running

MISSION  전체 상태값 전환
GOAL     파일 전체의 TODO를 DONE으로 바꾸세요.
TOOLS    기억할 명령 :%s
SIGNAL   [learn]--*-[console]  입력: :%s/TODO/DONE/g

RUNBOOK CONSOLE
> [T]ODO api
  TODO worker

NORMAL · running · cursor 0:0

ACTIONS  [?] 힌트 - grade 영향   [q] 종료
```

원칙:

- `MISSION`은 제목, `GOAL`은 행동 목표다.
- `TOOLS`는 tutorial에서만 기본 노출한다. Incident에서는 hint/failure 전까지 숨긴다.
- `SIGNAL`은 화면 반응이며 본문 설명을 대체하지 않는다.
- `ACTIONS`는 항상 하단 또는 modal footer의 별도 행으로 둔다.
- review/daily는 running 화면 기본형에서 과감히 낮춘다. 80-column에서는 success/debrief 또는 wide rail에서 더 잘 보이게 한다.

### Wide Layout

100 column 이상에서는 오른쪽 rail을 열 수 있다.

```text
ADVIMTURE | Runbook Dispatch | Relay-003 | 2/5 | running

MISSION  오염 구간 격리
GOAL     stale block 전체를 linewise로 제거하세요.
SIGNAL   [relay]-*--[console]  입력: V -> range armed

RUNBOOK CONSOLE                         RUNBOOK ARC
> {stale cache}                          001 signal found     sealed
  {stale token}                          002 structure sync   sealed
  live route                             003 isolate block    active

NORMAL · visual · cursor 1:0             CLOCKS
                                         input 1/3
ACTIONS [?] 힌트 - grade 영향 [q] 종료   review debt 3
```

Wide rail은 보조 정보의 집이다. 현재 목표보다 먼저 눈에 들어오면 실패다.

## Screen System

### 1. Tutorial Running

목표: 새 Vim command를 숨기지 않고, 학습자가 바로 실행할 수 있게 한다.

표시:

- `MISSION`: exercise title
- `GOAL`: 짧은 조작 목표
- `TOOLS`: `기억할 명령`, 필요하면 `훈련 키`
- `SIGNAL`: learning channel, input echo
- `CONSOLE`: buffer
- `ACTIONS`: hint/quit

금지:

- review/daily detail이 console 전 접근을 늦추는 것
- 여러 문장 briefing이 계속 화면 상단을 차지하는 것

좋은 문구:

```text
GOAL   다음 단어 시작으로 이동하세요.
TOOLS  기억할 명령 w
```

나쁜 문구:

```text
TRAINING BRIEF · 기억할 명령: w · 보조 행동 힌트: ? · 종료: q
복구 메모: 재점검 3건 · 다음: 전체 파일 상태값 치환하기
```

문제는 정보가 틀린 게 아니라, 같은 밀도의 텍스트가 너무 많이 붙어 있다는 점이다.

### 2. Incident Running

목표: 정답 key보다 판단 질문이 먼저 보이게 한다.

표시:

- `MISSION`: incident node name
- `SITUATION`: 상황 1문장
- `JUDGMENT`: 선택해야 할 Vim 도구 관점
- `SIGNAL`: relay channel, input echo
- `ACTIONS`: hint/quit

예:

```text
MISSION    오염 구간 격리
SITUATION  stale block이 live route 앞을 막고 있습니다.
JUDGMENT   한 줄씩 지우지 말고 줄 묶음 범위를 선택하세요.
SIGNAL     [relay]-*--[console]  대기: selection
```

Incident에서 `TOOLS`는 기본 노출하지 않는다. hint를 누르면 다음 순서로 연다.

```text
HINT 1  범위가 여러 줄이면 먼저 linewise 범위를 잡으세요.
HINT 2  visual line mode를 떠올리세요.
HINT 3  V로 줄 선택을 시작할 수 있습니다.
```

### 3. Command/Search/Insert/Visual Mode

Mode-specific UI는 매우 중요하다. 현재 mode에 따라 action bar가 바뀌어야 한다.

Command mode:

```text
COMMAND  :%s/TODO/DONE/g
ACTIONS  [enter] 실행   [esc] 취소
```

Search mode:

```text
SEARCH   /timeout
ACTIONS  [enter] 찾기   [esc] 취소
```

Insert mode:

```text
INSERT   입력 중
ACTIONS  [esc] normal로 복귀
```

Visual mode:

```text
SELECT   linewise range armed
ACTIONS  [d] 제거   [y] 복사   [esc] 취소
```

이 상태 안내는 일반 hint/quit보다 우선한다. 왜냐하면 지금 누를 수 있는 키가 mode마다 달라지기 때문이다.

### 4. Failure Report

실패 화면은 감탄사나 세계관보다 선택 이유를 먼저 말한다.

```text
RECOVERY REPORT
원인   현재 줄만 바꿔 두 번째 TODO가 남았습니다.
판단   전체 상태값이면 % range로 파일 전체를 대상으로 잡아야 합니다.
증거   TODO worker

ACTIONS  [enter] 다시 시도   [?] 힌트   [q] 종료
```

구조:

- `원인`: 무엇이 틀렸는가
- `판단`: 어떤 Vim 판단을 해야 하는가
- `증거`: buffer/state에서 보이는 증거
- `ACTIONS`: retry/hint/quit

### 5. Success Report

성공 화면은 점수판이 아니라 복구 리포트다.

```text
RUNBOOK SEALED
복구   TODO 2건을 DONE으로 정리했습니다.
판단   % range와 g flag를 조합해 반복 상태값을 한 번에 처리했습니다.
기록   16 keys · best 16 · grade A
다음   구조 재동기화 node로 이동

ACTIONS  [enter] 다음 단계   [q] 종료
```

성공 메시지는 "잘했다"보다 "무엇을 배웠는가"가 먼저다.

## Interaction Language

### Action Bar Contract

Action은 모두 같은 grammar로 표시한다.

```text
ACTIONS  [key] label   [key] label   [key] label
```

예:

- `[enter] 다음 단계`
- `[r] 다시 시도`
- `[?] 힌트 - grade 영향`
- `[q] 종료`
- `[esc] 취소`
- `[d] 선택 제거`

규칙:

- key는 항상 brackets 안에 둔다.
- label은 동사형 또는 명령형으로 짧게 쓴다.
- cost가 있으면 label 뒤에 붙인다.
- primary action은 가장 왼쪽에 둔다.
- E2E는 기존 `action.id`를 유지하고 screen label은 별도 검증한다.

### Signal Rail Contract

Signal은 현재 입력 상태를 짧게 반응시킨다.

Tutorial:

```text
SIGNAL [learn]-*--[console]  입력: w -> cursor advanced
```

Incident:

```text
SIGNAL [relay]--*-[console]  입력: V -> range armed
```

Failure 직전 또는 forbidden input:

```text
SIGNAL [relay]x---[console]  blocked: arrow key
```

규칙:

- Signal은 1줄을 넘지 않는다.
- 성공/실패 판단을 대신하지 않는다.
- Animation frame은 의미를 약하게 보조한다.
- `reduce-motion` 또는 non-tty 환경에서는 frame 변화 없이 고정 rail이어도 의미가 살아야 한다.

## Visual Design Direction

Advimture는 호러/사이버펑크 터미널이 아니라 "원격 시설 복구국의 낡지만 신뢰 가능한 작업 콘솔"에 가깝다.

톤:

- 차분한 operational UI
- 고장 난 장비의 작은 신호감
- 과한 glitch, matrix green, full-screen ASCII 폭발은 금지
- 색상은 의미 보조만 한다

권장 스타일:

- Primary action: brighter foreground or bold
- Warning/failure: amber/red only on label or border
- Success: green/blue accent, body는 과하지 않게
- Signal rail: subtle dim style
- Console cursor/selection: bracket fallback 유지, style 가능할 때만 background emphasis

ASCII 장식 원칙:

- 장식은 정보 구조를 만들어야 한다.
- frame, rail, route node, clock 정도만 쓴다.
- 커다란 ASCII art는 pre-start scene이나 episode transition에만 제한한다.

## Hint Ladder

현재 content schema를 바로 바꾸지 않고도 1차 ladder를 만들 수 있다.

### Without Schema Change

사용 가능한 재료:

- scenario hint
- trained/reviewed commands
- required key
- failure reason
- command goal

파생:

- Hint 1: scenario hint 또는 판단 관점
- Hint 2: trained/reviewed command family
- Hint 3: required key/command memory

예:

```text
Hint 1  같은 상태값이 여러 줄에 반복됩니다. 범위를 먼저 생각하세요.
Hint 2  substitute는 range를 붙일 수 있습니다.
Hint 3  :%s/old/new/g 형태를 떠올리세요.
```

### With Future Schema Change

나중에는 content YAML에 `hints[]`를 도입할 수 있다. 단, 이는 content schema 변경이므로 별도 승인과 migration 검토가 필요하다.

## Dispatch Route

Incident 001-003은 첫 복구 arc로 이미 정리됐다. UI에서는 이를 route로 보이게 해야 한다.

표현:

```text
ROUTE  001 signal trace [sealed] -> 002 structure sync [active] -> 003 isolate [locked]
```

80-column에서는 축약:

```text
ROUTE  002/003 structure sync
```

Wide rail에서는 node list:

```text
RUNBOOK ARC
001 signal trace      sealed
002 structure sync    active
003 isolate block     locked
```

효과:

- 다음 exercise가 의미 있는 복구 단계로 읽힌다.
- success에서 "다음 단계"가 그냥 버튼이 아니라 route progression이 된다.

## Implementation Roadmap

### Slice 1: UI-ACTION-HUD-001

목표: 가장 먼저 보이는 running 화면을 action bar + dense HUD로 정리한다.

범위:

- `ACTIONS [key] label` renderer 추가
- running FocusPanel action 표시를 action bar로 전환
- Mission HUD copy line을 `MISSION`, `GOAL`, `TOOLS/JUDGMENT`, `SIGNAL`로 분해
- review/daily running line은 기본 80-column에서 축약 또는 숨김

완성 기준:

- 80x24 첫 tutorial 화면에서 `MISSION`, `GOAL`, `SIGNAL`, `ACTIONS`가 모두 보인다.
- `힌트`, `종료`, `다음`, `재시도`가 본문과 구분된다.
- 기존 action id 계약은 유지된다.
- `playable_adventure_signal_input_echo`, viewport success/failure E2E 통과.

### Slice 2: UI-REPORT-001

목표: failure/success modal을 recovery report 구조로 바꾼다.

범위:

- failure modal label을 `RECOVERY REPORT` 중심으로 정리
- `원인`, `판단`, `증거` line 도입
- success modal에 `복구`, `판단`, `기록`, `다음` line 도입

완성 기준:

- 실패/성공 화면의 첫 두 줄만 읽어도 왜 성공/실패했는지 안다.
- action bar가 modal footer에서 보존된다.
- 80x24에서 clipping 없음.

### Slice 3: UI-HINT-LADDER-001

목표: hint를 한 번짜리 도움말에서 점진 공개로 바꾼다.

범위:

- schema 변경 없이 existing hint + command memory + required keys로 단계형 표시
- hints used count에 따라 다음 hint level 표시
- hint cost affordance 강화

완성 기준:

- 첫 hint는 exact command를 바로 말하지 않는다.
- 마지막 hint는 tutorial에서 command form을 충분히 알려준다.
- incident는 judgment first 원칙을 유지한다.

### Slice 4: UI-PRESTART-DISPATCH-001

목표: incident/arc 시작을 짧은 briefing scene으로 만든다.

범위:

- incident 첫 exercise 또는 arc entry 전 pre-start state
- `enter`는 runtime key trace에 들어가지 않음
- tutorial에는 기본 적용하지 않음

완성 기준:

- incident 시작 전 `DISPATCH`, `상황`, `판단`, `ACTIONS`가 보인다.
- pre-start 상태에서 Vim buffer는 아직 조작되지 않는다.

### Slice 5: UI-DISPATCH-RAIL-001

목표: wide viewport에 route/tools/clocks/memory side rail을 추가한다.

범위:

- 100 column 이상에서만 side rail
- 80 column은 기존 vertical layout 유지
- rail은 app_state 검증 필드 또는 renderer test로 의미 검증

완성 기준:

- Console line index가 mobile/narrow에서 유지된다.
- Wide에서 route progress와 clocks가 오른쪽에 보인다.

### Slice 6: UI-STYLE-MOTION-001

목표: color/emphasis/reduce-motion pass.

범위:

- lipgloss style pass
- reduce-motion option 검토
- color 없는 snapshot 의미 보존

완성 기준:

- color off에서도 action/signal/report 의미가 유지된다.
- animation은 정보 전달을 방해하지 않는다.

## Open Decisions

다음 구현 전에 결정이 필요한 것:

- `review/daily` running line을 80-column에서 완전히 숨길지, `ROUTE` 한 줄로 대체할지
- Action bar를 항상 최하단에 둘지, console 직후에 둘지
- Hint ladder 1차를 schema 변경 없이 할지, 바로 content schema 변경을 열지
- Incident pre-start를 모든 incident에 둘지, arc 시작에만 둘지
- Wide rail breakpoint를 100, 110, 120 중 어디로 잡을지

추천:

- `review/daily`는 80-column running에서 숨기고 success/debrief 및 wide rail로 이동한다.
- Action bar는 terminal height가 있으면 최하단, legacy render에서는 console 직후에 둔다.
- Hint ladder 1차는 schema 변경 없이 한다.
- Pre-start는 arc 시작과 incident playlist 첫 entry에만 둔다.
- Wide rail breakpoint는 110으로 시작한다.

## Next Best Step

가장 좋은 다음 작업은 `UI-ACTION-HUD-001`이다.

이유:

- 사용자가 바로 느끼는 문제인 "무엇을 눌러야 하는가"를 해결한다.
- 텍스트 밀도를 줄인다.
- 현재 action id와 FocusPanel 구조를 재사용할 수 있다.
- 이후 report, hint ladder, pre-start, side rail의 기반 UI grammar가 된다.

예상 변경 범위:

- `docs/exec-plans/active/ui-action-hud-001.md`
- `docs/gameplay/spec.md`
- `docs/gameplay/tui-screen-contract.md`
- `docs/gameplay/tui-ux-direction.md`
- `internal/playableview/render.go`
- `internal/playableview/render_test.go`
- `internal/playable/model.go`
- `internal/playable/model_test.go`
- focused E2E 2-4개

금지 경계:

- progress schema 변경 없음
- content schema 변경 없음
- 새 dependency 없음
- 새 Vim engine capability 없음
