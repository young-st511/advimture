# PLAYPACK-008 — Search Basic Tutorial

Slice-ID: PLAYPACK-008
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/playpack-008-search-basic.md
- docs/exec-plans/completed/playpack-008-search-basic.md
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

`search-basic`을 playable tutorial로 승격한다. 플레이어는 `/`, `n`, `N`으로 로그와 설정에서 명시된 literal token을 찾는 흐름을 익힌다.

## 범위

- 포함:
  - `content/command_clusters/search-basic.yaml`
  - `content/exercises/search-basic.yaml` 4문항
  - `content/scenarios/search-basic.yaml`
  - `content/playlists/search-basic.yaml`
  - content loader count, playable ID/order, coverage 테스트 갱신
  - full playlist E2E
- 제외:
  - `?` backward search
  - regex/highlight/search history
  - progress 저장 포맷 변경

## 수용 기준

- completed: `search-basic` command cluster는 approved + implemented content로 로드된다.
- completed: `/`, `n`, `N`이 coverage를 모두 만족한다.
- completed: 모든 exercise는 `replay_status: pass`이고 E2E assertion을 가진다.
- completed: playlist는 8문항 이하이며 `tutorial-9-repeat-last-change` 다음에 실행된다.
- completed: full E2E는 전체 playlist 완료, progress 저장, key trace, final app state를 검증한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/content/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_search_basic_full.yaml`

최종 전체 검증은 PLATFORM-RFC-001 완료 후 함께 수행한다.
