# ExecPlan: Insert mode text engine

Slice-ID: VIM-014
Created: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/command-catalog.md
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/vim-014-insert-mode-text-engine.md
- docs/exec-plans/completed/vim-014-insert-mode-text-engine.md
- internal/vimengine/engine.go
- internal/vimengine/engine_test.go
- internal/tuiadapter/adapter.go
- internal/tuiadapter/adapter_test.go

## 목표

`insert-mode-entry` playpack 후보를 위해 Normal mode `i`, `a`, `A`와 Insert mode printable text insertion을 구현한다.

## 수용 기준

- `i`는 현재 cursor 위치 앞에 삽입할 수 있는 Insert mode로 진입한다.
- `a`는 현재 cursor 위치 뒤에 삽입할 수 있는 Insert mode로 진입한다.
- `A`는 현재 줄 끝에 삽입할 수 있는 Insert mode로 진입한다.
- Insert mode의 단일 printable key는 현재 insert cursor 위치에 들어가고 cursor를 다음 insert 위치로 이동한다.
- `esc`는 Insert mode에서 Normal mode로 복귀하고 cursor를 유효한 Normal mode 위치로 clamp한다.
- Insert mode에서 `q`는 app quit이 아니라 텍스트 입력으로 전달된다.
- newline, backspace, 복잡한 insert editing은 이번 루프에서 제외한다.

## 범위

- 포함: vimengine insert transition
- 포함: TUI adapter insert-mode mapping
- 포함: unit tests
- 제외: content YAML exercise
- 제외: E2E playpack
- 제외: backspace/newline/count prefix

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

- [x] `i`, `a`, `A` 진입 구현
- [x] insert printable mutation 구현
- [x] TUI adapter insert mode mapping 보강
- [x] command catalog 상태 갱신
- [x] 검증 명령 실행
- [x] ExecPlan completed 이동
