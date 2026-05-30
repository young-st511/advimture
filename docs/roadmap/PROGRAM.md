# Program — 현재 Phase

> 가장 자주 읽히는 파일이다. 현재 phase, active slice, 다음 권장 후보, 최근 완료만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 긴 이력은 `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/CHANGES.md`, `docs/exec-plans/completed/`, `docs/roadmap/archive/`를 본다.

## 현재 Phase

Phase: Vim Learning Foundation

목표: Vim 학습 게임의 핵심 설계 단위인 command cluster, exercise, scenario를 축적하고, 첫 출시 가능한 foundation loop의 출구 조건을 닫는다.

## 활성 슬라이스

현재 활성 구현 slice 없음. `PLAYTEST-GATE-001`은 P0/P1 blocker 없이 완료됐고, 다음 권장 작업은 `FIRST-RUN-POLISH-001`로 첫 실행 cue와 viewport smoke를 좁게 다듬는 것이다.

Rolling plan: `docs/roadmap/FORWARD_PLAN.md`

## 다음 권장 후보

### FIRST-RUN-POLISH-001. First Run Polish

- 상태: proposed
- 목표: 새 engine/content/schema 없이 첫 tutorial cue 밀도, review/daily line 길이, viewport smoke evidence를 좁게 다듬는다.
- 입력 문서:
  - `docs/roadmap/PLAYTEST_GATE_2026-05-30.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - `docs/verification/tui-ui-qa-contract.md`
  - `docs/gameplay/tui-screen-contract.md`
- 제외:
  - 저장 포맷 변경
  - 새 Vim engine
  - 새 content schema
  - 기능 확장
- 주의: 화면 문구 변경은 app_state/focus_panel assertion을 약화하지 않는다.

### RC-BLOCKER-FIX-001. Release Candidate Blocker Fix

- 상태: skipped for now
- 목표: `PLAYTEST-GATE-001`에서 P0/P1 blocker가 나오면 새 기능을 멈추고 해당 blocker만 고친다.
- 조건: fresh run review에서 진행 중단, 저장 손상, 오해를 부르는 필수 조작, terminal clipping 중 하나가 확인될 때 다시 연다.
- 제외: 새 engine, 새 schema, progress 저장 포맷 변경

## 최근 완료

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
