# VIM-026 — Quote Text Object Engine

Slice-ID: VIM-026
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: engine-only
Allowed-Paths:
- docs/exec-plans/active/vim-026-quote-text-object-engine.md
- docs/exec-plans/completed/vim-026-quote-text-object-engine.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- content/command_clusters/text-object-quote-pair.yaml
- internal/vimengine/
- internal/runtime/

## 목표

`di"`, `ci"`, `yi"`를 Vim-like engine과 runtime replay에서 지원한다.

## 범위

- 포함:
  - operator inner text object pending 상태에서 `"` 처리
  - double quote 내부 range 탐색
  - delete/change/yank semantics
  - undo와 last-change recording
  - runtime replay smoke
- 제외:
  - content exercise/scenario/playlist
  - E2E
  - nested pair
  - escaped quote
  - single quote, parenthesis, brace
  - around object
  - count prefix
  - visual selection

## 수용 기준

- completed: `di"`는 같은 줄의 double quote 내부 값만 삭제하고 Normal mode를 유지한다.
- completed: `ci"`는 같은 줄의 double quote 내부 값만 삭제하고 Insert mode로 진입한다.
- completed: `yi"`는 같은 줄의 double quote 내부 값을 unnamed register에 저장하고 buffer를 변경하지 않는다.
- completed: quote 내부 range를 찾지 못하면 buffer/register를 변경하지 않는다.
- completed: `ci"` 변경은 `esc` 이후 `.` 반복 대상으로 기록된다.
- completed: runtime은 `d i "`, `c i "`, `y i "` required key replay를 성공으로 평가할 수 있다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/vimengine/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/runtime/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `git diff --check`
