# Scenario Bank

> 승인된 exercise를 어드벤처 사건으로 감싸는 저장소다. 시나리오는 exercise의 학습 목표를 바꾸지 않는다.

## 운영 규칙

- 새 scenario는 `[draft]` 상태로 추가한다.
- `exercise_id`가 `approved`가 아니면 scenario를 `approved`로 올릴 수 없다.
- scenario는 exercise의 정답, target state, allowed keys를 바꾸지 않는다.
- briefing과 피드백은 Vim command 학습 목표를 강화해야 한다.
- 스토리를 제거해도 exercise는 독립적으로 성립해야 한다.

## 상태 값

| 상태 | 의미 |
|------|------|
| `draft` | 시나리오 초안 |
| `approved` | 게임 데이터로 구현 가능 |
| `implemented` | 게임 데이터/코드/E2E와 연결됨 |
| `retired` | 더 이상 사용하지 않음 |

## Scenario 스키마

```yaml
scenario:
  id: <exercise-id>-scenario
  status: draft | approved | implemented | retired
  exercise_id: <exercise-id>
  engine_support: implemented | planned | unsupported
  learning_reinforcement: <which Vim concept the copy reinforces>
  does_not_change:
    - target_state
    - optimal_keys
    - allowed_keys
  mission_title: <미션 제목>
  briefing: <플레이어에게 보여줄 상황 설명>
  context_role: <텍스트가 게임 세계에서 의미하는 것>
  mentor_success: <성공 피드백>
  mentor_failure: <실패/힌트 피드백>
  story_constraints:
    - <문항을 방해하지 않기 위한 제약>
```

## Scenarios

### survival-save-quit-001-scenario

```yaml
scenario:
  id: survival-save-quit-001-scenario
  status: approved
  exercise_id: survival-save-quit-001
  engine_support: implemented
  learning_reinforcement: "`esc`는 현재 mode를 빠져나와 Normal mode로 돌아온다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "일단 빠져나오기"
  briefing: "입력 모드에 커서가 묶였습니다. 더 입력하기 전에 먼저 Normal mode로 복귀하세요."
  context_role: "당황한 터미널 입력"
  mentor_success: "좋습니다. 터미널에서 당황했을 때 첫 수습은 esc로 Normal mode를 되찾는 것입니다."
  mentor_failure: "아직 편집을 계속하는 흐름입니다. 빠져나오려면 esc로 mode를 먼저 정리하세요."
  story_constraints:
    - "저장이나 종료를 요구하지 않는다."
    - "esc의 mode 복귀 의미만 강조한다."
```

### normal-motion-basic-001-scenario

```yaml
scenario:
  id: normal-motion-basic-001-scenario
  status: draft
  exercise_id: normal-motion-basic-001
  engine_support: implemented
  learning_reinforcement: "`l`은 Normal mode에서 오른쪽으로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "커서 위치 맞추기"
  briefing: "로그 줄에서 표시된 지점까지 커서를 이동해야 합니다. 화살표 키 대신 Vim 기본 이동을 사용하세요."
  context_role: "짧은 로그 라인"
  mentor_success: "좋습니다. 손을 홈 포지션에 둔 채 위치를 잡았습니다."
  mentor_failure: "화살표 키를 쓰지 않고 `h/j/k/l` 네 방향을 몸에 익히는 단계입니다."
  story_constraints:
    - "긴 거리 이동을 요구하지 않는다."
    - "기본 이동의 방향 학습에 집중한다."
```

### survival-save-quit-002-scenario

