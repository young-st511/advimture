# ExecPlan: Single-char edit engine

Slice-ID: VIM-013
Created: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/command-catalog.md
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/vim-013-single-char-edit-engine.md
- docs/exec-plans/completed/vim-013-single-char-edit-engine.md
- internal/vimengine/engine.go
- internal/vimengine/engine_test.go
- internal/tuiadapter/adapter.go
- internal/tuiadapter/adapter_test.go

## 목표

다음 small edits playpack의 첫 기반으로 Normal mode `x`, `r{char}`를 vimengine에 구현한다.

## 수용 기준

- `x`는 현재 cursor 아래 문자를 삭제한다.
- `x`는 삭제 후 cursor를 유효한 위치로 clamp한다.
- 빈 줄에서 `x`는 buffer를 변경하지 않고 boundary event를 낸다.
- `r`은 pending key 상태에 들어간다.
- `r{char}`는 현재 cursor 아래 문자를 `{char}` 하나로 교체하고 Normal mode를 유지한다.
- `esc`는 pending `r` 상태를 취소한다.
- unsupported `r` sequence는 buffer를 변경하지 않고 pending state를 정리한다.
- TUI adapter는 normal mode에서 `x`, `r`, 그리고 pending replacement용 단일 printable key를 engine에 전달할 수 있다.

## 범위

- 포함: vimengine state transition
- 포함: TUI input mapping
- 포함: unit tests
- 제외: content YAML exercise
- 제외: replay/E2E playpack
- 제외: undo/redo history
- 제외: count prefix

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
- `go test ./...`
- `git diff --check`

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/tuiadapter/...`: pass
- `go test ./internal/playable/...`: pass
- `go test ./internal/runtime/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## 작업 항목

- [x] `x` delete current char 구현
- [x] `r{char}` replace current char 구현
- [x] TUI adapter mapping 보강
- [x] command catalog 상태 갱신
- [x] 검증 명령 실행
- [x] ExecPlan completed 이동
