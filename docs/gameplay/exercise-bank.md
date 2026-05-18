# Exercise Bank

> Vim command cluster를 반복 훈련 가능한 문항으로 바꾸는 저장소다. 모든 exercise는 `command-catalog.md`의 cluster를 참조한다.

## 운영 규칙

- 새 exercise는 `[draft]` 상태로 추가한다.
- `command_cluster`가 `approved`가 아니면 exercise를 `approved`로 올릴 수 없다.
- exercise는 scenario 없이도 독립적으로 성립해야 한다.
- 각 exercise는 기계 검증 가능한 target state와 optimal key trace를 가진다.
- scenario는 exercise 승인 이후에만 붙인다.

## 상태 값

| 상태 | 의미 |
|------|------|
| `draft` | 문항 초안 |
| `approved` | scenario 설계에 사용할 수 있음 |
| `implemented` | 게임 데이터/코드/E2E와 연결됨 |
| `retired` | 더 이상 사용하지 않음 |

## Exercise 스키마

```yaml
exercise:
  id: <command-cluster-id>-NNN
  status: draft | approved | implemented | retired
  command_cluster: <command-cluster-id>
  engine_support: implemented | planned | unsupported
  trained_commands: ["<commands used in optimal_keys>"]
  reviewed_commands: ["<commands allowed for review or recovery>"]
  mistake_focus: ["<common mistake this exercise teaches against>"]
  replay_status: pending | pass | fail
  title: <문항 이름>
  goal_for_player: <플레이어에게 보여줄 목표 문장>
  initial_state:
    mode: NORMAL | INSERT | COMMAND
    cursor:
      row: <0-based-row>
      col: <0-based-col>
    buffer: |
      <initial text>
  target_state:
    mode: <expected mode, optional>
    cursor:
      row: <0-based-row>
      col: <0-based-col>
    buffer: |
      <target text, optional>
  optimal_keys: "<space-separated key trace>"
  allowed_keys: ["<key>"]
  forbidden_keys: ["<key>"]
  hints:
    - <1단계 힌트>
    - <2단계 힌트>
  grading:
    pass_condition: <기계 검증 조건>
    optimal_key_count: <number>
```

## Exercises

### survival-save-quit-001

```yaml
exercise:
  id: survival-save-quit-001
  status: draft
  command_cluster: survival-save-quit
  engine_support: planned
  trained_commands: [":q!"]
  reviewed_commands: ["esc"]
  mistake_focus: ["saving when the goal is to discard changes"]
  replay_status: pending
  title: 저장하지 않고 빠져나오기
  goal_for_player: "변경하지 않고 Vim에서 빠져나오세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      server_name wrong.example.com;
  target_state:
    mode: NORMAL
    buffer: |
      server_name wrong.example.com;
  optimal_keys: ": q ! enter"
  allowed_keys: ["esc", ":", "q", "!", "enter"]
  forbidden_keys: ["i", "a", "o", "x", "d", "c"]
  hints:
    - "수정하지 않고 나갈 때는 저장 명령이 필요하지 않습니다."
    - "`:q!`는 변경을 버리고 종료합니다."
  grading:
    pass_condition: "app exits with code 0 && buffer unchanged && no progress file required"
    optimal_key_count: 4
```

### normal-motion-basic-001

```yaml
exercise:
  id: normal-motion-basic-001
  status: draft
  command_cluster: normal-motion-basic
  engine_support: implemented
  trained_commands: ["l"]
  reviewed_commands: ["h"]
  mistake_focus: ["confusing h and l"]
  replay_status: pending
  title: 목표 문자까지 이동하기
  goal_for_player: "커서를 X 표시 위로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      ....X
  target_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 4
  optimal_keys: "l l l l"
  allowed_keys: ["h", "j", "k", "l", "esc"]
  forbidden_keys: ["right", "left", "up", "down"]
  hints:
    - "오른쪽 이동은 홈 포지션의 오른손 새끼손가락 쪽 키입니다."
    - "`l`은 오른쪽으로 한 칸 이동합니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 4"
    optimal_key_count: 4
```

