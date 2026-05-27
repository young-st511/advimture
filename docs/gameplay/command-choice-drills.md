# Command Choice Drills

> 배운 Vim command 중 "무엇을 써야 하는가"를 훈련하는 mixed drill 설계 기준이다.

## 목적

Command choice drill은 새 Vim command를 소개하지 않는다. 이미 학습한 command cluster를 한 상황에 섞어 두고, 플레이어가 편집 범위와 반복 패턴을 보고 적절한 도구를 고르게 만든다.

이 drill의 핵심 질문은 다음이다.

- 한 글자, 단어, 따옴표 내부, 줄 묶음 중 어느 범위를 고쳐야 하는가?
- 직접 다시 입력하는 것이 좋은가, yank/put 또는 `.` 반복이 좋은가?
- 먼저 `/`로 대상을 찾고 편집해야 하는가?
- `:s`, `:%s`, range substitute가 visual/operator 조작보다 더 적합한가?

## 설계 원칙

- 설계 순서는 기존 제품 철학처럼 `Command set -> Exercise choice -> Scenario skin`을 따른다.
- command choice drill은 `command_cluster`가 아니라 적용 레이어다. 새 engine 기능이나 progress 저장 필드를 요구하지 않는다.
- 첫 playable 승격 전까지 새 YAML schema 필드를 추가하지 않는다. 필요한 정보는 설계 문서와 scenario/hint 문구로 표현한다.
- playable 콘텐츠로 만들 때는 기존 `constraints.required_keys`와 `constraints.forbidden_keys`로 의도한 선택을 고정한다.
- 문제 첫 문장은 정답 key sequence를 직접 말하지 않는다. 실패/힌트 단계에서는 선택 이유를 설명할 수 있다.
- 정답 암기보다 "왜 이 도구가 이 상황에 맞는가"를 성공/실패 피드백에 남긴다.

## Drill Types

| Type | 훈련 질문 | 후보 command | 적합한 상황 |
|------|----------|--------------|-------------|
| `scope-choice` | 편집 범위가 무엇인가? | `x`, `r`, `ciw`, `ci"`, `v...d`, `V...d` | 문자 하나, 단어, quote 값, 줄 묶음이 섞여 있을 때 |
| `reuse-choice` | 같은 내용을 다시 쓰지 않고 재사용할 수 있는가? | `yy/p`, `yw/P`, `yi"/P`, `.` | 같은 줄, 단어, quote 값, 반복 변경이 이어질 때 |
| `search-then-act` | 먼저 찾아야 하는가? | `/`, `n`, `N` + edit command | 로그/설정에서 target token 위치가 멀거나 여러 개인 경우 |
| `range-choice` | 범위 명령이 더 좋은가? | `:s`, `:%s`, `:2,3s`, `V...d` | 한 줄, 전체 파일, 줄 범위, 블록 삭제 중 선택해야 할 때 |
| `inline-target-choice` | 같은 줄 delimiter를 기준으로 어디까지 고쳐야 하는가? | `ct,`, `cf,`, `f=`, `t,` | hyphenated 값이나 delimiter 보존이 필요한 한 줄 설정을 고칠 때 |

## Authoring Rubric

Playable YAML에 새 schema를 추가하기 전까지, 기획자는 각 command choice 후보를 다음 형태로 검토한다.

```yaml
choice_drill_draft:
  id: choice-001-scope-triage
  trained_clusters:
    - text-object-inner-word
    - text-object-quote-pair
    - visual-line-basic
  choice_focus: scope-choice
  viable_options:
    - ciw
    - ci"
    - Vd
  intended_option: Vd
  why_intended: 삭제 대상이 단어가 아니라 연속된 설정 줄 묶음이다.
  forbidden_shortcuts:
    - repeated-dd
    - manual-retype
  coaching_copy: 먼저 고칠 범위가 글자/단어/줄 중 무엇인지 판단하세요.
```

이 구조는 당장 loader가 읽는 canonical schema가 아니다. `COMMAND-CHOICE-001`은 이 정보를 콘텐츠 제작자가 공유하는 작성 계약으로 고정하고, 실제 playable 승격은 별도 ExecPlan에서 결정한다.

## First Drill Candidates

