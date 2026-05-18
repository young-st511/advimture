# PLAYPACK-002 — Small Edits Tutorial

## 상태

- 상태: completed
- 시작일: 2026-05-19
- 완료일: 2026-05-19
- 사용자 승인:
  - Ex command보다 작은 편집 튜토리얼을 먼저 배치한다.
  - 필요하면 기존 tutorial playlist ID를 더 이상적인 순서로 변경한다.
  - 7문항 구성: `x`, `r`, `i`, `a`, `A`, `u`, `ctrl+r`.
  - 튜토리얼 톤은 첫 투어 느낌을 유지한다.

## 배경

현재 playable tutorial은 이동, 생존, 빠른 이동 뒤 곧바로 Ex substitute로 넘어간다. Vim 학습 순서상 `:s`보다 먼저 글자 하나 삭제/교체, Insert mode 진입, undo/redo를 경험하는 편이 자연스럽다.

## 목표

작은 수정 중심의 7문항 tutorial playpack을 추가하고, playable tutorial 순서를 다음으로 고정한다.

1. `tutorial-0-movement`
2. `tutorial-1-survival`
3. `tutorial-2-fast-navigation`
4. `tutorial-3-small-edits`
5. `tutorial-4-ex-command`

## 범위

- `content/command_clusters/first-five-minutes.yaml`
- `content/exercises/first-five-minutes.yaml`
- `content/scenarios/first-five-minutes.yaml`
- `content/playlists/first-five-minutes.yaml`
- `docs/gameplay/*`
- `docs/roadmap/*`
- `internal/runtime/*`
- `internal/content/*_test.go`
- `cmd/e2e-runner/*`
- `test/e2e/*.yaml`
- `Makefile`

## 비범위

- 저장 포맷 변경
- 새 Go 의존성 추가
- 실제 Vim/Neovim 런타임 연결
- Insert mode의 backspace/newline/multibyte hardening
- Ex command semantics 확장

## 완료 내용

- `single-char-edit`, `insert-mode-entry`, `undo-redo-basic` YAML command cluster를 구현했다.
- 다음 7개 exercise를 approved + implemented + replay pass 상태로 추가했다.
  - `single-char-edit-001`: `x`
  - `single-char-edit-002`: `r`
  - `insert-mode-entry-001`: `i`
  - `insert-mode-entry-002`: `a`
  - `insert-mode-entry-003`: `A`
  - `undo-redo-basic-001`: `u`
  - `undo-redo-basic-002`: `ctrl+r` 중심, 마지막 `h`로 결과 위치 확인
- `tutorial-3-small-edits`를 추가하고 기존 Ex command tutorial을 `tutorial-4-ex-command`로 이동했다.
- E2E runner에 `ctrl+r` 전송을 추가했다.
- required key가 있는 초기 목표 상태 일치 문항이 시작 즉시 성공하지 않도록 runtime guard를 보강했다.

## 검증

- `go test ./cmd/e2e-runner/...`: pass
- `go test ./internal/runtime/...`: pass
- `go test ./internal/content/...`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass
- `git diff --check`: pass
