# INCIDENT-RUN-003 — Visual Recovery Incident

Slice-ID: INCIDENT-RUN-003
Created: 2026-05-22
Status: completed
Scope-Mode: content-and-e2e
Allowed-Paths:
- docs/exec-plans/active/incident-run-003-visual-recovery.md
- docs/exec-plans/completed/incident-run-003-visual-recovery.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- content/exercises/incident-visual.yaml
- content/scenarios/incident-visual.yaml
- content/playlists/incident-visual.yaml
- test/e2e/playable_incident_003_full.yaml
- internal/content/loader_test.go
- Makefile

## 목표

visual selection을 tutorial 전용 기능이 아니라 Runbook Dispatch incident에서 쓰는 적용 도구로 승격한다.

## 수용 기준

- 새 Vim engine 기능을 추가하지 않는다.
- incident는 5~6 beat 이하로 유지한다.
- 최소 3개 beat에서 visual selection을 실제 복구 동작으로 사용한다.
- search 또는 substitute 같은 기존 command를 visual과 함께 조합한다.
- full incident E2E가 progress/app_state/key trace를 검증한다.

## 제외한 것

- 새 command cluster 추가
- linewise `V`, visual block, multi-line visual 구현
- progress 저장 포맷 변경
- incident 001/002 wording 변경

## 검증 계획

- `go test ./internal/content/...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_003_full.yaml`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 결과

- `incident-003-visual-recovery` playlist를 추가했다.
- 5개 beat는 `/contam` 검색, visual `d`, visual `y` + `p`, backward visual `d`, `:%s`로 구성한다.
- 새 Vim engine 기능이나 progress 저장 포맷 변경은 추가하지 않았다.
