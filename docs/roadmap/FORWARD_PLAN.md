# Forward Plan — Foundation to Release Quality

Last reviewed: 2026-06-20
Status: current rolling plan

## 목적

이 문서는 앞으로 2~8주 동안 어떤 순서로 Advimture를 출시 가능한 품질의 Vim 학습 게임으로 다듬을지 정리한다.

- `PROGRAM.md`: 지금 무엇이 active인지 확인한다.
- `MIDTERM_TODO.md`: 현재 중기 보드를 확인한다.
- `FORWARD_PLAN.md`: 왜 그 순서로 가는지, 다음 몇 주의 방향을 확인한다.

작업 시작 전에는 `PROGRAM.md -> MIDTERM_TODO.md -> FORWARD_PLAN.md` 순서로 읽는다.

## 현재 판단

Foundation engine과 E2E loop는 충분히 튼튼해졌다. v0.2.0 배포 후 다음 병목은 release candidate 포장이 아니라, UX 재검토에서 발견된 modal/hint/final evidence blind spot을 닫고 fresh evidence로 다음 구현 후보를 고르는 일이다.

현재 상태:

- Vim engine: tutorial/incident를 만들 수 있을 만큼 충분히 닫힘
- Content: tutorial coverage와 incident 001~008이 있음
- E2E: long route final/timeline evidence까지 보강됨
- UI/UX: Mission HUD, Runbook Console, viewport overlay modal, action footer, running hint/quit action, final-frame evidence가 있으며, modal placement/hint wrapping/wide-char evidence 후속 보강이 active
- 출시감: mission/review loop, content breadth, quote/pair hardening, UI polish, release readiness, fresh playtest, first-run polish, pre-RC hardening, playable quality baseline, modal/action hierarchy hardening, Homebrew `v0.2.0` 배포는 한 차례 닫혔다. 다음 병목은 `UI-MODAL-ACTION-HIERARCHY-002`를 닫은 뒤 fresh evidence 기반으로 다음 content/UX 후보를 선택하는 것이다.

`FOUNDATION-EXIT-001` review 결과 Foundation은 조건부 통과했고, `PLAYTEST-GATE-001`에서 P0/P1 blocker는 확인되지 않았다. `FIRST-RUN-POLISH-001`로 첫 실행 cue와 viewport smoke를 닫았고, `PRE-RC-HARDENING-001`로 첫 5분/대표 incident evidence도 한 번 더 보강했다. `PLAYABLE-QUALITY-BASELINE-001` 이후 `CONTENT-ARC-POLISH-001 -> JUDGMENT-DRILL-REVIEW-001 -> UI-CONSOLE-POLISH-001 -> POST-POLISH-PLAYTEST-001 -> LINE-REUSE-APPLIED-001 -> SEARCH-THEN-SCOPE-APPLIED-001 -> BRACKET-PAIR-HARDEN-001 -> NEXT-PLAYTEST-REVIEW-001 -> REVIEW-LOOP-MOTIVATION-001 -> COMMAND-CHOICE-BREADTH-002 -> POST-BREADTH-PLAYTEST-REVIEW-001 -> UI-MODAL-ACTION-HIERARCHY-001`까지 완료했다. v0.2.0 배포 후에는 **`UI-MODAL-ACTION-HIERARCHY-002`로 modal placement, hint utility line, final-frame evidence를 닫고 나서 기존 engine 기반 applied content breadth를 검토한다**로 간다.

## 0. 운영 원칙

- 새 기능보다 먼저 현재 evidence를 본다.
- E2E가 부족하다고 느껴지면 구현을 멈추고 verification을 보강한다.
- 새 engine은 하나의 command contract만 다룬다.
- progress 저장 포맷은 사용자 승인 전까지 변경하지 않는다.
- long route E2E는 `screen_timeline.txt`와 `screen_final.txt`를 남긴다.
- failed/succeeded/debrief modal은 문자열 존재만으로 통과시키지 않고, final viewport에서 primary action이 보이는지 확인한다.
- final viewport evidence는 한글 wide-width continuation cell 때문에 사람이 읽는 문구가 깨져 보이지 않아야 한다.
- 문서가 stale해질 수 있는 변경을 하면 `PROGRAM.md`, `MIDTERM_TODO.md`, 이 문서를 함께 확인한다.

