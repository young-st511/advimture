# Program — 현재 Phase

> 가장 자주 읽히는 파일이다. 현재 phase, active slice, 다음 권장 후보, 최근 완료만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 긴 이력은 `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/CHANGES.md`, `docs/exec-plans/completed/`, `docs/roadmap/archive/`를 본다.

## 현재 Phase

Phase: Vim Learning Foundation

목표: Vim 학습 게임의 핵심 설계 단위인 command cluster, exercise, scenario를 축적하고, 첫 출시 가능한 foundation loop의 출구 조건을 닫는다.

## 활성 슬라이스

현재 active slice는 없다. `PLAYABLE-QUALITY-BASELINE-001` 이후의 중기 보강과 applied content 보강(`POST-POLISH-PLAYTEST-001`, `LINE-REUSE-APPLIED-001`, `SEARCH-THEN-SCOPE-APPLIED-001`, `BRACKET-PAIR-HARDEN-001`)은 완료됐고, `NEXT-PLAYTEST-REVIEW-001` 이후의 `REVIEW-LOOP-MOTIVATION-001`, `COMMAND-CHOICE-BREADTH-002`, `POST-BREADTH-PLAYTEST-REVIEW-001`도 완료됐다. 다음은 구현 slice가 아니라 user decision checkpoint다. 바로 출시 후보를 포장하지 않는다는 방향은 유지한다.

Completed ExecPlan: `docs/exec-plans/completed/playable-quality-baseline-001-release-quality-baseline.md`
Completed Midterm Polish:
- `docs/exec-plans/completed/content-arc-polish-001-first-dispatch-arc.md`
- `docs/exec-plans/completed/judgment-drill-review-001-command-choice.md`
- `docs/exec-plans/completed/ui-console-polish-001-action-identity.md`
- `docs/exec-plans/completed/post-polish-playtest-001-fresh-product-loop-review.md`
- `docs/exec-plans/completed/line-reuse-applied-001-line-reuse-drill.md`
- `docs/exec-plans/completed/search-then-scope-applied-001.md`
- `docs/exec-plans/completed/bracket-pair-harden-001.md`
- `docs/exec-plans/completed/next-playtest-review-001-fresh-direction.md`
- `docs/exec-plans/completed/review-loop-motivation-001.md`
- `docs/exec-plans/completed/command-choice-breadth-002.md`
- `docs/exec-plans/completed/post-breadth-playtest-review-001.md`
Review: `docs/roadmap/PLAYABLE_QUALITY_BASELINE_2026-06-02.md`
Audit: `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md`
Evidence Bundle: `docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md`
Latest Review: `docs/roadmap/POST_BREADTH_PLAYTEST_REVIEW_2026-06-07.md`

Rolling plan: `docs/roadmap/FORWARD_PLAN.md`

## 다음 권장 후보

### USER-DECISION-CHECKPOINT. Commit or New Direction

- 상태: recommended decision
- 목표: 현재 완료된 중기 목표 묶음을 커밋할지, 새 방향을 열지, 나중의 release candidate 준비로 갈지 결정한다.
- 근거:
  - `docs/roadmap/POST_BREADTH_PLAYTEST_REVIEW_2026-06-07.md`
- 제외:
  - 사용자 요청 없는 commit
  - 바로 release/tag 작업
- 결론: evidence상 추가 hardening은 지금 열 필요가 없다.

### POST-BREADTH-PLAYTEST-REVIEW-001. Post Breadth Playtest Review

- 상태: completed
- 목표: command-choice 7-beat route와 최근 review motivation polish evidence를 읽고, deeper hardening이 실제로 필요한지 판정한다.
- 완료: P0/P1 blocker는 보이지 않으며, deeper hardening은 현재 evidence가 요구하지 않는다고 판정했다.
- Review: `docs/roadmap/POST_BREADTH_PLAYTEST_REVIEW_2026-06-07.md`
- ExecPlan: `docs/exec-plans/completed/post-breadth-playtest-review-001.md`

### COMMAND-CHOICE-BREADTH-002. Command Choice Breadth

