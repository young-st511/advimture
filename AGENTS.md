# AGENTS.md

## Project Overview
- Advimture는 Go + Bubble Tea 기반 TUI adventure 게임으로 재기획 중인 프로젝트다.
- 현재 구현은 참고 자료이며, 새 제품 방향의 canonical source는 `docs/roadmap/`과 `docs/*/spec.md`다.
- 이 저장소는 단일 Go 모듈이다. 웹 프런트엔드나 서비스 모노레포가 아니다.

## Commands
- Run: `go run .`
- Build: `go build .`
- Test all: `go test ./...`
- Test editor: `go test ./internal/editor/...`
- Test game: `go test ./internal/game/...`
- Test data: `go test ./internal/data/...`
- Test progress: `go test ./internal/progress/...`
- E2E smoke: `go run ./cmd/e2e-runner --scenario test/e2e/ftue_ctrl_c_quit.yaml` 또는 `make e2e-smoke`
- Makefile: `make run`, `make e2e-smoke`
- Lint: 전용 린터 설정 없음. 필요 시 `go vet ./...`를 수동 실행한다.

## Project Structure
- 현재 구조는 작업 시작 시 `rg --files`로 직접 확인할 것. 정적 트리를 외워서 판단하지 말 것.
- 새 Go 앱 코드는 `internal/` 아래 책임별 패키지에 둔다.
- TUI 화면 전환은 `internal/app/`, 화면 컴포넌트는 `internal/ui/`, 게임 상태/미션 진행은 `internal/game/`에 둔다.
- Vim 유사 편집 동작은 `internal/editor/`에 둔다.
- 튜토리얼/미션 데이터와 검증 로직은 `internal/data/`에 둔다.
- 진행 저장은 `internal/progress/`에 둔다.
- 테스트는 같은 패키지의 `*_test.go`로 둔다.
- 기획, 실행 계획, 검증 루프 문서는 `docs/` 아래에 둔다.

## Code Style
- Go 표준 포맷을 따를 것. 파일 수정 후 Go 파일은 `gofmt` 대상이다.
- 에러 메시지는 기존 코드처럼 한국어 문맥을 유지하고, 원인 보존이 필요하면 `fmt.Errorf("...: %w", err)` 형태로 래핑할 것.
- Bubble Tea 모델은 `Model`, `Update`, `View`, `Init` 패턴을 따른다.
- 사용자 노출 문구는 한국어 톤을 우선한다.
- 새 의존성은 ExecPlan 또는 명시적 사용자 승인 없이 추가하지 말 것.

## Testing
- Go 단위 테스트는 `go test ./...`를 기본 검증으로 사용한다.
- 변경 범위가 좁으면 먼저 변경 패키지 테스트를 실행하고, 공통 동작에 영향이 있으면 전체 테스트까지 실행한다.
- 새 기능은 승인된 spec 또는 ExecPlan의 수용 기준을 테스트로 먼저 옮긴 뒤 구현한다.
- TUI E2E 검증은 `make e2e-smoke`와 `docs/verification/tui-e2e-loop.md`의 루프 설계를 따른다. 새 시나리오를 추가하면 `test/e2e/`에 YAML을 두고 evidence를 `artifacts/e2e/`에서 확인한다.

## Git Workflow
- 브랜치는 기본적으로 `codex/<short-topic>` 형식을 사용한다.
- 커밋은 사용자가 요청했을 때만 만든다.
- PR에는 spec/ExecPlan 변경과 코드 변경을 함께 포함한다.
- 리뷰어는 `docs/` diff를 먼저 확인한 뒤 코드를 리뷰한다.

## Work Start Protocol
작업을 시작하기 전에 반드시 다음 순서로 확인할 것:

1. **AGENTS.md → Boundaries**를 확인한다.
2. **`docs/roadmap/PRODUCT.md`**에서 제품 기둥과 현재 재기획 가정을 확인한다.
3. **`docs/roadmap/PROGRAM.md`**에서 현재 phase와 활성 slice를 확인한다.
4. Plan-first 대상이면 **`docs/exec-plans/active/`**에 해당 ExecPlan이 있는지 확인한다.
5. 관련 도메인의 **`docs/{domain}/domain-contract.md`**를 읽는다.
6. 관련 도메인의 **`docs/{domain}/spec.md`**에서 승인된 수용 기준을 확인한다.
7. 수용 기준이 없거나 `[draft]` 상태면 구현을 시작하지 말고, 초안 작성 또는 사용자 승인을 먼저 요청한다.

