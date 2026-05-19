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

현재 진행 중: Yank / Put and Text Object Bridge.

다음 후보는 `ENGINE-GAP-001`에서 아래 순서로 검토한다.

| 우선순위 | 후보 | 필요한 엔진 계약 | 비고 |
|----------|------|------------------|------|
| 1 | `single-char-edit` | `x`, `r`, buffer mutation, cursor clamp | 다음 playpack의 가장 작은 편집 성공 경험 |
| 2 | `insert-mode-entry` | `i`, `a`, `A`, insert text input, `esc` 복귀 | Bubble Tea key mapping과 runtime replay 표현 점검 필요 |
| 3 | `undo-redo-basic` | undo stack, redo stack, mutation event history | 억까/회복 상황을 게임적으로 만들기 좋음 |
| 4 | `open-line-edit` | `o`, `O`, newline insertion, insert mode entry | insert 모델 안정 후 진행 |
| 5 | operator grammar | operator pending, `dw`, `dd`, `cw`, `cc` | 작은 수정 playpack 이후 adventure intro 후보 |

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
