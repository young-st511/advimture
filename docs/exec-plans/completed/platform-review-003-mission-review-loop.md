# PLATFORM-REVIEW-003 — Mission/Review Game Loop

Slice-ID: PLATFORM-REVIEW-003
Created: 2026-05-30
Completed: 2026-05-30
Status: completed
Scope-Mode: gameplay-system
Allowed-Paths:
- docs/exec-plans/active/platform-review-003-mission-review-loop.md
- docs/exec-plans/completed/platform-review-003-mission-review-loop.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/UX_BACKLOG_001.md
- docs/roadmap/PLATFORM_RFC_001.md
- docs/roadmap/CHANGES.md
- internal/playable/model.go
- internal/playable/model_test.go
- internal/playableview/render.go
- internal/playableview/render_test.go
- test/e2e/playable_review_queue.yaml
- test/e2e/playable_debrief_success.yaml
- test/e2e/playable_coaching_panel.yaml
- test/e2e/playable_ftue_first_five_route.yaml

## 영향 도메인

- Gameplay: 저장 포맷 변경 없이 mission/review/daily/best record를 하나의 반복 플레이 루프로 묶는다.
- Verification: 기존 app_state `review`와 `ui.focus_panel` assertion을 사용해 루프 의미를 검증한다.

Domain-contract 핵심 제약:

- `internal/progress/` 저장 JSON 필드 변경 금지.
- 새 content schema, 새 Vim engine, 새 의존성 추가 금지.
- 사용자 노출 문구는 한국어 톤과 Runbook Dispatch 세계관을 유지한다.
- TUI 조작은 E2E key trace와 app_state로 재현 가능해야 한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: review queue, daily route, success debrief, progress v1 저장 경계
- `docs/verification/spec.md`: review/daily app_state, focus panel app_state, playable E2E evidence
- `docs/roadmap/FOUNDATION_EXIT_REVIEW_2026-05-30.md`: 다음 병목은 mission/review/game loop

## 목표

현재 흩어져 있는 `재점검`, `잔류 리스크`, `오늘의 복구 루트`, best record, playlist 완료 화면을 저장 포맷 변경 없이 하나의 플레이 루프로 묶는다. 성공 화면은 다음 출격 이유를 분명히 보여주고, 마지막 dispatch에서 review 후보가 남아 있으면 `enter`로 primary review exercise에 재진입한다.

## Step 1: Spec and Plan Gate

- 목표: 이번 slice의 acceptance criteria와 제한 범위를 문서로 고정한다.
- 변경 파일:
  - `docs/exec-plans/active/platform-review-003-mission-review-loop.md`
  - `docs/gameplay/spec.md`
- 테스트 파일: 없음
- 충족 기준: 구현 전에 저장 포맷 변경 없는 mission/review loop 기준이 문서화된다.
- Boundaries 주의: `internal/progress/` 저장 필드 변경 금지.
- 상세 작업:
  - [x] active ExecPlan 작성
  - [x] gameplay spec에 이번 루프의 현재 동작 기준 추가

## Step 2: Review Dispatch Loop

- 목표: 성공 debrief와 마지막 dispatch enter 동작을 review loop에 연결한다.
- 변경 파일:
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
- 테스트 파일:
  - `internal/playable/model_test.go`
- 충족 기준:
  - 성공 debrief는 `이번 복구`, `최단 복구`, `목표 입력`, `잔류 리스크`, `다음 출격`을 표시한다.
  - 마지막 entry 성공 상태에서 review 후보가 있으면 action은 `Next dispatch: enter`를 표시한다.
  - 마지막 entry 성공 상태에서 `enter`를 누르면 primary review exercise로 이동한다.
  - review 후보가 없으면 기존 `Dispatch complete`/`Playlist complete` 종료 안내를 유지한다.
- Boundaries 주의: progress schema v2, daily streak, persisted review due date 금지.
- 상세 작업:
  - [x] Red test 작성
  - [x] review dispatch target 계산 구현
  - [x] 성공 상태 enter 처리 확장
  - [x] success debrief 문구 정리

