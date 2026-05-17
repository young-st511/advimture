# ExecPlan: E2E runner assertion hardening

Slice-ID: E2E-002
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- cmd/e2e-runner/**
- test/e2e/**
- docs/verification/**
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/e2e-002-runner-assertion-hardening.md

## 목표

첫 playable slice를 실제 TUI에 연결하기 전에 E2E runner의 assertion과 evidence를 강화한다. 이번 slice는 app wiring이 아니라 runner 자체의 신뢰도를 높이는 작업이다.

## 범위

- 포함: summary JSON evidence
- 포함: unsafe HOME guard
- 포함: key trace assertion
- 포함: progress file content assertion
- 포함: smoke scenario에 key trace assertion 추가
- 제외: 새 Bubble Tea flow 구현
- 제외: Vim buffer/cursor/mode app assertion 활성화
- 제외: CI integration

## 구현 계획

1. `evidence.summary.json`을 항상 남길 수 있게 한다.
2. `setup.home`이 실제 사용자 HOME을 가리키면 기본적으로 거부한다.
3. `assert.key_trace`로 전송된 키 trace를 검증한다.
4. `assert.progress_file_contains`로 progress JSON 텍스트 포함 여부를 검증한다.
5. verification docs에 runner assertion capability를 갱신한다.

## E2E Stop Rule

첫 playable app wiring에서 buffer/cursor/mode 성공 여부를 E2E로 신뢰할 수 없으면 구현을 멈추고, 앱이 테스트용 state summary를 노출하도록 E2E assertion을 먼저 보강한다.

## 검증 계획

- `go test ./cmd/e2e-runner`
- `go test ./...`
- `make e2e-smoke`

## 승인 체크

- [x] summary JSON evidence가 성공/실패 상태와 핵심 관측값을 기록한다.
- [x] 실제 HOME 사용이 기본적으로 차단된다.
- [x] key trace assertion이 동작한다.
- [x] progress file content assertion이 동작한다.
- [x] smoke scenario가 강화된 assertion을 사용한다.
- [x] 전체 테스트와 smoke E2E가 통과한다.

## 의사결정 로그

- 2026-05-18: `summary.json`은 evidence flag와 무관하게 항상 기록한다. 실패 루프에서 Agent가 가장 먼저 읽을 파일이기 때문이다.
- 2026-05-18: app state summary 기반 buffer/cursor/mode assertion은 이번 slice에서 구현하지 않았다. 아직 새 runtime이 Bubble Tea app에 연결되지 않았으므로, 실제 playable wiring slice에서 테스트 전용 state export와 함께 추가한다.
