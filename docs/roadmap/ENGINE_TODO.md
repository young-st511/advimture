# Engine / Module Status

> Advimture는 엔진, 미들웨어, 어댑터를 분리해서 만든다. 이 문서는 현재 engine 상태와 hardening 후보만 둔다. 활성 slice는 항상 `docs/roadmap/PROGRAM.md`가 우선한다.

Last reviewed: 2026-06-02

## 진행 원칙

- 한 루프는 하나의 엔진 또는 모듈 계약만 다룬다.
- 새 Vim command 추가와 runtime/UI 연결을 같은 루프에서 섞지 않는다.
- TUI로 연결되는 순간 E2E assertion이 부족하면 구현을 멈추고 E2E를 보강한다.
- 새 engine 후보는 `command catalog -> vimengine -> runtime/tuiadapter -> content/playable -> E2E` 순서로 진행한다.
- progress 저장 포맷 변경은 engine 작업과 섞지 않는다.

## 현재 판단

현재 활성 engine slice는 없다.

Foundation engine은 tutorial과 incident 적용 런을 만들 수 있을 정도로 충분히 닫혔다. `PLATFORM-REVIEW-003`으로 mission/review/game loop는 한 차례 닫혔고, `PLAYABLE-QUALITY-BASELINE-001`도 완료됐다. 새 engine은 content 작성 병목이 확인될 때만 연다. 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 따른다.

## Release-Quality 모듈화 판정

현재 모듈 분리는 release-quality baseline을 진행하기에 충분하다.

Review: `docs/roadmap/MODULE_QUALITY_REVIEW_2026-06-02.md`

- `internal/vimengine`은 Vim-like state transition을 담당한다.
- `internal/runtime`은 exercise session, constraints, retry/hint를 담당한다.
- `internal/content`는 YAML loader, replay gate, coverage validator를 담당한다.
- `internal/scenario`는 exercise를 runbook layer로 감싸지만 target/keys/constraints를 바꾸지 않는다.
- `internal/playable`은 Bubble Tea model과 progress/review/playable flow를 조율한다.
- `internal/playableview`는 renderer와 viewport 안정성을 테스트할 수 있는 순수 surface다.
- `cmd/e2e-runner`와 `internal/e2estate`는 사람이 읽는 evidence와 typed app_state evidence를 함께 남긴다.

따라서 다음 품질 개선은 대개 narrow slice로 열 수 있다. 단, 아래 변경은 별도 ExecPlan/checkpoint가 필요하다.

- progress 저장 포맷 변경
- content schema 변경
- 새 의존성 추가
- pre-start modal처럼 Bubble Tea input routing과 runtime key trace를 분리해야 하는 UI 구조 변경
- visual block, count/register prefix처럼 engine blast radius가 큰 Vim capability

큰 구조 변경은 release-quality 목표에 필요하면 허용한다. 현재 기준은 큰 수정을 피하는 것이 아니라, evidence 없이 저장 포맷/schema/dependency처럼 제품 계약을 흔드는 변경을 피하는 것이다.

현재 판정상 release-quality baseline을 위해 새 Vim engine을 바로 열 필요는 없다. 다음 큰 구조 후보는 새 engine보다 `internal/playable` orchestration 분리 또는 pre-start modal input boundary다.

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
| `text-object-quote-pair` | `di"`, `ci"`, `yi"`, `di'`, `ci'`, `yi'` | tutorial/E2E completed |
| `open-line-edit` | `o`, `O` | tutorial/E2E completed |
| `repeat-last-change` | `.` | tutorial/E2E completed |
| `search-basic` | `/`, `n`, `N` | tutorial/E2E completed |
| `visual-char-line` | `v`, visual `d/y` | tutorial/E2E completed |
| `visual-line-basic` | `V`, linewise `d/y` | tutorial/E2E completed |
| `char-find-line` | `f/t`, `df/dt/cf/ct` | tutorial/E2E completed |

## Hardening Candidates

| 우선순위 | 후보 | 필요한 계약 | 주의 |
|----------|------|-------------|------|
| 1 | `quote-pair-hardening` | `ci(`, `ci{`의 pair parsing 범위 | escaped/nested pair는 별도 scope로 분리. `ci'`, `di'`, `yi'`는 completed |
| 2 | `command-choice-breadth` | 기존 implemented command 조합 | 새 engine 없이 content/constraints/E2E로 처리 |
| 3 | `platform-review-loop` | 저장 포맷 변경 없는 mission/review/game loop | completed, 후속은 progress schema 필요성 확인 시 재검토 |
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
