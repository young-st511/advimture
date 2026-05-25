# HUD-DENSITY-001 — Mission HUD Priority and Wrapping

Slice-ID: HUD-DENSITY-001
Created: 2026-05-26
Status: completed
Completed: 2026-05-26
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/tui-ui-qa-contract.md
- docs/exec-plans/active/hud-density-001-mission-hud-priority.md
- docs/exec-plans/completed/playtest-evidence-001-foundation-routes.md
- internal/playable/
- internal/playableview/
- test/e2e/

## 목표

Mission HUD가 “현재 해야 할 Vim 조작”을 가장 먼저 읽히게 만든다. RedTeam이 지적한 review/daily 과노출과 어색한 줄바꿈을 줄이되, 반복 학습 동기 정보는 완전히 제거하지 않는다.

## 해결할 문제

1. FTUE 초반에서 `복구 현황: 재점검 대상... 오늘의 복구 루트...`가 briefing 직후 길게 노출되어 새 key 학습 목표보다 먼저 보인다.
2. 긴 briefing과 review/daily가 terminal 폭에서 끊기며 `오염된 줄 묶음을 골`, `외 2건`처럼 미완성 문장처럼 보인다.
3. incident에서는 review/daily가 세계관 메타 정보로 유용하지만, 현재 목표와 console 접근을 늦추면 안 된다.

## UX 원칙

- Mission HUD의 첫 정보는 항상 `mission title`과 `briefing`이다.
- tutorial running 화면에서는 review/daily를 “복구 메모” 수준의 짧은 보조 정보로 접는다.
- incident running 화면에서는 review/daily를 유지하되 한 줄 길이를 제한한다.
- running cue line(`TRAINING BRIEF`, `OPERATOR JUDGMENT`)은 current action과 hint affordance를 담고, review/daily보다 눈에 잘 들어와야 한다.
- 좁은 width에서는 긴 보조 정보가 잘리는 것보다 축약되는 편이 낫다.

## 구체 설계

### 1. Recovery Summary 축약

현재:

```text
복구 현황: 재점검 대상: 경고 지점으로 이동하기: 미복구 · 오늘의 복구 루트: 경고 지점으로 이동하기(미복구) 외 2건 대기
```

변경 후보:

```text
복구 메모: 재점검 3건 · 다음: 경고 지점으로 이동하기
```

규칙:
- tutorial running에서는 `복구 메모: 재점검 N건 · 다음: <primary title>` 형식.
- incident running에서는 `복구 현황: 재점검 N건 · 잔류: <primary title>` 형식.
- succeeded/failed modal의 debrief는 기존 상세 정보를 유지하되, HUD running 화면에서는 축약한다.
- app_state review의 상세 값은 유지한다.

### 2. Briefing Wrapping

현재 renderer가 긴 문장을 terminal line에 그대로 흘려보내 일부 환경에서 어색하게 보인다.

변경 후보:
- HUD briefing은 terminal width를 기준으로 2줄까지 wrap한다.
- 2줄을 초과하면 마지막 줄을 `...`로 줄인다.
- 줄바꿈 단위는 rune width가 아니라 우선 단순 rune count 기반으로 시작한다. 한글 wide width exact 계산은 후속으로 미룬다.
- console buffer와 modal 위치를 밀지 않도록 renderer test에서 line count를 고정한다.

### 3. Action Cue 우선순위

현재:

```text
OPERATOR JUDGMENT · Inputs left: 3/3 · 판단: 목표 상태를 보고 이미 배운 Vim 동작을 선택하세요. · ?: hint  q: quit
```

변경 후보:
- cue line은 유지하되 review/daily보다 아래가 아니라 current briefing 직후 시각적으로 가까운 위치에 둔다.
- `복구 메모`는 cue line 다음 또는 더 짧은 secondary line으로 둔다.

## 제외

- modal layout 재설계
- success modal 중복 제거
- command-choice scenario 내용 변경
- `?` help 동작 변경
- progress 저장 포맷 변경
- 새 의존성 추가

## 수용 기준

- FTUE running screen에서 current mission title/briefing/cue가 review/daily보다 먼저 읽힌다.
- tutorial running screen의 review/daily summary는 한 줄로 축약된다.
- incident running screen도 review/daily가 한 줄을 넘지 않는다.
- command-choice briefing이 `오염된 줄 묶음을 골`처럼 미완성 문장처럼 끊기지 않는다.
- `app_state.review` assertion은 기존 상세 값을 유지한다.
- renderer tests와 focused E2E가 축약 문구와 wrapping 정책을 검증한다.
- `make e2e-playable`을 통과한다.

## Step 1: Contract and Tests

- [x] `tui-screen-contract`에 running HUD 축약 원칙 반영
- [x] renderer test에 tutorial/incident recovery summary 축약 expectation 추가
- [x] command-choice briefing wrap/truncate expectation 추가

## Step 2: Renderer Implementation

- [x] recovery summary formatter 추가
- [x] HUD briefing wrap/truncate helper 추가
- [x] tutorial/incident running 화면에서 축약 summary 사용

## Step 3: E2E Verification

- [x] `playable_ftue_first_five_route` 갱신/검증
- [x] `playable_incident_001_full` 갱신/검증
- [x] `playable_command_choice_scope` 갱신/검증
- [x] `go test ./...`
- [x] `make e2e-playable`
- [x] `git diff --check`

## 구현 결과

- `playable.Model`이 `Screen`에 playlist category, review count, primary review title을 넘긴다.
- running HUD는 tutorial에서 `복구 메모: 재점검 N건 · 다음: <title>`, incident에서 `복구 현황: 재점검 N건 · 잔류: <title>`로 축약한다.
- command/search/visual 같은 mode cue에서도 현재 playlist category에 맞춰 축약 summary를 유지한다.
- HUD briefing은 terminal width 기준 최대 2줄로 wrap하고, 초과 시 `...`로 축약한다.
- app_state review assertion은 기존 상세 daily route를 유지한다.

## 검증 결과

- `go test ./internal/playableview ./internal/playable`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 후속 Backlog

- `CHOICE-JUDGMENT-001`: command-choice cue/scenario를 비교 판단 중심으로 강화
- `SUCCESS-MODAL-001`: success modal heading과 record density 정리
- `HELP-AFFORDANCE-001`: `?`, retry, quit 실제 입력 UX 검증
