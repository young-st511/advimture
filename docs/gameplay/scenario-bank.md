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
  status: draft
  exercise_id: survival-save-quit-001
  engine_support: planned
  learning_reinforcement: "`:q!`는 변경을 버리고 나가는 command다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "일단 빠져나오기"
  briefing: "프로덕션 서버에서 낯선 파일이 열렸습니다. 지금은 수정하지 말고 안전하게 빠져나오는 것이 목표입니다."
  context_role: "처음 마주친 nginx 설정 파일"
  mentor_success: "좋아요. 저장하지 않고 빠져나오는 법을 알면, 적어도 더 망치지는 않습니다."
  mentor_failure: "지금은 편집이 아니라 탈출이 목표입니다. 변경을 버리고 나가는 명령을 떠올려보세요."
  story_constraints:
    - "플레이어에게 파일 수정을 요구하지 않는다."
    - "`:q!`와 `:wq`의 차이를 흐리지 않는다."
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
  status: draft
  exercise_id: survival-save-quit-002
  engine_support: planned
  learning_reinforcement: "`:wq`는 저장한 뒤 종료한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "작업 마무리"
  briefing: "설정 확인을 마쳤습니다. 이제 변경 결과를 저장하고 안전하게 종료하세요."
  context_role: "검증된 서버 설정 파일"
  mentor_success: "좋습니다. 필요한 작업을 저장하고 세션을 닫았습니다."
  mentor_failure: "이번 목표는 버리고 나가는 것이 아닙니다. 저장 후 종료하는 명령을 떠올려보세요."
  story_constraints:
    - "`:q!`와 `:wq`의 차이를 분명히 보여준다."
    - "추가 편집을 요구하지 않는다."
```

### normal-motion-basic-002-scenario

```yaml
scenario:
  id: normal-motion-basic-002-scenario
  status: draft
  exercise_id: normal-motion-basic-002
  engine_support: implemented
  learning_reinforcement: "`j`는 아래 줄로 이동한다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "경고 줄 찾기"
  briefing: "부팅 로그에서 경고가 난 줄로 이동해야 합니다. 아래 줄로 내려가 WARN 표시를 확인하세요."
  context_role: "짧은 시스템 로그"
  mentor_success: "좋습니다. 위아래 이동을 써서 필요한 줄을 잡았습니다."
  mentor_failure: "지금은 아래 줄로 내려가는 연습입니다. `j` 방향을 떠올려보세요."
  story_constraints:
    - "편집이나 검색을 요구하지 않는다."
    - "`j`의 아래 방향 학습에 집중한다."
```

### normal-motion-basic-003-scenario

```yaml
scenario:
  id: normal-motion-basic-003-scenario
  status: draft
  exercise_id: normal-motion-basic-003
  engine_support: implemented
  learning_reinforcement: "`j`와 `l`을 조합해 행과 열을 맞춘다."
  does_not_change: ["target_state", "optimal_keys", "allowed_keys"]
  mission_title: "좌표 맞추기"
  briefing: "터미널 지도에서 X 표시로 커서를 옮기세요. 짧은 이동을 조합하면 됩니다."
  context_role: "터미널 지도"
  mentor_success: "좋아요. 짧은 이동을 조합해서 정확한 좌표에 도착했습니다."
  mentor_failure: "한 번에 멀리 가려 하지 말고, 아래 이동과 오른쪽 이동을 나누어 생각하세요."
  story_constraints:
    - "긴 거리 반복 입력을 요구하지 않는다."
    - "단어 이동이나 검색을 아직 요구하지 않는다."
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
