# Scenario Production Harness

> `Vim command -> Exercise -> Scenario` 제작 워크플로우를 높은 품질로 반복하기 위한 하네스다. 이 문서는 콘텐츠를 직접 만들기보다, Agent가 콘텐츠를 만들 때 통과해야 하는 절차와 검증 기준을 정의한다.

## 핵심 규칙

1. 시나리오 제목이나 세계관에서 시작하지 않는다.
2. 모든 산출물은 하나 이상의 command cluster에 연결된다.
3. command cluster는 실무 유용성, 선행 관계, 다음 조합 가능성을 가져야 한다.
4. exercise는 scenario 없이도 독립적으로 풀리고 검증 가능해야 한다.
5. scenario는 exercise의 목표 상태, 정답 키, 허용 키를 바꾸지 않는다.
6. 현재 엔진이 지원하지 않는 command는 `engine_support: planned`로 남기고 playable 후보에서 제외한다.
7. `OK`는 작성 Agent가 아니라 별도 검증 Agent가 준다.
8. approved/implemented exercise는 optimal key replay로 target state와 일치해야 한다.
9. cluster 안의 모든 command는 최소 1개 exercise의 optimal key trace에 실제로 등장해야 한다.

## Roles

### Producer Agent

책임:

- command 후보를 curriculum map의 cluster로 환원한다.
- command cluster draft를 작성한다.
- approved 또는 review 대상 command에 대해 exercise draft를 만든다.
- exercise를 scenario skin으로 감싼다.
- 각 단계의 self-check 결과를 기록한다.

금지:

- scenario를 먼저 만들고 나중에 command를 끼워 맞추기
- 기계 검증 불가능한 목표를 exercise로 승인하기
- `engine_support: planned`를 playable로 연결하기
- 기존 archived 구현을 canonical source처럼 인용하기

### Verifier Agent

책임:

- 산출물이 command-first 순서를 지켰는지 확인한다.
- Vim 실전 유용성과 선행 관계가 빈약한 cluster를 reject한다.
- 정답/목표/허용 키가 모호한 exercise를 reject한다.
- scenario가 학습 목표를 흐리거나 목표 상태를 바꾸면 reject한다.
- OK/REVISE/BLOCKED 중 하나로 판정한다.

### Human

책임:

- `draft` 항목을 `approved`로 승격할지 결정한다.
- 제품 톤, 난이도, 커리큘럼 우선순위의 최종 판단을 한다.
- 저장 포맷, 의존성, 엔진 범위 변경을 승인한다.

## Iteration Protocol

```text
1. Producer Agent drafts command cluster candidates.
2. Verifier Agent reviews command layer only.
3. If REVISE, Producer updates command layer and repeats.
4. If OK, Producer drafts exercises for OK command clusters.
5. Verifier reviews exercise layer only.
6. If REVISE, Producer updates exercises and repeats.
7. If OK, Producer drafts scenario skins.
8. Verifier reviews scenario layer.
9. If OK, Human may approve by removing draft status.
```

반복 제한:

- 같은 layer에서 3회 연속 `REVISE`면 사람에게 질문한다.
- `BLOCKED`는 엔진 미지원, 제품 결정 필요, 저장 포맷 변경 같은 이유가 있을 때만 사용한다.
- 질문은 한 번에 1-3개로 제한하고, 구현을 막는 결정만 묻는다.

## OK Gate: Command

필수 조건:

- `id`, `status`, `compatibility_tier`, `engine_support`, `commands`, `purpose`, `prerequisite`, `difficulty`가 있다.
- 하나의 학습 목표로 묶인다.
- 실무에서 언제 쓰는지 최소 2개 상황이 있다.
- 선행 cluster가 명시되어 있다.
- 다음 조합 경로가 있다.
- 초보자 실수와 compatibility note가 있다.

Reject 조건:

- 단순 명령어 나열이다.
- “재미있을 것 같아서”가 주된 선정 이유다.
- 이전 학습 없이 사용할 수 없는 command인데 prerequisite이 비어 있다.
- engine support가 planned인데 implemented처럼 다룬다.
- exact tier인데 oracle 검증 전략이 없다.
- cluster의 일부 command가 allowed key에만 있고 optimal path에 등장하지 않는다.

## OK Gate: Exercise

필수 조건:

- `command_cluster`가 존재한다.
- initial state와 target state가 있다.
- target state는 cursor/mode/buffer 중 하나 이상을 기계 검증 가능하게 가진다.
- optimal key trace가 있고 key count가 계산 가능하다.
- allowed keys와 forbidden keys가 있다.
- 최소 2단계 힌트가 있다.
- pass condition이 app state나 buffer/cursor state로 판정 가능하다.
- 같은 command cluster 안에서 최소 2-3개 반복 변주를 만들 수 있다.
- optimal key trace를 현재 엔진 또는 oracle로 replay했을 때 target state와 일치한다.
- optimal key count가 normalized key trace 길이와 일치한다.
- optimal key는 모두 allowed keys 안에 있고 forbidden keys와 충돌하지 않는다.

