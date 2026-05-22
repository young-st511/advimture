# VISUAL-OP-001 — Charwise Visual Operator Scope

Slice-ID: VISUAL-OP-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/active/visual-op-001-charwise-operator-scope.md
- docs/exec-plans/completed/visual-op-001-charwise-operator-scope.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md

## 목표

charwise visual selection foundation 다음 단계로, visual selection에 적용할 operator 범위와 제외 항목을 고정한다.

## 결정

- 다음 engine slice는 `VIM-028`이다.
- 포함 범위는 charwise visual selection에 `d`와 `y`를 적용하는 것이다.
- `d`는 inclusive selection range를 삭제하고 normal mode로 돌아간다.
- `y`는 inclusive selection range를 unnamed register에 저장하고 normal mode로 돌아간다.
- 첫 구현은 같은 줄 selection만 대상으로 한다.
- multi-line charwise selection은 selection state로는 표현되더라도 operator 적용은 후속 hardening으로 분리한다.
- visual operator 결과는 app_state `mode`, `buffer`, `cursor`, `selection`, 필요 시 register unit test로 검증한다.

## 제외 항목

- linewise `V`
- visual block `<C-v>`
- multi-line charwise delete/yank
- `c`, `>`, `<`, `p`, `P` visual mode operator
- count/register prefix
- `gv`, `o`, `O` selection endpoint swap
- dot repeat 연계

## 수용 기준

- VISUAL-OP-001은 VIM-028과 PLAYPACK-010의 범위를 분리한다.
- VIM-028은 engine/runtime/TUI 기반만 다루며 content/playpack은 추가하지 않는다.
- PLAYPACK-010은 VIM-028 통과 이후 3~4문항 visual tutorial로 별도 진행한다.

## 검증 결과

- passed: `git diff --check`
