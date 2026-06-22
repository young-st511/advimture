# UX/UI Research 2026-06-20

## 목적

Advimture의 다음 UX/UI 개선 방향을 외부 게임 UI 레퍼런스와 일반 UX 원칙을 바탕으로 정리한다. 목표는 다른 게임의 화면을 복제하는 것이 아니라, Vim 학습 TUI에 맞는 상호작용 원칙으로 번역하는 것이다.

## 현재 진단

현재 Advimture는 조작 목표, console, 실패/성공 modal, action id, signal rail을 갖췄다. 기능과 검증성은 좋아졌지만, 실제 플레이 감각은 아직 설명형 콘솔에 가깝다.

주요 문제:

- 화면의 텍스트 계층이 여전히 비슷한 무게로 읽힌다.
- running 화면에서 "무엇을 봐야 하는가"가 mission, cue, review line, signal, console 사이에 분산된다.
- `SIGNAL` rail은 반응성을 주기 시작했지만 아직 mission state와 강하게 연결되지 않는다.
- action은 구조화됐지만 시각적으로는 아직 텍스트 footer에 가깝다.
- incident arc의 공간감과 진행감이 약하다.
- 힌트는 존재하지만 "점진 공개되는 추리/판단 도움"이라기보다 일반 도움말에 가깝다.

## 레퍼런스 관찰

### Duskers

관찰: 플레이어는 낡은 장비와 command-line interface를 통해 원격 드론을 조작한다. Steam 설명도 "old gritty tech"와 "only eyes and ears", command-line control을 핵심으로 둔다.

Advimture 적용:

- Vim 입력을 단순 키 입력이 아니라 원격 복구 장비에 보내는 command로 읽히게 만든다.
- console 주변에 `signal`, `link`, `scan`, `risk` 같은 작은 상태 신호를 붙인다.
- 단, Duskers처럼 조작법을 숨기면 Vim 학습을 방해하므로 초반 tutorial은 더 명시적이어야 한다.

### Hacknet

관찰: terminal-driven UI를 세계 그 자체로 취급한다. 공식 설명은 "no obvious game elements"와 "support system"을 함께 말한다.

Advimture 적용:

- 게임 UI와 개발 도구 UI를 분리하지 말고, tool UI 자체가 세계관 UI가 되게 한다.
- novice support는 숨기지 않는다. `기억할 명령`, `참고 명령`, command memory는 세계관 안의 operator assist로 표현한다.

### Seoul 2033

관찰: 텍스트, 위기, 선택, 결과 변화가 핵심이다. 스토어 설명은 선택과 판단이 운명과 도시의 미래를 바꾼다고 설명한다.

Advimture 적용:

- incident는 "정답 키 입력"보다 "어떤 Vim 도구를 선택할지"가 먼저 보이게 한다.
- command-choice beat는 선택지 UI처럼 보이게 할 수 있다. 예: `범위 선택`, `구조 편집`, `inline target`, `반복 재사용`.
- failure는 틀렸다는 판정이 아니라 선택 결과 리포트로 보여준다.

### 80 Days

관찰: route, clock, resource, replayability가 이야기와 UI를 묶는다. 공식 페이지는 수많은 route와 strategic planning을 강조한다.

Advimture 적용:

- incident 001-003 같은 arc를 "dispatch route"로 시각화한다.
- 한 exercise 성공이 다음 route node를 켜는 느낌을 준다.
- review queue는 단순 복습 목록이 아니라 오늘의 우회 경로, 잔류 리스크, 재출격 후보로 보이게 한다.

### Citizen Sleeper

관찰: Dice, Clocks, Drives로 선택 가능한 행동과 진행도를 명확하게 만든다. Steam 설명은 clocks가 플레이어 행동과 타인의 행동을 추적한다고 설명한다.

Advimture 적용:

- 각 incident에 작은 progress clock을 붙인다. 예: `LINK SYNC 2/5`, `RISK 1`, `REVIEW DEBT 3`.
- key count, best record, input left를 점수판이 아니라 operation clock으로 표시한다.
- "이번 입력이 무엇을 진전시켰는지"를 input echo에 연결한다.

### Professor Layton

관찰: 퍼즐은 장면 안에 있고, 미해결 퍼즐과 힌트 경제가 플레이 흐름을 만든다. 공식 신작 페이지에서도 캐릭터가 미해결 퍼즐을 추적하는 역할을 가진다.

Advimture 적용:

- hint는 단일 도움말이 아니라 ladder로 만든다.
- 1단계: 판단 관점, 2단계: command family, 3단계: exact form.
- hint 비용은 숨기지 말고 "점수 영향"을 action 옆에 붙인다.

## UX 원칙

- Recognition over recall: 플레이어가 Vim 명령을 머리에서 꺼내기만 기다리지 말고, 현재 맥락에서 떠올릴 단서를 보여준다.
- 명확한 interactive signifier: 누를 수 있는 것은 본문과 다른 형태로 보여준다.
- 현재 목표와 controls reminder는 항상 접근 가능해야 한다.
- 중요한 일시 정보는 시야 밖에 두지 않는다.
- 배경 animation은 의미를 보조해야 하며, 끌 수 있거나 의미 없이 읽혀야 한다.

