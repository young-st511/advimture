# VIM-029 — Linewise Visual Engine

Slice-ID: VIM-029
Created: 2026-05-23
Status: completed
Scope-Mode: engine-runtime-tui
Allowed-Paths:
- docs/exec-plans/active/vim-029-linewise-visual-engine.md
- docs/exec-plans/completed/vim-029-linewise-visual-engine.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- docs/verification/selection-app-state-contract.md
- internal/vimengine/selection.go
- internal/vimengine/visual.go
- internal/vimengine/engine.go
- internal/vimengine/engine_test.go
- internal/runtime/session_test.go
- internal/tuiadapter/adapter.go
- internal/tuiadapter/adapter_test.go
- internal/playable/model.go
- internal/playable/model_test.go

## 목표

VISUAL-LINE-001 수용 기준에 맞춰 linewise `V` visual mode와 linewise `d/y`를 구현한다.

## 수용 기준

- normal mode `V`가 visual mode로 진입하고 `selection.kind: linewise`를 만든다.
- linewise visual mode의 `V`/`esc`는 normal mode로 돌아가며 selection을 clear한다.
- linewise visual mode의 `j/k/gg/G`는 cursor/head를 row 단위로 이동하고 full-line normalized range를 갱신한다.
- linewise `d`는 선택 줄 전체를 삭제하고 unnamed linewise register에 저장한다.
- linewise `y`는 선택 줄 전체를 unnamed linewise register에 저장하고 buffer를 변경하지 않는다.
- linewise `d/y` 성공 후 normal mode, nil selection, cursor col 0을 보장한다.
- TUI와 app_state는 `selection.kind: linewise`를 노출한다.

## 제외 항목

- multi-line charwise `v` operator
- visual block `<C-v>`
- visual `c`, `>`, `<`, `p`, `P`
- count/register prefix
- dot repeat 연계
- content/playpack 추가

## 테스트 Red 계획

- vimengine:
  - `V` starts linewise selection
  - `Vj` normalizes full-line selection
  - backward `Vk` selection
  - linewise `y` stores linewise register without mutation
  - linewise `d` deletes selected lines, stores linewise register, clears selection
  - deleting all lines leaves one empty fallback line
  - `Vgg`/`VG` row motion works in linewise visual mode
- runtime:
  - linewise visual delete trace can satisfy goal and required keys
- tuiadapter/playable:
  - uppercase `V` maps to runtime key
  - linewise selection is rendered/exposed in view model

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/runtime/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `git diff --check`

## 완료 내용

- `SelectionLinewise`와 `KeyShiftV`를 추가했다.
- normal mode `V`가 linewise visual selection으로 진입한다.
- linewise visual mode에서 `j/k/G/gg` row motion과 full-line normalized selection을 지원한다.
- linewise `d/y`가 unnamed linewise register와 cursor landing을 보장한다.
- linewise delete는 undo 가능하며, 전체 줄 삭제 시 빈 fallback line을 남긴다.
- TUI input mapping과 render selection 표시가 linewise kind를 처리한다.

## 검증 결과

- passed: `go test ./internal/vimengine/...`
- passed: `go test ./internal/runtime/...`
- passed: `go test ./internal/tuiadapter/...`
- passed: `go test ./internal/playable/...`
- passed: `go test ./...`
