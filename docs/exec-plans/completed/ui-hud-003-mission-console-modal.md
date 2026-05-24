# UI-HUD-003 — Mission HUD + Floating Feedback Modal

Slice-ID: UI-HUD-003
Created: 2026-05-24
Status: completed
Completed: 2026-05-24
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- internal/playable/
- internal/playableview/
- test/e2e/

## 영향 도메인

- Gameplay: 화면 위계와 사용자 노출 문구를 바꾼다. Vim 학습 목표가 가장 먼저 보여야 한다.
- Verification: 화면 구조 변경은 renderer unit test와 playable E2E로 검증한다.

## 목표

현재 화면은 header, briefing, 복구 현황, FocusPanel, console, status가 세로로 나열되어 디버그 대시보드처럼 보인다. 이 slice는 화면을 `Mission HUD + Console Core`로 재정렬하고, 실패/성공/완료 피드백은 Zellij floating pane처럼 console core 안에 뜨는 modal로 표현한다.

## 범위

- 포함:
  - terminal size가 있는 playable screen은 mission HUD 구조로 렌더링한다.
  - 현재 목표와 학습/제한 cue를 상단 mission 영역에 둔다.
  - `RUNBOOK CONSOLE`을 화면의 중심 조작 영역으로 둔다.
  - `failure`, `success` FocusPanel은 console core 안의 floating modal로 렌더링한다.
  - modal은 `Mistake/Concept/Next` 또는 `Result/Learned/Next`에 준하는 읽기 구조를 가진다.
  - 기존 key trace, progress, app_state 계약은 유지한다.
  - 기존 full playable E2E가 통과하도록 screen assertion을 필요한 만큼 보강한다.
- 제외:
  - split cockpit wide layout
  - 실제 입력 차단 modal state
  - content schema 변경
  - progress 저장 포맷 변경
  - 새 dependency 추가

## 수용 기준

- terminal width/height가 있는 화면은 `MISSION`, `RUNBOOK CONSOLE`, status bar 순서로 렌더링된다.
- `복구 현황`은 별도 큰 섹션으로 console 앞을 차지하지 않고 mission HUD 내부의 보조 line으로 낮아진다.
- failed 상태의 `RECOVERY REQUIRED`는 `RUNBOOK CONSOLE` 이후 floating modal처럼 나타난다.
- succeeded 상태의 `STEP SEALED`는 `RUNBOOK CONSOLE` 이후 floating modal처럼 나타난다.
- failed modal은 실패 이유와 retry action을 모두 보존한다.
- succeeded modal은 복구 기록과 next action을 모두 보존한다.
- modal 내용이 길어져도 action line(`Retry`, `Next`, `Next tutorial`)은 잘리지 않는다.
- 기존 playable E2E suite는 통과한다.

## Step 1: Renderer Contract Test

- 목표: 새 UI 위계를 테스트로 먼저 고정한다.
- 변경 파일:
  - `internal/playableview/render_test.go`
- 상세 작업:
  - [x] mission HUD가 console 앞에 오는 테스트
  - [x] 실패 modal이 console 뒤에 뜨는 테스트
  - [x] 성공 modal이 action line을 보존하는 테스트

## Step 2: HUD Renderer

- 목표: terminal size가 있는 경우 mission HUD + console core renderer를 사용한다.
- 변경 파일:
  - `internal/playableview/render.go`
- 상세 작업:
  - [x] mission area 렌더링 함수 분리
  - [x] console core 렌더링 함수 분리
  - [x] floating modal 렌더링 함수 분리
  - [x] terminal size가 없는 fallback 유지

## Step 3: E2E/Docs Sync

- 목표: full playable E2E와 문서를 새 UI 기준으로 맞춘다.
- 변경 파일:
  - `test/e2e/`
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/roadmap/PROGRAM.md`
- 상세 작업:
  - [x] 필요한 screen assertion 안정화
  - [x] gameplay spec 현재 동작 반영
  - [x] screen contract 업데이트
  - [x] `go test ./...`, `make e2e-playable`, `git diff --check`

## 검증 결과

- `go test ./internal/playable ./internal/playableview ./internal/tuiadapter`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_mismatch_feedback.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_forbidden_retry.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- `make e2e-playable`
- `go test ./...`
- `git diff --check`

## 실행 규칙

테스트 작성 → 구현 → 변경 범위 테스트 → full playable E2E → 문서 동기화 → completed 이동 → 커밋/푸시 순서로 진행한다.
