# UI-CONSOLE-POLISH-001 — Runbook Console Product Feel

Status: completed
Created: 2026-06-06

## 영향 도메인

- Gameplay UI: FocusPanel action line을 화면 label과 internal action identity로 분리한다.
- E2E QA: app_state와 assertion이 action 의미를 문자열보다 안정적으로 검증하게 한다.
- Docs: TUI screen contract, UX direction, UX backlog를 최신화한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: failed/succeeded FocusPanel은 modal로 렌더링되고 action line이 잘리지 않아야 한다.
- `docs/gameplay/tui-screen-contract.md`: screen label과 검증 가능한 UI contract.
- 사용자 제공 중기 Plan: 저장 포맷 변경 없이 internal DTO 수준 action identity를 추가한다.

## Step 1: FocusPanel Action Identity

- 목표: `Retry`, `Next`, `Next runbook`, `Dispatch complete` 등 화면 label과 E2E action marker를 분리한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playable/model.go`
  - `internal/e2estate/state.go`
  - `cmd/e2e-runner/main.go`
- 테스트 파일:
  - `internal/playableview/render_test.go`
  - `internal/playable/model_test.go`
  - `internal/e2estate/state_test.go`
  - `cmd/e2e-runner/main_test.go`
- 충족 기준:
  - 사용자에게 보이는 action label이 제품 톤과 맞는다.
  - E2E는 `action.id`로 안정적인 의미를 검증할 수 있다.
  - app_state 확장은 내부 QA DTO에만 머물고 progress 저장 포맷을 변경하지 않는다.
- Boundaries 주의:
  - `internal/progress/`, `go.mod`, `go.sum` 변경 금지.
- 상세 작업:
  - [x] `FocusPanel`에 action DTO를 추가한다.
  - [x] success/failure action 생성 로직을 label + id로 분리한다.
  - [x] renderer overflow priority를 action identity 기반으로 바꾼다.
  - [x] e2e app_state와 YAML assertion에 `actions` 검증을 추가한다.
  - [x] viewport/failure/success focused E2E를 최신화한다.

## Step 2: Console Label/Docs Polish

- 목표: 콘솔 action language와 문서 계약을 맞춘다.
- 변경 파일:
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - focused E2E YAML
- 충족 기준:
  - 80x24 성공/실패 modal에서 다음 행동이 잘리지 않는다.
  - color/style이 없어도 동일한 의미를 읽을 수 있다.
  - action label과 internal action id 계약이 문서화된다.
- 상세 작업:
  - [x] 제품 톤 action label을 문서화한다.
  - [x] E2E가 action id로 의미를 검증하도록 focused YAML을 갱신한다.
  - [x] UX backlog에서 action line 분리 항목을 완료 처리한다.

## 검증 계획

- `go test ./internal/playableview ./internal/playable ./cmd/e2e-runner`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_success_modal_80x24.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_failure_modal_80x24.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_hint_affordance.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_mismatch_feedback.yaml`
- `git diff --check`

## 완료 기록

- Completed on 2026-06-06.
- `FocusPanel.actions`와 `app_state.ui.focus_panel.actions`를 추가했다.
- `retry`, `next`, `next_tutorial`, `next_runbook`, `next_dispatch`, `dispatch_complete`, `playlist_complete`, `quit` action id를 label과 분리했다.
- `playable_viewport_success_modal_80x24`, `playable_viewport_failure_modal_80x24`, `playable_incident_hint_affordance`, `playable_command_mismatch_feedback` 통과.
- `go test ./internal/playableview ./internal/playable ./cmd/e2e-runner` 통과.
