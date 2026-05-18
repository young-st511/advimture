# ExecPlan: Archive old implementation

Slice-ID: LEGACY-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/app/**
- internal/progress/**
- internal/editor/**
- internal/game/**
- internal/data/**
- internal/ui/**
- docs/archived/legacy-code/**
- docs/archived/legacy-e2e/**
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/LEGACY_INVENTORY.md
- docs/gameplay/domain-contract.md
- docs/gameplay/spec.md
- docs/guardrails.md
- docs/verification/spec.md
- docs/exec-plans/completed/e2e-runner-bootstrap.md
- docs/exec-plans/completed/play-001-first-playable-slice.md
- docs/exec-plans/completed/legacy-001-archive-old-implementation.md
- test/e2e/**
- Makefile
- AGENTS.md

## 목표

새 playable path가 E2E로 검증됐으므로 기존 editor/game/data/ui 구현을 active code path에서 제거하고 archive한다.

## 범위

- 포함: app 기본 경로를 새 playable path로 전환
- 포함: 기존 `internal/editor`, `internal/game`, `internal/data`, `internal/ui` archive
- 포함: legacy `.go` 파일을 `.go.txt`로 보관
- 포함: obsolete FTUE E2E scenario archive
- 제외: `internal/progress` save schema 변경
- 제외: 새 command cluster 추가
- 제외: multi-mission flow

## 검증 계획

- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`

## 승인 체크

- [x] active app path가 새 playable model을 기본으로 사용한다.
- [x] legacy Go code가 `go test ./...` 대상에서 빠진다.
- [x] legacy source는 archive에 남는다.
- [x] obsolete FTUE E2E는 archive로 이동한다.
- [x] 전체 테스트와 E2E가 통과한다.

## 완료 결과

- `internal/app` 기본 경로를 새 `internal/playable` model로 전환했다.
- 기존 `internal/editor`, `internal/game`, `internal/data`, `internal/ui`를 `docs/archived/legacy-code/2026-05-18/`에 보관했다.
- archived Go source는 `*.go.txt`로 바꿔 `go test ./...` 대상에서 제외했다.
- obsolete FTUE E2E scenario는 `docs/archived/legacy-e2e/2026-05-18/`에 보관했다.
- E2E cache flake는 E2E-004로 분리해 보강했다.
