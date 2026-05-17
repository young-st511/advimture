# Advimture

> Go + Bubble Tea 기반 TUI adventure 게임. 현재는 기존 Vim 학습 게임 구현을 참고 자료로 두고, 제품 기획과 검증 워크플로우를 다시 세팅하는 단계입니다.

## 개발 워크플로우

이 프로젝트는 **plan-first + 승인된 spec-first** 워크플로우를 사용합니다. 구현 전에 기획 의도, 수용 기준, 검증 루프를 문서로 먼저 고정해서 Agent가 코드와 테스트를 동시에 자기합리화하지 않도록 합니다.

### Quick Start

1. `docs/roadmap/PRODUCT.md`에 제품의 본질과 기둥을 정리합니다.
2. `docs/roadmap/PROGRAM.md`에 현재 phase와 활성 slice를 둡니다.
3. 비사소한 작업은 `docs/exec-plans/templates/`에서 템플릿을 골라 `docs/exec-plans/active/`에 만듭니다.
4. 구현 기준은 관련 `docs/{domain}/spec.md`의 `[draft]` 수용 기준으로 먼저 작성합니다.
5. 사람이 승인해 `[draft]`가 제거된 뒤 Agent가 테스트와 구현을 시작합니다.
6. TUI 전체 흐름은 `docs/verification/tui-e2e-loop.md`의 QA Loop 설계를 기준으로 검증합니다.

### E2E Smoke 실행

```sh
make e2e-smoke
```

실패/성공 evidence는 `artifacts/e2e/`에 남습니다. 이 디렉토리는 git에서 제외됩니다.

### 문서 구조

| 파일 | 누가 읽나 | 역할 |
|------|----------|------|
| `README.md` | 사람 | 워크플로우 온보딩 |
| `AGENTS.md` | AI Agent | 작업 지침 |
| `docs/README.md` | 사람 + AI | docs 구조와 작성 규칙 |
| `docs/roadmap/` | 사람 + AI | 재기획의 제품/phase/결정 기록 |
| `docs/exec-plans/` | 사람 + AI | 비사소한 작업의 실행 계획 |
| `docs/gameplay/` | 사람 + AI | 게임플레이 도메인 계약과 spec |
| `docs/verification/` | 사람 + AI | TUI E2E 검증 루프와 QA contract |
| `docs/archived/` | 사람 + AI | 과거 구현/기획 스냅샷 |

기존 `docs/archived/PLAN.md`, `docs/archived/GAME_DESIGN.md`는 과거 구현/기획 참고 자료입니다. 새 방향으로 유지할 아이디어만 `docs/`로 옮긴 뒤 canonical 기준으로 사용합니다.
