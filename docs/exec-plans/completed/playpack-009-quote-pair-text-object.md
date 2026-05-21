# PLAYPACK-009 — Quote/Pair Text Object Tutorial

Slice-ID: PLAYPACK-009
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/playpack-009-quote-pair-text-object.md
- docs/exec-plans/completed/playpack-009-quote-pair-text-object.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- content/exercises/text-object-quote-pair.yaml
- content/scenarios/text-object-quote-pair.yaml
- content/playlists/text-object-quote-pair.yaml
- internal/content/loader_test.go
- test/e2e/playable_text_object_quote_pair_full.yaml
- test/e2e/playable_incident_001_full.yaml
- Makefile

## 목표

`di"`, `ci"`, `yi"`를 짧은 튜토리얼 playpack으로 연결하고, full playlist E2E로 검증한다.

## 범위

- 포함:
  - 4문항 quote text object exercise
  - scenario skin
  - `tutorial-91-text-object-quote-pair` playlist
  - content count/coverage test 갱신
  - full E2E
- 제외:
  - 새 Vim engine 기능
  - incident content
  - single quote, parenthesis, brace
  - visual selection
  - progress 저장 포맷 변경

## 수용 기준

- completed: tutorial playlist는 8문항 이하이며 search tutorial 뒤, incident run 앞에 실행된다.
- completed: `di"`, `ci"`, `yi"`는 각각 optimal trace와 trained command에 등장한다.
- completed: 각 문항은 `constraints.required_keys`로 quote text object 사용을 고정한다.
- completed: full E2E는 progress 저장, key trace, app state를 검증한다.
- completed: 기존 incident E2E fixture는 새 tutorial 추가 후에도 incident부터 시작한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/content/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_quote_pair_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
