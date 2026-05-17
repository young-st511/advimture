# Vim Learning Design Loops

> Advimture의 핵심 설계 루프다. 목표는 스토리 있는 튜토리얼이 아니라, Vim을 유용하게 쓰게 만드는 반복 훈련 어드벤처를 만드는 것이다.

## 제품 철학

> 유용하고 재미있는 Vim 훈련 게임을 먼저 만들고, 거기에 맞는 스토리를 붙인다.

이 문서는 모든 튜토리얼, 미션, 스테이지를 만들 때 따르는 순서를 정의한다.

```text
Vim Command Loop
  → Exercise Design Loop
  → Scenario Skinning Loop
```

역방향 설계는 금지한다. 시나리오가 먼저 떠올랐다면, 먼저 어떤 Vim command를 훈련하는지로 환원한 뒤 다시 시작한다.

## Loop 1. Vim Command Loop

### 목적

플레이어가 실제 Vim 사용에서 유용하게 쓰는 단축어와 명령어를 고른다.

### 입력

- Vim command 후보
- 기존 command catalog
- 플레이어 숙련도 단계
- 이전 미션에서 이미 학습한 조작

### 산출물

```yaml
command_cluster:
  id: word-motion-basic
  commands: ["w", "b", "e"]
  purpose: 단어 단위 이동
  prerequisite: ["hjkl"]
  useful_when:
    - 긴 줄에서 단어 단위로 빠르게 이동할 때
    - 삭제/변경 operator와 조합할 때
  combo_paths:
    - ["dw", "cw", "yw"]
  difficulty: beginner
```

### 강화 질문

- 이 command는 실전에서 자주 쓰이는가?
- 이전 command와 자연스럽게 조합되는가?
- 플레이어가 `hjkl`보다 빠르다는 체감을 얻는가?
- command 하나만이 아니라 다음 문법으로 이어지는가?
- 이 command를 모르면 이후 미션 이해가 막히는가?

### 완료 조건

- command cluster가 하나의 학습 목표로 묶인다.
- 선행 command가 명시된다.
- 실무 유용성이 한 문장으로 설명된다.
- 다음 exercise에서 검증 가능한 조작으로 바뀔 수 있다.

## Loop 2. Exercise Design Loop

### 목적

선택한 command cluster를 반복 훈련할 수 있는 문항과 정답을 만든다.

### 입력

- 승인된 command cluster
- 플레이어가 이미 배운 command 목록
- 원하는 난이도와 반복 횟수

### 산출물

```yaml
exercise:
  id: word-motion-basic-001
  command_cluster: word-motion-basic
  initial_text: "server_name old.example.com;"
  target_state:
    cursor:
      row: 0
      col: 12
  optimal_keys: "w w"
  allowed_keys: ["w", "b", "e", "h", "j", "k", "l", "esc"]
  hints:
    - "한 글자씩 가지 말고 단어 단위 이동을 써보세요."
    - "`w`는 다음 단어의 시작으로 이동합니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 12"
    optimal_key_count: 2
```

### 강화 질문

- 정답 상태가 기계적으로 검증 가능한가?
- 최적 입력이 명확한가?
- 허용 키가 학습 목표를 흐리지 않는가?
- 이전 command를 자연스럽게 복습하는가?
- 같은 command를 다른 텍스트/맥락에서 2~3회 반복할 수 있는가?
- 실패했을 때 힌트가 정답을 바로 말하지 않고 사고를 돕는가?

### 완료 조건

- initial state와 target state가 있다.
- optimal key trace가 있다.
- 허용 키와 금지 키 기준이 있다.
- 최소 2단계 힌트가 있다.
- 채점 기준이 화면 상태 또는 buffer/cursor state로 검증 가능하다.

## Loop 3. Scenario Skinning Loop

### 목적

검증된 exercise를 기억에 남는 어드벤처 사건으로 감싼다.

### 입력

- 승인된 exercise
- 현재 월드/챕터 톤
- 캐릭터 목록
- 플레이어의 진행 단계

### 산출물

```yaml
scenario:
  exercise_id: word-motion-basic-001
  mission_title: "잘못된 서버 이름"
  briefing: "배포 설정에서 서버 이름 위치를 찾아야 합니다. 단어 단위 이동으로 목표 지점까지 이동하세요."
  context_text: "server_name old.example.com;"
  mentor_success: "좋아요. 한 글자씩 걷지 않고 단어 단위로 뛰어넘었네요."
  mentor_failure: "지금은 빠른 이동을 익히는 중이에요. `w`로 다음 단어 시작점을 찾아보세요."
```

### 강화 질문

- 시나리오가 command 학습 목표를 더 선명하게 만드는가?
- 플레이어가 왜 이 텍스트를 편집/탐색하는지 이해하는가?
- 대사가 문항 풀이를 방해하지 않는가?
- 성공 피드백이 사용한 Vim 개념을 다시 언급하는가?
- 실패 피드백이 세계관 농담보다 학습 회복을 우선하는가?

### 완료 조건

- briefing이 exercise 목표와 일치한다.
- 성공/실패 피드백이 command 학습을 강화한다.
- 스토리를 제거해도 exercise가 독립적으로 성립한다.
- 시나리오가 exercise의 정답을 바꾸지 않는다.

## 설계 산출물 흐름

```text
docs/gameplay/spec.md
  현재 동작 / 수용 기준

docs/workflows/vim-learning-loops.md
  설계 프로세스

docs/gameplay/command-catalog.md
  command cluster 목록

docs/gameplay/exercise-bank.md
  검증된 exercise 목록

docs/gameplay/scenario-bank.md
  exercise에 입힌 scenario 목록
```

## Agent 운영 규칙

- 새 미션을 만들 때 scenario 제목부터 쓰지 않는다.
- 먼저 command cluster를 제안한다.
- command cluster가 승인되면 exercise를 만든다.
- exercise가 승인되면 scenario를 입힌다.
- 각 loop는 독립적으로 리뷰 가능해야 한다.
- 한 loop의 출력이 약하면 다음 loop로 넘어가지 않는다.
