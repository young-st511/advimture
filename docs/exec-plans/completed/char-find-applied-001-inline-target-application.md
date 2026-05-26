# CHAR-FIND-APPLIED-001 — Inline Target Application Candidate

Slice-ID: CHAR-FIND-APPLIED-001
Created: 2026-05-26
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/gameplay/scenario-tone.md
- docs/exec-plans/active/char-find-applied-001-inline-target-application.md

## 목표

`char-find-line` tutorial 직후에 바로 새 engine을 늘리지 않고, 배운 `f/t`와 operator 조합을 실제 runbook 판단 상황에 적용할 후보를 고른다.

## 포함

- command-choice drill 후보 1개
- incident beat 후보 1개
- 각 후보의 학습 의도, 필요한 command 조합, E2E 검증 방식
- 다음 구현 slice 분리 여부 판단

## 제외

- 새 Vim engine 기능
- content YAML 구현
- E2E runner 변경
- progress 저장 포맷 변경

## 수용 기준

- 적용 후보는 기존 구현 command만 사용한다.
- `f/t`가 “빠른 이동”이 아니라 “적절한 범위 선택” 문제로 쓰인다.
- command-choice와 incident 중 어느 쪽을 먼저 구현할지 명시한다.
- 다음 구현 slice가 필요하면 별도 ExecPlan 후보로 분리한다.

## Step 1: Candidate Ideation

- [x] command-choice 적용 후보 작성
- [x] incident 적용 후보 작성

## Step 2: Selection

- [x] 학습 가치와 구현 리스크 비교
- [x] 추천 후보 결정

## Step 3: Next Slice Split

- [x] 구현 범위와 제외 항목 정의
- [x] 검증 계획 정의

## 결정

첫 적용 구현은 `choice-005-inline-target-range`를 `incident-005-command-choice`의 third beat로 추가한다. `ct,`가 의도 선택이며, `cf,`와 `cw`는 각각 delimiter 삭제/부분 단어 변경 때문에 부적절하다는 판단을 훈련한다.

후속 incident 후보는 `incident-006-inline-target-repair`로 남긴다. 이 후보는 `/target`으로 줄을 찾은 뒤 `ct,`로 comma 앞 값만 바꾸는 applied run이다.

## 다음 구현 Slice 후보

```yaml
exec_plan_candidate:
  slice_id: CHOICE-PLAY-002
  goal: incident-005-command-choice에 inline target range choice beat를 추가한다.
  include:
    - command-choice-inline-target-001 exercise
    - command-choice-inline-target-001-scenario
    - incident-005-command-choice third beat
    - focused E2E update
  exclude:
    - 새 engine 기능
    - 새 playlist category
    - progress 저장 포맷 변경
    - search + inline target 복합 incident
  acceptance_criteria:
    - optimal trace에 ct,가 들어간다.
    - cf, 우회는 forbidden route로 막는다.
    - command-choice focused E2E가 3-beat 흐름을 검증한다.
```
