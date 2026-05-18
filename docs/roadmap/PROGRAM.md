# Program — 현재 Phase

> 가장 자주 읽히는 파일이다. 현재 phase와 활성 slice만 둔다. 과거 phase는 `archive/`로 이동한다.

## 현재 Phase

Phase: Vim Learning Foundation

목표: Vim 학습 게임의 핵심 설계 단위인 command cluster, exercise, scenario를 축적하고, 첫 5분 플레이 루프의 기반을 만든다.

## 활성 슬라이스

### VIM-001. 첫 command cluster 설계
- 상태: active
- ExecPlan: `docs/exec-plans/active/vim-001-first-command-cluster.md`
- 목표: 첫 5분 플레이 루프에 들어갈 command cluster 후보를 선정하고, `docs/gameplay/command-catalog.md`의 draft 항목을 승인 가능한 수준으로 다듬는다.
- 범위: command catalog 중심. exercise/scenario는 필요한 최소 초안만 다룬다.

## 완료된 초기 세팅

### P0-001. Harness docs bootstrap
- 상태: completed
- ExecPlan: 없음. 초기 문서 세팅 작업.
- 완료일: 2026-05-17

### E2E-001. TUI E2E runner bootstrap
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/e2e-runner-bootstrap.md`
- 완료일: 2026-05-18

### E2E-002. E2E runner assertion hardening
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/e2e-002-runner-assertion-hardening.md`
- 완료일: 2026-05-18

### E2E-003. App state summary assertion
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/e2e-003-app-state-summary-assertion.md`
- 완료일: 2026-05-18

### E2E-004. Go cache isolation for E2E runner
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/e2e-004-go-cache-isolation.md`
- 완료일: 2026-05-18

### PLAY-001. First playable vertical slice
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/play-001-first-playable-slice.md`
- 완료일: 2026-05-18

### VIM-002. Vim engine foundation
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-002-vimengine-foundation.md`
- 완료일: 2026-05-18

### VIM-003. Vim engine contract hardening
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-003-vimengine-contract-hardening.md`
- 완료일: 2026-05-18

### VIM-004. Exercise runtime foundation
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-004-exercise-runtime-foundation.md`
- 완료일: 2026-05-18

### VIM-005. Content schema foundation
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-005-content-schema-foundation.md`
- 완료일: 2026-05-18

### VIM-006. Grader / scoring engine
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-006-grader-scoring-engine.md`
- 완료일: 2026-05-18

### VIM-007. Scenario runtime
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-007-scenario-runtime.md`
- 완료일: 2026-05-18

### VIM-008. TUI adapter
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-008-tui-adapter.md`
- 완료일: 2026-05-18

### VIM-009. Progress adapter
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-009-progress-adapter.md`
- 완료일: 2026-05-18

### VIM-010. Neovim oracle runner
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-010-neovim-oracle-runner.md`
- 완료일: 2026-05-18

### VIM-011. Legacy archive pass
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/vim-011-legacy-archive-pass.md`
- 완료일: 2026-05-18

### LEGACY-001. Archive old implementation
- 상태: completed
- ExecPlan: `docs/exec-plans/completed/legacy-001-archive-old-implementation.md`
- 완료일: 2026-05-18

## Backlog

| ID | 제목 | 우선순위 |
|----|------|---------|
| B-001 | 첫 command cluster 승인 | P0 |
| B-002 | 첫 exercise set 승인 | P0 |
| B-003 | 첫 scenario skin 승인 | P1 |
| B-004 | 첫 playable slice에서 app state summary export 연결 | P0 |
| B-005 | 기존 구현에서 유지할 모듈/버릴 모듈 분류 | P1 |
| B-006 | Neovim oracle test runner 설계 | P1 |