- 상태: completed
- 목표: command-choice를 더 넓은 상황 판단 훈련으로 확장하되, 이미 구현된 engine command만 사용한다.
- 완료: `incident-005-command-choice`에 `command-choice-bracket-scope-001`을 seventh beat로 추가했다. hyphenated 괄호 인자 전체를 `ci(`로 교체하는 scope 판단을 검증한다.
- ExecPlan: `docs/exec-plans/completed/command-choice-breadth-002.md`

### REVIEW-LOOP-MOTIVATION-001. Review Loop Motivation Polish

- 상태: completed
- 목표: success debrief, 잔류 리스크, 오늘의 복구 루트, 다음 출격 언어를 다듬어 반복 플레이 동기를 더 자연스럽게 만든다.
- 완료: tutorial success는 `재점검 메모`/`나중에 다시 풀기`, incident success는 `잔류 리스크`/`다음 출격 후보`를 표시한다. action id/label 계약과 progress 저장 포맷은 유지했다.
- ExecPlan: `docs/exec-plans/completed/review-loop-motivation-001.md`

### NEXT-PLAYTEST-REVIEW-001. Fresh Direction Review

- 상태: completed
- 목표: `SEARCH-THEN-SCOPE-APPLIED-001`과 `BRACKET-PAIR-HARDEN-001` 완료 후 fresh evidence를 보고 다음 보강 축을 고른다.
- 완료: 다음 후보를 `REVIEW-LOOP-MOTIVATION-001`로 판정했다.
- Review: `docs/roadmap/NEXT_PLAYTEST_REVIEW_2026-06-07.md`
- ExecPlan: `docs/exec-plans/completed/next-playtest-review-001-fresh-direction.md`

### BRACKET-PAIR-HARDEN-001. Parenthesis/Brace Inner Text Object

- 상태: completed
- 목표: `text-object-quote-pair`를 같은 줄의 비중첩 parenthesis/brace 내부 object `i(`, `i{`까지 확장한다.
- 완료: `tutorial-95-bracket-pair` 6문항과 `playable_bracket_pair_full` E2E를 추가했다.
- ExecPlan: `docs/exec-plans/completed/bracket-pair-harden-001.md`
- 제외:
  - 바로 출시하거나 tag를 찍는 작업
  - nested pair, escaped delimiter, around object, multi-line pair
  - 사용자 확인 없는 저장 포맷/content schema/의존성 변경

### SEARCH-THEN-SCOPE-APPLIED-001. Search Then Scope Drill

- 상태: completed
- 목표: marker를 먼저 찾고 그 주변 줄 묶음을 판단해 처리하는 applied run을 추가한다.
- 완료: `incident-008-search-scope`를 별도 1-beat incident로 추가했다.
- 제외: 새 command, 새 schema, progress 저장 포맷 변경

### POST-POLISH-PLAYTEST-001. Fresh Product Loop Review

- 상태: completed
- 목표: 새 구현 없이 first tour, first dispatch, judgment drill, review loop evidence를 사람이 한 번 더 읽고 다음 병목이 content breadth인지 UI polish인지 판정한다.
- 완료: 다음 slice를 `LINE-REUSE-APPLIED-001`로 판정했다.
- Review: `docs/roadmap/POST_POLISH_PLAYTEST_2026-06-06.md`

### LINE-REUSE-APPLIED-001. Line Reuse Choice Drill

- 상태: completed
- 목표: `incident-005-command-choice`에 검증된 줄 전체를 직접 재입력하지 않고 linewise yank/put으로 재사용하는 sixth beat를 추가한다.
- 완료: `command-choice-line-reuse-001`을 추가하고 focused command-choice E2E를 6-beat route로 갱신했다.
- 제외: 새 command, 새 schema, progress 저장 포맷 변경

### CONTENT-ARC-POLISH-001. First Dispatch Arc Polish

- 상태: completed
- 목표: 새 YAML 콘텐츠 없이 incident 001~003의 기존 title/briefing/feedback을 world-frame 기준으로 spot polish한다.
- 완료: 원인 신호 추적 -> 구조 재동기화 -> 오염 구간 격리 흐름으로 첫 dispatch arc를 정리했다.
- 제외: 새 command, 새 content schema, progress 저장 포맷 변경

