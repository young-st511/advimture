# OPEN-LINE-001 — Open Line Edit Gap

Slice-ID: OPEN-LINE-001
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/open-line-001-open-line-edit-gap.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md

## 목표

`o/O` open-line edit의 구현 범위와 제외 항목을 고정한다. 플레이어는 현재 줄 주변에 새 줄을 열고 즉시 입력하는 실무 편집 흐름을 배운다.

## 결정

- `open-line-edit`을 approved + planned command cluster로 승격한다.
- 첫 구현은 `o`, `O`를 모두 포함한다.
- `o`는 현재 줄 아래에 빈 줄을 삽입하고 Insert mode로 진입한다.
- `O`는 현재 줄 위에 빈 줄을 삽입하고 Insert mode로 진입한다.
- `esc`로 Normal mode 복귀하는 기존 insert-mode 계약을 재사용한다.
- 삭제/변경 계열처럼 open-line mutation도 undo snapshot을 남긴다.
- content schema 변경은 필요 없다. 기존 `optimal_keys`, `constraints`, `e2e_assertions`로 표현한다.

## 제외

- indentation
- auto-comment
- count prefix
- insert-mode Enter
- backspace
- dot repeat 연계
- progress schema 변경

## 후속 루프

| 루프 | 목표 |
|------|------|
| VIM-023 | `o`, `O` engine/tuiadapter/runtime 구현 |
| PLAYPACK-006 | 4~6문항 open-line tutorial content/E2E 구현 |
| DEBRIEF-001 | 저장 포맷 변경 없는 success/playlist debrief |

## VIM-023 수용 기준 초안

- Normal mode에서 `o`는 현재 줄 아래에 빈 줄을 삽입하고 Insert mode로 진입한다.
- Normal mode에서 `O`는 현재 줄 위에 빈 줄을 삽입하고 Insert mode로 진입한다.
- 입력 후 `esc`는 Normal mode로 돌아온다.
- `u`는 open-line 삽입과 입력 결과를 되돌릴 수 있다.
- TUI adapter는 lowercase `o`와 uppercase `O`를 구분해 runtime key로 전달한다.

## PLAYPACK-006 수용 기준 초안

- `open-line-edit`은 `o`, `O` coverage를 모두 만족한다.
- tutorial playlist는 8문항 이하로 유지한다.
- 각 exercise는 `replay_status: pass`와 E2E assertion을 가진다.
- full E2E는 `o/O` 방향, final buffer, cursor, mode, progress를 검증한다.

## 검증

- `git diff --check`
