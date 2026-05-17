# 0001 — Harness-first Replan

Status: accepted

## Context

기존 Advimture 구현과 기획은 Go + Bubble Tea 기반 Vim 학습 TUI 게임을 어느 정도 구성했지만, 제품 경험이 충분히 만족스럽지 않다. 바로 코드를 갈아엎으면 Agent가 기존 구현을 기준으로 새 작업을 정당화할 위험이 있다.

## Decision

기존 구현은 참고 자료로만 둔다. 새 작업은 `docs/roadmap/`, `docs/exec-plans/`, 도메인별 spec과 domain contract를 먼저 갱신한 뒤 진행한다.

## Consequences

- 구현 속도보다 기획과 검증 기준의 선명도를 우선한다.
- 기존 `docs/archived/PLAN.md`, `docs/archived/GAME_DESIGN.md`의 아이디어는 승격 절차를 거쳐야 재사용된다.
- TUI E2E QA Loop는 별도 ExecPlan을 만든 뒤 도입한다.
