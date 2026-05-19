# VIM-019 — Yank Register Foundation

Slice-ID: VIM-019
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-019-yank-register-foundation.md
- docs/exec-plans/completed/vim-019-yank-register-foundation.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- internal/vimengine/
- internal/tuiadapter/

## 목표

Unnamed register와 `y` operator pending을 추가하고, buffer를 변경하지 않는 yank 명령 `yw`, `y$`, `yy`를 구현한다. 붙여넣기(`p`, `P`)는 VIM-020에서 다룬다.

## 범위

- 포함:
  - `vimengine.State`에 unnamed register 추가
  - register의 charwise/linewise 구분
  - `KeyY` 상수와 `y` pending mode
  - `yw`, `y$`, `yy` yank semantics
  - yank는 buffer/cursor/mode를 변경하지 않고 `EventYanked` 반환
  - `tuiadapter` normal mode `y` mapping
- 제외:
  - `p`, `P`
  - named register, system clipboard
  - count prefix
  - text object
  - content/E2E 연결

## 구현 결과

- `Register` 타입을 추가해 charwise `Text`와 linewise `Lines`를 구분했다.
- `State.Register`를 `copyState`에서 deep-copy한다.
- `KeyY`와 `EventYanked`를 추가했다.
- `yw`, `y$`, `yy`를 pending operator로 구현했다.
- `tuiadapter` normal mode `y` mapping을 추가했다.

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/tuiadapter/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## E2E Evidence

콘텐츠 연결 전 엔진 루프이므로 E2E는 추가하지 않았다. `PLAYPACK-004`에서 `y/p` content와 함께 추가한다.

## 의사결정 로그

- 2026-05-19: yank는 buffer mutation이 아니므로 undo stack에 snapshot을 추가하지 않는다.
- 2026-05-19: `yw`, `y$`는 charwise register, `yy`는 linewise register로 저장한다.