Reject 조건:

- “적절히 이동한다”, “자연스럽게 고친다”처럼 관찰 불가능한 목표다.
- optimal key trace가 Vim 동작과 맞지 않는다.
- 허용 키가 학습 목표를 우회하게 만든다.
- 정답이 scenario 문맥을 알아야만 이해된다.
- 지나치게 긴 `hjkl` 반복으로 빠른 motion 학습을 방해한다.
- cluster command 중 일부가 실제 최적해에서 한 번도 훈련되지 않는다.

## OK Gate: Scenario

필수 조건:

- `exercise_id`를 참조한다.
- briefing이 exercise의 player goal과 일치한다.
- success/failure feedback이 사용한 Vim 개념을 다시 언급한다.
- learning reinforcement가 명시되어 있다.
- `does_not_change` 체크가 target/keys/pass condition을 보호한다.
- story constraints가 command 학습을 보호한다.
- scenario를 제거해도 exercise가 독립적으로 성립한다.

Reject 조건:

- scenario가 exercise target state를 바꾼다.
- 세계관 설명이 목표보다 길고 조작을 흐린다.
- 실패 피드백이 농담이나 분위기만 있고 회복 단서를 주지 않는다.
- command 학습보다 텍스트 내용 이해가 더 어려워진다.

## OK Gate: Playable Readiness

필수 조건:

- 관련 command cluster가 `approved` 이상이다.
- exercise가 `approved` 이상이다.
- scenario가 `approved` 이상이다.
- 모든 command가 현재 `engine_support: implemented`다.
- E2E key trace와 app state assertion이 있다.
- progress 영향이 있다면 저장 포맷 변경이 필요 없는지 확인했다.
- lifecycle gate를 통과했다: approved scenario는 approved/implemented exercise만 참조하고, approved exercise는 approved/implemented command cluster만 참조한다.

Reject 조건:

- `planned` command를 playable playlist에 넣는다.
- E2E assertion이 screen text만 있고 cursor/buffer/status를 보지 않는다.
- progress file이나 app state summary를 실제 HOME에 쓰게 만든다.

## SubAgent Prompt Templates

### Producer

```text
Advimture의 Vim 학습 콘텐츠를 command -> exercise -> scenario 순서로 설계하세요.
먼저 command cluster만 작성하고, scenario 제목이나 세계관에서 시작하지 마세요.
각 cluster는 실무 유용성, prerequisite, combo_paths, common_mistakes, engine_support를 포함해야 합니다.
산출물은 draft 상태로 두고, 구현 가능/계획 필요를 분리하세요.
```

### Verifier

```text
Advimture 콘텐츠 산출물을 검증하세요.
command-first 순서, 기계 검증 가능성, Vim 실전 유용성, scenario가 exercise 목표를 바꾸지 않는지 확인하세요.
판정은 OK / REVISE / BLOCKED 중 하나로 내리고, REVISE면 반드시 수정해야 할 항목만 짧게 제시하세요.
```

## Output Contract

Producer는 다음 형식으로 보고한다.

```yaml
production_report:
  layer: command | exercise | scenario
  status: draft
  inputs:
    - <source>
  outputs:
    - <doc path or item id>
  self_check:
    command_first: pass | fail
    mechanically_verifiable: pass | fail | n/a
    engine_support_separated: pass | fail
  open_questions:
    - <only blocking questions>
```

Verifier는 다음 형식으로 보고한다.

```yaml
verification_report:
  verdict: OK | REVISE | BLOCKED
  layer: command | exercise | scenario
  must_fix:
    - <blocking issue>
  should_fix:
    - <quality issue>
  missing_decisions:
    - <human decision needed>
```

## Coverage Matrix

각 콘텐츠 팩은 아래 matrix를 함께 제출해야 한다.

```text
cluster_id
  commands: [h, j, k, l]
  trained_in_optimal_trace:
    h: [normal-motion-basic-004]
    j: [normal-motion-basic-002, normal-motion-basic-003]
    k: [normal-motion-basic-005]
    l: [normal-motion-basic-001, normal-motion-basic-003]
  edge_cases:
    - line boundary
    - short line clamp
  replay_status: pass | fail | pending
  oracle_status: required | optional | none | pending
```

allowed key에 포함된 것만으로는 훈련된 것으로 보지 않는다. command는 optimal trace에 등장해야 coverage로 인정한다.

## Hint Policy

- 1단계 힌트는 개념 또는 방향을 알려준다.
- 2단계 힌트는 막힌 플레이어를 회복시키기 위해 직접 키를 알려줄 수 있다.
- 3단계가 필요한 경우에는 정답 키보다 “왜 그 키인지”를 설명한다.
- 힌트 사용은 scoring penalty에 반영한다.
