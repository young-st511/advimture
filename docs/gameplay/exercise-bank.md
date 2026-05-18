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
  status: approved
  command_cluster: survival-save-quit
  engine_support: implemented
  trained_commands: ["esc"]
  reviewed_commands: []
  mistake_focus: ["Insert mode에서 Normal mode로 빠져나오지 못한다."]
  replay_status: pass
  title: Normal mode로 돌아오기
  goal_for_player: "Insert mode에서 빠져나와 Normal mode로 돌아오세요."
  initial_state:
    mode: INSERT
    cursor:
      row: 0
      col: 0
    buffer: |
      draft
  target_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
  optimal_keys: "esc"
  allowed_keys: ["esc"]
  forbidden_keys: ["ctrl+c"]
  hints:
    - "esc는 현재 mode를 빠져나와 Normal mode로 돌아갑니다."
  grading:
    pass_condition: "mode == normal"
    optimal_key_count: 1
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
  status: approved
  command_cluster: survival-save-quit
  engine_support: implemented
  trained_commands: [":q!"]
  reviewed_commands: ["esc"]
  mistake_focus: ["저장하지 않고 나가야 할 때 :q!를 입력하지 못한다."]
  replay_status: pass
  title: 변경 버리고 나가기
  goal_for_player: "command-line에 :q!를 실행하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      throwaway draft
  target_state:
    mode: NORMAL
    command: ":q!"
  optimal_keys: ": q ! enter"
  allowed_keys: [":", "q", "!", "enter", "esc"]
  forbidden_keys: ["ctrl+c"]
  hints:
    - ":로 command-line mode에 들어갑니다."
    - "q!는 변경을 버리고 종료한다는 뜻입니다."
  grading:
    pass_condition: "last_command == :q!"
    optimal_key_count: 4
```

### survival-save-quit-003

```yaml
exercise:
  id: survival-save-quit-003
  status: approved
  command_cluster: survival-save-quit
  engine_support: implemented
  trained_commands: [":wq"]
  reviewed_commands: ["esc", ":q!"]
  mistake_focus: ["저장 후 종료와 버리고 종료를 구분하지 못한다."]
  replay_status: pass
  title: 저장하고 나가기
  goal_for_player: "command-line에 :wq를 실행하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      edited config
  target_state:
    mode: NORMAL
    command: ":wq"
  optimal_keys: ": w q enter"
  allowed_keys: [":", "w", "q", "enter", "esc"]
  forbidden_keys: ["ctrl+c"]
  hints:
    - ":로 command-line mode에 들어갑니다."
    - "wq는 저장(write)하고 종료(quit)한다는 뜻입니다."
  grading:
    pass_condition: "last_command == :wq"
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

### whole-file-navigation-001

```yaml
exercise:
  id: whole-file-navigation-001
  status: approved
  command_cluster: whole-file-navigation
  engine_support: implemented
  trained_commands: ["gg"]
  reviewed_commands: ["G"]
  mistake_focus: ["파일 처음으로 갈 때 k를 반복한다."]
  replay_status: pass
  title: 파일 처음으로 돌아가기
  goal_for_player: "커서를 파일 첫 줄 첫 칸으로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor: {row: 2, col: 4}
    buffer: |
      server {
      route api
      error here
  target_state:
    mode: NORMAL
    cursor: {row: 0, col: 0}
  optimal_keys: "g g"
  allowed_keys: ["g", "G", "0", "$", "h", "j", "k", "l", "esc"]
  forbidden_keys: ["up", "down", "left", "right"]
  hints:
    - "g는 prefix입니다. 한 번 더 g를 눌러 파일 처음으로 이동합니다."
    - "gg는 파일의 처음으로 이동합니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 0"
    optimal_key_count: 2
```

### whole-file-navigation-002

```yaml
exercise:
  id: whole-file-navigation-002
  status: approved
  command_cluster: whole-file-navigation
  engine_support: implemented
  trained_commands: ["G"]
  reviewed_commands: ["gg"]
  mistake_focus: ["파일 끝으로 갈 때 j를 반복한다."]
  replay_status: pass
  title: 파일 끝으로 이동하기
  goal_for_player: "커서를 파일 마지막 줄 첫 칸으로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor: {row: 0, col: 2}
    buffer: |
      server {
      route api
      status ok
  target_state:
    mode: NORMAL
    cursor: {row: 2, col: 0}
  optimal_keys: "G"
  allowed_keys: ["g", "G", "0", "$", "h", "j", "k", "l", "esc"]
  forbidden_keys: ["up", "down", "left", "right"]
  hints:
    - "G는 파일 마지막 줄로 이동합니다."
    - "여러 번 j를 누르지 않아도 됩니다."
  grading:
    pass_condition: "cursor.row == 2 && cursor.col == 0"
    optimal_key_count: 1
```

### whole-file-navigation-003

```yaml
exercise:
  id: whole-file-navigation-003
  status: approved
  command_cluster: whole-file-navigation
  engine_support: implemented
  trained_commands: ["0"]
  reviewed_commands: ["$"]
  mistake_focus: ["줄 시작으로 갈 때 h를 반복한다."]
  replay_status: pass
  title: 현재 줄 시작으로 이동하기
  goal_for_player: "커서를 현재 줄 첫 칸으로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor: {row: 0, col: 10}
    buffer: |
      route api backend
  target_state:
    mode: NORMAL
    cursor: {row: 0, col: 0}
  optimal_keys: "0"
  allowed_keys: ["0", "$", "h", "l", "w", "b", "e", "esc"]
  forbidden_keys: ["left", "right"]
  hints:
    - "0은 현재 줄의 첫 칸으로 이동합니다."
    - "긴 줄에서 h를 반복하지 않아도 됩니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 0"
    optimal_key_count: 1
```

### whole-file-navigation-004

```yaml
exercise:
  id: whole-file-navigation-004
  status: approved
  command_cluster: whole-file-navigation
  engine_support: implemented
  trained_commands: ["$"]
  reviewed_commands: ["0"]
  mistake_focus: ["줄 끝으로 갈 때 l을 반복한다."]
  replay_status: pass
  title: 현재 줄 끝으로 이동하기
  goal_for_player: "커서를 현재 줄 마지막 글자로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor: {row: 0, col: 0}
    buffer: |
      route api backend
  target_state:
    mode: NORMAL
    cursor: {row: 0, col: 16}
  optimal_keys: "$"
  allowed_keys: ["0", "$", "h", "l", "w", "b", "e", "esc"]
  forbidden_keys: ["left", "right"]
  hints:
    - "$는 현재 줄의 마지막 글자로 이동합니다."
    - "긴 줄에서 l을 반복하지 않아도 됩니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 16"
    optimal_key_count: 1
```
