# REPEAT-GAP-001 — Repeat Last Change Transaction

Slice-ID: REPEAT-GAP-001
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/repeat-gap-001-last-change-transaction.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- docs/gameplay/vim-curriculum-map.md

## 목표

`.` repeat-last-change 구현 전에 “마지막 변경”의 저장 단위, 첫 구현 subset, 제외 항목을 고정한다.

## 결정

첫 구현은 `State + Key -> State + Events` 계약을 유지하면서 Vim engine 안에 replayable change transaction을 기록한다. transaction은 Normal mode에서 시작한 변경 입력의 key sequence이며, `.`는 그 sequence를 재생한다.

### 포함할 transaction

- 단일 Normal mode 변경:
  - `x`
  - `r<char>`
- Insert transaction:
  - `i ... esc`
  - `a ... esc`
  - `A ... esc`
- Change transaction:
  - `cw ... esc`
  - `c$ ... esc`
  - `cc ... esc`
  - `ciw ... esc`
- Open-line transaction:
  - `o ... esc`
  - `O ... esc`

### 첫 구현에서 제외

- `d` delete 계열 repeat
- `y` yank 계열과 `p/P` put repeat
- Ex command repeat
- search repeat
- macro/register/count prefix
- multi-line Insert mode Enter
- Vim exact undo block semantics

## Transaction 규칙

- transaction은 Normal mode에서 시작한다.
- Insert mode에 들어가는 transaction은 `esc`로 Normal mode에 돌아올 때 commit한다.
- mutation이 없는 빈 transaction은 기록하지 않는다.
- `.` 재생 중 발생한 key는 새로운 transaction으로 기록하지 않는다.
- `.` 자체는 last change에 포함하지 않는다.
- last change가 없을 때 `.`는 boundary event로 끝난다.
- `u`와 `ctrl+r`는 last change를 갱신하지 않는다.
- 첫 구현은 pedagogical undo를 유지한다. `.`가 만든 반복 변경은 기존 undo stack 위에 쌓이며, exact Vim undo block 통합은 후속 hardening으로 미룬다.

## VIM-024 구현 후보

- `vimengine.State`에 last change sequence와 recording state를 추가한다.
- engine 내부에 change recorder를 두되 runtime/content/progress에는 노출하지 않는다.
- `KeyDot = "."`를 추가한다.
- `Apply`가 mutation event를 반환하는 기존 변경 command에서 recorder를 갱신한다.
- Insert transaction은 Insert mode 진입 key부터 `esc`까지의 printable keys를 sequence로 보관한다.
- dot replay는 `ApplyKeys`와 유사하게 sequence를 재생하되 recorder를 비활성화한다.

## PLAYPACK-007 후보

- 4~6문항 efficiency tutorial로 구성한다.
- 초반 문항은 `A ... esc .`로 같은 suffix를 여러 줄에 붙이는 감각을 준다.
- 중반 문항은 `cw ... esc .` 또는 `ciw ... esc .`로 여러 값 교체를 반복한다.
- constraints는 수동 재입력 우회를 금지하고 `.` required key를 요구한다.

## 수용 기준

- [x] `repeat-last-change` command cluster의 첫 구현 subset이 문서화된다.
- [x] last-change transaction commit/cancel/replay 규칙이 문서화된다.
- [x] VIM-024와 PLAYPACK-007의 구현/콘텐츠 경계가 분리된다.
- [x] delete/yank/put/search/macro/register/count prefix는 첫 구현에서 제외한다.
- [x] progress 저장 포맷 변경이 필요 없음을 확인한다.

## 검증 결과

- `git diff --check`

## 후속

VIM-024에서 `.` 엔진 구현을 시작한다.
