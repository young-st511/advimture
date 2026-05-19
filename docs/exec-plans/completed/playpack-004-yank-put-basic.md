# PLAYPACK-004 — Yank / Put Basic Tutorial

Slice-ID: PLAYPACK-004
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/playpack-004-yank-put-basic.md
- docs/exec-plans/completed/playpack-004-yank-put-basic.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- content/
- test/e2e/
- internal/content/

## 목표

`yank-put-basic`을 playable tutorial로 승격한다. 플레이어가 정상 설정 줄이나 토큰을 다시 타이핑하지 않고 `y`로 보관하고 `p/P`로 재사용하는 감각을 얻는다.

## 범위

- 포함:
  - `yank-put-basic` content cluster implemented 승격
  - 5문항 exercise: `yy+p`, `yy+P`, `yw+P`, `y$+p`, `yw+p`
  - scenario YAML
  - `tutorial-6-yank-put` playlist
  - full playlist E2E
  - content loader count/coverage tests 갱신
- 제외:
  - text object
  - named register / clipboard
  - count prefix
  - visual selection

## 수용 기준

- `yank-put-basic`은 `yw`, `y$`, `yy`, `p`, `P` coverage를 모두 만족한다.
- 모든 exercise는 `replay_status: pass`이고 replay gate를 통과한다.
- playlist는 8문항 이하이며 ID 순서상 `tutorial-5-operator-grammar` 다음에 실행된다.
- E2E는 5문항을 순서대로 완료하고 progress, key trace, final app state를 검증한다.
- text object 관련 command는 이 playpack에 포함하지 않는다.

## 결과

- `content/command_clusters/yank-put.yaml`, `content/exercises/yank-put.yaml`, `content/scenarios/yank-put.yaml`, `content/playlists/yank-put.yaml`를 추가했다.
- `internal/content/loader_test.go`에 root content count, playable playlist, `yank-put-basic` coverage expectation을 추가했다.
- `test/e2e/playable_yank_put_full.yaml`로 다섯 문항 전체 흐름과 progress 저장을 검증한다.
- `docs/gameplay/spec.md`와 `docs/gameplay/command-catalog.md`를 implemented 상태로 동기화했다.

## 검증

- `go test ./internal/content/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_yank_put_full.yaml`
- `make e2e-smoke`