## Implementation Cycle
새 기능/변경은 아래 순서를 따른다:

1. **의도 수신**: 사람의 의도나 활성 ExecPlan을 확인한다.
2. **수용 기준 초안 작성**: `spec.md` 또는 ExecPlan에 `[draft]` 항목으로 작성한다.
3. **사람 승인 대기**: 사람이 `[draft]`를 제거하거나 명시적으로 승인할 때까지 구현하지 않는다.
4. **테스트 작성 (Red)**: 승인된 기준을 테스트로 옮긴다.
5. **구현 (Green)**: 테스트가 통과하도록 최소 범위로 구현한다.
6. **검증**: 관련 Go 테스트, 필요 시 TUI E2E 루프를 실행한다.
7. **문서 동기화**: 구현된 기준을 현재 동작으로 옮기고 소스 참조를 추가한다.

## Documentation Maintenance
- `docs/README.md`에 정의된 문서별 업데이트 시점을 따른다.
- 새 기획은 `docs/roadmap/PRODUCT.md`와 `docs/roadmap/PROGRAM.md`를 먼저 갱신한다.
- 비사소한 작업은 `docs/exec-plans/active/`에 계획을 둔 뒤 진행한다.
- 기존 `docs/archived/PLAN.md`, `docs/archived/GAME_DESIGN.md`는 과거 참고 자료다. 새 구현 기준으로 인용하지 말고, 필요한 아이디어만 새 docs로 승격한다.
- 프로덕션 버그나 반복된 리뷰 지적은 관련 `domain-contract.md` 또는 `docs/guardrails.md`로 승격할지 검토한다.

### Feedback Loop
- 놓친 엣지 케이스 발견 시 관련 spec의 수용 기준에 추가한다.
- TUI에서 사람이 발견한 어색함은 `docs/verification/tui-e2e-loop.md`의 검사 항목으로 승격한다.
- phase 종료 시 `docs/roadmap/CHANGES.md`에 가정 변경을 append-only로 기록한다.

## Plan-First Rule
다음 작업은 코드 변경 전에 `docs/exec-plans/active/`에 ExecPlan을 만든다:
- 다중 파일 변경
- 게임 루프, 화면 전환, 저장 포맷, 미션 데이터 스키마 변경
- Bubble Tea 프로그램 구조 변경
- TUI E2E 러너 또는 QA 자동화 변경
- 새 의존성 추가
- 공개 contract나 roadmap에 영향을 주는 변경

다음은 ExecPlan을 건너뛸 수 있다:
- 단순 typo 수정
- 명백한 단일 파일 버그
- 테스트 전용 변경
- 문서 표현만 다듬는 변경

ExecPlan은 작성 → `active/` 진행 → 완료 시 `completed/`로 이동한다.

## Boundaries
> 이 섹션은 "무엇을 건드리지 말 것"의 quick reference다. 근거와 검증 방법은 각 domain contract를 참조할 것.

- `internal/progress/` 저장 포맷 변경 — 사용자 승인 필수 (→ `docs/gameplay/domain-contract.md`)
- `go.mod`, `go.sum` 의존성 변경 — ExecPlan과 사용자 승인 필수 (→ `docs/guardrails.md`)
- `internal/data/tutorials/`, `internal/data/missions/` YAML ID/파일명 변경 — 기획 승인 필수 (→ `docs/gameplay/domain-contract.md`)
- TUI E2E 러너가 실제 `~/.advimture`를 쓰도록 만드는 변경 금지. 테스트 전용 HOME 또는 progress path 주입을 사용할 것 (→ `docs/verification/domain-contract.md`)
- 기존 구현을 canonical spec처럼 문서화하는 작업 금지. 재기획 승인 전까지 현재 코드는 참고 자료로만 취급할 것.

## Reporting
최종 보고는 항상 `변경 내용 / 이유 / 검증 결과` 순서로 간결하게 작성한다.
