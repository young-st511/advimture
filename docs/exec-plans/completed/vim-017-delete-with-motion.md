# VIM-017 — Delete With Motion

Slice-ID: VIM-017
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-017-delete-with-motion.md
- docs/exec-plans/completed/vim-017-delete-with-motion.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- internal/vimengine/
- internal/runtime/

## 목표

`d + motion`을 통해 플레이어가 단어, 줄 끝, 현재 줄 전체를 한 번에 삭제할 수 있는 엔진 semantics를 만든다. 이번 루프는 `delete-with-motion`만 다루고, change operator는 VIM-018로 분리한다.

## 범위

- 포함:
  - `dw`: 현재 cursor부터 다음 단어 시작 전까지 삭제
  - `d$`: 현재 cursor부터 현재 줄 끝까지 삭제
  - `dd`: 현재 줄 전체 삭제
  - 삭제 동작의 undo snapshot 저장과 redo stack 초기화
  - pending 해제, cursor clamp, mode 유지
  - runtime replay smoke로 `d`, motion key trace가 exercise 경로에서 보존되는지 확인
- 제외:
  - `cw`, `c$`, `cc`
  - count prefix, text object, multi-line `dw`
  - register/yank semantics
  - YAML content와 TUI E2E 연결

## 수용 기준

- `d` 다음 `w`는 현재 cursor부터 다음 단어 시작 전까지 삭제하고 `EventChanged`를 반환한다.
- `d` 다음 `$`는 현재 cursor부터 현재 줄 끝까지 삭제하고 `EventChanged`를 반환한다.
- `d` 다음 `d`는 현재 줄 전체를 삭제하고 다음 줄을 같은 row로 당긴다. 마지막 줄 삭제 시 빈 줄 하나를 남긴다.
- 삭제 후 cursor는 유효한 Normal mode 위치로 clamp되고 `PendingKey`는 비워진다.
- 삭제는 undo/redo로 되돌릴 수 있다.
- `c + key` unsupported 처리와 기존 `g`, `r` pending 동작은 유지된다.
- runtime session은 `dw` key trace를 `["d", "w"]`로 보존하고 target buffer를 판정할 수 있다.

## 구현 결과

- `deleteWithMotion` dispatcher를 추가해 `dw`, `d$`, `dd`만 지원하도록 했다.
- `deleteWordForward`, `deleteToLineEnd`, `deleteCurrentLine`, `deleteLineRange` helper를 추가했다.
- 삭제 동작은 기존 undo/redo snapshot 경로를 사용한다.
- runtime session smoke test로 `dw` trace와 target buffer 성공 판정을 고정했다.

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/runtime/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## E2E Evidence

콘텐츠 연결 전 엔진 루프이므로 E2E는 추가하지 않았다. `PLAYPACK-003`에서 `dw`, `d$`, `dd` playable content와 함께 추가한다.

## 의사결정 로그

- 2026-05-19: 첫 `dw`는 pedagogical semantics로 현재 줄 안의 다음 단어 시작 전까지만 삭제한다.
