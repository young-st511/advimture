# UI-MODAL-001 — Focus Modal Foundation

Slice-ID: UI-MODAL-001
Created: 2026-05-23
Status: completed
Completed: 2026-05-23
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- internal/playable/
- internal/playableview/
- internal/e2estate/
- cmd/e2e-runner/
- test/e2e/

## 영향 도메인

- Gameplay: 사용자 노출 문구는 한국어 톤을 우선하며, 저장 포맷과 content schema는 변경하지 않는다.
- Verification: TUI E2E는 테스트 전용 HOME을 사용하고, 화면 문자열뿐 아니라 typed app_state를 함께 검증한다.

## 목표

현재 안내/성공/실패 패널이 `RUNBOOK CONSOLE` 아래에 렌더링되어 플레이어가 안내 변화를 놓치기 쉽다. 이 slice는 안내를 console 위의 modal-like `FocusPanel`로 승격해 현재 상태의 다음 행동을 먼저 보이게 만들고, E2E가 해당 UI intent를 구조적으로 검증할 수 있게 한다.

## 범위

- 포함:
  - `playableview.Screen`에 structured focus panel 모델 추가
  - `internal/playable`에서 tutorial/incident/failure/success/mode별 focus panel intent 생성
  - focus panel을 briefing/복구 현황 아래, `RUNBOOK CONSOLE` 위에 렌더링
  - `e2estate.State`와 E2E runner에 `ui.focus_panel` assertion 추가
  - 대표 E2E fixture로 tutorial/failure/incident focus panel 검증
  - 관련 gameplay/verification 문서 동기화
- 제외:
  - `tea.WindowSizeMsg` 기반 true 중앙 overlay
  - modal open 상태에서 Vim runtime 입력을 막는 pre-start modal
  - progress 저장 포맷 변경
  - content schema 변경
  - 새 Go dependency 추가

## 수용 기준

- `FocusPanel`은 `kind`, `title`, `lines`를 가진 구조화 모델이다.
- running tutorial은 `training` kind와 `TRAINING BRIEF` title을 보여주며 required-key coaching과 `?: hint  q: quit`을 포함한다.
- running incident는 `incident` kind와 `OPERATOR JUDGMENT` title을 보여주며 정답 key sequence를 기본 노출하지 않고 판단 cue를 우선한다.
- failed 상태는 `failure` kind와 `RECOVERY REQUIRED` title을 보여주며 실패 회복 정보, `Retry: r or enter`, `?: hint  q: quit`을 포함한다.
- succeeded 상태는 `success` kind와 `STEP SEALED` title을 보여주며 기존 debrief와 `Next`/`Next tutorial`/`Playlist complete` 안내를 유지한다.
- insert/visual/command/search mode는 `mode` kind와 상태별 title을 보여주며 실제 입력 처리와 맞지 않는 hint/quit 안내를 섞지 않는다.
- focus panel은 최종 화면에서 `RUNBOOK CONSOLE`보다 먼저 렌더링된다.
- E2E app_state는 `ui.focus_panel.kind`, `title`, `lines`를 노출하고 runner가 이를 typed assertion으로 검증할 수 있다.
- 기존 key trace, progress assertion, review/daily app_state assertion은 유지된다.

## Step 1: FocusPanel 렌더링 기반

- 목표: 화면 아래 action panel을 console 위 focus panel로 승격한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
- 테스트 파일:
  - `internal/playableview/render_test.go`
  - `internal/playable/model_test.go`
- 충족 기준: `FocusPanel` 구조화 모델, 상태별 title/kind, console 이전 렌더링
- Boundaries 주의: 저장 포맷, content schema, dependency 변경 금지
- 상세 작업:
  - [x] `FocusPanel` 타입과 renderer 추가
  - [x] `Model`에서 상태별 focus panel 생성
  - [x] 기존 문구 호환성을 유지하며 `ACTION` 디버그 label 제거
  - [x] command/search/insert mode에서 실제 입력 처리와 맞는 안내만 표시

## Step 2: E2E UI state assertion

- 목표: 화면 텍스트뿐 아니라 app_state로 focus panel intent를 검증한다.
- 변경 파일:
  - `internal/e2estate/state.go`
  - `internal/e2estate/state_test.go`
  - `internal/playable/model.go`
  - `cmd/e2e-runner/main.go`
  - `cmd/e2e-runner/main_test.go`
  - `test/e2e/*.yaml`
- 테스트 파일:
  - `internal/e2estate/state_test.go`
  - `cmd/e2e-runner/main_test.go`
  - 대표 `test/e2e/` fixture
- 충족 기준: `ui.focus_panel` typed assertion, tutorial/failure/incident 대표 E2E
- Boundaries 주의: E2E HOME은 temp 또는 injected path만 사용
- 상세 작업:
  - [x] `e2estate.UI`와 `FocusPanel` 추가
  - [x] runner YAML assertion에 `app_state.ui.focus_panel` 추가
  - [x] 대표 fixture에 focus panel assertion 추가
  - [x] evidence와 기존 app_state assertion 회귀 확인

## Step 3: 문서 동기화와 다음 slice 분리

- 목표: 구현된 focus panel 동작을 canonical docs에 반영하고, true overlay/intro modal을 후속 slice로 분리한다.
- 변경 파일:
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/verification/spec.md`
  - `docs/roadmap/PROGRAM.md`
- 테스트 파일: n/a
- 충족 기준: 현재 동작과 후속 제외 범위가 문서에서 모순 없이 읽힌다.
- Boundaries 주의: 완료된 과거 구현을 canonical로 역인용하지 않는다.
- 상세 작업:
  - [x] `ACTION` panel 기준을 `FocusPanel` 기준으로 갱신
  - [x] true 중앙 overlay와 pre-start modal을 후속 slice로 기록
  - [x] active slice를 completed로 이동할 준비

## 검증 결과

- `go test ./internal/playableview ./internal/playable`
- `go test ./internal/e2estate ./internal/playable ./cmd/e2e-runner`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_arrow_forbidden.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ctrl_c_does_not_quit.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_focus_panel.yaml`
- `make e2e-playable`
- `git diff --check`

---

## 실행 규칙 (하네스 모드)

각 Step은 아래 사이클을 순서대로 수행한다.

### 1. Spec 기반 테스트 작성

수용 기준을 먼저 테스트로 옮긴다. 구현 전 실패하는 테스트를 확인한다.

### 2. 구현

테스트가 통과하도록 최소 범위로 구현한다.

### 3. 검증

- 변경 범위 패키지 테스트를 먼저 실행한다.
- 공통 동작 영향이 있으면 `go test ./...`와 대표 E2E를 실행한다.
- 필요 시 SubAgent로 Plan 정합성 검증과 코드 리뷰를 분리한다.

### 4. 문서 동기화

구현된 현재 동작을 gameplay/verification docs에 반영한다.

### 5. 커밋 및 푸시

검증 통과 후 관련 변경을 하나의 커밋으로 묶고 push한다.
