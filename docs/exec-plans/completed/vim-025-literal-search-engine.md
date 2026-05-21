# VIM-025 — Literal Search Engine

Slice-ID: VIM-025
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/vim-025-literal-search-engine.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- internal/vimengine/
- internal/tuiadapter/
- internal/runtime/
- internal/playable/

## 목표

SEARCH-GAP-001에서 결정한 범위대로 `/`, `n`, `N` literal search state와 cursor movement를 구현한다.

## 결과

- `ModeSearch`를 추가했다.
- `/query enter` literal forward search를 구현했다.
- `n`, `N` last search repeat movement를 구현했다.
- search는 줄 경계를 넘고 wrap-around한다.
- TUI adapter는 search mode 입력과 `/`, `n`, `N` normal mode key를 전달한다.
- playable은 search prompt `/query`와 search action panel을 표시한다.

## 수용 기준

- [x] `/query enter`는 현재 cursor 이후의 다음 literal match로 이동한다.
- [x] `n`은 마지막 search 방향으로 다음 match로 이동한다.
- [x] `N`은 마지막 search 방향의 반대로 이전 match로 이동한다.
- [x] search 입력 중 `esc`는 검색을 취소하고 Normal mode로 돌아간다.
- [x] `?`는 search key로 우회하지 않고 기존 hint key로 남는다.

## 검증 결과

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/runtime/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `git diff --check`

## 후속

PLAYPACK-008에서 search-basic tutorial content와 E2E를 연결한다.