```yaml
scenario:
  id: survival-save-quit-002-scenario
  status: approved
  exercise_id: survival-save-quit-002
  engine_support: implemented
  learning_reinforcement: "`:q!`는 변경을 버리고 나가는 command-line 명령이다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "실험 파일 버리고 나가기"
  briefing: "임시 설정을 잘못 열었습니다. 이 변경은 버려도 되니 저장하지 않고 command-line에서 나가세요."
  context_role: "버려도 되는 임시 설정"
  mentor_success: "좋습니다. :q!는 저장하지 않는다는 판단이 끝난 뒤 쓰는 안전한 종료입니다."
  mentor_failure: "저장하면 안 되는 임시 작업입니다. command-line을 열고 버리고 나가는 흐름을 떠올려보세요."
  story_constraints:
    - "실제 파일 삭제나 저장을 수행하지 않는다."
    - "앱 종료 단축키 q와 Vim 명령 :q!를 구분한다."
```

### survival-save-quit-003-scenario

```yaml
scenario:
  id: survival-save-quit-003-scenario
  status: approved
  exercise_id: survival-save-quit-003
  engine_support: implemented
  learning_reinforcement: "`:wq`는 저장하고 나가는 command-line 명령이다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "수정 저장 후 종료"
  briefing: "배포 전 설정 수정이 끝났습니다. 변경을 남긴 채 세션을 닫아야 하니 저장 후 종료하세요."
  context_role: "수정한 배포 설정"
  mentor_success: "좋습니다. :wq는 검증한 변경을 저장하고 세션을 닫는 기본 마무리입니다."
  mentor_failure: "이번 작업은 버리면 안 됩니다. command-line에서 저장 후 종료하는 명령을 조립하세요."
  story_constraints:
    - "실제 파일 저장을 수행하지 않는다."
    - ":q!와 :wq의 의도 차이를 강조한다."
```

### normal-motion-basic-002-scenario

```yaml
scenario:
  id: normal-motion-basic-002-scenario
  status: approved
  exercise_id: normal-motion-basic-002
  engine_support: implemented
  learning_reinforcement: "`j`는 아래 줄로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "경고 줄 찾기"
  briefing: "부팅 로그에서 WARN 줄을 놓쳤습니다. 아래 줄로 내려가 경고 표시의 첫 글자를 확인하세요."
  context_role: "짧은 시스템 로그"
  mentor_success: "좋습니다. `j`는 아래 줄의 단서를 빠르게 잡을 때 쓰는 기본 이동입니다."
  mentor_failure: "목표는 아래 줄입니다. 커서를 한 줄 내리는 Normal mode 이동을 떠올려보세요."
  story_constraints:
    - "편집이나 검색을 요구하지 않는다."
    - "`j`의 아래 방향 학습에 집중한다."
```

### normal-motion-basic-003-scenario

```yaml
scenario:
  id: normal-motion-basic-003-scenario
  status: approved
  exercise_id: normal-motion-basic-003
  engine_support: implemented
  learning_reinforcement: "`h`는 Normal mode에서 왼쪽으로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "되돌아온 커서 정렬"
  briefing: "표시 지점을 오른쪽으로 지나쳤습니다. 화살표 키를 쓰지 말고 왼쪽으로 돌아가 첫 칸 단서를 다시 잡으세요."
  context_role: "짧은 터미널 지도"
  mentor_success: "좋습니다. `h`는 오른쪽으로 지나쳤을 때 커서를 왼쪽으로 되돌리는 기본 이동입니다."
  mentor_failure: "지금은 왼쪽으로 되돌아가는 연습입니다. `l`의 반대 방향을 떠올려보세요."
  story_constraints:
    - "긴 거리 이동을 요구하지 않는다."
    - "`h`의 왼쪽 방향 학습에 집중한다."
```

### normal-motion-basic-004-scenario

```yaml
scenario:
  id: normal-motion-basic-004-scenario
  status: approved
  exercise_id: normal-motion-basic-004
  engine_support: implemented
  learning_reinforcement: "`k`는 Normal mode에서 위 줄로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "위쪽 경고 줄 복귀"
  briefing: "로그를 한 줄 더 내려가 버렸습니다. 아래쪽 INFO에서 위쪽 WARN 줄로 다시 올라가세요."
  context_role: "짧은 시스템 로그"
  mentor_success: "좋습니다. `k`는 아래로 지나쳤을 때 위쪽 줄의 단서를 다시 잡는 기본 이동입니다."
  mentor_failure: "목표는 위쪽 줄입니다. `j`의 반대 방향으로 한 줄 올라가세요."
  story_constraints:
    - "편집이나 검색을 요구하지 않는다."
    - "`k`의 위 방향 학습에 집중한다."
```

