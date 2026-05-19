# VIM-022 — Inner Word Text Object Semantics

Slice-ID: VIM-022
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-022-inner-word-text-object-semantics.md
- docs/exec-plans/completed/vim-022-inner-word-text-object-semantics.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/command-catalog.md
- internal/vimengine/
- internal/runtime/

## 목표

`diw`, `ciw`, `yiw`를 엔진 동작으로 구현한다. 플레이어는 단어 시작으로 이동하지 않고도 현재 단어 전체를 삭제, 변경, 복사할 수 있어야 한다.

## 결과

- `diw`는 현재 cursor가 놓인 keyword/symbol run 전체를 삭제한다.
- `ciw`는 현재 inner word를 삭제한 뒤 Insert mode로 진입한다.
- `yiw`는 현재 inner word를 unnamed charwise register에 저장하고 buffer/cursor를 유지한다.
- 공백 위에서 `iw`를 실행하면 buffer를 변경하지 않고 boundary event로 처리한다.
- 삭제/변경은 undo stack에 기록된다.
- runtime replay smoke에서 `diw`, `ciw`, `yiw + p` 흐름을 검증했다.

## 제외

- `aw`
- quote/pair text object
- count prefix
- clipboard/named register

## 검증

- `go test ./internal/vimengine/...`
- `go test ./internal/runtime/...`
