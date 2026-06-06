# Forward Plan — Foundation to Release Quality

Last reviewed: 2026-06-06
Status: current rolling plan

## 목적

이 문서는 앞으로 2~8주 동안 어떤 순서로 Advimture를 출시 가능한 품질의 Vim 학습 게임으로 다듬을지 정리한다.

- `PROGRAM.md`: 지금 무엇이 active인지 확인한다.
- `MIDTERM_TODO.md`: 현재 중기 보드를 확인한다.
- `FORWARD_PLAN.md`: 왜 그 순서로 가는지, 다음 몇 주의 방향을 확인한다.

작업 시작 전에는 `PROGRAM.md -> MIDTERM_TODO.md -> FORWARD_PLAN.md` 순서로 읽는다.

## 현재 판단

Foundation engine과 E2E loop는 충분히 튼튼해졌다. 다음 병목은 새 Vim command 수나 release candidate 포장이 아니라, 현재 playable loop가 "출시 가능한 품질"처럼 느껴지게 만드는 기준선이다.

현재 상태:

- Vim engine: tutorial/incident를 만들 수 있을 만큼 충분히 닫힘
- Content: tutorial coverage와 incident 001~007이 있음
- E2E: long route final/timeline evidence까지 보강됨
- UI/UX: Mission HUD, Runbook Console, floating modal 기반은 있음
- 출시감: mission/review loop, content breadth, quote/pair hardening, UI polish, release readiness, fresh playtest, first-run polish, pre-RC hardening, playable quality baseline은 한 차례 닫혔다. 하지만 바로 출시할 계획은 없으므로 다음 병목은 release candidate 문서가 아니라 후속 world/UX/content polish를 필요한 순간에 좁게 여는 것이다.

`FOUNDATION-EXIT-001` review 결과 Foundation은 조건부 통과했고, `PLAYTEST-GATE-001`에서 P0/P1 blocker는 확인되지 않았다. `FIRST-RUN-POLISH-001`로 첫 실행 cue와 viewport smoke를 닫았고, `PRE-RC-HARDENING-001`로 첫 5분/대표 incident evidence도 한 번 더 보강했다. `PLAYABLE-QUALITY-BASELINE-001`로 세계관/UX/UI/콘텐츠 기획/엔진 모듈화 기준을 닫은 뒤, 2026-06-06에는 `CONTENT-ARC-POLISH-001 -> JUDGMENT-DRILL-REVIEW-001 -> UI-CONSOLE-POLISH-001`까지 완료했다. 따라서 다음 순서는 **fresh evidence review -> narrow applied content or UI polish -> later release candidate prep**으로 간다.

## 0. 운영 원칙

- 새 기능보다 먼저 현재 evidence를 본다.
- E2E가 부족하다고 느껴지면 구현을 멈추고 verification을 보강한다.
- 새 engine은 하나의 command contract만 다룬다.
- progress 저장 포맷은 사용자 승인 전까지 변경하지 않는다.
- long route E2E는 `screen_timeline.txt`와 `screen_final.txt`를 남긴다.
- 문서가 stale해질 수 있는 변경을 하면 `PROGRAM.md`, `MIDTERM_TODO.md`, 이 문서를 함께 확인한다.

## 1. Foundation Exit Result

### FOUNDATION-EXIT-001 — Foundation Exit Review

Status: completed
Review: `docs/roadmap/FOUNDATION_EXIT_REVIEW_2026-05-30.md`

판정:

- Foundation은 다음 단계로 넘어가도 된다.
- 판정은 "출시 가능"이 아니라 "출시 가능한 게임 루프를 만들 수 있는 foundation 통과"다.
- P0 blocker는 없다.
- 다음 병목은 새 Vim command 수가 아니라 mission/review/game loop다.

확인한 기준:

- `go test ./...`: pass
- `make e2e-playable`: pass
- long incident final/timeline evidence spot review 완료

## 2. Platform Review Result

### PLATFORM-REVIEW-003 — Mission/Review Game Loop

Status: completed
ExecPlan: `docs/exec-plans/completed/platform-review-003-mission-review-loop.md`

결과:

- 성공 debrief가 `이번 복구`, `최단 복구`, `목표 입력`, `잔류 리스크`, `다음 출격`을 보여준다.
- 마지막 dispatch에서 review 후보가 남아 있으면 `다음 출격: enter`로 primary review exercise에 재진입한다.
- progress schema v2, daily streak, persisted review due date는 여전히 도입하지 않았다.

검증:

