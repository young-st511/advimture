# ExecPlan: Undo redo engine

Slice-ID: VIM-015
Created: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/command-catalog.md
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/vim-015-undo-redo-engine.md
- docs/exec-plans/completed/vim-015-undo-redo-engine.md
- internal/vimengine/engine.go
- internal/vimengine/engine_test.go
- internal/tuiadapter/adapter.go
- internal/tuiadapter/adapter_test.go

## 목표

small edits playpack의 실패 회복 감각을 위해 Normal mode `u`, `<C-r>` undo/redo를 구현한다.

## 수용 기준

- buffer mutation 명령은 undo snapshot을 남긴다.
- `u`는 마지막 mutation 이전 buffer/cursor로 되돌린다.
- `<C-r>`는 직전 undo를 다시 적용한다.
- 새 mutation이 발생하면 redo stack은 비워진다.
- undo/redo는 Normal mode로 복귀하고 command/pending state를 정리한다.
- undo할 항목이 없으면 boundary event를 낸다.
- redo할 항목이 없으면 boundary event를 낸다.

## 범위

- 포함: vimengine history state
- 포함: TUI adapter `ctrl+r` mapping
- 포함: unit tests
- 제외: Vim undo block exact semantics
- 제외: content YAML exercise
- 제외: E2E playpack

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
- `go test ./...`
- `git diff --check`

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/tuiadapter/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## 작업 항목

- [x] mutation history snapshot 구현
- [x] `u`, `<C-r>` 구현
- [x] redo invalidation 구현
- [x] command catalog 상태 갱신
- [x] 검증 명령 실행
- [x] ExecPlan completed 이동