## 0.4. Active Post-release Hardening

### UI-MODAL-ACTION-HIERARCHY-002 — Modal/Hint Post-release Hardening

Status: active
ExecPlan: `docs/exec-plans/active/ui-modal-action-hierarchy-002-post-release.md`

목표:

- failed/succeeded modal이 buffer/status/grade 뒤에 append된 것처럼 보이지 않게 console surface 위 overlay placement를 고정한다.
- running hint/quit utility action이 cue/hint body와 같은 wrapping group에 섞이지 않게 한다.
- `screen_final.txt`의 한글 wide-width evidence가 사람이 읽을 수 있게 유지된다.
- action label과 summary evidence 의미를 정리한다.

제외:

- release/tag/push
- progress 저장 포맷 변경
- content schema 변경
- 새 dependency

## 0.5. Completed UX Hardening

### UI-MODAL-ACTION-HIERARCHY-001 — Modal Action Hierarchy Hardening

Status: completed
ExecPlan: `docs/exec-plans/completed/ui-modal-action-hierarchy-001.md`

결과:

- 실패/성공/debrief 화면을 console 뒤에 붙은 블록이 아니라 viewport 기준 modal decision surface로 보이게 한다.
- `다시 시도`, `다음 단계`, `힌트`, `종료`를 action footer로 분리한다.
- running 상태의 hint/quit affordance도 일반 문구가 아니라 utility action으로 읽히게 한다.
- final-frame/viewport QA를 보강해 이 문제가 다시 문자열 존재만으로 통과하지 않게 한다.
- incident briefing은 exact command sequence보다 상황과 판단 목표를 먼저 말한다.

검증:

- `go test ./internal/content ./internal/playable ./internal/playableview`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass

제외:

- release/tag/push
- progress 저장 포맷 변경
- content schema 변경
- 새 dependency

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
- 당시 `i(`, `i{`는 후속 hardening으로 남겼고, 이후 `BRACKET-PAIR-HARDEN-001`에서 같은 줄의 비중첩 scope를 완료했다.

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

### 0. POST-POLISH-PLAYTEST-001 — Fresh Product Loop Review

Status: completed
ExecPlan: `docs/exec-plans/completed/post-polish-playtest-001-fresh-product-loop-review.md`
Review: `docs/roadmap/POST_POLISH_PLAYTEST_2026-06-06.md`

결과:

- first tour, first dispatch, judgment drill, review loop evidence를 fresh review했다.
- P0/P1 blocker는 보이지 않았다.
- 다음 최선 후보를 `LINE-REUSE-APPLIED-001`로 판정했다.

### 0.1. LINE-REUSE-APPLIED-001 — Line Reuse Choice Drill

Status: completed
ExecPlan: `docs/exec-plans/completed/line-reuse-applied-001-line-reuse-drill.md`

결과:

- `incident-005-command-choice`에 sixth beat `command-choice-line-reuse-001`을 추가했다.
- 검증된 route 줄 전체를 직접 다시 입력하지 않고 linewise `V` + `y` + `p`로 backup 아래에 재사용하는 판단을 훈련한다.
- focused command-choice E2E는 6-beat route와 final/timeline/app_state evidence를 남긴다.

검증:

- `go test ./internal/content ./internal/playable`: pass
- `go test ./...`: pass
- `make release-check`: pass

### 0.2. SEARCH-THEN-SCOPE-APPLIED-001 — Search Then Scope Drill

Status: completed
ExecPlan: `docs/exec-plans/completed/search-then-scope-applied-001.md`

결과:

- `incident-008-search-scope` 1-beat applied incident를 추가했다.
- `/breach`로 marker를 찾은 뒤 linewise `V`, `j`, `d`로 줄 묶음을 격리하는 판단을 검증한다.
- focused E2E는 final/timeline/app_state/key trace evidence를 남긴다.

검증:

