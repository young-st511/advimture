# Post-Polish Playtest Review — 2026-06-06

Status: completed
Related ExecPlan: `docs/exec-plans/completed/post-polish-playtest-001-fresh-product-loop-review.md`

## 질문

First Dispatch/Judgment/UI console polish 이후, 다음에 열어야 할 작업은 무엇인가?

## Evidence Inventory

| Evidence | 읽는 흐름 | 판정 |
|----------|-----------|------|
| `playable_ftue_first_five_route` | first tour가 tutorial 초반에서 Runbook Dispatch 진입감으로 이어지는가? | 유지 |
| `playable_incident_001_full` | 원인 신호 추적이 첫 hotfix 작전으로 읽히는가? | 유지 |
| `playable_incident_002_full` | 구조 재동기화가 quote/mirror 보존 판단으로 읽히는가? | 유지 |
| `playable_incident_003_full` | 오염 구간 격리가 visual/range 판단으로 읽히는가? | 유지 |
| `playable_command_choice_scope` | command-choice가 command 암기가 아니라 도구 선택 판단으로 읽히는가? | 보강 필요 |
| `playable_review_queue` | 성공 후 residual risk와 next dispatch가 반복 동기로 이어지는가? | 유지 |

P0/P1 blocker는 보이지 않는다. 다음 병목은 release candidate 포장이 아니라 command-choice의 판단 폭을 한 칸 더 선명하게 넓히는 것이다.

## Options

| Option | 장점 | 리스크 | 판정 |
|--------|------|--------|------|
| Line reuse applied drill | 기존 `visual-line-basic`만 사용하고, 직접 재입력 vs 줄 전체 재사용 판단을 선명하게 추가한다. | 기존 linewise tutorial과 겹칠 수 있다. scenario copy로 "판단 문제"임을 분리해야 한다. | 선택 |
| Search-then-scope applied drill | incident 감각이 강하고 search + linewise scope 조합이 실전적이다. | command-choice route가 길어지고, search/navigation cue가 섞여 이번 scope보다 크다. | 후속 |
| Bracket-pair hardening | 구조 편집의 실용성이 크다. | 새 engine capability가 필요하다. 이번 제한과 맞지 않는다. | 보류 |
| 추가 UI console polish | 제품 느낌을 더 다듬을 수 있다. | 직전 slice에서 action language 계약을 닫았다. 지금은 콘텐츠 판단 폭이 더 큰 병목이다. | 보류 |

## Decision

다음 slice는 `LINE-REUSE-APPLIED-001`로 한다.

이유:

- 새 Vim engine, 새 dependency, content schema, progress schema를 열지 않는다.
- `incident-005-command-choice`의 기존 정체성인 "상황에 맞는 Vim 도구 선택"을 더 직접적으로 강화한다.
- quote value reuse와 repeat-change reuse 사이에서 빠진 "줄 전체 재사용" 판단을 evidence로 검증할 수 있다.
- focused E2E 하나로 final/timeline/app_state/key trace/progress evidence를 모두 확인할 수 있다.

## Exit Criteria

- `LINE-REUSE-APPLIED-001`이 completed로 이동한다.
- `playable_command_choice_scope`가 여섯 beat route를 검증한다.
- `go test ./...`, `make release-check`, `git diff --check`가 통과한다.
- `go.mod`, `go.sum`, `internal/progress/`, content schema 변경이 없다.
