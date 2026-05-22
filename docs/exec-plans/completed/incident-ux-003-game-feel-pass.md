# INCIDENT-UX-003 — Incident Game-Feel Pass

Slice-ID: INCIDENT-UX-003
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/incident-ux-003-game-feel-pass.md
- docs/exec-plans/completed/incident-ux-003-game-feel-pass.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- content/exercises/incident-run.yaml
- content/exercises/incident-structure.yaml
- content/scenarios/incident-run.yaml
- content/scenarios/incident-structure.yaml
- content/playlists/incident-run.yaml
- content/playlists/incident-structure.yaml
- test/e2e/playable_incident_001_full.yaml
- test/e2e/playable_incident_002_full.yaml

## 목표

Incident 001/002를 “문제 세트”가 아니라 `원격 시설 복구국 / Runbook Dispatch` 세계관의 짧은 복구 작전처럼 느끼게 한다.

## 완료 내용

- `incident-001-hotfix` playlist title을 “릴레이 기지 001: 야간 핫픽스 복구”로 리프레이밍했다.
- `incident-002-structure-recovery` playlist title을 “릴레이 기지 002: 구조 설정 재동기화”로 리프레이밍했다.
- incident 001/002의 scenario briefing, mission title, success/failure feedback을 복구 작전 톤에 맞췄다.
- incident 001/002의 모든 exercise hint를 2단계 이상으로 보강했다.
- focused E2E의 wait/assert 문구를 갱신했다.

## 제외한 것

- target_state 변경
- optimal_keys 변경
- allowed/forbidden/constraints 변경
- 새 Vim engine 기능
- progress 저장 포맷 변경
- visual mode 구현

## 검증 결과

- passed: `go test ./internal/content/...`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`

