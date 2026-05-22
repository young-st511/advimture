# UI-HIERARCHY-001 — Current Task First Hierarchy

Slice-ID: UI-HIERARCHY-001
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: renderer-and-tests
Allowed-Paths:
- docs/exec-plans/active/ui-hierarchy-001-current-task-first.md
- docs/exec-plans/completed/ui-hierarchy-001-current-task-first.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/playableview/
- internal/playable/model_test.go

## 목표

현재 exercise의 제목과 briefing을 화면 상위로 올리고, review queue와 daily route는 보조 운영 정보로 낮춘다. 기존 game loop, progress 저장, content schema, E2E state shape는 변경하지 않는다.

## 완료 내용

- Header를 `ADVIMTURE | <playlist> | Exercise: n/m | Status: <status>`로 압축했다.
- 현재 exercise title/message를 review/daily route보다 먼저 렌더링한다.
- review/daily route는 `OPS` 영역으로 낮췄다.
- buffer는 `RUNBOOK CONSOLE` 영역 아래에 모았다.
- 기존 `Exercise:`, `Mode:`, `Status:`, review/daily 문구는 화면에서 계속 찾을 수 있게 유지했다.

## 검증 결과

- `go test ./internal/playableview/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `git diff --check`

참고: `go run .`은 Codex 비대화형 실행 환경에서 `/dev/tty`가 없어 실패했다. 이 slice 검증은 TUI E2E runner와 단위 테스트로 수행했다.

## 후속 작업

다음 slice는 `UI-MODE-001`이다. tutorial과 incident의 action panel 언어를 분리해, tutorial은 훈련 키를 명시하고 incident는 판단 cue와 점진 힌트를 우선한다.
