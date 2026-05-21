# PLAYLIST-ORDER-001 — Explicit Playlist Ordering

Slice-ID: PLAYLIST-ORDER-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/playlist-order-001-explicit-order.md
- docs/exec-plans/completed/playlist-order-001-explicit-order.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- content/playlists/
- internal/content/
- internal/playable/
- test/e2e/

## 목표

Incident Run을 `incident-*` 별도 카테고리로 추가할 수 있도록 playable playlist 정렬을 ID 사전순이 아니라 명시적 `category` + `order` 기준으로 바꾼다.

## 범위

- 포함:
  - playlist YAML의 `category`, `order` 필드
  - loader schema와 정렬 로직
  - active tutorial playlist의 명시 order
  - content/playable test 갱신
- 제외:
  - playlist ID 변경
  - exercise/scenario/progress mission ID 변경
  - Incident content 추가

## 수용 기준

- completed: approved/implemented playlist는 `category`와 `order`를 가진다.
- completed: playable playlist는 `category` rank, `order`, `id` 순서로 정렬된다.
- completed: tutorial category는 incident category보다 먼저 실행된다.
- completed: 기존 tutorial 실행 순서는 유지된다.
- completed: `tutorial-90-search-basic` ID는 유지하되 `order: 10`으로 정렬 의미를 표현한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/content/... ./internal/playable/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
