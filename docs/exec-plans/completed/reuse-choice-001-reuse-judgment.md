# REUSE-CHOICE-001 — Reuse Judgment Drill Design

Slice-ID: REUSE-CHOICE-001
Created: 2026-05-28
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/reuse-choice-001-reuse-judgment.md

## 목표

직접 다시 입력하지 않고 이미 있는 텍스트를 재사용해야 하는 상황을 command-choice drill로 설계한다. 후보는 `yi" + P`, `yy/p`, `.`, retype이며, 첫 playable 구현 후보를 하나로 고른다.

## 포함

- reuse-choice 후보 2~3개
- intended command와 부적합한 대안의 이유
- forbidden shortcut 기준
- 다음 구현 slice 분리

## 제외

- content YAML 구현
- 새 Vim engine 기능
- 새 schema
- progress 저장 포맷 변경

## 수용 기준

- 첫 playable 후보는 이미 implemented command만 사용한다.
- 성공/실패 copy는 “왜 재사용해야 하는가”를 설명한다.
- 다음 구현 slice `CHOICE-PLAY-003`의 include/exclude/검증 계획이 정의된다.

## Step 1: Candidate Design

- [x] reuse-choice 후보 작성
- [x] 우회/금지 입력 기준 작성

## Step 2: Selection

- [x] 구현 리스크 비교
- [x] 첫 playable 후보 결정

## Step 3: Next Slice

- [x] CHOICE-PLAY-003 구현 범위 정의
- [x] 검증 계획 정의

## 결정

첫 playable 후보는 `choice-006-quote-value-reuse`로 한다. 정답은 `yi"` + `P`이며, 검증된 quote 내부 token을 빈 quote 위치에 그대로 복제해야 하는 상황으로 설계한다.

보류:

- `choice-007-line-reuse`: linewise 복제는 이미 tutorial/incident에서 반복 노출됐다.
- `choice-008-repeat-change-reuse`: dot repeat는 별도 repeat-choice 루프로 분리하는 편이 선명하다.

## 다음 구현 Slice 후보

```yaml
exec_plan_candidate:
  slice_id: CHOICE-PLAY-003
  goal: incident-005-command-choice에 quote value reuse-choice beat를 추가한다.
  include:
    - command-choice-quote-reuse-001 exercise
    - command-choice-quote-reuse-001-scenario
    - incident-005-command-choice fourth beat
    - focused E2E update
  exclude:
    - 새 engine 기능
    - 새 schema
    - linewise reuse-choice
    - dot repeat-choice
  acceptance_criteria:
    - optimal trace에 yi"와 P가 들어간다.
    - 직접 재입력/수정 command는 forbidden route로 막는다.
    - command-choice focused E2E가 4-beat 흐름을 검증한다.
```