### JUDGMENT-DRILL-REVIEW-001. Command Choice as Core Identity

- 상태: completed
- 목표: `incident-005-command-choice`의 기존 beat가 scope/range/inline/reuse/repeat-change 판단 질문을 설명하게 한다.
- 완료: success/failure copy와 `docs/gameplay/command-choice-drills.md`의 mapping을 선택 이유 중심으로 최신화했다. 이후 `LINE-REUSE-APPLIED-001`와 `COMMAND-CHOICE-BREADTH-002`가 추가되어 현재 route는 일곱 beat다.
- 제외: 새 command, 새 schema, 새 exercise

### UI-CONSOLE-POLISH-001. Runbook Console Product Feel

- 상태: completed
- 목표: success/failure action line의 화면 label과 E2E action 의미를 분리한다.
- 완료: `FocusPanel.actions` internal DTO와 app_state assertion을 추가하고, 화면 label을 한국어 action language로 정리했다.
- 제외: progress 저장 포맷 변경, 새 dependency

### RELEASE-CANDIDATE-001. Release Candidate Prep

- 상태: later
- 목표: 실제로 공개 후보를 묶기로 결정했을 때 release note, known limitations, final evidence bundle, tag 후보 기준을 정리한다.
- 조건: 사용자가 실제 release candidate 준비를 원할 때 연다.
- 제외: 새 engine, 새 schema, progress 저장 포맷 변경

## 최근 완료

### PLAYABLE-QUALITY-BASELINE-001. Release-Quality Playable Baseline

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/playable-quality-baseline-001-release-quality-baseline.md`
- Review: `docs/roadmap/PLAYABLE_QUALITY_BASELINE_2026-06-02.md`
- Audit: `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md`
- 완료일: 2026-06-02
- 결론: 세계관/UX/UI/콘텐츠 기획/엔진 모듈화 기준을 release-quality baseline으로 정렬했고, mode/floating modal/header track UX를 보강했다. `make release-check` 통과.

### First Dispatch/Judgment/UI Console Midterm Polish

- 상태: completed
- ExecPlans:
  - `docs/exec-plans/completed/content-arc-polish-001-first-dispatch-arc.md`
  - `docs/exec-plans/completed/judgment-drill-review-001-command-choice.md`
  - `docs/exec-plans/completed/ui-console-polish-001-action-identity.md`
- 완료일: 2026-06-06
- 결론: 새 engine/schema/progress 저장 포맷 없이 첫 dispatch arc, judgment drill identity, action label/app_state 계약을 보강했다.

### PRE-RC-HARDENING-001. First Release Hardening

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/pre-rc-hardening-001-first-release-hardening.md`
- Review: `docs/roadmap/PRE_RC_HARDENING_2026-06-02.md`
- 완료일: 2026-06-02
- 결론: P0/P1 blocker 없음. Incident hint cue wrapping과 long incident app_state evidence gap을 보강했다. 당시 다음 후보는 `RELEASE-CANDIDATE-001`이었으나, 현재 active 방향은 `PLAYABLE-QUALITY-BASELINE-001`이다.

### FIRST-RUN-POLISH-001. Cue Density and Viewport Evidence

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/first-run-polish-001-cue-viewport.md`
- 완료일: 2026-06-01
- 결론: tutorial running cue의 `기억할 명령`/`Coach` 중복을 제거했고, floating modal 주변 detailed review/daily line을 숨겼으며, 80x24 success/failure viewport smoke와 mid tutorial final/timeline/app_state evidence를 보강했다.

### PLAYTEST-GATE-001. Fresh Playtest Release Gate

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/playtest-gate-001-fresh-release-gate.md`
- Review: `docs/roadmap/PLAYTEST_GATE_2026-05-30.md`
- 완료일: 2026-05-30
- 결론: P0/P1 blocker 없음. P2 first-run polish 후보는 running cue density, review/daily line length, viewport smoke evidence 등이다.

