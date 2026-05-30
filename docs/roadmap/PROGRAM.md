# Program — 현재 Phase

> 가장 자주 읽히는 파일이다. 현재 phase, active slice, 다음 권장 후보, 최근 완료만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 긴 이력은 `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/CHANGES.md`, `docs/exec-plans/completed/`, `docs/roadmap/archive/`를 본다.

## 현재 Phase

Phase: Vim Learning Foundation

목표: Vim 학습 게임의 핵심 설계 단위인 command cluster, exercise, scenario를 축적하고, 첫 출시 가능한 foundation loop의 출구 조건을 닫는다.

## 활성 슬라이스

현재 활성 구현 slice 없음. `RELEASE-READINESS-001`은 완료됐고, 다음 권장 작업은 `PLAYTEST-GATE-001`로 release candidate를 실제 fresh run 관점에서 점검하는 것이다.

Rolling plan: `docs/roadmap/FORWARD_PLAN.md`

## 다음 권장 후보

### PLAYTEST-GATE-001. Fresh Playtest Release Gate

- 상태: proposed
- 목표: README 기준으로 처음 실행하는 플레이어 관점에서 첫 5분, tutorial 확장, incident 3개 이상을 직접 훑고 출시 전 UX/content blocker를 분리한다.
- 입력 문서:
  - `README.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - `docs/verification/tui-ui-qa-contract.md`
  - `artifacts/e2e/`
- 제외:
  - 저장 포맷 변경
  - 새 Vim engine
  - 새 content schema
  - 기능 확장
- 주의: 버그는 즉시 별도 fix slice로 분리하고, 취향/세계관/콘텐츠 확장은 release 이후 후보로 나눈다.

## 최근 완료

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
