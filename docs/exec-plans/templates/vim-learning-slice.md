# ExecPlan: <Vim learning slice 제목>

Slice-ID: VIM-<NNN>
Created: <YYYY-MM-DD>
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/gameplay/command-catalog.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- test/e2e/
- internal/

## 목표

<이번 slice에서 플레이어가 어떤 Vim command를 유용하게 쓸 수 있게 되는지 작성한다. 스토리 목표가 아니라 학습 목표를 먼저 쓴다.>

## 범위

- 포함: <command cluster, exercise, scenario, 구현/검증 범위>
- 제외: <이번 slice에서 다루지 않을 command, UX, 저장 포맷, 시나리오 확장>

## Loop 1 — Vim Command

### 입력

- 후보 command:
- 선행 command:
- 플레이어 숙련도:

### 산출물

`docs/gameplay/command-catalog.md`에 다음 항목을 추가하거나 갱신한다.

```yaml
command_cluster:
  id: <id>
  status: draft
  title: <title>
  commands: []
  purpose: <purpose>
  prerequisite: []
  difficulty: beginner | intermediate | advanced
  useful_when: []
  combo_paths: []
  common_mistakes: []
  design_notes: []
```

### 승인 체크

- [ ] command cluster가 하나의 학습 목표로 묶인다.
- [ ] 실무 유용성이 명확하다.
- [ ] 선행 관계가 명시되어 있다.
- [ ] 사람이 승인하여 `status: approved`로 바꿨다.

## Loop 2 — Exercise

### 입력

- 승인된 command cluster:
- 반복 횟수:
- 난이도:

### 산출물

`docs/gameplay/exercise-bank.md`에 다음 항목을 추가하거나 갱신한다.

```yaml
exercise:
  id: <id>
  status: draft
  command_cluster: <command-cluster-id>
  title: <title>
  goal_for_player: <goal>
  initial_state:
    mode: NORMAL
    cursor:
      row: 0
      col: 0
    buffer: |
      <text>
  target_state:
    cursor:
      row: 0
      col: 0
  optimal_keys: "<keys>"
  allowed_keys: []
  forbidden_keys: []
  hints: []
  grading:
    pass_condition: <condition>
    optimal_key_count: <number>
```

### 승인 체크

- [ ] initial state와 target state가 기계적으로 검증 가능하다.
- [ ] optimal key trace가 명확하다.
- [ ] allowed/forbidden keys가 학습 목표를 흐리지 않는다.
- [ ] 힌트가 최소 2단계다.
- [ ] 사람이 승인하여 `status: approved`로 바꿨다.

## Loop 3 — Scenario

### 입력

- 승인된 exercise:
- 챕터/월드 톤:
- 등장 캐릭터:

### 산출물

`docs/gameplay/scenario-bank.md`에 다음 항목을 추가하거나 갱신한다.

```yaml
scenario:
  id: <id>
  status: draft
  exercise_id: <exercise-id>
  mission_title: <title>
  briefing: <briefing>
  context_role: <role>
  mentor_success: <message>
  mentor_failure: <message>
  story_constraints: []
```

### 승인 체크

- [ ] briefing이 exercise 목표와 일치한다.
- [ ] 성공/실패 피드백이 command 학습을 강화한다.
- [ ] 스토리를 제거해도 exercise가 독립적으로 성립한다.
- [ ] 사람이 승인하여 `status: approved`로 바꿨다.

## 구현 계획

- <승인된 산출물을 어떤 코드/data 파일에 반영할지 작성한다.>

## 검증 계획

- `go test ./...`
- `make e2e-smoke`
- <필요 시 추가 E2E scenario>

## E2E Evidence

- scenario:
- artifact path:
- 확인할 screen/buffer/cursor assertion:

## 의사결정 로그

(작업 중 추가)

## 미해결 질문

(작업 중 추가)
