# SEARCH-GAP-001 — Literal Search Scope

Slice-ID: SEARCH-GAP-001
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/search-gap-001-literal-search.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- docs/gameplay/vim-curriculum-map.md

## 목표

literal `/`, `n`, `N` search의 첫 구현 범위와 `?` hint 충돌 처리 방침을 고정한다.

## 결정

첫 search 구현은 forward literal search만 시작 명령으로 지원한다. `?`는 현재 TUI hint key이므로 backward search 시작 명령으로 사용하지 않는다.

## 포함

- `/query enter`: 현재 cursor 이후의 다음 literal match로 이동한다.
- `n`: 마지막 search 방향으로 다음 match로 이동한다.
- `N`: 마지막 search 방향의 반대로 이전 match로 이동한다.
- 검색은 case-sensitive literal match다.
- 검색은 줄 경계를 넘고, 문서 끝/처음에서 wrap-around한다.
- 성공 목표는 cursor/mode/key trace로 검증한다.

## 제외

- `?query enter`
- regex
- highlight 검증
- search offset
- `*`, `#`
- ignorecase/smartcase 옵션
- search history UI
- search state app_state schema 확장

## VIM-025 구현 후보

- `ModeSearch`를 추가한다.
- `/`는 search 입력 모드에 진입하고, prompt는 `/query`로 표시한다.
- search 입력 중 `enter`는 literal query를 실행한다.
- `esc`는 search 입력을 취소하고 Normal mode로 돌아간다.
- `n/N`은 `LastSearch`를 사용해 이동한다.

## PLAYPACK-008 후보

- 4~6문항 search tutorial.
- `/ERROR enter`, `n`, `N`을 각각 최소 1문항 이상 다룬다.
- `?`는 allowed/required key에 포함하지 않는다.
- 문항은 로그/설정에서 명시된 token을 찾는 상황으로 구성한다.

## 수용 기준

- [x] `/`, `n`, `N` scope가 literal search로 고정된다.
- [x] `?` backward search는 hint 충돌 때문에 첫 구현에서 제외한다.
- [x] regex/highlight/search history는 첫 구현에서 제외한다.
- [x] VIM-025와 PLAYPACK-008의 경계가 분리된다.

## 검증 결과

- `git diff --check`

## 후속

VIM-025에서 literal search engine/runtime을 구현한다.
