# Verification Domain Contract

> 검증 도메인은 Go 테스트, TUI E2E 루프, evidence 산출물, Agent QA 반복 방식을 다룬다.

## 공통 불변 규칙

| 규칙 | 근거 | 검증 방법 |
|------|------|----------|
| 검증은 변경 파일 범위에서 시작하고, 공통 동작 영향 시 전체 테스트로 확장한다. | 작은 변경의 피드백 속도를 유지하면서 회귀를 막는다. | 최종 보고에 실행한 명령을 기록한다. |
| TUI E2E는 실제 사용자 progress를 건드리지 않는다. | 로컬 `~/.advimture/progress.json` 손상을 막아야 한다. | E2E 러너가 임시 HOME 또는 path injection을 쓰는지 확인한다. |
| E2E 실패 evidence는 임시 산출물로 격리한다. | raw ANSI log와 screenshot류가 docs에 섞이면 잡음이 누적된다. | `artifacts/` 아래 생성하고 git stage 여부를 확인한다. |
| TUI assertion은 화면 텍스트만이 아니라 key trace와 state 변화도 함께 본다. | ANSI 렌더링만으로는 실제 게임 성공 여부를 오판할 수 있다. | E2E scenario에 input trace, screen assertion, exit/progress assertion이 있는지 확인한다. |

## Verification 고유 규칙

| 규칙 | 근거 | 검증 방법 |
|------|------|----------|
| TUI E2E 러너 도입 전에는 설계를 문서로 합의한다. | pseudo terminal, ANSI parsing, timing flake 정책을 먼저 정해야 한다. | `docs/verification/tui-e2e-loop.md`와 ExecPlan을 확인한다. |
| timing 기반 assertion을 최소화한다. | TUI 애니메이션과 터미널 렌더링은 환경별 흔들림이 있다. | scenario에서 sleep보다 상태 텍스트/프롬프트 도달 대기를 우선하는지 확인한다. |
| flake는 재시도보다 원인 evidence를 우선 남긴다. | 재시도만 추가하면 UX 문제를 숨길 수 있다. | 실패 시 cleaned screen과 raw log가 남는지 확인한다. |
