# VIM-027-TUI-003 — Charwise Visual Foundation

Slice-ID: VIM-027-TUI-003
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: engine-and-tui
Allowed-Paths:
- docs/exec-plans/active/vim-027-tui-003-charwise-visual-foundation.md
- docs/exec-plans/completed/vim-027-tui-003-charwise-visual-foundation.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/vimengine/engine.go
- internal/vimengine/engine_test.go
- internal/tuiadapter/adapter.go
- internal/tuiadapter/adapter_test.go
- internal/playable/model.go
- internal/playable/model_test.go

## 목표

VISUAL-GAP-002/E2E-007 계약에 맞춰 charwise `v` visual mode foundation과 최소 TUI 표시를 구현한다.

## 완료 내용

- `vimengine.ModeVisual`과 charwise `Selection` state를 추가했다.
- normal mode `v`가 visual mode로 진입하고 현재 cursor를 anchor/head로 잡는다.
- visual mode motion key가 cursor/head를 이동하고 normalized inclusive range를 갱신한다.
- visual mode `esc`와 `v`가 normal mode로 돌아가며 selection을 clear한다.
- `tuiadapter.ViewModel`이 selection을 전달한다.
- playable TUI가 selection range line과 `{x}` selected cell 표시를 제공한다.
- `Model.State()`가 selection을 E2E state summary로 전달한다.

## 제외한 것

- visual selection에 `d`/`y` operator 적용
- linewise `V`
- visual block `<C-v>`
- visual mode content/playpack

## 검증 결과

- passed: `go test ./internal/vimengine/...`
- passed: `go test ./internal/tuiadapter/...`
- passed: `go test ./internal/playable/...`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