## 추천 개선 방향

### 1. Action Bar를 먼저 만든다

가장 먼저 고칠 부분은 action footer다. `힌트: ?`, `종료: q`, `다음 단계: enter`를 본문 텍스트가 아니라 action bar로 보이게 한다.

제안:

```text
ACTIONS  [enter] 다음 단계   [?] 힌트 - grade 영향   [q] 종료
```

완성 기준:

- 80x24에서 action bar가 잘리지 않는다.
- action id는 기존 `FocusPanel.actions[].id`를 유지한다.
- color 없이도 action과 본문이 구분된다.

### 2. Running HUD를 더 덜어낸다

현재 running HUD는 mission, cue, review, signal이 모두 보인다. 첫 화면의 기본형은 다음 정도로 줄인다.

```text
MISSION  커서 위치 맞추기
GOAL     목표 문자까지 Vim 기본 이동으로 이동
SIGNAL   [learn]-*--[console]  입력: l -> cursor advanced
```

review/daily는 wide rail 또는 success/debrief에서 더 잘 보이게 하는 편이 낫다.

### 3. Pre-start Briefing Scene을 incident에만 도입한다

모든 exercise에 modal을 띄우면 답답하다. 대신 incident 시작 또는 arc 시작에만 짧은 scene을 둔다.

제안:

```text
DISPATCH  RELAY-001
상황      error 신호가 relay log에서 반복됩니다.
판단      먼저 원인 신호를 찾고, 다음 복구 명령을 결정하세요.

[enter] 출격   [?] 작전 힌트   [q] 종료
```

완성 기준:

- pre-start 상태의 `enter`는 Vim key trace에 들어가지 않는다.
- tutorial 기본 흐름은 지연시키지 않는다.

### 4. Wide Layout Side Rail을 연다

80 column에서는 지금처럼 세로형을 유지하고, 100 column 이상에서만 오른쪽 rail을 둔다.

Rail 후보:

- `RUNBOOK ARC`: 현재 incident node와 남은 node
- `TOOLS`: 현재 학습/추천 command family
- `CLOCKS`: input left, risk, review debt
- `MEMORY`: 최근 성공 command 3개

### 5. Hint Ladder를 만든다

현재 hint는 하나의 문장이다. adventure/puzzle game 감각을 위해 단계형 힌트가 좋다.

제안:

- Hint 1: 판단 관점
- Hint 2: Vim 도구 범주
- Hint 3: command form

첫 구현은 content schema를 바꾸지 않고 renderer/model에서 existing hint와 command memory를 조합해도 된다. schema 변경이 필요하면 별도 승인 대상으로 둔다.

### 6. Feedback을 "판정"이 아니라 "결과 리포트"로 바꾼다

실패 modal은 틀렸다는 감정보다 "왜 이 도구 선택이 위험했는지"를 먼저 말해야 한다.

제안:

```text
RECOVERY REPORT
원인   현재 줄만 바꿔 잔류 TODO가 남았습니다.
판단   전체 상태값이면 % range가 필요합니다.

ACTIONS [enter] 다시 시도   [?] 힌트   [q] 종료
```

## 추천 순서

1. `UI-ACTION-BAR-001`: action footer를 action bar로 승격
2. `UI-HUD-DENSITY-002`: running HUD 3줄 기준으로 재정리
3. `UI-PRESTART-DISPATCH-001`: incident/arc 시작 briefing scene
4. `UI-DISPATCH-RAIL-001`: wide viewport side rail
5. `UI-HINT-LADDER-001`: 단계형 힌트와 hint cost affordance
6. `UI-STYLE-002`: color/emphasis/reduce-motion pass

## 가장 좋은 다음 한 수

`UI-ACTION-BAR-001`과 `UI-HUD-DENSITY-002`를 하나의 짧은 slice로 묶는 것이 좋다. 이유는 다음과 같다.

- 사용자가 즉시 보는 "뭘 눌러야 하지?" 문제를 직접 줄인다.
- 새 content/schema/progress 없이 가능하다.
- 기존 action id 계약을 활용해 E2E 안정성을 유지할 수 있다.
- 이후 pre-start scene과 wide rail의 기반이 된다.

## 참고 출처

- Duskers: https://store.steampowered.com/app/254320/Duskers/
- Hacknet: https://hacknet-os.com/
- Seoul 2033: https://play.google.com/store/apps/details?id=com.banjigamaes.seoul2033_global
- 80 Days: https://www.inklestudios.com/80days/
- Citizen Sleeper: https://store.steampowered.com/app/1578650/Citizen_Sleeper/
- Professor Layton and the New World of Steam: https://www.layton.jp/jouki/en/
- Game Accessibility Guidelines: https://gameaccessibilityguidelines.com/full-list/
- NN/g Recognition and Recall: https://www.nngroup.com/articles/recognition-and-recall/
