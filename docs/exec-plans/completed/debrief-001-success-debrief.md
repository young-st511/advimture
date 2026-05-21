# DEBRIEF-001 — Success Debrief Without Save Schema Change

Slice-ID: DEBRIEF-001
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/debrief-001-success-debrief.md
- docs/exec-plans/completed/debrief-001-success-debrief.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/spec.md
- internal/playable/
- test/e2e/

## 목표

성공 상태와 playlist 완료 상태에서 학습 회고 정보를 더 잘 보이게 한다. 저장 포맷은 바꾸지 않고, 현재 score/runtime trace와 기존 `Progress.Missions` best record만 읽어 표시한다.

## 결과

- 성공 action panel에 현재 grade/key count debrief를 추가했다.
- 기존 progress의 best grade / best keystrokes를 표시한다.
- 현재 playlist 완료 수를 표시한다.
- playlist 완료 화면에서도 debrief와 best record가 유지된다.
- `test/e2e/playable_debrief_success.yaml` focused E2E를 추가했다.

## 수용 기준

- [x] 성공 상태 action panel은 현재 grade와 key count를 표시한다.
- [x] 성공 상태 action panel은 progress에 저장된 best grade와 best keystrokes를 표시한다.
- [x] 성공 상태 action panel은 현재 playlist 완료 수와 전체 수를 표시한다.
- [x] 마지막 exercise 성공 후 playlist complete 화면에도 debrief와 best record가 보인다.
- [x] 실패/진행 중/action hint UI는 기존 의미를 유지한다.
- [x] progress 저장 JSON 필드는 변경하지 않는다.

## 검증 결과

- `go test ./internal/playable/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_debrief_success.yaml`
- `make e2e-smoke`
- `git diff --check`

## 후속

REPEAT-GAP-001에서 `.` repeat-last-change의 transaction 경계를 먼저 고정한다.