- `go test ./...`: pass
- `make e2e-playable`: pass
- focused review/debrief E2E: pass

## 3. Immediate Plan

### CONTENT-BREADTH-002 — Repeat Change Choice

Status: completed
ExecPlan: `docs/exec-plans/completed/content-breadth-002-repeat-choice.md`

결과:

- `incident-005-command-choice`에 fifth beat `command-choice-repeat-change-001`을 추가했다.
- 같은 단어 교체가 이어질 때 두 번째 변경을 다시 입력하지 않고 `.`으로 반복하는 판단을 훈련한다.
- focused command-choice E2E는 final/timeline/app_state evidence를 남긴다.

검증:

- `go test ./internal/content ./internal/playable`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass

### QUOTE-PAIR-HARDEN-001 — Quote/Pair Text Object Hardening

Status: completed
ExecPlan: `docs/exec-plans/completed/quote-pair-harden-001-single-quote.md`

결과:

- `di'`, `ci'`, `yi'` single quote 내부 object를 engine/runtime/content/E2E까지 연결했다.
- `tutorial-91-text-object-quote-pair`는 7문항으로 확장됐고 8문항 제한을 유지한다.
- `i(`, `i{`, nested/escaped/around/count/register는 후속 hardening으로 남긴다.

검증:

- `go test ./internal/vimengine ./internal/runtime ./internal/content`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_quote_pair_full.yaml`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass

### UI-POLISH-002 — Command Memory Cue

Status: completed
ExecPlan: `docs/exec-plans/completed/ui-polish-002-command-memory.md`

결과:

- tutorial running 화면은 `기억할 명령: ...`으로 current exercise command memory를 보여준다.
- incident running 기본 화면은 command memory를 숨기고, hint/failure 후 `참고 명령: ...`으로 점진 공개한다.
- 저장 포맷, content schema, app_state schema는 바꾸지 않았다.

검증:

- `go test ./internal/playable`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_coaching_panel.yaml`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_hint_affordance.yaml`: pass

### RELEASE-READINESS-001 — First Release Readiness

Status: completed
ExecPlan: `docs/exec-plans/completed/release-readiness-001-first-release.md`

목표: 첫 공개 전 설치/실행/검증/터미널 크기/known limitations/release build 기준을 정리한다.

완료 결과:

- README가 현재 실행 가능한 Vim adventure game, 진행 저장, reset, known limitations를 설명한다.
- Makefile에 `build`, `test`, `release-check` target을 추가했다.
- `make release-check`는 `make test`, `make build`, `make e2e-playable`을 순서대로 실행한다.

검증:

- `go test ./...`: pass
- `make release-check`: pass

### PLAYTEST-GATE-001 — Fresh Playtest Release Gate

Status: completed
ExecPlan: `docs/exec-plans/completed/playtest-gate-001-fresh-release-gate.md`
Review: `docs/roadmap/PLAYTEST_GATE_2026-05-30.md`

목표: README 기준으로 처음 실행하는 플레이어 관점에서 첫 5분, tutorial 확장, incident 3개 이상을 직접 훑고 출시 전 blocker와 후속 wishlist를 분리한다.

완료 결과:

- P0/P1 blocker는 없다.
- P2 first-run polish 후보가 있었다.
- 후속 `FIRST-RUN-POLISH-001`은 완료됐다.

## 4. Next Midterm Sequence

### 1. PLAYABLE-QUALITY-BASELINE-001 — Release-Quality Playable Baseline

Status: completed
ExecPlan: `docs/exec-plans/completed/playable-quality-baseline-001-release-quality-baseline.md`
Audit: `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md`

목표: 바로 출시하지 않더라도 현재 playable loop가 출시 가능한 품질에 가까워지도록 세계관, UX/UI, 콘텐츠 목표, 엔진/모듈화 기준을 세우고 개선한다.

포함 후보:

- world-frame/scenario-tone release-quality 기준 정리
- first-run/tutorial/incident/review UI 기준과 evidence spot review
- 출시할 만한 콘텐츠 양과 episode arc 기획
- 현재 엔진/모듈화가 품질 목표를 감당하는지 판정
- P0/P1/P2를 나눠 focused fix 후보 선정

제외:

- 바로 release candidate를 묶거나 tag를 찍는 작업
- 신규 대형 tutorial/incident 구현
- 사용자 확인 없는 progress 저장 포맷/content schema/의존성 변경

출구:

- 세계관/UX/UI/콘텐츠/엔진 기준과 현재 판정이 문서화됨
- 필요한 small fix가 적용되거나 큰 구조 변경 후보가 별도 ExecPlan으로 분리됨
- `make release-check`, `git diff --check` pass

### 2. CONTENT-QUALITY-PLAN-001 — Release-Quality Content Plan

Status: completed as planning
Plan: `docs/roadmap/CONTENT_QUALITY_PLAN_001.md`

목표: 신규 콘텐츠를 바로 구현하지 않고, 출시할 만한 첫 콘텐츠 양과 arc를 기획한다.

포함 후보:

- tutorial first tour에서 first dispatch까지의 전환 설계
- incident 001~003 relay station arc 재정리
- command-choice judgment drill 후보 목록
- 다음 engine candidate를 콘텐츠 병목 기준으로 선정

제외:

- content YAML 구현
- 새 engine 구현
- schema 변경

### 3. CONTENT-ARC-POLISH-001 — First Dispatch Arc Polish

Status: completed
ExecPlan: `docs/exec-plans/completed/content-arc-polish-001-first-dispatch-arc.md`

목표: 새 YAML 콘텐츠 없이 incident 001~003의 기존 title/briefing/feedback을 world-frame 기준으로 spot polish한다.

완료 결과:

- incident 001은 원인 신호 추적, 002는 구조 재동기화, 003은 오염 구간 격리로 읽힌다.
- 변경 범위는 scenario title/briefing/feedback copy와 docs에 머물렀다.
- exercise target, optimal keys, constraints는 변경하지 않았다.

### 4. JUDGMENT-DRILL-REVIEW-001 — Command Choice as Core Identity

Status: completed
ExecPlan: `docs/exec-plans/completed/judgment-drill-review-001-command-choice.md`

목표: command-choice를 Advimture의 핵심 차별점인 "상황에 맞는 Vim 도구 선택"으로 강화한다.

완료 결과:

- `incident-005-command-choice`의 5 beat는 scope/range/inline/reuse/repeat-change 판단 질문에 매핑된다.
- success/failure copy는 command 이름보다 선택 이유를 먼저 설명한다.
- 후속 후보는 line reuse, search-then-scope, bracket-pair hardening으로 문서에만 정리했다.

### 5. UI-CONSOLE-POLISH-001 — Runbook Console Product Feel

Status: completed
ExecPlan: `docs/exec-plans/completed/ui-console-polish-001-action-identity.md`

목표: action line의 화면 label과 E2E action 의미를 분리한다.

완료 결과:

- `FocusPanel.actions` internal DTO와 `app_state.ui.focus_panel.actions`를 추가했다.
- 사용자 화면 label은 `다음 단계`, `다시 시도`, `다음 runbook`, `출격 완료` 같은 제품 톤을 사용한다.
- progress schema, content schema, dependency 변경은 없다.

### 6. RELEASE-CANDIDATE-001 — Release Candidate Prep

Status: later

목표: 실제로 공개 후보를 묶기로 결정했을 때 release note, known limitations, final evidence bundle, tag 후보 기준을 정리한다.

다시 열리는 조건:

- 사용자가 실제 공개 후보 준비를 원함
- 현재 known limitations와 release gate가 충분히 정리됨

### 7. POST-MVP-CONTENT-001 — Post Baseline Content/Engine Expansion

Status: later

목표: 첫 공개 후보가 닫힌 뒤 content breadth 또는 다음 engine hardening을 evidence 기반으로 고른다.

후보:

- line reuse applied drill
- search-then-act incident
- bracket pair text object hardening
- progress v2 decision 재검토

## 5. Completed Midterm Sequence

### 1. PRE-RC-HARDENING-001 — First Release Hardening

Status: completed
ExecPlan: `docs/exec-plans/completed/pre-rc-hardening-001-first-release-hardening.md`
Review: `docs/roadmap/PRE_RC_HARDENING_2026-06-02.md`

완료 결과:

- 첫 5분 route와 대표 incident route evidence를 spot review했고 P0/P1 blocker는 없었다.
- 긴 incident hint cue가 한 줄에서 잘리는 문제를 terminal width 기반 wrapping으로 보강했다.
- incident hint affordance를 80x30 viewport smoke와 final/timeline/app_state evidence로 고정했다.
- 긴 incident full route의 app_state evidence gap을 닫았다.

검증:

- focused E2E: pass
- `go test ./...`: pass
- `make release-check`: pass
- `git diff --check`: pass

### 2. FIRST-RUN-POLISH-001 — First Run Polish

Status: completed
ExecPlan: `docs/exec-plans/completed/first-run-polish-001-cue-viewport.md`

목표: 새 engine/content/schema 없이 첫 실행 tutorial cue와 release evidence만 좁게 다듬는다.

완료 결과:

- tutorial running cue에서 `기억할 명령`과 `Coach` key 중복을 제거했다.
- success/failure floating modal 주변의 detailed review/daily line을 숨겼다.
- 80x24 success/failure viewport smoke fixture를 `make e2e-playable`에 추가했다.
- open-line/repeat/search/visual/char-find full route의 final/timeline/app_state evidence를 보강했다.

검증:

- focused E2E: pass
- `go test ./...`: pass
- `make release-check`: pass
- `git diff --check`: pass

### 3. CONTENT-BREADTH-002 — Applied Content Expansion

Status: completed

목표: 새 engine 없이 기존 command를 조합하는 applied incident와 command-choice를 늘린다.

후보:

- line reuse choice: 검증된 줄 전체를 `V` + `y` + `p`로 재사용
- repeat-change choice: 같은 변경을 `.`로 반복할지 판단
- search-then-act incident: `/`, `n`, `N`으로 찾고 적절한 edit command 선택
- mixed incident 008: 3~5 beat 이하로 제한한 생존 어드벤처 run

품질 기준:

- 한 beat는 하나의 판단만 요구한다.
- 새 command를 소개하지 않는다.
- long route에는 final/timeline evidence를 남긴다.

완료 결과:

- repeat-change choice를 fifth beat로 추가했다.
- 남은 후보인 line reuse, search-then-act, mixed incident 008은 release 전 content polish 후보로 유지한다.

### 4. QUOTE-PAIR-HARDEN-001 — Quote/Pair Text Object Hardening

Status: completed

목표: 기존 `i"` text object를 `i'`, `i(`, `i{`로 확장한다.

