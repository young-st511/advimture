# ExecPlan: Small edits engine gap

Slice-ID: ENGINE-GAP-001
Created: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/gameplay/vim-curriculum-map.md
- docs/exec-plans/active/engine-gap-001-small-edits-gap.md
- docs/exec-plans/completed/engine-gap-001-small-edits-gap.md

## 목표

`PLAYPACK-002`를 구현하기 전에 `single-char-edit`, `insert-mode-entry`, `undo-redo-basic`에 필요한 vimengine/runtime gap과 구현 순서를 결정한다.

## 현재 관찰

- `internal/vimengine`은 movement, command-line, substitute는 지원한다.
- Normal mode printable editing key 중 `x`, `r`, `i`, `a`, `A`, `u`, `<C-r>`는 아직 지원하지 않는다.
- `ModeInsert`는 상태 값으로만 존재하고 printable text insertion은 아직 없다.
- engine state에는 undo/redo history가 없다.
- TUI adapter는 normal mode에서 `x`, `i`, `a`, `u` 등을 아직 ActionKey로 매핑하지 않는다.

## 수용 기준

- 작은 수정 playpack에 필요한 engine gap을 우선순위별로 정리한다.
- 각 gap은 구현 대상 module, 필요한 test, E2E 필요 여부를 가진다.
- 다음 구현 루프 하나를 명확히 선택한다.
- `PLAYPACK-002`는 engine 구현이 끝난 cluster만 playable로 승격한다.

## 결정

다음 구현 루프는 `VIM-013: single-char-edit engine`으로 한다.

이유:

- `x`, `r`은 insert text input 모델 없이도 buffer mutation을 검증할 수 있다.
- undo/redo와 insert mode가 의존할 “mutation event” 감각을 가장 작게 만들 수 있다.
- 2문항짜리 playpack 조각으로 빠르게 replay/E2E까지 닫을 수 있다.

## 후속 구현 순서

1. `VIM-013`: `x`, `r` buffer mutation engine
2. `VIM-014`: Insert mode entry and printable text insertion (`i`, `a`, `A`, `esc`)
3. `VIM-015`: Undo/redo stack (`u`, `<C-r>`)
4. `PLAYPACK-002`: small edits tutorial content and E2E

## 검증 계획

- `rg "VIM-013|VIM-014|VIM-015|PLAYPACK-002" docs/roadmap docs/exec-plans docs/gameplay`
- `go test ./...`
- `git diff --check`

## 검증 결과

- `rg "VIM-013|VIM-014|VIM-015|PLAYPACK-002" docs/roadmap docs/exec-plans docs/gameplay`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## 작업 항목

- [x] 현재 engine/tuiadapter/runtime gap을 문서화한다.
- [x] 다음 구현 루프를 선택한다.
- [x] ENGINE_TODO와 MIDTERM/PROGRAM을 갱신한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동한다.
