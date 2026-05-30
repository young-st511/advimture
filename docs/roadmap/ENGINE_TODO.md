# Engine / Module Status

> Advimture는 엔진, 미들웨어, 어댑터를 분리해서 만든다. 이 문서는 현재 engine 상태와 hardening 후보만 둔다. 활성 slice는 항상 `docs/roadmap/PROGRAM.md`가 우선한다.

Last reviewed: 2026-05-30

## 진행 원칙

- 한 루프는 하나의 엔진 또는 모듈 계약만 다룬다.
- 새 Vim command 추가와 runtime/UI 연결을 같은 루프에서 섞지 않는다.
- TUI로 연결되는 순간 E2E assertion이 부족하면 구현을 멈추고 E2E를 보강한다.
- 새 engine 후보는 `command catalog -> vimengine -> runtime/tuiadapter -> content/playable -> E2E` 순서로 진행한다.
- progress 저장 포맷 변경은 engine 작업과 섞지 않는다.

## 현재 판단

현재 활성 engine slice는 없다.

Foundation engine은 tutorial과 incident 적용 런을 만들 수 있을 정도로 충분히 닫혔다. 다음 작업은 바로 새 engine을 여는 것보다 `PLAN-REFRESH-009`에서 Foundation exit review를 먼저 수행한 뒤 고른다.

## 완료된 핵심 모듈

| 영역 | 상태 | 비고 |
|------|------|------|
| `internal/vimengine` foundation | completed | `State + Key -> State + Events` 순수 전이 |
| `internal/runtime` | completed | exercise session, constraints, retry/hint |
| `internal/content` | completed | YAML loader, replay gate, coverage validator |
| `internal/scoring` | completed | grade/key count/intent scoring |
| `internal/scenario` | completed | exercise를 scenario/runbook layer로 감싸기 |
| `internal/tuiadapter` | completed | Bubble Tea input/view model 변환 |
| `internal/progressadapter` / `internal/progress` | completed | progress v1 저장 경계 유지 |
| `internal/vimoracle` | completed | optional Neovim oracle comparison |
| `cmd/e2e-runner` | completed | pty E2E, app_state, final/timeline evidence |

## Implemented Vim Clusters

| Cluster | 핵심 command | 상태 |
|---------|--------------|------|
| `normal-motion-basic` | `h`, `j`, `k`, `l` | tutorial/E2E completed |
| `survival-save-quit` | `esc`, `:q!`, `:wq` | tutorial/E2E completed |
| `word-motion-basic` | `w`, `b`, `e` | tutorial/E2E completed |
| `whole-file-navigation` | `gg`, `G`, `0`, `$` | tutorial/E2E completed |
| `vim-ex-command-substitute` | `:s`, `:%s`, `:2,3s` | tutorial/E2E completed |
| `single-char-edit` / `insert-mode-entry` / `undo-redo-basic` | `x`, `r`, `i`, `a`, `A`, `u`, `ctrl+r` | tutorial/E2E completed |
| `delete-with-motion` / `change-with-motion` | `dw`, `d$`, `dd`, `cw`, `c$`, `cc` | tutorial/E2E completed |
| `yank-put-basic` | `yw`, `y$`, `yy`, `p`, `P` | tutorial/E2E completed |
| `text-object-inner-word` | `diw`, `ciw`, `yiw` | tutorial/E2E completed |
| `text-object-quote-pair` | `di"`, `ci"`, `yi"` | tutorial/E2E completed |
| `open-line-edit` | `o`, `O` | tutorial/E2E completed |
| `repeat-last-change` | `.` | tutorial/E2E completed |
| `search-basic` | `/`, `n`, `N` | tutorial/E2E completed |
| `visual-char-line` | `v`, visual `d/y` | tutorial/E2E completed |
| `visual-line-basic` | `V`, linewise `d/y` | tutorial/E2E completed |
| `char-find-line` | `f/t`, `df/dt/cf/ct` | tutorial/E2E completed |

## Hardening Candidates

| 우선순위 | 후보 | 필요한 계약 | 주의 |
|----------|------|-------------|------|
| 1 | `quote-pair-hardening` | `ci'`, `di'`, `yi'`, `ci(`, `ci{`의 pair parsing 범위 | escaped/nested pair는 별도 scope로 분리 |
| 2 | `platform-review-loop` | 저장 포맷 변경 없는 mission/review/game loop | engine 작업이 아니며 progress schema 변경 금지 |
| 3 | `command-choice-breadth` | 기존 implemented command 조합 | 새 engine 없이 content/constraints/E2E로 처리 |
| 4 | `search-hardening` | `?`, regex, highlight, history 중 일부 | `?`는 현재 hint key와 충돌하므로 UX 결정 필요 |
| 5 | `visual-advanced` | visual block, multi-line charwise, count/register prefix | blast radius 큼, E2E selection contract 선행 |
| 6 | `automation-advanced` | macro/register/count prefix | 출시 전 필수 아님 |

## Engine Opening Checklist

새 engine slice를 열기 전에:

- `PROGRAM.md`에 active slice가 있다.
- command cluster가 `docs/gameplay/command-catalog.md` 또는 ExecPlan에서 승인됐다.
- 제외 범위가 명확하다.
- engine unit test와 runtime replay test가 먼저 정의된다.
- content/playable/E2E는 engine 통과 뒤 별도 step으로 연결한다.
- TUI E2E assertion이 부족하면 먼저 verification 계약을 확장한다.
