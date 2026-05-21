# VISUAL-GAP-001 — Visual Mode Scope

Slice-ID: VISUAL-GAP-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/completed/visual-gap-001-visual-mode-scope.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md

## 목표

다음 대형 Vim capability 후보인 visual mode를 바로 구현하지 않고, 최소 scope와 engine/TUI 영향도를 문서로 분리한다.

## 결정

- 다음 구현 후보는 `visual-char-line` gap planning이다.
- 첫 visual slice 후보는 `v`, `V`, `d`, `y`, `>` 중에서 다시 좁힌다.
- `<C-v>` visual block은 첫 visual slice에서 제외한다.
- selection rendering, cursor anchor, operator application, E2E app_state 확장이 필요하므로 engine-only slice 전에 gap planning을 둔다.
- visual mode는 text object보다 UI 표시 의존도가 크므로 focused TUI E2E가 필수다.

## 후보 범위

- 최소 후보 A: `v` + charwise selection + `d`/`y`
- 최소 후보 B: `V` + linewise selection + `d`/`y`
- 보류 후보: `<C-v>`, block insert, indentation, multi-cursor-like edits

## 제외 항목

- visual block
- count prefix
- register prefix
- indentation command
- selection highlight 고도화
- mouse/terminal selection 연동

## 수용 기준

- completed: command catalog에 visual 후보가 draft로 기록된다.
- completed: roadmap에는 visual 구현이 다음 중기 플랜으로 분리된다.
- completed: visual mode 구현 전에 필요한 engine/TUI/E2E 영향을 명시한다.

## 검증 결과

- passed: `git diff --check`
