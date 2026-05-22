# VISUAL-LINE-GAP-001 — Linewise Visual Scope

Slice-ID: VISUAL-LINE-GAP-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/completed/visual-line-gap-001-linewise-scope.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- docs/gameplay/vim-curriculum-map.md
- docs/verification/selection-app-state-contract.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md

## 목표

linewise `V`와 multi-line visual을 구현하기 전에 첫 범위, 제외 항목, 검증 표면을 결정한다.

## 결정

- 다음 visual 구현 후보는 `visual-line-basic`이다.
- 첫 범위는 linewise `V` + row motion + `d/y`로 제한한다.
- multi-line charwise `v` + `d/y`는 linewise 구현 이후로 미룬다.
- visual block `<C-v>`는 장기 보류한다.
- linewise selection은 app_state selection object의 `kind: linewise`를 사용한다.
- `anchor`와 `head`는 cursor row를 보존하되, normalized `start`는 선택 첫 줄의 col 0, `end`는 선택 마지막 줄의 마지막 col로 표현한다. empty line의 `end.col`은 0으로 둔다.
- linewise `d`는 선택 줄 전체를 삭제하고 unnamed linewise register에 저장한다.
- linewise `y`는 선택 줄 전체를 unnamed linewise register에 저장하고 buffer를 변경하지 않는다.
- 성공한 `d/y`는 normal mode로 돌아가고 selection을 clear한다.
- cursor landing은 normalized start row의 col 0으로 둔다. 선택 삭제가 파일 끝까지 닿으면 남은 마지막 줄 또는 빈 fallback 줄의 col 0으로 clamp한다.

## 첫 구현 포함 후보

- `V` 진입과 visual mode 중 `V`/`esc` 해제
- visual-line mode에서 `j/k`, `gg/G` row motion
- 같은 줄 linewise `Vd`, `Vy`
- 여러 줄 linewise `Vj`, `Vk`, `VG`, `Vgg` selection
- linewise `d/y` 후 기존 `p/P`와의 linewise register 호환성
- TUI 표시: `Selection: linewise <start.row>,0 -> <end.row>,<end.col>`

## 제외 항목

- multi-line charwise `v` operator
- visual block `<C-v>`
- visual `c`, `>`, `<`, `p`, `P`
- count prefix, register prefix
- `gv`, `o`, `O` endpoint swap
- dot repeat 연계
- indentation command
- mouse/terminal selection 연동

## 후속 실행 후보

1. `VISUAL-LINE-001`: linewise `V` engine/runtime/TUI scope approval
2. `VIM-029`: linewise visual selection foundation and `d/y`
3. `PLAYPACK-011`: 3문항 이하 linewise visual tutorial
4. `INCIDENT-RUN-004`: linewise visual을 적용하는 config block 복구 run

## 검증 계획

- engine unit: forward/backward line range, delete all lines fallback, yank register kind, cursor landing
- runtime unit: required key `V`, `d/y`, goal success와 unsupported key preservation
- TUI/app_state: `kind: linewise`, normalized start/end, selection clear after operator
- content replay: `e2e_assertions.selection`이 linewise kind를 비교
- E2E: tutorial full run에서 buffer/cursor/mode/progress/key trace 검증

## 검증 결과

- passed: `git diff --check`
