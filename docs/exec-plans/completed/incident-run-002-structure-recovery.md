# INCIDENT-RUN-002 — Structure Recovery Incident

Slice-ID: INCIDENT-RUN-002
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/incident-run-002-structure-recovery.md
- docs/exec-plans/completed/incident-run-002-structure-recovery.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- content/exercises/incident-structure.yaml
- content/scenarios/incident-structure.yaml
- content/playlists/incident-structure.yaml
- internal/content/loader_test.go
- test/e2e/playable_incident_001_full.yaml
- test/e2e/playable_incident_002_full.yaml
- Makefile

## 목표

새 engine 기능 없이 search, quote text object, substitute, repeat을 섞은 두 번째 incident run을 만든다.

## 범위

- 포함:
  - 5문항 incident exercise
  - scenario skin
  - `incident-002-structure-recovery` playlist
  - full incident E2E
  - existing incident transition assertion 갱신
- 제외:
  - 새 command cluster
  - 새 Vim engine 기능
  - progress 저장 포맷 변경
  - world lore 장문 확장

## 수용 기준

- completed: incident playlist는 `category: incident`, `order: 2`로 incident 001 뒤에 실행된다.
- completed: 흐름은 search -> `ci"` -> `yi"`/`P` -> `:%s` -> `ci"`/`.`를 다룬다.
- completed: full E2E는 progress 저장과 app state를 검증한다.
- completed: incident 001 E2E는 다음 incident 존재를 반영한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/content/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