| ID | Focus | Intended choice | 후보 상황 |
|----|-------|-----------------|-----------|
| `choice-001-scope-triage` | `scope-choice` | `V...d` | 정상 route 값 사이의 quarantine block을 보고 값/단어가 아니라 줄 묶음을 복구 범위로 먼저 구분한다. Playable: `incident-005-command-choice` / `command-choice-scope-001`. |
| `choice-002-repeat-or-substitute` | `range-choice` | `:%s` | 같은 literal이 파일 전체와 한 줄 안에 반복되어 `.`보다 substitute가 자연스럽다. Playable: `incident-005-command-choice` / `command-choice-repeat-substitute-001`. |
| `choice-003-copy-or-retype` | `reuse-choice` | `yi"` + `P` | 이미 검증된 quote 값이 있고, 새 위치에 그대로 붙여야 한다. |
| `choice-004-search-then-scope` | `search-then-act` | `/token` + `V...d` | 먼저 marker를 찾은 뒤 marker 아래 블록을 linewise로 격리한다. |
| `choice-005-inline-target-range` | `inline-target-choice` | `ct,` | comma 뒤 route는 정상이고 comma는 보존해야 한다. hyphenated 값만 바꾸려면 `cw`나 `cf,`가 아니라 `ct,`가 적합하다. Playable: `incident-005-command-choice` / `command-choice-inline-target-001`. |
| `incident-006-inline-target-repair` | `search-then-act` + `inline-target-choice` | `/target` + `ct,` | 먼저 손상된 target 값을 찾고, 같은 줄 comma 앞 값만 교체한다. Playable: `incident-006-inline-target-repair`. |
| `choice-006-quote-value-reuse` | `reuse-choice` | `yi"` + `P` | 검증된 quote 내부 token을 빈 quote 위치에 그대로 복제해야 한다. 직접 재입력은 token 길이/오타 리스크가 크다. Next playable candidate: `incident-005-command-choice` fourth beat. |
| `choice-007-line-reuse` | `reuse-choice` | `V` + `y` + `p` | 검증된 route 줄 전체를 다음 위치에 복제해야 한다. 단어/quote 값이 아니라 줄 전체 재사용 문제다. |
| `choice-008-repeat-change-reuse` | `reuse-choice` | `.` | 같은 quote 값 변경이 여러 줄에 반복된다. 재입력보다 last change repeat가 적합하다. |

## Inline Target Application Decision

`CHAR-FIND-APPLIED-001`의 첫 적용 후보는 `choice-005-inline-target-range`로 한다.

이유:

- command-choice의 본래 목적이 “무슨 도구를 써야 하는가”라서 `ct,`와 `cf,`의 범위 판단을 가장 직접적으로 훈련한다.
- 새 incident를 만들기 전 기존 `incident-005-command-choice`에 1 beat를 추가하면 흐름과 E2E blast radius가 작다.
- `incident-006-inline-target-repair`는 `/target`, `n`, `ct,`를 조합하는 applied run으로 승격했다. 다음 command-choice 후보는 직접 재입력 대신 기존 텍스트를 재사용할지 판단하는 `reuse-choice`다.

## Reuse Choice Decision

`REUSE-CHOICE-001`의 첫 적용 후보는 `choice-006-quote-value-reuse`로 한다.

이유:

- 기존 engine만으로 `yi"` + `P`를 안정적으로 검증할 수 있고, quote 내부 값 재사용은 실무 config/token 편집에서 자주 등장한다.
- `choice-007-line-reuse`는 이미 linewise tutorial/incident에서 비슷한 감각을 여러 번 다뤘다.
- `choice-008-repeat-change-reuse`는 `.`의 효용이 크지만 “재사용”보다 “반복 변경” 판단에 가까워 별도 repeat-choice로 분리하는 편이 낫다.

첫 playable 구현은 `incident-005-command-choice`에 fourth beat로 붙인다. 우회 방지는 `i/a/A/o/O`, `c/d/x/r`, `:`를 금지하고 `y`, `i`, `"`, `P`를 required key로 고정한다.

## Playable Gate

Command choice drill을 실제 content로 승격할 때는 다음을 만족해야 한다.

- 새 command를 소개하지 않고 기존 implemented cluster만 사용한다.
- 각 exercise는 `constraints.required_keys`로 의도 command를 포함한다.
- 우회가 학습 목표를 무너뜨리면 `constraints.forbidden_keys`로 막는다.
- 최종 buffer/cursor뿐 아니라 key trace coverage를 E2E에서 검증한다.
- scenario briefing은 `상황 1문장 + 선택 판단 목표 1문장`으로 제한한다.
- success/failure copy는 사용한 command 이름보다 선택 이유를 우선 설명한다.
