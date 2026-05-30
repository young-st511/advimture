# UI-POLISH-002 — Command Memory Cue

Slice-ID: UI-POLISH-002
Created: 2026-05-30
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/CHANGES.md
- docs/roadmap/UX_BACKLOG_001.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/spec.md
- docs/exec-plans/active/ui-polish-002-command-memory.md
- docs/exec-plans/completed/ui-polish-002-command-memory.md
- internal/playable/model.go
- internal/playable/model_test.go
- test/e2e/playable_coaching_panel.yaml
- test/e2e/playable_incident_hint_affordance.yaml

## 목표

출시 전 UI polish의 첫 slice로 `FocusPanel`에 command memory cue를 추가한다. Tutorial은 현재 배우는 command를 짧게 노출하고, Incident는 기본 running 화면에서 정답 key를 노출하지 않되 hint/failure 경로에서만 참고 command를 점진 공개한다.

## 포함

- `gameEntry`에 exercise `trained_commands` / `reviewed_commands` metadata 연결
- tutorial running `FocusPanel`에 `기억할 명령: ...` line 추가
- incident hint/failure `FocusPanel`에 `참고 명령: ...` line 추가
- focused model tests
- focused E2E 갱신

## 제외

- 새 content schema
- progress 저장 포맷 변경
- renderer layout 재구성
- side rail / pre-start modal
- incident running 기본 화면의 정답 key 노출

## 수용 기준

- tutorial running 화면과 `app_state.ui.focus_panel.lines`는 current exercise의 `trained_commands`를 `기억할 명령: ...`으로 노출한다.
- incident running 기본 화면은 `기억할 명령` / `참고 명령` / `Coach: 훈련 키`를 노출하지 않는다.
- incident에서 `?` hint를 요청하거나 실패하면 `참고 명령: ...`을 노출한다.
- 저장 포맷, content schema, E2E app_state schema는 변경하지 않는다.
- `go test ./internal/playable`, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Scope Decision

- [x] SubAgent 검토로 command memory 최소판을 1순위로 선택한다.
- [x] side rail, color pass, pre-start modal은 후속으로 분리한다.

## Step 2: Red/Green

- [x] model Red tests 추가
- [x] command metadata wiring 구현
- [x] focus panel line 구현

## Step 3: E2E

- [x] tutorial focused E2E 갱신
- [x] incident hint focused E2E 갱신

## Step 4: Validation

- [x] focused Go tests
- [x] focused E2E
- [x] full regression

## Step 5: Closeout

- [x] docs 동기화
- [x] ExecPlan completed 이동
- [x] 커밋/푸시
