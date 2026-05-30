# Content Requirements

> Scenario workflow를 돌리며 발견한 콘텐츠 loader 요구사항이다. 이 문서는 CONTENT-001의 입력으로 사용한다.

Status: historical discovery + schema reference
Last reviewed: 2026-05-30

CONTENT-001 loader와 replay/coverage validator는 이미 구현됐다. 이 문서는 새 content schema 요구가 발견될 때 참고하는 discovery record로 유지한다. 현재 playable 동작의 canonical source는 `content/` YAML, `docs/gameplay/spec.md`, 그리고 관련 completed ExecPlan이다.

## 원칙

- 콘텐츠는 항상 `command_cluster -> exercise -> scenario -> playlist` 순서로 연결된다.
- scenario는 exercise의 목표, 정답, 허용 키를 바꾸지 않는다.
- 콘텐츠 파일은 YAML을 우선하며 repo root의 `content/` 아래에 둔다.
- draft 콘텐츠도 파일로 보관할 수 있다. 단 `status: draft` 또는 `engine_support: planned`는 playable 후보에서 제외한다.
- loader는 사람이 읽기 쉬운 작성 단위와 엔진이 소비하기 쉬운 compiled 단위를 분리해야 한다.
- 현재 엔진이 지원하지 않는 command를 포함한 콘텐츠는 `engine_support: planned`로 남기고 playable path에 연결하지 않는다.
- 모든 playable 콘텐츠는 TUI E2E가 재현할 수 있는 key trace와 app state assertion을 가져야 한다.
- 실제 학습용 tutorial playlist는 8문항 이하로 나누며, approved/implemented playlist가 이를 초과하면 loader가 거부한다.
- retired playlist는 과거 스냅샷으로 보관할 수 있지만 default playable path에는 포함하지 않는다.
- 새 command 학습 문항은 최대 입력 수, 필수 command, 금지 입력, 금지 우회 전략, 실패 코칭을 함께 정의해야 한다.
- 최대 입력 수 초과나 금지 입력 사용은 즉시 실패이며, 초반 튜토리얼 재시도는 기본 무제한이다.
- `constraints.max_inputs`는 입력 횟수 제한, `constraints.required_keys`는 의도 command 사용 여부, `constraints.forbidden_keys`는 명시적으로 막을 우회 입력을 표현한다.
- `constraints.attempt_limit: 0`은 무제한 재시도이며, 후반 콘텐츠에서 제한을 켤 여지를 남긴다.

## First 5-Minute Loop

ID: `first-5-minute`

목적: 플레이어가 Vim에서 당황하지 않고, Normal mode에서 커서를 움직이며, 짧은 목표를 정확히 달성하는 첫 성공 경험을 만든다.

> 현재 구현 메모: 이 섹션의 beat 예시는 CONTENT-001 이전 discovery 기록이다. 현재 playable content의 canonical source는 `content/` YAML과 `docs/gameplay/spec.md`다. `first-5-minute` playlist는 `h/j/k/l`, survival, word motion, file navigation, substitute까지 포함한 legacy 17개 exercise 스냅샷으로 retired 상태이며, default playable path는 8문항 이하 tutorial episode playlist들을 실행한다. Substitute 계열은 중반 고급 튜토리얼로 분리했다.

### Beat 1. Panic Exit

```yaml
content_beat:
  id: first-5-minute-001
  role: safety
  command_cluster: survival-save-quit
  exercise_id: survival-save-quit-001
  scenario_id: survival-save-quit-001-scenario
  engine_support: planned
  player_lesson: "수정하지 않을 때는 변경을 버리고 나갈 수 있다."
  loader_needs:
    - command-line key trace 표현
    - app exit expectation
    - progress를 만들지 않는 smoke assertion
```

발견: `:q!`와 `:wq`는 일반 key 단위 입력만으로는 부족하다. loader는 command-line sequence를 key trace로 표현할 수 있어야 하고, runtime은 “성공 후 app 종료”와 “성공 후 다음 beat”를 구분해야 한다.