### RELEASE-READINESS-001. First Release Readiness

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/release-readiness-001-first-release.md`
- 완료일: 2026-05-30
- 결론: README를 실행 가능한 게임 안내로 갱신하고 `make build`, `make test`, `make release-check`를 추가했다. release gate는 `make test`, `make build`, `make e2e-playable`을 묶는다.

### UI-POLISH-002. Command Memory Cue

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/ui-polish-002-command-memory.md`
- 완료일: 2026-05-30
- 결론: tutorial은 `기억할 명령`, incident는 hint/failure 후 `참고 명령`으로 command memory를 점진 공개한다. 저장/schema 변경은 없다.

### QUOTE-PAIR-HARDEN-001. Single Quote Text Object

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/quote-pair-harden-001-single-quote.md`
- 완료일: 2026-05-30
- 결론: `di'`, `ci'`, `yi'` single quote 내부 object를 engine/runtime/content/E2E까지 연결했다. 이후 `BRACKET-PAIR-HARDEN-001`에서 같은 줄의 비중첩 `i(`, `i{` scope도 완료했다.

### CONTENT-BREADTH-002. Repeat Change Choice

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/content-breadth-002-repeat-choice.md`
- 완료일: 2026-05-30
- 결론: `incident-005-command-choice`에 `command-choice-repeat-change-001` fifth beat를 추가했다. 같은 단어 교체가 이어질 때 `.`으로 마지막 변경을 반복하는 판단을 검증한다.

### PLATFORM-REVIEW-003. Mission/Review Game Loop

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/platform-review-003-mission-review-loop.md`
- 완료일: 2026-05-30
- 결론: 성공 debrief와 마지막 dispatch를 review queue로 연결했다. 저장 포맷 변경은 없다.

### FOUNDATION-EXIT-001. Foundation Exit Review

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/foundation-exit-001-review.md`
- Review: `docs/roadmap/FOUNDATION_EXIT_REVIEW_2026-05-30.md`
- 완료일: 2026-05-30
- 결론: Foundation은 조건부 통과. 다음 작업은 새 engine보다 `PLATFORM-REVIEW-003`을 권장한다.

### PLAN-REFRESH-009. Foundation Forward Plan

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/plan-refresh-009-foundation-forward-plan.md`
- Plan: `docs/roadmap/FORWARD_PLAN.md`
- 완료일: 2026-05-30

### E2E-EVIDENCE-008. Long Incident Final/Timeline Evidence

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/e2e-evidence-008-long-incident-evidence.md`
- Review: `docs/roadmap/PLAYTEST_REVIEW_2026-05-29.md`
- 완료일: 2026-05-29

### PLAYTEST-REVIEW-002. Applied Mastery Review

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/playtest-review-002-applied-mastery.md`
- 완료일: 2026-05-28

### INCIDENT-007. Mixed Recovery Run

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/incident-007-mixed-recovery.md`
- 완료일: 2026-05-28

### CHOICE-PLAY-003. Quote Value Reuse Choice

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/choice-play-003-reuse-choice.md`
- 완료일: 2026-05-28

### INCIDENT-006. Inline Target Repair Run

- 상태: completed
- ExecPlan: `docs/exec-plans/completed/incident-006-inline-target-repair.md`
- 완료일: 2026-05-28

### PLAYPACK-012 / VIM-030. Char Find Line

- 상태: completed
- ExecPlans:
  - `docs/exec-plans/completed/vim-030-char-find-engine.md`
  - `docs/exec-plans/completed/playpack-012-char-find-line.md`
- 완료일: 2026-05-26

## 문서 신선도 규칙

- 이 파일에는 최근 완료 5~10개만 유지한다.
- 과거 health check/review는 `docs/roadmap/archive/`로 이동한다.
- 새 active slice가 생기면 이 파일의 `활성 슬라이스`를 먼저 갱신한다.
- 다음 후보가 바뀌면 `docs/roadmap/CHANGES.md`에 append-only로 이유를 남긴다.
- 2~8주 방향이 바뀌면 `docs/roadmap/FORWARD_PLAN.md`를 함께 갱신한다.
