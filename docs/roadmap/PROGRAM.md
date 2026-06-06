# Program — 현재 Phase

> 가장 자주 읽히는 파일이다. 현재 phase, active slice, 다음 권장 후보, 최근 완료만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 긴 이력은 `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/CHANGES.md`, `docs/exec-plans/completed/`, `docs/roadmap/archive/`를 본다.

## 현재 Phase

Phase: Vim Learning Foundation

목표: Vim 학습 게임의 핵심 설계 단위인 command cluster, exercise, scenario를 축적하고, 첫 출시 가능한 foundation loop의 출구 조건을 닫는다.

## 활성 슬라이스

현재 active slice는 없다. `PLAYABLE-QUALITY-BASELINE-001` 이후의 중기 보강 3단계(`CONTENT-ARC-POLISH-001`, `JUDGMENT-DRILL-REVIEW-001`, `UI-CONSOLE-POLISH-001`)도 완료됐고, 바로 출시 후보를 포장하지 않는다는 방향은 유지한다.

Completed ExecPlan: `docs/exec-plans/completed/playable-quality-baseline-001-release-quality-baseline.md`
Completed Midterm Polish:
- `docs/exec-plans/completed/content-arc-polish-001-first-dispatch-arc.md`
- `docs/exec-plans/completed/judgment-drill-review-001-command-choice.md`
- `docs/exec-plans/completed/ui-console-polish-001-action-identity.md`
Review: `docs/roadmap/PLAYABLE_QUALITY_BASELINE_2026-06-02.md`
Audit: `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md`
Evidence Bundle: `docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md`

Rolling plan: `docs/roadmap/FORWARD_PLAN.md`

## 다음 권장 후보

### POST-POLISH-PLAYTEST-001. Fresh Product Loop Review

- 상태: next candidate
- 목표: 새 구현 없이 first tour, first dispatch, judgment drill, review loop evidence를 사람이 한 번 더 읽고 다음 병목이 content breadth인지 UI polish인지 판정한다.
- 입력 문서:
  - `docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md`
  - `docs/gameplay/world-frame.md`
  - `docs/gameplay/command-choice-drills.md`
  - `docs/gameplay/tui-screen-contract.md`
- 제외:
  - 바로 출시하거나 tag를 찍는 작업
  - 신규 대형 tutorial/incident 구현
  - 사용자 확인 없는 저장 포맷/content schema/의존성 변경
- 결론: 구현보다 다음 병목 판정이 먼저다.

### CONTENT-ARC-POLISH-001. First Dispatch Arc Polish

- 상태: completed
- 목표: 새 YAML 콘텐츠 없이 incident 001~003의 기존 title/briefing/feedback을 world-frame 기준으로 spot polish한다.
- 완료: 원인 신호 추적 -> 구조 재동기화 -> 오염 구간 격리 흐름으로 첫 dispatch arc를 정리했다.
- 제외: 새 command, 새 content schema, progress 저장 포맷 변경

### JUDGMENT-DRILL-REVIEW-001. Command Choice as Core Identity

- 상태: completed
- 목표: `incident-005-command-choice`의 다섯 beat가 scope/range/inline/reuse/repeat-change 판단 질문을 설명하게 한다.
- 완료: success/failure copy와 `docs/gameplay/command-choice-drills.md`의 mapping을 선택 이유 중심으로 최신화했다.
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
- 결론: `di'`, `ci'`, `yi'` single quote 내부 object를 engine/runtime/content/E2E까지 연결했다. `i(`, `i{`는 후속 hardening으로 남긴다.

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
