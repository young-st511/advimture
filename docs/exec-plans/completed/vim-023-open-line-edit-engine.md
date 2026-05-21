# VIM-023 — Open Line Edit Engine

Slice-ID: VIM-023
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-023-open-line-edit-engine.md
- docs/exec-plans/completed/vim-023-open-line-edit-engine.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- internal/vimengine/
- internal/tuiadapter/
- internal/runtime/

## 목표

`o`, `O` open-line edit을 engine/runtime 수준에서 구현한다. 플레이어는 현재 줄 아래/위에 빈 줄을 만들고 즉시 Insert mode로 들어갈 수 있어야 한다.

## 범위

- 포함:
  - `o`: 현재 줄 아래 빈 줄 삽입 후 Insert mode 진입
  - `O`: 현재 줄 위 빈 줄 삽입 후 Insert mode 진입
  - open-line mutation undo snapshot
  - TUI adapter lowercase/uppercase mapping
  - runtime replay smoke
  - command catalog/spec engine support 동기화
- 제외:
  - content YAML
  - playlist/E2E
  - indentation
  - auto-comment
  - count prefix
  - insert-mode Enter
  - dot repeat

## 수용 기준

- [x] Normal mode에서 `o`는 현재 줄 아래에 빈 줄을 만들고 Insert mode로 들어간다.
- [x] Normal mode에서 `O`는 현재 줄 위에 빈 줄을 만들고 Insert mode로 들어간다.
- [x] 새 줄에 입력 후 `esc`로 Normal mode에 돌아올 수 있다.
- [x] `o` 또는 `O` 직후 `esc`, `u`는 삽입한 빈 줄을 제거한다.
- [x] TUI adapter는 `o`와 `O`를 구분해 runtime key로 전달한다.
- [x] runtime session은 `o text esc`, `O text esc` trace를 재생해 목표 상태에 도달한다.

## 검증 결과

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/runtime/...`
- `go test ./...`
- `git diff --check`

## 후속

PLAYPACK-006에서 4~6문항 open-line tutorial content, coverage gate, full playlist E2E를 연결한다.
