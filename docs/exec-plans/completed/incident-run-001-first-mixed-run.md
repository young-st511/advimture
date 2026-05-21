# INCIDENT-RUN-001 — First Mixed Incident Run

Slice-ID: INCIDENT-RUN-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/incident-run-001-first-mixed-run.md
- docs/exec-plans/completed/incident-run-001-first-mixed-run.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- content/exercises/
- content/scenarios/
- content/playlists/
- internal/content/
- test/e2e/

## 목표

새 Vim engine 기능을 추가하지 않고 기존 command를 조합해 첫 생존 어드벤처 적용 런 `incident-001`을 만든다.

주제: 장애 로그에서 원인을 찾고 설정을 핫픽스한다.

## 범위

- 포함:
  - 6문항 incident exercise
  - scenario skin
  - `incident-001` playlist
  - content count/order/coverage test 갱신
  - full incident E2E
- 제외:
  - 새 command cluster
  - 새 Vim engine 기능
  - progress 저장 포맷 변경
  - Practice/Challenge schema 분리

## 수용 기준

- completed: incident playlist는 `category: incident`, `order: 1`로 tutorial 뒤에 실행된다.
- completed: 각 exercise는 기존 implemented command cluster 중 하나에 연결된다.
- completed: 흐름은 `/error` -> `/timeout` + `n` -> `ciw` -> `o` -> `yy/p` -> `:2,3s`를 다룬다.
- completed: Incident constraints는 tutorial보다 약간 완화하되 핵심 required key는 유지한다.
- completed: full incident E2E는 progress 저장과 app state를 검증한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/content/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_search_basic_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