- `go test ./internal/content`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_008_full.yaml`: pass

### 0.3. BRACKET-PAIR-HARDEN-001 — Parenthesis/Brace Inner Text Object

Status: completed
ExecPlan: `docs/exec-plans/completed/bracket-pair-harden-001.md`

결과:

- `text-object-quote-pair` cluster를 같은 줄의 비중첩 parenthesis/brace 내부 object까지 확장했다.
- `di(`, `ci(`, `yi(`, `di{`, `ci{`, `yi{` 6문항을 `tutorial-95-bracket-pair`로 분리했다.
- `playable_bracket_pair_full` E2E는 final/timeline/app_state/key trace evidence를 남긴다.

검증:

- `go test ./internal/vimengine ./internal/runtime ./internal/content`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_bracket_pair_full.yaml`: pass

### 0.4. NEXT-PLAYTEST-REVIEW-001 — Fresh Direction Review

Status: completed
ExecPlan: `docs/exec-plans/completed/next-playtest-review-001-fresh-direction.md`
Review: `docs/roadmap/NEXT_PLAYTEST_REVIEW_2026-06-07.md`

결과:

- first tour, bracket pair tutorial, first dispatch, judgment drill, incident 008, review loop evidence를 spot review했다.
- P0/P1 blocker는 확인되지 않았다.
- 다음 최선 후보는 `REVIEW-LOOP-MOTIVATION-001`로 판정했다.

검증:

- `git diff --check`: pass
- `go test ./...`: pass

### 0.5. REVIEW-LOOP-MOTIVATION-001 — Review Loop Motivation Polish

Status: completed
ExecPlan: `docs/exec-plans/completed/review-loop-motivation-001.md`

목표:

- success debrief, 잔류 리스크, 오늘의 복구 루트, 다음 출격 언어를 더 자연스럽게 만든다.
- tutorial success의 review 안내가 실제 primary action처럼 읽히지 않게 한다.
- incident success의 Runbook Dispatch 톤과 app_state action id 계약은 유지한다.

결과:

- tutorial success는 `재점검 메모`/`나중에 다시 풀기`로 review 후보를 보조 동기로 표시한다.
- incident success는 `잔류 리스크`/`다음 출격 후보`로 review 후보를 Runbook Dispatch 톤에 맞춰 표시한다.
- `next`, `next_dispatch` 등 action id/label 계약과 progress 저장 포맷은 변경하지 않았다.

제외:

- progress schema 변경
- content schema 변경
- 새 Vim command/engine capability
- 새 dependency
- release candidate/tag 작업

검증:

- `go test ./internal/playable ./internal/playableview ./internal/e2estate`: pass
- `playable_review_queue`, `playable_debrief_success`, `playable_ftue_first_five_route`, `playable_viewport_success_modal_80x24`: pass

### 0.6. COMMAND-CHOICE-BREADTH-002 — Command Choice Breadth

Status: completed
ExecPlan: `docs/exec-plans/completed/command-choice-breadth-002.md`

목표:

- 새 Vim command를 열기보다 기존 engine command만 사용해 command-choice 판단 breadth를 늘린다.
- 후보는 line/search/scope/reuse 판단이 겹치지 않도록, `docs/gameplay/command-choice-drills.md`의 판단 질문을 기준으로 고른다.

결과:

- `incident-005-command-choice`를 일곱 beat route로 확장했다.
- `command-choice-bracket-scope-001`은 `old-value`처럼 단어 단위로는 충분하지 않은 괄호 내부 인자 전체를 `ci(`로 교체하는 scope 판단을 검증한다.
- 새 command, 새 schema, progress 저장 포맷, dependency는 추가하지 않았다.

제외:

- progress schema 변경
- content schema 변경
- 새 Vim command/engine capability
- 새 dependency
- release candidate/tag 작업

검증:

- `go test ./internal/content ./internal/playable`: pass
- `playable_command_choice_scope`: pass

### 0.7. POST-BREADTH-PLAYTEST-REVIEW-001 — Post Breadth Playtest Review

Status: completed
ExecPlan: `docs/exec-plans/completed/post-breadth-playtest-review-001.md`
Review: `docs/roadmap/POST_BREADTH_PLAYTEST_REVIEW_2026-06-07.md`

목표:

- command-choice 7-beat route와 review motivation evidence를 읽고, deeper hardening이 실제로 필요한지 판정한다.
- P0/P1 blocker가 없으면 새 hardening 구현을 열지 않고 현 상태를 중기 목표 완료 후보로 본다.

결과:

- `playable_command_choice_scope` final/app_state evidence는 7번째 bracket scope beat를 안정적으로 설명한다.
- `playable_review_queue` evidence는 tutorial success의 review motivation과 실제 `next` action이 분리되어 있음을 보여준다.
- `make release-check`는 exit code 0으로 통과했다.
- deeper hardening은 현재 evidence가 요구하지 않는다.

제외:

- progress schema 변경
- content schema 변경
- 새 Vim command/engine capability
- 새 dependency
- release candidate/tag 작업

### 0.8. USER-DECISION-CHECKPOINT — Commit or New Direction

Status: completed, followed by `NEXT-DIRECTION-CHECKPOINT`

목표:

- 현재 완료된 중기 목표 묶음을 커밋할지, 새 방향을 열지, 나중의 release candidate 준비로 갈지 선택한다.
- 사용자 스크린샷 이후에는 modal/action hierarchy hardening을 먼저 열었고, 지금은 완료됐다.

제외:

- 사용자 요청 없는 commit
- 바로 release/tag 작업

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

- `incident-005-command-choice`는 후속 breadth까지 포함해 7 beat route이며 scope/range/inline/reuse/repeat-change/line-reuse/bracket-pair scope 판단 질문에 매핑된다.
- success/failure copy는 command 이름보다 선택 이유를 먼저 설명한다.
- 후속 후보였던 line reuse와 bracket-pair scope는 각각 `LINE-REUSE-APPLIED-001`, `COMMAND-CHOICE-BREADTH-002`로 승격했다. search-then-scope는 별도 `incident-008-search-scope`로 완료됐다.

### 5. UI-CONSOLE-POLISH-001 — Runbook Console Product Feel

Status: completed
ExecPlan: `docs/exec-plans/completed/ui-console-polish-001-action-identity.md`

목표: action line의 화면 label과 E2E action 의미를 분리한다.

완료 결과:

- `FocusPanel.actions` internal DTO와 `app_state.ui.focus_panel.actions`를 추가했다.
- 사용자 화면 label은 `다음 단계`, `다시 시도`, `다음 런북`, `출격 완료` 같은 제품 톤을 사용한다.
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

- repeat-change choice: 같은 변경을 `.`로 반복할지 판단
- line reuse choice: 검증된 줄 전체를 `V` + `y` + `p`로 재사용
- search-then-act incident: `/`, `n`, `N`으로 찾고 적절한 edit command 선택
- mixed incident 008: 3~5 beat 이하로 제한한 생존 어드벤처 run

품질 기준:

- 한 beat는 하나의 판단만 요구한다.
- 새 command를 소개하지 않는다.
- long route에는 final/timeline evidence를 남긴다.

완료 결과:

- repeat-change choice를 fifth beat로 추가했다.
- line reuse choice는 `LINE-REUSE-APPLIED-001`로 추가했다.
- search-then-scope는 `SEARCH-THEN-SCOPE-APPLIED-001`로 별도 incident에 승격했다. 이후 fresh review 결과, 새 route를 바로 늘리기보다 `REVIEW-LOOP-MOTIVATION-001`을 먼저 여는 것으로 판정했다.

### 4. QUOTE-PAIR-HARDEN-001 — Quote/Pair Text Object Hardening

Status: completed

목표: 기존 `i"` text object를 단계적으로 `i'`, `i(`, `i{`로 확장한다.

포함:

- `ci'`, `di'`, `yi'`
- `di(`, `ci(`, `yi(`, `di{`, `ci{`, `yi{`
- config/JSON/function-argument style exercise

제외:

- nested pair
- escaped delimiter
- around object
- count/register prefix

완료 결과:

- 첫 hardening scope는 `i'`로 닫았다.
- 두 번째 hardening scope는 같은 줄의 비중첩 `i(`, `i{`로 닫았다.

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

현재 release readiness, first-run polish, pre-RC hardening, modal/action hierarchy hardening은 닫혔다. 다만 공개 후보 설명과 evidence 묶기는 사용자가 release candidate 준비를 원할 때만 연다.

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
