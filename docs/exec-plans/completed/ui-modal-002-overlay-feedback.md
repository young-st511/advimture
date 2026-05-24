# UI-MODAL-002 — Overlay Feedback Modal

Slice-ID: UI-MODAL-002
Created: 2026-05-24
Status: completed
Completed: 2026-05-24
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- cmd/e2e-runner/
- internal/playable/
- internal/playableview/
- internal/runtime/
- internal/scenario/
- test/e2e/

## 영향 도메인

- Gameplay: 사용자 노출 피드백은 한국어 톤을 유지하고 Vim 학습 의도를 먼저 설명한다.
- Verification: TUI feedback 변경은 key trace, app_state, screen evidence로 검증한다.

## 목표

현재 FocusPanel은 console 위에 flow로 삽입되어 상태마다 화면 높이를 밀어낸다. 또한 실패/성공 피드백이 briefing 영역에 섞여 안내처럼 보이지 않고, `:q!` 문항에서 `:wq` 같은 반대 의도 command를 실행해도 명확한 거절이 없다. 이 slice는 FocusPanel을 overlay처럼 렌더링하고, 학습/거절 피드백을 FocusPanel로 모아 플레이어가 즉시 알아차리게 만든다.

## 범위

- 포함:
  - terminal width/height가 있을 때 FocusPanel을 base layout 위에 overlay로 합성
  - failed/succeeded 피드백을 FocusPanel 안에 표시하고 briefing은 원래 미션 설명으로 유지
  - command goal 문항에서 실행된 command가 목표 command와 다르면 즉시 실패 피드백
  - `:q!` 목표에서 `:wq`, `:wq` 목표에서 `:q!`를 구분하는 문구
  - 모달 줄바꿈 때문에 E2E 문장 검증이 흔들리지 않도록 screen text 정규화
  - renderer/model/runtime/E2E 테스트
- 제외:
  - modal open 중 입력 차단
  - content schema 변경
  - progress 저장 포맷 변경
  - 새 dependency 추가

## 수용 기준

- terminal height가 있을 때 FocusPanel은 layout flow에 줄을 추가하지 않고 기존 화면 위에 overlay로 합성된다.
- overlay가 켜져도 `RUNBOOK CONSOLE`의 시작 위치는 FocusPanel line count에 따라 밀리지 않는다.
- failed 상태의 scenario/runtime 피드백은 FocusPanel 안에 보이고, briefing 영역은 원래 미션 briefing을 유지한다.
- succeeded 상태의 success text는 FocusPanel 안에 보이고, briefing 영역은 원래 미션 briefing을 유지한다.
- command goal이 `:q!`인데 `:wq`를 실행하면 실패하고 “저장하면 안 되는 임시 작업” 맥락의 거절 피드백을 보여준다.
- command goal이 `:wq`인데 `:q!`를 실행하면 실패하고 “저장해야 하는 작업” 맥락의 거절 피드백을 보여준다.
- 기존 full playable E2E는 통과한다.

## Step 1: Overlay Rendering

- 목표: FocusPanel을 화면 흐름을 밀지 않는 overlay로 렌더링한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
- 상세 작업:
  - [x] base layout과 overlay 합성 분리
  - [x] terminal size가 없는 경우 fallback 유지
  - [x] console 위치가 panel line count에 의존하지 않는 테스트 추가

## Step 2: Feedback Placement

- 목표: 실패/성공 피드백을 FocusPanel로 이동한다.
- 변경 파일:
  - `internal/scenario/run.go`
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
- 상세 작업:
  - [x] scenario state에 원래 briefing을 노출
  - [x] playable briefing은 원래 briefing을 사용
  - [x] failed/succeeded message를 FocusPanel 첫 줄에 표시

## Step 3: Command Goal Mismatch

- 목표: `:q!`/`:wq` 반대 command 실행을 명확히 거절한다.
- 변경 파일:
  - `internal/runtime/session.go`
  - `internal/runtime/session_test.go`
  - `test/e2e/`
- 상세 작업:
  - [x] command goal mismatch 감지
  - [x] `:q!` vs `:wq` 전용 피드백 추가
  - [x] E2E fixture 추가

## Step 4: 문서/검증/완료

- 목표: docs와 Program 상태를 갱신하고 full E2E로 회귀를 확인한다.
- 상세 작업:
  - [x] gameplay docs 갱신
  - [x] ExecPlan completed 이동
  - [x] `go test ./...`, `make e2e-playable`

## 검증 결과

- `go test ./internal/playableview ./internal/runtime ./internal/scenario ./internal/tuiadapter ./internal/playable`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_mismatch_feedback.yaml`
- `go test ./cmd/e2e-runner ./internal/playableview`
- `make e2e-playable`
- `go test ./...`
- `git diff --check`

---

## 실행 규칙

각 Step은 테스트 작성 → 구현 → 변경 범위 테스트 → E2E → 문서 동기화 → 커밋/푸시 순서로 진행한다.
