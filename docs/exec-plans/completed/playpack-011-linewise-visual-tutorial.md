# PLAYPACK-011 — Linewise Visual Tutorial

Slice-ID: PLAYPACK-011
Created: 2026-05-23
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/playpack-011-linewise-visual-tutorial.md
- docs/exec-plans/completed/playpack-011-linewise-visual-tutorial.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- content/command_clusters/visual-line-basic.yaml
- content/exercises/visual-line-basic.yaml
- content/scenarios/visual-line-basic.yaml
- content/playlists/visual-line-basic.yaml
- internal/content/loader_test.go
- test/e2e/playable_visual_line_full.yaml
- Makefile

## 목표

linewise `V`를 3문항 이하 tutorial로 승격하고 full E2E로 검증한다.

## 수용 기준

- `visual-line-basic` command cluster는 approved + implemented 상태다.
- tutorial은 3문항 이하로 유지한다.
- `Vd`, `Vy`, linewise register + `p`, `VGd`를 다룬다.
- 각 exercise는 `V`를 required key로 요구하고 lowercase `v` 우회를 금지한다.
- full E2E는 progress, key_trace, app_state buffer/cursor/mode/progress를 검증한다.

## 제외 항목

- multi-line charwise visual
- visual block
- visual `c`, indent, count/register prefix
- incident 적용 run

## 검증 계획

- `go test ./internal/content/...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_line_full.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 완료 내용

- `visual-line-basic` command cluster를 approved + implemented content로 추가했다.
- `visual-line-basic-001..003` exercise로 linewise delete, linewise yank+put, `VGd` tail delete를 구성했다.
- `tutorial-93-visual-line` playlist를 추가했다.
- `playable_visual_line_full.yaml` E2E를 추가하고 `make e2e-playable`에 연결했다.

## 검증 결과

- passed: `go test ./internal/content/...`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_line_full.yaml`
- passed: `go test ./...`
- passed: `make e2e-playable`
