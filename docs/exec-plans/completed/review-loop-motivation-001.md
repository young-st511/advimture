# REVIEW-LOOP-MOTIVATION-001 — Review Loop Motivation Polish

Slice-ID: REVIEW-LOOP-MOTIVATION-001
Created: 2026-06-07
Status: completed
Scope-Mode: normal

## 영향 도메인

- Gameplay: success debrief의 review motivation wording을 다듬는다. progress 저장 포맷과 review queue 계산은 유지한다.
- Verification: focused model tests와 E2E가 화면 문구, app_state action id, review state를 함께 검증한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: FocusPanel success modal, review queue, action id 계약.
- `docs/gameplay/tui-screen-contract.md`: review/daily는 보조 정보이며 action label은 `actions.id`와 분리한다.
- `docs/roadmap/NEXT_PLAYTEST_REVIEW_2026-06-07.md`: 다음 후보를 `REVIEW-LOOP-MOTIVATION-001`로 판정했다.

## 목표

success modal에서 primary action과 review motivation을 더 명확히 분리한다. Tutorial success는 실제 버튼처럼 보이는 `다음 출격` 문구를 피하고, incident success는 Runbook Dispatch 톤을 유지하되 후보성 문구로 표현한다.

## 포함

- tutorial success debrief의 review line copy 변경
- incident success debrief의 review line copy 변경
- app_state `ui.focus_panel.actions[].id` 계약 유지
- focused E2E expectation 갱신
- gameplay/roadmap/UX docs 동기화

## 제외

- progress schema 변경
- content schema 변경
- review queue 후보 계산 변경
- 새 Vim command/engine capability
- 새 dependency
- release candidate/tag 작업

## Step 1: Spec 기반 테스트

- 목표: 성공 modal review motivation copy를 테스트로 먼저 고정한다.
- 변경 파일:
  - `internal/playable/model_test.go`
  - `test/e2e/playable_review_queue.yaml`
  - `test/e2e/playable_debrief_success.yaml`
- 충족 기준: tutorial success에서는 `나중에 다시 풀기`를 표시하고, incident success에서는 `다음 출격 후보`를 표시한다. action id/label은 기존 계약을 유지한다.
- Boundaries 주의: progress 저장 포맷 변경 금지.
- 상세 작업:
  - [x] tutorial success debrief unit expectation을 추가/수정한다.
  - [x] incident success debrief unit expectation을 추가한다.
  - [x] focused E2E expected screen text를 새 copy로 갱신한다.

## Step 2: 구현

- 목표: review wording을 current playlist category에 따라 분리한다.
- 변경 파일:
  - `internal/playable/model.go`
- 테스트 파일:
  - `internal/playable/model_test.go`
- 충족 기준: review queue와 action id는 그대로이고, success panel lines만 제품 톤에 맞게 바뀐다.
- Boundaries 주의: `internal/progress/` 접근 금지.
- 상세 작업:
  - [x] tutorial success residual line을 `재점검 메모`로 표현한다.
  - [x] tutorial success next review line을 `나중에 다시 풀기`로 표현한다.
  - [x] incident success next review line을 `다음 출격 후보`로 표현한다.
  - [x] 마지막 dispatch의 `next_dispatch` action label은 유지한다.

## Step 3: 문서/검증/Closeout

- 목표: 현재 동작 문서와 roadmap을 업데이트하고 검증한다.
- 변경 파일:
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/CHANGES.md`
  - `docs/exec-plans/active/review-loop-motivation-001.md`
  - `docs/exec-plans/completed/review-loop-motivation-001.md`
- 충족 기준: focused Go tests, focused E2E, full Go tests, diff check가 통과한다.
- Boundaries 주의: release/RC/tag 작업은 하지 않는다.
- 상세 작업:
  - [x] 현재 동작 문서를 갱신한다.
  - [x] 검증 명령을 실행한다.
  - [x] ExecPlan을 completed로 이동한다.

## 완료 결과

- Tutorial success debrief는 `재점검 메모`와 `나중에 다시 풀기`를 표시한다.
- Incident success debrief는 `잔류 리스크`와 `다음 출격 후보`를 표시한다.
- `next`, `next_dispatch` 등 `app_state.ui.focus_panel.actions[].id`와 화면 action label은 유지했다.
- progress schema, content schema, dependency, engine capability는 변경하지 않았다.

## 검증 계획

- `go test ./internal/playable ./internal/playableview ./internal/e2estate`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_debrief_success.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_success_modal_80x24.yaml`
- `go test ./...`
- `git diff --check`

## 실행 규칙 (하네스 모드)

각 Step은 Spec 기반 테스트 작성 -> 구현 -> 검증 -> 문서 동기화 순서로 수행한다. 저장 포맷, content schema, dependency 변경은 금지한다.
