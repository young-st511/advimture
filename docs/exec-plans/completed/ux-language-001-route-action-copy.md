# UX-LANGUAGE-001 — Route Action Copy

Slice-ID: UX-LANGUAGE-001
Created: 2026-05-25
Status: completed
Completed: 2026-05-25
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- docs/exec-plans/active/ux-language-001-route-action-copy.md
- internal/playable/
- internal/playableview/
- test/e2e/

## 목표

성공 modal의 다음 행동 문구를 tutorial/incident 맥락에 맞게 정리한다. 특히 incident 사이에서 `Next tutorial: enter`가 보이고, 마지막 incident에서 `Playlist complete`가 보이는 어색함을 제거한다.

## 범위

- 포함:
  - 같은 playlist 내 다음 exercise는 기존 `Next: enter` 유지
  - 다음 playlist가 tutorial이면 `Next tutorial: enter`
  - 다음 playlist가 incident이면 `Next runbook: enter`
  - 마지막 incident는 `Dispatch complete`
  - 관련 model/E2E 테스트 갱신
- 제외:
  - modal 레이아웃 변경
  - 진행 저장 포맷 변경
  - playlist ordering 변경

## 수용 기준

- tutorial에서 다음 tutorial로 넘어갈 때는 `Next tutorial: enter`를 유지한다.
- tutorial에서 incident로 넘어갈 때와 incident 사이에서는 `Next runbook: enter`를 표시한다.
- 마지막 incident 완료 시 `Dispatch complete`를 표시한다.
- `make e2e-playable`이 통과한다.

## Step 1: Contract Update

- [x] spec/contract 문구 갱신
- [x] action copy 규칙을 model 테스트로 고정

## Step 2: Implementation

- [x] success action line helper 구현
- [x] stale E2E fixture 갱신

## Step 3: Verification

- [x] focused tests
- [x] `make e2e-playable`
- [x] `git diff --check`

## 결정

- 같은 playlist 내 다음 exercise는 `Next: enter`를 유지한다.
- 다음 playlist가 tutorial이면 `Next tutorial: enter`를 유지한다.
- 다음 playlist가 incident이면 `Next runbook: enter`를 표시한다.
- 마지막 incident는 `Dispatch complete`를 표시한다.

## 검증 결과

- `go test ./internal/playable ./internal/playableview`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_line_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_004_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`
