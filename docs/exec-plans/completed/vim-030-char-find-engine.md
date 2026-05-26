# VIM-030 — Char Find Engine

Slice-ID: VIM-030
Created: 2026-05-26
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/spec.md
- docs/exec-plans/active/vim-030-char-find-engine.md
- docs/exec-plans/completed/char-find-gap-001-line-char-find.md
- internal/vimengine/
- internal/tuiadapter/
- internal/runtime/

## 목표

CHAR-FIND-GAP-001에서 고정한 `char-find-line` 첫 scope를 engine 레벨에서 구현한다.

## 포함

- normal mode `f{char}`, `t{char}`
- operator 결합 `df{char}`, `dt{char}`, `cf{char}`, `ct{char}`
- same-line forward literal rune find
- boundary/not found handling
- undo/redo for delete/change range mutation
- tuiadapter key mapping
- runtime replay smoke

## 제외

- `F/T`
- `;`, `,`
- count prefix
- visual mode
- `yf/yt`
- cross-line search
- content YAML/playable E2E

## 수용 기준

- `f{char}`는 현재 cursor 오른쪽의 target char로 이동한다.
- `t{char}`는 target char 직전으로 이동한다.
- `df{char}`는 target char를 포함해 삭제한다.
- `dt{char}`는 target char 직전까지만 삭제한다.
- `cf{char}`/`ct{char}`는 같은 범위를 삭제하고 Insert mode로 진입한다.
- target이 없거나 유효 범위가 없으면 boundary/unsupported event로 state를 보존한다.
- `go test ./internal/vimengine ./internal/tuiadapter ./internal/runtime`과 `go test ./...`를 통과한다.

## Step 1: Red Tests

- [x] vimengine normal `f/t` tests
- [x] vimengine `df/dt/cf/ct` tests
- [x] tuiadapter/runtime replay tests

## Step 2: Green Implementation

- [x] key constants and pending states
- [x] char-find helper
- [x] normal motion and operator mutation integration
- [x] change recording support

## Step 3: Verification

- [x] focused tests
- [x] `go test ./...`
- [x] `git diff --check`

## 결과

- `f/t` normal motion과 `df/dt/cf/ct` operator 결합을 구현했다.
- `cf/ct`는 insert recording을 통해 `c`, motion, target char, inserted text, `esc`를 last change로 남긴다.
- `F/T`, `;`, `,`, count, visual, yank 결합, cross-line search는 구현하지 않았다.
