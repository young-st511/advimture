# TEXT-OBJECT-001 — Inner Word Text Object Gap

Slice-ID: TEXT-OBJECT-001
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md

## 목표

`yank-put-basic` 이후 operator grammar를 motion 대상에서 text object 대상으로 확장한다. 첫 범위는 플레이어가 커서 위치와 무관하게 현재 단어 내부를 대상으로 잡는 `iw`다.

## 결정

- `text-object-inner-word`를 approved + planned로 승격한다.
- 첫 구현 command는 `diw`, `ciw`, `yiw`다.
- 엔진은 `operator -> i -> w` 3-key pending sequence를 지원한다.
- `iw` selection은 현재 커서가 놓인 keyword 또는 symbol run 전체를 잡는다.
- 공백 위에서 `iw`를 실행하는 경우는 unsupported/boundary로 처리한다.

## 제외

- `i"`, `i'`, `i(`, `i{`
- `aw`
- visual selection
- count prefix
- named register / clipboard
- Vim whitespace 세부 semantics exact hardening

## 후속 루프

| 루프 | 목표 |
|------|------|
| VIM-021 | `operator -> i -> w` pending sequence와 inner-word selection 기반 |
| VIM-022 | `diw`, `ciw`, `yiw` semantics |
| PLAYPACK-005 | 6문항 text object inner word tutorial |
| E2E-PLAYPACK-005 | full playlist와 forbidden route 검증 |

## 검증

- 문서 정합성: `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/ENGINE_TODO.md`, `docs/gameplay/command-catalog.md`, `docs/gameplay/vim-curriculum-map.md`
