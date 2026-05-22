# INCIDENT-RUN-004 — Linewise Block Recovery Incident

Slice-ID: INCIDENT-RUN-004
Created: 2026-05-23
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/incident-run-004-linewise-block-recovery.md
- docs/exec-plans/completed/incident-run-004-linewise-block-recovery.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- content/exercises/incident-linewise.yaml
- content/scenarios/incident-linewise.yaml
- content/playlists/incident-linewise.yaml
- internal/content/loader_test.go
- test/e2e/playable_incident_004_full.yaml
- Makefile

## 목표

linewise visual을 tutorial 전용 기능이 아니라 Runbook Dispatch incident에서 config block 복구 도구로 적용한다.

## 수용 기준

- 새 Vim engine 기능을 추가하지 않는다.
- incident는 5 beat 이하로 유지한다.
- 최소 3개 beat에서 linewise `V`를 실제 복구 동작으로 사용한다.
- search 또는 substitute 같은 기존 command를 linewise visual과 함께 조합한다.
- full incident E2E가 progress/app_state/key trace를 검증한다.

## 제외 항목

- 새 command cluster
- multi-line charwise visual
- visual block
- progress schema 변경
- incident 001~003 wording 변경

## 검증 계획

- `go test ./internal/content/...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_004_full.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 완료 내용

- `incident-004-linewise-block-recovery` playlist를 추가했다.
- 5개 beat는 `/block` 검색, linewise `Vd`, linewise `Vy+p`, linewise `VGd`, `:%s`로 구성한다.
- 새 Vim engine 기능이나 progress 저장 포맷 변경은 추가하지 않았다.

## 검증 결과

- passed: `go test ./internal/content/...`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_004_full.yaml`
- passed: `go test ./...`
- passed: `make e2e-playable`