## Step 3: E2E Evidence

- 목표: focused E2E가 새 success debrief와 review app_state를 고정한다.
- 변경 파일:
  - `test/e2e/playable_review_queue.yaml`
  - `test/e2e/playable_debrief_success.yaml`
  - `test/e2e/playable_coaching_panel.yaml`
  - `test/e2e/playable_ftue_first_five_route.yaml`
  - `docs/verification/spec.md`
- 테스트 파일:
  - `test/e2e/playable_review_queue.yaml`
  - `test/e2e/playable_debrief_success.yaml`
- 충족 기준:
  - review queue E2E는 success focus panel kind/title과 review app_state를 검증한다.
  - debrief E2E는 `이번 복구`, `최단 복구`, `목표 입력`, `다음 출격` 문구를 검증한다.
  - full playable E2E 통과 기준은 유지된다.
- Boundaries 주의: 실제 사용자 HOME 사용 금지.
- 상세 작업:
  - [x] focused E2E assertion 갱신
  - [x] focused E2E 실행
  - [x] `make e2e-playable` 실행

## Step 4: Roadmap Sync and Complete

- 목표: slice 완료 후 current roadmap과 change log를 동기화한다.
- 변경 파일:
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/ENGINE_TODO.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - `docs/roadmap/PLATFORM_RFC_001.md`
  - `docs/roadmap/CHANGES.md`
  - `docs/exec-plans/completed/platform-review-003-mission-review-loop.md`
- 테스트 파일: 없음
- 충족 기준: 다음 후보가 `CONTENT-BREADTH-002`로 이동하고, `PLATFORM-REVIEW-003` 완료가 기록된다.
- Boundaries 주의: stale next pointer를 남기지 않는다.
- 상세 작업:
  - [x] roadmap entry point 갱신
  - [x] ExecPlan 검증 결과 기록
  - [x] completed로 이동

---

## 실행 규칙 (하네스 모드)

각 Step은 아래 사이클을 순서대로 수행한다.

### 1. Spec 기반 테스트 작성

구현 전 spec/ExecPlan 수용 기준을 테스트로 변환한다. 테스트가 먼저 실패하는 Red 상태를 확인한다.

### 2. 구현

테스트가 통과하도록 최소 범위로 구현한다.

### 3. 검증

- 관련 Go 패키지 테스트
- focused E2E
- 필요 시 `make e2e-playable`
- `git diff --check`

### 4. 문서 동기화

구현된 동작은 `docs/gameplay/spec.md`, `docs/verification/spec.md`, roadmap 문서에 반영한다.

### 5. 커밋 및 푸시

검증 통과 후 커밋하고 push한다.

## 의사결정 로그

- 2026-05-30: SubAgent 아이디에이션 결과, 이번 slice는 새 저장 시스템보다 `성공 -> 기록 -> 잔류 리스크 -> 다음 출격` 행동 연결을 우선한다.
- 2026-05-30: app_state 새 필드는 필수로 보지 않고 기존 `review`와 `ui.focus_panel`로 1차 검증한다.
- 2026-05-30: 마지막 dispatch에서 review 후보가 없으면 기존 `Dispatch complete`를 유지하고, review 후보가 있으면 `Next dispatch: enter`로 primary review exercise에 재진입한다.

## 검증 결과

- Red: `go test ./internal/playable` failed as expected before implementation.
- Green: `go test ./internal/playable ./internal/playableview`: pass
- Focused E2E:
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_debrief_success.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_coaching_panel.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`: pass
- Related package tests: `go test ./internal/review ./internal/playable ./internal/playableview ./internal/e2estate`: pass
- Full Go tests: `go test ./...`: pass
- Full playable E2E: `make e2e-playable`: pass
- `git diff --check`: pass

## 미해결 질문

- 없음. progress schema v2와 public release gate는 후속 slice로 둔다.