### survival-save-quit-002

```yaml
exercise:
  id: survival-save-quit-002
  status: draft
  command_cluster: survival-save-quit
  engine_support: planned
  trained_commands: [":wq"]
  reviewed_commands: ["esc", ":q!"]
  mistake_focus: ["discarding when the goal is to save"]
  replay_status: pending
  title: 저장하고 종료하기
  goal_for_player: "성공한 변경을 저장하고 종료하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      server_name correct.example.com;
  target_state:
    mode: NORMAL
    buffer: |
      server_name correct.example.com;
  optimal_keys: ": w q enter"
  allowed_keys: ["esc", ":", "w", "q", "enter"]
  forbidden_keys: ["i", "a", "o", "x", "d", "c"]
  hints:
    - "이번에는 결과를 저장하고 종료해야 합니다."
    - "`:wq`는 저장한 뒤 종료합니다."
  grading:
    pass_condition: "app exits with code 0 && buffer unchanged && mission completed"
    optimal_key_count: 4
```

### normal-motion-basic-002

```yaml
exercise:
  id: normal-motion-basic-002
  status: draft
  command_cluster: normal-motion-basic
  engine_support: implemented
  trained_commands: ["j"]
  reviewed_commands: ["k"]
  mistake_focus: ["confusing j and k"]
  replay_status: pending
  title: 경고 지점으로 이동하기
  goal_for_player: "커서를 WARN 표시의 W 위로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      INFO boot
      WARN disk
      INFO done
  target_state:
    mode: NORMAL
    cursor:
      row: 1
      col: 0
  optimal_keys: "j"
  allowed_keys: ["h", "j", "k", "l", "esc"]
  forbidden_keys: ["right", "left", "up", "down"]
  hints:
    - "아래 줄로 내려가야 합니다."
    - "`j`는 아래로 한 줄 이동합니다."
  grading:
    pass_condition: "cursor.row == 1 && cursor.col == 0"
    optimal_key_count: 1
```

### normal-motion-basic-003

```yaml
exercise:
  id: normal-motion-basic-003
  status: draft
  command_cluster: normal-motion-basic
  engine_support: implemented
  trained_commands: ["j", "l"]
  reviewed_commands: ["h", "k"]
  mistake_focus: ["losing row/column orientation"]
  replay_status: pending
  title: 짧은 경로 조합하기
  goal_for_player: "커서를 X 표시 위로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      .....
      ..X..
  target_state:
    mode: NORMAL
    cursor:
      row: 1
      col: 2
  optimal_keys: "j l l"
  allowed_keys: ["h", "j", "k", "l", "esc"]
  forbidden_keys: ["right", "left", "up", "down"]
  hints:
    - "아래로 이동한 뒤 오른쪽으로 이동하면 됩니다."
    - "`j` 다음 `l`을 두 번 눌러보세요."
  grading:
    pass_condition: "cursor.row == 1 && cursor.col == 2"
    optimal_key_count: 3
```

### word-motion-basic-001

```yaml
exercise:
  id: word-motion-basic-001
  status: draft
  command_cluster: word-motion-basic
  engine_support: planned
  trained_commands: ["w"]
  reviewed_commands: ["b", "e"]
  mistake_focus: ["using repeated l instead of word motion"]
  replay_status: pending
  title: 단어 시작점으로 뛰어가기
  goal_for_player: "커서를 `backend` 단어의 시작으로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      service api backend enabled
  target_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 12
  optimal_keys: "w w"
  allowed_keys: ["h", "j", "k", "l", "w", "b", "e", "esc"]
  forbidden_keys: ["right", "left", "up", "down"]
  hints:
    - "한 글자씩 이동하지 말고 단어 단위로 이동해보세요."
    - "`w`는 다음 단어의 시작으로 이동합니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 12"
    optimal_key_count: 2
```
