# ENGINE-SPLIT-001 — Vimengine Selection/Range Split

Slice-ID: ENGINE-SPLIT-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: refactor-only
Allowed-Paths:
- docs/exec-plans/active/engine-split-001-vimengine-selection-range.md
- docs/exec-plans/completed/engine-split-001-vimengine-selection-range.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/vimengine/engine.go
- internal/vimengine/selection.go
- internal/vimengine/visual.go

## 목표

`internal/vimengine/engine.go`에 누적된 visual selection/range/operator helper를 behavior 변화 없이 별도 파일로 분리한다.

## 수용 기준

- `SelectionKind`, `Selection`, selection 정규화 helper를 `selection.go`로 분리한다.
- visual mode 진입, visual key dispatch, charwise visual `d/y` helper를 `visual.go`로 분리한다.
- public API와 engine behavior는 바꾸지 않는다.
- 새 visual 기능을 추가하지 않는다.

## 결과

- selection type과 normalization/clamp helper를 `internal/vimengine/selection.go`로 이동했다.
- visual mode dispatch와 charwise visual `d/y` helper를 `internal/vimengine/visual.go`로 이동했다.
- `engine.go`의 behavior는 변경하지 않았다.

## 제외한 것

- `V`, visual block, multi-line visual operator 구현
- undo/register semantics 변경
- runtime/content/TUI behavior 변경
- progress 저장 포맷 변경

## 검증

- `go test ./internal/vimengine/...`
- `go test ./...`
- `git diff --check`
