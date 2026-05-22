# INCIDENT-FLOW-001 — Runbook Continuity Pass

Slice-ID: INCIDENT-FLOW-001
Created: 2026-05-22
Status: completed
Scope-Mode: content-wording-and-e2e
Allowed-Paths:
- docs/exec-plans/active/incident-flow-001-runbook-continuity.md
- docs/exec-plans/completed/incident-flow-001-runbook-continuity.md
- docs/gameplay/spec.md
- docs/gameplay/world-frame.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- content/scenarios/incident-run.yaml
- content/scenarios/incident-structure.yaml
- content/scenarios/incident-visual.yaml
- test/e2e/playable_incident_001_full.yaml
- test/e2e/playable_incident_002_full.yaml
- test/e2e/playable_incident_003_full.yaml

## 목표

incident 001~003이 독립 exercise 묶음이 아니라 하나의 Runbook Dispatch 복구 흐름처럼 보이도록 briefing, success, failure 문구를 정렬한다.

## 수용 기준

- target_state, optimal_keys, constraints를 변경하지 않는다.
- progress 저장 포맷을 변경하지 않는다.
- 각 incident의 첫 beat는 진입 조치, 중간 beat는 후속 조치, 마지막 beat는 마감 조치로 읽힌다.
- briefing은 계속 “상황 1문장 + Vim 조작 목표 1문장” 수준으로 유지한다.
- incident focused E2E가 최신 화면 문구와 progress/app_state를 검증한다.

## 제외한 것

- 새 command 또는 engine 기능
- exercise buffer/target 변경
- progress schema 변경
- 별도 narrative screen 또는 cutscene

## 검증 계획

- `go test ./internal/content/...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_003_full.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 결과

- incident 001~003의 briefing/success/failure 문구를 runbook 단계 흐름으로 정렬했다.
- `docs/gameplay/world-frame.md`에 incident 003 reframe을 추가했다.
- target_state, optimal_keys, constraints, progress 저장 포맷은 변경하지 않았다.