### Beat 2. First Cursor Target

```yaml
content_beat:
  id: first-5-minute-002
  role: first_success
  command_cluster: normal-motion-basic
  exercise_id: normal-motion-basic-001
  scenario_id: normal-motion-basic-001-scenario
  engine_support: implemented
  player_lesson: "`l`로 오른쪽 목표까지 이동한다."
  loader_needs:
    - cursor target assertion
    - optimal key trace
    - success feedback copy
```

발견: 현재 playable slice와 가장 가까운 콘텐츠다. CONTENT-001은 이 beat를 파일 기반 데이터로 로드해 현재 hardcoded playable exercise를 대체할 수 있어야 한다.

### Beat 3. Two-Dimensional Movement

```yaml
content_beat:
  id: first-5-minute-003
  role: spatial_practice
  command_cluster: normal-motion-basic
  exercise_id: normal-motion-basic-002
  scenario_id: normal-motion-basic-002-scenario
  engine_support: implemented
  player_lesson: "`h/j/k/l` 네 방향을 짧은 거리에서 조합한다."
  loader_needs:
    - multi-line buffer
    - cursor row/col target
    - boundary mistake feedback
```

발견: 단일 줄 이동만으로는 `j/k`를 학습할 수 없다. loader는 multi-line buffer와 cursor target을 안정적으로 표현해야 한다.

### Beat 4. Word Jump Teaser

```yaml
content_beat:
  id: first-5-minute-004
  role: efficiency_teaser
  command_cluster: word-motion-basic
  exercise_id: word-motion-basic-001
  scenario_id: word-motion-basic-001-scenario
  engine_support: planned
  player_lesson: "한 글자씩 걷는 것보다 단어 단위 이동이 빠르다."
  loader_needs:
    - compatibility tier 표시
    - oracle comparison eligibility
    - efficiency comparison copy
```

발견: `word-motion-basic`은 첫 5분 끝에서 “다음에 배울 맛보기”로 매우 좋지만, 엔진 구현 전에는 playable로 연결하면 안 된다. loader는 콘텐츠를 읽더라도 `engine_support`나 command support matrix로 실행 가능 여부를 판단해야 한다.

### Beat 5. Save And Continue

```yaml
content_beat:
  id: first-5-minute-005
  role: closure
  command_cluster: survival-save-quit
  exercise_id: survival-save-quit-002
  scenario_id: survival-save-quit-002-scenario
  engine_support: planned
  player_lesson: "성공한 변경은 저장하고 종료한다."
  loader_needs:
    - completion action
    - mission completion metadata
    - progress reward copy
```

발견: 저장/종료는 단순 buffer target이 아니라 mission completion과 progress 저장을 연결한다. CONTENT-001은 progression metadata를 schema 밖으로 미루더라도, scenario가 어느 mission/progress event에 속하는지는 표현할 수 있어야 한다.

## Required Content Files

CONTENT-001은 최소한 아래 네 종류를 고려한다.

```text
content/
  command_clusters/*.yaml
  exercises/*.yaml
  scenarios/*.yaml
  playlists/*.yaml
```

### command_clusters

필수 필드:

- `id`
- `status`
- `compatibility_tier`
- `engine_support`: `implemented | planned | unsupported`
- `commands`
- `purpose`
- `prerequisite`
- `difficulty`
- `common_mistakes`
- `oracle`: `required | optional | none`

### exercises

필수 필드:

- `id`
- `status`
- `command_cluster`
- `engine_support`
- `title`
- `goal_for_player`
- `initial_state.mode`
- `initial_state.cursor`
- `initial_state.buffer`
- `target_state.mode`
- `target_state.cursor`
- `target_state.buffer`
- `optimal_keys`
- `allowed_keys`
- `forbidden_keys`
- `hints`
- `grading.pass_condition`
- `grading.optimal_key_count`
- `e2e_assertions`

발견된 추가 요구:

- `optimal_keys`는 사람이 쓰기 쉬운 문자열과 runner가 쓰기 쉬운 배열 표현 중 하나로 정규화되어야 한다.
- `target_state.buffer`는 선택 값이어야 한다. cursor-only exercise가 필요하다.
- `forbidden_keys`는 과거 호환 필드로 유지하며, compile 시 `constraints.forbidden_keys`와 합쳐 runtime enforcement에 사용한다.
- command-line sequence는 `":"`, `"q"`, `"!"`, `"enter"`처럼 key trace로 풀어 저장해야 한다.
- `constraints.max_inputs`
- `constraints.required_keys`
- `constraints.forbidden_keys`
- `constraints.attempt_limit`

### scenarios

필수 필드:

- `id`
- `status`
- `exercise_id`
- `mission_title`
- `briefing`
- `context_role`
- `mentor_success`
- `mentor_failure`
- `story_constraints`

발견된 추가 요구:

- `mentor_hint` 또는 hint copy override가 필요할 수 있다.
- scenario가 exercise의 target을 바꾸지 않는다는 validation이 필요하다.
- same exercise를 다른 scenario skin으로 재사용할 수 있게 `exercise_id`와 narrative copy를 분리한다.
- TUI에 표시할 짧은 제목과 문서용 설명을 분리할 수 있어야 한다.

### playlists

필수 필드:

- `id`
- `status`
- `title`
- `beats`
- `unlock_policy`
- `completion_policy`

발견된 추가 요구:

- first playable에서는 단일 beat만 실행해도 되지만, 첫 5분 루프에는 episode 순서와 unlock이 필요하다.
- `engine_support: planned` beat는 playlist에 남길 수 있지만 현재 playable build에서는 skip 또는 locked로 표시해야 한다.
- approved/implemented tutorial playlist는 8문항 이하이어야 한다.
- retired playlist는 콘텐츠 이력 보관용이며 playable 후보에서 제외한다.

## CONTENT-001 Acceptance Draft

Status: implemented by `docs/exec-plans/completed/content-001-yaml-loader.md` and follow-up content validation slices.

- [draft] loader는 repo root `content/` 아래의 YAML 파일을 우선 읽어야 한다.
- [draft] loader는 command cluster, exercise, scenario, playlist 파일을 각각 읽을 수 있어야 한다.
- [draft] loader는 `engine_support: planned` 콘텐츠를 로드하되 playable 후보에서는 제외할 수 있어야 한다.
- [draft] loader는 `normal-motion-basic-001`을 현재 hardcoded playable exercise와 같은 compiled exercise로 변환할 수 있어야 한다.
- [draft] loader는 exercise가 참조하는 command cluster가 없으면 오류를 반환해야 한다.
- [draft] loader는 scenario가 참조하는 exercise가 없으면 오류를 반환해야 한다.
- [draft] loader는 cursor target이 buffer 범위를 벗어나면 오류를 반환해야 한다.
- [draft] loader는 E2E assertion에 필요한 buffer/cursor/mode/status/score/progress 기대값을 보존해야 한다.
- [draft] loader는 approved exercise가 approved 또는 implemented command cluster만 참조하도록 검증해야 한다.
- [draft] loader는 approved scenario가 approved 또는 implemented exercise만 참조하도록 검증해야 한다.
- [draft] loader는 optimal key trace를 정규화하고 `optimal_key_count`와 길이가 다르면 오류를 반환해야 한다.
- [draft] loader는 optimal key가 allowed keys에 없거나 forbidden keys와 충돌하면 오류를 반환해야 한다.
- [draft] loader는 command cluster의 `coverage_required`가 approved exercise들의 optimal trace에서 모두 등장하는지 보고해야 한다.
- [draft] loader 또는 후속 검증기는 `replay_status: pass` 없이 approved/implemented exercise를 playable로 올리지 않아야 한다.

## Open Questions

- JSON import/export를 언제 지원할지 결정이 필요하다.