포함:

- `ci'`, `di'`, `yi'`
- config/JSON/function-argument style exercise

제외:

- `ci(`, `ci{`
- nested pair
- escaped quote
- around object
- count/register prefix

완료 결과:

- 첫 hardening scope는 `i'`로 닫았다.
- `i(`, `i{`는 bracket pair hardening 후보로 유지한다.

### 5. UI-POLISH-002 — Release UI Polish

Status: completed

목표: 출시 전 화면을 개발 UI가 아니라 Vim adventure console처럼 읽히게 다듬는다.

후보:

- color/emphasis pass
- learned command memory
- wide layout side rail
- pre-start briefing modal

주의:

- 화면 문구보다 `app_state` 검증을 우선한다.
- color 없는 환경에서도 의미가 보존되어야 한다.

완료 결과:

- 첫 release polish는 command memory cue로 닫았다.
- color/emphasis, side rail, pre-start modal은 후속 UI 후보로 유지한다.

## 6. Release Readiness

첫 공개 전 필요 항목:

- `README.md`에 설치/실행/테스트 안내: completed
- progress 파일 안전성 점검: completed
- release build command: completed
- known limitations 정리: completed
- 첫 실행 경험 검증: completed
- 터미널 크기별 smoke: completed for 80x24 success/failure modal

