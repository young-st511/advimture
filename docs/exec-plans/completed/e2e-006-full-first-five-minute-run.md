# ExecPlan: Full first 5-minute playlist E2E

Slice-ID: E2E-006
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- test/e2e/**
- Makefile
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/e2e-006-full-first-five-minute-run.md
- docs/exec-plans/completed/e2e-006-full-first-five-minute-run.md

## 목표

첫 5분 playlist의 모든 playable exercise를 실제 TUI 입력으로 처음부터 끝까지 완주한다. 단일 smoke와 부분 command-mode E2E를 넘어, 플레이어가 첫 시나리오 팩을 끊김 없이 완료할 수 있음을 검증한다.

## 영향 도메인

- Verification: full playlist pty scenario를 추가하고 playable E2E suite에 포함한다.
- Gameplay: success 후 `enter` next transition이 마지막 문제까지 안정적으로 이어져야 한다.
- Safety: `setup.home: temp`만 사용하고 실제 progress는 건드리지 않는다.

## 수용 기준

- full E2E는 14개 playable exercise를 순서대로 성공시킨다.
- final screen은 마지막 range substitute 성공 문구와 `Playlist complete`를 포함한다.
- progress file에는 첫 exercise와 마지막 exercise completion이 모두 남는다.
- app state summary는 마지막 exercise의 buffer, command, status, progress를 검증한다.
- `make e2e-playable`은 full playlist scenario까지 실행한다.
- full scenario는 최소 2회 연속 통과한다.

## 범위

- 포함: full-run E2E YAML
- 포함: Makefile suite 연결
- 포함: verification docs 갱신
- 제외: 새 runner DSL 추가
- 제외: content polish 문구 수정
- 제외: CI workflow 추가

## 검증 계획

- `go test ./cmd/e2e-runner`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `make e2e-playable` 반복 실행
- `git diff --check`

## 작업 항목

- [x] full playlist E2E scenario를 작성한다.
- [x] `make e2e-playable`에 full scenario를 포함한다.
- [x] 실패 evidence를 보고 flaky step이나 assertion을 보강한다.
- [x] full scenario 2회 연속 통과를 확인한다.
- [x] 문서를 completed 상태로 동기화한다.
