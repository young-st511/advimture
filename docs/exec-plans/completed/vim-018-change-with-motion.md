# VIM-018 — Change With Motion

Slice-ID: VIM-018
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-018-change-with-motion.md
- docs/exec-plans/completed/vim-018-change-with-motion.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- internal/vimengine/
- internal/runtime/

## 목표

`c + motion`을 삭제와 Insert mode 진입을 결합한 operator grammar로 구현한다. 플레이어는 잘못된 단어, 줄 끝, 현재 줄 전체를 바로 새 텍스트로 교체할 수 있어야 한다.

## 범위

- 포함:
  - `cw`: 현재 cursor부터 다음 단어 시작 전까지 삭제하고 Insert mode 진입
  - `c$`: 현재 cursor부터 현재 줄 끝까지 삭제하고 Insert mode 진입
  - `cc`: 현재 줄을 빈 줄로 만들고 Insert mode 진입
  - `esc`로 Normal mode 복귀 후 cursor clamp
  - undo/redo로 change 결과 복원
  - runtime replay smoke로 `cw`, printable 입력, `esc` trace 검증
- 제외:
  - count prefix, text object, register/yank semantics
  - multi-line change
  - YAML content와 TUI E2E 연결

## 수용 기준

- `c` 다음 `w`는 `dw`와 같은 범위를 삭제한 뒤 Insert mode로 진입한다.
- `c` 다음 `$`는 `d$`와 같은 범위를 삭제한 뒤 Insert mode로 진입한다.
- `c` 다음 `c`는 현재 줄을 빈 줄로 만들고 Insert mode로 진입한다.
- change 후 printable 입력은 삭제 위치에 삽입된다.
- `esc` 후 mode는 Normal이고 cursor는 유효 범위로 clamp된다.
- change 동작은 undo/redo로 되돌릴 수 있다.
- `d + key` 지원 범위와 기존 `g`, `r` pending 동작은 유지된다.
- runtime session은 `cw`, 텍스트 입력, `esc` trace를 보존하고 target buffer/mode를 판정할 수 있다.

## 구현 결과

- `changeWithMotion` dispatcher를 추가해 `cw`, `c$`, `cc`를 지원했다.
- `cw`와 `c$`는 delete range와 같은 pedagogical 범위를 삭제한 뒤 Insert mode로 진입한다.
- `cc`는 현재 줄을 빈 줄로 만든 뒤 Insert mode로 진입한다.
- runtime smoke에서 `cw`, `omega ` 입력, `esc`까지의 trace와 target buffer/mode를 검증했다.

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/runtime/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## E2E Evidence

콘텐츠 연결 전 엔진 루프이므로 E2E는 추가하지 않았다. `PLAYPACK-003`에서 실제 operator grammar adventure intro content와 함께 추가한다.

## 의사결정 로그

- 2026-05-19: `cc`는 줄을 삭제하지 않고 현재 줄을 빈 줄로 만든 뒤 Insert mode로 진입하는 교체 semantics로 구현한다.
- 2026-05-19: pedagogical `cw`는 `dw`와 같은 범위를 사용하므로 뒤 공백까지 삭제한다. 자연스러운 단어 교체 문항은 새 단어 뒤 공백 입력을 optimal trace에 포함한다.
