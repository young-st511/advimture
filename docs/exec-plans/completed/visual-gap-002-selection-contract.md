# VISUAL-GAP-002 — Selection State Contract

Slice-ID: VISUAL-GAP-002
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/active/visual-gap-002-selection-contract.md
- docs/exec-plans/completed/visual-gap-002-selection-contract.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/verification/selection-app-state-contract.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md

## 목표

visual mode 구현 전에 selection state, TUI 표시, content/E2E app_state assertion 계약을 같은 용어로 고정한다.

## 완료 내용

- 첫 구현 범위를 charwise `v`로 제한했다.
- selection state를 `active`, `kind`, `anchor`, `head`, `start`, `end`로 고정했다.
- `start`/`end`를 normalized inclusive range로 정의했다.
- TUI 최소 표시를 `Mode: visual`, selection range line, selected non-cursor cell `{x}`로 정의했다.
- E2E는 화면 텍스트가 아니라 app_state `selection` object로 selection을 검증하도록 고정했다.
- `V`, visual block, operator application은 후속 slice로 분리했다.

## 검증 결과

- passed: `git diff --check`
