# Program — 현재 Phase

> 가장 자주 읽히는 파일이다. 현재 phase, active slice, 다음 권장 후보, 최근 완료만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 긴 이력은 `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/CHANGES.md`, `docs/exec-plans/completed/`, `docs/roadmap/archive/`를 본다.

## 현재 Phase

Phase: Vim Learning Foundation

목표: Vim 학습 게임의 핵심 설계 단위인 command cluster, exercise, scenario를 축적하고, 첫 출시 가능한 foundation loop의 출구 조건을 닫는다.

## 활성 슬라이스

현재 활성 구현 slice 없음. `PLAN-REFRESH-009`의 forward plan 문서화는 완료됐고, 다음 실제 작업은 `FOUNDATION-EXIT-001`로 foundation exit review를 수행하는 것이다.

Rolling plan: `docs/roadmap/FORWARD_PLAN.md`

## 다음 권장 후보

### FOUNDATION-EXIT-001. Foundation Exit Review

- 상태: proposed
- 목표: 현재 engine/content/UI/E2E가 foundation 출시 후보로 충분한지 점검하고 다음 중기 플랜을 고른다.
- 입력 문서:
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/PLAYTEST_REVIEW_2026-05-29.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/gameplay/spec.md`
  - `docs/gameplay/vim-curriculum-map.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
- 선택 후보:
  - `QUOTE-PAIR-HARDEN-001`: `ci'`, `ci(`, `ci{` 등 quote/pair text object 확장
  - `PLATFORM-REVIEW-003`: 저장 포맷 변경 없는 mission/review/game loop 강화
  - content breadth: 기존 engine만 사용하는 applied incident/command-choice 확장
- 주의: evidence 없이 새 engine/content를 열지 않는다.

## 최근 완료

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