### word-motion-basic-001-scenario

```yaml
scenario:
  id: word-motion-basic-001-scenario
  status: draft
  exercise_id: word-motion-basic-001
  engine_support: planned
  learning_reinforcement: "`w`는 다음 단어 시작점으로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "서비스 이름 찾기"
  briefing: "배포 설정 한 줄에서 `backend` 항목으로 빠르게 이동해야 합니다. 한 글자씩 걷지 말고 단어 단위로 이동하세요."
  context_role: "서비스 라우팅 설정"
  mentor_success: "좋아요. 단어 단위 이동을 쓰면 긴 설정 줄에서도 금방 위치를 잡을 수 있습니다."
  mentor_failure: "지금 목표는 속도입니다. `w`로 다음 단어 시작점까지 뛰어가보세요."
  story_constraints:
    - "편집을 요구하지 않는다."
    - "`w`의 효율을 `l` 반복과 비교해 체감하게 한다."
```

### whole-file-navigation-001-scenario

```yaml
scenario:
  id: whole-file-navigation-001-scenario
  status: approved
  exercise_id: whole-file-navigation-001
  engine_support: implemented
  learning_reinforcement: "`gg`는 파일 처음으로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "설정 맨 위로 복귀"
  briefing: "오류 줄까지 내려왔지만 진짜 단서는 파일 맨 위 선언부에 있습니다. k를 반복하지 말고 파일 처음으로 돌아가세요."
  context_role: "짧은 서버 설정 파일"
  mentor_success: "좋습니다. gg를 알면 깊이 내려온 파일에서도 바로 처음으로 복귀할 수 있습니다."
  mentor_failure: "지금은 한 줄씩 올라가는 문제가 아닙니다. g를 두 번 눌러 파일 처음으로 이동하세요."
  story_constraints:
    - "편집을 요구하지 않는다."
    - "`gg`의 prefix 입력을 명확히 드러낸다."
```

### whole-file-navigation-002-scenario

```yaml
scenario:
  id: whole-file-navigation-002-scenario
  status: approved
  exercise_id: whole-file-navigation-002
  engine_support: implemented
  learning_reinforcement: "`G`는 파일 마지막 줄로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "마지막 상태 확인"
  briefing: "설정 파일 마지막 줄에 현재 상태가 적혀 있습니다. j를 반복하지 말고 바로 마지막 줄로 이동하세요."
  context_role: "짧은 서버 설정 파일"
  mentor_success: "좋습니다. G는 긴 파일에서 마지막 줄을 확인할 때 바로 손에 잡히는 이동입니다."
  mentor_failure: "이번에는 대문자 G입니다. 파일 끝으로 한 번에 내려가세요."
  story_constraints:
    - "대문자 G와 소문자 g를 혼동하지 않게 한다."
    - "줄 끝 이동과 파일 끝 이동을 섞지 않는다."
```

### whole-file-navigation-003-scenario

```yaml
scenario:
  id: whole-file-navigation-003-scenario
  status: approved
  exercise_id: whole-file-navigation-003
  engine_support: implemented
  learning_reinforcement: "`0`은 현재 줄의 첫 column으로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "키 이름으로 복귀"
  briefing: "설정 값 쪽에 커서가 있습니다. 현재 줄의 키 이름을 다시 확인해야 하니 줄 시작으로 이동하세요."
  context_role: "서비스 라우팅 설정"
  mentor_success: "좋습니다. 0을 쓰면 현재 줄의 시작으로 바로 돌아올 수 있습니다."
  mentor_failure: "한 글자씩 왼쪽으로 가지 않아도 됩니다. 0을 눌러 현재 줄 첫 칸으로 이동하세요."
  story_constraints:
    - "파일 처음으로 이동하라고 오해시키지 않는다."
    - "`0`의 현재 줄 범위를 강조한다."
```

