# PLAYPACK-006 — Open Line Edit Tutorial

Slice-ID: PLAYPACK-006
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/playpack-006-open-line-edit.md
- docs/exec-plans/completed/playpack-006-open-line-edit.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- docs/gameplay/vim-curriculum-map.md
- content/
- internal/content/
- test/e2e/

## 목표

`open-line-edit`을 playable tutorial로 승격한다. 플레이어는 `o`와 `O`로 현재 줄 아래/위에 새 줄을 열고, 즉시 필요한 값을 입력한 뒤 `esc`로 Normal mode에 돌아오는 흐름을 익힌다.

## 결과

- `content/command_clusters/open-line-edit.yaml`를 추가했다.
- `content/exercises/open-line-edit.yaml`에 5문항을 추가했다.
- `content/scenarios/open-line-edit.yaml`와 `content/playlists/open-line-edit.yaml`를 추가했다.
- `internal/content/loader_test.go`에 root count, playable IDs, playlist order, coverage expectation을 갱신했다.
- `test/e2e/playable_open_line_full.yaml`로 전체 playlist 완료를 검증했다.
- `test/e2e/playable_open_line_forbidden_route.yaml`로 잘못된 `O` 방향 입력 후 retry 회복을 검증했다.

## 수용 기준

- [x] `open-line-edit` command cluster는 approved + implemented content로 로드된다.
- [x] playlist는 8문항 이하이며 `tutorial-7-text-object-inner-word` 다음에 실행된다.
- [x] 각 exercise는 `replay_status: pass`, `optimal_keys`, `constraints.required_keys`, `e2e_assertions`를 가진다.
- [x] `o`와 `O`가 각각 최소 2문항 이상 optimal trace와 coverage에 등장한다.
- [x] 잘못된 방향의 `o/O` 입력은 즉시 실패 후 retry로 회복할 수 있다.
- [x] full E2E는 전체 playlist 완료, progress 저장, key trace, final app state를 검증한다.

## 검증 결과

- `go test ./internal/content/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_open_line_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_open_line_forbidden_route.yaml`
- `make e2e-smoke`
- `git diff --check`

## 후속

DEBRIEF-001에서 저장 포맷 변경 없이 성공/playlist 완료 debrief와 기존 best record 비교를 추가한다.