첫 공개 기준:

- tutorial route가 막힘 없이 진행된다.
- incident route가 3개 이상 게임처럼 읽힌다.
- 실패/재시도/힌트가 플레이를 막지 않는다.
- `make e2e-playable`이 통과한다.
- long incident evidence가 남는다.
- progress schema 변경 없이 저장/재개가 안전하다.

현재 release readiness, first-run polish, pre-RC hardening은 닫혔다. 다음은 release candidate prep으로 공개 후보 설명과 evidence를 한곳에 묶는다.

## 7. Long-Run Candidates

아래는 출시 전 필수가 아니다.

- progress schema v2
- spaced review due date
- daily streak/history
- macros/register/count prefix
- visual block
- regex search/highlight/history
- terminal cell-grid viewport parser

이 후보들은 실제 플레이 evidence로 병목이 확인될 때만 연다.

## 8. 문서 업데이트 규칙

각 slice 종료 시:

1. `PROGRAM.md`: active/recent completed/next 후보 갱신
2. `MIDTERM_TODO.md`: 현재 중기 보드 상태 갱신
3. `FORWARD_PLAN.md`: 추천 순서나 gate가 바뀌었으면 `Last reviewed`와 관련 섹션 갱신
4. `CHANGES.md`: 가정 변경은 append-only로 기록
5. 오래된 review/health 문서는 `docs/roadmap/archive/`로 이동

이 규칙을 지키지 못하면 다음 작업 전에 docs cleanup slice를 먼저 연다.
