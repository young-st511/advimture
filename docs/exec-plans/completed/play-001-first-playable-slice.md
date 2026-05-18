# ExecPlan: First playable vertical slice

Slice-ID: PLAY-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/app/**
- internal/playable/**
- internal/e2estate/**
- test/e2e/**
- Makefile
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/play-001-first-playable-slice.md

## 목표

새 엔진 경로를 실제 TUI 앱에 아주 작게 연결한다. 플레이어는 `l`, `l`을 눌러 cursor를 목표 위치로 이동하고, 성공 message, score, progress, app state summary를 확인할 수 있어야 한다.

## 범위

- 포함: env-gated playable mode
- 포함: `content -> scenario -> tuiadapter -> Bubble Tea -> progressadapter`
- 포함: `ADVIMTURE_E2E=1`일 때 app state summary export
- 포함: playable E2E scenario
- 제외: 기존 FTUE/menu 제거
- 제외: 기존 editor/game archive
- 제외: 복수 mission flow
- 제외: 새 Vim command

## 구현 계획

1. `internal/playable`에 첫 playable Bubble Tea model을 만든다.
2. model은 hardcoded `normal-motion-basic` exercise 하나만 가진다.
3. `internal/app`은 `ADVIMTURE_PLAYABLE_SLICE=1`일 때 playable model로 시작한다.
4. 성공 시 `progressadapter`로 mission completion을 만들고 기존 progress 저장을 호출한다.
5. `ADVIMTURE_E2E=1`이면 `.advimture/e2e_state.json`을 쓴다.
6. E2E scenario는 state summary와 progress file을 검증한다.

## 검증 계획

- `go test ./internal/playable`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml`
- `make e2e-smoke`

## 승인 체크

- [x] `ADVIMTURE_PLAYABLE_SLICE=1`에서 새 playable model로 시작한다.
- [x] `l`, `l` 입력으로 scenario가 succeeded가 된다.
- [x] 성공 시 score `S`가 생긴다.
- [x] success 시 progress file이 생성된다.
- [x] E2E app state summary가 buffer/cursor/mode/status/score/progress를 검증한다.
- [x] 기존 FTUE smoke E2E는 유지된다.

## 의사결정 로그

- 2026-05-18: 기존 FTUE/menu 흐름은 유지하고 `ADVIMTURE_PLAYABLE_SLICE=1`에서만 새 playable path로 시작하게 했다. 기존 구현을 즉시 archive하기 전에 새 path를 검증하기 위함이다.
- 2026-05-18: E2E 입력 사이에 짧은 wait를 넣었다. Bubble Tea raw input 준비 직후 키 손실을 줄이기 위한 flake 방지다.
- 2026-05-18: LEGACY-001 이후 playable path가 기본 앱 경로가 되었고 `ADVIMTURE_PLAYABLE_SLICE` gate는 제거됐다. 이 문서는 PLAY-001 당시의 완료 기록으로 보존한다.
