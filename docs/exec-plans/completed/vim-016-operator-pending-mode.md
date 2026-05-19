# VIM-016 — Operator Pending Mode

Slice-ID: VIM-016
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-016-operator-pending-mode.md
- docs/exec-plans/completed/vim-016-operator-pending-mode.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- internal/vimengine/
- internal/tuiadapter/

## 목표

`d`와 `c`를 일반 단일 키가 아니라 Vim operator grammar의 시작으로 인식하게 만든다. 이번 루프는 실제 삭제/변경 범위를 구현하지 않고, 이후 `dw`, `d$`, `dd`, `cw`, `c$`, `cc`가 안전하게 올라갈 pending contract를 고정한다.

## 범위

- 포함:
  - `internal/vimengine`에 `KeyD`, `KeyC` 상수 추가
  - Normal mode에서 `d`, `c` 입력 시 `PendingKey`를 설정하고 `EventPendingKey` 반환
  - 아직 지원하지 않는 `operator + key` 조합 입력 시 buffer/cursor를 변경하지 않고 pending을 해제하며 `EventUnsupportedKey` 반환
  - `esc`로 operator pending을 취소하는 회귀 테스트
  - `internal/tuiadapter`에서 `d`, `c` normal-mode key mapping을 명시적으로 고정
- 제외:
  - `dw`, `d$`, `dd` 삭제 semantics
  - `cw`, `c$`, `cc` 변경 semantics와 Insert mode 진입
  - count prefix, text object, yank/put
  - YAML content, playable, E2E 연결
  - 저장 포맷 변경

## 수용 기준

- Normal mode에서 `d`는 cursor/buffer를 변경하지 않고 `PendingKey == "d"`와 `EventPendingKey`를 반환한다.
- Normal mode에서 `c`는 cursor/buffer를 변경하지 않고 `PendingKey == "c"`와 `EventPendingKey`를 반환한다.
- `d` 또는 `c` pending 상태에서 아직 구현되지 않은 다음 key를 입력하면 pending이 해제되고 기존 buffer/cursor/mode는 보존된다.
- `esc`는 `d` 또는 `c` pending 상태를 해제하고 Normal mode를 유지한다.
- 기존 `g` pending, `r` pending, insert mode printable 입력 동작은 유지된다.
- TUI adapter는 normal mode의 `d`, `c`를 vimengine key로 매핑하고, insert mode의 `d`, `c`는 printable input으로 보존한다.

## 구현 결과

- `internal/vimengine`에 `KeyD`, `KeyC`를 추가하고 Normal mode pending 전이에 연결했다.
- `applyPendingKey`는 아직 구현되지 않은 `d/c + key` 조합을 unsupported로 처리하며 pending을 해제한다.
- `internal/tuiadapter`는 normal mode `d`, `c`를 명시적으로 vimengine key에 매핑한다.
- insert mode printable 입력 회귀 테스트에 `d`, `c`를 포함했다.

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/tuiadapter/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## E2E Evidence

콘텐츠 연결 전 엔진 루프이므로 E2E는 추가하지 않았다. `PLAYPACK-003`에서 실제 operator grammar content를 연결할 때 `test/e2e/` scenario를 추가한다.

## 의사결정 로그

- 2026-05-19: VIM-016은 operator grammar의 "상태 전이 기반"만 다루고 mutation semantics는 VIM-017/VIM-018로 분리한다.
