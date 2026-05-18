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

## PLAYPACK-002 Small Edits — Implemented Set

| Exercise ID | Command cluster | Trained command | 핵심 목표 | Optimal keys |
|-------------|-----------------|-----------------|-----------|--------------|
| `single-char-edit-001` | `single-char-edit` | `x` | 커서 아래 불필요한 문자 하나 삭제 | `x` |
| `single-char-edit-002` | `single-char-edit` | `r` | 커서 아래 문자 하나 교체 | `r i` |
| `insert-mode-entry-001` | `insert-mode-entry` | `i`, `esc` | 커서 앞에 한 글자 삽입 후 Normal mode 복귀 | `i b esc` |
| `insert-mode-entry-002` | `insert-mode-entry` | `a`, `esc` | 커서 뒤에 한 글자 추가 후 Normal mode 복귀 | `a i esc` |
| `insert-mode-entry-003` | `insert-mode-entry` | `A`, `esc` | 줄 끝에 즉시 추가 후 Normal mode 복귀 | `A ! esc` |
| `undo-redo-basic-001` | `undo-redo-basic` | `u` | 실수 변경을 undo로 복구 | `x u` |
| `undo-redo-basic-002` | `undo-redo-basic` | `ctrl+r` | undo한 변경을 redo로 다시 적용하고 결과 위치 확인 | `x u ctrl+r h` |

공통 제약:

- 모든 문항은 `approved`, `engine_support: implemented`, `replay_status: pass` 상태다.
- 모든 문항은 `constraints.required_keys`로 의도 command 사용을 고정한다.
- `undo-redo-basic-002`는 required key 없이 최종 목표에 먼저 도착하지 않도록 target cursor를 함께 검증한다.

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
  status: approved
  command_cluster: normal-motion-basic
  engine_support: implemented
  trained_commands: ["j"]
  reviewed_commands: ["k"]
  mistake_focus: ["confusing j and k"]
  replay_status: pass
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
  status: approved
  command_cluster: normal-motion-basic
  engine_support: implemented
  trained_commands: ["h"]
  reviewed_commands: ["l"]
  mistake_focus: ["왼쪽으로 돌아갈 때 h와 l 방향을 헷갈린다."]
  replay_status: pass
  title: 줄 앞쪽 단서로 되돌아가기
  goal_for_player: "커서를 첫 번째 a 위로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 2
    buffer: |
      abc
  target_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
  optimal_keys: "h h"
  allowed_keys: ["h", "j", "k", "l", "esc"]
  forbidden_keys: ["right", "left", "up", "down"]
  hints:
    - "목표는 왼쪽입니다. 오른쪽 이동과 반대 방향을 떠올려보세요."
    - "`h`는 왼쪽으로 한 칸 이동합니다."
  grading:
    pass_condition: "cursor.row == 0 && cursor.col == 0"
    optimal_key_count: 2
```

### normal-motion-basic-004

```yaml
exercise:
  id: normal-motion-basic-004
  status: approved
  command_cluster: normal-motion-basic
  engine_support: implemented
  trained_commands: ["k"]
  reviewed_commands: ["j"]
  mistake_focus: ["위로 올라갈 때 j와 k 방향을 헷갈린다."]
  replay_status: pass
  title: 위쪽 로그 줄로 복귀하기
  goal_for_player: "커서를 WARN 표시의 W 위로 이동하세요."
  initial_state:
    mode: NORMAL
    cursor:
      row: 2
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
  optimal_keys: "k"
  allowed_keys: ["h", "j", "k", "l", "esc"]
  forbidden_keys: ["right", "left", "up", "down"]
  hints:
    - "목표는 위쪽 줄입니다. 아래 이동과 반대 방향을 떠올려보세요."
    - "`k`는 위로 한 줄 이동합니다."
  grading:
    pass_condition: "cursor.row == 1 && cursor.col == 0"
    optimal_key_count: 1
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

### vim-ex-command-substitute-001

```yaml
exercise:
  id: vim-ex-command-substitute-001
  status: approved
  command_cluster: vim-ex-command-substitute
  engine_support: implemented
  trained_commands: [":s"]
  reviewed_commands: [":%s"]
  mistake_focus: ["현재 줄만 바꿔야 하는데 전체 파일 치환을 떠올린다."]
  replay_status: pass
  title: 현재 줄 토큰 치환하기
  goal_for_player: "command-line에서 현재 줄의 첫 api를 web으로 바꾸세요."
  target_state:
    mode: NORMAL
    buffer: |
      web api
      api worker
  optimal_keys: ": s / a p i / w e b / enter"
  grading:
    pass_condition: "buffer[0] == web api && buffer[1] == api worker"
    optimal_key_count: 12
```

### vim-ex-command-substitute-002

```yaml
exercise:
  id: vim-ex-command-substitute-002
  status: approved
  command_cluster: vim-ex-command-substitute
  engine_support: implemented
  trained_commands: [":%s"]
  reviewed_commands: [":s"]
  mistake_focus: ["전체 파일 치환에 % range를 붙이지 않는다."]
  replay_status: pass
  title: 전체 파일 상태값 치환하기
  goal_for_player: "command-line에서 모든 TODO를 DONE으로 바꾸세요."
  target_state:
    mode: NORMAL
    buffer: |
      DONE api
      DONE worker
  optimal_keys: ": % s / T O D O / D O N E / g enter"
  grading:
    pass_condition: "buffer has no TODO"
    optimal_key_count: 16
```

### vim-ex-command-substitute-003

```yaml
exercise:
  id: vim-ex-command-substitute-003
  status: approved
  command_cluster: vim-ex-command-substitute
  engine_support: implemented
  trained_commands: [":2,3s"]
  reviewed_commands: [":s", ":%s"]
  mistake_focus: ["특정 줄 범위만 바꿔야 하는데 전체 파일을 바꾼다."]
  replay_status: pass
  title: 줄 범위 안에서 치환하기
  goal_for_player: "command-line에서 2~3번째 줄의 첫 error를 ok로 바꾸세요."
  target_state:
    mode: NORMAL
    buffer: |
      error one
      ok two
      ok three
  optimal_keys: ": 2 , 3 s / e r r o r / o k / enter"
  grading:
    pass_condition: "buffer[0] unchanged && buffer[1:3] changed"
    optimal_key_count: 16
```
