# VIM-024 — Repeat Last Change Engine

Slice-ID: VIM-024
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-024-repeat-last-change-engine.md
- docs/exec-plans/completed/vim-024-repeat-last-change-engine.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- internal/vimengine/
- internal/tuiadapter/
- internal/runtime/

## 목표

REPEAT-GAP-001에서 결정한 최소 subset으로 `.` repeat-last-change를 구현한다.

## 결과

- `KeyDot = "."`를 추가했다.
- `vimengine.State`에 last-change sequence와 recording state를 추가했다.
- `x`, `r<char>`, insert transaction, change transaction, open-line transaction을 기록한다.
- `.` replay는 recording을 비활성화해 last-change를 `.` 자신으로 덮어쓰지 않는다.
- TUI adapter는 `.`를 runtime key로 전달한다.
- runtime replay smoke를 추가했다.

## 수용 기준

- [x] last change가 없을 때 `.`는 boundary event를 낸다.
- [x] `A text esc`, `i text esc`, `o text esc`, `O text esc` 이후 `.`가 같은 변경을 현재 위치에 반복한다.
- [x] `r<char>` 이후 `.`가 같은 문자 교체를 반복한다.
- [x] `ciw text esc` 이후 `.`가 현재 단어를 같은 값으로 교체한다.
- [x] `.` 재생은 last-change sequence를 `.` 자신으로 덮어쓰지 않는다.
- [x] TUI adapter는 `.`를 runtime key로 전달한다.

## 검증 결과

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/runtime/...`
- `go test ./...`
- `git diff --check`

## 후속

PLAYPACK-007에서 repeat-last-change efficiency tutorial content와 E2E를 연결한다.
