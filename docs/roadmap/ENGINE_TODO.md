# Engine / Module Todo

> Advimture는 엔진, 미들웨어, 어댑터를 분리해서 만든다. 각 항목은 독립 ExecPlan으로 열고, 테스트와 커밋까지 닫은 뒤 다음 항목으로 넘어간다.

## 진행 원칙

- 한 루프는 하나의 엔진 또는 모듈 계약만 다룬다.
- 새 Vim command 추가와 runtime/UI 연결을 같은 루프에서 섞지 않는다.
- 기존 구현이 새 경계를 방해하면 먼저 legacy/archive 격리를 검토한다.
- TUI 연결 전까지는 unit test와 model-level test를 우선한다.
- TUI로 연결되는 순간 E2E assertion이 부족하면 구현을 멈추고 E2E를 보강한다.

## 순서

| 순서 | ID | 엔진/모듈 | 상태 | 핵심 책임 |
|------|----|-----------|------|-----------|
| 1 | VIM-002 | Vim Engine Foundation | completed | `State + Key -> State + Events`, `h/j/k/l` |
| 2 | VIM-003 | Vim Engine Contract Hardening | completed | 상태 주입, key trace replay, copy/normalize 경계 |
| 3 | VIM-004 | Exercise Runtime Foundation | completed | 문항 세션, 목표 판정, key trace, retry, hint |
| 4 | VIM-005 | Content Schema Foundation | completed | command/exercise/scenario 데이터를 runtime spec으로 변환 |
| 5 | VIM-006 | Grader / Scoring Engine | completed | 정답, 최적 키 대비 평가, 실수/힌트 평가 |
| 6 | VIM-007 | Scenario Runtime | completed | 검증된 exercise를 어드벤처 사건으로 감싸기 |
| 7 | VIM-008 | TUI Adapter | completed | Bubble Tea input/output과 runtime 연결 |
| 8 | VIM-009 | Progress Adapter | completed | runtime 결과를 로컬 저장 포맷으로 변환 |
| 9 | VIM-010 | Neovim Oracle Runner | completed | `exact` tier command를 Neovim 결과와 비교 |
| 10 | VIM-011 | Legacy Archive Pass | completed | 기존 `internal/editor/game/app/ui/data` 유지/격리 결정 |

## 다음 루프 후보

현재 진행 중: Utility Commands and Long-Run Platform.

다음 후보는 2026-05-21 토론 결과에 따라 아래 순서로 검토한다.

| 우선순위 | 후보 | 필요한 엔진 계약 | 비고 |
|----------|------|------------------|------|
| 1 | `open-line-edit` | `o`, `O`, newline insertion, insert mode entry, undo snapshot | 현재 next playable candidate |
| 2 | `repeat-last-change` | last-change transaction, replayable mutation, undo/redo 상호작용 | efficiency run의 핵심 |
| 3 | `search-basic` | literal search state, `/` query input, `n/N` match navigation | `?` hint 충돌과 regex는 보류 |
| 4 | platform debrief | 기존 progress 기반 best record/read-only summary | 저장 포맷 변경 없이 게임성 강화 |
| 5 | long-run review | mastery/spaced review/daily run progress 후보 | RFC와 사용자 승인 필요 |

진행 원칙:

- 새 Vim command를 늘릴 때는 `command catalog -> vimengine -> oracle comparison -> exercise` 순서로 진행한다.
- TUI로 연결되는 순간 E2E assertion이 부족하면 구현을 멈추고 E2E를 보강한다.

## ENGINE-GAP-001 결정

다음 구현 루프는 `VIM-013: single-char-edit engine`으로 한다.

| 루프 | 범위 | 필수 테스트 | E2E |
|------|------|-------------|-----|
| VIM-013 | `x`, `r`, buffer mutation, cursor clamp | `internal/vimengine` unit, runtime replay smoke | completed: engine/tuiadapter unit까지 연결. content 승격 시 E2E 추가 |
| VIM-014 | `i`, `a`, `A`, printable insertion, `esc` 복귀 | vimengine/tuiadapter/runtime unit | completed: insert-mode printable input까지 연결. content 승격 시 E2E 추가 |
| VIM-015 | `u`, `ctrl+r`, mutation history | vimengine unit, scenario scoring regression | completed: mutation snapshot undo/redo 구현. content 승격 시 E2E 추가 |

`PLAYPACK-002`는 VIM-013~015 중 실제 implemented cluster만 playable로 승격한다. engine support가 planned인 cluster는 YAML에 남길 수 있지만 playable path에는 연결하지 않는다.

## OPERATOR-GAP-001 결정

