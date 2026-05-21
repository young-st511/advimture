# PLAYPACK-007 — Repeat Last Change Efficiency Tutorial

Slice-ID: PLAYPACK-007
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/playpack-007-repeat-last-change.md
- docs/exec-plans/completed/playpack-007-repeat-last-change.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- content/
- internal/content/
- test/e2e/

## 목표

`repeat-last-change`를 playable tutorial로 승격한다. 플레이어는 `.`를 사용해 같은 변경을 다시 타이핑하지 않고 반복 적용하는 효율을 체감해야 한다.

## 결과

- `content/command_clusters/repeat-last-change.yaml`를 추가했다.
- `content/exercises/repeat-last-change.yaml`에 4문항을 추가했다.
- `content/scenarios/repeat-last-change.yaml`와 `content/playlists/repeat-last-change.yaml`를 추가했다.
- `internal/content/loader_test.go`에 root count, playable IDs, playlist order, coverage expectation을 갱신했다.
- `test/e2e/playable_repeat_last_change_full.yaml`로 전체 playlist 완료를 검증했다.

## 수용 기준

- [x] `repeat-last-change` command cluster는 approved + implemented content로 로드된다.
- [x] 모든 exercise는 `.`를 `trained_commands`와 `constraints.required_keys`에 포함한다.
- [x] 문항은 insert, change, open-line, replace-char repeat을 각각 최소 1회 다룬다.
- [x] playlist는 8문항 이하이며 `tutorial-8-open-line-edit` 다음에 실행된다.
- [x] full E2E는 전체 playlist 완료, progress 저장, key trace, final app state를 검증한다.

## 검증 결과

- `go test ./internal/content/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_repeat_last_change_full.yaml`
- `make e2e-smoke`
- `git diff --check`

## 후속

SEARCH-GAP-001에서 literal `/`, `n`, `N` search 범위와 `?` hint 충돌 처리 방침을 고정한다.
