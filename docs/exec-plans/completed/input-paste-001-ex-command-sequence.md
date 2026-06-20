# ExecPlan: Pasted Ex Command Input

Slice-ID: INPUT-PASTE-001
Created: 2026-06-20
Status: completed
Scope-Mode: strict
Allowed-Paths:
- `docs/exec-plans/completed/input-paste-001-ex-command-sequence.md`
- `internal/tuiadapter/**`
- `internal/playable/**`
- `test/e2e/**`

## 영향 도메인

- Gameplay: playable TUI 입력을 Vim key sequence로 전달하는 어댑터 계약을 다룬다.
- Verification: tutorial `:%s` 진행 불가 재현을 unit/E2E 수준에서 고정한다.

## 공통 제한

- `go.mod`, `go.sum` 변경 없음.
- `internal/progress/` 저장 포맷 변경 없음.
- content schema 변경 없음.
- 새 Vim engine capability 추가 없음.
- 새 dependency 추가 없음.
- 기존 exercise target state, optimal keys, constraints 변경 없음.

## 배경

Tutorial `vim-ex-command-substitute-002`는 한 글자씩 `:%s/TODO/DONE/g`를 입력하면 통과하지만, 터미널이 여러 rune을 하나의 key message로 전달하거나 사용자가 명령을 붙여넣는 경우 `tuiadapter`가 긴 문자열을 `ignored`로 처리해 진행이 막힌다. 기존 E2E는 한 글자씩만 보내 이 사각지대를 잡지 못했다.

## Todo 1: Reproduce Pasted Command Regression

- 상태: completed
- 목표: `:%s/TODO/DONE/g`가 한 번에 들어와도 tutorial exercise가 성공해야 한다는 기준을 테스트로 고정한다.
- 충족 기준:
  - command mode에서 `%s/TODO/DONE/g`가 key sequence로 매핑된다.
  - normal mode에서 `:%s/TODO/DONE/g`가 key sequence로 매핑된다.
  - playable model에서 붙여넣은 전체 파일 substitute 명령이 성공 상태와 DONE buffer를 만든다.

## Todo 2: Apply Key Sequence Through Playable

- 상태: completed
- 목표: 입력 어댑터와 playable update loop가 여러 key를 순서대로 처리한다.
- 충족 기준:
  - single-key 동작과 기존 action 계약은 유지된다.
  - normal mode는 `:` 또는 `/`로 시작하는 붙여넣기만 sequence로 해석한다.
  - command/search/insert mode의 printable pasted text는 sequence로 해석한다.
  - focused Go tests가 통과한다.

## Todo 3: Close and Verify

- 상태: completed
- 목표: 전체 검증 후 completed plan으로 이동한다.
- 충족 기준:
  - `go test ./internal/tuiadapter ./internal/playable` 통과.
  - `go test ./...` 통과.
  - `git diff --check` 통과.

## 검증

- `go test ./internal/tuiadapter ./internal/playable`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_substitute_whole_file_paste.yaml`
- `go test ./...`
- `git diff --check`
