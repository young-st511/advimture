# UI-RENDER-001 — Extract Playable Renderer

Slice-ID: UI-RENDER-001
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: refactor-only
Allowed-Paths:
- docs/exec-plans/active/ui-render-001-playable-renderer.md
- docs/exec-plans/completed/ui-render-001-playable-renderer.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/playable/model.go
- internal/playable/model_test.go
- internal/playableview/

## 목표

`internal/playable`의 Bubble Tea state/update 책임과 TUI string rendering 책임을 분리한다. 이번 slice는 동작 변경 없이 renderer만 새 패키지로 옮긴다.

## 완료 내용

- `internal/playableview` 순수 rendering package를 추가했다.
- `internal/playable.Model.View()`는 screen data와 action line만 구성하고 renderer에 출력을 위임한다.
- cursor, visual selection, action panel rendering 테스트를 `internal/playableview`로 이동했다.
- game state, progress, content schema, E2E state shape는 변경하지 않았다.

## 검증 결과

- `go test ./internal/playableview/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `git diff --check`

## 후속 작업

다음 slice는 `UI-HIERARCHY-001`이다. renderer가 분리됐으므로 현재 exercise 목표를 화면 상위로 올리고 review/daily route를 보조 정보로 낮추는 실제 정보 위계 변경을 진행한다.
