# VISUAL-LINE-001 — Linewise Visual Scope Approval

Slice-ID: VISUAL-LINE-001
Created: 2026-05-23
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/active/visual-line-001-linewise-scope.md
- docs/exec-plans/completed/visual-line-001-linewise-scope.md
- docs/gameplay/spec.md
- docs/verification/selection-app-state-contract.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md

## 목표

VIM-029 구현 전에 linewise `V`의 정확한 engine/runtime/TUI/E2E 수용 기준을 고정한다.

## 영향 도메인

- Gameplay: linewise visual은 새 command cluster 후보 `visual-line-basic`의 첫 구현이다.
- Verification: visual selection은 app_state selection object와 content replay assertion으로 검증해야 한다.

## 수용 기준

- normal mode `V`는 visual mode로 진입하고 selection `kind: linewise`를 만든다.
- linewise visual mode에서 `V`와 `esc`는 normal mode로 돌아가며 selection을 clear한다.
- linewise visual mode에서 `j/k/gg/G`는 cursor와 selection head를 row 단위로 이동한다.
- linewise selection의 normalized `start`는 첫 선택 줄의 col 0이다.
- linewise selection의 normalized `end`는 마지막 선택 줄의 마지막 col이며, empty line은 col 0이다.
- linewise `d`는 선택 줄 전체를 삭제하고 unnamed linewise register에 저장한다.
- linewise `y`는 선택 줄 전체를 unnamed linewise register에 저장하고 buffer를 변경하지 않는다.
- 성공한 linewise `d/y`는 normal mode로 돌아가고 selection을 clear한다.
- linewise `d/y` 후 cursor는 normalized start row의 col 0에 둔다. 삭제가 파일 끝까지 닿으면 남은 마지막 줄 또는 빈 fallback 줄의 col 0으로 clamp한다.
- TUI/app_state는 `selection.kind: linewise`를 노출하고 E2E에서 검증할 수 있어야 한다.

## 제외 항목

- multi-line charwise `v` operator
- visual block `<C-v>`
- visual `c`, `>`, `<`, `p`, `P`
- count prefix, register prefix
- `gv`, `o`, `O` endpoint swap
- dot repeat 연계
- indentation command
- mouse/terminal selection 연동

## 후속 구현 계획

- `VIM-029`: engine/runtime/TUI 구현과 단위 테스트
- `PLAYPACK-011`: 3문항 이하 linewise visual tutorial과 full E2E
- `INCIDENT-RUN-004`: linewise visual 적용 incident

## 검증 계획

- `git diff --check`

## 검증 결과

- planned: `git diff --check`
