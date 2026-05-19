# PLAYPACK-005 — Text Object Inner Word Tutorial

Slice-ID: PLAYPACK-005
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/playpack-005-text-object-inner-word.md
- docs/exec-plans/completed/playpack-005-text-object-inner-word.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- content/
- internal/content/
- test/e2e/

## 목표

`text-object-inner-word`를 playable tutorial로 승격한다. 플레이어는 단어 시작으로 이동하지 않고 커서가 단어 중간에 있어도 `diw`, `ciw`, `yiw`로 단어 전체를 대상으로 잡는 감각을 얻는다.

## 결과

- `content/command_clusters/text-object-inner-word.yaml`를 추가했다.
- `content/exercises/text-object-inner-word.yaml`에 6문항을 추가했다.
- `content/scenarios/text-object-inner-word.yaml`와 `content/playlists/text-object-inner-word.yaml`를 추가했다.
- `internal/content/loader_test.go`에 root count, playable IDs, playlist order, coverage expectation을 갱신했다.
- `test/e2e/playable_text_object_inner_word_full.yaml`로 전체 playlist 완료를 검증했다.
- `test/e2e/playable_text_object_inner_word_forbidden_route.yaml`로 단어 시작 이동 우회를 실패 처리하는지 검증했다.

## 수용 기준

- `text-object-inner-word`는 `diw`, `ciw`, `yiw` coverage를 모두 만족한다.
- 모든 exercise는 `replay_status: pass`이고 replay gate를 통과한다.
- playlist는 8문항 이하이며 ID 순서상 `tutorial-6-yank-put` 다음에 실행된다.
- full E2E는 6문항을 순서대로 완료하고 progress, key trace, final app state를 검증한다.
- forbidden route E2E는 단어 시작으로 이동하는 우회를 실패로 검증한다.

## 검증

- `go test ./internal/content/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_inner_word_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_inner_word_forbidden_route.yaml`
- `make e2e-smoke`
- `git diff --check`