### whole-file-navigation-004-scenario

```yaml
scenario:
  id: whole-file-navigation-004-scenario
  status: approved
  exercise_id: whole-file-navigation-004
  engine_support: implemented
  learning_reinforcement: "`$`는 현재 줄의 마지막 글자로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "값 끝 확인"
  briefing: "현재 줄의 마지막 토큰 끝을 확인해야 합니다. l을 반복하지 말고 줄 끝으로 이동하세요."
  context_role: "서비스 라우팅 설정"
  mentor_success: "좋습니다. $는 현재 줄의 끝을 확인할 때 가장 짧은 길입니다."
  mentor_failure: "오른쪽으로 계속 걷지 말고 $로 줄 끝에 도착하세요."
  story_constraints:
    - "파일 마지막 줄과 현재 줄 끝을 섞지 않는다."
    - "`$`의 현재 줄 범위를 강조한다."
```

### vim-ex-command-substitute-001-scenario

```yaml
scenario:
  id: vim-ex-command-substitute-001-scenario
  status: approved
  exercise_id: vim-ex-command-substitute-001
  engine_support: implemented
  learning_reinforcement: "`:s/old/new/`는 현재 줄의 첫 번째 literal match를 바꾼다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "현재 줄만 긴급 수정"
  briefing: "첫 줄의 라우트 타입 하나만 잘못 들어갔습니다. 전체 파일을 건드리지 말고 현재 줄의 첫 api만 web으로 바꾸세요."
  context_role: "서비스 라우팅 설정"
  mentor_success: "좋습니다. :s는 현재 줄의 작은 치환을 빠르게 처리할 때 유용합니다."
  mentor_failure: "이번 목표는 현재 줄의 첫 match입니다. :s/api/web/ 형태를 떠올려보세요."
  story_constraints:
    - "전체 파일 치환으로 오해시키지 않는다."
    - "literal substitute 범위를 강조한다."
```

### vim-ex-command-substitute-002-scenario

```yaml
scenario:
  id: vim-ex-command-substitute-002-scenario
  status: approved
  exercise_id: vim-ex-command-substitute-002
  engine_support: implemented
  learning_reinforcement: "`:%s/old/new/g`는 전체 파일의 모든 literal match를 바꾼다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "전체 상태값 전환"
  briefing: "두 작업 항목이 모두 끝났습니다. 파일 전체의 TODO를 DONE으로 바꾸세요."
  context_role: "배포 체크리스트"
  mentor_success: "좋습니다. % range와 g flag를 조합하면 반복된 상태값을 한 번에 정리할 수 있습니다."
  mentor_failure: "전체 파일을 대상으로 해야 합니다. :%s/TODO/DONE/g 명령을 조립해보세요."
  story_constraints:
    - "현재 줄 치환과 전체 파일 치환의 차이를 강조한다."
    - "regex 설명으로 새지 않는다."
```

### vim-ex-command-substitute-003-scenario

```yaml
scenario:
  id: vim-ex-command-substitute-003-scenario
  status: approved
  exercise_id: vim-ex-command-substitute-003
  engine_support: implemented
  learning_reinforcement: "`:2,3s/old/new/`는 지정한 1-based 줄 범위 안에서 치환한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "문제 구간만 수정"
  briefing: "첫 줄은 아직 조사 중이라 건드리면 안 됩니다. 2~3번째 줄의 error만 ok로 바꾸세요."
  context_role: "장애 처리 메모"
  mentor_success: "좋습니다. range를 붙이면 필요한 구간만 조심스럽게 고칠 수 있습니다."
  mentor_failure: "이번에는 전체 파일이 아닙니다. 2,3 range를 붙여 문제 구간만 바꾸세요."
  story_constraints:
    - "첫 줄이 그대로 남아야 함을 명확히 한다."
    - "range가 1-based inclusive임을 강화한다."
```