첫 operator grammar 구현은 `d/c + motion`으로 제한한다. text object, yank/put, count prefix는 다음 milestone로 미룬다.

| 루프 | 범위 | 필수 테스트 | E2E |
|------|------|-------------|-----|
| VIM-016 | operator pending mode, `d/c` key mapping, unsupported combo | completed: `internal/vimengine`, `internal/tuiadapter` unit | content 연결 전 E2E 없음 |
| VIM-017 | `dw`, `d$`, `dd` 삭제 semantics | completed: vimengine unit, runtime replay smoke | PLAYPACK-003에서 추가 |
| VIM-018 | `cw`, `c$`, `cc` change semantics | completed: vimengine unit, runtime replay smoke | PLAYPACK-003에서 추가 |

`PLAYPACK-003`은 VIM-016~018이 completed된 뒤 YAML content와 E2E를 연결한다.

## YANK-TEXT-001 결정

다음 구현 루프는 `yank-put-basic`을 먼저 다룬다. text object는 register와 put semantics가 안정된 뒤 별도 gap planning으로 넘긴다.

| 루프 | 범위 | 필수 테스트 | E2E |
|------|------|-------------|-----|
| VIM-019 | `y` pending, unnamed register, `yw`, `y$`, `yy` | completed: `internal/vimengine`, `internal/tuiadapter` unit | content 연결 전 E2E 없음 |
| VIM-020 | `p`, `P`, charwise/linewise put, undo/redo | completed: vimengine unit, runtime replay smoke | PLAYPACK-004에서 추가 |
| PLAYPACK-004 | `yank-put-basic` tutorial content | completed: content replay, coverage, playable model | completed: full playlist E2E |
| TEXT-OBJECT-001 | `iw` 기반 text object gap planning | docs + scope review | 없음 |

첫 `y/p` 구현은 unnamed register만 다루며 named register, clipboard, count prefix, visual selection은 제외한다.

## TEXT-OBJECT-001 결정

첫 text object 구현은 `iw`만 다룬다. operator grammar는 기존 `d/c/y + motion`에서 `d/c/y + i + w` 3-key sequence로 확장한다.

| 루프 | 범위 | 필수 테스트 | E2E |
|------|------|-------------|-----|
| VIM-021 | `operator -> i -> w` pending sequence, inner-word selection helper | completed: `internal/vimengine`, `internal/tuiadapter` unit | content 연결 전 E2E 없음 |
| VIM-022 | `diw`, `ciw`, `yiw` semantics, undo/register/runtime replay | completed: vimengine unit, runtime replay smoke | PLAYPACK-005에서 추가 |
| PLAYPACK-005 | `text-object-inner-word` tutorial content | completed: content replay, coverage, playable model | completed: full playlist E2E |
| E2E-PLAYPACK-005 | text object playpack QA hardening | completed: E2E evidence and assertions | completed: full playlist + forbidden route |

quote/pair text object(`i"`, `i'`, `i(`, `i{`), around object(`aw`), visual selection, count prefix는 후속 milestone으로 넘긴다.

## OPEN-LINE-001 이후 결정 예정

첫 utility command 구현은 `open-line-edit`로 한다. `o/O`는 현재 줄 주변에 새 줄을 열고 즉시 Insert mode로 진입하는 실무 편집 흐름이다.

| 루프 | 범위 | 필수 테스트 | E2E |
|------|------|-------------|-----|
| OPEN-LINE-001 | `o/O` scope, 제외 항목, VIM-023/PLAYPACK-006 분리 계획 | completed: docs + scope review | 없음 |
| VIM-023 | `o`, `O`, newline insertion, insert mode entry, undo snapshot | completed: `internal/vimengine`, `internal/tuiadapter`, `internal/runtime` unit | content 연결 전 E2E 없음 |
| PLAYPACK-006 | `open-line-edit` tutorial content | completed: content replay, coverage, playable model | completed: full playlist E2E + forbidden route |
| DEBRIEF-001 | 저장 변경 없는 success/playlist completion debrief | completed: playable model tests | completed: focused E2E |

첫 `o/O` 구현에서는 indentation, auto-comment, count prefix, insert-mode Enter, `.` repeat 연계를 제외한다.

## REPEAT / SEARCH 장기 후보

`repeat-last-change`는 REPEAT-GAP-001에서 x, r<char>, insert/change/open-line transaction을 첫 구현 범위로 고정했고, VIM-024/PLAYPACK-007에서 playable tutorial까지 연결했다. `search-basic`은 SEARCH-GAP-001에서 literal `/`, `n`, `N`만 첫 scope로 고정했고, VIM-025에서 engine support를 구현했다. `?`는 현재 hint key와 충돌하므로 제외한다.
