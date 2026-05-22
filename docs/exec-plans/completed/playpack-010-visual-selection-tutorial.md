# PLAYPACK-010 — Visual Selection Tutorial

Slice-ID: PLAYPACK-010
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/playpack-010-visual-selection-tutorial.md
- docs/exec-plans/completed/playpack-010-visual-selection-tutorial.md
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- content/command_clusters/visual-char-line.yaml
- content/exercises/visual-char-line.yaml
- content/scenarios/visual-char-line.yaml
- content/playlists/visual-char-line.yaml
- internal/content/loader_test.go
- test/e2e/playable_visual_selection_full.yaml
- Makefile

## 목표

charwise visual `v` + motion + `d/y`를 짧은 tutorial playpack으로 연결한다.

## 수용 기준

- `visual-char-line` command cluster를 approved + implemented로 승격한다.
- tutorial은 3~4문항으로 제한한다.
- 최소 한 문항은 `v + motion + d`를 훈련한다.
- 최소 한 문항은 `v + motion + y + p/P`를 훈련한다.
- backward selection normalization을 한 문항에서 다룬다.
- focused E2E는 full visual tutorial을 완료하고 progress/app_state를 검증한다.

## 결과

- `visual-char-line` cluster를 approved + implemented content로 추가했다.
- `visual-char-line-001..003` exercise로 forward deletion, visual yank-put, backward deletion을 구성했다.
- `tutorial-92-visual-selection` playlist를 추가했다.
- `playable_visual_selection_full.yaml` E2E를 추가하고 `make e2e-playable`에 연결했다.

## 제외한 것

- linewise `V`
- visual block `<C-v>`
- multi-line visual operator
- visual `c`, indentation, register prefix
- incident run 편입

## 검증

- `go test ./internal/content/...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_selection_full.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`
