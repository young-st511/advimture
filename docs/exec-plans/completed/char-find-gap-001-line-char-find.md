# CHAR-FIND-GAP-001 — Line Character Find Scope

Slice-ID: CHAR-FIND-GAP-001
Created: 2026-05-26
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md
- docs/exec-plans/active/char-find-gap-001-line-char-find.md
- content/command_clusters/char-find-line.yaml

## 목표

한 줄 안에서 목표 문자까지 빠르게 이동하거나 operator 범위를 잡는 `f/t` command cluster의 첫 구현 범위와 제외 항목을 고정한다.

## 결정

첫 구현은 forward same-line char find만 다룬다. `f{char}`는 현재 cursor 오른쪽의 다음 target char로 이동하고, `t{char}`는 target char 바로 앞까지 이동한다. operator와 결합할 때 `df{char}`는 target char까지 포함해 삭제하고, `dt{char}`는 target char 직전까지만 삭제한다. `cf{char}`, `ct{char}`도 같은 범위를 삭제한 뒤 Insert mode에 들어간다.

## 포함

- normal mode `f{char}`: 같은 줄 오른쪽의 다음 target char로 이동
- normal mode `t{char}`: 같은 줄 오른쪽의 다음 target char 직전으로 이동
- `df{char}`, `dt{char}`: 같은 줄 범위 삭제
- `cf{char}`, `ct{char}`: 같은 줄 범위 변경 후 Insert mode 진입
- target char는 single rune literal로 처리
- not found, empty line, target이 유효 범위를 만들지 않는 경우 boundary event
- undo/redo는 기존 range mutation 계약을 따른다

## 제외

- backward `F`, `T`
- repeat find `;`, `,`
- count prefix
- cross-line search
- visual mode `f/t`
- yank 결합 `yf/yt`
- search highlight 또는 find history
- exact Vim option semantics

## VIM-030 구현 후보

- `KeyF`, `KeyT`를 추가한다.
- normal mode에서 `f/t`는 target char pending state로 들어간다.
- operator pending 상태에서 `f/t`를 받으면 operator+motion+target char pending state로 들어간다.
- char-find helper는 현재 줄과 cursor column만 읽는 순수 함수로 둔다.
- `cf/ct`는 insert recording이 `c`, motion, target char, inserted text, `esc`를 포함하도록 한다.

## PLAYPACK-012 후보

- 4~6문항 tutorial.
- `f=`, `t,`, `df,`, `ct"`를 각각 최소 1문항 이상 다룬다.
- `F/T/;/,`는 allowed/required key에 포함하지 않는다.
- 문항은 config delimiter, comma-separated route, quoted value 같은 한 줄 복구 상황으로 구성한다.

## 수용 기준

- [x] `f/t` scope가 forward same-line literal char find로 고정된다.
- [x] operator 결합 범위가 `df/dt/cf/ct`로 고정된다.
- [x] `F/T/;/,/count/cross-line/visual/yank`는 첫 구현에서 제외된다.
- [x] VIM-030과 PLAYPACK-012의 경계가 분리된다.

## 검증 계획

- `git diff --check`

## 결과

- `char-find-line` command cluster를 approved/planned로 추가했다.
- VIM-030은 engine/runtime/tuiadapter 구현, PLAYPACK-012는 content/E2E 구현으로 분리한다.
- 첫 scope는 forward same-line `f/t`와 `df/dt/cf/ct`로 잠그고, reverse/repeat/count/visual/yank/cross-line은 후속 hardening으로 남겼다.

## 후속

VIM-030에서 char-find engine/runtime/tuiadapter를 구현한다.
